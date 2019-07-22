# Directories - Where is what? - What is where?

- an evolution:
	- [`ssss`](../ssss.none/) a first 'theoretical' approach - does not get You anywhere
	- [`sss`](../sss.naive/) a naive approach - as seen in popular slides, talks and blogs
	- [`ss`](../ss.notyet/) a better way to do it

- the essence
	- [`s`](../s/) package-level functions for You to use - [Batteries](batteries.md) included!
	- [`m`](../m/) You prefer to use methods? - Same [batteries](batteries.md) included!
	- [`l`](../l/) *lazy* evolution - see also `xxl` below

- toys and tests
	- [`examples`](../examples/)

- notes and explanations
	- [`readme`](../readme/) - various information: essays, references ... You're looking at it right now.

- internals
	- [`internal`](../internal/) hides [bundledotgo](../internal/cmd/bundledotgo), a quick&dirty patch of bundle
	- [`.generate.xx`](../.generate.xx/) prepares the extended xx?-versions
	- [`.generic`](../.generic/) prepares the all-together package-level function version [`pipe.go`](../pipe.go)
	- [`.generic.m`](../.generic.m/) provides an all-together method-based version [`pipe.go`](../.generic.m/pipe.go)

- extended
	- [`xxl`](../xxl/) demand-driven channel - **lazy** evaluation
	- [`xxs`](../xxs/) supply-driven counter-part - as a true drop-in equivalent
	- [`xxsl`](../xxsl/) the super luxury version has both: demand- and supply-driven channels in one place

---
[Back to overview](overview.md)

