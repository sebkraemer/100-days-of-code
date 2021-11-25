module github.com/sebkraemer/100-days-of-code/projects/fints-go

go 1.17

require github.com/sebkraemer/100-days-of-code/projects/fints-go/pkg/parser v0.0.0-00010101000000-000000000000

require github.com/antlr/antlr4/runtime/Go/antlr v0.0.0-20211115101625-aeaa445b4d4f // indirect

replace github.com/sebkraemer/100-days-of-code/projects/fints-go/pkg/parser => ./pkg/parser
