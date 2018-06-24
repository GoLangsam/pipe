// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// package pipe provides functions
// useful to build a network of concurrent pipe processes
// the components of which are connected by channels.
//
// Just - all these beautiful and useful generic definitions
// cannot be seen via godoc here,
// as they need to be private - initially.
//
// This is due to the 'funny' way "genny" handles
// identifier-casing:
//
//  If the original generic type has a public (uppercase) name,
//  then generated identifiers will remain uppercase (and thus public)
//  and you may never generate private non-public (lowercased) identifiers.
//
// Thus we need to start with a private id (such as `anyThing` here)
// and with private function names in order to respect Your freedom of choice.
//
// I am awfully sorry for this inconvenience
// and provide two alternatives:
//
//  In the root of the repo is a complete generated version of `package pipe`
//  with a public generic type (`Any`). Thus, everything public is visible
//  (just: it's not intended for further use - it's way too large, isn't it?).
//
//  Under the `examples` folder there are directories with samples generated
//  for public types - and thus provide meaningful `godoc` documentation
//  for what is used in the particular context at hand.
//
// Please enjoy to study and use what You find here.
//
// And please feel free and encouraged to suggest, improve, comment or ask,
// You'll be welcome!
//
// Think deep - code happy - be simple - see clear :-)
package pipe
