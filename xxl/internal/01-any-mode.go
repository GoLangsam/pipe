// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

// ===========================================================================
// Beg of anyDemand channel object

// anyDemand is a
// demand channel
type anyDemand struct {
	ch chan anyThing
	req chan struct{}
}

// anyDemandMakeChan returns
// a (pointer to a) fresh
// unbuffered
// demand channel.
func anyDemandMakeChan() *anyDemand {
	d := anyDemand{
		ch:  make(chan anyThing),
		req: make(chan struct{}),
	}
	return &d
}

// anyDemandMakeBuff returns
// a (pointer to a) fresh
// buffered (with capacity=`cap`)
// demand channel.
func anyDemandMakeBuff(cap int) *anyDemand {
	d := anyDemand{
		ch:  make(chan anyThing, cap),
		req: make(chan struct{}),
	}
	return &d
}

// ---------------------------------------------------------------------------

// Get is the comma-ok multi-valued form to receive from the channel and
// reports whether a received value was sent before the channel was closed.
//
// Get blocks until the request is accepted and value `val` has been received from `from`.
func (from *anyDemand) Get() (val anyThing, open bool) {
	from.req <- struct{}{}
	val, open = <-from.ch
	return
}

// From returns the handshaking channels
// (for use in `select` statements)
// to receive values:
//  `req` to send a request `req <- struct{}{}` and
//  `rcv` to reveive such requested value from.
func (from *anyDemand) From() (req chan<- struct{}, rcv <-chan anyThing) {
	return from.req, from.ch
}

// ---------------------------------------------------------------------------

// Put is the send-upon-request method
// - aka "myAnyChan <- myAny".
//
// Put blocks until requsted to send value `val` into `into`.
func (into *anyDemand) Put(val anyThing) {
	<-into.req
	into.ch <- val
}

// Into returns the handshaking channels
// (for use in `select` statements)
// to send values:
//  `req` to receive a request `<-req` and
//  `snd` to send such requested value into.
func (into *anyDemand) Into() (req <-chan struct{}, snd chan<- anyThing) {
	return into.req, into.ch
}

// Close is to be called by a producer when finished sending.
// The value channel is closed in order to broadcast this.
func (into *anyDemand) Close() {
	close(into.ch)
}

// ---------------------------------------------------------------------------

// Cap reports the capacity of the underlying value channel.
func (c *anyDemand) Cap() int {
	return cap(c.ch)
}

// Len reports the length of the underlying value channel.
func (c *anyDemand) Len() int {
	return len(c.ch)
}

// End of anyDemand channel object
// ===========================================================================
