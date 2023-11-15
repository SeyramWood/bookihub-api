package database

import (
	"fmt"
	"log"
	"os"

	"github.com/SeyramWood/bookibus/config"
	"github.com/SeyramWood/bookibus/ent"
	"github.com/SeyramWood/bookibus/utils/env"
)

func postgresConnector(dBDriver string) *ent.Client {
	conf := config.DB()
	var DbSSLMode string
	if os.Getenv("APP_ENV") == "production" {
		DbSSLMode = os.Getenv("DB_SSLMODE")
	} else {
		DbSSLMode = env.Get("DB_SSLMODE", "disable")
	}
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		conf.Host,
		conf.Port,
		conf.Username,
		conf.Name,
		conf.Password,
		DbSSLMode,
	)
	client, err := ent.Open(dBDriver, dsn)
	if err != nil {
		log.Fatalf("failed connecting to postgres: %v", err)
	}
	return client
}
