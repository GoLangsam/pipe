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
		ch: make(chan anyThing),
		// : make(chan struct{}),
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
		ch: make(chan anyThing, cap),
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

// From returns the handshaking channels
// (for use in `select` statements)
// to receive values:
//  `req` to send a request `req <- struct{}{}` and
//  `rcv` to reveive such requested value from.
func (from *anySupply) From() (req chan<- struct{}, rcv <-chan anyThing) {
	cha := make(chan struct{})
	close(cha)
	return cha, from.ch
}

// ---------------------------------------------------------------------------

// NextGetFrom `from` for `into` and report success.
// Follow it with `into.Send( f(val) )`, if ok.
func (into *anySupply) NextGetFrom(from *anySupply) (val anyThing, ok bool) {
	if ok = into.Next(); ok {
		val, ok = from.Get()
	}
	if !ok {
		from.Drop()
		into.Close()
	}
	return
}

// Put is the send-upon-request method
// - aka "myAnyChan <- myAny".
//
// Put blocks until requested to send value `val` into `into` and
// reports whether the request channel was open.
//
// Put is a convenience for
//  if Next() { Send(v) } else { Close() }
//
func (into *anySupply) Put(val anyThing) (ok bool) {
	ok = true // <-into.req
	if ok {
		into.ch <- val
	} else {
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

// Provide is the low-level send-upon-request method
// - aka "myAnyChan <- myAny".
//
// Note: Provide is low-level and differs from Put
// as the latter closes the channel upon nok.
// Use with care.
func (into *anySupply) Provide(val anyThing) (ok bool) {
	ok = true // <-into.req
	if ok {
		into.ch <- val
	}
	return ok
}

// Into returns the handshaking channels
// (for use in `select` statements)
// to send values:
//  `req` to receive a request `<-req` and
//  `snd` to send such requested value into.
func (into *anySupply) Into() (req <-chan struct{}, snd chan<- anyThing) {
	cha := make(chan struct{})
	close(cha)
	return cha, into.ch
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

// MyAnySupply returns itself.
func (c *anySupply) MyAnySupply() *anySupply {
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
