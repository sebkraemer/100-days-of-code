package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/sebkraemer/100-days-of-code/projects/sudoku-go/pkg/sudokusolver"
)

func GetSolveHandler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/solve" {
		res.WriteHeader(http.StatusNotFound)
		// todo add HATEOAS
		return
	}
	if req.Body == nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("no data"))
		return
	}
	if req.Method != http.MethodGet {
		res.WriteHeader(http.StatusMethodNotAllowed)
		// todo add HATEOAS
		return
	}
	var board sudokusolver.SudokuBoard
	err := json.NewDecoder(req.Body).Decode(&board)
	if err != nil {
		fmt.Printf("ERROR %v", err)            // log error
		res.WriteHeader(http.StatusBadRequest) // not implemented as head
		res.Write([]byte("could not decode your data"))
		return
	}

	solved := board.Solve()
	if !solved {
		// ? solver is too stupid to check if it cannot solve a board
		// and the API is not able to return a specific error
		// with current implementation, the solver could try forever on a wrong crafted input
		// so we assume that the input was ok and we failed some other way, just for sake of this project example
		res.WriteHeader(http.StatusInternalServerError) // not implemented as head
		res.Write([]byte("could not solve your riddle, sorry"))
	}
	resBody, err := json.Marshal(board)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError) // not implemented as head
		res.Write([]byte(fmt.Sprintf("could not marshal computation result, this should not happen! result: %v", board)))
	}
	res.Write(resBody)
}

type handler struct{}

func (h handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	GetSolveHandler(w, req)
}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Starting sudoku server on port %v.\n", port)

	server := http.Server{
		Addr:    ":" + port,
		Handler: handler{},
	}

	fmt.Println(server.ListenAndServe())
}
