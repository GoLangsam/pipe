// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

// ===========================================================================
// Beg of anyDemand channel object

// anyDemand is a
// demand channel
type anyDemand struct {
	dat chan anyThing
	req chan struct{}
}

// anyDemandMakeChan returns
// a (pointer to a) fresh
// unbuffered
// demand channel
func anyDemandMakeChan() *anyDemand {
	d := anyDemand{
		dat: make(chan anyThing),
		req: make(chan struct{}),
	}
	return &d
}

// anyDemandMakeBuff returns
// a (pointer to a) fresh
// buffered (with capacity=`cap`)
// demand channel
func anyDemandMakeBuff(cap int) *anyDemand {
	d := anyDemand{
		dat: make(chan anyThing, cap),
		req: make(chan struct{}),
	}
	return &d
}

// Provide is the send method
// - aka "myAnyChan <- myAny"
func (c *anyDemand) Provide(dat anyThing) {
	<-c.req
	c.dat <- dat
}

// Receive is the receive operator as method
// - aka "myAny := <-myAnyChan"
func (c *anyDemand) Receive() (dat anyThing) {
	c.req <- struct{}{}
	return <-c.dat
}

// Request is the comma-ok multi-valued form of Receive and
// reports whether a received value was sent before the anyThing channel was closed
func (c *anyDemand) Request() (dat anyThing, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// Close closes the underlying anyThing channel
func (c *anyDemand) Close() {
	close(c.dat)
}

// Cap reports the capacity of the underlying anyThing channel
func (c *anyDemand) Cap() int {
	return cap(c.dat)
}

// Len reports the length of the underlying anyThing channel
func (c *anyDemand) Len() int {
	return len(c.dat)
}

// End of anyDemand channel object
// ===========================================================================
