package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"cse5234/order-processing/pkg/config"
	"cse5234/order-processing/pkg/mysql"
)

func main() {
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

	// create a context for database operations
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// execute SQL queries
	service := mysql.NewService(db)
	service.UpdateOrderStatus(ctx)

	log.Println("order processing complete...")

}
