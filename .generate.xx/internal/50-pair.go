// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

// ===========================================================================
// Beg of PairAny functions

// PairAny returns a pair of channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PairAny(inp Anymode) (out1, out2 Anymode) {
	cha1 := MakeAnymodeChan()
	cha2 := MakeAnymodeChan()
	go pairAny(cha1, cha2, inp)
	return cha1, cha2
}

func pairAny(out1, out2 Anymode, inp Anymode) {
	defer out1.Close()
	defer out2.Close()
	for i, ok := inp.Request(); ok; i, ok = inp.Request() {
		out1.Provide(i)
		out2.Provide(i)
	}
}

// End of PairAny functions
// ===========================================================================
