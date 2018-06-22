// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sites

import (
	"net/url"
	"sync"
)

// Traffic as it goes around inside a circular site pipe network,
// e. g. a crawling Crawler.
// Composed of Travel, a channel for those who travel in the traffic,
// and an embedded *sync.WaitGroup to keep track of congestion.
type Traffic struct {
	Travel          chan Site // to be processed
	*sync.WaitGroup           // monitor SiteEnter & SiteLeave
}

// Feed registers new entries and launches their dispatcher
// (which we intentionally left untouched).
func (t *Traffic) Feed(urls []*url.URL, parent *url.URL, depth int) {
	queueURLs(t.Travel, urls, parent, depth)
}

// Processor builds the site traffic processing network;
// it is cirular if crawl uses Feed to provide feedback.
func (t *Traffic) Processor(crawl func(s Site), parallel int) {
	sites, seen := SiteForkSeenAttr(SitePipeEnter(t.Travel, t), Site.Attr)
	for _, inp := range SiteStrew(SitePipeAdjust(sites), parallel) {
		SiteDoneFunc(inp, crawl) // `sites` leave inside crawl
	}
	SiteDoneLeave(seen, t) // `seen` leave without further processing
}
