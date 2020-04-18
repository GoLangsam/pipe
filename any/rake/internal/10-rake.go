// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rake

import (
	"sync"
)

// Rake represents a fanned out circular pipe network
// with a flexibly adjusting buffer.
// Any item is processed once only -
// items seen before are filtered out.
//
// A Rake may be used e.g. as a crawling Crawler
// where every link shall be visited only once.
type Rake struct {
	items chan item                // to be processed
	wg    *sync.WaitGroup          // monitor SiteEnter & SiteLeave
	done  chan struct{}            // to signal termination due to traffic having subsided
	once  *sync.Once               // to close Done only once - lauched from first feed
	rake  func(a item)             // function to be applied
	attr  func(a item) interface{} // attribute to discriminate seen
	runs  bool                     // am I running?
	many  int                      // # of parallel raking endpoints of the Rake
}

// New returns a (pointer to a) new operational Rake.
//
// `rake` is the operation to be executed in parallel on any item
// which has not been seen before.
// Have it use `myrake.Feed(items...)` in order to provide feed-back.
//
// `attr` allows to specify an attribute for the seen filter.
// Pass `nil` to filter on any item itself.
//
// `somany` is the # of parallel processes - the parallelism
// of the network built by Rake,
// the # of parallel raking endpoints of the Rake.
func New(
	rake func(a item),
	attr func(a item) interface{},
	somany int,
) (
	my *Rake,
) {
	if somany < 1 {
		somany = 1
	}
	my = &Rake{
		make(chan item),
		new(sync.WaitGroup),
		make(chan struct{}),
		new(sync.Once),
		rake,
		attr,
		false,
		somany,
	}

	return my
}

// init builds the network
func (my *Rake) init() *Rake {
	proc := func(a item) { // wrap rake:
		my.rake(a)   // apply original rake
		my.wg.Done() // have this item leave
	}

	// build the concurrent pipe network
	items, seen := (itemFrom)(my.items).ForkSeenAttr(my.attr)
	_ = seen.DoneLeave(my.wg) // `seen` leave without further processing

	for _, items := range items.PipeAdjust().Strew(my.many) {
		_ = items.Done(proc) // strewed `items` leave in wrapped `crawl`
	}

	return my
}

// start builds the network and spawns the closer
func (my *Rake) start() {
	my = my.init()
	my.runs = true
	go my.closer()
}

func (my *Rake) closer() *Rake {
	my.done <- <-(itemInto)(my.items).DoneWait(my.wg)
	close(my.done)
	return my
}

// checkRuns for paranoids
func (my *Rake) checkRuns() *Rake {
	if my.runs {
		panic("Rake is running already")
	}
	return my
}

// Rake sets the rake function to be applied (in parallel).
//
// `rake` is the operation to be executed in parallel on any item
// which has not been seen before.
//
// You may provide `nil` here and call `Rake(..)` later to provide it.
// Or have it use `myrake.Feed(items...)` in order to provide feed-back.
//
// Rake panics iff called after first nonempty `Feed(...)`
func (my *Rake) Rake(rake func(a item)) *Rake {
	my.checkRuns()
	my.rake = rake
	return my
}

// Attr sets the (optional) attribute to discriminate 'seen'.
//
// `attr` allows to specify an attribute for the 'seen' filter.
// If not set 'seen' will discriminate any item by itself.
//
// Seen panics iff called after first nonempty `Feed(...)`
func (my *Rake) Attr(attr func(a item) interface{}) *Rake {
	my.checkRuns()
	my.attr = attr
	return my
}

// Done returns a channel which will be signalled and closed
// when traffic has subsided, nothing is left to be processed
// and consequently all goroutines have terminated.
func (my *Rake) Done() (done <-chan struct{}) {
	return my.done
}

// Feed registers new items on the network.
func (my *Rake) Feed(items ...item) *Rake {

	if len(items) == 0 {
		return my // nothing to do
	}

	my.wg.Add(len(items)) // items enter

	my.once.Do(my.start) // lazy init: build & start the network

	for _, i := range items {
		my.items <- i
	}

	return my
}

// End of Rake
// ===========================================================================
