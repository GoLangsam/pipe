// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

import (
	"github.com/cheekybits/genny/generic"
)

type Any generic.Type

// ===========================================================================
// Beg of has nil versions

// Functions suitable only for types which can be == nil.
// Thus, do not use for basic built-in's such as int, string, ...

// ChanAnyFuncNil returns a channel to receive
// all results of generator `gen`
// until nil
// before close.
func ChanAnyFuncNil(gen func() Any) (out <-chan Any) {
	cha := make(chan Any)
	go chanAnyFuncNil(cha, gen)
	return cha
}

func chanAnyFuncNil(out chan<- Any, gen func() Any) {
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
