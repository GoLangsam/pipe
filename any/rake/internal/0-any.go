// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate bundle -pkg rake -o .\..\rake.go .

package rake

import (
	"github.com/cheekybits/genny/generic"
)

// Any is the generic type flowing thru the pipe network.
type Any generic.Type

type item = Any
