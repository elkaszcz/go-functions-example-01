package samplefunctions

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestAlwaysError tests TestHelloWorld GET function
func TestHelloWorldGET(t *testing.T) {
	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)
	HelloWorld(rr, req)
	if rr.Result().StatusCode != http.StatusOK {
		t.Errorf("HelloWorld StatusCode = %v, want %v", rr.Result().StatusCode, http.StatusOK)
	}
}

// TestAlwaysError tests TestHelloWorld GET function
func TestHelloWorldPOST(t *testing.T) {
	rr := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/", nil)
	HelloWorld(rr, req)
	if rr.Result().StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("HelloWorld StatusCode = %v, want %v", rr.Result().StatusCode, http.StatusMethodNotAllowed)
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
