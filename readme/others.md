# Other related works

An incomplete collection of other related works we came accross - with brief notes, subjective notes.

- github.com\bouk\Idiomatic Generics in Go.html
- github.com\bouk\gonerics
  - gonerics.io serves type-specific *.go - given a gist.github - Types are T U V W X Y Z. AST-Rewrite - no comment handling?

- github.com\vladimirvivien\learning-go
  - ch09\
    - pattern7.go - similar basic pattern; word-count histogram
    - sync3 ff nice samples - service start / stop / with chache ...
    - sync5 6 7: nice calculation sample:
      - See https://projecteuler.net/problem=1
      - TODO use as Example

  - ch01/divide.go Euclidean division 
    - http://en.wikipedia.org/wiki/Division_algorithm
    - math/big/int.go



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
