# Examples seen in the wild

### Multiplexing - Fan-in

```go
func fanIn(input1, input2 <-chan string) <-chan string {
    c := make(chan string)
    go func() { for { c <- <-input1 } }()
    go func() { for { c <- <-input2 } }()
    return c
}
```

[Go Concurrency Patterns #27](https://talks.golang.org/2012/concurrency.slide#27)

### Fan-in using select

```go
func fanIn(input1, input2 <-chan string) <-chan string {
    c := make(chan string)
    go func() {
        for {
            select {
            case s := <-input1:  c <- s
            case s := <-input2:  c <- s
            }
        }
    }()
    return c
}

```

[Go Concurrency Patterns #34](https://talks.golang.org/2012/concurrency.slide#34)

### Daisy-chain

```go
func f(left, right chan int) {
    left <- 1 + <-right
}

func main() {
    const n = 10000
    leftmost := make(chan int)
    right := leftmost
    left := leftmost
    for i := 0; i < n; i++ {
        right = make(chan int)
        go f(left, right)
        left = right
    }
    go func(c chan int) { c <- 1 }(right)
    fmt.Println(<-leftmost)
}
```

[Go Concurrency Patterns #39](https://talks.golang.org/2012/concurrency.slide#39)


---
[Back to overview](overview.md)
