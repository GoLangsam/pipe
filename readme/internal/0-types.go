// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

// ===========================================================================
// Directional channel

// anyThingFrom is a receive-only anyThing channel
type anyThingFrom <-chan anyThing

// anyThingInto is a send-only anyThing channel
type anyThingInto chan<- anyThing

// ===========================================================================
// Signalling

var done chan struct{} // returned from inside to outside

var quit chan struct{} // passed from outside to inside
var stop chan struct{} // passed from outside to deep inside (kill/abort)

// anyThingWait is where a process broadcasts it's done (and then quits)
type anyThingWait <-chan struct{}

// anyThingStop is where an outer process sends it's request to stop
type anyThingStop chan<- struct{}

// ===========================================================================
// Function signatures

// Oper represents an operation
// as in anyThingDoneFunc
type anyThingOper func(a anyThing)

// Func represents a function
// usually called `act` - action
// as in anyThingPipeFunc
type anyThingFunc func(a anyThing) anyThing

// Attr returns an attribute for discrimination
type anyThingAttr func(a anyThing) interface{}

// ===========================================================================
// Process signatures

// into <- <-from
type anyThingProc func(into anyThingInto, from anyThingFrom)

// ===========================================================================
// Pipe

type anyThingChan func(args ...interface{}) anyThingFrom
type anyThingPipe func(inp anyThingFrom, args ...interface{}) anyThingFrom
type anyThingDone func(inp anyThingFrom, args ...interface{}) anyThingDone

// ===========================================================================
// Closures

type anyThingTube func(inp anyThingFrom) anyThingFrom
type anyThingFini func(inp anyThingFrom) anyThingDone
