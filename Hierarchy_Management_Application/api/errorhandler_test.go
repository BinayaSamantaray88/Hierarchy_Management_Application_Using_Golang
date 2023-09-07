package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	// Import your actual package
)

func TestSendErrorResponse(t *testing.T) {
	// Prepare a mock HTTP response recorder
	recorder := httptest.NewRecorder()

	// Call the sendErrorResponse function
	statusCode := http.StatusOK
	message := "Not Found"
	data := map[string]interface{}{"error": "Resource not found"}
	sendErrorResponse(recorder, statusCode, message, data)

	// Check the response status code and header
	assert.Equal(t, statusCode, recorder.Code)
	assert.Equal(t, "application/json", recorder.Header().Get("Content-Type"))

	// Parse the response body
	var respBody Response
	err := json.Unmarshal(recorder.Body.Bytes(), &respBody)
	assert.NoError(t, err)

	// Check the response message and data
	assert.Equal(t, statusCode, respBody.StatusCode)
	assert.Equal(t, message, respBody.Message)
	assert.Equal(t, data, respBody.Data)
}
func TestSendSuccessResponse(t *testing.T) {
	// Prepare a mock HTTP response recorder
	recorder := httptest.NewRecorder()

	// Call the sendSuccessResponse function
	statusCode := http.StatusOK
	message := "Success"
	data := map[string]interface{}{"result": "Data processed successfully"}
	sendSuccessResponse(recorder, statusCode, message, data)

	// Check the response status code and header
	assert.Equal(t, statusCode, recorder.Code)
	assert.Equal(t, "application/json", recorder.Header().Get("Content-Type"))

	// Parse the response body
	var respBody Response
	err := json.Unmarshal(recorder.Body.Bytes(), &respBody)
	assert.NoError(t, err)

	// Check the response message and data
	assert.Equal(t, statusCode, respBody.StatusCode)
	assert.Equal(t, message, respBody.Message)
	assert.Equal(t, data, respBody.Data)
}
