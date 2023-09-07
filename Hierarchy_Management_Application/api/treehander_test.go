package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetCategoryTreeHandler(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery("SELECT id, name, parentid FROM categories WHERE parentid=0").
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "parentid"}).AddRow(1, "Root", 0))

	mock.ExpectQuery("SELECT id, name, parentid FROM categories WHERE parentid=?").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "parentid"}).AddRow(2, "Child", 1))

	req := httptest.NewRequest("GET", "/categories", nil)
	w := httptest.NewRecorder()

	handler := GetCategoryTreeHandler(db)
	handler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

}
