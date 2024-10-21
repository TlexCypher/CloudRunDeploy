package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	recorder := httptest.NewRecorder()

	req := httptest.NewRequest("GET", "/", nil)

	Handler(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("expected status code %v, got %v", http.StatusOK, status)
	}

	expectedBody := "Success to start go http server and auto-deploy to Cloud Run."
	if strings.TrimSpace(recorder.Body.String()) != expectedBody {
		t.Errorf("expected body '%v', got '%v'", expectedBody, recorder.Body.String())
	}
}
