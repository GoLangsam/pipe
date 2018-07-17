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
// Beg of DoneFreq - receive a frequency histogram

// DoneFreq returns a channel to receive
// a frequency histogram (as a `map[anyThing]int64`)
// upon close.
func (inp anyThingFrom) DoneFreq() (freq <-chan map[anyThing]int64) {
	cha := make(chan map[anyThing]int64)
	go inp.doneFreq(cha)
	return cha
}

// DoneFreqAttr returns a channel to receive
// a frequency histogram (as a `map[interface{}]int64`)
// upon close.
//
// `attr` provides the key to the frequency map.
// If `nil` is passed as `attr` then anyThing is used as key.
func (inp anyThingFrom) DoneFreqAttr(attr func(a anyThing) interface{}) (freq <-chan map[interface{}]int64) {
	cha := make(chan map[interface{}]int64)
	go inp.doneFreqAttr(cha, attr)
	return cha
}

func (inp anyThingFrom) doneFreq(out chan<- map[anyThing]int64) {
	defer close(out)
	freq := make(map[anyThing]int64)

	for i := range inp {
		freq[i]++
	}
	out <- freq
}

func (inp anyThingFrom) doneFreqAttr(out chan<- map[interface{}]int64, attr func(a anyThing) interface{}) {
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

// End of DoneFreq - receive a frequency histogram
// ===========================================================================
