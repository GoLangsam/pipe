// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Oh my god!		-in github.com/GoLangsam/pipe/s/pipe.go does not work :-(
// have to keep using	-in ../../../s/pipe.go

//go:generate genny     -in ../../../s/pipe.go		-out gen-pipe.go	-pkg httpsyet gen "Any=site"
//go:generate genny     -in ../../../s/pipe.go		-out gen-pipe-result.go	-pkg httpsyet gen "Any=result"
//go:generate genny	-in ../../../s/flap/flap.go	-out gen-flap.go	-pkg httpsyet gen "Any=site"
//go:generate genny	-in ../../../s/strew/strew.go	-out gen-strew.go	-pkg httpsyet gen "Any=site"
//go:generate genny	-in ../../../s/seen/seen.go	-out gen-seen.go	-pkg httpsyet gen "Any=site"

package httpsyet

// This file uses geanny to pull the type specific generic code
