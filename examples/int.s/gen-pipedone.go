// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

// ===========================================================================
// Beg of intPipeDone

// intPipeDone returns a channel to receive every `inp` before close and a channel to signal this closing.
func intPipeDone(inp <-chan int) (out <-chan int, done <-chan struct{}) {
	cha := make(chan int)
	doit := make(chan struct{})
	go pipeIntDone(cha, doit, inp)
	return cha, doit
}

func pipeIntDone(out chan<- int, done chan<- struct{}, inp <-chan int) {
	defer close(out)
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// End of intPipeDone
// ===========================================================================
