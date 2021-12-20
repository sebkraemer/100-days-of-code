module github.com/sebkraemer/100-days-of-code/projects/sudoku-go/cmd/sudokuserver

go 1.17

require (
	github.com/sebkraemer/100-days-of-code/projects/sudoku-go/pkg/sudokusolver v0.0.0
	github.com/sirupsen/logrus v1.8.1
)

require golang.org/x/sys v0.0.0-20191026070338-33540a1f6037 // indirect

replace github.com/sebkraemer/100-days-of-code/projects/sudoku-go/pkg/sudokusolver => ../../pkg/sudokusolver
