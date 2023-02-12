package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"

	"cse5234/order/pkg/config"
	"cse5234/order/pkg/mysql"
	"cse5234/order/pkg/routes"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// get a copy of configuations
	options := config.NewConfig()
	if options == nil {
		panic(fmt.Errorf("invalid configuation"))
	}

	// Set up database connection
	db, err := mysql.Connect(options.Database.User, options.Database.Password, options.Database.Name, options.Database.Address)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// set up Rabbitmq connection
	conn, err := amqp.Dial(options.ShipmentMQ)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// open a Rabbitmq channel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"shipment", // name
		false,      // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	failOnError(err, "Failed to declare a queue")

	// Run migrations
	if err := mysql.Migrate(db); err != nil {
		panic(err)
	}

	// Set up database service
	dbService := mysql.NewService(db)
	if dbService == nil {
		panic(fmt.Errorf("cannot create DB Service"))
	}

	// Create a new handler
	handler := routes.NewHandler(dbService, ch, q.Name)

	serverAddr := fmt.Sprintf(":%d", options.Port)
	log.Printf("Listening on %s", serverAddr)

	// Create an HTTP server
	srv := &http.Server{
		Addr:    serverAddr,
		Handler: handler.Routes(),
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
