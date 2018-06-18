// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

import (
	"github.com/cheekybits/genny/generic"
)

// Any is the generic type flowing thru the pipe network.
type Any generic.Type

// ===========================================================================
// Beg of PipeAnyDone

// PipeAnyDone returns a channel to receive every `inp` before close and a channel to signal this closing.
func PipeAnyDone(inp <-chan Any) (out <-chan Any, done <-chan struct{}) {
	cha := make(chan Any)
	doit := make(chan struct{})
	go pipeAnyDone(cha, doit, inp)
	return cha, doit
}

func pipeAnyDone(out chan<- Any, done chan<- struct{}, inp <-chan Any) {
	defer close(out)
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// End of PipeAnyDone
// ===========================================================================
