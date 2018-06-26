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

func Example_Rake() {

	var r *Rake
	crawl := func(item Any) {
		log.Println("have:", item)
		for i := 0; i < rand.Intn(9) + 2 ; i++ {
			r.Feed(rand.Intn(2000)) // up to 10 new numbers < 2.000
		}
		time.Sleep(time.Millisecond * 10)
	}

	r = New(crawl, nil, 80)

	r.Feed(1)

	<-r.Done()

	fmt.Println("Done")
	// Output:
	// Done
}
