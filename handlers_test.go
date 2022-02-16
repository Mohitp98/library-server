package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Mohitp98/library-server/models"
	"github.com/gorilla/mux"
	"gotest.tools/assert"
)

func TestGetAllBooksHandler(t *testing.T) {
	t.Parallel()

	r, _ := http.NewRequest("GET", "/books", nil)
	w := httptest.NewRecorder()

	GetAllBooksHandler(w, r)
	print(w)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetBookHandler(t *testing.T) {
	t.Parallel()

	r, _ := http.NewRequest("GET", "/book/620bc4b8e32f74f8ccd8d98c", nil)
	w := httptest.NewRecorder()

	vars := map[string]string{
		"book_id": "620bc4b8e32f74f8ccd8d98c",
	}

	r = mux.SetURLVars(r, vars)

	GetBookHandler(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestAddBookHandler(t *testing.T) {
	t.Parallel()

	payload := &models.Book{
		Name:     "think like a monk",
		Author:   "jay shetty",
		Language: "english",
		Price:    299,
		Domain:   "self-growth",
		Pages:    300,
	}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(payload)
	if err != nil {
		log.Println(err)
	}
	r, _ := http.NewRequest("POST", "/books", &buf)
	w := httptest.NewRecorder()

	AddBookHandler(w, r)
	assert.Equal(t, http.StatusCreated, w.Code)
}
