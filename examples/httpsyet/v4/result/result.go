// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package result

// Result represents (a secondary) observation
type Result string

type result = Result // to keep generated functions private

// DoneFunc returns a channel to receive
// one signal after `act` has been applied to every `inp`
// before close.
func DoneFunc(inp <-chan Result, act func(a Result)) (done <-chan struct{}) {
	return resultDone(inp, act)
}
