// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package httpsyet

import (
	"net/url"
	"sync"
)

// Traffic as it goes around inside a circular site pipe network,
// e. g. a crawling Crawler.
// Composed of Travel, a channel for those who travel in the traffic,
// and an embedded *sync.WaitGroup to keep track of congestion.
type Traffic struct {
	Travel          chan site // to be processed
	*sync.WaitGroup           // monitor SiteEnter & SiteLeave
}

// Feed registers new entries and launches their dispatcher
// (which we intentionally left untouched).
func (t *Traffic) Feed(urls []*url.URL, parent *url.URL, depth int) {
	t.Add(len(urls))
	go queueURLs(t.Travel, urls, parent, depth)
}

// Processor builds the site traffic processing network;
// it is cirular if crawl uses Feed to provide feedback.
func (t *Traffic) Processor(crawl func(s site), parallel int) {
	sites, seen := siteForkSeenAttr(t.Travel, site.Attr)
	for _, inp := range siteStrew(sites, parallel) {
		siteDone(inp, crawl) // `sites` leave inside crawl
	}
	siteDoneLeave(seen, t) // `seen` leave without further processing
}
