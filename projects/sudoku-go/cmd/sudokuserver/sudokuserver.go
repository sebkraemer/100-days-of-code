package main

import (
	"encoding/json"
	"fmt"
	"net/http"

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

func main() {
	//field := SudokuBoard{} // fails or runs forever with empty board?
	field := sudokusolver.SudokuBoard{
		Data: [sudokusolver.N][sudokusolver.N]int{{3, 0, 6, 5, 0, 8, 4, 0, 0},
			{5, 2, 0, 0, 0, 0, 0, 0, 0},
			{0, 8, 7, 0, 0, 0, 0, 3, 1},
			{0, 0, 3, 0, 1, 0, 0, 8, 0},
			{9, 0, 0, 8, 6, 3, 0, 0, 5},
			{0, 5, 0, 0, 9, 0, 6, 0, 0},
			{1, 3, 0, 0, 0, 0, 2, 5, 0},
			{0, 0, 0, 0, 0, 0, 0, 7, 4},
			{0, 0, 5, 2, 0, 6, 3, 0, 0}},
	}

	fmt.Printf("%s\n", field.String())
	solved := field.Solve()
	fmt.Printf("solved: %t\n%s\n", solved, field.String())

	asBytes, err := json.Marshal(field.Data)
	if err != nil {
		fmt.Errorf("could not marshal solved sudoku data")
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(asBytes)
	})
	err = http.ListenAndServe(":5000", nil)
	if err != nil {
		fmt.Errorf("%w", err)
	}
}
