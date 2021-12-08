package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// these could probably go into a data driven table test

// instead of giving same handler I should probably put them into a muxer and test that one or the whole server

func TestGetSolveHandlerReturnsStatusBadRequest(t *testing.T) {
	handler := http.HandlerFunc(GetSolveHandler)
	rr := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/solve", nil)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Error("Did not return status 400 - Bad Request")
	}
}

func TestGetSolveHandlerReturnsStatusNotFound(t *testing.T) {
	handler := http.HandlerFunc(GetSolveHandler)
	rr := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/notsolve", nil)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNotFound {
		t.Error("Did not return status 404 - Not Found")
	}
}

func TestGetSolveHandlerReturnsStatusMethodNotAllowed(t *testing.T) {
	handler := http.HandlerFunc(GetSolveHandler)
	rr := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/solve", nil)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Error("Did not return status 405 - Method Not Allowed")
	}
}

func TestGetSolveHandlerSuccessfulRequest(t *testing.T) {
	handler := http.HandlerFunc(GetSolveHandler)
	rr := httptest.NewRecorder()
	// TODO: I should probably json-ify that once I'm using json encode/decode in other places
	payload := strings.NewReader(`{
  "field": "[ {3,1,6,5,7,8,4,9,2},{5,2,9,1,3,4,7,6,8},{4,8,7,6,2,9,5,3,1},{2,6,3,4,1,5,9,8,7},{9,7,4,8,6,3,1,2,5},{8,5,1,7,9,2,6,4,3},{1,3,8,9,4,7,2,5,6},{6,9,2,3,5,1,8,7,4},{7,4,5,2,8,6,3,1,9}]"
}`)
	req, _ := http.NewRequest(http.MethodGet, "/solve", payload)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Error("Did not return status 200 - OK")
	}
	// todo check response data, valid sudoko field
}
