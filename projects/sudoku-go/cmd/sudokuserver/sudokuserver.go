package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sebkraemer/100-days-of-code/projects/sudoku-go/pkg/sudokusolver"
)

func GetSolveHandler(res http.ResponseWriter, req *http.Request) {
	payload, err := ioutil.ReadAll(req.Body)
	if err == nil {
		fmt.Printf("%v", payload)
	}
	res.WriteHeader(http.StatusInternalServerError) // not implemented as head
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
