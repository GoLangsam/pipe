// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file uses geanny to pull the type specific generic code

// Oh my god!		-in github.com/GoLangsam/pipe/s/pipe.go does not work :-(
// have to keep using	-in ../../../m/pipe.go

//go:generate genny     -in ../../../m/pipe.go		-out result-pipe.go	-pkg httpsyet gen "anyThing=result"
//go:generate genny     -in ../../../any/rake/rake.go	-out rake.go		-pkg httpsyet gen "Any=site"

package httpsyet
