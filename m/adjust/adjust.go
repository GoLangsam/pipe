// adjustments and embedding:
// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// original found in $GOROOT/test/chan/sieve2.go

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

import (
	"container/ring"

	"github.com/cheekybits/genny/generic"
)

// Note: pipeAdjust imports "container/ring" for the expanding buffer.

// anyThing is the generic type flowing thru the pipe network.
type anyThing generic.Type

// ===========================================================================
// Beg of PipeAdjust

// PipeAdjust returns a channel to receive
// all `inp`
// buffered by a anyThingSendProxy process
// before close.
func (inp anyThingFrom) PipeAdjust(sizes ...int) (out anyThingFrom) {
	cap, que := sendanyThingProxySizes(sizes...)
	cha := make(chan anyThing, cap)
	go inp.pipeAdjust(cha, que)
	return cha
}

// TubeAdjust returns a closure around PipeAdjust (_, sizes ...int).
func (inp anyThingFrom) TubeAdjust(sizes ...int) (tube func(inp anyThingFrom) (out anyThingFrom)) {

	return func(inp anyThingFrom) (out anyThingFrom) {
		return inp.PipeAdjust(sizes...)
	}
}

// End of PipeAdjust
// ===========================================================================

// ===========================================================================
// Beg of sendanyThingProxy

func sendanyThingProxySizes(sizes ...int) (cap, que int) {

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
// even so it does not fit the pipe tube pattern as PipeAdjust does.
func anyThingSendProxy(out anyThingInto, sizes ...int) (send anyThingInto) {
	cap, que := sendanyThingProxySizes(sizes...)
	cha := make(chan anyThing, cap)
	go (anyThingFrom)(cha).pipeAdjust(out, que)
	return cha
}

// pipeAdjust uses an adjusting buffer to receive from 'inp'
// even so 'out' is not ready to receive yet. The buffer may grow
// until 'inp' is closed and then will shrink by every send to 'out'.
//  Note: the adjusting buffer is implemented via "container/ring"
func (inp anyThingFrom) pipeAdjust(out anyThingInto, QUE int) {
	defer close(out)
	n := QUE // the allocated size of the circular queue
	first := ring.New(n)
	last := first
	var c anyThingInto
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
