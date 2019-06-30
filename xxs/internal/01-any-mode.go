// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

// ===========================================================================
// Beg of anySupply channel object

// anySupply is a
// supply channel
type anySupply struct {
	ch chan anyThing
	//  chan struct{}
}

// anySupplyMakeChan returns
// a (pointer to a) fresh
// unbuffered
// supply channel.
func anySupplyMakeChan() *anySupply {
	d := anySupply{
		ch:  make(chan anyThing),
		// : make(chan struct{}),
	}
	return &d
}

// anySupplyMakeBuff returns
// a (pointer to a) fresh
// buffered (with capacity=`cap`)
// supply channel.
func anySupplyMakeBuff(cap int) *anySupply {
	d := anySupply{
		ch:  make(chan anyThing, cap),
		// : make(chan struct{}),
	}
	return &d
}

// ---------------------------------------------------------------------------

// Get is the comma-ok multi-valued form to receive from the channel and
// reports whether a received value was sent before the channel was closed.
//
// Get blocks until the request is accepted and value `val` has been received from `from`.
func (from *anySupply) Get() (val anyThing, open bool) {
	// m.req <- struct{}{}
	val, open = <-from.ch
	return
}

// ---------------------------------------------------------------------------

// Put is the send-upon-request method
// - aka "myAnyChan <- myAny".
//
// Put blocks until requsted to send value `val` into `into`.
func (into *anySupply) Put(val anyThing) {
	// nto.req
	into.ch <- val
}

// Close is to be called by a producer when finished sending.
// The value channel is closed in order to broadcast this.
func (into *anySupply) Close() {
	close(into.ch)
}

// ---------------------------------------------------------------------------

// Cap reports the capacity of the underlying value channel.
func (c *anySupply) Cap() int {
	return cap(c.ch)
}

// Len reports the length of the underlying value channel.
func (c *anySupply) Len() int {
	return len(c.ch)
}

// End of anySupply channel object
// ===========================================================================
