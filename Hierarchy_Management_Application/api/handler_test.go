package api

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestFindCategoryByID(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	query := "SELECT id, name, parentid FROM categories WHERE id=?"

	id := 1

	rows := sqlmock.NewRows([]string{"id", "name", "parentid"}).AddRow(1, "BeautyProducts", 0)
	mock.ExpectQuery(query).WithArgs(id).WillReturnRows(rows)

	rr := httptest.NewRecorder()
	category, err := findCategoryByID(rr, db, id)
	assert.NoError(t, err)
	assert.NotNil(t, category.ID)
	assert.Equal(t, 1, category.ID)
	assert.Equal(t, "BeautyProducts", category.Name)
	assert.Equal(t, 0, category.ParentID)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectations: %s", err)
	}
}

func TestUpdateCategoryByID(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/categories/{id}", UpdateCategoryByid(db)).Methods("PUT")

	requestBody := `{"name": "UpdatedName", "parentid": 0}`
	request := httptest.NewRequest("PUT", "/categories/1", nil)
	request.Body = ioutil.NopCloser(strings.NewReader(requestBody))
	request.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	mock.ExpectQuery("SELECT id FROM categories WHERE id=?").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	mock.ExpectExec("UPDATE categories set name=?,parentid=? WHERE id=?").
		WithArgs("UpdatedName", 0, 1).
		WillReturnResult(sqlmock.NewResult(0, 1))

	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, http.StatusOK, recorder.Code)
	expectedResponseBody := `{"statuscode":0,"message":"Category Updated Successfully","data":null}`
	actualResponseBody := strings.TrimSpace(recorder.Body.String())
	assert.Equal(t, expectedResponseBody, actualResponseBody)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}

func TestCreateCategory(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error occurred while creating the mock database: %v", err)
	}
	defer db.Close()

	requestBody := `{"name": "NewCategory", "parentID": 0}`
	req, err := http.NewRequest("POST", "/categories", strings.NewReader(requestBody))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()

	mock.ExpectExec("INSERT INTO categories").
		WithArgs("NewCategory", 0).
		WillReturnResult(sqlmock.NewResult(1, 1))

	handler := CreateCategory(db)
	handler(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	expectedResponseBody := `{"statuscode":0,"message":"Successfully Created","data":null}`
	actualResponseBody := strings.TrimSpace(rr.Body.String())
	assert.Equal(t, expectedResponseBody, actualResponseBody)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}

func TestDeleteCategoryByID(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/categories/{id}", DeleteCategoryById(db)).Methods("DELETE")

	request := httptest.NewRequest("DELETE", "/categories/1", nil)

	recorder := httptest.NewRecorder()

	mock.ExpectQuery("SELECT id FROM categories where id=?").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	mock.ExpectExec("DELETE FROM categories WHERE id=?").
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(0, 1))
	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
	expectedResponseBody := `{"statuscode":0,"message":"Category Deleted Successfully","data":null}`
	actualResponseBody := strings.TrimSpace(recorder.Body.String())
	assert.Equal(t, expectedResponseBody, actualResponseBody)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}
