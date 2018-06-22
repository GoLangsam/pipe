// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

import "time"

// ===========================================================================
// Beg of intPlugAfter - graceful terminator

// intPlugAfter returns a channel to receive every `inp` before close and a channel to signal this closing.
// Upon receipt of a time signal
// (e.g. from `time.After(...)`),
// output is immediately closed,
// and for graceful termination
// any remaining input is drained before done is signalled.
func intPlugAfter(inp <-chan int, after <-chan time.Time) (out <-chan int, done <-chan struct{}) {
	cha := make(chan int)
	doit := make(chan struct{})
	go plugintAfter(cha, doit, inp, after)
	return cha, doit
}

func plugintAfter(out chan<- int, done chan<- struct{}, inp <-chan int, after <-chan time.Time) {
	defer close(done)

	var end bool // shall we end?
	var ok bool  // did we read successfully?
	var e int    // what we've read

	for !end {
		select {
		case e, ok = <-inp:
			if ok {
				out <- e
			} else {
				end = true
			}
		case <-after:
			end = true
		}
	}

	close(out)

	for range inp {
		// drain inp
	}

	done <- struct{}{}
}

// End of intPlugAfter - graceful terminator