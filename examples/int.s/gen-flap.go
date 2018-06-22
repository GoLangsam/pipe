// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

// ===========================================================================
// Beg of intPipeEnter/Leave - Flapdoors observed by a Waiter

// intWaiter - as implemented by `*sync.WaitGroup` -
// attends Flapdoors and keeps counting
// who enters and who leaves.
//
// Use intDoneWait to learn about
// when the facilities are closed.
//
// Note: You may also use Your provided `*sync.WaitGroup.Wait()`
// to know when to close the facilities.
// Just: intDoneWait is more convenient
// as it also closes the primary channel for You.
//
// Just make sure to have _all_ entrances and exits attended,
// and `Wait()` only *after* You've started flooding the facilities.
type intWaiter interface {
	Add(delta int)
	Done()
	Wait()
}

// Note: The name is intentionally generic in order to avoid eventual multiple-declaration clashes.

// intPipeEnter returns a channel to receive
// all `inp`
// and registers throughput
// as arrival
// on the given `sync.WaitGroup`
// until close.
func intPipeEnter(inp <-chan int, wg intWaiter) (out <-chan int) {
	cha := make(chan int)
	go pipeintEnter(cha, wg, inp)
	return cha
}

// intPipeLeave returns a channel to receive
// all `inp`
// and registers throughput
// as departure
// on the given `sync.WaitGroup`
// until close.
func intPipeLeave(inp <-chan int, wg intWaiter) (out <-chan int) {
	cha := make(chan int)
	go pipeintLeave(cha, wg, inp)
	return cha
}

// intDoneLeave returns a channel to receive
// one signal after
// all throughput on `inp`
// has been registered
// as departure
// on the given `sync.WaitGroup`
// before close.
func intDoneLeave(inp <-chan int, wg intWaiter) (done <-chan struct{}) {
	sig := make(chan struct{})
	go doneintLeave(sig, wg, inp)
	return sig
}

func pipeintEnter(out chan<- int, wg intWaiter, inp <-chan int) {
	defer close(out)
	for i := range inp {
		wg.Add(1)
		out <- i
	}
}

func pipeintLeave(out chan<- int, wg intWaiter, inp <-chan int) {
	defer close(out)
	for i := range inp {
		out <- i
		wg.Done()
	}
}

func doneintLeave(done chan<- struct{}, wg intWaiter, inp <-chan int) {
	defer close(done)
	for i := range inp {
		_ = i // discard
		wg.Done()
	}
	done <- struct{}{}
}

// intTubeEnter returns a closure around intPipeEnter (_, wg)
// registering throughput
// as arrival
// on the given `sync.WaitGroup`.
func intTubeEnter(wg intWaiter) (tube func(inp <-chan int) (out <-chan int)) {

	return func(inp <-chan int) (out <-chan int) {
		return intPipeEnter(inp, wg)
	}
}

// intTubeLeave returns a closure around intPipeLeave (_, wg)
// registering throughput
// as departure
// on the given `sync.WaitGroup`.
func intTubeLeave(wg intWaiter) (tube func(inp <-chan int) (out <-chan int)) {

	return func(inp <-chan int) (out <-chan int) {
		return intPipeLeave(inp, wg)
	}
}

// intFiniLeave returns a closure around `intDoneLeave(_, wg)`
// registering throughput
// as departure
// on the given `sync.WaitGroup`.
func intFiniLeave(wg intWaiter) func(inp <-chan int) (done <-chan struct{}) {

	return func(inp <-chan int) (done <-chan struct{}) {
		return intDoneLeave(inp, wg)
	}
}

// intDoneWait returns a channel to receive
// one signal
// after wg.Wait() has returned and inp has been closed
// before close.
//
// Note: Use only *after* You've started flooding the facilities.
func intDoneWait(inp chan<- int, wg intWaiter) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doneintWait(cha, inp, wg)
	return cha
}

func doneintWait(done chan<- struct{}, inp chan<- int, wg intWaiter) {
	defer close(done)
	wg.Wait()
	close(inp)
	done <- struct{}{} // not really needed - but looks better
}

// intFiniWait returns a closure around `DoneintWait(_, wg)`.
func intFiniWait(wg intWaiter) func(inp chan<- int) (done <-chan struct{}) {

	return func(inp chan<- int) (done <-chan struct{}) {
		return intDoneWait(inp, wg)
	}
}

// End of intPipeEnter/Leave - Flapdoors observed by a Waiter
