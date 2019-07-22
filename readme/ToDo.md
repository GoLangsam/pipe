# ToDo - all kinds of notes and scribblings

## Consistency:

- pipe/s versus pipe/m - not in m:
	- balance.forever - compare with any/balance
	- bool (empty!)	- see Partition below
	- proc	 // TODO: We do not know whether or not proc closed out

- pipe/s versus pipe/m - unintended differences:
	- internal/pipe-make: comments differ 

	- internal/strew: comments differ
	- internal/merge: mergeThing inp-args differ

- generic/internal/genny.go versus generic/internal/genny.go
	- m/sema is not bundled yet!

## General: 

--- pipe/m/strew/strew.go
Line 52: warning: receiver name static should be consistent with previous receiver name inp for anyThingFrom (golint)
Ignore it?

--- pipe/m/adjust 
-- func anyThingSendProxy(out anyThingInto, sizes ...int) (send anyThingInto) {
=> func (out anyThingInto)anyThingSendProxy(sizes ...int) (send anyThingInto) {
??? no good, may be

- rename sizes
- sendanyThingProxySizes
  - => 'static' method!?!
  - use ?Into - it's more rare ;-)
  - golint will complain :-(

=> methods? How?
- func ThingDaisyChaiN(inp chan Thing, somany int, procs ...func(out ThingInto, inp ThingFrom)) (out chan Thing)
- func ThingDaisyChain(inp chan Thing, procs ...func(out ThingInto, inp ThingFrom)) (out chan Thing)

---
## Partition
github.com\tobyhede\go-underscore\

?PipeTrue: func(item Any) bool => chan - filter function
?ForkBool: func(item Any) bool => chan & chan aka true & fail - discriminator function

---
## => README.md
rfc1925.txt
- (10)  One size never fits all.

--- add `Done`
- FanIn2
- fan-in
- fan-in1
- has-nil

---
## Nomenklatura

- `?Chan`
  - pour
    - gush fill

### => Terminology

- strew - to cover a surface with things
  - scatter spray splash(no) splatter(no)
  - a la fan-out

- Adjust / Adapt / Tailor / Proxy? / Buffer / Stretch / Extend / Expand / 

pour flow pump 

pump/extract/siphon/draw sth from sth

**gush** (of sth)   a large amount of liquid suddenly and quickly flowing or pouring out of sth 
		a gush of blood - blood gushing from a wound 

**seep** - to flow slowly and in small quantities through sth or into sth ( especially of liquids )
Syn: trickle
- Blood was beginning to _seep_ through the bandages. 
- Water _seeped_ from a crack in the pipe. 
- Gradually the pain _seeped_ away. ( figurative )

### Plug, Rupt or Fuse
- stop sending, emit `done` and keep draining
  - detector func(a) (a, bool)
  - incoming stop signal

---
---todo: ????
Note: len(inpS) may be zero! - If > 1 ... ??? / Or nil / one channel?
type Component?	func(inpS ...<-chan anyThing) (outS []<-chan interface{})
type Processor?	func(inpS ...<-chan anyThing) (done <-chan struct{})

---todo: Whirl
aka Rake: Whirl / spin(no) / reel(no)

type WhirlPool? func() // represents the circular network

type Whirl interface {
	Proc(WhirlPool func()) done     // receives the processing network / returns the done channel
./.	Done() <-chan struct{}
Rake =>	Todo(func (item anyThing))	// NO - this is part of the network - after init, before first non-empty Feed(...)
<=	Feed(items ...anyThing)         // for initial food and (via Todo) for feedback
<=	Gone(items ...anyThing)
}


w := New() *Whirl, feed, gone
use feed (&done) to build:
- Your Todo(feed) - feedback, e.g. store it in Your struct
- Your Proc(done) - network
<-w.Proc(proc).Todo(todo).Feed(1,2,3).Done()

Feed:		panic iff todo == nil || proc == nil
Todo & Proc:	panic iff called after non-empty Feed(...)

????: return 3 args, or a struct with .Feed() & .Gone() & Done() ???
Lesson learned: Not a struct with methods! The pointer in the embedding struct (which may be used in proc's feedback) arrives too late.
The problem is feed - the struct may have Gone() & Done() as we do lazy init upon first feed

The network rule: Any Feed(item) *MUST* go either to Todo(item) or to Gone() 

---
[Back to overview](overview.md)