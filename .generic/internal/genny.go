// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file uses geanny to pull the type specific generic code

//go:generate genny	-in ../../s/pipe.go			-out 1-pipe.go		-pkg pipe gen "anyThing=Thing"
//go:generate genny	-in ../../s/.fan2/fan2.go		-out 6-fan2.go		-pkg pipe gen "anyThing=Thing"
//go:generate genny	-in ../../s/adjust/adjust.go		-out 3-adjust.go	-pkg pipe gen "anyThing=Thing"
//go:generate genny	-in ../../s/buffered/buffered.go	-out 2-buffered.go	-pkg pipe gen "anyThing=Thing"
//go:generate genny	-in ../../s/daisy/daisy.go		-out 9-daisy.go		-pkg pipe gen "anyThing=Thing"
//go:generate genny	-in ../../s/fan-in/fan-in.go		-out 6-fan-in.go	-pkg pipe gen "anyThing=Thing"
//go:generate genny	-in ../../s/fan-in1/fan-in1.go		-out 6-fan-in1.go	-pkg pipe gen "anyThing=Thing"
//go:generate genny	-in ../../s/fan-out/fan-out.go		-out 4-fan-out.go	-pkg pipe gen "anyThing=Thing"
//go:generate genny	-in ../../s/flap/flap.go		-out 2-flap.go		-pkg pipe gen "anyThing=Thing"
//  :generate genny	-in ../../s/is-nil/is-nil.go		-out x-is-nil.go	-pkg pipe gen "anyThing=Thing"
//go:generate genny	-in ../../s/join/join.go		-out 8-join.go		-pkg pipe gen "anyThing=Thing"
//go:generate genny	-in ../../s/merge/merge.go		-out 7-merge.go		-pkg pipe gen "anyThing=Thing"
//go:generate genny	-in ../../s/pipedone/pipedone.go	-out 2-pipedone.go	-pkg pipe gen "anyThing=Thing"
//go:generate genny	-in ../../s/plug/plug.go		-out 2-plug.go		-pkg pipe gen "anyThing=Thing"
//go:generate genny	-in ../../s/plugafter/plugafter.go	-out 2-plugafter.go	-pkg pipe gen "anyThing=Thing"
//go:generate genny	-in ../../s/same/same.go		-out 7-same.go		-pkg pipe gen "anyThing=Thing"
//go:generate genny	-in ../../s/seen/seen.go		-out 5-seen.go		-pkg pipe gen "anyThing=Thing"
//go:generate genny	-in ../../s/strew/strew.go		-out 4-strew.go		-pkg pipe gen "anyThing=Thing"

package pipe
