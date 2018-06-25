# A Crawler

## [Use Go Channels to Build a Crawler](https://jorin.me/use-go-channels-to-build-a-crawler/)

[@jorinvo](https://github.com/jorinvo) aims to provide "_an implementation that abstracts the coordination using channels_"
in his [original post](https://jorin.me/use-go-channels-to-build-a-crawler/) (also [here](https://dev.to/jorinvo/use-go-channels-to-build-a-crawler-8jl)),

Having looked at it while ago, having read the explanations and justifications,
I thought it might be a nice usecase for [`pipe/s`](https://github.com/GoLangsam/pipe) generic functions.

## tl;dr
Using functions generated from `pipe/s` (see `genny.go` below),
the actual concurrent network to process the site traffic becomes:

```go
	sites, seen := siteForkSeenAttr(c.sites, site.Attr)
	for _, inp := range siteStrew(sites, size) {
		siteDoneFunc(inp, c.crawl) // sites leave inside crawler's crawl
	}
	siteDoneLeave(seen, c) // seen leave without further processing
```

_Simple, is it not? ;-)_

"No locks. No condition variables. No callbacks." - as Rob Pike loves to reiterate in his slides.

Please note: no `sync-WaitGroup` is needed around the parallel processes (as did the original),
Also Done...'s results may be safely discarded. (The traffic congestion is monitored another way.)

So, how to get there?

- [v0](v0/) - The Original
- [v1](v1/) - The separation of concerns: new struct `crawling`
- [v2](v2/) - A new home for the processing network: new struct `traffic`
- [v3](v3/) - Improve the `traffic` network
- [v4](v4/) - Teach `traffic` to be more gentle and easy to use
- [v5](v5/) - Just for the fun of it: use `pipe/m` to generate methods on `traffic` (instead of package functions)
- And now?

## Overview

In order to abstract (latin: abstrahere = "to take away") "the coordination using channels" it is taken away into a new sourcefile, and a new struct `crawling` gives home to the relevant data.

This leaves `crawler.go` focused on the actual crawl functionality.

"The coordination using channels" is redone on a slightly higher level of abstration: using a pipe network.
Yes, channels connect the parts - but they are not the main focus anymore,

And no `wait` channel (and it's management) is needed to keep track of the traffic.
A properly used `*sync.WaitGroup` is all we need.

And the code for the arrangement of parts is shown above - the body of a `func (c *crawling)` :
- what arrives in `c.sites` is separated (by the `site.Attr` method)
- new (previously unseen) sites get strewed (scattered) onto a slice of channels - size determines the parallelism
  - for each ranged channel all what needs to be done is to apply the crawl function (a method of crawling)
- what had been seen before leaves without further processing (last line)

Please note the four vital functions are taken from the `pipe/s` collection!

[v1](v1/) has more details - please have a look.

And please notice: we've accomplished to relive `crawler.go` and may leave it untoched during what follows.
---

### [v2](v2/) - A new home for the processing network: new struct `traffic` 

As the processing network is mainly about the site channel and it's observing WaitGroup,
Thus, it's natural to factor it out into a new struct `traffic` and into a new sourcefile.
Anonymous embedding allows seamless use in `crawling` - and it's sourcefile gets a little more compact.

Last, but not least there is also this string called `result` passing through a channel.
In order to improve clarity we give it a named type - go is a type safe language.

[v2](v2/) has more details - please have a look.

And please notice: no need here to touch the previosly refactured `crawler.go` except for the one line where 
the result is sent and we simply need to cast it to the new type.

---

### [v3](v3/) - Improve the `traffic` network

It's not easy to see at first, but there is still an interesting issue hiding in the original implementation.

It may become more obvious when we imagine to crawl huge numbers -
e.g. finding a thousand new urls on new each page, and going down a couple of levels.

Look: Each site's feedback spawns a new goroutine in order to `queueURLs` the (e.g. 1000) urls found.

And most likely these will block!

Almost each and every of these many goroutines will block most of the time,
as there are still urls discovered earlier to be fetched and crawled.

And each `queueURLs` carries the full slice of urls - no matter how many have been sent to the processing yet.
(The implementation uses a straightforward `range` and does not attempt to shrink the slice. Idiomatic go.)

We can do differently. No need to wast many huge slices across plenty of blocked goroutines.

A battery called `àdjust` provides a flexibly buffered pipe. so we use `SitePipeAdjust` in our network
and do not need to have `Feed` spawn the `queueURLs` function any more. We may feed synchonously now!

But now we also do not need to bother `Feed` with the need of registering new traffic (using `t.Add(len(urls))`) up front.
Instead we use `SitePipeEnter` (a companion of `SiteDoneLeave`) at the entrnace of our network processor.

Thus, the network becomes more flexible and more self-contained and gives less burden to it's surroundings.

Pushing the types `site` and it's related sites traffic and `result` into separate sub-packages
is just a little more tidying - respecting the original `Crawler` and it's living space.

[v3](v3/) has more details - please have a look.

And please notice: no need here to touch the previosly refactured `crawler.go`.

---

### [v4](v4/) - Teach `traffic` to be more gentle and easy to use

Proper use of the new struct `traffic` is still awkward for [`crawling`](v3/crawling.go).

Thus, time to teach `traffic` to behave better and more robust. Sepcifically:
- new constructor `New()`: don't bother the client with initialisation - and thus now no need anymore to import "sync"
- have `Processor` return a signal channel to broadcast "traffic has subsided and nothing is left to be processed"
- lazy initialisation of this mechanism upon first Feed, and Do() only `synce.Once` 
- new method `Done()` - just a convenince to receive the broadcast channel another way
- wrap the `crawl` function passed to `Processor` and have it register the site having left - thus: no need anymore for crawling to do so in it's crawl method.

So, [`crawling` now](v4/crawling.go) is 20% shorter more more focused on it's own subject, is it not?

Please also note: "launch the results closer" now happens happily before the first "feed initial urls" -
no need anymore to worry for something like "goWaitAndClose is to be used after initial traffic has been added".

The client (crawling) is free to use the channel returned from `Processor` (as it does now)
or may even use `<-crawling.Done()` at *any* time he likes or seems fit (even before(!)) the `Processor` is build.

And: `Done()` is a method familar e.g. from the "context" package - thus easy to use and understand.
Easier as is `<-sites.SiteDoneWait(c.Travel, c)`, is it not?

[v4](v4/) has more details - please have a look.

And please notice: no need here to touch the previosly refactured `crawler.go`.

---

### [v5](v5/) - Just for the fun of it

As You see in [`genny.go`](v5/sites/genny.go) now `pipe/m` is used -
and `m` generate methods on `traffic` (instead of package functions).

Further, using `anyThing=Site` (watch the change to the public type!)
generated stuff becomes public (where intended) and can be seen in godoc.

---

### And now?

Well, it's enough, is it not? Let's call it a day.

Even so ... if we look back into the smallified `crawler.go` it may be seen,
that there is still room to use a channel and to apply a little concurrency.

You may even spot some anti-concurrency pattern - as may come from 'classical' education and experience.

Where? Look at these `links` and `urls` - collections - passed around only after having finished building / appending them.

Any such anchor might as well be sent on a (new) channel right away from where it was found,
get parsed (and filtered) and might become a direct feed for the sites processor.
Just one (more) single go routine would be more than happy to do so ... concurrently.

On the other hand - using concurrency too much and or in the wrong place may lead to very poor results easily.
See e.g. [Go code refactoring : the 23x performance hunt](https://medium.com/@val_deleplace/go-code-refactoring-the-23x-performance-hunt-156746b522f7)
for some very interesting observations and refactorings around concurrent .

----

And now? Well - Let's keep this for another day ;-)

----
