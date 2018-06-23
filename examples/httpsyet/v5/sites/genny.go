// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Oh my god!		-in github.com/GoLangsam/pipe/s/pipe.go does not work :-(
// have to keep using	-in ../../../../m/pipe.go

//go:generate genny     -in ../../../../m/pipe.go		-out gen-pipe.go	-pkg sites gen "anyThing=Site anyOwner=*Traffic"
//go:generate genny	-in ../../../../m/adjust/adjust.go	-out gen-adjust.go	-pkg sites gen "anyThing=Site anyOwner=*Traffic"
//go:generate genny	-in ../../../../m/flap/flap.go		-out gen-flap.go	-pkg sites gen "anyThing=Site anyOwner=*Traffic"
//go:generate genny	-in ../../../../m/strew/strew.go	-out gen-strew.go	-pkg sites gen "anyThing=Site anyOwner=*Traffic"
//go:generate genny	-in ../../../../m/seen/seen.go		-out gen-seen.go	-pkg sites gen "anyThing=Site anyOwner=*Traffic"

package sites

// This file uses geanny to pull the type specific generic code
