// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

import (
	"sync"

	"github.com/cheekybits/genny/generic"
)

type Any generic.Type

// ===========================================================================
// Beg of FanAnysIn

// FanAnysIn returns a channel to receive all inputs arriving
// on variadic inps
// before close.
//
//  Ref: https://blog.golang.org/pipelines
//  Ref: https://github.com/QuentinPerez/go-stuff/channel/Fan-out-Fan-in/main.go
func FanAnysIn(inps ...<-chan Any) (out <-chan Any) {
	cha := make(chan Any)

	wg := new(sync.WaitGroup)
	wg.Add(len(inps))

	go func(wg *sync.WaitGroup, out chan Any) { // Spawn "close(out)" once all inps are done
		wg.Wait()
		close(out)
	}(wg, cha)

	for i := range inps {
		go func(out chan<- Any, inp <-chan Any) { // Spawn "output(c)"s
			defer wg.Done()
			for i := range inp {
				out <- i
			}
		}(cha, inps[i])
	}

	return cha
}

// End of FanAnysIn
// ===========================================================================
