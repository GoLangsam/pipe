# Resources
Some links to related informations.

## Books
- [Communicating Sequential Processes (CSP)](http://www.usingcsp.com/cspbook.pdf) **The** CSPbook
    * Author: C.A.R Hoare
    * Publication Date: May 18, 2015 - first published in 1985

- **The Go Programming Language** *(Addison-Wesley Professional Computing Series)*
    * Author: Alan A.A. Donovan and Brian Kernighan
    * Publication Date: November, 2015
    * ISBN: 978-0134190440
    * Reference: http://www.gopl.io/

## [blog.golang](https://blog.golang.org/) - Articles
- [Pipelines](https://blog.golang.org/pipelines)
- [Share Memory By Communicating](https://blog.golang.org/share-memory-by-communicating)
- [Concurrency is not parallelism](http://blog.golang.org/concurrency-is-not-parallelism)
- [Go Concurrency Patterns: Timing out, moving on](https://blog.golang.org/go-concurrency-patterns-timing-out-and)
- [Go Concurrency Patterns: Context](https://blog.golang.org/context)

## [YouTube](http://www.youtube.com/) - Videos
- [Go Concurrency Patterns](http://www.youtube.com/watch?v=f6kdp27TYZs)
- [Advanced Go Concurrency Patterns](http://www.youtube.com/watch?v=QDDwwePbDtw)

## other blogs
- [Resources for new go programmers](https://dave.cheney.net/resources-for-new-go-programmers) by [Dave Cheney](https://dave.cheney.net/)

- [Golang channels tutorial](http://guzalexander.com/2013/12/06/golang-channels-tutorial.html) by [Alexander Guz](http://guzalexander.com/)

## further readings
- [a laundry list of common mistakes](https://github.com/golang/go/wiki/CodeReviewComments), not a style guide.
- [package names](http://golang.org/doc/effective_go.html#package-names)

- [Golang Internals Part 2: Nice benefits of named return values](https://blog.minio.io/golang-internals-part-2-nice-benefits-of-named-return-values-1e95305c8687)

- [Go by Example: Channels](https://gobyexample.com/channels)

  "*Channels* are the pipes that connect concurrent goroutines. You can send values into channels from one goroutine and receive those values into another goroutine."
  "Channels are typed by the values they convey."

- [Go101](http://Go101.org/ "Go101")

  Includes interesting stuff such as basic concurrency patterns. 

---
- [Simple Data Processing Pipeline with Golang](https://www.hugopicado.com/2016/09/26/simple-data-processing-pipeline-with-golang.html) by [Hugo Picado](https://www.hugopicado.com/) Sep. 26, 2016
- [Sources](https://github.com/picadoh/gostreamer)

  "In this example we are building a simple processing pipeline that consumes a text line from a socket and sends it through a series of processes to extract independent words, filter the ones starting with # and printing the result to the console. For this, a set of structures and functions were created so we can try around and build other kind of pipelines at will."

  - Has a `Collector.Execute` as `Fan-In(cap=1)` and a `Processor.Execute`,and a `ChannelDemux.Execute` for non-random FanOut.
  - Uses `type ProcessFunction func(name string, input Message, out chan Message)`
  - Code flavour is `ssss`

---
- [Fan-out-Fan-in/package](https://go.hotlibs.com/github.com/QuentinPerez/go-stuff/channel/Fan-out-Fan-in/package)
  
  [repo](github.com/QuentinPerez/go-stuff/channel/Fan-out-Fan-in)

  `merge` - a Fan-In, not -Out! (with sync.WaitGroup for closer, and 'done' as context)

  *Warning*: These are just misnamed copies:
  - [Fan-out example](https://gist.github.com/mchirico/df9fad3e7a5ea0c4527a)
  - [same](https://www.snip2code.com/Snippet/1043022/Go-(Golang)-Fan-out-example/)

  - [Fan-out-Fan-in/package](https://go.hotlibs.com/github.com/QuentinPerez/go-stuff/channel/Fan-out-Fan-in/package)
    [repo](github.com/QuentinPerez/go-stuff/channel/Fan-out-Fan-in)

- [Channels Are Not Enough](https://gist.github.com/kachayev/21e7fe149bc5ae0bd878) ... or Why Pipelining Is Not That Easy -
by [@kachayev](https://twitter.com/kachayev)
  - Unicorn Cartoon :-)
  - Fan-In sample from above, "(shamelessly stolen from [here](http://blog.golang.org/pipelines))" 
  - Delves into "channel is a functor" and "Futures & Promises", but does not distinguish supply and demand (but uses it)

- [Buffered Channels In Go — What Are They Good For?](https://medium.com/capital-one-developers/buffered-channels-in-go-what-are-they-good-for-43703871828) by [Jon Bodner](https://medium.com/@jon_43067)

  Buffered channels are useful when you know how many goroutines you have launched, want to limit the number of goroutines you will launch, or want to limit the amount of work that is queued up.

  - Parallel Processing
  - Creating a Pool

- [Abundant concurrency in Go](https://hunterloftis.github.io/2017/07/14/abundant-concurrency/)

  Contrasts his 'JavaScripter’s mindset' with an 'abundance mindset'. 

- [Buffered Channels](https://medium.com/capital-one-developers/buffered-channels-in-go-what-are-they-good-for-43703871828)

---
[Back to overview](overview.md)
