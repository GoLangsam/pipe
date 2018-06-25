// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

import (
	"fmt"
)

func Naturals() func() int {
	var n int
	return func() int {
		n++
		return n
	}
}

func TimesN(times int) func(int) int {
	return func(i int) int {
		return times * i
	}
}

func p(i int) {
	fmt.Println(i)
}

func Example_one() {

	t := intChan(1, 2, 3, 4, 5)
	t = intPipeFunc(t, TimesN(4))
	<-intDoneFunc(t, p)

	// Output:
	// 4
	// 8
	// 12
	// 16
	// 20
}

func Example_two() {

	t := intChan(1, 2, 3, 4, 5)
	t4 := intTubeFunc(TimesN(4))
	<-intDoneFunc(t4(t), p)

	// Output:
	// 4
	// 8
	// 12
	// 16
	// 20
}

func Example_three() {

	t := intChan(1, 2, 3, 4, 5)
	t2 := intTubeFunc(TimesN(2))
	<-intDoneFunc(t2(t2(t)), p)

	// Output:
	// 4
	// 8
	// 12
	// 16
	// 20
}

func Example_daisy() {

	t := intChan(1, 2, 3, 4, 5)
	t2 := intTubeFunc(TimesN(2))
	<-intDoneFunc(t2(t2(t)), p)

	// Output:
	// 4
	// 8
	// 12
	// 16
	// 20
}

func Example_same1() {

	c1 := intChan(1, 2, 3, 4, 5)
	c2 := intChan(1, 2, 3, 4, 5)

	same := func(a, b int) bool { return a == b }
	fmt.Println(<-intSame(same, c1, c2))

	// Output:
	// true
}

func Example_same2() {

	c1 := intChan(1, 2, 3, 4, 5)
	c2 := intChan(1, 2, 3, 4, 5, 6)

	same := func(a, b int) bool { return a == b }
	fmt.Println(<-intSame(same, c1, c2))

	// Output:
	// false
}
