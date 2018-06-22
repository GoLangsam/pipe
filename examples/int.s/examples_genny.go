// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate genny -in ../../s/pipe.go			-out pipe-int.go	-pkg pipe gen "anyThing=int"
//go:generate genny -in ../../s/.fan2/fan2.go		-out fan2-int.go	-pkg pipe gen "anyThing=int"
//go:generate genny -in ../../s/adjust/adjust.go	-out adjust-int.go	-pkg pipe gen "anyThing=int"
//go:generate genny -in ../../s/buffered/buffered.go	-out buffered-int.go	-pkg pipe gen "anyThing=int"
//go:generate genny -in ../../s/daisy/daisy.go		-out daisy-int.go	-pkg pipe gen "anyThing=int"
//go:generate genny -in ../../s/fan-in/fan-in.go	-out fan-in-int.go	-pkg pipe gen "anyThing=int"
//go:generate genny -in ../../s/fan-out/fan-out.go	-out fan-out-int.go	-pkg pipe gen "anyThing=int"
//go:generate genny -in ../../s/flap/flap.go		-out flap-int.go	-pkg pipe gen "anyThing=int"
//  :generate genny -in ../../s/is-nil/is-nil.go	-out is-nil-int.go	-pkg pipe gen "anyThing=int"
//go:generate genny -in ../../s/join/join.go		-out join-int.go	-pkg pipe gen "anyThing=int"
//go:generate genny -in ../../s/merge/merge.go		-out merge-int.go	-pkg pipe gen "anyThing=int"
//go:generate genny -in ../../s/pipedone/pipedone.go	-out pipedone-int.go	-pkg pipe gen "anyThing=int"
//go:generate genny -in ../../s/plug/plug.go		-out plug-int.go	-pkg pipe gen "anyThing=int"
//go:generate genny -in ../../s/plugafter/plugafter.go	-out plugafter-int.go	-pkg pipe gen "anyThing=int"
//go:generate genny -in ../../s/same/same.go		-out same-int.go	-pkg pipe gen "anyThing=int"
//go:generate genny -in ../../s/strew/strew.go		-out strew-int.go	-pkg pipe gen "anyThing=int"
//go:generate genny -in ../../s/seen/seen.go		-out seen-int.go	-pkg pipe gen "anyThing=int"

package pipe

// This file uses geanny to pull the type specific generic code
