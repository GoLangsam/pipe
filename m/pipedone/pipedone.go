// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

import (
	"github.com/cheekybits/genny/generic"
)

// anyThing is the generic type flowing thru the pipe network.
type anyThing generic.Type

// ===========================================================================
// Beg of PipeDone

// PipeDone returns a channel to receive every `inp` before close and a channel to signal this closing.
func (inp anyThingFrom) PipeDone() (out anyThingFrom, done <-chan struct{}) {
	cha := make(chan anyThing)
	doit := make(chan struct{})
	go inp.pipeDone(cha, doit)
	return cha, doit
}

func (inp anyThingFrom) pipeDone(out anyThingInto, done chan<- struct{}) {
	defer close(out)
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// End of PipeDone
// ===========================================================================
