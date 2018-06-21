// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sites

// ===========================================================================
// Beg of sitePipeEnter/Leave - Flapdoors observed by a Waiter

// siteWaiter - as implemented by `*sync.WaitGroup` -
// attends Flapdoors and keeps counting
// who enters and who leaves.
//
// Use DoneSiteWait to learn about
// when the facilities are closed.
//
// Note: You may also use Your provided `*sync.WaitGroup.Wait()`
// to know when to close the facilities.
// Just: DoneSiteWait is more convenient
// as it also closes the primary channel.
//
// Just make sure to have _all_ entrances and exits attended,
// and `Wait()` only *after* You've started flooding the facilities.
type siteWaiter interface {
	Add(delta int)
	Done()
	Wait()
}

// Note: Name is generic in order to avoid multiple-declaration clashes.

// sitePipeEnter returns a channel to receive
// all `inp`
// and registers throughput
// as arrival
// on the given `sync.WaitGroup`
// until close.
func sitePipeEnter(inp <-chan site, wg siteWaiter) (out <-chan site) {
	cha := make(chan site)
	go pipesiteEnter(cha, wg, inp)
	return cha
}

// sitePipeLeave returns a channel to receive
// all `inp`
// and registers throughput
// as departure
// on the given `sync.WaitGroup`
// until close.
func sitePipeLeave(inp <-chan site, wg siteWaiter) (out <-chan site) {
	cha := make(chan site)
	go pipesiteLeave(cha, wg, inp)
	return cha
}

func pipesiteEnter(out chan<- site, wg siteWaiter, inp <-chan site) {
	defer close(out)
	for i := range inp {
		wg.Add(1)
		out <- i
	}
}

func pipesiteLeave(out chan<- site, wg siteWaiter, inp <-chan site) {
	defer close(out)
	for i := range inp {
		out <- i
		wg.Done()
	}
}

// siteTubeEnter returns a closure around sitePipeEnter (_, wg)
// registering throughput
// on the given `sync.WaitGroup`
// as arrival.
func siteTubeEnter(wg siteWaiter) (tube func(inp <-chan site) (out <-chan site)) {

	return func(inp <-chan site) (out <-chan site) {
		return sitePipeEnter(inp, wg)
	}
}

// siteTubeLeave returns a closure around sitePipeLeave (_, wg)
// registering throughput
// on the given `sync.WaitGroup`
// as departure.
func siteTubeLeave(wg siteWaiter) (tube func(inp <-chan site) (out <-chan site)) {

	return func(inp <-chan site) (out <-chan site) {
		return sitePipeLeave(inp, wg)
	}
}

// siteDoneWait returns a channel to receive
// one signal
// after wg.Wait() has returned and inp has been closed
// before close.
//
// Note: Use only *after* You've started flooding the facilities.
func siteDoneWait(inp chan<- site, wg siteWaiter) (done <-chan struct{}) {
	cha := make(chan struct{})
	go donesiteWait(cha, inp, wg)
	return cha
}

func donesiteWait(done chan<- struct{}, inp chan<- site, wg siteWaiter) {
	defer close(done)
	wg.Wait()
	close(inp)
	done <- struct{}{} // not really needed - but looks better
}

// siteFiniWait returns a closure around `DonesiteWait(_, wg)`.
func siteFiniWait(wg siteWaiter) func(inp chan<- site) (done <-chan struct{}) {

	return func(inp chan<- site) (done <-chan struct{}) {
		return siteDoneWait(inp, wg)
	}
}

// End of sitePipeEnter/Leave - Flapdoors observed by a Waiter
