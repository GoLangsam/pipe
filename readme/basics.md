# Basics

_"Function returning channel" is an important idiom._ Rob Pike

---
## Overview:

What comes with a basic "package pipe" (in some file named `pipe.go`)?

Which convention is used to name the functions (resp. methods)?

Note: `?` represents the name of the (generic) type.

### Create & generate
- [Creator](#creator) - ?Make*
- [Generator](#generator) - ?Chan*

---
### Pipe line 'til done
- [Pipe tube](#pipe-tube) - ?Pipe*
- [Finaliser](#finaliser) - ?Done*

---
### Spread & combine 
- [One to Two](#one-to-two) - ?Fork* & ?Pair*
- [Two to One](#two-to-one) - ?Fan2* - a binary fan-in

---
## Details

Functions and their names as they come with any basic "package pipe".

Note: All the following functions (except `?MakeChan`) spawn a process.

Note: For brevity functions output is not mentioned. It's a `?` channel.

### Creator

Some like (or need) "Do-it-yourself":
create a channel and write a generator

- `?MakeChan()` - aka `make(chan ?)` - returns a channel of `?`.

Note: No process is spawned here, of course.

### Generator

- `?Chan...` is the prefix common for generators.

Use some items,
or some slices of items.

- `?Chan(     item ...?)`
- `?ChanSlice(itemS ...[]?)`

Use some function 
which produces items together with an indication, typically a `bool` or an `error`.

- `?ChanFuncNok(func() ?, bool)`
- `?ChanFuncErr(func() ?, error)`

Typical candidates for such functions are
- some iterator like `Next() (item ?, ok bool)`
- some scanner busy scanning something 'till it encounters an error (e.g. `io.EOF`).

### Pipe tube

- `?Pipe...` is the prefix common for pipe tubes.

Please remember: These return one receive-only channel (named `out`).

- `?Pipe(     func(item ?)   ...)` - apply operations (such as print/log)
- `?PipeFunc( func(item ?) ? ...)` - chain actions which return an item for an item

Mind You: [Batteries](batteries.md) are included
with many special purpose pipe-tubes.

### Finaliser

- `?Done...` is the prefix common for finalisers.

Please remember: These return a signal channel (named `done`), not a `?` channel.

- `?Done(     func(item ?)   ...)` - apply operations (such as print/log) 'till close
- `?DoneFunc( func(item ?) ? ...)` - chain actions 'till close
- `?DoneSlice()` - returns a slice of all items received

Mind You: [Batteries](batteries.md) are included 
with special purpose finalisers such as `freq`.

### One to Two

These return **two** channels (not just a single one, as usual):

- `?Fork` - send to any one (of the two)
- `?Pair` - send to both (in lockstep)

Notice the difference:
- `?Fork` enables two identical processes each to process part of the total workload concurrently.
- `?Pair` enables two different futures: two roads -typically different ones- to be pursued concurrently.

Use several times to spread out. Or see below. 

### Two to One

These accept **two** channels (and return one, as usual):

- `?Fan2` - a binary fan-in.

Notice that `?Fan2` is kind of a closing bracket where `?Fork` is the opening one.

### One to Many & Many to One

Sure, two is not enough - more is needed.

Well, this is beyond basics, is it not?

Mind You: [Batteries](batteries.md) are included, and there are "the big brothers':
- `?Fork` becomes `?Strew` 
- `?Pair` becomes `?FanOut` 
- `?Fan2` becomes `?FanIn` or `?FanIn1`

And there is more, such as `?Merge` or `?Same`.

---
## [Closures](closures.md)

- `?Tube...` - closures around ?Pipe*
- `?Fini...` - closures around ?Done*

---
[Back to overview](overview.md)
