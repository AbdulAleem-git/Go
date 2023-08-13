package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// curl -H "Content-Type: application/json" -X POST -d {"Id":0, "Title":"Java", "Author":"ramayana murthy","ISBN":"Rampur publications"} http://127.0.0.1:8080/books
// create table
type Book struct {
	Id                  int
	Title, Author, ISBN string
}

var library map[int]Book

// database
func GetDBConnections() *sql.DB {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=abdul password=password dbname=practice_database sslmode=disable")

	if err != nil {
		fmt.Println("Error while opening a connection: ", err)
		return nil
	}
	// ping for five times
	for i := 0; i < 3; i++ {
		err = db.Ping()
		if err != nil {
			fmt.Println("Error while pinging a connection: ", err)
		}
		fmt.Println("ping")
	}

	return db
}

func RunQuery(query string) (*sql.Rows, error) {
	con := GetDBConnections()
	defer con.Close()
	return con.Query(query)
}

func updateRow(query string) (sql.Result, error) {
	con := GetDBConnections()
	defer con.Close()
	return con.Exec(query)
}

//

func listbooks(w http.ResponseWriter, r *http.Request) {
	// page 0 limit 10 start = 0 end = 9
	// page 1 limit 10 start = 10 end = 19
	// page 2 limit 10 start = 20 end = 29

	page, err := strconv.ParseInt(r.URL.Query().Get("page"), 10, 64)
	if err != nil {

		w.Write([]byte("bad request"))
	}

	limit, err := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 64)
	if err != nil {

		w.Write([]byte("bad request"))
	}

	if library == nil || len(library) == 0 {
		fmt.Fprintf(w, "Empty no library")
		return
	}

	slices := make([]Book, limit)
	start := page * limit
	end := start + limit
	idx := 0
	for start < end && idx < int(limit) {
		if b, available := library[int(start)]; available {
			slices[idx] = b
			idx++
		}
		start++
	}
	// for id, Book := range library {
	// 	fmt.Fprintf(w, "ID: %d, Book: %s, author: %s, ISBN: %s\n", id, Book.Title, Book.Author, Book.ISBN)
	// }
	err = json.NewEncoder(w).Encode(slices)
	if err != nil {
		w.Write([]byte("failed to encode to json"))
	}
	fmt.Fprintf(w, "All Done :)")

}
func listbookById(w http.ResponseWriter, r *http.Request) {
	ids := mux.Vars(r)

	// if bookId, available := ids["id"]; !available {
	// 	fmt.Fprintf(w, "ID: %d Not found", bookId)
	// 	return
	// }
	id, err := strconv.ParseInt(ids["id"], 10, 64)

	if err != nil {
		fmt.Fprintf(w, "Internal error :")
		return
	}
	book, isAvailable := library[int(id)]
	if !isAvailable {
		fmt.Fprintf(w, "Not found")
		return
	}
	fmt.Fprintf(w, "ID: %d, Book: %s, author: %s, ISBN: %s\n", id, book.Title, book.Author, book.ISBN)

}
func insertBooks(w http.ResponseWriter, r *http.Request) {
	var book Book

	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	query := fmt.Sprintf("INSERT INTO books (id, title, author, isbn) VALUES (%d, '%s', '%s', '%s')", book.Id, book.Title, book.Author, book.ISBN)
	fmt.Println(query)
	_, err = updateRow(query)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Inserted the documents!!\n"))
}
func updateExistingBook(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		fmt.Fprintf(w, "wrong data being input")
		return
	}

	var newBook Book

	if err = json.NewDecoder(r.Body).Decode(&newBook); err != nil {
		fmt.Fprintf(w, "error while decoding")
		return
	}

	query := fmt.Sprintf("UPDATE books SET title = '%s', author = '%s', ISBN = %s WHERE id = %d;", newBook.Title, newBook.Author, newBook.ISBN, id)
	_, err = updateRow(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "record is updated successfuly")
}
func deleteExistingBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)

	if err != nil {
		fmt.Fprintf(w, "wrong data")
		return
	}

	query := fmt.Sprintf("DELETE FROM books WHERE id=%d", id)
	result, err := updateRow(query)
	rowaffected, _ := result.RowsAffected()
	fmt.Fprintf(w, "%d record deleted successfuly\n", rowaffected)
}
func setHttpRoutes(r *mux.Router) {

	r.HandleFunc("/books", listbooks).Methods("GET")
	r.HandleFunc("/books/{id}", listbookById).Methods("GET")
	r.HandleFunc("/books", insertBooks).Methods("POST")
	r.HandleFunc("/books/{id}", updateExistingBook).Methods("PUT")
	r.HandleFunc("/books/{id}", deleteExistingBook).Methods("DELETE")

}

func main() {
	r := mux.NewRouter()

	setHttpRoutes(r)

	http.ListenAndServe(":8080", r)
}
