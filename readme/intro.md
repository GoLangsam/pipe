# Intro

_"Function returning channel" is an important idiom._ Rob Pike

A process with only two channels, namely an input channel `into` and an output channel `from` is called a "pipe",
_(Chapter 4.4 of C.A.R. Hoare's "CSP-Book - "Communicating Sequential Processes" - see [resources](resources.md)._

Initially, we shall confine our attention to such processes and their composition.

- Note: C.A.R. Hoare uses a left-to-right notation: `P >> Q`

  Go uses a right-to-left notation: `Q(into) <- <-P(from)`
  (due to some subconcious _arabic_ influence, may be).

Initially, we shall further confine our attention to "homogeneous" pipe networks - composed of processes operating on one common type.

`?` shall represent the name of this (generic) type.

Instances of `?` "flow" along channels.
(Often any such instance is called an `item`, or sometimes simply `a` or `i`.)

Channels connect pipes. Thus: they become the glue to build some network composed of pipes.

So, let's see what we need and what [we get](basics.md) to build:

## A "pipe" - a middle piece

"Pipe" is the central component.

We model a "pipe" as a `func ( <-chan ) <-chan`:
- receive a channel
- return a channel

Or - a little more specific:
- receive some receive-only channel
  - and some more arguments, may be
- return (another) receive-only channel

Such "pipe" is a concurrent process:
- it spawns a goroutine when invoked, which
  - upon receive of some `?`
  - does some magic which
  - sends none, one or several `?`
  on the outgoing channel.

### A "pipeline" - a longer middle piece

Compose into a "pipeline":
What one "pipe" returns becomes argument of another "pipe":

`c-pipe(b-pipe(a-pipe( ... )))` - function composition - right-to-left

`...a-pipe().b-pipe().c-pipe()` - method chaining - left-to-right

### A beginning: "chan"

Well, to begin any pipeline. it needs a channel from which some `?` may be received.

A process doing such is often called "generator" or "source".

### An end: "done"

"Anything has an end - only a hotdog has two." (a non-serious German saying)

Any endpoint of a pipeline always returns a signal channel - usually named `done`.

Once the flow subsides, (which means: the incoming channel gets closed):
- one (and only one) value is sent on that `done` channel,
- and `done` is closed in order to broadcast: flow has subsided
- and function in the spawned process terminates
- and so does the goroutine - it does not leak.

Some authors use what they call a "sink": something flows into the "sink", nothing comes out from such "sink".
We intentionally avoid such - we want to know when the flow has subsided.

---
[Back to overview](overview.md)
