// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate bundle -pkg pipe -o ..\pipe.go.gen .

////go:generate genny -in $GOFILE	-out ../s/internal/$GOFILE.supply	gen "Anymode=*AnySupply"
////go:generate genny -in $GOFILE	-out ../l/internal/$GOFILE.demand	gen "Anymode=*AnyDemand"

package pipe

import (
	"github.com/cheekybits/genny/generic"
)

type Any interface{}

type Anymode generic.Type
