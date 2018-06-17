# `pipe/s`
A pipers bag - generic functions to gain concurrency - batteries included :-)

```
    go get -u github.com/GoLangsam/pipe
```

- an evolution:
	- [sss](sss.naive/) a naive approach - as seen in popular slides, talks and blogs.
	- [ss](ss.notyet/) a better way to code it

- the essence
	- [s](s/) is for You to use - Batteries included!

- toys and tests
	- [expamples](expamples/)

- notes and explanations
	- [readme](readme/) contains further background documentation 

- internals
	- [internal](internal/) hides [bundledotgo](internal/cmd/bundledotgo), a quick&dirty patch of bundle
	- [.generate](.generate/) prepares the extended versions
	- [.generic](.generic/) prepares the all-together pipe.go

- extended
	- [xxl](xxl/) demand-driven channel - **lazy** evaluation
	- [xxs](xxs/) supply-driven counter-part
	- [xxsl](xsl/) the super luxury version has both: demand- and supply-driven channels

---

## [sss](sss.naive/) - simply super stupid small


## [ss](ss.notyet/) - still simply small


## [s](s/) - smart & useful - batteries included


### Batteries

	- [pipedone](s/pipedone/)
	- [plug](s/plug/)
	- [plugafter](s/plugafter/)
	- [flap](s/flap/)

	- [buffer](s/buffer/)

	- [scatter](s/scatter/)
	- [fan-out](s/fan-out/)

	- [seen](s/seen/)

	- [fan2](s/fan2/)
	- [fan-in](s/fan-in/)

	- [merge](s/merge/)
	- [same](s/same/)

	- [join](s/join/)

	- [daisy](s/daisy/)


## [examples](examples/)

	- [a web-crawler](examples/httpsyet/) - a refactored real world example

## [internal](internal/)

	- [bundle.go](internal/cmd/bundledotgo/)

