# Useful commands

format code:
- `goimports -l -w .`
- `go fmt` (builtin)

creating and starting a program:
- `go mod init <foo/bar>`
- `go run hello.go`

automatically add/remote modules to go.mod: `go mod tidy`


# Open questions

- debug value which is declared in an if statement, see day 3.
- need for clarification: compare my flash card definition for block (scope) with the switch description in the book, look it up language spec!
- when to use which variable declaration?
- on what types does the for range loop work?
- iterating over map is not deterministic (anti hash DOS mechanism), but fmt.Print print it sorted. how? is it possible without getting and sorting all values first?
- Pointers chapter notices that the interface{} parameter needs a pointer and that this API pattern is used because of lack of generics. Generics being almost there, how does that compare? will e.g. the JSON.Unmarshal method change?

# TODOs
- Many great articles and talks are referenced in the *Learning Go*'s Pointers chapter.
- revisit Interfaces chapter
- `go generate`
- local packages OK after all? (ref. `./projects/golang-fifa-world-cup-web-service`)
