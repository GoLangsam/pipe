// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

// ===========================================================================
// Beg of ThingPipeEnter/Leave - Flapdoors observed by a Waiter

// ThingWaiter - as implemented by `*sync.WaitGroup` -
// attends Flapdoors and keeps counting
// who enters and who leaves.
//
// Use ThingDoneWait to learn about
// when the facilities are closed.
//
// Note: You may also use Your provided `*sync.WaitGroup.Wait()`
// to know when to close the facilities.
// Just: ThingDoneWait is more convenient
// as it also closes the primary channel for You.
//
// Just make sure to have _all_ entrances and exits attended,
// and `Wait()` only *after* You've started flooding the facilities.
type ThingWaiter interface {
	Add(delta int)
	Done()
	Wait()
}

// Note: The name is intentionally generic in order to avoid eventual multiple-declaration clashes.

// ThingPipeEnter returns a channel to receive
// all `inp`
// and registers throughput
// as arrival
// on the given `sync.WaitGroup`
// until close.
func ThingPipeEnter(inp <-chan Thing, wg ThingWaiter) (out <-chan Thing) {
	cha := make(chan Thing)
	go pipeThingEnter(cha, wg, inp)
	return cha
}

// ThingPipeLeave returns a channel to receive
// all `inp`
// and registers throughput
// as departure
// on the given `sync.WaitGroup`
// until close.
func ThingPipeLeave(inp <-chan Thing, wg ThingWaiter) (out <-chan Thing) {
	cha := make(chan Thing)
	go pipeThingLeave(cha, wg, inp)
	return cha
}

// ThingDoneLeave returns a channel to receive
// one signal after
// all throughput on `inp`
// has been registered
// as departure
// on the given `sync.WaitGroup`
// before close.
func ThingDoneLeave(inp <-chan Thing, wg ThingWaiter) (done <-chan struct{}) {
	sig := make(chan struct{})
	go doneThingLeave(sig, wg, inp)
	return sig
}

func pipeThingEnter(out chan<- Thing, wg ThingWaiter, inp <-chan Thing) {
	defer close(out)
	for i := range inp {
		wg.Add(1)
		out <- i
	}
}

func pipeThingLeave(out chan<- Thing, wg ThingWaiter, inp <-chan Thing) {
	defer close(out)
	for i := range inp {
		out <- i
		wg.Done()
	}
}

func doneThingLeave(done chan<- struct{}, wg ThingWaiter, inp <-chan Thing) {
	defer close(done)
	for i := range inp {
		_ = i // discard
		wg.Done()
	}
	done <- struct{}{}
}

// ThingTubeEnter returns a closure around ThingPipeEnter (_, wg)
// registering throughput
// as arrival
// on the given `sync.WaitGroup`.
func ThingTubeEnter(wg ThingWaiter) (tube func(inp <-chan Thing) (out <-chan Thing)) {

	return func(inp <-chan Thing) (out <-chan Thing) {
		return ThingPipeEnter(inp, wg)
	}
}

// ThingTubeLeave returns a closure around ThingPipeLeave (_, wg)
// registering throughput
// as departure
// on the given `sync.WaitGroup`.
func ThingTubeLeave(wg ThingWaiter) (tube func(inp <-chan Thing) (out <-chan Thing)) {

	return func(inp <-chan Thing) (out <-chan Thing) {
		return ThingPipeLeave(inp, wg)
	}
}

// ThingFiniLeave returns a closure around `ThingDoneLeave(_, wg)`
// registering throughput
// as departure
// on the given `sync.WaitGroup`.
func ThingFiniLeave(wg ThingWaiter) func(inp <-chan Thing) (done <-chan struct{}) {

	return func(inp <-chan Thing) (done <-chan struct{}) {
		return ThingDoneLeave(inp, wg)
	}
}

// ThingDoneWait returns a channel to receive
// one signal
// after wg.Wait() has returned and inp has been closed
// before close.
//
// Note: Use only *after* You've started flooding the facilities.
func ThingDoneWait(inp chan<- Thing, wg ThingWaiter) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doneThingWait(cha, inp, wg)
	return cha
}

func doneThingWait(done chan<- struct{}, inp chan<- Thing, wg ThingWaiter) {
	defer close(done)
	wg.Wait()
	close(inp)
	done <- struct{}{} // not really needed - but looks better
}

// ThingFiniWait returns a closure around `DoneThingWait(_, wg)`.
func ThingFiniWait(wg ThingWaiter) func(inp chan<- Thing) (done <-chan struct{}) {

	return func(inp chan<- Thing) (done <-chan struct{}) {
		return ThingDoneWait(inp, wg)
	}
}

// End of ThingPipeEnter/Leave - Flapdoors observed by a Waiter
// ===========================================================================
