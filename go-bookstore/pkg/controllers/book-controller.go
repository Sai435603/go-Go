package controllers

import (
	"bookstore/pkg/models"
	"bookstore/pkg/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	books := models.GetAllBooks()
	json.NewEncoder(w).Encode(books)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	// fmt.Println("here is the id" + strconv.FormatInt(id, 10))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	book, db := models.GetBookById(int(id))
	if db.Error != nil {
		http.Error(w, db.Error.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(book)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book models.Book
	utils.ParseBody(r, &book)
	b := book.CreateBook()
	json.NewEncoder(w).Encode(b)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    id, err := strconv.ParseInt(params["id"], 10, 64)
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }
    bookDetails, db := models.GetBookById(int(id))
    if db.Error != nil {
        http.Error(w, db.Error.Error(), http.StatusNotFound)
        return
    }
    var updatedBook models.Book
    utils.ParseBody(r, &updatedBook)
    if updatedBook.Name != "" {
        bookDetails.Name = updatedBook.Name
    }
    if updatedBook.Author != "" {
        bookDetails.Author = updatedBook.Author
    }
    if updatedBook.Description != "" {
        bookDetails.Description = updatedBook.Description
    }
    db.Save(&bookDetails)
    json.NewEncoder(w).Encode(bookDetails)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	book := models.DeleteBook(int(id))
	json.NewEncoder(w).Encode(book)
}
