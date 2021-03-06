// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

import "sync"

// ===========================================================================
// Beg of ThingPipeSeen/ThingForkSeen - an "I've seen this Thing before" filter / forker

// ThingPipeSeen returns a channel to receive
// all `inp`
// not been seen before
// while silently dropping everything seen before
// (internally growing a `sync.Map` to discriminate)
// until close.
// Note: ThingPipeFilterNotSeenYet might be a better name, but is fairly long.
func (inp ThingFrom) ThingPipeSeen() (out ThingFrom) {
	cha := make(chan Thing)
	go inp.pipeThingSeenAttr(cha, nil)
	return cha
}

// ThingPipeSeenAttr returns a channel to receive
// all `inp`
// whose attribute `attr` has
// not been seen before
// while silently dropping everything seen before
// (internally growing a `sync.Map` to discriminate)
// until close.
// Note: ThingPipeFilterAttrNotSeenYet might be a better name, but is fairly long.
func (inp ThingFrom) ThingPipeSeenAttr(attr func(a Thing) interface{}) (out ThingFrom) {
	cha := make(chan Thing)
	go inp.pipeThingSeenAttr(cha, attr)
	return cha
}

// ThingForkSeen returns two channels, `new` and `old`,
// where `new` is to receive
// all `inp`
// not been seen before
// and `old`
// all `inp`
// seen before
// (internally growing a `sync.Map` to discriminate)
// until close.
func (inp ThingFrom) ThingForkSeen() (new, old ThingFrom) {
	cha1 := make(chan Thing)
	cha2 := make(chan Thing)
	go inp.forkThingSeenAttr(cha1, cha2, nil)
	return cha1, cha2
}

// ThingForkSeenAttr returns two channels, `new` and `old`,
// where `new` is to receive
// all `inp`
// whose attribute `attr` has
// not been seen before
// and `old`
// all `inp`
// seen before
// (internally growing a `sync.Map` to discriminate)
// until close.
func (inp ThingFrom) ThingForkSeenAttr(attr func(a Thing) interface{}) (new, old ThingFrom) {
	cha1 := make(chan Thing)
	cha2 := make(chan Thing)
	go inp.forkThingSeenAttr(cha1, cha2, attr)
	return cha1, cha2
}

func (inp ThingFrom) pipeThingSeenAttr(out ThingInto, attr func(a Thing) interface{}) {
	defer close(out)

	if attr == nil { // Make `nil` value useful
		attr = func(a Thing) interface{} { return a }
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

func (inp ThingFrom) forkThingSeenAttr(new, old ThingInto, attr func(a Thing) interface{}) {
	defer close(new)
	defer close(old)

	if attr == nil { // Make `nil` value useful
		attr = func(a Thing) interface{} { return a }
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

// ThingTubeSeen returns a closure around ThingPipeSeen()
// (silently dropping every Thing seen before).
func (inp ThingFrom) ThingTubeSeen() (tube func(inp ThingFrom) (out ThingFrom)) {

	return func(inp ThingFrom) (out ThingFrom) {
		return inp.ThingPipeSeen()
	}
}

// ThingTubeSeenAttr returns a closure around ThingPipeSeenAttr(attr)
// (silently dropping every Thing
// whose attribute `attr` was
// seen before).
func (inp ThingFrom) ThingTubeSeenAttr(attr func(a Thing) interface{}) (tube func(inp ThingFrom) (out ThingFrom)) {

	return func(inp ThingFrom) (out ThingFrom) {
		return inp.ThingPipeSeenAttr(attr)
	}
}

// End of ThingPipeSeen/ThingForkSeen - an "I've seen this Thing before" filter / forker
// ===========================================================================
