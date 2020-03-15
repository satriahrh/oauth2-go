package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/satriahrh/oauth2-go/handler"
)

func TestHealthzHandler(t *testing.T) {
	request, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()

	handler.HealthzHandler(responseRecorder, request)

	// Check the status code is what we expect.
	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"message":"OK"}` + "\n"
	actual := responseRecorder.Body.String()
	if actual != expected {
		t.Errorf(
			"handler returned unexpected body:\ngot\nwant\n%v\n%v",
			actual, expected,
		)
	}
}
