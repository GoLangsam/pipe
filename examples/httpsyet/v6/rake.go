// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by golang.org/x/tools/cmd/bundle. DO NOT EDIT.

package httpsyet

import (
	"container/ring"
	"sync"
	"time"
)

type item = site

// Rake represents a fanned out circular pipe network
// with a flexibly adjusting buffer.
// site item is processed once only -
// items seen before are filtered out.
//
// A Rake may be used e.g. as a crawling Crawler
// where every link shall be visited only once.
type Rake struct {
	items chan item                // to be processed
	wg    *sync.WaitGroup          // monitor SiteEnter & SiteLeave
	done  chan struct{}            // to signal termination due to traffic having subsided
	once  *sync.Once               // to close Done only once - lauched from first feed
	rake  func(a item)             // function to be applied
	attr  func(a item) interface{} // attribute to discriminate seen
	runs  bool                     // am I running?
	many  int                      // # of parallel raking endpoints of the Rake
}

// New returns a (pointer to a) new operational Rake.
//
// `rake` is the operation to be executed in parallel on any item
// which has not been seen before.
// Have it use `myrake.Feed(items...)` in order to provide feed-back.
//
// `attr` allows to specify an attribute for the seen filter.
// Pass `nil` to filter on any item itself.
//
// `somany` is the # of parallel processes - the parallelism
// of the network built by Rake,
// the # of parallel raking endpoints of the Rake.
func New(
	rake func(a item),
	attr func(a item) interface{},
	somany int,
) (
	my *Rake,
) {
	if somany < 1 {
		somany = 1
	}
	my = &Rake{
		make(chan item),
		new(sync.WaitGroup),
		make(chan struct{}),
		new(sync.Once),
		rake,
		attr,
		false,
		somany,
	}

	return my
}

// init builds the network
func (my *Rake) init() *Rake {
	proc := func(a item) { // wrap rake:
		my.rake(a)   // apply original rake
		my.wg.Done() // have this item leave
	}

	// build the concurrent pipe network
	items, seen := (itemFrom)(my.items).itemForkSeenAttr(my.attr)
	_ = seen.itemDoneLeave(my.wg) // `seen` leave without further processing

	for _, items := range items.itemPipeAdjust().itemStrew(my.many) {
		_ = items.itemDoneFunc(proc) // strewed `items` leave in wrapped `crawl`
	}

	return my
}

// start builds the network and spawns the closer
func (my *Rake) start() {
	my = my.init()
	my.runs = true
	go my.closer()
}

func (my *Rake) closer() *Rake {
	my.done <- <-(itemInto)(my.items).itemDoneWait(my.wg)
	close(my.done)
	return my
}

// checkRuns for paranoids
func (my *Rake) checkRuns() *Rake {
	if my.runs {
		panic("Rake is running already")
	}
	return my
}

// Rake sets the rake function to be applied (in parallel).
//
// `rake` is the operation to be executed in parallel on any item
// which has not been seen before.
//
// You may provide `nil` here and call `Rake(..)` later to provide it.
// Or have it use `myrake.Feed(items...)` in order to provide feed-back.
//
// Rake panics iff called after first nonempty `Feed(...)`
func (my *Rake) Rake(rake func(a item)) *Rake {
	my.checkRuns()
	my.rake = rake
	return my
}

// Attr sets the (optional) attribute to discriminate 'seen'.
//
// `attr` allows to specify an attribute for the 'seen' filter.
// If not set 'seen' will discriminate any item by itself.
//
// Seen panics iff called after first nonempty `Feed(...)`
func (my *Rake) Attr(attr func(a item) interface{}) *Rake {
	my.checkRuns()
	my.attr = attr
	return my
}

// Done returns a channel which will be signalled and closed
// when traffic has subsided, nothing is left to be processed
// and consequently all goroutines have terminated.
func (my *Rake) Done() (done <-chan struct{}) {
	return my.done
}

// Feed registers new items on the network.
func (my *Rake) Feed(items ...item) *Rake {

	if len(items) == 0 {
		return my // nothing to do
	}

	my.wg.Add(len(items)) // items enter

	my.once.Do(my.start) // lazy init: build & start the network

	for _, i := range items {
		my.items <- i
	}

	return my
}

// End of Rake
// ===========================================================================

// itemFrom is a receive-only item channel
type itemFrom <-chan item

// itemInto is a send-only item channel
type itemInto chan<- item

// ===========================================================================
// Beg of itemMake creators

// itemMakeChan returns a new open channel
// (simply a 'chan item' that is).
// Note: No 'item-producer' is launched here yet! (as is in all the other functions).
//  This is useful to easily create corresponding variables such as:
/*
var myitemPipelineStartsHere := itemMakeChan()
// ... lot's of code to design and build Your favourite "myitemWorkflowPipeline"
   // ...
   // ... *before* You start pouring data into it, e.g. simply via:
   for drop := range water {
myitemPipelineStartsHere <- drop
   }
close(myitemPipelineStartsHere)
*/
//  Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
//  (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for itemPipeBuffer) the channel is unbuffered.
//
func itemMakeChan() (out chan item) {
	return make(chan item)
}

// End of itemMake creators
// ===========================================================================

// ===========================================================================
// Beg of itemChan producers

// itemChan returns a channel to receive
// all inputs
// before close.
func itemChan(inp ...item) (out itemFrom) {
	cha := make(chan item)
	go chanitem(cha, inp...)
	return cha
}

func chanitem(out itemInto, inp ...item) {
	defer close(out)
	for i := range inp {
		out <- inp[i]
	}
}

// itemChanSlice returns a channel to receive
// all inputs
// before close.
func itemChanSlice(inp ...[]item) (out itemFrom) {
	cha := make(chan item)
	go chanitemSlice(cha, inp...)
	return cha
}

func chanitemSlice(out itemInto, inp ...[]item) {
	defer close(out)
	for i := range inp {
		for j := range inp[i] {
			out <- inp[i][j]
		}
	}
}

// itemChanFuncNok returns a channel to receive
// all results of generator `gen`
// until `!ok`
// before close.
func itemChanFuncNok(gen func() (item, bool)) (out itemFrom) {
	cha := make(chan item)
	go chanitemFuncNok(cha, gen)
	return cha
}

func chanitemFuncNok(out itemInto, gen func() (item, bool)) {
	defer close(out)
	for {
		res, ok := gen() // generate
		if !ok {
			return
		}
		out <- res
	}
}

// itemChanFuncErr returns a channel to receive
// all results of generator `gen`
// until `err != nil`
// before close.
func itemChanFuncErr(gen func() (item, error)) (out itemFrom) {
	cha := make(chan item)
	go chanitemFuncErr(cha, gen)
	return cha
}

func chanitemFuncErr(out itemInto, gen func() (item, error)) {
	defer close(out)
	for {
		res, err := gen() // generate
		if err != nil {
			return
		}
		out <- res
	}
}

// End of itemChan producers
// ===========================================================================

// ===========================================================================
// Beg of itemPipe functions

// itemPipeFunc returns a channel to receive
// every result of action `act` applied to `inp`
// before close.
// Note: it 'could' be itemPipeMap for functional people,
// but 'map' has a very different meaning in go lang.
func (inp itemFrom) itemPipeFunc(act func(a item) item) (out itemFrom) {
	cha := make(chan item)
	if act == nil { // Make `nil` value useful
		act = func(a item) item { return a }
	}
	go inp.pipeitemFunc(cha, act)
	return cha
}

func (inp itemFrom) pipeitemFunc(out itemInto, act func(a item) item) {
	defer close(out)
	for i := range inp {
		out <- act(i) // apply action
	}
}

// End of itemPipe functions
// ===========================================================================

// ===========================================================================
// Beg of itemTube closures around itemPipe

// itemTubeFunc returns a closure around PipeItemFunc (_, act).
func itemTubeFunc(act func(a item) item) (tube func(inp itemFrom) (out itemFrom)) {

	return func(inp itemFrom) (out itemFrom) {
		return inp.itemPipeFunc(act)
	}
}

// End of itemTube closures around itemPipe
// ===========================================================================

// ===========================================================================
// Beg of itemDone terminators

// itemDone returns a channel to receive
// one signal
// upon close
// and after `inp` has been drained.
func (inp itemFrom) itemDone() (done <-chan struct{}) {
	sig := make(chan struct{})
	go inp.doneitem(sig)
	return sig
}

func (inp itemFrom) doneitem(done chan<- struct{}) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// itemDoneSlice returns a channel to receive
// a slice with every item received on `inp`
// upon close.
//
// Note: Unlike itemDone, itemDoneSlice sends the fully accumulated slice, not just an event, once upon close of inp.
func (inp itemFrom) itemDoneSlice() (done <-chan []item) {
	sig := make(chan []item)
	go inp.doneitemSlice(sig)
	return sig
}

func (inp itemFrom) doneitemSlice(done chan<- []item) {
	defer close(done)
	slice := []item{}
	for i := range inp {
		slice = append(slice, i)
	}
	done <- slice
}

// itemDoneFunc
// will apply `act` to every `inp` and
// returns a channel to receive
// one signal
// upon close.
func (inp itemFrom) itemDoneFunc(act func(a item)) (done <-chan struct{}) {
	sig := make(chan struct{})
	if act == nil {
		act = func(a item) { return }
	}
	go inp.doneitemFunc(sig, act)
	return sig
}

func (inp itemFrom) doneitemFunc(done chan<- struct{}, act func(a item)) {
	defer close(done)
	for i := range inp {
		act(i) // apply action
	}
	done <- struct{}{}
}

// End of itemDone terminators
// ===========================================================================

// ===========================================================================
// Beg of itemFini closures

// itemFini returns a closure around `itemDone()`.
func (inp itemFrom) itemFini() func(inp itemFrom) (done <-chan struct{}) {

	return func(inp itemFrom) (done <-chan struct{}) {
		return inp.itemDone()
	}
}

// itemFiniSlice returns a closure around `itemDoneSlice()`.
func (inp itemFrom) itemFiniSlice() func(inp itemFrom) (done <-chan []item) {

	return func(inp itemFrom) (done <-chan []item) {
		return inp.itemDoneSlice()
	}
}

// itemFiniFunc returns a closure around `itemDoneFunc(act)`.
func (inp itemFrom) itemFiniFunc(act func(a item)) func(inp itemFrom) (done <-chan struct{}) {

	return func(inp itemFrom) (done <-chan struct{}) {
		return inp.itemDoneFunc(act)
	}
}

// End of itemFini closures
// ===========================================================================

// ===========================================================================
// Beg of itemPair functions

// itemPair returns a pair of channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func (inp itemFrom) itemPair() (out1, out2 itemFrom) {
	cha1 := make(chan item)
	cha2 := make(chan item)
	go inp.pairitem(cha1, cha2)
	return cha1, cha2
}

/* not used - kept for reference only.
func (inp itemFrom) pairitem(out1, out2 itemInto, inp itemFrom) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
} */

func (inp itemFrom) pairitem(out1, out2 itemInto) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		select { // send first to whomever is ready to receive
		case out1 <- i:
			out2 <- i
		case out2 <- i:
			out1 <- i
		}
	}
}

// End of itemPair functions
// ===========================================================================

// ===========================================================================
// Beg of itemFork functions

// itemFork returns two channels
// either of which is to receive
// every result of inp
// before close.
func (inp itemFrom) itemFork() (out1, out2 itemFrom) {
	cha1 := make(chan item)
	cha2 := make(chan item)
	go inp.forkitem(cha1, cha2)
	return cha1, cha2
}

/* not used - kept for reference only.
func (inp itemFrom) forkitem(out1, out2 itemInto) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
} */

func (inp itemFrom) forkitem(out1, out2 itemInto) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		select { // send first to whomever is ready to receive
		case out1 <- i:
			out2 <- i
		case out2 <- i:
			out1 <- i
		}
	}
}

// End of itemFork functions
// ===========================================================================

// ===========================================================================
// Beg of itemFanIn2 simple binary Fan-In

// itemFanIn2 returns a channel to receive
// all from both `inp` and `inp2`
// before close.
func (inp itemFrom) itemFanIn2(inp2 itemFrom) (out itemFrom) {
	cha := make(chan item)
	go inp.fanIn2item(cha, inp2)
	return cha
}

/* not used - kept for reference only.
// (inp itemFrom) fanin2item as seen in Go Concurrency Patterns
func fanin2item(out itemInto, inp, inp2 itemFrom) {
	for {
		select {
		case e := <-inp:
			out <- e
		case e := <-inp2:
			out <- e
		}
	}
} */

func (inp itemFrom) fanIn2item(out itemInto, inp2 itemFrom) {
	defer close(out)

	var (
		closed bool // we found a chan closed
		ok     bool // did we read successfully?
		e      item // what we've read
	)

	for !closed {
		select {
		case e, ok = <-inp:
			if ok {
				out <- e
			} else {
				inp = inp2    // swap inp2 into inp
				closed = true // break out of the loop
			}
		case e, ok = <-inp2:
			if ok {
				out <- e
			} else {
				closed = true // break out of the loop				}
			}
		}
	}

	// inp might not be closed yet. Drain it.
	for e = range inp {
		out <- e
	}
}

// End of itemFanIn2 simple binary Fan-In
// ===========================================================================

// Note: pipeitemAdjust imports "container/ring" for the expanding buffer.

// ===========================================================================
// Beg of itemPipeAdjust

// itemPipeAdjust returns a channel to receive
// all `inp`
// buffered by a itemSendProxy process
// before close.
func (inp itemFrom) itemPipeAdjust(sizes ...int) (out itemFrom) {
	cap, que := senditemProxySizes(sizes...)
	cha := make(chan item, cap)
	go inp.pipeitemAdjust(cha, que)
	return cha
}

// itemTubeAdjust returns a closure around itemPipeAdjust (_, sizes ...int).
func (inp itemFrom) itemTubeAdjust(sizes ...int) (tube func(inp itemFrom) (out itemFrom)) {

	return func(inp itemFrom) (out itemFrom) {
		return inp.itemPipeAdjust(sizes...)
	}
}

// End of itemPipeAdjust
// ===========================================================================

// ===========================================================================
// Beg of senditemProxy

func senditemProxySizes(sizes ...int) (cap, que int) {

	// CAP is the minimum capacity of the buffered proxy channel in `itemSendProxy`
	const CAP = 10

	// QUE is the minimum initially allocated size of the circular queue in `itemSendProxy`
	const QUE = 16

	cap = CAP
	que = QUE

	if len(sizes) > 0 && sizes[0] > CAP {
		que = sizes[0]
	}

	if len(sizes) > 1 && sizes[1] > QUE {
		que = sizes[1]
	}

	if len(sizes) > 2 {
		panic("itemSendProxy: too many sizes")
	}

	return
}

// itemSendProxy returns a channel to serve as a sending proxy to 'out'.
// Uses a goroutine to receive values from 'out' and store them
// in an expanding buffer, so that sending to 'out' never blocks.
//  Note: the expanding buffer is implemented via "container/ring"
//
// Note: itemSendProxy is kept for the Sieve example
// and other dynamic use to be discovered
// even so it does not fit the pipe tube pattern as itemPipeAdjust does.
func itemSendProxy(out itemInto, sizes ...int) (send itemInto) {
	cap, que := senditemProxySizes(sizes...)
	cha := make(chan item, cap)
	go (itemFrom)(cha).pipeitemAdjust(out, que)
	return cha
}

// pipeitemAdjust uses an adjusting buffer to receive from 'inp'
// even so 'out' is not ready to receive yet. The buffer may grow
// until 'inp' is closed and then will shrink by every send to 'out'.
//  Note: the adjusting buffer is implemented via "container/ring"
func (inp itemFrom) pipeitemAdjust(out itemInto, QUE int) {
	defer close(out)
	n := QUE // the allocated size of the circular queue
	first := ring.New(n)
	last := first
	var c itemInto
	var e item
	ok := true
	for ok {
		c = out
		if first == last {
			c = nil // buffer empty: disable output
		} else {
			e = first.Value.(item)
		}
		select {
		case e, ok = <-inp:
			if ok {
				last.Value = e
				if last.Next() == first {
					last.Link(ring.New(n)) // buffer full: expand it
					n *= 2
				}
				last = last.Next()
			}
		case c <- e:
			first = first.Next()
		}
	}

	for first != last {
		out <- first.Value.(item)
		first = first.Unlink(1) // first.Next()
	}
}

// End of senditemProxy
// ===========================================================================

// ===========================================================================
// Beg of itemPipeEnter/Leave - Flapdoors observed by a Waiter

// itemWaiter - as implemented by `*sync.WaitGroup` -
// attends Flapdoors and keeps counting
// who enters and who leaves.
//
// Use itemDoneWait to learn about
// when the facilities are closed.
//
// Note: You may also use Your provided `*sync.WaitGroup.Wait()`
// to know when to close the facilities.
// Just: itemDoneWait is more convenient
// as it also closes the primary channel for You.
//
// Just make sure to have _all_ entrances and exits attended,
// and `Wait()` only *after* You've started flooding the facilities.
type itemWaiter interface {
	Add(delta int)
	Done()
	Wait()
}

// Note: The name is intentionally generic in order to avoid eventual multiple-declaration clashes.

// itemPipeEnter returns a channel to receive
// all `inp`
// and registers throughput
// as arrival
// on the given `sync.WaitGroup`
// until close.
func (inp itemFrom) itemPipeEnter(wg itemWaiter) (out itemFrom) {
	cha := make(chan item)
	go inp.pipeitemEnter(cha, wg)
	return cha
}

// itemPipeLeave returns a channel to receive
// all `inp`
// and registers throughput
// as departure
// on the given `sync.WaitGroup`
// until close.
func (inp itemFrom) itemPipeLeave(wg itemWaiter) (out itemFrom) {
	cha := make(chan item)
	go inp.pipeitemLeave(cha, wg)
	return cha
}

// itemDoneLeave returns a channel to receive
// one signal after
// all throughput on `inp`
// has been registered
// as departure
// on the given `sync.WaitGroup`
// before close.
func (inp itemFrom) itemDoneLeave(wg itemWaiter) (done <-chan struct{}) {
	sig := make(chan struct{})
	go inp.doneitemLeave(sig, wg)
	return sig
}

func (inp itemFrom) pipeitemEnter(out itemInto, wg itemWaiter) {
	defer close(out)
	for i := range inp {
		wg.Add(1)
		out <- i
	}
}

func (inp itemFrom) pipeitemLeave(out itemInto, wg itemWaiter) {
	defer close(out)
	for i := range inp {
		out <- i
		wg.Done()
	}
}

func (inp itemFrom) doneitemLeave(done chan<- struct{}, wg itemWaiter) {
	defer close(done)
	for i := range inp {
		_ = i // discard
		wg.Done()
	}
	done <- struct{}{}
}

// itemTubeEnter returns a closure around itemPipeEnter (wg)
// registering throughput
// as arrival
// on the given `sync.WaitGroup`.
func (inp itemFrom) itemTubeEnter(wg itemWaiter) (tube func(inp itemFrom) (out itemFrom)) {

	return func(inp itemFrom) (out itemFrom) {
		return inp.itemPipeEnter(wg)
	}
}

// itemTubeLeave returns a closure around itemPipeLeave (wg)
// registering throughput
// as departure
// on the given `sync.WaitGroup`.
func (inp itemFrom) itemTubeLeave(wg itemWaiter) (tube func(inp itemFrom) (out itemFrom)) {

	return func(inp itemFrom) (out itemFrom) {
		return inp.itemPipeLeave(wg)
	}
}

// itemFiniLeave returns a closure around `itemDoneLeave(wg)`
// registering throughput
// as departure
// on the given `sync.WaitGroup`.
func (inp itemFrom) itemFiniLeave(wg itemWaiter) func(inp itemFrom) (done <-chan struct{}) {

	return func(inp itemFrom) (done <-chan struct{}) {
		return inp.itemDoneLeave(wg)
	}
}

// itemDoneWait returns a channel to receive
// one signal
// after wg.Wait() has returned and out has been closed
// before close.
//
// Note: Use only *after* You've started flooding the facilities.
func (out itemInto) itemDoneWait(wg itemWaiter) (done <-chan struct{}) {
	cha := make(chan struct{})
	go out.doneitemWait(cha, wg)
	return cha
}

func (out itemInto) doneitemWait(done chan<- struct{}, wg itemWaiter) {
	defer close(done)
	wg.Wait()
	close(out)
	done <- struct{}{} // not really needed - but looks better
}

// itemFiniWait returns a closure around `itemDoneWait(wg)`.
func (out itemInto) itemFiniWait(wg itemWaiter) func(out itemInto) (done <-chan struct{}) {

	return func(out itemInto) (done <-chan struct{}) {
		return out.itemDoneWait(wg)
	}
}

// End of itemPipeEnter/Leave - Flapdoors observed by a Waiter
// ===========================================================================

// ===========================================================================
// Beg of itemStrew - scatter them

// itemStrew returns a slice (of size = size) of channels
// one of which shall receive each inp before close.
func (inp itemFrom) itemStrew(size int) (outS []itemFrom) {
	chaS := make(map[chan item]struct{}, size)
	for i := 0; i < size; i++ {
		chaS[make(chan item)] = struct{}{}
	}

	go inp.strewitem(chaS)

	outS = make([]itemFrom, size)
	i := 0
	for c := range chaS {
		outS[i] = (itemFrom)(c) // convert `chan item` to itemFrom
		i++
	}

	return outS
}

func (inp itemFrom) strewitem(outS map[chan item]struct{}) {

	for i := range inp {
		for !inp.trySenditem(i, outS) {
			time.Sleep(time.Millisecond * 10) // wait a little before retry
		} // !sent
	} // inp

	for o := range outS {
		close(o)
	}
}

func (static itemFrom) trySenditem(inp item, outS map[chan item]struct{}) bool {

	for o := range outS {

		select { // try to send
		case o <- inp:
			return true
		default:
			// keep trying
		}

	} // outS
	return false
}

// End of itemStrew - scatter them
// ===========================================================================

// ===========================================================================
// Beg of itemPipeSeen/itemForkSeen - an "I've seen this item before" filter / forker

// itemPipeSeen returns a channel to receive
// all `inp`
// not been seen before
// while silently dropping everything seen before
// (internally growing a `sync.Map` to discriminate)
// until close.
// Note: itemPipeFilterNotSeenYet might be a better name, but is fairly long.
func (inp itemFrom) itemPipeSeen() (out itemFrom) {
	cha := make(chan item)
	go inp.pipeitemSeenAttr(cha, nil)
	return cha
}

// itemPipeSeenAttr returns a channel to receive
// all `inp`
// whose attribute `attr` has
// not been seen before
// while silently dropping everything seen before
// (internally growing a `sync.Map` to discriminate)
// until close.
// Note: itemPipeFilterAttrNotSeenYet might be a better name, but is fairly long.
func (inp itemFrom) itemPipeSeenAttr(attr func(a item) interface{}) (out itemFrom) {
	cha := make(chan item)
	go inp.pipeitemSeenAttr(cha, attr)
	return cha
}

// itemForkSeen returns two channels, `new` and `old`,
// where `new` is to receive
// all `inp`
// not been seen before
// and `old`
// all `inp`
// seen before
// (internally growing a `sync.Map` to discriminate)
// until close.
func (inp itemFrom) itemForkSeen() (new, old itemFrom) {
	cha1 := make(chan item)
	cha2 := make(chan item)
	go inp.forkitemSeenAttr(cha1, cha2, nil)
	return cha1, cha2
}

// itemForkSeenAttr returns two channels, `new` and `old`,
// where `new` is to receive
// all `inp`
// whose attribute `attr` has
// not been seen before
// and `old`
// all `inp`
// seen before
// (internally growing a `sync.Map` to discriminate)
// until close.
func (inp itemFrom) itemForkSeenAttr(attr func(a item) interface{}) (new, old itemFrom) {
	cha1 := make(chan item)
	cha2 := make(chan item)
	go inp.forkitemSeenAttr(cha1, cha2, attr)
	return cha1, cha2
}

func (inp itemFrom) pipeitemSeenAttr(out itemInto, attr func(a item) interface{}) {
	defer close(out)

	if attr == nil { // Make `nil` value useful
		attr = func(a item) interface{} { return a }
	}

	seen := sync.Map{}
	for i := range inp {
		if _, visited := seen.LoadOrStore(attr(i), struct{}{}); visited {
			// drop i silently
		} else {
			out <- i
		}
	}
}

func (inp itemFrom) forkitemSeenAttr(new, old itemInto, attr func(a item) interface{}) {
	defer close(new)
	defer close(old)

	if attr == nil { // Make `nil` value useful
		attr = func(a item) interface{} { return a }
	}

	seen := sync.Map{}
	for i := range inp {
		if _, visited := seen.LoadOrStore(attr(i), struct{}{}); visited {
			old <- i
		} else {
			new <- i
		}
	}
}

// itemTubeSeen returns a closure around itemPipeSeen()
// (silently dropping every item seen before).
func (inp itemFrom) itemTubeSeen() (tube func(inp itemFrom) (out itemFrom)) {

	return func(inp itemFrom) (out itemFrom) {
		return inp.itemPipeSeen()
	}
}

// itemTubeSeenAttr returns a closure around itemPipeSeenAttr(attr)
// (silently dropping every item
// whose attribute `attr` was
// seen before).
func (inp itemFrom) itemTubeSeenAttr(attr func(a item) interface{}) (tube func(inp itemFrom) (out itemFrom)) {

	return func(inp itemFrom) (out itemFrom) {
		return inp.itemPipeSeenAttr(attr)
	}
}

// End of itemPipeSeen/itemForkSeen - an "I've seen this item before" filter / forker
// ===========================================================================
