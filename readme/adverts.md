# adverts and advices

## Such says he - Rob Pike - a founding member of the go team

### Rob Pike's Go Course Day 3 # 17

- "Function returning channel" is an important idiom.

### Lesson

- A complex problem can be broken down into easy-to-understand components.
- The pieces can be composed concurrently.
- The result is easy to understand, efficient, scalable, and correct.
- Maybe even parallel.

[Concurrency is not Parallelism #55](https://talks.golang.org/2012/waza.slide#55)

### Conclusion

- Concurrency is powerful.
- Concurrency is not parallelism.
- Concurrency enables parallelism.
- Concurrency makes parallelism (and scaling and everything else) easy.

[Concurrency is not Parallelism #58](https://talks.golang.org/2012/waza.slide#58)

### Don't overdo it

- They're fun to play with, but don't overuse these ideas.
- Goroutines and channels are big ideas. They're tools for program construction.
- But sometimes all you need is a reference counter.
- Go has "sync" and "sync/atomic" packages that provide mutexes, condition variables, etc. They provide tools for smaller problems.
- Often, these things will work together to solve a bigger problem.

Always use the right tool for the job.

[Go Concurrency Patterns #54](https://talks.golang.org/2012/concurrency.slide#54)

---
[Back to overview](overview.md)
