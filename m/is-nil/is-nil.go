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
// Beg of has nil versions

// Functions suitable only for types which can be == nil.
// Thus, do not use for basic built-in's such as int, string, ...

// ChanFuncNil returns a channel to receive
// all results of generator `gen`
// until nil
// before close.
func ChanFuncNil(gen func() anyThing) (out anyThingFrom) {
	cha := make(chan anyThing)
	go chanFuncNil(cha, gen)
	return cha
}

func chanFuncNil(out anyThingInto, gen func() anyThing) {
	defer close(out)
	for {
		res := gen() // generate
		if res == nil {
			return
		}
		out <- res
	}
}

// End of has nil versions
// ===========================================================================
