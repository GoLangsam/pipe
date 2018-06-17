# `pipe/s`
A pipers bag - generic functions to gain concurrency - batteries included :-)

[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](LICENSE.md)
[![Go Report Card](https://goreportcard.com/badge/github.com/GoLangsam/pipe)](https://goreportcard.com/report/github.com/GoLangsam/pipe)
[![Build Status](https://travis-ci.org/GoLangsam/pipe.svg?branch=master)](https://travis-ci.org/GoLangsam/pipe)
[![GoDoc](https://godoc.org/github.com/GoLangsam/pipe?status.svg)](https://godoc.org/github.com/GoLangsam/pipe)


```
    go get -u github.com/GoLangsam/pipe
```

Please feel free and encouraged to suggest, improve, comment or ask - You'll be welcome!

## Overview

- an evolution:
	- [sss](sss.naive/) a naive approach - as seen in popular slides, talks and blogs.
	- [ss](ss.notyet/) a better way to code it

- the essence
	- [s](s/) is for You to use - [Batteries](#batteries) included!

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

---
Think deep - code happy - be simple - see clear :-)
## Support on Beerpay
Hey dude! Help me out for a couple of :beers:!

[![Beerpay](https://beerpay.io/GoLangsam/pipe/badge.svg?style=beer-square)](https://beerpay.io/GoLangsam/pipe)  [![Beerpay](https://beerpay.io/GoLangsam/pipe/make-wish.svg?style=flat-square)](https://beerpay.io/GoLangsam/pipe?focus=wish)