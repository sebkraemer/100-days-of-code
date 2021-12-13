package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetSolveHandler(t *testing.T) {
	cases := []struct {
		testName       string
		clientMethod   string
		clientUri      string
		clientBody     io.Reader
		expectedStatus int
		expectedBody   *bytes.Buffer
	}{
		{"empty request", http.MethodGet, "/solve", nil, http.StatusBadRequest, bytes.NewBufferString("no data")},
		{"no root handler", http.MethodGet, "/", nil, http.StatusNotFound, nil},
	}
	for _, test := range cases {
		t.Run(test.testName, func(t *testing.T) {
			handler := http.HandlerFunc(GetSolveHandler)
			rr := httptest.NewRecorder()
			req, _ := http.NewRequest(test.clientMethod, test.clientUri, test.clientBody)
			handler.ServeHTTP(rr, req)

			if got, expected := rr.Code, test.expectedStatus; got != expected {
				t.Errorf("status mismatch: got %d, expected %d", got, expected)
			}

			var resBody []byte
			var expectedBody []byte
			if rr.Body != nil {
				resBody, _ = ioutil.ReadAll(rr.Body)
			}
			if test.expectedBody != nil {
				expectedBody, _ = ioutil.ReadAll(test.expectedBody)
			}

			if !bytes.Equal(resBody, expectedBody) {
				t.Errorf("response body mismatch, got %v, expected %v", resBody, expectedBody)
			}
		})
	}
}

func TestGetSolveHandlerReturnsStatusBadRequestInvalidData(t *testing.T) {
	handler := http.HandlerFunc(GetSolveHandler)
	rr := httptest.NewRecorder()
	payload := strings.NewReader(`{ "data": "[ {3,1,6,5,7,8,4,9,2},{5,2,9,1,3,4,7,6,8},{4,8,7,6,2,9,5,3,1},{2,6,3,4,1,5,9,8,7},{9,7,4,8,6,3,1,2,5},{8,5,1,7,9,2,6,4,3},{1,3,8,9,4,7,2,5,6},{6,9,2,3,5,1,8,7,4},{7,4,5,2,8,6,3,1,9}]"}`)
	req, _ := http.NewRequest(http.MethodGet, "/solve", payload)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Did not return status 400 - Bad Request, returned %d instead", status)
	}
}

func TestGetSolveHandlerReturnsStatusMethodNotAllowed(t *testing.T) {
	handler := http.HandlerFunc(GetSolveHandler)
	rr := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/solve", strings.NewReader("nonempty"))
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Error("Did not return status 405 - Method Not Allowed")
	}
}

func TestGetSolveHandlerSuccessfulRequest(t *testing.T) {
	handler := http.HandlerFunc(GetSolveHandler)
	rr := httptest.NewRecorder()
	payload := strings.NewReader(`{ "data": [ [3, 0, 6, 5, 0, 8, 4, 0, 0], [5, 2, 0, 0, 0, 0, 0, 0, 0], [0, 8, 7, 0, 0, 0, 0, 3, 1], [0, 0, 3, 0, 1, 0, 0, 8, 0], [9, 0, 0, 8, 6, 3, 0, 0, 5], [0, 5, 0, 0, 9, 0, 6, 0, 0], [1, 3, 0, 0, 0, 0, 2, 5, 0], [0, 0, 0, 0, 0, 0, 0, 7, 4], [0, 0, 5, 2, 0, 6, 3, 0, 0] ]}`)
	req, _ := http.NewRequest(http.MethodGet, "/solve", payload)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Error("Did not return status 200 - OK")
	}

	expected := `{"data":[[3,1,6,5,7,8,4,9,2],[5,2,9,1,3,4,7,6,8],[4,8,7,6,2,9,5,3,1],[2,6,3,4,1,5,9,8,7],[9,7,4,8,6,3,1,2,5],[8,5,1,7,9,2,6,4,3],[1,3,8,9,4,7,2,5,6],[6,9,2,3,5,1,8,7,4],[7,4,5,2,8,6,3,1,9]]}`
	bodyBytes, err := ioutil.ReadAll(rr.Body)
	if err != nil {
		t.Error(err)
	} else {
		got := string(bodyBytes)
		if expected != got {
			t.Errorf("unexpected solved response: got\n%v, expected\n%v", got, expected)
		}
	}
}
