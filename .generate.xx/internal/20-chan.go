// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

// ===========================================================================
// Beg of ChanAny producers

// ChanAny returns a channel to receive
// all inputs
// before close.
func ChanAny(inp ...Any) (out Anymode) {
	cha := MakeAnymodeChan()
	go chanAny(cha, inp...)
	return cha
}

func chanAny(out Anymode, inp ...Any) {
	defer out.Close()
	for i := range inp {
		out.Provide(inp[i])
	}
}

// ChanAnySlice returns a channel to receive
// all inputs
// before close.
func ChanAnySlice(inp ...[]Any) (out Anymode) {
	cha := MakeAnymodeChan()
	go chanAnySlice(cha, inp...)
	return cha
}

func chanAnySlice(out Anymode, inp ...[]Any) {
	defer out.Close()
	for i := range inp {
		for j := range inp[i] {
			out.Provide(inp[i][j])
		}
	}
}

// ChanAnyFuncNok returns a channel to receive
// all results of generator `gen`
// until `!ok`
// before close.
func ChanAnyFuncNok(gen func() (Any, bool)) (out Anymode) {
	cha := MakeAnymodeChan()
	go chanAnyFuncNok(cha, gen)
	return cha
}

func chanAnyFuncNok(out Anymode, gen func() (Any, bool)) {
	defer out.Close()
	for {
		res, ok := gen() // generate
		if !ok {
			return
		}
		out.Provide(res)
	}
}

// ChanAnyFuncErr returns a channel to receive
// all results of generator `gen`
// until `err != nil`
// before close.
func ChanAnyFuncErr(gen func() (Any, error)) (out Anymode) {
	cha := MakeAnymodeChan()
	go chanAnyFuncErr(cha, gen)
	return cha
}

func chanAnyFuncErr(out Anymode, gen func() (Any, error)) {
	defer out.Close()
	for {
		res, err := gen() // generate
		if err != nil {
			return
		}
		out.Provide(res)
	}
}

// End of ChanAny producers
// ===========================================================================
