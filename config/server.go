package config

import (
	"encoding/json"
	"log"
	"os"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"go.uber.org/zap"
)

type server struct {
	HTTP   *fiber.App
	WG     *sync.WaitGroup
	Logger *zap.Logger
}

func NewServer() *server {
	var logger *zap.Logger
	var wg = sync.WaitGroup{}
	var err error
	if os.Getenv("APP_ENV") == "production" {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}
	if err != nil {
		panic(err)
	}
	return &server{
		HTTP: fiber.New(
			fiber.Config{
				Prefork:           Server().Prefork,
				CaseSensitive:     Server().CaseSensitive,
				StrictRouting:     Server().StrictRouting,
				ServerHeader:      Server().ServerHeader,
				AppName:           App().Name,
				BodyLimit:         10485760,
				StreamRequestBody: true,
				JSONEncoder:       json.Marshal,
				JSONDecoder:       json.Unmarshal,
				Views:             html.New("./app/framework/http/templates", ".html"),
			},
		),
		WG:     &wg,
		Logger: logger,
	}

}

func (http *server) Run() {
	if os.Getenv("APP_ENV") == "production" {
		port := os.Getenv("PORT")
		if port == "" {
			port = App().PORT
		}
		log.Fatal(http.HTTP.Listen(":" + port))
	} else {
		log.Fatal(http.HTTP.Listen(":" + App().PORT))
	}
}
