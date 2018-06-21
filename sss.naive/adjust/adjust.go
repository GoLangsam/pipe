// minimal adjustments:
// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// original found in $GOROOT/test/chan/sieve2.go

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

// Note: anyThingSendProxy imports "container/ring" for the expanding buffer.
import (
	"container/ring"

	"github.com/cheekybits/genny/generic"
)

// anyThing is the generic type flowing thru the pipe network.
type anyThing generic.Type

// ===========================================================================
// Beg of anyThingSendProxy

// BufferanyThingCAP is the capacity of the buffered proxy channel in `anyThingSendProxy`
const BufferanyThingCAP = 10

// BufferanyThingQUE is the allocated size of the circular queue in `anyThingSendProxy`
const BufferanyThingQUE = 16

// anyThingSendProxy returns a channel to serve as a sending proxy to 'out'.
// Uses a goroutine to receive values from 'out' and store them
// in an expanding buffer, so that sending to 'out' never blocks.
//  Note: the expanding buffer is implemented via "container/ring"
func anyThingSendProxy(out chan<- anyThing) chan<- anyThing {
	proxy := make(chan anyThing, BufferanyThingCAP)
	go func() {
		n := BufferanyThingQUE // the allocated size of the circular queue
		first := ring.New(n)
		last := first
		var c chan<- anyThing
		var e anyThing
		for {
			c = out
			if first == last {
				// buffer empty: disable output
				c = nil
			} else {
				e = first.Value.(anyThing)
			}
			select {
			case e = <-proxy:
				last.Value = e
				if last.Next() == first {
					// buffer full: expand it
					last.Link(ring.New(n))
					n *= 2
				}
				last = last.Next()
			case c <- e:
				first = first.Next()
			}
		}
	}()
	return proxy
}

// End of anyThingSendProxy
// ===========================================================================
