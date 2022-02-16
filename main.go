package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Mohitp98/library-server/middlewares"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	URL_PATH        string = "/api/"
	DB_NAME         string = "library"
	COLLECTION_NAME string = "books"
)

func InitConnection() *mongo.Client {

	MONGO_URI := os.Getenv("MONGO_URI")
	if MONGO_URI == "" {
		log.Fatal("MONGO_URI is not configured!")
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(MONGO_URI))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// check connection with db
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Failed to connect database: ", err)
	}

	log.Println("Connected to Database")
	return client
}

// get database client
var Client = InitConnection()

func main() {
	r := mux.NewRouter()

	// enable middlerware to custom header in response
	r.Use(middlewares.JsonEncoding)

	// registering our routes
	r.HandleFunc("/books", GetAllBooksHandler).Methods(http.MethodGet)
	r.HandleFunc("/books", AddBookHandler).Methods(http.MethodPost)
	r.HandleFunc("/book/{book_id}", GetBookHandler).Methods(http.MethodGet)

	// default server configuration
	srv := &http.Server{
		Handler:      r,
		Addr:         "localhost:5000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Server starting with port no :5000...")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}
}
