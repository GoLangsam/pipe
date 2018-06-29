// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file uses geanny to pull the type specific generic code

//go:generate genny	-in ../../m/pipe.go			-out 1-pipe.go		-pkg pipe gen "anyThing=Thing"
//go:generate genny	-in ../../m/.fan2/fan2.go		-out 6-fan2.go		-pkg pipe gen "anyThing=Thing"
//go:generate genny	-in ../../m/adjust/adjust.go		-out 3-adjust.go	-pkg pipe gen "anyThing=Thing"
//go:generate genny	-in ../../m/buffered/buffered.go	-out 2-buffered.go	-pkg pipe gen "anyThing=Thing"
//go:generate genny	-in ../../m/daisy/daisy.go		-out 9-daisy.go		-pkg pipe gen "anyThing=Thing"
//go:generate genny	-in ../../m/fan-in/fan-in.go		-out 6-fan-in.go	-pkg pipe gen "anyThing=Thing"
//go:generate genny	-in ../../m/fan-in1/fan-in1.go		-out 6-fan-in1.go	-pkg pipe gen "anyThing=Thing"
//go:generate genny	-in ../../m/fan-out/fan-out.go		-out 4-fan-out.go	-pkg pipe gen "anyThing=Thing"
//go:generate genny	-in ../../m/flap/flap.go		-out 2-flap.go		-pkg pipe gen "anyThing=Thing"
//go:generate genny	-in ../../m/freq/freq.go		-out 2-freq.go		-pkg pipe gen "anyThing=Thing"
//  :generate genny	-in ../../m/is-nil/is-nil.go		-out x-is-nil.go	-pkg pipe gen "anyThing=Thing"
//go:generate genny	-in ../../m/join/join.go		-out 8-join.go		-pkg pipe gen "anyThing=Thing"
//go:generate genny	-in ../../m/merge/merge.go		-out 7-merge.go		-pkg pipe gen "anyThing=Thing"
//go:generate genny	-in ../../m/pipedone/pipedone.go	-out 2-pipedone.go	-pkg pipe gen "anyThing=Thing"
//go:generate genny	-in ../../m/plug/plug.go		-out 2-plug.go		-pkg pipe gen "anyThing=Thing"
//go:generate genny	-in ../../m/plugafter/plugafter.go	-out 2-plugafter.go	-pkg pipe gen "anyThing=Thing"
//go:generate genny	-in ../../m/same/same.go		-out 7-same.go		-pkg pipe gen "anyThing=Thing"
//go:generate genny	-in ../../m/seen/seen.go		-out 5-seen.go		-pkg pipe gen "anyThing=Thing"
//go:generate genny	-in ../../m/strew/strew.go		-out 4-strew.go		-pkg pipe gen "anyThing=Thing"

package pipe
