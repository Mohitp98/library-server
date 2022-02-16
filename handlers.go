package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Mohitp98/library-server/models"
	"github.com/Mohitp98/library-server/responses"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var validate = validator.New()
var bookCollection *mongo.Collection = Client.Database(DB_NAME).Collection(COLLECTION_NAME)

func GetAllBooksHandler(rw http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var books []models.Book
	defer cancel()

	results, err := bookCollection.Find(ctx, bson.M{})
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		response := responses.Default{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"details": err.Error()}}
		json.NewEncoder(rw).Encode(response)
		return
	}

	defer results.Close(ctx)
	for results.Next(ctx) {
		var book models.Book
		if err = results.Decode(&book); err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			response := responses.Default{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"details": err.Error()}}
			json.NewEncoder(rw).Encode(response)
		}
		books = append(books, book)
	}

	rw.WriteHeader(http.StatusOK)
	response := responses.Default{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"books": books}}
	json.NewEncoder(rw).Encode(response)
}

func GetBookHandler(rw http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	params := mux.Vars(r)
	bookId := params["book_id"]
	var book models.Book
	defer cancel()

	objID, _ := primitive.ObjectIDFromHex(bookId)
	log.Printf("Fetching record for using id : `%v`", bookId)
	err := bookCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&book)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		response := responses.Default{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"details": err.Error()}}
		json.NewEncoder(rw).Encode(response)
		return
	}

	rw.WriteHeader(http.StatusOK)
	response := responses.Default{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"book": book}}
	json.NewEncoder(rw).Encode(response)
}

func AddBookHandler(rw http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var book models.Book
	defer cancel()

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		response := responses.Default{Status: http.StatusBadRequest, Message: "invalid json", Data: map[string]interface{}{"details": err.Error()}}
		json.NewEncoder(rw).Encode(response)
		return
	}

	if validationError := validate.Struct(&book); validationError != nil {
		rw.WriteHeader(http.StatusBadRequest)
		response := responses.Default{Status: http.StatusBadRequest, Message: "invalid json", Data: map[string]interface{}{"details": validationError.Error()}}
		json.NewEncoder(rw).Encode(response)
		return
	}

	newBook := models.Book{
		ID:       primitive.NewObjectID(),
		Name:     book.Name,
		Author:   book.Author,
		Pages:    book.Pages,
		Price:    book.Price,
		Language: book.Language,
		Domain:   book.Domain,
	}

	result, err := bookCollection.InsertOne(ctx, newBook)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		response := responses.Default{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"details": err.Error()}}
		json.NewEncoder(rw).Encode(response)
		return
	}

	rw.WriteHeader(http.StatusCreated)
	response := responses.Default{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"book": result}}
	json.NewEncoder(rw).Encode(response)
}
