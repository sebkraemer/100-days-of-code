# Useful commands

format code:
- `goimports -l -w .`
- `go fmt` (builtin)

creating and starting a program:
- `go mod init <foo/bar>`
- `go run hello.go`

automatically add/remote modules to go.mod: `go mod tidy`

- `-ldflags="-s -w" 	 See go help build, omit the symbol table, debug information to get a smaller binary` (https://www.go-on-aws.com/lambda-go/lambda_function/simple-lambda/)


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
- how does one implement a function with a `interface{}` parameter?
- how about disabling/enabling optimizations explicitly? ("debug vs. release build")


# TODOs
- Many great articles and talks are referenced in the *Learning Go*'s Pointers chapter.
- revisit Interfaces chapter, see also https://www.airs.com/blog/archives/281
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
- on page 245, learning Go, on bottom performance benefits are mentioned when using strings, but the example from the docs is only a string of a JSON array, not an array type, so like before it works with a reader on a string. I don't see the huge benefit and wonder if the json.Decoder works so much faster on a JSON array than on line-separated json objects.
  could be a candidate for a benchmark
- difference zero value and default value (none?)
- reading and flash-carding all https://pkg.go.dev/fmt
- aws with docker: https://aws-blog.de/2021/11/lambda-container-deployment-with-cdk-using-arm-based-lambda-with-go.html
- article idea: go modules and versioning and "GOPATH mode"
  - https://go.dev/blog/v2-go-modules: ".. tools that are not version-aware — including the go command in GOPATH mode — may not distinguish between major versions"
  - https://utcc.utoronto.ca/~cks/space/blog/programming/Go117StillGopathMode
  - possibly refer day 37
- escape analysis: https://www.ardanlabs.com/blog/2017/05/language-mechanics-on-escape-analysis.html
- splay tree: https://en.wikipedia.org/wiki/Splay_tree
- memory layout
  - https://go101.org/article/memory-layout.html
  - https://golangbyexample.com/two-dimensional-array-slice-golang/#How_multidimensional_array_is_stored_in_memory
- PR for fluentd + docker compose example in fluentd docs, see github issue where I commented
- have a look for useful stuff: https://gitlab.com/Figur81/100dayscodegolang
- check suspended and todo-tagged flash cards, aim for unsuspending
- checkout go tags! (build configs?)
- repository pattern, "constructor" vs. DDD abstraction


# Quiz ideas
- zero/default values
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
- something about memory layout, see above
- memory sharing properties of subslices: https://go.dev/play/p/txAYy0bwmRE
- can you take the address of a map element? (no. why? rehashing might invalidate)
- what is unnamed pointer type? -> unnamed type: composition of previously declared types -> type literal (p. 135 Go programming lang book)
- anon function. a) function that is not assigned to a variable, b) value as the result of function expression (YES), c) ..
  - bonus: why is it important and useful: named functions are only allowed on package level, so for defer() and other closure uses, anonymous functions are necessary.

# Other / links
- GopherCon 2018: Kat Zien - How Do You Structure Your Go Apps: https://www.youtube.com/watch?v=oL6JBUk6tj0
- Semantic Import Versioning: https://research.swtch.com/vgo-import
- Minimal Version Selection: https://www.ardanlabs.com/blog/2019/12/modules-03-minimal-version-selection.html
- https://go.dev/doc/
- cloud native
  - https://www.capitalone.com/tech/cloud/what-is-cloud-native/
- https://gophercises.com/

# Commands

`minikube start --listen-address=0.0.0.0 --driver=hyperkit --no-kubernetes`
`eval $(minikube -p minikube docker-env)` # or move this os bash profile

## with vmware driver!

`minikube start --listen-address=0.0.0.0 --driver=vmware --no-kubernetes --memory 4096 --cpus 4`
`eval $(minikube docker-env)`

to make mount available to docker: `minikube mount /Users/s/src:/Users/s/src`

