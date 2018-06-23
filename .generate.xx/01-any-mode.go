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
	dat chan anyThing
	req chan struct{}
}
*/

// anyModeMakeChan returns
// a (pointer to a) fresh
// unbuffered
// mode channel
func anyModeMakeChan() *anyMode {
	d := anyMode{
		dat: make(chan anyThing),
		req: make(chan struct{}),
	}
	return &d
}

// anyModeMakeBuff returns
// a (pointer to a) fresh
// buffered (with capacity=`cap`)
// mode channel
func anyModeMakeBuff(cap int) *anyMode {
	d := anyMode{
		dat: make(chan anyThing, cap),
		req: make(chan struct{}),
	}
	return &d
}

// Provide is the send method
// - aka "myAnyChan <- myAny"
func (c *anyMode) Provide(dat anyThing) {
	<-c.req
	c.dat <- dat
}

// Receive is the receive operator as method
// - aka "myAny := <-myAnyChan"
func (c *anyMode) Receive() (dat anyThing) {
	c.req <- struct{}{}
	return <-c.dat
}

// Request is the comma-ok multi-valued form of Receive and
// reports whether a received value was sent before the anyThing channel was closed
func (c *anyMode) Request() (dat anyThing, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// Close closes the underlying anyThing channel
func (c *anyMode) Close() {
	close(c.dat)
}

// Cap reports the capacity of the underlying anyThing channel
func (c *anyMode) Cap() int {
	return cap(c.dat)
}

// Len reports the length of the underlying anyThing channel
func (c *anyMode) Len() int {
	return len(c.dat)
}

// End of anyMode channel object
// ===========================================================================
//
