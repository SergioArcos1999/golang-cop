package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Books []Book

type Book struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Pages int    `json:"pages"`
}

type GetAllBooksResponse struct {
	BookList Books `json:"books"`
	Quantity int   `json:"quantity"`
}

var bookList Books

func main() {
	http.HandleFunc("/books", handleRead)
	http.HandleFunc("/add-book", handleWrite)
	http.ListenAndServe(":8000", nil)
}

func handleWrite(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "POST":
		decoder := json.NewDecoder(request.Body)
		var book Book
		decoder.Decode(&book)

		bookList = append(bookList, book)
		fmt.Println(bookList)
	default:
		message := request.Method + " method not allowed"
		http.Error(writer, message, http.StatusMethodNotAllowed)
	}
}

func handleRead(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		response := GetAllBooksResponse{bookList, len(bookList)}
		json.NewEncoder(writer).Encode(response)
	default:
		message := request.Method + " method not allowed"
		http.Error(writer, message, http.StatusMethodNotAllowed)
	}
}
