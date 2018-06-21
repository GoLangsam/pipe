// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package httpsyet

import (
	"fmt"
	"net/url"
)

// site represents what travels: an URL
// which may have a Parent URL, and a Depth.
type site struct {
	URL    *url.URL
	Parent *url.URL
	Depth  int
}

// ---------------------------------------------------------------------------
// site learning a little behaviour

// Attr implements the attribute relevant for ForkSiteSeenAttr,
// the "I've seen this site before" discriminator.
func (s site) Attr() interface{} {
	return s.URL.String()
}

// Print may be used via e.g. PipeSiteFunc(sites, site.print) for tracing
func (s site) Print() site {
	fmt.Println(s)
	return s
}

// queueURLs sends urls on the given queue
func queueURLs(queue chan<- Site, urls []*url.URL, parent *url.URL, depth int) {
	for _, u := range urls {
		queue <- Site{
			URL:    u,
			Parent: parent,
			Depth:  depth,
		}
	}
}
