package validator

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/mail"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/gabriel-vasile/mimetype"

	"github.com/SeyramWood/bookibus/app/framework/database"
)

func Validate(i any, rs ...*http.Request) any {
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)

	if t.Kind() == reflect.Ptr {
		t = t.Elem() // Gets the type in the type pointer
	}
	if v.Kind() == reflect.Ptr {
		v = v.Elem() // Get the value in the value address
	}
	if t.Kind() != reflect.Struct {
		log.Panicln("Please provide a struct type")
	}

	if rs != nil {
		r := rs[0]
		if strings.ToLower(strings.SplitN(r.Header.Get("Content-Type"), ";", 2)[0]) == "multipart/form-data" {
			wg := &sync.WaitGroup{}
			for i := 0; i < t.NumField(); i++ {
				wg.Add(1)
				go func(i int) {
					defer wg.Done()
					if val, ok := t.Field(i).Tag.Lookup("json"); ok {
						fieldType := t.Field(i).Type.String()
						if multipartTypes(fieldType) {
							switch fieldType {
							case "[]*multipart.FileHeader":
								fh := r.MultipartForm.File[val]
								v.Field(i).Set(reflect.ValueOf(fh))
							default:
								_, fh, _ := r.FormFile(val)
								v.Field(i).Set(reflect.ValueOf(fh))
							}
						} else if floatTypes(fieldType) {
							switch fieldType {
							case "float32":
								fv, _ := strconv.ParseFloat(r.FormValue(val), 32)
								v.Field(i).Set(reflect.ValueOf(float32(fv)))
							default:
								fv, _ := strconv.ParseFloat(r.FormValue(val), 64)
								v.Field(i).Set(reflect.ValueOf(fv))
							}
						} else if intTypes(fieldType) {
							switch fieldType {
							case "int":
								iv, _ := strconv.Atoi(r.FormValue(val))
								v.Field(i).Set(reflect.ValueOf(iv))
							case "int8":
								iv, _ := strconv.ParseInt(r.FormValue(val), 10, 8)
								v.Field(i).Set(reflect.ValueOf(int8(iv)))
							case "int16":
								iv, _ := strconv.ParseInt(r.FormValue(val), 10, 16)
								v.Field(i).Set(reflect.ValueOf(int16(iv)))
							case "int32":
								iv, _ := strconv.ParseInt(r.FormValue(val), 10, 32)
								v.Field(i).Set(reflect.ValueOf(int32(iv)))
							case "uint":
								iv, _ := strconv.Atoi(r.FormValue(val))
								v.Field(i).Set(reflect.ValueOf(uint(iv)))
							case "uint8":
								iv, _ := strconv.ParseInt(r.FormValue(val), 10, 8)
								v.Field(i).Set(reflect.ValueOf(uint8(iv)))
							case "uint16":
								iv, _ := strconv.ParseInt(r.FormValue(val), 10, 16)
								v.Field(i).Set(reflect.ValueOf(uint16(iv)))
							case "uint32":
								iv, _ := strconv.ParseInt(r.FormValue(val), 10, 32)
								v.Field(i).Set(reflect.ValueOf(uint32(iv)))
							case "uint64":
								iv, _ := strconv.ParseInt(r.FormValue(val), 10, 64)
								v.Field(i).Set(reflect.ValueOf(uint64(iv)))
							default:
								iv, _ := strconv.ParseInt(r.FormValue(val), 10, 64)
								v.Field(i).Set(reflect.ValueOf(iv))
							}
						} else {
							v.Field(i).Set(reflect.ValueOf(r.FormValue(val)))
						}

					}
				}(i)
			}
			wg.Wait()
		}
	}

	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}

	var structure = make(map[string]any)
	var errors = make([]string, 0, t.NumField())

	for i := 0; i < t.NumField(); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			if _, ok := t.Field(i).Tag.Lookup("json"); ok {
				msg := validator(i, t, v)
				mut.Lock()
				if msg == "" {
					errors = append(errors, t.Field(i).Tag.Get("json"))
				}
				structure[t.Field(i).Tag.Get("json")] = msg
				mut.Unlock()
			}
		}(i)
	}
	wg.Wait()

	if len(errors) == t.NumField() {
		structure = make(map[string]any)
		errors = []string{}
		return nil
	}

	return structure

}

func validator(index int, t reflect.Type, v reflect.Value) any {
	rules := strings.Split(t.Field(index).Tag.Get("validate"), "|")
	field := t.Field(index).Name
	fieldType := t.Field(index).Type.String()
	value := v.Field(index)
	formattedField := formatFieldName(t.Field(index).Tag.Get("json")) // t.Field(index).Tag.Get("json")
	for _, rule := range rules {
		if rule == "required" {
			if (value.Kind().String() == "string" || value.Kind().String() == "slice") && value.Len() == 0 {
				return fmt.Sprintf("The %s field is required", formattedField)
			} else if fieldType == "bool" && !value.Bool() {
				return fmt.Sprintf("The %s field must be true", formattedField)
			} else if numberTypes(fieldType) && value.IsZero() {
				return fmt.Sprintf("The %s field is required", formattedField)
			} else if multipartTypes(fieldType) && value.IsNil() {
				return fmt.Sprintf("The %s field is required", formattedField)
			}
		}

		if fieldType == "string" && value.Len() != 0 {
			switch rule {
			case "string":
				r, _ := regexp.Compile("^[0-9a-zA-Z-+ .]+$")
				if !r.MatchString(value.String()) {
					return fmt.Sprintf("The %s must be a string.", formattedField)
				}
			case "ascii":
				r, _ := regexp.Compile("^[\\x00-\\x7F]*$")
				if !r.MatchString(value.String()) {
					return fmt.Sprintf("The %s must be a string.", formattedField)
				}
			case "email":
				if _, err := mail.ParseAddress(value.String()); err != nil {
					return fmt.Sprintf("The %s must be a valid email address.", formattedField)
				}
			case "phone_with_code":
				if phone, _ := regexp.Compile(`^\+\d{12}$`); !phone.MatchString(value.String()) {
					return fmt.Sprintf("The %s field must be a valid phone number with county code.", formattedField)
				}
			case "phone":
				if phone, _ := regexp.Compile(`^0\d{9}$`); !phone.MatchString(value.String()) {
					return fmt.Sprintf("The %s field must be a valid phone number.", formattedField)
				}
			case "email_phone":
				phone, _ := regexp.Compile(`^0\d{9}$`)
				if strings.Contains(value.String(), "@") {
					if _, err := mail.ParseAddress(value.String()); err != nil {
						return fmt.Sprintf("The %s must be a valid email address.", formattedField)
					}
				} else {
					if !phone.MatchString(value.String()) {
						return fmt.Sprintf("The %s field must be a valid phone number.", formattedField)
					}
				}
			case "id_card":
				r, _ := regexp.Compile(`^GHA-\d{9}-\d{1}$`)
				if !r.MatchString(value.String()) {
					return "The ID field must be a valid Ghana Card."
				}
			case "digital_address":
				r, _ := regexp.Compile(`[A-Z]{2}-\d{1,4}-\d{4}$`)
				if !r.MatchString(value.String()) {
					return "The address field must be a valid digital address."
				}
			default:
				if strings.Contains(rule, ":") {
					err := validateSecondaryStringRules(rule, formattedField, field, t, value, v)
					if err != "" {
						return err
					}
				}
			}
		}
		if numberTypes(fieldType) && !value.IsZero() {
			if strings.Contains(rule, ":") {
				err := validateSecondaryNumberRules(rule, formattedField, fieldType, value)
				if err != "" {
					return err
				}
			}
		}
		if multipartTypes(fieldType) && !value.IsNil() {
			if rule == "image" {
				if fieldType == "*multipart.FileHeader" {
					fh := value.Interface().(*multipart.FileHeader)
					buffer, _, err := checkMIMES(fh, formattedField)
					if err != "" {
						return err
					}
					if !mimetype.EqualsAny(mimetype.Detect(buffer).Extension(), ".jpg", ".jpeg", ".png") {
						return fmt.Sprintf("The %s must be a file of type: .jpg, .jpeg, .png", formattedField)
					}
				}
				if fieldType == "[]*multipart.FileHeader" {
					for _, fh := range value.Interface().([]*multipart.FileHeader) {
						buffer, _, err := checkMIMES(fh, formattedField)
						if err != "" {
							return err
						}
						if !mimetype.EqualsAny(mimetype.Detect(buffer).Extension(), ".jpg", ".jpeg", ".png") {
							return fmt.Sprintf("The %s must be a file of type: .jpg, .jpeg, .png", formattedField)
						}
					}
				}
			}
			if strings.Contains(rule, ":") {
				err := validateSecondaryMimeRules(rule, formattedField, fieldType, value)
				if err != "" {
					return err
				}
			}
		}
		if value.Kind().String() == "slice" {
			if strings.Contains(rule, ":") {
				err := validateSecondarySliceRules(rule, formattedField, fieldType, value)
				if err != "" {
					return err
				}
			}
		}

	}
	return ""
}

func validateSecondaryStringRules(rule, formattedField, dbField string, t reflect.Type, value, v reflect.Value) any {

	r := strings.Split(rule, ":")
	switch r[0] {
	case "max":
		val, _ := strconv.Atoi(r[1])
		if value.Len() > val {
			return fmt.Sprintf("The %s must not be greater than %v characters", formattedField, val)
		}
	case "min":
		val, _ := strconv.Atoi(r[1])
		if value.Len() < val {
			return fmt.Sprintf("The %s must be at least %v characters", formattedField, val)
		}
	case "match":
		var val reflect.Value
		for i := 0; i < t.NumField(); i++ {
			if t.Field(i).Name != r[1] {
				continue
			}
			val = v.Field(i)
			break
		}
		if strings.TrimSpace(value.String()) != strings.TrimSpace(val.String()) {
			return fmt.Sprintf("The %s does not matched", formatFieldName(r[1]))
		}
	case "unique":
		table := r[1]
		if !strings.Contains(table, ".") {
			if r := isUsernameExist(value.String(), dbField, table, formattedField); r != nil {
				return r
			}
		}
		if tc := strings.SplitN(table, ".", 2); len(tc) == 2 {
			if r := isValueExist(value.String(), tc[1], tc[0], formattedField); r != nil {
				return r
			}
		}

	}

	return ""
}
func validateSecondaryNumberRules(
	rule, formattedField, fieldType string, value reflect.Value,
) any {
	r := strings.Split(rule, ":")
	switch r[0] {
	case "max":
		if intTypes(fieldType) {
			val, _ := strconv.ParseInt(r[1], 10, 64)
			if !(value.Int() <= val) {
				return fmt.Sprintf("The %s must be less or equal to %v", formattedField, val)
			}
		}
		if floatTypes(fieldType) {
			val, _ := strconv.ParseFloat(r[1], 64)
			if !(value.Float() <= val) {
				return fmt.Sprintf("The %s must be less or equal to %v", formattedField, val)
			}
		}
	case "min":
		if intTypes(fieldType) {
			val, _ := strconv.ParseInt(r[1], 10, 64)
			if !(value.Int() >= val) {
				return fmt.Sprintf("The %s must be greater or equal to %v", formattedField, val)
			}
		}
		if floatTypes(fieldType) {
			val, _ := strconv.ParseFloat(r[1], 64)
			if !(value.Float() >= val) {
				return fmt.Sprintf("The %s must be greater or equal to %v", formattedField, val)
			}
		}
	case "equal":
		if intTypes(fieldType) {
			val, _ := strconv.ParseInt(r[1], 10, 64)
			if value.Int() != val {
				return fmt.Sprintf("The %s must be equal to %v", formattedField, val)
			}
		}
		if floatTypes(fieldType) {
			val, _ := strconv.ParseFloat(r[1], 64)
			if value.Float() != val {
				return fmt.Sprintf("The %s must be equal to %v", formattedField, val)
			}
		}

	}
	return ""
}

func validateSecondaryMimeRules(
	rule, formattedField, fieldType string, value reflect.Value,
) any {
	r := strings.Split(rule, ":")
	switch r[0] {
	case "image", "file", "mimes":
		if fieldType == "*multipart.FileHeader" {
			fh := value.Interface().(*multipart.FileHeader)
			buffer, _, err := checkMIMES(fh, formattedField)
			if err != "" {
				return err
			}
			if !mimetype.EqualsAny(mimetype.Detect(buffer).Extension(), prepareMIMES(r[1])...) {
				return fmt.Sprintf(
					"The %s must be a file of type: %v", formattedField,
					strings.TrimSpace(strings.Join(prepareMIMES(r[1]), ", ")),
				)
			}
		}
		if fieldType == "[]*multipart.FileHeader" {
			for _, fh := range value.Interface().([]*multipart.FileHeader) {
				buffer, _, err := checkMIMES(fh, formattedField)
				if err != "" {
					return err
				}
				if !mimetype.EqualsAny(mimetype.Detect(buffer).Extension(), prepareMIMES(r[1])...) {
					return fmt.Sprintf(
						"The %s must be a file of type: %v", formattedField,
						strings.TrimSpace(strings.Join(prepareMIMES(r[1]), ", ")),
					)
				}
			}
		}
	case "size":
		sValue := strings.Split(r[1], "")
		var rs int
		vsymbol := strings.ToLower(strings.Join(sValue[len(sValue)-2:], ""))
		if len(sValue) == 4 {
			if crs, err := strconv.Atoi(strings.Join(sValue[:len(sValue)-2], "")); err == nil {
				rs = crs
			}
		} else {
			if crs, err := strconv.Atoi(strings.Join(sValue[:len(sValue)-2], "")); err == nil {
				rs = crs
			}
		}
		if fieldType == "*multipart.FileHeader" {
			fh := value.Interface().(*multipart.FileHeader)
			switch vsymbol {
			case "mb":
				if fh.Size > int64(1024*1024*rs) {
					if rs > 1 {
						return fmt.Sprintf("The %s must be %d megabytes", formattedField, rs)
					}
					return fmt.Sprintf("The %s must be %d megabyte", formattedField, rs)
				}
			default:
				return fmt.Sprintf("The rule %s is not supported, try this: %dMB", r[1], rs)
			}
		}
		if fieldType == "[]*multipart.FileHeader" {
			switch vsymbol {
			case "mb":
				for _, fh := range value.Interface().([]*multipart.FileHeader) {
					if fh.Size > int64(1024*1024*rs) {
						if rs > 1 {
							return fmt.Sprintf("The %s must be %d megabytes", formattedField, rs)
						}
						return fmt.Sprintf("The %s must be %d megabyte", formattedField, rs)
					}
				}
			default:
				return fmt.Sprintf("The rule %s is not supported, try this: %dMB", r[1], rs)
			}
		}
	case "min":
		if fieldType == "[]*multipart.FileHeader" {
			if fh, ok := value.Interface().([]*multipart.FileHeader); ok {
				minVal, err := strconv.Atoi(r[1])
				if err != nil {
					log.Panicln(err)
				}
				if len(fh) < minVal {
					return fmt.Sprintf("Only %v or more upload(s) allowed", minVal)
				}
			}
		}
	case "max":
		if fieldType == "[]*multipart.FileHeader" {
			if fh, ok := value.Interface().([]*multipart.FileHeader); ok {
				minVal, err := strconv.Atoi(r[1])
				if err != nil {
					log.Panicln(err)
				}
				if len(fh) > minVal {
					return fmt.Sprintf("Only %v upload(s) allowed", minVal)
				}
			}
		}
	}
	return ""
}

func validateSecondarySliceRules(
	rule, formattedField, fieldType string, value reflect.Value,
) any {
	r := strings.Split(rule, ":")
	switch r[0] {
	case "min":
		minVal, err := strconv.Atoi(r[1])
		if err != nil {
			log.Panicln(err)
		}
		if value.Len() < minVal {
			return fmt.Sprintf("The %s must have at least %v items", formattedField, minVal)
		}
	case "max":
		minVal, err := strconv.Atoi(r[1])
		if err != nil {
			log.Panicln(err)
		}
		if value.Len() > minVal {
			return fmt.Sprintf("The %s may not have more than %v items", formattedField, minVal)
		}
	}
	return ""
}

func formatFieldName(field string) string {
	var text string
	for i := 0; i < len(field); i++ {
		c := string([]byte{field[i]})
		if c == strings.ToUpper(c) {
			if len(text) != 0 {
				text += " "
			}
			text += strings.ToLower(c)
		} else {
			text += strings.ToLower(c)
		}
	}
	return text

}

func snakeCase(field string) string {
	var text string
	for i := 0; i < len(field); i++ {
		c := string([]byte{field[i]})
		if c == strings.ToUpper(c) {
			if len(text) != 0 {
				text += "_"
			}
			text += strings.ToLower(c)
		} else {
			text += strings.ToLower(c)
		}
	}
	return text
}
func camelCase(field string) string {
	var text string
	for i := 0; i < len(field); i++ {
		c := string([]byte{field[i]})
		if c == strings.ToUpper(c) {
			if len(text) != 0 {
				text += "_"
			}
			text += strings.ToLower(c)
		} else {
			text += strings.ToLower(c)
		}
	}
	return text

}
func numberTypes(numberType string) bool {
	nt := []string{"int", "int8", "int16", "int32", "int64", "float32", "float64"}
	for _, s := range nt {
		if numberType != s {
			continue
		}
		return true
	}
	return false
}
func intTypes(numberType string) bool {
	nt := []string{"int", "int8", "int16", "int32", "int64"}
	for _, s := range nt {
		if numberType != s {
			continue
		}
		return true
	}
	return false
}
func floatTypes(numberType string) bool {
	nt := []string{"float32", "float64"}
	for _, s := range nt {
		if numberType != s {
			continue
		}
		return true
	}
	return false
}

func multipartTypes(multipartType string) bool {
	mt := []string{"*multipart.FileHeader", "[]*multipart.FileHeader"}
	for _, s := range mt {
		if multipartType != s {
			continue
		}
		return true
	}
	return false
}
func prepareMIMES(mimes string) []string {
	buffer := make([]string, 0, len(mimes))
	for _, m := range strings.Split(mimes, ",") {
		buffer = append(buffer, fmt.Sprintf(".%s", m))
	}
	return buffer
}
func checkMIMES(fh *multipart.FileHeader, formattedField string) ([]byte, int64, string) {
	file, err := fh.Open()
	if err != nil {
		return nil, 0, fmt.Sprintf("The %s must be a file", formattedField)
	}
	buffer, err := io.ReadAll(file)
	if err != nil {
		return nil, 0, fmt.Sprintf("The %s must be a file", formattedField)
	}
	return buffer, 0, ""
}

func isUsernameExist(username, field, table, formattedField string) any {
	dbField := snakeCase(field)
	queryStr := fmt.Sprintf("SELECT %s FROM %s WHERE %s=?", dbField, table, dbField)
	db := database.Connect()
	err := db.QueryRow(queryStr, username).Scan(&username)
	if err == nil {
		return fmt.Sprintf("The %s is already taken.", formattedField)
	}
	defer db.Close()
	return nil
}
func isValueExist(value, field, table, formattedField string) any {
	dbField := snakeCase(field)
	queryStr := fmt.Sprintf("SELECT %s FROM %s WHERE %s=?", dbField, table, dbField)
	db := database.Connect()
	err := db.QueryRow(queryStr, value).Scan(&value)
	if err == nil {
		return fmt.Sprintf("The %s is already taken.", formattedField)
	}
	defer db.Close()
	return nil
}
