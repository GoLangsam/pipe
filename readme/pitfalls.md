# Pitfalls and popular easy-to-do mistakes - vaccines and protective habits

Note: As of now this collection is in it's infancy. Your kind understanding is appreciated.


## Ownership and sharing

Go does not do too much to enforce a discipline of ownership.

This opens a can of worms - and go leaves it to the skill and discipline of the programmer to avoid worms becoming nasty bugs.

### `go vet -copylocks`
Yes, `go vet -copylocks` can detect disallowed copying of values who implement `Lock()`,
such as the mutexes of the std `sync` package.

Note: Pointers to such values are safe to copy - 

`sync` also uses an internal `noCopy` struct in order to protect other concurrency primitives such as
"condition variables", any "Pool - a set of temporary objects" and any "WaitGroup"

https://golang.org/issues/8005#issuecomment-190753527


### Who owns the duty (or privilege) to close this channel?

This duty may be left to the garbage collector. In rare cases, this is the only solution available.

Just: this is lazy thinking, and most likely results in dangerous buggy implementation:
it's all too easy to overlook (and thus: leak) some go routine working on such channel.

In the context of our pipe implementation here, this important question has a simple answer:

> The function which provided the channel (and its associated processing) has to close the provided channel(s).

This simplicity is an advantage of the fact that we 'follow the flow': functions implement a process and supply some read-only channel(s).

## Blog posts and articles

### `case <-done` - The Ordering Trap

A Pitfall in `select` re order between `case <-inp` and `case <-done`

NoGood:

```go
	func (f *Foo) Number() (int, error) {
	    select {
	    case <-f.doneChan:
	        return 0, errors.New("canceled")
	    case <-f.numberChan:
	        return f.number, nil // Deliver result
	    }
	}
```

Better: Deliver also upon done, if available

```go
	func (f *Foo) Number() (int, error) {
	    select {
	    case <-f.doneChan:
	        select {
	        case <-f.numberChan:
	            return f.number, nil // Deliver result
	        default:
	            return 0, errors.New("canceled")
	        }
	    case <-f.numberChan:
	        return f.number, nil // Deliver result
	    }
	}
```

[Notifications on the channels in Golang](http://blog.atte.ro/2017/07/09/golang-channel-notification-select.html) - The Ordering Trap

---
### Strong typing - Lesson learned

"Imagine you have a toy, and the object of the toy is to fill it by passing its contents through appropriate slots.
Ideally, you’d want the toy to tell you that you’re doing it wrong by using different shapes for the different objects and holes.
Instead, with CockroachDB, all the objects and holes were shaped the same and the rule was instead to “pay attention to the color.”
What’s worse is that Go didn’t help us realize we were also color blind.
A schrodinbug in CockroachDB!.

**Strong(er) typing really helped us put things back in shape.**

We’ll do it more from now on."

[Squashing a Schroedinbug with strong typing](https://www.cockroachlabs.com/blog/squashing-a-schroedinbug-with-strong-typing/ "Cockroach Labs")

---
[Back to overview](overview.md)
