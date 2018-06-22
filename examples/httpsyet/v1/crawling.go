// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package httpsyet

import (
	"fmt"
	"net/url"
	"sync"
	"time"
)

// ===========================================================================

// report prints a result to Crawler.Out;
// used by `DoneStringFunc`.
func (c Crawler) report(result string) {
	if _, err := fmt.Fprintln(c.Out, result); err != nil {
		c.Log.Printf("failed to write output '%s': %v\n", result, err)
	}
}

// ---------------------------------------------------------------------------

// crawling represents a crawling Crawler ... busy crawling ...
type crawling struct {
	Crawler                     // config
	sites           chan site   // to be crawled
	*sync.WaitGroup             // for SiteEnter & SiteLeave
	results         chan string // to be reported
}

// ---------------------------------------------------------------------------
// teach `*crawling` some straight-forward behaviour, and how to crawl :-)

// add registers new entries and launches their dispatcher
// (which we intentionally left untouched).
func (c *crawling) add(urls []*url.URL, parent *url.URL, depth int) {
	c.Add(len(urls))
	go queueURLs(c.sites, urls, parent, depth)
}

// goWaitAndClose is to be used after initial traffic has been added.
func (c *crawling) goWaitAndClose() {
	go func(c *crawling) {
		c.Wait()
		close(c.sites)
		close(c.results)
	}(c)
}

// ===========================================================================

// crawl performs a crawling Crawler's main function: crawl.
//
// This version attempts to respect the original implementation.
// Thus, it is still too busy - catering with too many concerns;
// concerns which might better be catered for in the processor's pipe network;
// used by `DoneSiteFunc`.
func (c *crawling) crawl(s site) {
	urls := c.crawlSite(s)        // the core crawl process
	c.add(urls, s.URL, s.Depth-1) // new urls enter crawling - circular feedback
	c.Done()                      // site leaves crawling
	time.Sleep(c.Delay)           // have a gentle nap
}

// ===========================================================================

// a Crawler crawling: crawl traffic, emit results
// and signal when done
func (c Crawler) crawling(urls []*url.URL) (done <-chan struct{}) {
	crawling := crawling{
		c,                   // "Crawler is used as configuration ..."
		make(chan site),     // the feedback traffic
		new(sync.WaitGroup), // monitor traffic
		make(chan string),   // the (secondary) output
	}
	return crawling.crawling(urls, parallel(c.Parallel))
}

// ===========================================================================

// crawling builds and feeds the network, and
// returns a channel to receive
// one signal after
// all traffic has subsided.
func (c *crawling) crawling(urls []*url.URL, size int) (done <-chan struct{}) {
	c.processor(size)         // build the process network
	c.add(urls, nil, c.Depth) // feed initial urls
	c.goWaitAndClose()        // launch the closer
	return stringDoneFunc(c.results, c.report)
}

// processor builds our little site processing network;
// its cirular due to c.crawl's feedback.
func (c *crawling) processor(size int) {
	sites, seen := siteForkSeenAttr(c.sites, site.Attr)
	for _, inp := range siteStrew(sites, size) {
		siteDoneFunc(inp, c.crawl) // sites leave inside crawler's crawl
	}
	siteDoneLeave(seen, c) // seen leave without further processing
}

// ===========================================================================
