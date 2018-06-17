// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

// ===========================================================================
// Beg of MakeAny creators

// MakeAnyChannelChan returns a new open channel
// (simply a 'chan Any' that is).
//  Note: No 'Any-producer' is launched here yet! (as is in all the other functions).
//  This is useful to easily create corresponding variables such as:
/*
   var myAnyPipelineStartsHere := MakeAnyChan()
   // ... lot's of code to design and build Your favourite "myAnyWorkflowPipeline"
   // ...
   // ... *before* You start pouring data into it, e.g. simply via:
   for drop := range water {
       myAnyPipelineStartsHere <- drop
   }
   close(myAnyPipelineStartsHere)
*/
//  Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
//  (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
//  Note: as always (except for PipeAnyBuffer) the channel is unbuffered.
//
func MakeAnyChannelChan() (out AnyChannel) {
	return &AnySupply{make(chan Any)}
}

// MakeAnyChannelBuff returns a new open buffered channel with capacity `cap`.
func MakeAnyChannelBuff(cap int) (out AnyChannel) {
	return &AnySupply{make(chan Any, cap)}
}

// End of MakeAny creators
// ===========================================================================
