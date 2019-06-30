// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate genny -in $GOFILE	-out ../xxs/internal/$GOFILE.supply	gen "mode=supply anyMode=anySupply"
//go:generate genny -in $GOFILE	-out ../xxl/internal/$GOFILE.demand	gen "mode=demand anyMode=anyDemand"

package pipe

import "github.com/cheekybits/genny/generic"

type anyThing interface{}

type mode generic.Type
type anyMode generic.Type

// ===========================================================================
// Beg of anyMode channel object

/*
// anyMode is a
// mode channel
type anyMode struct {
	ch  chan anyThing
	req chan struct{}
}
*/

// anyModeMakeChan returns
// a (pointer to a) fresh
// unbuffered
// mode channel.
func anyModeMakeChan() *anyMode {
	d := anyMode{
		ch:  make(chan anyThing),
		req: make(chan struct{}),
	}
	return &d
}

// anyModeMakeBuff returns
// a (pointer to a) fresh
// buffered (with capacity=`cap`)
// mode channel.
func anyModeMakeBuff(cap int) *anyMode {
	d := anyMode{
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
func (from *anyMode) Get() (val anyThing, open bool) {
	from.req <- struct{}{}
	val, open = <-from.ch
	return
}

// From returns the handshaking channels
// (for use in `select` statements)
// to receive values:
//  `req` to send a request `req <- struct{}{}` and
//  `rcv` to reveive such requested value from.
func (from *anyMode) From() (req chan<- struct{}, rcv <-chan anyThing) {
	return from.req, from.ch
}

// ---------------------------------------------------------------------------

// Put is the send-upon-request method
// - aka "myAnyChan <- myAny".
//
// Put blocks until requsted to send value `val` into `into`.
func (into *anyMode) Put(val anyThing) {
	<-into.req
	into.ch <- val
}

// Into returns the handshaking channels
// (for use in `select` statements)
// to send values:
//  `req` to receive a request `<-req` and
//  `snd` to send such requested value into.
func (into *anyMode) Into() (req <-chan struct{}, snd chan<- anyThing) {
	return into.req, into.ch
}

// Close is to be called by a producer when finished sending.
// The value channel is closed in order to broadcast this.
func (into *anyMode) Close() {
	close(into.ch)
}

// ---------------------------------------------------------------------------

// Cap reports the capacity of the underlying value channel.
func (c *anyMode) Cap() int {
	return cap(c.ch)
}

// Len reports the length of the underlying value channel.
func (c *anyMode) Len() int {
	return len(c.ch)
}

// End of anyMode channel object
// ===========================================================================
