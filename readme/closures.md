# Closures

Functions are first citizens in go.

A Function returning a function is a powerful idiom. 

Closures takes this on a higher level.

These concepts are orthogonal to concurrency.

Functions and closures are used a lot throughout this repo.

## What others say

- [Closures are the Generics for Go](https://medium.com/capital-one-developers/closures-are-the-generics-for-go-cb32021fb5b5) by [Jon Bodner](https://medium.com/@jon_43067)
  "We could use `interface{}` as a way to pass untyped input and output parameters around but that misses the point. It creates ugly code that requires casts and subverts the type system that helps us write correct code."

  Shows clearly, that 'generics' can be attacked and solved using
  - generated code - for generic algorithms & data structures
  - closures - as shown here
    - see below
    - `type sorter struct { ... }` - see below - aka `dancing` in dlx

- [use a closure](https://play.golang.org/p/dNmhg_6x9T)

```go
	package main

	import (
		"fmt"
	)

	func outer2(name string) int {
		var total int
		myClosure := func(x int) {
			total = len(name) * x
		}
		helper(myClosure)
		return total
	}

	func helper(f func(int)) {
		f(4)
	}

	func main() {
		fmt.Println(outer2("hello"))
	}
```

Result: 20 
- = 5 * 4
- = _`(len("hello"))`_ * 4 )
- = _`(len("hello") * x )`_(4)
- = _`(len(name)`("hello")` * x )`_(4)

"Look at `outer2`. Its local variable `total` was modified when `myClosure` was passed to `helper` and called from there.
There’s no reference to `total` in `helper`, but using a closure allowed it to be modified.
Just like structs in Go, closures have state.
This state provides the solution to our problem."

---
- [Closure: sorter](https://play.golang.org/p/hYMcQ81AvN)

```go
	package main

	import (
		"fmt"
		"sort"
	)

	type sorter struct {
		len  int
		swap func(i, j int)
		less func(i, j int) bool
	}
	
	func (x sorter) Len() int           { return x.len }
	func (x sorter) Swap(i, j int)      { x.swap(i, j) }
	func (x sorter) Less(i, j int) bool { return x.less(i, j) }
	
	func Sort(n int, swap func(i, j int), less func(i, j int) bool) {
		sort.Sort(sorter{len: n, swap: swap, less: less})
	}

	func main() {
		a := []int{5, 4, 3, 2, 1}
		Sort(
			len(a),
			func(i, j int) {
				temp := a[i]
				a[i] = a[j]
				a[j] = temp
			},
			func(i, j int) bool {
				return a[i] < a[j]
			})
		fmt.Println(a)

		b := []string{"bear", "cow", "ant", "chicken", "dog"}
		Sort(
			len(b),
			func(i, j int) {
				temp := b[i]
				b[i] = b[j]
				b[j] = temp
			},
			func(i, j int) bool {
				return b[i] < b[j]
			})
		fmt.Println(b)
	}

```

---
[Back to overview](overview.md)


