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

// anyOwner is the generic who shall own the methods.
//  Note: Need to use `generic.Number` here as `generic.Type` is an interface and cannot have any method.
type anyOwner generic.Number

// ===========================================================================
// Beg of anyThingFanIn

// anyThingFanIn returns a channel to receive all inputs arriving
// on variadic inps
// before close.
//
//  Ref: https://blog.golang.org/pipelines
//  Ref: https://github.com/QuentinPerez/go-stuff/channel/Fan-out-Fan-in/main.go
func (my anyOwner) anyThingFanIn(inps ...<-chan anyThing) (out <-chan anyThing) {
	cha := make(chan anyThing)

	wg := new(sync.WaitGroup)
	wg.Add(len(inps))

	go my.fanInanyThingWaitAndClose(cha, wg) // Spawn "close(out)" once all inps are done

	for i := range inps {
		go my.fanInanyThing(cha, inps[i], wg) // Spawn "output(c)"s
	}

	return cha
}

func (my anyOwner) fanInanyThing(out chan<- anyThing, inp <-chan anyThing, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range inp {
		out <- i
	}
}

func (my anyOwner) fanInanyThingWaitAndClose(out chan<- anyThing, wg *sync.WaitGroup) {
	wg.Wait()
	close(out)
}

// End of anyThingFanIn
// ===========================================================================
