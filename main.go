package main

import (
	"errors"
	"fmt"
	"github.com/luqus/s/api"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/luqus/s/storage"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("invalid port")
	}

	listenAddr := fmt.Sprintf(":%s", port)

	dbString := os.Getenv("MONGO_CONN_STR")
	if dbString == "" {
		log.Fatal(errors.New("MONGO_CONN_STR not found in .env"))
	}

	redisConnStr := os.Getenv("REDIS_CONN_STR")
	if redisConnStr == "" {
		log.Fatal(errors.New("REDIS_CONN_STR  not found in .env"))
	}

	// Database connection
	client, err := storage.Init(dbString)
	if err != nil {
		log.Fatal(err)
	}

	redisClient := storage.RedisInit(redisConnStr, "", 0)

	userCollection := client.Database("stayAlert").Collection("users")
	alertCollection := client.Database("stayAlert").Collection("alert")

	authstore := storage.NewAuthStorage(userCollection)
	alertStore := storage.NewAlertStorage(alertCollection)
	pubsub := storage.NewRedisPubSub(redisClient)
	// Server Instance
	server := api.NewAPIServer(listenAddr, authstore, alertStore, pubsub)

	// Start server
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}

}
