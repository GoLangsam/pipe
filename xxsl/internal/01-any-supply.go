// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

// ===========================================================================
// Beg of AnySupply channel object

// AnySupply is a
// supply channel
type AnySupply struct {
	dat chan Any
	//  chan struct{}
}

// MakeAnySupplyChan returns
// a (pointer to a) fresh
// unbuffered
// supply channel
func MakeAnySupplyChan() *AnySupply {
	d := AnySupply{
		dat: make(chan Any),
	//	req: make(chan struct{}),
	}
	return &d
}

// MakeAnySupplyBuff returns
// a (pointer to a) fresh
// buffered (with capacity=`cap`)
// supply channel
func MakeAnySupplyBuff(cap int) *AnySupply {
	d := AnySupply{
		dat: make(chan Any, cap),
	//	req: make(chan struct{}),
	}
	return &d
}

// Provide is the send method
// - aka "myAnyChan <- myAny"
func (c *AnySupply) Provide(dat Any) {
	// .req
	c.dat <- dat
}

// Receive is the receive operator as method
// - aka "myAny := <-myAnyChan"
func (c *AnySupply) Receive() (dat Any) {
	// eq <- struct{}{}
	return <-c.dat
}

// Request is the comma-ok multi-valued form of Receive and
// reports whether a received value was sent before the Any channel was closed
func (c *AnySupply) Request() (dat Any, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// Close closes the underlying Any channel
func (c *AnySupply) Close() {
	close(c.dat)
}

// Cap reports the capacity of the underlying Any channel
func (c *AnySupply) Cap() int {
	return cap(c.dat)
}

// Len reports the length of the underlying Any channel
func (c *AnySupply) Len() int {
	return len(c.dat)
}

// End of AnySupply channel object
// ===========================================================================
