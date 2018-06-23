// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

// ===========================================================================
// Beg of anySupply channel object

// anySupply is a
// supply channel
type anySupply struct {
	dat chan anyThing
	//  chan struct{}
}

// anySupplyMakeChan returns
// a (pointer to a) fresh
// unbuffered
// supply channel
func anySupplyMakeChan() *anySupply {
	d := anySupply{
		dat: make(chan anyThing),
		// : make(chan struct{}),
	}
	return &d
}

// anySupplyMakeBuff returns
// a (pointer to a) fresh
// buffered (with capacity=`cap`)
// supply channel
func anySupplyMakeBuff(cap int) *anySupply {
	d := anySupply{
		dat: make(chan anyThing, cap),
		// : make(chan struct{}),
	}
	return &d
}

// Provide is the send method
// - aka "myAnyChan <- myAny"
func (c *anySupply) Provide(dat anyThing) {
	// .req
	c.dat <- dat
}

// Receive is the receive operator as method
// - aka "myAny := <-myAnyChan"
func (c *anySupply) Receive() (dat anyThing) {
	// eq <- struct{}{}
	return <-c.dat
}

// Request is the comma-ok multi-valued form of Receive and
// reports whether a received value was sent before the anyThing channel was closed
func (c *anySupply) Request() (dat anyThing, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// Close closes the underlying anyThing channel
func (c *anySupply) Close() {
	close(c.dat)
}

// Cap reports the capacity of the underlying anyThing channel
func (c *anySupply) Cap() int {
	return cap(c.dat)
}

// Len reports the length of the underlying anyThing channel
func (c *anySupply) Len() int {
	return len(c.dat)
}

// End of anySupply channel object
// ===========================================================================
