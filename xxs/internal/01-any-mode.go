// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

// ===========================================================================
// Beg of anySupply channel object

// anySupply is a
// supply channel
type anySupply struct {
	//  chan struct{}
	ch chan anyThing
}

// anySupplyMakeChan returns
// a (pointer to a) fresh
// unbuffered
// supply channel.
func anySupplyMakeChan() *anySupply {
	d := anySupply{
		// : make(chan struct{}),
		ch: make(chan anyThing),
	}
	return &d
}

// anySupplyMakeBuff returns
// a (pointer to a) fresh
// buffered
// supply channel
// (with capacity=`cap`).
func anySupplyMakeBuff(cap int) *anySupply {
	d := anySupply{
		// : make(chan struct{}),
		ch: make(chan anyThing, cap),
	}
	return &d
}

// ---------------------------------------------------------------------------

// Get is the comma-ok multi-valued form to receive from the channel and
// reports whether a value was received from an open channel
// or not (as it has been closed).
//
// Get blocks until the request is accepted and value `val` has been received from `from`.
func (from *anySupply) Get() (val anyThing, open bool) {
	// m.req <- struct{}{}
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
func (into *anySupply) Put(val anyThing) (ok bool) {
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
func (into *anySupply) Provide(val anyThing) (ok bool) {
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
func (into *anySupply) NextGetFrom(from *anySupply) (val anyThing, ok bool) {
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
func (into *anySupply) Next() (ok bool) {
	ok = true // <-into.req
	return
}

// Send is to be used after a successful Next()
func (into *anySupply) Send(val anyThing) {
	into.ch <- val
}

// ---------------------------------------------------------------------------
// Drop & Close signal requesting / sending as being finished

// Drop is to be called by a consumer when finished requesting.
// The request channel is closed in order to broadcast this.
//
// In order to avoid deadlock, pending sends are drained.
func (from *anySupply) Drop() {
	// se(from.req)
	go func(from *anySupply) {
		for range from.ch {
		} // drain values - there could be some
	}(from)
}

// Close is to be called by a producer when finished sending.
// The value channel is closed in order to broadcast this.
//
// In order to avoid deadlock, pending requests are drained.
func (into *anySupply) Close() {
	close(into.ch)
	/*
		go func(into *anySupply) {
			for range into.req {
			} // drain requests - there could be some
		}(into)
	*/
}

// ---------------------------------------------------------------------------
// obtain directional handshaking channels
// (for use e.g. in `select` statements)

// From returns the handshaking channels
// (for use e.g. in `select` statements)
// to receive values:
//  `req` to send a request `req <- struct{}{}` and
//  `rcv` to reveive such requested value from.
func (from *anySupply) From() (req chan<- struct{}, rcv <-chan anyThing) {
	cha := make(chan struct{})
	close(cha)
	return cha, from.ch
}

// Into returns the handshaking channels
// (for use e.g. in `select` statements)
// to send values:
//  `req` to receive a request `<-req` and
//  `snd` to send such requested value into.
func (into *anySupply) Into() (req <-chan struct{}, snd chan<- anyThing) {
	cha := make(chan struct{})
	close(cha)
	return cha, into.ch
}

// ---------------------------------------------------------------------------

// New returns a new similar channel.
//
// Useful e.g. when embedded anonymously.
func (c *anySupply) New() *anySupply {
	return anySupplyMakeChan()
}

// Self returns itself.
//
// Useful e.g. when embedded anonymously
// and e.g. wrappers for multi-value methods are required.
func (c *anySupply) Self() *anySupply {
	return c
}

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
