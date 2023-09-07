package api

import (
	"database/sql"
	"net/http"
)

func buildCategoryTree(db *sql.DB) ([]Category, error) {
	rootCategories := []Category{}

	rows, err := db.Query("SELECT id, name, parentid FROM categories WHERE parentid=0")
	if err != nil {
		return rootCategories, err
	}
	defer rows.Close()

	for rows.Next() {
		var root Category
		err := rows.Scan(&root.ID, &root.Name, &root.ParentID)
		if err != nil {
			return rootCategories, err
		}

		err = buildChildren(db, &root)
		if err != nil {
			return rootCategories, err
		}

		rootCategories = append(rootCategories, root)
	}

	return rootCategories, nil
}

func buildChildren(db *sql.DB, parent *Category) error {
	rows, err := db.Query("SELECT id, name, parentid FROM categories WHERE parentid=?", parent.ID)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var child Category
		err := rows.Scan(&child.ID, &child.Name, &child.ParentID)
		if err != nil {
			return err
		}

		err = buildChildren(db, &child)
		if err != nil {
			return err
		}

		parent.Children = append(parent.Children, child)
	}

	return nil
}
func GetCategoryTreeHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rootCategory, err := buildCategoryTree(db)
		if err != nil {
			statuscode := 1
			message := err.Error()
			sendErrorResponse(w, statuscode, message, nil)
			return
		}

		statuscode := 0
		message := "Tree generated successfully"
		sendSuccessResponse(w, statuscode, message, rootCategory)
	}
}
