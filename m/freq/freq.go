// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

import (
	"github.com/cheekybits/genny/generic"
)

// anyThing is the generic type flowing thru the pipe network.
type anyThing generic.Type

// ===========================================================================
// Beg of anyThingDoneFreq - receive a frequency histogram

// anyThingDoneFreq returns a channel to receive
// a frequency histogram (as a `map[anyThing]int64`)
// upon close.
func (inp anyThingFrom) anyThingDoneFreq() (freq <-chan map[anyThing]int64) {
	cha := make(chan map[anyThing]int64)
	go inp.doneanyThingFreq(cha)
	return cha
}

// anyThingDoneFreqAttr returns a channel to receive
// a frequency histogram (as a `map[interface{}]int64`)
// upon close.
//
// `attr` provides the key to the frequency map.
// If `nil` is passed as `attr` then anyThing is used as key.
func (inp anyThingFrom) anyThingDoneFreqAttr(attr func(a anyThing) interface{}) (freq <-chan map[interface{}]int64) {
	cha := make(chan map[interface{}]int64)
	go inp.doneanyThingFreqAttr(cha, attr)
	return cha
}

func (inp anyThingFrom) doneanyThingFreq(out chan<- map[anyThing]int64) {
	defer close(out)
	freq := make(map[anyThing]int64)

	for i := range inp {
		freq[i]++
	}
	out <- freq
}

func (inp anyThingFrom) doneanyThingFreqAttr(out chan<- map[interface{}]int64, attr func(a anyThing) interface{}) {
	defer close(out)
	freq := make(map[interface{}]int64)

	if attr == nil { // Make `nil` value useful
		attr = func(a anyThing) interface{} { return a }
	}

	for i := range inp {
		freq[attr(i)]++
	}
	out <- freq
}

// End of anyThingDoneFreq - receive a frequency histogram
// ===========================================================================
