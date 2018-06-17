// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

// ===========================================================================
// Beg of ChanAny producers

// ChanAny returns a channel to receive
// all inputs
// before close.
func ChanAny(inp ...Any) (out <-chan Any) {
	cha := make(chan Any)
	go chanAny(cha, inp...)
	return cha
}

func chanAny(out chan<- Any, inp ...Any) {
	defer close(out)
	for i := range inp {
		out <- inp[i]
	}
}

// ChanAnySlice returns a channel to receive
// all inputs
// before close.
func ChanAnySlice(inp ...[]Any) (out <-chan Any) {
	cha := make(chan Any)
	go chanAnySlice(cha, inp...)
	return cha
}

func chanAnySlice(out chan<- Any, inp ...[]Any) {
	defer close(out)
	for i := range inp {
		for j := range inp[i] {
			out <- inp[i][j]
		}
	}
}

// ChanAnyFuncNok returns a channel to receive
// all results of generator `gen`
// until `!ok`
// before close.
func ChanAnyFuncNok(gen func() (Any, bool)) (out <-chan Any) {
	cha := make(chan Any)
	go chanAnyFuncNok(cha, gen)
	return cha
}

func chanAnyFuncNok(out chan<- Any, gen func() (Any, bool)) {
	defer close(out)
	for {
		res, ok := gen() // generate
		if !ok {
			return
		}
		out <- res
	}
}

// ChanAnyFuncErr returns a channel to receive
// all results of generator `gen`
// until `err != nil`
// before close.
func ChanAnyFuncErr(gen func() (Any, error)) (out <-chan Any) {
	cha := make(chan Any)
	go chanAnyFuncErr(cha, gen)
	return cha
}

func chanAnyFuncErr(out chan<- Any, gen func() (Any, error)) {
	defer close(out)
	for {
		res, err := gen() // generate
		if err != nil {
			return
		}
		out <- res
	}
}

// End of ChanAny producers
// ===========================================================================
