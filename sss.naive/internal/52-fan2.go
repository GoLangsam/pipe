// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

// ===========================================================================
// FanIn2Any as seen in Go Concurrency Patterns
//
// Warning: For instruction and teaching only!
// Do not use in any serious project, as
// it hangs forever upon close of both inputs.
// Thus: it leaks it's goroutine!
// (And never closes it's output)
func FanIn2Any(inp1, inp2 chan Any) chan Any {
	out := make(chan Any)
	go func() {
		for {
			select {
			case e := <-inp1:
				out <- e
			case e := <-inp2:
				out <- e
			}
		}
	}()
	return out
}

// ===========================================================================
