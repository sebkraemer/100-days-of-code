# 100 Days Of Code - Log

## Day 1

November 11, 2021

**Today's Progress**

Installed Go and setup VS Code, followed along the tutorial at https://golang.org/doc/code and the related Youtube video https://www.youtube.com/watch?v=1MXIGYrMk80.

**Thoughts**

I enjoyed getting started. Tools and commands are a lot to grasp at once. Regarding coding, I really noticed how useful it was that I have started with the *Learning Go* book a while ago. I felt I could apply some of the things and did not just copy some content.

**Link to work:** [day001](./day001)

## Day 2

November 12, 2021

**Today's Progress**

Less of a Go than a DevOps day.
Noticed that the one test actually failed. Fixed it and added a github workflow to build and test automatically. Added the badge to the README, that just looks nice.
A bit of ground work, reading about different kinds of for loops in Go. Sorry, no code.

**Thoughts**

Always test. And make sure that tests do run!
I like doing DevOps.

**Link to work:**
- [ReverseString bugfix](https://github.com/sebkraemer/100-days-of-code/commit/db6c39ffa216d7cb17193e3dd90ae10e5ad6fd74)
- [day001 test and build workflow](./actions/runs/1454887039)

## Day 3

November 13, 2021

**Today's Progress**
- Continued to read the *Learning Go* book but didn't find the time to write more flash cards.
- `controls.go` program with playground code with if, for ("four for loops"), break, continue, , labels, switch/case

**Thoughts**
- I need some getting used to opening brackets on the same line..
- wonder how to see the value in the debugger for values declared in scope of if statement, e.g. `if n := getValue(); ..`, VS Code won't show me the value.
- I wonder if fmt.Print does get and sort all values before outputting, because it's always sorted but iterating over the map has a random portion and is nondeterministic.

**Link to work:** [day003](./day003)


## Day 4

November 14, 2021

**Today's Progress**
- updated (added) flash cards
- started chapter about functions
- started with a @codility lesson but couldn’t finish (too late, really 🥱)

**Link to work:** [day004](./day004)


## Day 5

November 15, 2021

**Today's Progress**
- Continued functions topic

**Thoughts**
- I like reading and learning from the book. Sometimes I look up a new topic at [the spec](https://golang.org/ref/spec#Function_types) to see what the language reference says. 🤓📘  Compared to C++, Go's ref is sooo easy to read.
- A quote by @jonbodner from the “Functions” chapter I want to share: *Error handling is what separates the professionals from the amateurs.* ☝️

**Link to work:** --


## Day 6

November 16, 2021

**Today's Progress**

- wrote the book's calculator example from memory, solved it a bit differently, I think
- updated (added) flash cards

**Thoughts**
- It didn't go too bad when writing the calculator while combining several topics that I conquered from the book!
- Still working with the book (which I think is the right thing) does not leave much room for starting more serious coding work. Maybe I should even prioritize the book part to be done with it sooner than later.

**Link to work:** [day006](./day006)


## Day 7

November 17, 2021

**Today's Progress**

- *Learning Go* pages 107-132/340
- added memory chapter flash cards

**Thoughts**
- Interesting chapter abount memory (heap) management and garbage collection.
  It appears Go finds a good tradeoff between performance and using the pass by value approach. Once more I recognize roots of C but leaving out most of its heritage. Even in C++ community, traditionally very performance concerned folks, I have read about the advantages of value based interfaces compared to optimizing performance but making compiler's and developer's work harder to track those variables.

## Day 8

November 18, 2021

**Today's Progress**

- played around with slice memory management, verifying my understanding of subslicing, appending, forced copying of memory blocks (cap exceeded cases)
- did my bit repeating flash cards (as I always do)
- read the article about @antlr with @Golang and discovered struct's embedded fields (just a few pages ahead in my book, hadn't read that yet)
- #100DaysOfCode socializing
- listened to most of "Go time" podcast from @changelog: Building actually maintainable software https://gotime.fm/196

**Thoughts**
- What's with all the WakaTime (sic) tweets in my #100DaysOfCode feed all of a sudden? 🤨
- Behaviour that is hard to describe in detail (many words) are easy to grasp once you get to the roots of it (i.e., understand slice's memory layout and behavior)
- Regarding the podcast I liked the most "Untestable code is unmaintainable code" (uncredited) and "fight for your space and time to keep your software maintainable" (uncredited). And in general it's interesting to see that maintainability has several aspects and how to approach the problem differs in teams and depending on the kind of development phase and product.

**Link to work:** [day008](./day008)


## Day 9

November 19, 2021

**Today's Progress**

- revisited (and solved) codility's binarygap problem (I'm no competitive programmer at this point, that's for sure 😆)
- dabbled with #vscode Go unit testing, but leaving the topic for another day

**Thoughts**

Competitive programming tasks are quite unusual for me so getting to the solution was rather time intensive, given it wasn't a difficult problem.
While practicing these kind of things are generally good, I'm not sure it's the best use of my time *right now*. Programming the solution is rather easy (language-wise) but getting to a working algorithm is a much greater part of the challenge. And I'm here for learning to apply the language.

I also had a bit of a problem with the local `binarygap` module but it was fixed by not registering the package in the subfolder.

No book reading today, I'm afraid.

**Link to work:** [day009](./day009)


## Day 10

November 20, 2021

**Today's Progress**

- solved confusion about %q and %c format specifiers
- revistied method calling on nil pointer, can be OK!
- *Learning Go* pages 133-145/340

**Thoughts**

Possibly genius but it needed a second though to grasp this:

```
var it *IntTree // zero value: nil
it = it.Insert(5) // might panic or not
```

works because a method can be called on a nil instance. If it is implemented to support the case, good things happen.


## Day 11

November 21, 2021

**Today's Progress**

- *Learning Go* pages 146-157/340

**Thoughts**

Nontrivial stuff in the Interfaces chapter. It’s slowing me down 😞 but I also realise that this is essential stuff.

It’s true that Go is always famous for its concurrency model but not for its implicit interfaces.


## Day 12

November 22, 2021

**Today's Progress**

- *Learning Go* pages 157-160/340, finishing Interfaces chapter with Dependency Injection
- Bootstrapped fints-go, a parsing project in Go.

**Thoughts**

Apparently functions are a so much first class citizens that you can even declare methods on them, and this can be used to implement an interface. 🙇‍♂️ It's quite fascinating.

I don't find the tweet right now that I stumbled upon today (it was a #100DaysOfCode one!), anyway it recommends a 70/30 practicing to stuying ratio for learning a new topic. I'm too much on the learning side. On the other hand, maybe that ration changes and it's good to have a stronger theory part when starting off to code your project.

Interface chapter is not so straightforward to put into flash cards and I certainly don't want that holding me back any longer. Will have to get back to it. However, I believe some concepts have become clear by now.

Build!

**Link to work:**

Today's commit: [81402b5](https://github.com/sebkraemer/100-days-of-code/commit/81402b580ca06f3c05c80321dedb9fb3740a93f8)

Project: [fints-go](./projects/fints-go) (This will contain more than I wrote today when you check it out later.)


## Day 13

November 23, 2021

**Today's Progress**

- *Learning Go* pages 161-170/340, starting *Errors* chapter
- Playing with type and interface
- Filing an errata on the book ☝️
- ~~continued with fints-go~~

**Thoughts**

The Go parser from ANTKR needs some love, apparently there is generated AST visitor code but it's not as straightforward and easy to setup (and not as idiomatic?) in the Go version as for other languages.
I had dabbled with it with python and JavaScript and also found there that the quality isn't on par with the Java version, which is the prime citizen, I believe. See e.g. [issue #1807](https://github.com/antlr/antlr4/pull/1807).

Not sure what to make of it and how far I should take this project idea.

I was happy to feel comfortable writing an interface and type switch case and type conversion on the playground. Don't even know what led me to do it, but please see [today's log](./day013) for the story.

**Link to work:**

Today's commits: 
- [day013](./day013)
- Project: [fints-go](./projects/fints-go)


## Day 14

November 24, 2021

**Today's Progress**

- *Learning Go* pages 171-202/340, finishing Errors chapter and also completing Packaging chapter. 📖

**Thoughts**

No new flash cards and no coding but I hope this does count anyway..

Next up: Concurrency! 🤩  In a way I'm of course looking forward to it, on the other hand I feed it's not the most important thing to cover next..


## Day 15

November 25, 2021

**Today's Progress**

- had to create a few flash cards, I'm running out of reviews soon
- picked up the fints-go idea and made the project compile

**Thoughts**

For getting the project to build, I had to fight a bit with local package dependency; in the end I successfully used `replace()` in my `go.mod` 😀 (The book mentions `replace` but unfortunately doesn't explain it further.)

```go.mod
replace (
   github.com/sebkraemer/100-days-of-code/projects/fints-go/pkg/parser => ./pkg/parser
)
```

Going further, I did notice that the Go ANTLR target, at least the visitor implementation I intended to use, has bigger problems than anticipated (see [here](https://github.com/antlr/antlr4/issues/2504) and extensive comments [here](https://github.com/antlr/antlr4/pull/1841#pullrequestreview-66582914)) and I'm not sure I want to go down that path. I could still try to come up with a POC but reimplementing every base method is not feasable for a grammar that will grow large quickly, and the FinTS one would. An example of someone who made the effort for a simple calulator (minum set of rules) can be seen [here](https://github.com/DavidGamba/go-antlr-calc).

Using SomeType.New() object creation seems to be a thing.

Have to look up `go generate`.

**Link to work:**

- Today's commits: https://github.com/sebkraemer/100-days-of-code/commit/8ca9d1dfdd5c77bfc287f1cb6b2598956772b64a
- Project: [fints-go](./projects/fints-go)


## Day 16

November 26, 2021

**Today's Progress**

- Started sudoku solving project but got a bit stuck in the early stage. Better progress next time! 🤞


## Day 17

November 26, 2021

**Today's Progress**

- ❌ no coding today 🤭

- ✅ spent the whole evening with my wife 😇

- Started chapter about standard library and had a brief look at backtracking.


## Day 18

November 28, 2021

**Today's Non-Progress**

intentionally left blank


## Day 19

November 29, 2021

**Today's Progress**

- continued on sudoku
- early stage: print support, test layout assumptions
- used stringer interface for printing support
- used strings.Builder
- inspected array memory layout (row major, of course, like in C 🙂)

**Thoughts**

After a break, coding was fun and I feel motivated with some things in mind where to go with this toy project.

Implementing String() made me realize what I already knew: It has nothing to do with operator overloading, which go does not support.
I believe that would be against the philosophy of Go that things should not be hidden or obfuscated, even if that means more code.
This does not mean that odules like `fmt` can't use whatever mechanism (reflection?) to check for an interface and handle things comfortably.

**Link to work:**

- Today's commits: https://github.com/sebkraemer/100-days-of-code/commit/b4806f9491889806f2d8bb29d726765ce8118564
- Project: [sudoku-go](./projects/sudoku-go)


## Day 20

November 30, 2021

**Today's Progress**

- finished first version of backtracking sudoku solver, it solved an example sudoku ✅

**Thoughts**

- Finally implemented this. Had backtracking on my list for too long!
- The mix of bool and error type which I started with felt weird for the helper functions and recursion results; ended up with bools only.
- Stringer interface does only work with `fmt` if fields are not exported!

**Link to work:**

- Today's commits: https://github.com/sebkraemer/100-days-of-code/commit/11125a10752a101371b78d03c1050482e6bbbc5a
- Project: [sudoku-go](./projects/sudoku-go) with updated README


## Day 21

December 1, 2021

**Today's Progress**

- Couldn't resist and started the testing chapter
- turned a sudoku test into a benchmark test
- flash-carded the benchmark section

**Thoughts**

I don't see a good way to have code written both for a test and a benchmark without code duplication or extracting into a yet another function. I feel that the latter would hurt readibility, but maybe it doesn't.

**Link to work:**

- Today's commits: https://github.com/sebkraemer/100-days-of-code/commit/621ec13634d1fa72727b2d7e8dc284df990e9ebf with benchmark output
- Project: [sudoku-go](./projects/sudoku-go)


## Day 22

December 2, 2021

**Today's Progress**

- Revisiting modules and flash-carding
- Finally understood Minimal Version Selection 🎯 after reading a great article by @goinggodotnet, the example in the book had me confused.
  https://t.co/kJoulWbWN1

no code to share today, sorry

**Thoughts**

It felt good to create new cards and work with the book again. While it feels a bit wrong to skip on the coding/practical side, I notice that I'm writing down stuff on cards for repetition that I had read but rather forgotten the specifics about.

**Link to work:**

- Today's commits: https://github.com/sebkraemer/100-days-of-code/commit/621ec13634d1fa72727b2d7e8dc284df990e9ebf with benchmark output
- Project: [sudoku-go](./projects/sudoku-go)


## Day 23

December 3, 2021

**Today's Progress**

- Started looking into web services with go and implemented a *hello world* for you guys.
- Also peeked into AWS Lambda as a candidate for future serving my sudoko solver but I concluded quickly that this is a larger topic on its own.
  I consider it as a serverless option for the soduko solver.

**Thoughts**

It appears I've gotten used to put the opening curly brace on the same line as the introductory declaration.

It's just fascinating how easy it is to ramp up a webservice in Go (coming from C++, that is)

**Link to work:**

- Today's commits: https://github.com/sebkraemer/100-days-of-code/commit/232c7d9a3f99050f852c86e242bd56913fed4d58
- Project: [day023](./day023)


## Day 27

December 7, 2021

**Today's Progress**

- Took the @pluralsight lab "Build a FIFA World Cup Web Service in Go"
  (which sounds like quite more work than it was)

**Thoughts**

The project was in a way too easy because it dictated what to write. Still, I got a good impression on how to write
the handlers. Also, there is http request testing code which I can use for my own future service tests.
I felt comfortable to write the Go code.

The project used local packages which seemed pretty straightforward. I wonder if my approach was too complicated, and maybe I misunderstood the best practices.

It's interesting to note that passing an interface as parameter can include (is always) pointer type although it does not look like it in the parameter list (concrete example: `http.ResponseWriter`).

**Link to work:**
- Commits: https://github.com/sebkraemer/golang-fifa-world-cup-web-service/commits/3b068f9f1c43ce0d8204faf7cc9c3f74010389aa
- Project: [golang-fifa-world-cup-web-service@3b068f9 submodule](./projects/)


## Day 31

December 11, 2021

**Today's Progress**

- Implemented the HTTP request handler to make checks and process the input, i.e., solving and returning the solved sudoku data

**Thoughts**

Feels good to code and when tests turn green ✅
Wonder how to architest the whole web service application, incl logging and mature muxer, HTTP API design..
The Reader interface of request body and json decoder, things are falling into place.

**Link to work:**
- Commits: https://github.com/sebkraemer/100-days-of-code/commit/9b003b3b14307572de1a94b2fc2c8b9a56cdb7b6
- Project: [sudoku-go//cmd/sudokuserver](./projects/sudoku-go/cmd/sudokuserver)


## Day 32

December 12, 2021

**Today's Progress**

- Pluralsight course *Creating Web Services with Go* by @AlexCSchultz


**Thoughts**

Two ways to start the http server:
- 'static'
  - http.Handler(...)
  - http.ListenAndServer(port, muxer=nil)
- custom with more control
  - s := http.Server{.., Handler: FooHandler{}}
  - s.ListenAndServer()

I should download and review th course source code, I wondered that one place where the handler was passed as `&handler` (I think), why was it done this way?

**Link to work:**
- Commits: https://github.com/sebkraemer/100-days-of-code/commit/ec98da5da2743b55fccddc2b297b187bb96d099c
- Project: [sudoku-go//cmd/sudokuserver](./projects/sudoku-go/cmd/sudokuserver)


## Day 33

December 13, 2021

**Today's Progress**

- Converted some handler tests to table tests format
- Produced fresh cards
- Some random topic reading, making notes for several things to get back to later.
- buffer, reader, strings usage

**Thoughts**

- I somehow confused httptest.NewServer as something I could use but I didn't need to simulate the server.
- Happy that VS Code supports table tests out of the box.
- Various specific notes were floating around my head, made notes and read spec and blogs. Things to come back to later.
- Surprisingly learned the debug console and watch feature in the "Getting Started" VS Code youtube video.

**Link to work:**
- Commits: https://github.com/sebkraemer/100-days-of-code/commit/0850530ae0edc296eedc9308ad9cf8bf371605e0
- Project: [sudoku-go//cmd/sudokuserver](./projects/sudoku-go/cmd/sudokuserver)


## Day 34

*TODO, get note from twitter*


## Day 34

*TODO, get note from twitter*


## Day 36

December 16, 2021

**Today's Progress**

- Learned about using environment variables inside Go program.
- Learned about crossplatform Go builds with GOOS and GOARCH.
- Learned about Docker multiplatform builds, didn't know that.

**Thoughts**

Couldn't really follow the course regarding the multiplatform docker build because I couldn't get the `buildx` command running with my brewed docker and minikube.
Not saying it's not possible but it's not my focus right now.
It's a bit sad though and probably mainly a konsequence of not having Docker Desktop.
Should I switch to my (kind of neglected) Linux system?

Multi-platform docker gets me in the mood to try that on my (also neglected Raspberry Pi! 😁

**Link to work:**
- Commits: https://github.com/sebkraemer/100-days-of-code/commit/2736c3edeed9f533ff4bbc7ca26f18fdd0e5b719
- Project: [sudoku-go//cmd/sudokuserver](./projects/sudoku-go/cmd/sudokuserver)


## Day 40

December 20, 2021

**Today's Progress**

- probably solved the local module mystery
- added logrus logging to service
- ran example for docker compose, want to utilize it for advanced logging infra
- looked into fluentd, prometheus, ELK

**Link to work:**
- Commits: https://github.com/sebkraemer/100-days-of-code/commit/0767a12064a36d04fa496e46d8aeb80cc1cdc122
- Project: [sudoku-go//cmd/sudokuserver](./projects/sudoku-go/cmd/sudokuserver)


## Day 44

December 24, 2021 🎄

**Today's Progress**

- Goroutines: experiment with done channel pattern and (as I call it) implicit cancellation

**Thoughts**

- Goroutines, finally 😎
- Gotcha! I'm glad I experienced the pitfall of capturing the reference to a loop counter in a closure instead of the value -- better learn that now than later
- Deadlock detection (experienced it) is an interesting feature


**Link to work:**
- Commits: https://github.com/sebkraemer/100-days-of-code/commit/bb7800b202bd27ae2c8a745a3e6a6baf6e29c02e
