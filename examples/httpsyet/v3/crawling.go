// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package httpsyet

import (
	"fmt"
	"net/url"
	"sync"
	"time"

	res "github.com/GoLangsam/pipe/examples/httpsyet/v3/result"
	"github.com/GoLangsam/pipe/examples/httpsyet/v3/sites"
)

// ---------------------------------------------------------------------------

// alias types in order to leave source code unaffected
type site = sites.Site
type traffic = sites.Traffic

// ---------------------------------------------------------------------------
// make `result` an explicit type, and
// teach `Crawler` how to `report` a result.

type result = res.Result

// report prints a result to Crawler.Out;
// used by `DoneResultFunc`.
func (c Crawler) report(r result) {
	if _, err := fmt.Fprintln(c.Out, r); err != nil {
		c.Log.Printf("failed to write output '%s': %v\n", r, err)
	}
}

// ---------------------------------------------------------------------------

// crawling represents a crawling Crawler ... busy crawling ...
type crawling struct {
	Crawler             // config
	traffic             // to be crawled
	results chan result // to be reported
}

// ---------------------------------------------------------------------------
// teach `*crawling` some straight-forward behaviour, and how to crawl :-)

// goWaitAndClose is to be used after initial traffic has been added.
func (c *crawling) goWaitAndClose() {
	go func(c *crawling) {
		<-sites.SiteDoneWait(c.Travel, c)
		close(c.results)
	}(c)
}

// ===========================================================================

// crawl performs a crawling Crawler's main function: crawl.
//
// This version attempts to respect the original implementation.
// Thus, it is still too busy - catering with too many concerns;
// concerns which might better be catered for in the traffic crawling processor's pipe network;
// passed to traffic.processor, and used by `DoneSiteFunc`.
func (c *crawling) crawl(s site) {
	urls := c.crawlSite(s)         // the core crawl process
	c.Feed(urls, s.URL, s.Depth-1) // new urls enter crawling - circular feedback
	c.Done()                       // site leaves crawling
	time.Sleep(c.Delay)            // have a gentle nap
}

// ===========================================================================

// a Crawler crawling: crawl traffic, emit results
// and signal when done
func (c Crawler) crawling(urls []*url.URL) (done <-chan struct{}) {
	crawling := crawling{
		c, // "Crawler is used as configuration ..."
		traffic{
			Travel:    make(chan site),     // the feedback traffic
			WaitGroup: new(sync.WaitGroup), // monitor traffic
		},
		make(chan result), // results - the (secondary) output
	}
	crawling.crawling(urls, parallel(c.Parallel))
	return res.ResultDone(crawling.results, c.report)
}

// ===========================================================================

// crawling builds and feeds the network,
// and returns after having launched the closer.
func (c *crawling) crawling(urls []*url.URL, parallel int) {
	c.Processor(c.crawl, parallel) // build the site traffic processing network
	c.Feed(urls, nil, c.Depth)     // feed initial urls
	c.goWaitAndClose()             // launch the closer
}

// ===========================================================================
