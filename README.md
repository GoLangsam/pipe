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


## [s](s/) - smart & useful - [batteries](#batteries) included!


### Batteries

- [pipedone](s/pipedone/)	- be signalled when flow subsides here
- [plug](s/plug/)	- pull the plug
- [plugafter](s/plugafter/) - pull the plug after some time
- [flap](s/flap/)	- keep track how many enter and leave
- [buffered](s/buffered/)	- insert a buffered channel

- [adjust](s/adjust/)	- insert an adjustable buffer 

- [strew](s/strew/)	- throw it at some available receiver
- [fan-out](s/fan-out/)	- throw it at all receivers

- [seen](s/seen/)	- an "I've seen this before" filter / forker

- [fan2](s/fan2/)	- feed some inputs into your channel
- [fan-in](s/fan-in/)	- gather all inputs into one output

- [merge](s/merge/)	- gather sorted streams into one
- [same](s/same/)	- compare two streams

- [join](s/join/)	- join stuff into the flow

- [daisy](s/daisy/)	- daisy chain tubes


## [examples](examples/)

- [a web-crawler](examples/httpsyet/) - a refactored real world example

## [internal](internal/)

- [bundle.go](internal/cmd/bundledotgo/) - a quick&dirty hack

---
Think deep - code happy - be simple - see clear :-)
## Support on Beerpay
Hey dude! Help me out for a couple of :beers:!

[![Beerpay](https://beerpay.io/GoLangsam/pipe/badge.svg?style=beer-square)](https://beerpay.io/GoLangsam/pipe)  [![Beerpay](https://beerpay.io/GoLangsam/pipe/make-wish.svg?style=flat-square)](https://beerpay.io/GoLangsam/pipe?focus=wish)