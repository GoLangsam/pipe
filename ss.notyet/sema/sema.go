// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

import (
	"sync"

	"github.com/cheekybits/genny/generic"
)

// anyThing is the generic type flowing thru the pipe network.
type anyThing generic.Type

// ===========================================================================
// Beg of anyThingSema - limited parallel execution

// anyThingPipeFuncMany returns a channel to receive
// every result of action `act` applied to `inp`
// by `many` parallel processing goroutines
// before close.
//
// anyThingPipeFuncMany - a parallel processing anyThingPipeFunc
//
//  ref: cmd/compile/internal/gc/noder.go
//
// Note: There is no need for a `WaitGroup` in `noder.go`.
// Each `noder` carries another channel `err`
// which will be closed upon completition
// and thus can safely be ranged over
// in that `range noders` loop which follows
// the spawning of the parallel processing goroutines.
// Another useful idiom.
//
func anyThingPipeFuncMany(inp <-chan anyThing, act func(a anyThing) anyThing, many int) (out <-chan anyThing) {
	cha := make(chan anyThing)

	if act == nil { // Make `nil` value useful
		act = func(a anyThing) anyThing { return a }
	}

	if many < 1 {
		many = 1
	}

	go pipeanyThingFuncMany(cha, inp, act, many)
	return cha
}

func pipeanyThingFuncMany(out chan<- anyThing, inp <-chan anyThing, act func(a anyThing) anyThing, many int) {
	defer close(out)

	sem := make(chan struct{}, many)
	var wg sync.WaitGroup

	for i := range inp {
		wg.Add(1)
		go func(i anyThing) {
			sem <- struct{}{}
			defer func() {
				<-sem
				wg.Done()
			}()
			out <- act(i) // apply action
		}(i)
	}
	wg.Wait()
}

// End of anyThingSema - limited parallel execution
// ===========================================================================
