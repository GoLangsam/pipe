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
	req chan struct{}
	ch  chan anyThing
}
*/

// anyModeMakeChan returns
// a (pointer to a) fresh
// unbuffered
// mode channel.
func anyModeMakeChan() *anyMode {
	d := anyMode{
		req: make(chan struct{}),
		ch:  make(chan anyThing),
	}
	return &d
}

// anyModeMakeBuff returns
// a (pointer to a) fresh
// buffered
// mode channel
// (with capacity=`cap`).
func anyModeMakeBuff(cap int) *anyMode {
	d := anyMode{
		req: make(chan struct{}),
		ch:  make(chan anyThing, cap),
	}
	return &d
}

// ---------------------------------------------------------------------------

// Get is the comma-ok multi-valued form to receive from the channel and
// reports whether a value was received from an open channel
// or not (as it has been closed).
//
// Get blocks until the request is accepted and value `val` has been received from `from`.
func (from *anyMode) Get() (val anyThing, open bool) {
	from.req <- struct{}{}
	val, open = <-from.ch
	return
}

// ---------------------------------------------------------------------------
// Put or {`defer into.Close()` and Provide }

// Put is the send-upon-request method
// - aka "myAnyChan <- myAny".
//
// Put blocks until requested to send value `val` into `into` and
// reports whether the request channel was open.
//
// Put is a convenience for
//  if Next() { Send(v) } else { Close() }
//
// Put includes housekeeping:
// If `into` has been dropped, `into` is closed. 
func (into *anyMode) Put(val anyThing) (ok bool) {
	ok = into.Next()
	if ok {
		into.ch <- val
	} else {
		into.Close()
	}
	return
}

// Provide is the low-level send-upon-request method
// - aka "myAnyChan <- myAny".
//
// Note: Provide is low-level - its cousin `Put`
// includes housekeeping: `Put`
// closes the channel upon nok.
//
// Hint: Provide is useful in constructors
// together with `defer into.Close()`.
func (into *anyMode) Provide(val anyThing) (ok bool) {
	ok = into.Next()
	if ok {
		into.ch <- val
	}
	return ok
}

// ---------------------------------------------------------------------------
// Next... => Send

// NextGetFrom `from` for `into` and report success.
//
// Follow it with `into.Send( f(val) )`, if ok.
//
// NextGetFrom includes housekeeping:
// If `into` has been dropped or `from` has been closed,
// `from` is dropped and `into` is closed.
func (into *anyMode) NextGetFrom(from *anyMode) (val anyThing, ok bool) {
	ok = into.Next()
	if ok {
		val, ok = from.Get()
	}
	if !ok {
		from.Drop()
		into.Close()
	}
	return
}

// Next is the request method.
// It blocks until a request is received and
// reports whether the request channel was open.
//
// A successful Next is to be followed by one Send(v).
func (into *anyMode) Next() (ok bool) {
	_, ok = <-into.req
	return
}

// Send is to be used after a successful Next()
func (into *anyMode) Send(val anyThing) {
	into.ch <- val
}

// ---------------------------------------------------------------------------
// Drop & Close signal requesting / sending as being finished

// Drop is to be called by a consumer when finished requesting.
// The request channel is closed in order to broadcast this.
//
// In order to avoid deadlock, pending sends are drained.
func (from *anyMode) Drop() {
	close(from.req)
	go func(from *anyMode) {
		for range from.ch {
		} // drain values - there could be some
	}(from)
}

// Close is to be called by a producer when finished sending.
// The value channel is closed in order to broadcast this.
//
// In order to avoid deadlock, pending requests are drained.
func (into *anyMode) Close() {
	close(into.ch)
	go func(into *anyMode) {
		for range into.req {
		} // drain requests - there could be some
	}(into)
}

// ---------------------------------------------------------------------------
// obtain directional handshaking channels
// (for use e.g. in `select` statements)

// From returns the handshaking channels
// (for use e.g. in `select` statements)
// to receive values:
//  `req` to send a request `req <- struct{}{}` and
//  `rcv` to reveive such requested value from.
func (from *anyMode) From() (req chan<- struct{}, rcv <-chan anyThing) {
	return from.req, from.ch
}

// Into returns the handshaking channels
// (for use e.g. in `select` statements)
// to send values:
//  `req` to receive a request `<-req` and
//  `snd` to send such requested value into.
func (into *anyMode) Into() (req <-chan struct{}, snd chan<- anyThing) {
	return into.req, into.ch
}

// ---------------------------------------------------------------------------

// New returns a new similar channel.
//
// Useful e.g. when embedded anonymously.
func (c *anyMode) New() *anyMode {
	return anyModeMakeChan()
}

// Self returns itself.
//
// Useful e.g. when embedded anonymously
// and e.g. wrappers for multi-value methods are required.
func (c *anyMode) Self() *anyMode {
	return c
}

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
