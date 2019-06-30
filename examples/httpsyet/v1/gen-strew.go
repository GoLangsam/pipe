// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package httpsyet

import "time"

// ===========================================================================
// Beg of siteStrew - scatter them

// siteStrew returns a slice (of size = size) of channels
// one of which shall receive each inp before close.
func siteStrew(inp <-chan site, size int) (outS [](<-chan site)) {
	chaS := make(map[chan site]struct{}, size)
	for i := 0; i < size; i++ {
		chaS[make(chan site)] = struct{}{}
	}

	go strewSite(inp, chaS)

	outS = make([]<-chan site, size)
	i := 0
	for c := range chaS {
		outS[i] = (<-chan site)(c) // convert `chan` to `<-chan`
		i++
	}

	return outS
}

// c strewSite(inp <-chan site, outS ...chan<- site) {
// Note: go does not convert the passed slice `[]chan site` to `[]chan<- site` automatically.
// So, we do neither here, as we are lazy (we just call an internal helper function).
func strewSite(inp <-chan site, outS map[chan site]struct{}) {

	for i := range inp {
		for !trySendSite(i, outS) {
			time.Sleep(time.Millisecond * 10) // wait a little before retry
		} // !sent
	} // inp

	for o := range outS {
		close(o)
	}
}

func trySendSite(inp site, outS map[chan site]struct{}) bool {

	for o := range outS {

		select { // try to send
		case o <- inp:
			return true
		default:
			// keep trying
		}

	} // outS
	return false
}

// End of siteStrew - scatter them
// ===========================================================================
