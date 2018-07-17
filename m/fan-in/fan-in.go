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
// Beg of FanIn

// FanIn returns a channel to receive all inputs arriving
// on variadic inps
// before close.
//
//  Note: For each input one go routine is spawned to forward arrivals.
//
// See FanIn1 in `fan-in1` for another implementation.
//
//  Ref: https://blog.golang.org/pipelines
//  Ref: https://github.com/QuentinPerez/go-stuff/channel/Fan-out-Fan-in/main.go
func (inp anyThingFrom) FanIn(inps ...anyThingFrom) (out anyThingFrom) {
	cha := make(chan anyThing)

	wg := new(sync.WaitGroup)
	wg.Add(len(inps) + 1)

	go inp.fanInWaitAndClose(cha, wg) // Spawn "close(out)" once all inps are done

	go inp.fanIn(cha, wg)
	for i := range inps {
		go inps[i].fanIn(cha, wg) // Spawn "output(c)"s
	}

	return cha
}

func (inp anyThingFrom) fanIn(out anyThingInto, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range inp {
		out <- i
	}
}

func (inp anyThingFrom) fanInWaitAndClose(out anyThingInto, wg *sync.WaitGroup) {
	wg.Wait()
	close(out)
}

// End of FanIn
// ===========================================================================
