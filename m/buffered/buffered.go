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
// Beg of PipeBuffered - a buffered channel with capacity `cap` to receive

// PipeBuffered returns a buffered channel with capacity `cap` to receive
// all `inp`
// before close.
func (inp anyThingFrom) PipeBuffered(cap int) (out anyThingFrom) {
	cha := make(chan anyThing, cap)
	go inp.pipeBuffered(cha)
	return cha
}

func (inp anyThingFrom) pipeBuffered(out anyThingInto) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// TubeBuffered returns a closure around PipeBuffer (cap).
func (inp anyThingFrom) TubeBuffered(cap int) (tube func(inp anyThingFrom) (out anyThingFrom)) {

	return func(inp anyThingFrom) (out anyThingFrom) {
		return inp.PipeBuffered(cap)
	}
}

// End of PipeBuffered - a buffered channel with capacity `cap` to receive
// ===========================================================================
