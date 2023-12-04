package config

import (
	"os"

	"github.com/SeyramWood/bookibus/utils/env"
)

type (
	app struct {
		Name             string
		Version          string
		AppURL           string
		AppWebsiteURL    string
		AppCompanyURL    string
		AppBookiBusURL   string
		Key              string
		TokenName        string
		PORT             string
		FilesystemDriver string
		Redis            string
		RabbitMQ         string
	}
	serverConfig struct {
		Prefork           bool
		CaseSensitive     bool
		StrictRouting     bool
		StreamRequestBody bool
		EnablePrintRoutes bool
		Concurrency       int64
		ServerHeader      string
		AppName           string
	}
	sms struct {
		Sender  string
		Gateway string
	}

	mailer struct {
		Mailer      string
		Host        string
		Port        string
		Username    string
		Password    string
		Encryption  string
		FromAddress string
		FromName    string
	}
	payment struct {
		Gateway string
	}
	paystack struct {
		URL    string
		PubKey string
		SecKey string
		Email  string
		Domain string
	}

	arkesel struct {
		APIKey string
		URL    string
	}
	contact struct {
		Email string
		Phone string
	}
)

func App() *app {
	if os.Getenv("APP_ENV") == "production" {
		return &app{
			Name:             os.Getenv("APP_NAME"),
			Version:          os.Getenv("APP_VERSION"),
			AppURL:           os.Getenv("APP_URL"),
			AppWebsiteURL:    os.Getenv("APP_WEBSITE_URL"),
			AppCompanyURL:    os.Getenv("APP_COMPANY_URL"),
			AppBookiBusURL:   os.Getenv("APP_BOOKIBUS_URL"),
			Key:              os.Getenv("APP_KEY"),
			TokenName:        os.Getenv("API_TOKEN_NAME"),
			PORT:             os.Getenv("SERVER_PORT"),
			FilesystemDriver: os.Getenv("FILESYSTEM_DRIVER"),
			RabbitMQ:         os.Getenv("RabbitMQ_HOST"),
			Redis:            os.Getenv("REDIS_HOST"),
		}
	}
	return &app{
		Name:             env.Get("APP_NAME", "My First API"),
		Version:          env.Get("APP_VERSION", "0.0.1"),
		AppURL:           env.Get("APP_URL", "http://127.0.0.1:80"),
		AppWebsiteURL:    env.Get("APP_WEBSITE_URL", ""),
		AppCompanyURL:    env.Get("APP_COMPANY_URL", ""),
		AppBookiBusURL:   env.Get("APP_BOOKIBUS_URL", ""),
		Key:              env.Get("APP_KEY", "secretKEY5465"),
		TokenName:        env.Get("API_TOKEN_NAME", "remember"),
		PORT:             env.Get("SERVER_PORT", "80"),
		FilesystemDriver: env.Get("FILESYSTEM_DRIVER", "local"),
		RabbitMQ:         env.Get("RabbitMQ_HOST", "rabbitmq"),
		Redis:            env.Get("REDIS_HOST", "redis"),
	}
}
func Server() *serverConfig {
	return &serverConfig{
		Prefork:           true,
		CaseSensitive:     true,
		StrictRouting:     true,
		StreamRequestBody: true,
		EnablePrintRoutes: true,
		Concurrency:       256 * 2048,
		ServerHeader:      "Bookibus Inc.",
		AppName:           env.Get("APP_NAME", "Bookihub Inc."),
	}
}

func Mailer() *mailer {
	if os.Getenv("APP_ENV") == "production" {
		return &mailer{
			Mailer:      os.Getenv("MAIL_MAILER"),
			Host:        os.Getenv("MAIL_HOST"),
			Port:        os.Getenv("MAIL_PORT"),
			Username:    os.Getenv("MAIL_USERNAME"),
			Password:    os.Getenv("MAIL_PASSWORD"),
			Encryption:  os.Getenv("MAIL_ENCRYPTION"),
			FromAddress: os.Getenv("MAIL_FROM_ADDRESS"),
			FromName:    os.Getenv("MAIL_FROM_NAME"),
		}
	}
	return &mailer{
		Mailer:      env.Get("MAIL_MAILER", "smtp"),
		Host:        env.Get("MAIL_HOST", ""),
		Port:        env.Get("MAIL_PORT", ""),
		Username:    env.Get("MAIL_USERNAME", ""),
		Password:    env.Get("MAIL_PASSWORD", ""),
		Encryption:  env.Get("MAIL_ENCRYPTION", ""),
		FromAddress: env.Get("MAIL_FROM_ADDRESS", "info@bookihub.com"),
		FromName:    env.Get("MAIL_FROM_NAME", "Bookihub Inc."),
	}
}

func SMS() *sms {
	if os.Getenv("APP_ENV") == "production" {
		return &sms{
			Sender:  os.Getenv("SMS_SENDER"),
			Gateway: os.Getenv("SMS_GATEWAY"),
		}
	}
	return &sms{
		Sender:  env.Get("SMS_SENDER", "Asinyo"),
		Gateway: env.Get("SMS_GATEWAY", "arkesel"),
	}
}
func Arkesel() *arkesel {
	if os.Getenv("APP_ENV") == "production" {
		return &arkesel{
			APIKey: os.Getenv("ARKESEL_API_KEY"),
			URL:    os.Getenv("ARKESEL_URL"),
		}
	}
	return &arkesel{
		APIKey: env.Get("ARKESEL_API_KEY", ""),
		URL:    env.Get("ARKESEL_URL", ""),
	}
}
func Payment() *payment {
	if os.Getenv("APP_ENV") == "production" {
		return &payment{
			Gateway: os.Getenv("PAYMENT_GATEWAY"),
		}
	}
	return &payment{
		Gateway: env.Get("PAYMENT_GATEWAY", "paystack"),
	}
}
func Paystack() *paystack {
	if os.Getenv("APP_ENV") == "production" {
		return &paystack{
			URL:    os.Getenv("PAYSTACK_URL"),
			PubKey: os.Getenv("PAYSTACK_PUB_KEY"),
			SecKey: os.Getenv("PAYSTACK_SEC_KEY"),
			Email:  os.Getenv("PAYSTACK_EMAIL"),
			Domain: os.Getenv("PAYSTACK_DOMAIN"),
		}
	}
	return &paystack{
		URL:    env.Get("PAYSTACK_URL", ""),
		PubKey: env.Get("PAYSTACK_PUB_KEY", ""),
		SecKey: env.Get("PAYSTACK_SEC_KEY", ""),
		Email:  env.Get("PAYSTACK_EMAIL", ""),
		Domain: env.Get("PAYSTACK_DOMAIN", ""),
	}

}
func Contact() *contact {
	if os.Getenv("APP_ENV") == "production" {
		return &contact{
			Email: os.Getenv("APP_EMAIL"),
			Phone: os.Getenv("APP_PHONE"),
		}
	}
	return &contact{
		Email: env.Get("APP_EMAIL", "tech@bookihub.com"),
		Phone: env.Get("APP_PHONE", ""),
	}
}
