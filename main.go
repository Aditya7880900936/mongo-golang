package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Aditya7880900936/mongo-golang/controllers"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client := getClient()
	uc := controllers.NewUserController(client)

	r := httprouter.New()
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)

	fmt.Println("🚀 Server running at http://localhost:8090")
	log.Fatal(http.ListenAndServe("localhost:8090", r))
}

func getClient() *mongo.Client {
	uri := "mongodb://127.0.0.1:27017"
	clientOptions := options.Client().ApplyURI(uri)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("❌ Could not connect to MongoDB:", err)
	}

	fmt.Println("✅ Connected to MongoDB successfully!")
	return client
}
