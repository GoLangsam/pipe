// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

import (
	"fmt"
)

func sendOne(snd anyThingInto) {
	defer close(snd)
	snd <- 1 // send a 1
}

func sendTwo(snd anyThingInto) {
	defer close(snd)
	snd <- 1 // send a 1
	snd <- 2 // send a 2
}

func ExampleanyThingDaisyChaiN_tenthousand() {

	right := make(chan anyThing)

	nPlusOne := func(left anyThingInto, right anyThingFrom) { // left <- 1 + <-right }
		r := <-right
		left <- 1 + r.(int)
	}

	leftmost := anyThingDaisyChaiN(right, 10000, nPlusOne) // the chain - right to left!

	go sendOne(right)
	// sendTwo(right)

	fmt.Println(<-leftmost)
	// Output:
	// 10001
}
