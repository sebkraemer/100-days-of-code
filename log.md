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
