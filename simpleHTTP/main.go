package main

// import package http
import (
	"encoding/json"
	"fmt"
	"net/http"
)

// buat struct untuk di encode menjadi json
type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Page   int    `json:"page"`
}

// function hello
func Hello(w http.ResponseWriter, r *http.Request) {

	// tipe header
	w.Header().Set("Content-Type", "text/html")

	// template
	newTemplate := `<h1 style= "color:red">Hello World</h1>`

	// serve template
	w.Write([]byte(newTemplate))

}

// function untuk menampilkan data json
func ShowBook(w http.ResponseWriter, r *http.Request) {

	// set header
	w.Header().Set("Content-Type", "application/json")

	// isi data ke sruct book
	book := Book{
		Title:  "Go Web Programming",
		Author: "Reza Abdas",
		Page:   1000,
	}

	// kita encode data ke json
	json.NewEncoder(w).Encode(book)

}

func main() {

	// sediakan route dan juga funcion handler
	http.HandleFunc("/", Hello)
	http.HandleFunc("/book", ShowBook)

	fmt.Println("server berjalan pada http://localhost:8080")

	// jalankan server
	http.ListenAndServe(":8080", nil)
}
