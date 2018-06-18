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

// crawling represents a crawling Crawler
type crawling struct {
	Crawler                     // config
	*sync.WaitGroup             // for SiteEnter & SiteLeave
	*sync.Once                  // for wait - and for the paranoid :-)
	sites           chan site   // to be crawled
	results         chan string // to be reported
}

// ---------------------------------------------------------------------------
// site learning a little behaviour

// attr returns the attribute relevant for ForkSiteSeenAttr,
// the "I've seen this site before" discriminator.
func (s site) attr() interface{} {
	return s.URL.String()
}

// print may be used via e.g. PipeSiteFunc(sites, site.print) for tracing
func (s site) print() site {
	fmt.Println(s)
	return s
}

// ---------------------------------------------------------------------------
// *crawling learning some straight-forward behaviour, and how to crawl :-)

// add registers new entries and launches their dispatcher
// (which we intentionally left untouched).
func (c *crawling) add(urls []*url.URL, parent *url.URL, depth int) {
	c.Add(len(urls))
	go queueURLs(c.sites, urls, parent, depth)
}

// wait is to be launched after initial traffic has been added.
//  Note: Any crawling assures `wait` to be launched only once :-)
func (c *crawling) wait() {
	c.Do(func() {
		go func() {
			c.Wait()
			close(c.sites)
			close(c.results)
		}()
	})
}

// crawl performs a crawling Crawler's main function: crawl.
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

// ---------------------------------------------------------------------------

// report prints a result to Crawler.Out;
// used by `DoneStringFunc`.
func (c *crawling) report(result string) {
	if _, err := fmt.Fprintln(c.Out, result); err != nil {
		c.Log.Printf("failed to write output '%s': %v\n", result, err)
	}
}

// ===========================================================================

// A crawling Crawler ... busy crawling ...
func (c Crawler) crawling(urls []*url.URL) (done <-chan struct{}) {
	crawling := crawling{
		c,                   // "Crawler is used as configuration ..."
		new(sync.WaitGroup), // monitor traffic
		new(sync.Once),      // `wait` & close only once
		make(chan site),     // the feedback traffic
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
	go c.wait()               // launch the closer
	return DoneStringFunc(c.results, c.report)
}

// processor builds our little site processing network;
// its cirular due to c.crawl's feedback.
func (c *crawling) processor(size int) {
	sites, seen := ForkSiteSeenAttr(c.sites, site.attr)
	for _, inp := range ScatterSite(sites, size) {
		DoneSiteFunc(inp, c.crawl) // sites leave inside crawler's crawl
	}
	DoneSite(PipeSiteLeave(seen, c)) // seen leave without further processing
}

// ===========================================================================
