// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

import (
	"sync"
)

// just to be sure :-)
var _ anyThingWaiter = new(sync.WaitGroup)
