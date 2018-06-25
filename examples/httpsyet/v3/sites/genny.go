// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file uses geanny to pull the type specific generic code

// Oh my god!		-in github.com/GoLangsam/pipe/s/pipe.go does not work :-(
// have to keep using	-in ../../../s/pipe.go

//go:generate genny     -in ../../../../s/pipe.go		-out gen-pipe.go	-pkg sites gen "anyThing=Site"
//go:generate genny	-in ../../../../s/adjust/adjust.go	-out gen-adjust.go	-pkg sites gen "anyThing=Site"
//go:generate genny	-in ../../../../s/flap/flap.go		-out gen-flap.go	-pkg sites gen "anyThing=Site"
//go:generate genny	-in ../../../../s/strew/strew.go	-out gen-strew.go	-pkg sites gen "anyThing=Site"
//go:generate genny	-in ../../../../s/seen/seen.go		-out gen-seen.go	-pkg sites gen "anyThing=Site"

package sites
