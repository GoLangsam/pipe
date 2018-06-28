// A concurrent prime sieve
// found in go/doc/play
// adapted Filter to return a tube function

package pipe

import (
	"fmt"
)

// Send the sequence 2, 3, 4, ... to channel 'ch'.
func Generate(ch anyThingInto) {
	for i := 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// Filter returns a tube function which does:
// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func Filter(prime anyThing) (tube func(anyThingInto, anyThingFrom)) {
	tube = func(out anyThingInto, inp anyThingFrom) {
		prime := prime.(int)
		for {
			n := <-inp // Receive value from 'in'.
			i := n.(int)
			if i%prime != 0 {
				out <- i // Send 'i' to 'out'.
			}
		}
	}
	return tube
}

// The prime sieve: Daisy-chain Filter processes.
func ExampleanyThingDaisyChain_sieve() {
	ch := make(chan anyThing) // Create a new channel.
	go Generate(ch)           // Launch Generate goroutine.
	for i := 0; i < 10; i++ {
		prime := <-ch
		fmt.Println(prime)
		ch = anyThingDaisyChain(ch, Filter(prime))
	}
	// Output:
	// 2
	// 3
	// 5
	// 7
	// 11
	// 13
	// 17
	// 19
	// 23
	// 29
}
