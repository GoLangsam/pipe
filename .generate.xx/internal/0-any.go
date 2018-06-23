// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate bundle -pkg pipe -o ..\pipe.go.gen .

////go:generate genny -in $GOFILE	-out ../s/internal/$GOFILE.supply	gen "anymode=*anySupply"
////go:generate genny -in $GOFILE	-out ../l/internal/$GOFILE.demand	gen "anymode=*anyDemand"

package pipe

import (
	"github.com/cheekybits/genny/generic"
)

// anyThing is the generic type flowing thru the pipe network.
type anyThing generic.Type

// anymode is the generic channel type connecting the pipe network components.
type anymode generic.Type
