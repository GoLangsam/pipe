// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Oh my god!		-in github.com/GoLangsam/pipe/s/pipe.go does not work :-(
// have to keep using	-in ../../s/pipe.go

//go:generate genny     -in ../../s/pipe.go		-out gen-pipe.go	-pkg httpsyet gen "Any=site"
//go:generate genny     -in ../../s/pipe.go		-out gen-pipe-str.go	-pkg httpsyet gen "Any=string"
//go:generate genny	-in ../../s/scatter/scatter.go	-out gen-scatter.go	-pkg httpsyet gen "Any=site"
//go:generate genny	-in ../../s/seen/seen.go	-out gen-seen.go	-pkg httpsyet gen "Any=site"

////go:generate genny     -in ../../s/buffer/buffer.go	-out gen-buffer.go	-pkg httpsyet gen "Any=site"
////go:generate genny	-in ../../s/daisy/daisy.go	-out daisy.go		-pkg httpsyet gen "Any=site"
////go:generate genny	-in ../../s/fan-in/fan-in.go	-out gen-fan-in.go	-pkg httpsyet gen "Any=site"
////go:generate genny	-in ../../s/fan-out/fan-out.go	-out fan-out.go		-pkg httpsyet gen "Any=site"
////go:generate genny	-in ../../s/flap/flap.go	-out gen-flap.go	-pkg httpsyet gen "Any=site"
////go:generate genny	-in ../../s/join/join.go	-out gen-join.go	-pkg httpsyet gen "Any=site"
////go:generate genny	-in ../../s/plug/plug.go	-out 2-plug.go		-pkg httpsyet gen "Any=site"
////go:generate genny	-in ../../s/plugafter/plugafter.go	-out 2-plugafter.go	-pkg httpsyet gen "Any=site"
////go:generate genny	-in ../../s/same/same.go	-out 7-same.go		-pkg httpsyet gen "Any=site"

package httpsyet

// This file uses geanny to pull the type specific generic code
