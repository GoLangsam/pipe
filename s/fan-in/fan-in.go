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
// Beg of anyThingFanIn

// anyThingFanIn returns a channel to receive all inputs arriving
// on variadic inps
// before close.
//
//  Note: For each input one go routine is spawned to forward arrivals.
//
// See anyThingFanIn1 in `fan-in1` for another implementation.
//
//  Ref: https://blog.golang.org/pipelines
//  Ref: https://github.com/QuentinPerez/go-stuff/channel/Fan-out-Fan-in/main.go
func anyThingFanIn(inps ...<-chan anyThing) (out <-chan anyThing) {
	cha := make(chan anyThing)

	wg := new(sync.WaitGroup)
	wg.Add(len(inps))

	go fanInanyThingWaitAndClose(cha, wg) // Spawn "close(out)" once all inps are done

	for i := range inps {
		go fanInanyThing(cha, inps[i], wg) // Spawn "output(c)"s
	}

	return cha
}

func fanInanyThing(out chan<- anyThing, inp <-chan anyThing, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range inp {
		out <- i
	}
}

func fanInanyThingWaitAndClose(out chan<- anyThing, wg *sync.WaitGroup) {
	wg.Wait()
	close(out)
}

// End of anyThingFanIn
// ===========================================================================
