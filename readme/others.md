# Other related works

An incomplete collection of other related works we came accross - with brief notes, subjective notes.

- [Idiomatic Generics in Go](https://github.com/bouk/Idiomatic Generics in Go.html)
- [github.com/bouk/gonerics](https://github.com/bouk/gonerics/)
  - gonerics.io serves a type-specific *.go file - given a 'gist.github'
    - Types are T U V W X Y Z.
    - Uses AST-Rewrite
    - No(?) comment handling

- [github.com/vladimirvivien/learning-go](https://github.com/vladimirvivien/learning-go)
  - ch09/
    - pattern7.go - similar basic pattern; word-count histogram
    - sync3 ff nice samples - service start / stop / with cache ...
    - sync5 6 7: nice [calculation sample](https://projecteuler.net/problem=1).

  - ch01/divide.go [Euclidean division](http://en.wikipedia.org/wiki/Division_algorithm) - see also `math/big/int.go`

- [github.com/urjitbhatia/gopipe](https://github.com/urjitbhatia/gopipe/)
  - `PipeLine` as a control stuct
  - one and only one "head" (aka "generator")
  - one and only one "tail"
  - `p.Enqueue(item interface{})` and `p.Dequeue() interface{}` as functions
  - channels are not exposed / returned
  - based on what we call `Proc` - pipe functions which
  - have to fullfill interface `Process(in chan interface{}, out chan interface{})`
  - no type safety - `interface{}` means everything and nothing
  - no directional safety (send-only / receive-only)
  - oddities such as uncontrolled `p.Close()` and `p.AttachSink(out chan interface{})` (aka drain / `?Done` - but without signal)
  + `DebugMode`

- [github.com/myntra/pipeline](https://github.com/myntra/pipeline)
  - a PipeLine has a squence of Stages
  - any Stage has a collection of Steps - sequential / concurrent
  - Stage and PipeLine act as flow governors.
  - Step is where the work is done.
  - no type safety - `interface{}` means everything and nothing

---
[Back to overview](overview.md)
