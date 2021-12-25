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
- what does & really do, always create a type on the heap? heap vs stack not explicitly mentioned in spec
- when exaclty does import execute anything from a package?
- with EXPOSE in a dockerfile, is it necessary to specify the port when running the container?
  https://docs.docker.com/engine/reference/run/#expose-incoming-ports
- goroutines that are still running, are a problem when returning from a function or quitting the program? (see `channels.go` play project)


# TODOs
- Many great articles and talks are referenced in the *Learning Go*'s Pointers chapter.
- revisit Interfaces chapter
- `go generate`
- local packages OK after all? (ref. `./projects/golang-fifa-world-cup-web-service`)
- (....) (I missed to note a few items, I'm sure..)
- example on page 235 (learning go), could that be changed into returning a ReadCloser?
- `retract` keyword is not mentioned in learning Go!
- read up on bytes.Buffer
- read up on bufio.{Reader,Writer,Scanner} as ioutil.Read/WriteAll replacement
- *prometheus*,
  write something to export /metrics endpoint and monitor it
  https://app.pluralsight.com/course-player?clipId=26bf6f6f-c9e5-46af-aa11-30304069d022
  using golang's templating mechanism maybe



# Quiz ideas
- difference zero value and default value (none?)
- different initializations of e.g. map or slice, when nil, when not?
- monotonic time vs. wall clock time, different handling in time.Time when e.g. captured with time.Now(), difference in calculation (because of time zone handling?) (time in io chapter)
- go implicit conversion with interface?
- execution of defer and capturing
  - https://stackoverflow.com/a/53219947
  - https://go.dev/ref/spec#Defer_statements
  - what's the difference between invoking and executing a function? (both words are used in spec when describing defer())
- `retract` keyword is not mentioned in learning Go as it's rather new
- var a = 10; pa := &a, what type is pa?
- difference !bytes.Equal() vs bytes.Compare(), linter told me to use the first
- encoding of zero value struct, zero value slice/map, zero size slice/map
- is it necessary to call close() on a channel after use? (no)
  https://stackoverflow.com/a/36613932/4312669
