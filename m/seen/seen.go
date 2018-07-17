// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

import (
	"sync"

	"github.com/cheekybits/genny/generic"
)

// anyThing is the generic type flowing thru the pipe network.
type anyThing generic.Type

// ===========================================================================
// Beg of PipeSeen/ForkSeen - an "I've seen this anyThing before" filter / forker

// PipeSeen returns a channel to receive
// all `inp`
// not been seen before
// while silently dropping everything seen before
// (internally growing a `sync.Map` to discriminate)
// until close.
//  Note: PipeFilterNotSeenYet might be a better name, but is fairly long.
func (inp anyThingFrom) PipeSeen() (out anyThingFrom) {
	cha := make(chan anyThing)
	go inp.pipeSeenAttr(cha, nil)
	return cha
}

// PipeSeenAttr returns a channel to receive
// all `inp`
// whose attribute `attr` has
// not been seen before
// while silently dropping everything seen before
// (internally growing a `sync.Map` to discriminate)
// until close.
//  Note: PipeFilterAttrNotSeenYet might be a better name, but is fairly long.
func (inp anyThingFrom) PipeSeenAttr(attr func(a anyThing) interface{}) (out anyThingFrom) {
	cha := make(chan anyThing)
	go inp.pipeSeenAttr(cha, attr)
	return cha
}

// ForkSeen returns two channels, `new` and `old`,
// where `new` is to receive
// all `inp`
// not been seen before
// and `old`
// all `inp`
// seen before
// (internally growing a `sync.Map` to discriminate)
// until close.
func (inp anyThingFrom) ForkSeen() (new, old anyThingFrom) {
	cha1 := make(chan anyThing)
	cha2 := make(chan anyThing)
	go inp.forkSeenAttr(cha1, cha2, nil)
	return cha1, cha2
}

// ForkSeenAttr returns two channels, `new` and `old`,
// where `new` is to receive
// all `inp`
// whose attribute `attr` has
// not been seen before
// and `old`
// all `inp`
// seen before
// (internally growing a `sync.Map` to discriminate)
// until close.
func (inp anyThingFrom) ForkSeenAttr(attr func(a anyThing) interface{}) (new, old anyThingFrom) {
	cha1 := make(chan anyThing)
	cha2 := make(chan anyThing)
	go inp.forkSeenAttr(cha1, cha2, attr)
	return cha1, cha2
}

func (inp anyThingFrom) pipeSeenAttr(out anyThingInto, attr func(a anyThing) interface{}) {
	defer close(out)

	if attr == nil { // Make `nil` value useful
		attr = func(a anyThing) interface{} { return a }
	}

	seen := sync.Map{}
	for i := range inp {
		if _, visited := seen.LoadOrStore(attr(i), struct{}{}); visited {
			// drop i silently
		} else {
			out <- i
		}
	}
}

func (inp anyThingFrom) forkSeenAttr(new, old anyThingInto, attr func(a anyThing) interface{}) {
	defer close(new)
	defer close(old)

	if attr == nil { // Make `nil` value useful
		attr = func(a anyThing) interface{} { return a }
	}

	seen := sync.Map{}
	for i := range inp {
		if _, visited := seen.LoadOrStore(attr(i), struct{}{}); visited {
			old <- i
		} else {
			new <- i
		}
	}
}

// TubeSeen returns a closure around PipeSeen()
// (silently dropping every anyThing seen before).
func (inp anyThingFrom) TubeSeen() (tube func(inp anyThingFrom) (out anyThingFrom)) {

	return func(inp anyThingFrom) (out anyThingFrom) {
		return inp.PipeSeen()
	}
}

// TubeSeenAttr returns a closure around PipeSeenAttr(attr)
// (silently dropping every anyThing
// whose attribute `attr` was
// seen before).
func (inp anyThingFrom) TubeSeenAttr(attr func(a anyThing) interface{}) (tube func(inp anyThingFrom) (out anyThingFrom)) {

	return func(inp anyThingFrom) (out anyThingFrom) {
		return inp.PipeSeenAttr(attr)
	}
}

// End of PipeSeen/ForkSeen - an "I've seen this anyThing before" filter / forker
// ===========================================================================
