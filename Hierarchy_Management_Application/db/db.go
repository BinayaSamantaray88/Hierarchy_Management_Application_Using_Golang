package db

import "database/sql"

func ConnectDB() (*sql.DB, error) {
	//Update with your MYSQL database details
	db, err := sql.Open("mysql", "root:1322@tcp(127.0.0.1:3306)/htmlcourse")
	if err != nil {
		return nil, err
	}
	return db, nil
}
