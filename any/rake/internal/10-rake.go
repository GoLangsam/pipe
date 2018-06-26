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
	items chan item       // to be processed
	wg    *sync.WaitGroup // monitor SiteEnter & SiteLeave
	done  chan struct{}   // to signal termination due to traffic having subsided
	once  *sync.Once      // to close Done only once - lauched from first feed
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
	}

	proc := func(a item) { // wrap rake:
		rake(a)      // apply original rake
		my.wg.Done() // have this item leave
	}

	// build the concurrent pipe network
	items, seen := my.itemForkSeenAttr(my.items, attr)
	_ = my.itemDoneLeave(seen, my.wg) // `seen` leave without further processing

	for _, items := range my.itemStrew(my.itemPipeAdjust(items), somany) {
		_ = my.itemDoneFunc(items, proc) // strewed `items` leave in wrapped `crawl`
	}

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
		return my
	}

	my.wg.Add(len(items))
	for _, i := range items {
		my.items <- i
	}

	my.once.Do(func() {
		go func(t *Rake) {
			my.done <- <-my.itemDoneWait(my.items, my.wg)
			close(my.done)
		}(my)
	})

	return my

}

// End of Rake
// ===========================================================================
