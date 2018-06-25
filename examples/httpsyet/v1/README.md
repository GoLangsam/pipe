# `v1` - The separation of concerns

## A refacturing

Inspired by [Jorin](https://jorin.me/about/)'s "qvl.io/httpsyet/httpsyet",
which I stumbled upon (via [GoLangWeekly](https://golangweekly.com/)).

So, as a "real life" example, the original got refactored with focus solely on the aspects of conncurrency,
and intentionally and respectfully each and all code related to the actual crawl functionality was left as untouched as possible.

Please feel free to compare the refactored `crawler.go` with the untouched original `crawler.go.ori.304` (304 LoC)
or see `crawler.go.mini.224` where the parts which became obsolete are removed entirely (and not only commented out).

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

Please note: no `sync-WaitGroup` is needed around the parallel processes (as did the original),
Also Done...'s results may be safely discarded. (The traffic congestion is monitored another way.)

So, how to get there?

## Overview

The original `Crawler` "is used as configuration for Run." only,
and this limited/focused purpose deserves respect.

### `crawling`
In order to give a home to the data structures relevant during crawling,
a new `type crawling struct` (in new `crawling.go` - see below) represents a crawling Crawler.

`Crawler` (the config) and `traffic` are embedded anonymously;
thus `crawling` inherits their respective methods.

Note: The original implementation uses four hand-made channels and
very cleverly orchestrates their handling. Too clever, may be.

Two channels become obsolete:

- `queue` becomes obsolete as feed-back is sent directly into `c.sites`.
- `wait` becomes a `*sync-WaitGroup` (to keep track of the traffic inside the circular net)

The remainig two channels `sites` and `results` got a new home in `crawling`:

- `crawling.add` registers entering urls (synchonously, and parallal!)
- `crawling.goWaitAndClose` patiently awaits it's Wait()
- `crawling.crawl` decrements every crawled site
- `sitePipeLeave` decrements the "I've seen your url before"-sites

----

Some remarks regarding source files follow:

## [`crawling.go`](crawling.go)
- defines `type crawling` to represent a crawling Crawler.

- `Crawler.crawling` instatiates a new `crawling`,
  and calls it `crawling.crawling` (please forgive the pun), which
  - builds the process network (see above)
  - feeds the initial urls (using the original `func queueURLs`)
  - launches the closer (who simply does a `crawling.Wait()` before he closes the channels owned by `crawling`)

## [`crawler_test.go`](crawler_test.go)
As we feed sites back into the crawling in parallel (which did not happen originally due to the use of channel `queue`)
the `visited` map needs to become a guarded map (defined at the end of the source file) in order pass the tests.

Feel free to compare with `crawler_test.go.ori`.

## [`genny.go`](genny.go)
Just contains the `go:generate` directives for `genny` to generate what we need from the generic `pipe/s` function library.

## Changes to [`crawler.go`](crawler.go)

- `func (c Crawler) Run() error`
  - typo corrected: "Run the cralwer." => "Run the crawler."
  - ca 30 LoC after initial validation removed,
  - finish with `<-c.crawling(urls)` instead - wait for crawling to finish

- `func makeQueue()`
  - completely removed - no need
  - ca 35 LoC

- `func (c Crawler) worker`
  - remove the `for s := range sites` loop
  - the loop body becomes a method of crawling: `func (c *crawling) crawlSite(s site) (urls []*url.URL)`
  - send results (now typed) into `c.results` (instead of old argument `results`)

- `func queueURLs`
  - is now launched in `func (c *crawling) add`
  
Thus, ca 80 LoC are removed / deactivated, and:
- no channel is created
- no go routine in launched
- only two send's remain:
  - `c.results <- ...` from `crawlSite(s site)`
  - `queue <- site` from `queueURLs`, now called with `c.sites` as argument `queue` (from `func (c *crawling) add`).

----
[Back to Overview](../)