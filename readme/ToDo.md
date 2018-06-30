# ToDo

all kinds of notes and scribblings

General: 

--- pipe/m/adjust 
-- func anyThingSendProxy(out anyThingInto, sizes ...int) (send anyThingInto) {
=> func (out anyThingInto)anyThingSendProxy(sizes ...int) (send anyThingInto) {

- rename sizes
- sendanyThingProxySizes => static method!?! - use ?Into - it's more rare ;-)

=> methods? How?
func ThingDaisyChaiN(inp chan Thing, somany int, procs ...func(out ThingInto, inp ThingFrom)) (out chan Thing)
func ThingDaisyChain(inp chan Thing, procs ...func(out ThingInto, inp ThingFrom)) (out chan Thing)

---todo: s/proc

---todo: s/bool & s/true

Partition: github.com\tobyhede\go-underscore\

?PipeTrue: func(item Any) bool => chan - filter function
?ForkBool: func(item Any) bool => chan & chan aka true & fail - discriminator function

---

### 0. Fork:

### 1. Get a local clone your fork:

    git clone git@github.com:YOUR-USERNAME/YOUR-FORKED-REPO.git

### 2. Add the remote original repository as `upstream` to your forked repository: 

    cd into/cloned/fork-repo
    git remote add upstream git://github.com/ORIGINAL-DEV-USERNAME/REPO-YOU-FORKED-FROM.git

### 3. Obtain latest changes from `upstream`
    git fetch upstream

### 4. Update your local clone of the fork from original `upstream` repo to keep up with their changes:

    git pull upstream master

### 5. Update your fork:

    git push	

- https://help.github.com/articles/configuring-a-remote-for-a-fork/
- https://help.github.com/articles/syncing-a-fork/

---

$ go test -bench=.

$ go test -bench=. -cpuprofile cpu.prof
$ go tool pprof -svg cpu.prof > cpu.svg

$ go test -bench=. -trace trace.out
$ go tool trace trace.out

$ go test -race
PASS

$ git gc --aggressive

---

golang.org\x\build\internal\lru\cache.go - a concurrency-safe lru cache

---

rfc1925.txt
- (10)  One size never fits all.

---
fan-in:  also *Done
fan-in1: also *Done
has-nil: also *Done

---
## Nomenklatura

- ?Make*

- ?Chan*		TODO => ?Fill* or ?Pour*
  pour
  - gush fill

- ?Pipe*
  - ?Tube*
- ?Done*
  - ?Fini*

- ?Fork*
- ?Pair*
- ?Fan2*
---

- Join*		added *FunkNok & *FunkErr

---

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

The network rule: Any Feed(item) *MUST* go either go to Todo(item) or to Gone() 

---
[Back to overview](overview.md)