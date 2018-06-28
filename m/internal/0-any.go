// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate bundle -pkg pipe -o .\..\pipe.go .

package pipe

import (
	"github.com/cheekybits/genny/generic"
)

// anyThing is the generic type flowing thru the pipe network.
type anyThing generic.Type

// anyThingFrom is a receive-only anyThing channel
type anyThingFrom <-chan anyThing

// anyThingInto is a send-only anyThing channel
type anyThingInto chan<- anyThing
