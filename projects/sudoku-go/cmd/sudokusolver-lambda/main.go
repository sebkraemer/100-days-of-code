package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/sebkraemer/100-days-of-code/projects/sudoku-go/pkg/sudokusolver"
)

func HandleRequest(ctx context.Context, board sudokusolver.SudokuBoard) (string, error) {
	fmt.Printf("Processing input sudoku board:\n%s\n", board.String())

	solved := board.Solve()
	if !solved {
		return "", errors.New("could not solve")
	}

	boardStr := board.String()
	fmt.Printf("Solution:\n%s\n", boardStr)
	return boardStr, nil
}

func main() {
	lambda.Start(HandleRequest)
}
