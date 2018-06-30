# Batteries - included

## Overview

- Special terminators
  - [freq](freq#)		- ???

- Special purpose pipe-tubes
  - [pipedone](#pipedone)	- be signalled when flow subsides here
  - [plug](#plug)		- pull the plug
  - [plugafter](#plugafter)	- pull the plug after some time
  - [flap](#flap)		- keep track how many enter and leave
  - [buffered](#buffered)	- insert a buffered channel

  - [adjust](#adjust)		- insert an adjusting buffer 

  - [seen](#seen)		- an "I've seen this before" filter / forker

- Parallel processing
  - [sema](#sema)		- ???

- One to Many
  - [strew](#strew)		- send each to one receiver available
  - [fan-out](#fan-out)		- send each to all receivers (lock-step)

- Many to One
  - [fan2](#fan2)		- feed some inputs into your channel
  - [fan-in](#fan-in)		- gather all inputs into one output
  - [fan-in1](#fan-in1)		- same - just using only one goroutine

  - [merge](#merge)		- gather sorted streams into one
  - [same](#same)		- compare two (or more) streams

- Circular Feedback
  - [join](#join)		- join stuff into the flow

- Daisy Chain
  - [daisy](#daisy)		- daisy chain tubes

---
## in alphabethical order

1.  [adjust](#adjust)		- insert an adjusting buffer
2.  [buffered](#buffered)	- insert a buffered channel
3.  [daisy](#daisy)		- daisy chain tubes
4.  [fan-in1](#fan-in1)		- same - just using only one goroutine
5.  [fan-in](#fan-in)		- gather all inputs into one output
6.  [fan-out](#fan-out)		- send each to all receivers (lock-step)
7.  [fan2](#fan2)		- feed some inputs into your channel
8.  [flap](#flap)		- keep track how many enter and leave
9.  [freq](freq#)		- ???
10.  [join](#join)		- join stuff into the flow
11.  [merge](#merge)		- gather sorted streams into one
12.  [pipedone](#pipedone)	- be signalled when flow subsides here
13.  [plug](#plug)		- pull the plug
14.  [plugafter](#plugafter)	- pull the plug after some time
15.  [same](#same)		- compare two (or more) streams
16.  [seen](#seen)		- an "I've seen this before" filter / forker
17.  [sema](#sema)		- ???
18.  [strew](#strew)		- send each to one receiver available

---
## Details

### `freq`
### `sema`
### `pipedone`
### `plug`
### `plugafter`
### `flap`
### `buffered`
### `adjust`
### `strew`
### `fan-out`
### `seen`
### `fan2`
### `fan-in`
### `fan-in1`
### `merge`
### `same`
### `join`
### `daisy`

---
[Back to overview](overview.md)
