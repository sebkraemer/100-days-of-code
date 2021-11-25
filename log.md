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
