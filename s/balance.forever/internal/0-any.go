// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate bundle -pkg balance -o .\..\balance.go .

package balance

import (
	"github.com/cheekybits/genny/generic"
)

type Any generic.Type
