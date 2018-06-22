// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

// ===========================================================================
// Beg of AnyDemand channel object

// AnyDemand is a
// demand channel
type AnyDemand struct {
	dat chan anyThing
	req chan struct{}
}

// MakeAnyDemandChan returns
// a (pointer to a) fresh
// unbuffered
// demand channel
func MakeAnyDemandChan() *AnyDemand {
	d := AnyDemand{
		dat: make(chan anyThing),
		req: make(chan struct{}),
	}
	return &d
}

// MakeAnyDemandBuff returns
// a (pointer to a) fresh
// buffered (with capacity=`cap`)
// demand channel
func MakeAnyDemandBuff(cap int) *AnyDemand {
	d := AnyDemand{
		dat: make(chan anyThing, cap),
		req: make(chan struct{}),
	}
	return &d
}

// Provide is the send method
// - aka "myAnyChan <- myAny"
func (c *AnyDemand) Provide(dat anyThing) {
	<-c.req
	c.dat <- dat
}

// Receive is the receive operator as method
// - aka "myAny := <-myAnyChan"
func (c *AnyDemand) Receive() (dat anyThing) {
	c.req <- struct{}{}
	return <-c.dat
}

// Request is the comma-ok multi-valued form of Receive and
// reports whether a received value was sent before the anyThing channel was closed
func (c *AnyDemand) Request() (dat anyThing, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// Close closes the underlying anyThing channel
func (c *AnyDemand) Close() {
	close(c.dat)
}

// Cap reports the capacity of the underlying anyThing channel
func (c *AnyDemand) Cap() int {
	return cap(c.dat)
}

// Len reports the length of the underlying anyThing channel
func (c *AnyDemand) Len() int {
	return len(c.dat)
}

// End of AnyDemand channel object
// ===========================================================================
