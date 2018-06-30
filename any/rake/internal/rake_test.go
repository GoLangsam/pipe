// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rake

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

// ===========================================================================

func ExampleRake_closure() {

	var r *Rake
	crawl := func(item Any) {
		if false {
			log.Println("have:", item)
		}
		for i := 0; i < rand.Intn(9)+2; i++ {
			r.Feed(rand.Intn(2000)) // up to 10 new numbers < 2.000
		}
		time.Sleep(time.Millisecond)
	}

	r = New(crawl, nil, 80)

	r.Feed(1)

	<-r.Done()

	fmt.Println("Done")
	// Output:
	// Done
}

// ===========================================================================

func ExampleRake_chained() {

	var r *Rake
	crawl := func(item Any) {
		if false {
			log.Println("have:", item)
		}
		for i := 0; i < rand.Intn(9)+2; i++ {
			r.Feed(rand.Intn(2000)) // up to 10 new numbers < 2.000
		}
		time.Sleep(time.Millisecond)
	}

	r = New(nil, nil, 80).Rake(crawl).Feed(1) // chain

	<-r.Done()

	fmt.Println("Done")
	// Output:
	// Done
}

// ===========================================================================

type crawling struct {
	// ... some config
	*Rake
	// ... more stuff - e.g. chan results
}

func (c *crawling) put(item Any) {
	if false {
		log.Println("have:", item)
	}
	for i := 0; i < rand.Intn(9)+2; i++ {
		c.Feed(rand.Intn(2000)) // note: c.Feed comes from anonymously embedded *Rake
	}
	time.Sleep(time.Millisecond)
}

func ExampleRake_closure_on_embedded() {

	var c *crawling // need to declare first so we may define the function with feed back

	crawl := func(item Any) { c.put(item) } // a closure

	c = &crawling{ // instantiate c
		New(crawl, nil, 80).Feed(1), // embed *Rake using closure `crawl`
	}

	<-c.Done()

	fmt.Println("Done")
	// Output:
	// Done
}

func ExampleRake_closure_as_arg_dangerous() {
	// Very dangerous!
	// works only because
	// it takes longer for `1` to progress in the network
	// than it takes for go to adjust the pointer `c`
	// and thus the closure will have it when invoked.
	var c *crawling                                               // need to declare `c` first so we may define the closure with feed back to `c`
	c = &crawling{New(func(a Any) { c.put(a) }, nil, 80).Feed(1)} // pass the closure directly as func
	<-c.Done()

	fmt.Println("Done")
	// Output:
	// Done
}

func ExampleRake_closure_as_arg_stepbystep() {
	var c *crawling                                       // need to declare `c` first so we may define the closure with feed back to `c`
	c = &crawling{New(func(a Any) { c.put(a) }, nil, 80)} // pass the closure directly as func
	c.Feed(1)
	<-c.Done()

	fmt.Println("Done")
	// Output:
	// Done
}

// ===========================================================================

type crawler struct {
	// ... some config
	rake *Rake
	// ... more stuff - e.g. chan results
}

func (c *crawler) Rake(r *Rake) *crawler {
	c.rake = r
	return c
}

func (c *crawler) put(item Any) {
	if false {
		log.Println("have:", item)
	}
	for i := 0; i < rand.Intn(9)+2; i++ {
		c.rake.Feed(rand.Intn(2000)) // note: c.Feed comes from anonymously embedded *Rake
	}
	time.Sleep(time.Millisecond)
}

func ExampleRake_assign_rake() {

	c := &crawler{}
	// c.Rake(New(c.put, nil, 80).Feed(1, 2, 3, 4, 5)) // this will bounce! nil!
	c.Rake(New(c.put, nil, 80))
	c.rake.Feed(1, 2, 3, 4, 5) // now everybody is in his place
	<-c.rake.Done()

	fmt.Println("Done")
	// Output:
	// Done
}

// ===========================================================================
