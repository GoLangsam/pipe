// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

import (
	"fmt"
)

func sendOne(snd chan<- anyThing) {
	defer close(snd)
	snd <- 1 // send a 1
}

func sendTwo(snd chan<- anyThing) {
	defer close(snd)
	snd <- 1 // send a 1
	snd <- 2 // send a 2
}

func ExampleanyThingDaisyChaiN_tenthousand() {
	var my anyOwner // so we may call his methods

	right := make(chan anyThing)

	nPlusOne := func(left chan<- anyThing, right <-chan anyThing) { // left <- 1 + <-right }
		r := <-right
		left <- 1 + r.(int)
	}

	leftmost := my.anyThingDaisyChaiN(right, 10000, nPlusOne) // the chain - right to left!

	go sendOne(right)
	// sendTwo(right)

	fmt.Println(<-leftmost)
	// Output:
	// 10001
}
