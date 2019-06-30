// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

// ===========================================================================
// Beg of anyThingMake creators

// anyThingMakeChan returns a new open channel
// (simply a 'chan anyThing' that is).
//
// Note: No 'anyThing-producer' is launched here yet! (as is in all the other functions).
//  This is useful to easily create corresponding variables such as:
//
// 	var myanyThingPipelineStartsHere := anyThingMakeChan()
// 	// ... lot's of code to design and build Your favourite "myanyThingWorkflowPipeline"
// 	// ...
// 	// ... *before* You start pouring data into it, e.g. simply via:
// 	for drop := range water {
// 	    myanyThingPipelineStartsHere <- drop
// 	}
// 	close(myanyThingPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for anyThingPipeBuffer) the channel is unbuffered.
//
func anyThingMakeChan() chan anyThing {
	return make(chan anyThing)
}

// End of anyThingMake creators
// ===========================================================================
