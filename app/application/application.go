package application

import (
	"crypto/rand"
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unsafe"

	"github.com/SeyramWood/app/adapters/presenters"
	"github.com/SeyramWood/ent"
)

func Rollback(tx *ent. /*  */ Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}
	return err
}
func OTP(size ...int) string {
	defaultSize := 6
	if size != nil {
		defaultSize = size[0]
	}
	chars := []byte("0123456789")
	b := make([]byte, defaultSize)
	_, _ = rand.Read(b)
	for i := 0; i < defaultSize; i++ {
		b[i] = chars[b[i]%byte(len(chars))]
	}
	return *(*string)(unsafe.Pointer(&b))
}
func RandomString(size int) string {
	chars := []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]byte, size)
	_, _ = rand.Read(b)
	for i := 0; i < size; i++ {
		b[i] = chars[b[i]%byte(len(chars))]
	}
	return *(*string)(unsafe.Pointer(&b))
}
func UsernameType(username, delimiter string) bool {
	if strings.Contains(username, "@") && delimiter == "email" {
		return true
	}
	phone, _ := regexp.Compile(`^0\d{9}$`)
	if phone.MatchString(username) && delimiter == "phone" {
		return true
	}
	return false
}
func Paginate(count int, results any) (*presenters.PaginationResponse, error) {
	return &presenters.PaginationResponse{
		Count: count,
		Data:  results,
	}, nil
}

func FormatSessionID(session any) int {
	if user, ok := session.(map[string]any); ok {
		userID, _ := strconv.Atoi(strconv.FormatFloat(user["id"].(float64), 'G', 'G', 64))
		return userID
	}
	return 0
}

func ConvertStructToMap(s any) map[string]interface{} {
	dataChan := make(chan map[string]interface{})
	go func(s any, dataChan chan map[string]interface{}) {
		v := reflect.ValueOf(s)
		values := make(map[string]interface{}, v.NumField())
		for i := 0; i < v.NumField(); i++ {
			values[v.Type().Field(i).Name] = v.Field(i).Interface()
		}
		dataChan <- values
	}(s, dataChan)
	for {
		select {
		case values := <-dataChan:
			return values
		}
	}
}

func combinations(input []string, prefix []string, index int, result *[][]string) {
	if index == len(input) {
		if len(prefix) > 0 {
			*result = append(*result, append([]string{}, prefix...))
		}
		return
	}

	// Include the current element in the combination
	combinations(input, append(prefix, input[index]), index+1, result)

	// Exclude the current element from the combination
	combinations(input, prefix, index+1, result)
}

// func generateCombinations(input []string) [][]string {
// 	var result [][]string
// 	combinations(input, []string{}, 0, &result)
// 	return result
// }

func FilterCombinations(input []string) [][]string {
	dataChan := make(chan [][]string)
	go func(dataChan chan [][]string) {
		var result [][]string
		combinations(input, []string{}, 0, &result)
		dataChan <- result
	}(dataChan)
	for {
		select {
		case result := <-dataChan:
			return result
		}
	}
}

func ParseRFC3339Datetime(rfc3339Datetime string) time.Time {
	if rfc3339Datetime == "" {
		return time.Now()
	}
	rfc3339Time, err := time.Parse(time.RFC3339, rfc3339Datetime)
	if err != nil {
		log.Panicln("Error parsing RFC3339 datetime:", err)
	}
	return rfc3339Time
}

func CompareFilter(value any) bool {
	switch value.(type) {
	case bool:
		return value.(bool)
	case int:
		if value != 0 {
			return true
		}
	case string:
		if value != "" {
			return true
		}
	}
	return false
}
