// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by golang.org/x/tools/cmd/bundle. DO NOT EDIT.

package httpsyet

// ===========================================================================
// Beg of siteMake creators

// siteMakeChan returns a new open channel
// (simply a 'chan site' that is).
// Note: No 'site-producer' is launched here yet! (as is in all the other functions).
//  This is useful to easily create corresponding variables such as:
/*
var mysitePipelineStartsHere := siteMakeChan()
// ... lot's of code to design and build Your favourite "mysiteWorkflowPipeline"
   // ...
   // ... *before* You start pouring data into it, e.g. simply via:
   for drop := range water {
mysitePipelineStartsHere <- drop
   }
close(mysitePipelineStartsHere)
*/
//  Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
//  (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for sitePipeBuffer) the channel is unbuffered.
//
func siteMakeChan() (out chan site) {
	return make(chan site)
}

// End of siteMake creators
// ===========================================================================

// ===========================================================================
// Beg of siteChan producers

// siteChan returns a channel to receive
// all inputs
// before close.
func siteChan(inp ...site) (out <-chan site) {
	cha := make(chan site)
	go chansite(cha, inp...)
	return cha
}

func chansite(out chan<- site, inp ...site) {
	defer close(out)
	for i := range inp {
		out <- inp[i]
	}
}

// siteChanSlice returns a channel to receive
// all inputs
// before close.
func siteChanSlice(inp ...[]site) (out <-chan site) {
	cha := make(chan site)
	go chansiteSlice(cha, inp...)
	return cha
}

func chansiteSlice(out chan<- site, inp ...[]site) {
	defer close(out)
	for i := range inp {
		for j := range inp[i] {
			out <- inp[i][j]
		}
	}
}

// siteChanFuncNok returns a channel to receive
// all results of generator `gen`
// until `!ok`
// before close.
func siteChanFuncNok(gen func() (site, bool)) (out <-chan site) {
	cha := make(chan site)
	go chansiteFuncNok(cha, gen)
	return cha
}

func chansiteFuncNok(out chan<- site, gen func() (site, bool)) {
	defer close(out)
	for {
		res, ok := gen() // generate
		if !ok {
			return
		}
		out <- res
	}
}

// siteChanFuncErr returns a channel to receive
// all results of generator `gen`
// until `err != nil`
// before close.
func siteChanFuncErr(gen func() (site, error)) (out <-chan site) {
	cha := make(chan site)
	go chansiteFuncErr(cha, gen)
	return cha
}

func chansiteFuncErr(out chan<- site, gen func() (site, error)) {
	defer close(out)
	for {
		res, err := gen() // generate
		if err != nil {
			return
		}
		out <- res
	}
}

// End of siteChan producers
// ===========================================================================

// ===========================================================================
// Beg of sitePipe functions

// sitePipeFunc returns a channel to receive
// every result of action `act` applied to `inp`
// before close.
// Note: it 'could' be PipeSiteMap for functional people,
// but 'map' has a very different meaning in go lang.
func sitePipeFunc(inp <-chan site, act func(a site) site) (out <-chan site) {
	cha := make(chan site)
	if act == nil { // Make `nil` value useful
		act = func(a site) site { return a }
	}
	go pipesiteFunc(cha, inp, act)
	return cha
}

func pipesiteFunc(out chan<- site, inp <-chan site, act func(a site) site) {
	defer close(out)
	for i := range inp {
		out <- act(i) // apply action
	}
}

// sitePipeBuffer returns a buffered channel with capacity `cap` to receive
// all `inp`
// before close.
func sitePipeBuffer(inp <-chan site, cap int) (out <-chan site) {
	cha := make(chan site, cap)
	go pipesiteBuffer(cha, inp)
	return cha
}

func pipesiteBuffer(out chan<- site, inp <-chan site) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// End of PipeSite functions
// ===========================================================================

// ===========================================================================
// Beg of siteTube closures around sitePipe

// siteTubeFunc returns a closure around PipeSiteFunc (_, act).
func siteTubeFunc(act func(a site) site) (tube func(inp <-chan site) (out <-chan site)) {

	return func(inp <-chan site) (out <-chan site) {
		return sitePipeFunc(inp, act)
	}
}

// siteTubeBuffer returns a closure around PipeSiteBuffer (_, cap).
func siteTubeBuffer(cap int) (tube func(inp <-chan site) (out <-chan site)) {

	return func(inp <-chan site) (out <-chan site) {
		return sitePipeBuffer(inp, cap)
	}
}

// End of siteTube closures around sitePipe
// ===========================================================================

// ===========================================================================
// Beg of siteDone terminators

// siteDone returns a channel to receive
// one signal before close after `inp` has been drained.
func siteDone(inp <-chan site) (done <-chan struct{}) {
	sig := make(chan struct{})
	go donesite(sig, inp)
	return sig
}

func donesite(done chan<- struct{}, inp <-chan site) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// siteDoneSlice returns a channel to receive
// a slice with every site received on `inp`
// before close.
//
// Note: Unlike siteDone, DoneSiteSlice sends the fully accumulated slice, not just an event, once upon close of inp.
func siteDoneSlice(inp <-chan site) (done <-chan []site) {
	sig := make(chan []site)
	go donesiteSlice(sig, inp)
	return sig
}

func donesiteSlice(done chan<- []site, inp <-chan site) {
	defer close(done)
	slice := []site{}
	for i := range inp {
		slice = append(slice, i)
	}
	done <- slice
}

// siteDoneFunc returns a channel to receive
// one signal after `act` has been applied to every `inp`
// before close.
func siteDoneFunc(inp <-chan site, act func(a site)) (done <-chan struct{}) {
	sig := make(chan struct{})
	if act == nil {
		act = func(a site) { return }
	}
	go donesiteFunc(sig, inp, act)
	return sig
}

func donesiteFunc(done chan<- struct{}, inp <-chan site, act func(a site)) {
	defer close(done)
	for i := range inp {
		act(i) // apply action
	}
	done <- struct{}{}
}

// End of siteDone terminators
// ===========================================================================

// ===========================================================================
// Beg of siteFini closures

// siteFini returns a closure around `siteDone(_)`.
func siteFini() func(inp <-chan site) (done <-chan struct{}) {

	return func(inp <-chan site) (done <-chan struct{}) {
		return siteDone(inp)
	}
}

// siteFiniSlice returns a closure around `siteDoneSlice(_)`.
func siteFiniSlice() func(inp <-chan site) (done <-chan []site) {

	return func(inp <-chan site) (done <-chan []site) {
		return siteDoneSlice(inp)
	}
}

// siteFiniFunc returns a closure around `siteDoneFunc(_, act)`.
func siteFiniFunc(act func(a site)) func(inp <-chan site) (done <-chan struct{}) {

	return func(inp <-chan site) (done <-chan struct{}) {
		return siteDoneFunc(inp, act)
	}
}

// End of siteFini closures
// ===========================================================================

// ===========================================================================
// Beg of sitePair functions

// sitePair returns a pair of channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func sitePair(inp <-chan site) (out1, out2 <-chan site) {
	cha1 := make(chan site)
	cha2 := make(chan site)
	go pairsite(cha1, cha2, inp)
	return cha1, cha2
}

/* not used - kept for reference only.
func pairsite(out1, out2 chan<- site, inp <-chan site) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
} */

func pairsite(out1, out2 chan<- site, inp <-chan site) {
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

// End of sitePair functions
// ===========================================================================

// ===========================================================================
// Beg of siteFork functions

// siteFork returns two channels
// either of which is to receive
// every result of inp
// before close.
func siteFork(inp <-chan site) (out1, out2 <-chan site) {
	cha1 := make(chan site)
	cha2 := make(chan site)
	go forksite(cha1, cha2, inp)
	return cha1, cha2
}

/* not used - kept for reference only.
func forksite(out1, out2 chan<- site, inp <-chan site) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
} */

func forksite(out1, out2 chan<- site, inp <-chan site) {
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

// End of siteFork functions
// ===========================================================================

// ===========================================================================
// Beg of siteFanIn2 simple binary Fan-In

// siteFanIn2 returns a channel to receive all to receive all from both `inp1` and `inp2` before close.
func siteFanIn2(inp1, inp2 <-chan site) (out <-chan site) {
	cha := make(chan site)
	go fanIn2site(cha, inp1, inp2)
	return cha
}

/* not used - kept for reference only.
// fanin2site as seen in Go Concurrency Patterns
func fanin2site(out chan<- site, inp1, inp2 <-chan site) {
	for {
		select {
		case e := <-inp1:
			out <- e
		case e := <-inp2:
			out <- e
		}
	}
} */

func fanIn2site(out chan<- site, inp1, inp2 <-chan site) {
	defer close(out)

	var (
		closed bool // we found a chan closed
		ok     bool // did we read successfully?
		e      site // what we've read
	)

	for !closed {
		select {
		case e, ok = <-inp1:
			if ok {
				out <- e
			} else {
				inp1 = inp2   // swap inp2 into inp1
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

	// inp1 might not be closed yet. Drain it.
	for e = range inp1 {
		out <- e
	}
}

// End of siteFanIn2 simple binary Fan-In
