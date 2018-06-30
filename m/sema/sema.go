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
//  ref: database/sql/sql_test.go
//  ref: cmd/compile/internal/gc/noder.go
//
func (inp anyThingFrom) anyThingPipeFuncMany(act func(a anyThing) anyThing, many int) (out anyThingFrom) {
	cha := make(chan anyThing)

	if act == nil { // Make `nil` value useful
		act = func(a anyThing) anyThing { return a }
	}

	if many < 1 {
		many = 1
	}

	go inp.pipeanyThingFuncMany(cha, act, many)
	return cha
}

func (inp anyThingFrom) pipeanyThingFuncMany(out anyThingInto, act func(a anyThing) anyThing, many int) {
	defer close(out)

	sem := make(chan struct{}, many)
	var wg sync.WaitGroup

	for i := range inp {
		sem <- struct{}{}
		wg.Add(1)
		go func(i anyThing) {
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
