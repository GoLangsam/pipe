// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

import "sync"

// ===========================================================================
// Beg of ThingFanIn

// ThingFanIn returns a channel to receive all inputs arriving
// on variadic inps
// before close.
//
//  Note: For each input one go routine is spawned to forward arrivals.
//
// See ThingFanIn1 in `fan-in1` for another implementation.
//
//  Ref: https://blog.golang.org/pipelines
//  Ref: https://github.com/QuentinPerez/go-stuff/channel/Fan-out-Fan-in/main.go
func ThingFanIn(inps ...<-chan Thing) (out <-chan Thing) {
	cha := make(chan Thing)

	wg := new(sync.WaitGroup)
	wg.Add(len(inps))

	go fanInThingWaitAndClose(cha, wg) // Spawn "close(out)" once all inps are done

	for i := range inps {
		go fanInThing(cha, inps[i], wg) // Spawn "output(c)"s
	}

	return cha
}

func fanInThing(out chan<- Thing, inp <-chan Thing, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range inp {
		out <- i
	}
}

func fanInThingWaitAndClose(out chan<- Thing, wg *sync.WaitGroup) {
	wg.Wait()
	close(out)
}

// End of ThingFanIn
// ===========================================================================
