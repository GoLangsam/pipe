// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package httpsyet

import (
	"fmt"
	"net/url"
	"time"
)

// ---------------------------------------------------------------------------

// site represents what travels: an URL
// which may have a Parent URL, and a Depth.
type site struct {
	URL    *url.URL
	Parent *url.URL
	Depth  int
}

// ---------------------------------------------------------------------------
// teach `Crawler` how to `report` a result.

// report prints a result to Crawler.Out;
// used by `DoneResultFunc`.
func (c *Crawler) report(r result) {
	if _, err := fmt.Fprintln(c.Out, r); err != nil {
		c.Log.Printf("failed to write output '%s': %v\n", r, err)
	}
}

// ---------------------------------------------------------------------------

// crawling represents a crawling Crawler ... busy crawling ...
type crawling struct {
	*Crawler             // config
	*Rake                // the cirular network
	results  chan result // to be reported
}

// ---------------------------------------------------------------------------
// teach `*crawling` how to feed urls

func (c *crawling) add(urls []*url.URL, parent *url.URL, depth int) {
	for _, u := range urls {
		c.Feed(site{
			URL:    u,
			Parent: parent,
			Depth:  depth,
		})
	}
}

// ===========================================================================

// a Crawler crawling: crawl traffic, emit results
// and signal when done
func (c *Crawler) crawling(urls []*url.URL) (done <-chan struct{}) {
	var w *crawling

	// how to perform a crawling Crawler's main function: crawl.
	rake := func(s site) {
		w.add(w.crawlSite(s), s.URL, s.Depth-1) // new urls enter crawling - circular feedback
		time.Sleep(c.Delay)                     // have a gentle nap
	}

	// how to discriminate 'seen before'
	attr := func(s site) interface{} {
		return s.URL.String()
	}

	// how many in parallel
	many := parallel(c.Parallel) // no idea what keeps Crawler from setting `Parallel` upon validation

	w = &crawling{
		c, // "Crawler is used as configuration ..."
		New(rake, attr, many), // the cirular network
		make(chan result),     // results - the (secondary) output
	}

	go func() { // launch the results closer
		<-w.Done() // block 'till crawling.Rake is done
		close(w.results)
	}()

	w.add(urls, nil, c.Depth) // feed initial urls

	return DoneFunc(w.results, c.report) // signal when results report are done
}

// ===========================================================================
