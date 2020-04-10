package samplefunctions

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestAlwaysError tests AlwaysError function
func TestHelloWorld(t *testing.T) {
	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)
	HelloWorld(rr, req)
	if rr.Result().StatusCode != http.StatusOK {
		t.Errorf("HelloWorld StatusCode = %v, want %v", rr.Result().StatusCode, http.StatusOK)
	}
}

// TestAlwaysError tests AlwaysError function
func TestAlwaysError(t *testing.T) {
	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)
	AlwaysError(rr, req)
	if rr.Result().StatusCode != http.StatusInternalServerError {
		t.Errorf("AlwaysError StatusCode = %v, want %v", rr.Result().StatusCode, http.StatusInternalServerError)
	}
}
