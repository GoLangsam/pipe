// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package httpsyet

// result represents (a secondary) observation
type result string

// DoneFunc returns a channel to receive
// one signal after `act` has been applied to every `inp`
// before close.
func DoneFunc(inp resultFrom, act func(a result)) (done <-chan struct{}) {
	return inp.Done(act)
}
