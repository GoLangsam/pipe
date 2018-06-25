# `v3` - Improve the `traffic` network

## Overview
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
Instead we use `SitePipeEnter` (a companion of `SiteDoneLeave`) at the entrance of our network processor.

Thus, the network becomes more flexible and more self-contained and gives less burden to it's surroundings.

Pushing the types `site` and it's related sites traffic and `result` into separate sub-packages
is just a little more tidying - respecting the original `Crawler` and it's living space.

----

Some remarks regarding changes to source files compared with the [previous](../v2) version:

## [`traffic.go`](traffic/traffic.go)
Simplify `Feed` as explained, and add two processes (`SitePipeEnter` and `SitePipeAdjust) to the network.

## [`genny.go`](traffic/genny.go) in `traffic/`
Just add a line to use `adjust.go`

## [`site.go`](site.go)
Only another package name

## [`crawling.go`](crawling.go)
Just import the new sub-packages, and adjust where need.

## [`crawler_test.go`](crawler_test.go)
Just the import path.

## Changes to [`crawler.go`](crawler.go)
No need to touch.

----
[Back to Overview](../)
