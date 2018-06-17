// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

// ===========================================================================
// Beg of PipeAny functions

// PipeAnyFunc returns a channel to receive
// every result of action `act` applied to `inp`
// before close.
// Note: it 'could' be PipeAnyMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeAnyFunc(inp Anymode, act func(a Any) Any) (out Anymode) {
	cha := MakeAnymodeChan()
	if act == nil {
		act = func(a Any) Any { return a }
	}
	go pipeAnyFunc(cha, inp, act)
	return cha
}

func pipeAnyFunc(out Anymode, inp Anymode, act func(a Any) Any) {
	defer out.Close()
	for i, ok := inp.Request(); ok; i, ok = inp.Request() {
		out.Provide(act(i))
	}
}

// PipeAnyBuffer returns a buffered channel with capacity `cap` to receive
// all `inp`
// before close.
func PipeAnyBuffer(inp Anymode, cap int) (out Anymode) {
	cha := MakeAnymodeBuff(cap)
	go pipeAnyBuffer(cha, inp)
	return cha
}

func pipeAnyBuffer(out Anymode, inp Anymode) {
	defer out.Close()
	for i, ok := inp.Request(); ok; i, ok = inp.Request() {
		out.Provide(i)
	}
}

// End of PipeAny functions
// ===========================================================================
