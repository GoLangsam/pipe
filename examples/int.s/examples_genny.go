// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate genny -in ../../s/pipe.go			-out gen-pipe.go	-pkg pipe gen "anyThing=int"
//go:generate genny -in ../../s/.fan2/fan2.go		-out gen-fan2.go	-pkg pipe gen "anyThing=int"
//go:generate genny -in ../../s/adjust/adjust.go	-out gen-adjust.go	-pkg pipe gen "anyThing=int"
//go:generate genny -in ../../s/buffered/buffered.go	-out gen-buffered.go	-pkg pipe gen "anyThing=int"
//go:generate genny -in ../../s/daisy/daisy.go		-out gen-daisy.go	-pkg pipe gen "anyThing=int"
//go:generate genny -in ../../s/fan-in/fan-in.go	-out gen-fan-in.go	-pkg pipe gen "anyThing=int"
//go:generate genny -in ../../s/fan-out/fan-out.go	-out gen-fan-out.go	-pkg pipe gen "anyThing=int"
//go:generate genny -in ../../s/flap/flap.go		-out gen-flap.go	-pkg pipe gen "anyThing=int"
//  :generate genny -in ../../s/is-nil/is-nil.go	-out gen-is-nil.go	-pkg pipe gen "anyThing=int"
//go:generate genny -in ../../s/join/join.go		-out gen-join.go	-pkg pipe gen "anyThing=int"
//go:generate genny -in ../../s/merge/merge.go		-out gen-merge.go	-pkg pipe gen "anyThing=int"
//go:generate genny -in ../../s/pipedone/pipedone.go	-out gen-pipedone.go	-pkg pipe gen "anyThing=int"
//go:generate genny -in ../../s/plug/plug.go		-out gen-plug.go	-pkg pipe gen "anyThing=int"
//go:generate genny -in ../../s/plugafter/plugafter.go	-out gen-plugafter.go	-pkg pipe gen "anyThing=int"
//go:generate genny -in ../../s/same/same.go		-out gen-same.go	-pkg pipe gen "anyThing=int"
//go:generate genny -in ../../s/strew/strew.go		-out gen-strew.go	-pkg pipe gen "anyThing=int"
//go:generate genny -in ../../s/seen/seen.go		-out gen-seen.go	-pkg pipe gen "anyThing=int"

package pipe

// This file uses geanny to pull the type specific generic code
