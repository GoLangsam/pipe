// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by golang.org/x/tools/cmd/bundle. DO NOT EDIT.

package sites

// ===========================================================================
// Beg of SiteMake creators

// SiteMakeChan returns a new open channel
// (simply a 'chan Site' that is).
//
// Note: No 'Site-producer' is launched here yet! (as is in all the other functions).
//  This is useful to easily create corresponding variables such as:
//
// var mySitePipelineStartsHere := SiteMakeChan()
// // ... lot's of code to design and build Your favourite "mySiteWorkflowPipeline"
// 	// ...
// 	// ... *before* You start pouring data into it, e.g. simply via:
// 	for drop := range water {
// mySitePipelineStartsHere <- drop
// 	}
// close(mySitePipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for SitePipeBuffer) the channel is unbuffered.
//
func SiteMakeChan() (out chan Site) {
	return make(chan Site)
}

// End of SiteMake creators
// ===========================================================================

// ===========================================================================
// Beg of SiteChan producers

// SiteChan returns a channel to receive
// all inputs
// before close.
func SiteChan(inp ...Site) (out <-chan Site) {
	cha := make(chan Site)
	go chanSite(cha, inp...)
	return cha
}

func chanSite(out chan<- Site, inp ...Site) {
	defer close(out)
	for i := range inp {
		out <- inp[i]
	}
}

// SiteChanSlice returns a channel to receive
// all inputs
// before close.
func SiteChanSlice(inp ...[]Site) (out <-chan Site) {
	cha := make(chan Site)
	go chanSiteSlice(cha, inp...)
	return cha
}

func chanSiteSlice(out chan<- Site, inp ...[]Site) {
	defer close(out)
	for i := range inp {
		for j := range inp[i] {
			out <- inp[i][j]
		}
	}
}

// SiteChanFuncNok returns a channel to receive
// all results of generator `gen`
// until `!ok`
// before close.
func SiteChanFuncNok(gen func() (Site, bool)) (out <-chan Site) {
	cha := make(chan Site)
	go chanSiteFuncNok(cha, gen)
	return cha
}

func chanSiteFuncNok(out chan<- Site, gen func() (Site, bool)) {
	defer close(out)
	for {
		res, ok := gen() // generate
		if !ok {
			return
		}
		out <- res
	}
}

// SiteChanFuncErr returns a channel to receive
// all results of generator `gen`
// until `err != nil`
// before close.
func SiteChanFuncErr(gen func() (Site, error)) (out <-chan Site) {
	cha := make(chan Site)
	go chanSiteFuncErr(cha, gen)
	return cha
}

func chanSiteFuncErr(out chan<- Site, gen func() (Site, error)) {
	defer close(out)
	for {
		res, err := gen() // generate
		if err != nil {
			return
		}
		out <- res
	}
}

// End of SiteChan producers
// ===========================================================================

// ===========================================================================
// Beg of SitePipe functions

// SitePipe
// will apply every `op` to every `inp` and
// returns a channel to receive
// each `inp`
// before close.
//
// Note: For functional people,
// this 'could' be named `SiteMap`.
// Just: 'map' has a very different meaning in go lang.
func SitePipe(inp <-chan Site, ops ...func(a Site)) (out <-chan Site) {
	cha := make(chan Site)
	go pipeSite(cha, inp, ops...)
	return cha
}

func pipeSite(out chan<- Site, inp <-chan Site, ops ...func(a Site)) {
	defer close(out)
	for i := range inp {
		for _, op := range ops {
			if op != nil {
				op(i) // chain action
			}
		}
		out <- i // send it
	}
}

// SitePipeFunc
// will chain every `act` to every `inp` and
// returns a channel to receive
// each result
// before close.
func SitePipeFunc(inp <-chan Site, acts ...func(a Site) Site) (out <-chan Site) {
	cha := make(chan Site)
	go pipeSiteFunc(cha, inp, acts...)
	return cha
}

func pipeSiteFunc(out chan<- Site, inp <-chan Site, acts ...func(a Site) Site) {
	defer close(out)
	for i := range inp {
		for _, act := range acts {
			if act != nil {
				i = act(i) // chain action
			}
		}
		out <- i // send result
	}
}

// End of SitePipe functions
// ===========================================================================

// ===========================================================================
// Beg of SiteTube closures around SitePipe

// SiteTube returns a closure around PipeSite (_, ops...).
func SiteTube(ops ...func(a Site)) (tube func(inp <-chan Site) (out <-chan Site)) {

	return func(inp <-chan Site) (out <-chan Site) {
		return SitePipe(inp, ops...)
	}
}

// SiteTubeFunc returns a closure around PipeSiteFunc (_, acts...).
func SiteTubeFunc(acts ...func(a Site) Site) (tube func(inp <-chan Site) (out <-chan Site)) {

	return func(inp <-chan Site) (out <-chan Site) {
		return SitePipeFunc(inp, acts...)
	}
}

// End of SiteTube closures around SitePipe
// ===========================================================================

// ===========================================================================
// Beg of SiteDone terminators

// SiteDone
// will apply every `op` to every `inp` and
// returns a channel to receive
// one signal
// upon close.
func SiteDone(inp <-chan Site, ops ...func(a Site)) (done <-chan struct{}) {
	sig := make(chan struct{})
	go doneSite(sig, inp, ops...)
	return sig
}

func doneSite(done chan<- struct{}, inp <-chan Site, ops ...func(a Site)) {
	defer close(done)
	for i := range inp {
		for _, op := range ops {
			if op != nil {
				op(i) // apply operation
			}
		}
	}
	done <- struct{}{}
}

// SiteDoneFunc
// will chain every `act` to every `inp` and
// returns a channel to receive
// one signal
// upon close.
func SiteDoneFunc(inp <-chan Site, acts ...func(a Site) Site) (done <-chan struct{}) {
	sig := make(chan struct{})
	go doneSiteFunc(sig, inp, acts...)
	return sig
}

func doneSiteFunc(done chan<- struct{}, inp <-chan Site, acts ...func(a Site) Site) {
	defer close(done)
	for i := range inp {
		for _, act := range acts {
			if act != nil {
				i = act(i) // chain action
			}
		}
	}
	done <- struct{}{}
}

// SiteDoneSlice returns a channel to receive
// a slice with every Site received on `inp`
// upon close.
//
// Note: Unlike SiteDone, SiteDoneSlice sends the fully accumulated slice, not just an event, once upon close of inp.
func SiteDoneSlice(inp <-chan Site) (done <-chan []Site) {
	sig := make(chan []Site)
	go doneSiteSlice(sig, inp)
	return sig
}

func doneSiteSlice(done chan<- []Site, inp <-chan Site) {
	defer close(done)
	slice := []Site{}
	for i := range inp {
		slice = append(slice, i)
	}
	done <- slice
}

// End of SiteDone terminators
// ===========================================================================

// ===========================================================================
// Beg of SiteFini closures

// SiteFini returns a closure around `SiteDone(_, ops...)`.
func SiteFini(ops ...func(a Site)) func(inp <-chan Site) (done <-chan struct{}) {

	return func(inp <-chan Site) (done <-chan struct{}) {
		return SiteDone(inp, ops...)
	}
}

// SiteFiniFunc returns a closure around `SiteDoneFunc(_, acts...)`.
func SiteFiniFunc(acts ...func(a Site) Site) func(inp <-chan Site) (done <-chan struct{}) {

	return func(inp <-chan Site) (done <-chan struct{}) {
		return SiteDoneFunc(inp, acts...)
	}
}

// SiteFiniSlice returns a closure around `SiteDoneSlice(_)`.
func SiteFiniSlice() func(inp <-chan Site) (done <-chan []Site) {

	return func(inp <-chan Site) (done <-chan []Site) {
		return SiteDoneSlice(inp)
	}
}

// End of SiteFini closures
// ===========================================================================

// ===========================================================================
// Beg of SitePair functions

// SitePair returns a pair of channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func SitePair(inp <-chan Site) (out1, out2 <-chan Site) {
	cha1 := make(chan Site)
	cha2 := make(chan Site)
	go pairSite(cha1, cha2, inp)
	return cha1, cha2
}

/* not used - kept for reference only.
func pairSite ( out1 , out2 chan <- Site , inp <- chan Site ) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
} */

func pairSite(out1, out2 chan<- Site, inp <-chan Site) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		select { // send first to whomever is ready to receive
		case out1 <- i:
			out2 <- i
		case out2 <- i:
			out1 <- i
		}
	}
}

// End of SitePair functions
// ===========================================================================

// ===========================================================================
// Beg of SiteFork functions

// SiteFork returns two channels
// either of which is to receive
// every result of inp
// before close.
func SiteFork(inp <-chan Site) (out1, out2 <-chan Site) {
	cha1 := make(chan Site)
	cha2 := make(chan Site)
	go forkSite(cha1, cha2, inp)
	return cha1, cha2
}

/* not used - kept for reference only.
func forkSite ( out1 , out2 chan <- Site , inp <- chan Site ) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
} */

func forkSite(out1, out2 chan<- Site, inp <-chan Site) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		select { // send first to whomever is ready to receive
		case out1 <- i:
			out2 <- i
		case out2 <- i:
			out1 <- i
		}
	}
}

// End of SiteFork functions
// ===========================================================================

// ===========================================================================
// Beg of SiteFanIn2 simple binary Fan-In

// SiteFanIn2 returns a channel to receive
// all from both `inp` and `inp2`
// before close.
func SiteFanIn2(inp, inp2 <-chan Site) (out <-chan Site) {
	cha := make(chan Site)
	go fanIn2Site(cha, inp, inp2)
	return cha
}

/* not used - kept for reference only.
// fanin2Site as seen in Go Concurrency Patterns
func fanin2Site ( out chan <- Site , inp , inp2 <- chan Site ) {
	for {
		select {
		case e := <-inp:
			out <- e
		case e := <-inp2:
			out <- e
		}
	}
} */

func fanIn2Site(out chan<- Site, inp, inp2 <-chan Site) {
	defer close(out)

	var (
		closed bool // we found a chan closed
		ok     bool // did we read successfully?
		e      Site // what we've read
	)

	for !closed {
		select {
		case e, ok = <-inp:
			if ok {
				out <- e
			} else {
				inp = inp2    // swap inp2 into inp
				closed = true // break out of the loop
			}
		case e, ok = <-inp2:
			if ok {
				out <- e
			} else {
				closed = true // break out of the loop				}
			}
		}
	}

	// inp might not be closed yet. Drain it.
	for e = range inp {
		out <- e
	}
}

// End of SiteFanIn2 simple binary Fan-In
// ===========================================================================
