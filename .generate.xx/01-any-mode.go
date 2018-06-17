// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate genny -in $GOFILE	-out ../xxs/internal/$GOFILE.supply	gen "mode=supply"
//go:generate genny -in $GOFILE	-out ../xxl/internal/$GOFILE.demand	gen "mode=demand"

package pipe

import "github.com/cheekybits/genny/generic"

type Any interface{}

type mode generic.Type

// ===========================================================================
// Beg of Anymode channel object

// Anymode is a
// mode channel
type Anymode struct {
	dat chan Any
	req chan struct{}
}

// MakeAnymodeChan returns
// a (pointer to a) fresh
// unbuffered
// mode channel
func MakeAnymodeChan() *Anymode {
	d := Anymode{
		dat: make(chan Any),
		req: make(chan struct{}),
		}
	return &d
}

// MakeAnymodeBuff returns
// a (pointer to a) fresh
// buffered (with capacity=`cap`)
// mode channel
func MakeAnymodeBuff(cap int) *Anymode {
	d := Anymode{
		dat: make(chan Any, cap),
		req: make(chan struct{}),
		}
	return &d
}

// Provide is the send method
// - aka "myAnyChan <- myAny"
func (c *Anymode) Provide(dat Any) {
	<-c.req
	c.dat <- dat
}

// Receive is the receive operator as method
// - aka "myAny := <-myAnyChan"
func (c *Anymode) Receive() (dat Any) {
	c.req <- struct{}{}
	return <-c.dat
}

// Request is the comma-ok multi-valued form of Receive and
// reports whether a received value was sent before the Any channel was closed
func (c *Anymode) Request() (dat Any, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// Close closes the underlying Any channel
func (c *Anymode) Close() {
	close(c.dat)
}

// Cap reports the capacity of the underlying Any channel
func (c *Anymode) Cap() int {
	return cap(c.dat)
}

// Len reports the length of the underlying Any channel
func (c *Anymode) Len() int {
	return len(c.dat)
}

// End of Anymode channel object
// ===========================================================================
//
