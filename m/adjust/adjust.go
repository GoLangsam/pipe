// adjustments and embedding:
// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// original found in $GOROOT/test/chan/sieve2.go

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

// Note: pipeanyThingAdjust imports "container/ring" for the expanding buffer.
import (
	"container/ring"

	"github.com/cheekybits/genny/generic"
)

// anyThing is the generic type flowing thru the pipe network.
type anyThing generic.Type

// anyOwner is the generic who shall own the methods.
//  Note: Need to use `generic.Number` here as `generic.Type` is an interface and cannot have any method.
type anyOwner generic.Number

// ===========================================================================
// Beg of anyThingPipeAdjust

// anyThingPipeAdjust returns a channel to receive
// all `inp`
// buffered by a anyThingSendProxy process
// before close.
func (my anyOwner) anyThingPipeAdjust(inp <-chan anyThing, sizes ...int) (out <-chan anyThing) {
	cap, que := my.sendanyThingProxySizes(sizes...)
	cha := make(chan anyThing, cap)
	go my.pipeanyThingAdjust(cha, inp, que)
	return cha
}

// anyThingTubeAdjust returns a closure around anyThingPipeAdjust (_, sizes ...int).
func (my anyOwner) anyThingTubeAdjust(sizes ...int) (tube func(inp <-chan anyThing) (out <-chan anyThing)) {

	return func(inp <-chan anyThing) (out <-chan anyThing) {
		return my.anyThingPipeAdjust(inp, sizes...)
	}
}

// End of anyThingPipeAdjust
// ===========================================================================

// ===========================================================================
// Beg of sendanyThingProxy

func (my anyOwner) sendanyThingProxySizes(sizes ...int) (cap, que int) {

	// CAP is the minimum capacity of the buffered proxy channel in `anyThingSendProxy`
	const CAP = 10

	// QUE is the minimum initially allocated size of the circular queue in `anyThingSendProxy`
	const QUE = 16

	cap = CAP
	que = QUE

	if len(sizes) > 0 && sizes[0] > CAP {
		que = sizes[0]
	}

	if len(sizes) > 1 && sizes[1] > QUE {
		que = sizes[1]
	}

	if len(sizes) > 2 {
		panic("anyThingSendProxy: too many sizes")
	}

	return
}

// anyThingSendProxy returns a channel to serve as a sending proxy to 'out'.
// Uses a goroutine to receive values from 'out' and store them
// in an expanding buffer, so that sending to 'out' never blocks.
//  Note: the expanding buffer is implemented via "container/ring"
//
// Note: anyThingSendProxy is kept for the Sieve example
// and other dynamic use to be discovered
// even so it does not fit the pipe tube pattern as anyThingPipeAdjust does.
func (my anyOwner) anyThingSendProxy(out chan<- anyThing, sizes ...int) chan<- anyThing {
	cap, que := my.sendanyThingProxySizes(sizes...)
	cha := make(chan anyThing, cap)
	go my.pipeanyThingAdjust(out, cha, que)
	return cha
}

// pipeanyThingAdjust uses an adjusting buffer to receive from 'inp'
// even so 'out' is not ready to receive yet. The buffer may grow
// until 'inp' is closed and then will shrink by every send to 'out'.
//  Note: the adjusting buffer is implemented via "container/ring"
func (my anyOwner) pipeanyThingAdjust(out chan<- anyThing, inp <-chan anyThing, QUE int) {
	defer close(out)
	n := QUE // the allocated size of the circular queue
	first := ring.New(n)
	last := first
	var c chan<- anyThing
	var e anyThing
	ok := true
	for ok {
		c = out
		if first == last {
			c = nil // buffer empty: disable output
		} else {
			e = first.Value.(anyThing)
		}
		select {
		case e, ok = <-inp:
			if ok {
				last.Value = e
				if last.Next() == first {
					last.Link(ring.New(n)) // buffer full: expand it
					n *= 2
				}
				last = last.Next()
			}
		case c <- e:
			first = first.Next()
		}
	}

	for first != last {
		out <- first.Value.(anyThing)
		first = first.Unlink(1) // first.Next()
	}
}

// End of sendanyThingProxy
// ===========================================================================
