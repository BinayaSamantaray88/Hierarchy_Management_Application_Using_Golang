package main

import (
	"database/sql"
	"log"
	"net/http"
	"task/api"
	"task/db"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	db, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err)

	}
	defer db.Close()
	createTable(db)

	r := mux.NewRouter()
	r.HandleFunc("/categories", api.CreateCategory(db)).Methods("POST")
	r.HandleFunc("/categories/category-tree", api.GetCategoryTreeHandler(db)).Methods("GET")
	r.HandleFunc("/categoriest/{id}", api.UpdateCategoryByid(db)).Methods("PUT")
	r.HandleFunc("/categories/{id}", api.DeleteCategoryById(db)).Methods("DELETE")
	log.Println("Server started on 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func createTable(db *sql.DB) {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS categories (
		id INT AUTO_INCREMENT PRIMARY KEY,
		parentid INT,
		name VARCHAR(255)
	);
`
	_, err := db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
}
