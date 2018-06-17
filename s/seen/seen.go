// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

import (
	"sync"

	"github.com/cheekybits/genny/generic"
)

type Any generic.Type

// ===========================================================================
// Beg of PipeAnySeen/ForkAnySeen - an "I've seen this Any before" filter / fork

// PipeAnySeen returns a channel to receive
// all `inp`
// not been seen before
// while silently dropping everything seen before
// (internally growing a `sync.Map` to discriminate)
// until close.
//  Note: PipeAnyFilterNotSeenYet might be a better name, but is fairly long.
func PipeAnySeen(inp <-chan Any) (out <-chan Any) {
	cha := make(chan Any)
	go pipeAnySeenAttr(cha, inp, nil)
	return cha
}

// PipeAnySeenAttr returns a channel to receive
// all `inp`
// whose attribute `attr` has
// not been seen before
// while silently dropping everything seen before
// (internally growing a `sync.Map` to discriminate)
// until close.
//  Note: PipeAnyFilterAttrNotSeenYet might be a better name, but is fairly long.
func PipeAnySeenAttr(inp <-chan Any, attr func(a Any) interface{}) (out <-chan Any) {
	cha := make(chan Any)
	go pipeAnySeenAttr(cha, inp, attr)
	return cha
}

// ForkAnySeen returns two channels, `new` and `old`,
// where `new` is to receive
// all `inp`
// not been seen before
// and `old`
// all `inp`
// seen before
// (internally growing a `sync.Map` to discriminate)
// until close.
func ForkAnySeen(inp <-chan Any) (new, old <-chan Any) {
	cha1 := make(chan Any)
	cha2 := make(chan Any)
	go forkAnySeenAttr(cha1, cha2, inp, nil)
	return cha1, cha2
}

// ForkAnySeenAttr returns two channels, `new` and `old`,
// where `new` is to receive
// all `inp`
// whose attribute `attr` has
// not been seen before
// and `old`
// all `inp`
// seen before
// (internally growing a `sync.Map` to discriminate)
// until close.
func ForkAnySeenAttr(inp <-chan Any, attr func(a Any) interface{}) (new, old <-chan Any) {
	cha1 := make(chan Any)
	cha2 := make(chan Any)
	go forkAnySeenAttr(cha1, cha2, inp, attr)
	return cha1, cha2
}

func pipeAnySeenAttr(out chan<- Any, inp <-chan Any, attr func(a Any) interface{}) {
	defer close(out)

	if attr == nil { // Make `nil` value useful
		attr = func(a Any) interface{} { return a }
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

func forkAnySeenAttr(new, old chan<- Any, inp <-chan Any, attr func(a Any) interface{}) {
	defer close(new)
	defer close(old)

	if attr == nil { // Make `nil` value useful
		attr = func(a Any) interface{} { return a }
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

// TubeAnySeen returns a closure around PipeAnySeen()
// (silently dropping every Any seen before).
func TubeAnySeen() (tube func(inp <-chan Any) (out <-chan Any)) {

	return func(inp <-chan Any) (out <-chan Any) {
		return PipeAnySeen(inp)
	}
}

// TubeAnySeenAttr returns a closure around PipeAnySeenAttr()
// (silently dropping every Any
// whose attribute `attr` was
// seen before).
func TubeAnySeenAttr(attr func(a Any) interface{}) (tube func(inp <-chan Any) (out <-chan Any)) {

	return func(inp <-chan Any) (out <-chan Any) {
		return PipeAnySeenAttr(inp, attr)
	}
}

// End of PipeAnySeen/ForkAnySeen - an "I've seen this Any before" filter / fork
// ===========================================================================
