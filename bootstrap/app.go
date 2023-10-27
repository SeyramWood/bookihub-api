package bootstrap

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/SeyramWood/app/application/cache"
	"github.com/SeyramWood/app/application/storage"
	"github.com/SeyramWood/app/application/worker"
	"github.com/SeyramWood/app/framework/database"
	"github.com/SeyramWood/config"
	"github.com/SeyramWood/ent/migrate"
	"github.com/SeyramWood/utils/env"
	"github.com/SeyramWood/utils/router"
)

func init() {
	env.Setup()
}

func App() {

	db := database.NewDB()

	ctx := context.Background()
	if err := db.DB.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	amqpConn, err := worker.Connect()
	if err != nil {
		log.Fatalf("failed to connect to RabbitMQ: %v", err)
	}
	eventProducer := worker.NewProducer(amqpConn, "events")

	server := config.NewServer()
	server.HTTP.Use(
		cors.New(
			cors.Config{
				AllowCredentials: true,
				AllowOrigins:     "*",
			},
		),
	)

	server.HTTP.Use(recover.New())

	server.HTTP.Use(logger.New())

	storageSrv := storage.NewService(server.WG)
	cacheSrv := cache.NewService()
	router.NewRouter(server.HTTP, db, eventProducer, storageSrv, cacheSrv)

	go server.Run()
	go storageSrv.Listen()

	c := make(chan os.Signal, 1) // Create channel to signify a signal being sent
	signal.Notify(
		c, syscall.SIGINT, syscall.SIGTERM,
	) // When an interrupt or termination signal is sent, notify the channel

	<-c // This blocks the main thread until an interrupt is received
	server.Logger.Info("Gracefully shutting down...")
	server.WG.Wait()
	storageSrv.Close()
	_ = server.HTTP.Shutdown()
	server.Logger.Info("Running cleanup tasks...")
	_ = db.DB.Close()
	amqpConn.Close()
	server.Logger.Info("Application successful shutdown.")

}
