// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Oh my god!		-in github.com/GoLangsam/pipe/s/pipe.go does not work :-(
// have to keep using	-in ../../../../s/pipe.go

//go:generate genny     -in ../../../../s/pipe.go		-out gen-pipe-result.go	-pkg result gen "anyThing=Result"

package result

// This file uses geanny to pull the type specific generic code
