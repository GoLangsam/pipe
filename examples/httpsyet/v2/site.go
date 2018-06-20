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

// Attr impelemnts the attribute relevant for ForksiteSeenAttr,
// the "I've seen this site before" discriminator.
func (s site) Attr() interface{} {
	return s.URL.String()
}

// print may be used via e.g. PipesiteFunc(sites, site.print) for tracing
func (s site) Print() site {
	fmt.Println(s)
	return s
}
