module github.com/sebkraemer/100-days-of-code/projects/sudoku-go/cmd/sudosolver-lambda

go 1.17

require (
	github.com/aws/aws-lambda-go v1.27.1
	github.com/sebkraemer/100-days-of-code/projects/sudoku-go/pkg/sudokusolver v0.0.0

)

replace github.com/sebkraemer/100-days-of-code/projects/sudoku-go/pkg/sudokusolver => ../../pkg/sudokusolver
