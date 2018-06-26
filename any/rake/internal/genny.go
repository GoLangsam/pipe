// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file uses geanny to pull the type specific generic code

// Oh my god!		-in github.com/GoLangsam/pipe/m/pipe.go does not work :-(
// have to keep using	-in ../../../m/pipe.go

//go:generate genny     -in ../../../m/pipe.go		-out 91-pipe.go		-pkg rake gen "anyThing=item anyOwner=*Rake"
//go:generate genny	-in ../../../m/adjust/adjust.go	-out 92-adjust.go	-pkg rake gen "anyThing=item anyOwner=*Rake"
//go:generate genny	-in ../../../m/flap/flap.go	-out 93-flap.go		-pkg rake gen "anyThing=item anyOwner=*Rake"
//go:generate genny	-in ../../../m/strew/strew.go	-out 94-strew.go	-pkg rake gen "anyThing=item anyOwner=*Rake"
//go:generate genny	-in ../../../m/seen/seen.go	-out 95-seen.go		-pkg rake gen "anyThing=item anyOwner=*Rake"

package rake
