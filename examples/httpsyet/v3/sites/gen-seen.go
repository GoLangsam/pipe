// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sites

import "sync"

// ===========================================================================
// Beg of PipeSiteSeen/ForkSiteSeen - an "I've seen this Site before" filter / fork

// PipeSiteSeen returns a channel to receive
// all `inp`
// not been seen before
// while silently dropping everything seen before
// (internally growing a `sync.Map` to discriminate)
// until close.
// Note: PipeSiteFilterNotSeenYet might be a better name, but is fairly long.
func PipeSiteSeen(inp <-chan Site) (out <-chan Site) {
	cha := make(chan Site)
	go pipeSiteSeenAttr(cha, inp, nil)
	return cha
}

// PipeSiteSeenAttr returns a channel to receive
// all `inp`
// whose attribute `attr` has
// not been seen before
// while silently dropping everything seen before
// (internally growing a `sync.Map` to discriminate)
// until close.
// Note: PipeSiteFilterAttrNotSeenYet might be a better name, but is fairly long.
func PipeSiteSeenAttr(inp <-chan Site, attr func(a Site) interface{}) (out <-chan Site) {
	cha := make(chan Site)
	go pipeSiteSeenAttr(cha, inp, attr)
	return cha
}

// ForkSiteSeen returns two channels, `new` and `old`,
// where `new` is to receive
// all `inp`
// not been seen before
// and `old`
// all `inp`
// seen before
// (internally growing a `sync.Map` to discriminate)
// until close.
func ForkSiteSeen(inp <-chan Site) (new, old <-chan Site) {
	cha1 := make(chan Site)
	cha2 := make(chan Site)
	go forkSiteSeenAttr(cha1, cha2, inp, nil)
	return cha1, cha2
}

// ForkSiteSeenAttr returns two channels, `new` and `old`,
// where `new` is to receive
// all `inp`
// whose attribute `attr` has
// not been seen before
// and `old`
// all `inp`
// seen before
// (internally growing a `sync.Map` to discriminate)
// until close.
func ForkSiteSeenAttr(inp <-chan Site, attr func(a Site) interface{}) (new, old <-chan Site) {
	cha1 := make(chan Site)
	cha2 := make(chan Site)
	go forkSiteSeenAttr(cha1, cha2, inp, attr)
	return cha1, cha2
}

func pipeSiteSeenAttr(out chan<- Site, inp <-chan Site, attr func(a Site) interface{}) {
	defer close(out)

	if attr == nil { // Make `nil` value useful
		attr = func(a Site) interface{} { return a }
	}

	seen := sync.Map{}
	for i := range inp {
		if _, visited := seen.LoadOrStore(attr(i), struct{}{}); visited {
			// drop i silently
		} else {
			out <- i
		}
	}
}

func forkSiteSeenAttr(new, old chan<- Site, inp <-chan Site, attr func(a Site) interface{}) {
	defer close(new)
	defer close(old)

	if attr == nil { // Make `nil` value useful
		attr = func(a Site) interface{} { return a }
	}

	seen := sync.Map{}
	for i := range inp {
		if _, visited := seen.LoadOrStore(attr(i), struct{}{}); visited {
			old <- i
		} else {
			new <- i
		}
	}
}

// TubeSiteSeen returns a closure around PipeSiteSeen()
// (silently dropping every Site seen before).
func TubeSiteSeen() (tube func(inp <-chan Site) (out <-chan Site)) {

	return func(inp <-chan Site) (out <-chan Site) {
		return PipeSiteSeen(inp)
	}
}

// TubeSiteSeenAttr returns a closure around PipeSiteSeenAttr()
// (silently dropping every Site
// whose attribute `attr` was
// seen before).
func TubeSiteSeenAttr(attr func(a Site) interface{}) (tube func(inp <-chan Site) (out <-chan Site)) {

	return func(inp <-chan Site) (out <-chan Site) {
		return PipeSiteSeenAttr(inp, attr)
	}
}

// End of PipeSiteSeen/ForkSiteSeen - an "I've seen this Site before" filter / fork