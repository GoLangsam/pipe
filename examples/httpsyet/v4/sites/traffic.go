// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sites

import (
	"net/url"
	"sync"
)

type site = Site

// Traffic goes around inside a circular Site pipe network,
// e. g. a crawling Crawler.
type Traffic struct {
	sites chan site       // to be processed
	wg    *sync.WaitGroup // monitor SiteEnter & SiteLeave
	done  chan struct{}   // to signal termination due to traffic having subsided
	once  *sync.Once      // to close Done only once - lauched from first feed
}

// New returns a new and operational Traffic processor.
func New() (t *Traffic) {
	return &Traffic{
		make(chan site),
		new(sync.WaitGroup),
		make(chan struct{}),
		new(sync.Once),
	}
}

// Done returns a channel which will be signalled and closed
// when traffic has subsided and nothing is left to be processed
// and consequently all goroutines have terminated.
func (t *Traffic) Done() (done <-chan struct{}) {
	return t.done
}

// Feed registers new entries.
func (t *Traffic) Feed(urls []*url.URL, parent *url.URL, depth int) {
	queueURLs(t.sites, urls, parent, depth)

	t.once.Do(func() {
		go func(t *Traffic) {
			t.done <- <-siteDoneWait(t.sites, t.wg)
			close(t.done)
		}(t)
	})
}

// Processor builds the site traffic processing network;
// it is cirular if crawl uses Feed to provide feedback.
func (t *Traffic) Processor(crawl func(s Site), parallel int) {
	proc := func(s Site) { // wrap crawl:
		crawl(s)    // apply original crawl
		t.wg.Done() // have this site leave
	}

	sites, seen := siteForkSeenAttr(sitePipeEnter(t.sites, t.wg), Site.attr)
	for _, sites := range siteStrew(sitePipeAdjust(sites), parallel) {
		siteDoneFunc(sites, proc) // strewed `sites` leave in wrapped `crawl`
	}
	siteDone(sitePipeLeave(seen, t.wg)) // `seen` leave without further processing
}
