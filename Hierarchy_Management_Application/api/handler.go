package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func findCategoryByID(w http.ResponseWriter, db *sql.DB, id int) (*Category, error) {
	var category Category
	err := db.QueryRow("SELECT id, name, parentid FROM categories WHERE id = ?", id).Scan(&category.ID, &category.Name, &category.ParentID)
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func CreateCategory(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newCategory Category
		err := json.NewDecoder(r.Body).Decode(&newCategory)
		if err != nil {
			statuscode := 1
			message := err.Error()
			sendErrorResponse(w, statuscode, message, nil)
			return
		}
		defer r.Body.Close()
		if newCategory.ParentID < 0 {
			statuscode := 1
			message := "Inalid Parent ID"
			sendErrorResponse(w, statuscode, message, nil)
			return

		}
		if newCategory.ParentID > 0 {
			parent, err := findCategoryByID(w, db, newCategory.ParentID)
			fmt.Println(parent)
			if err != nil {
				statuscode := 1
				message := "Parent Not Found"
				sendErrorResponse(w, statuscode, message, nil)
				return
			}

		}

		_, err = db.Exec("INSERT INTO categories(name, parentid) VALUES(?, ?)", newCategory.Name, newCategory.ParentID)
		fmt.Println(newCategory.Name)
		if err != nil {
			statuscode := 1
			message := err.Error()
			sendErrorResponse(w, statuscode, message, nil)
			return
		}

		statuscode := 0
		message := "Successfully Created"
		sendSuccessResponse(w, statuscode, message, nil)
	}

}
func UpdateCategoryByid(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var updatedcategory Category
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			message := err.Error()
			statuscode := 1
			sendErrorResponse(w, statuscode, message, nil)
			return
		}
		if id <= 0 {
			statuscode := 1
			message := "Invalid ID"
			sendErrorResponse(w, statuscode, message, nil)
			return
		}
		err = json.NewDecoder(r.Body).Decode(&updatedcategory)
		if err != nil {
			message := err.Error()
			statuscode := 1
			sendErrorResponse(w, statuscode, message, nil)
			return
		}
		rows, err := db.Query("SELECT id FROM categories WHERE id=?", id)
		if err != nil {
			message := err.Error()
			statuscode := 1
			sendErrorResponse(w, statuscode, message, nil)
			return
		}
		defer rows.Close()
		if !rows.Next() {
			message := "Category Not Found"
			statuscode := 1
			sendErrorResponse(w, statuscode, message, nil)
			return
		}
		_, err = db.Exec("UPDATE categories set name=? WHERE id=?", updatedcategory.Name, id)
		if err != nil {
			statuscode := 1
			message := err.Error()
			sendErrorResponse(w, statuscode, message, nil)
			return
		}
		message := "Category Updated Successfully"
		statuscode := 0
		sendSuccessResponse(w, statuscode, message, nil)

	}
}
func DeleteCategoryById(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			message := err.Error()
			statuscode := 1
			sendErrorResponse(w, statuscode, message, nil)
			return
		}
		if id <= 0 {
			statuscode := 1
			message := "Invalid ID"
			sendErrorResponse(w, statuscode, message, nil)
			return
		}
		rows, err := db.Query("SELECT id FROM categories where id=?", id)
		if err != nil {
			statuscode := 1
			message := err.Error()
			sendErrorResponse(w, statuscode, message, nil)
			return
		}
		defer rows.Close()

		if !rows.Next() {
			statuscode := 1
			message := "Category Not Found"
			sendErrorResponse(w, statuscode, message, nil)
			return
		}
		_, err = db.Exec("DELETE FROM categories WHERE id=?", id)
		if err != nil {
			statuscode := 1
			message := err.Error()
			sendErrorResponse(w, statuscode, message, nil)
			return
		}
		message := "Category Deleted Successfully"
		statuscode := 0
		sendSuccessResponse(w, statuscode, message, nil)
	}
}
