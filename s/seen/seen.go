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
// Beg of anyThingPipeSeen/anyThingForkSeen - an "I've seen this anyThing before" filter / forker

// anyThingPipeSeen returns a channel to receive
// all `inp`
// not been seen before
// while silently dropping everything seen before
// (internally growing a `sync.Map` to discriminate)
// until close.
//  Note: anyThingPipeFilterNotSeenYet might be a better name, but is fairly long.
func anyThingPipeSeen(inp <-chan anyThing) (out <-chan anyThing) {
	cha := make(chan anyThing)
	go pipeanyThingSeenAttr(cha, inp, nil)
	return cha
}

// anyThingPipeSeenAttr returns a channel to receive
// all `inp`
// whose attribute `attr` has
// not been seen before
// while silently dropping everything seen before
// (internally growing a `sync.Map` to discriminate)
// until close.
//  Note: anyThingPipeFilterAttrNotSeenYet might be a better name, but is fairly long.
func anyThingPipeSeenAttr(inp <-chan anyThing, attr func(a anyThing) interface{}) (out <-chan anyThing) {
	cha := make(chan anyThing)
	go pipeanyThingSeenAttr(cha, inp, attr)
	return cha
}

// anyThingForkSeen returns two channels, `new` and `old`,
// where `new` is to receive
// all `inp`
// not been seen before
// and `old`
// all `inp`
// seen before
// (internally growing a `sync.Map` to discriminate)
// until close.
func anyThingForkSeen(inp <-chan anyThing) (new, old <-chan anyThing) {
	cha1 := make(chan anyThing)
	cha2 := make(chan anyThing)
	go forkanyThingSeenAttr(cha1, cha2, inp, nil)
	return cha1, cha2
}

// anyThingForkSeenAttr returns two channels, `new` and `old`,
// where `new` is to receive
// all `inp`
// whose attribute `attr` has
// not been seen before
// and `old`
// all `inp`
// seen before
// (internally growing a `sync.Map` to discriminate)
// until close.
func anyThingForkSeenAttr(inp <-chan anyThing, attr func(a anyThing) interface{}) (new, old <-chan anyThing) {
	cha1 := make(chan anyThing)
	cha2 := make(chan anyThing)
	go forkanyThingSeenAttr(cha1, cha2, inp, attr)
	return cha1, cha2
}

func pipeanyThingSeenAttr(out chan<- anyThing, inp <-chan anyThing, attr func(a anyThing) interface{}) {
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

func forkanyThingSeenAttr(new, old chan<- anyThing, inp <-chan anyThing, attr func(a anyThing) interface{}) {
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

// anyThingTubeSeen returns a closure around anyThingPipeSeen()
// (silently dropping every anyThing seen before).
func anyThingTubeSeen() (tube func(inp <-chan anyThing) (out <-chan anyThing)) {

	return func(inp <-chan anyThing) (out <-chan anyThing) {
		return anyThingPipeSeen(inp)
	}
}

// anyThingTubeSeenAttr returns a closure around anyThingPipeSeenAttr()
// (silently dropping every anyThing
// whose attribute `attr` was
// seen before).
func anyThingTubeSeenAttr(attr func(a anyThing) interface{}) (tube func(inp <-chan anyThing) (out <-chan anyThing)) {

	return func(inp <-chan anyThing) (out <-chan anyThing) {
		return anyThingPipeSeenAttr(inp, attr)
	}
}

// End of anyThingPipeSeen/anyThingForkSeen - an "I've seen this anyThing before" filter / forker
// ===========================================================================
