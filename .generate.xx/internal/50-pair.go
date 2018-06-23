// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

// ===========================================================================
// Beg of anyThingPair functions

// anyThingPair returns a pair of channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func anyThingPair(inp anymode) (out1, out2 anymode) {
	cha1 := anymodeMakeChan()
	cha2 := anymodeMakeChan()
	go pairanyThing(cha1, cha2, inp)
	return cha1, cha2
}

func pairanyThing(out1, out2 anymode, inp anymode) {
	defer out1.Close()
	defer out2.Close()
	for i, ok := inp.Request(); ok; i, ok = inp.Request() {
		out1.Provide(i)
		out2.Provide(i)
	}
}

// End of anyThingPair functions
// ===========================================================================
