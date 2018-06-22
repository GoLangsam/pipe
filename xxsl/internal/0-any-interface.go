// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

// ===========================================================================
// Beg of anyThingChannel interface

// anyThingChannel represents a
// bidirectional
// channel of anyThing elements
type anyThingChannel interface {
	AnyChanCore // close, len & cap
	receiverAny // Receive / Request
	providerAny // Provide
}

// Note: Embedding AnyReceiver and AnyProvider directly would result in error: duplicate method Len Cap Close

// AnyReceiver represents a
// receive-only
// channel of anyThing elements
// - aka `<-chan`
type AnyReceiver interface {
	AnyChanCore // close, len & cap
	receiverAny // Receive / Request
}

type receiverAny interface {
	Receive() (data anyThing)              // the receive operator as method - aka `MyAny := <-myreceiverAny`
	Request() (data anyThing, isOpen bool) // the multi-valued comma-ok receive - aka `MyAny, ok := <-myreceiverAny`
}

// AnyProvider represents a
// send-enabled
// channel of anyThing elements
// - aka `chan<-`
type AnyProvider interface {
	AnyChanCore // close, len & cap
	providerAny // Provide
}

type providerAny interface {
	Provide(data anyThing) // the send method - aka `MyAnyproviderAny <- MyAny`
}

// AnyChanCore represents basic methods common to every
// channel of Any elements
type AnyChanCore interface {
	Close()
	Len() int
	Cap() int
}

// End of AnyChannel interface
// ===========================================================================
