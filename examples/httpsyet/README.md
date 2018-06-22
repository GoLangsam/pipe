# A Crawler

## [Use Go Channels to Build a Crawler](https://jorin.me/use-go-channels-to-build-a-crawler/)

[@jorinvo](https://github.com/jorinvo) aims to provide "_an implementation that abstracts the coordination using channels_"
in his [original post](https://jorin.me/use-go-channels-to-build-a-crawler/),

Having looked at it while ago, having read the explanations and justifications,
I thought it might be a nice usecase for [`pipe/s`](https://github.com/GoLangsam/pipe).

## A refacturing

Inspired by [Jorin](https://jorin.me/about/)'s "qvl.io/httpsyet/httpsyet",
which I stumbled upon (via [GoLangWeekly](https://golangweekly.com/)).

So, as a "real life" example, the original got refactored with focus solely on the aspects of conncurrency,
and intentionally and respectfully each and all code related to the actual crawling was left as untouched as possible.

Please feel free to compare the refactored `crawler.go` with the untouched original `crawler.go.ori.304` (304 LoC)
or see `crawler.go.mini.224` where the parts which became obsolete are not commented out, but entirely removed.

## Overview

The original `Crawler` "is used as configuration for Run." only,
and this limited/focused purpose deserves respect.

In order to give a home to the data structures needed during crawling,
a new `type crawling struct` (in new `crawling.go` - see below) represents a crawling Crawler.

`Crawler` (the config) and a `*sync-WaitGroup` are embedded anonymously;
thus `crawling` inherits all their respective methods.

Further, `crawling` becomes the new home for the (remaining) channels involved.

(Note: The original implementation uses four hand-made channels and
very cleverly orchestrates their handling. Too clever, may be.)

Two channels become obsolete:

- `queue` becomes obsolete as we feed-back directly into `c.sites`.
- `wait` becomes a `*sync-WaitGroup` (to keep track of the traffic inside the circular net)

The remainig two channels `sites` and `results` get a new home in the new `crawling`.

Using functions generated from `pipe/s` (see `genny.go` below),
the actual concurrent network to process the sites channel becomes:

```go
	sites, seen := siteForkSeenAttr(c.sites, site.Attr)
	for _, inp := range siteStrew(sites, size) {
		siteDoneFunc(inp, c.crawl) // sites leave inside crawler's crawl
	}
	siteDoneLeave(seen, c) // seen leave without further processing
```

_Simple, is it not? ;-)_

Please note: we do not need any `sync-WaitGroup` around the parallel processes (as did the original),
Also we may safely discard Done...'s results.

Our `*sync-WaitGroup` -embedded in `crawling`- controls the traffic:
- `crawling.add` registers entering urls (synchonously, and parallal!)
- PipeSiteLeave decrements the "I've seen your url before"-sites
- `crawling.crawl` decrements every crawled site
- `crawling.wait` patiently awaits it's Wait()


## `crawling.go`
- defines `type crawling` to represent a crawling Crawler.

- `Crawler.crawling` instatiates a new `crawling`,
  and calls it `crawling.crawling` (please forgive the pun), which
  - builds the process network (see above)
  - feeds the initial urls (using the original `func queueURLs`)
  - launches the closer (who simply does a `crawling.Wait()` before he closes the channels owned by `crawling`)
  - and returns a signal channel to receive a signal upon close of `results`
    (after each has gone thru crawlings `c.report`)

## `crawler_test.go`
As we feed sites back into the crawling in parallel (which did not happen originally due to the use of channel `queue`)
the `visited` map needs to become a guarded map (defined at the end of the source file).
Feel free to compare with `crawler_test.go.ori`.

## `genny.go`
Just contains the `go:generate` comments for `genny` to generate what we need from the `pipe/s` library.

## Changes to `crawler.go`

- `func (c Crawler) Run() error`
  - typo corrected: "Run the cralwer." => "// Run the crawler."
  - ca 30 LoC after initial validation removed,
  - finish with `<-c.crawling(urls)` instead - wait for crawling to finish

- `func makeQueue()`
  - completely removed - no need
  - ca 35 LoC

- `func (c Crawler) worker`
  - remove the `for s := range sites` loop
  - becomes a method of crawling: `func (c *crawling) crawlSite(s site) (urls []*url.URL)`
  - sends into `c.results` (instead of `results`)

- `func queueURLs`
  - is now launched in `func (c *crawling) add`
  
Thus, ca 80 LoC are removed / deactivated, and:
- no channel is created
- no go routine in launched
- only two send's remain:
  - `c.results <- ...` from `crawlSite(s site)`
  - `queue <- site` from `queueURLs`, now called with `c.sites` as argument `queue` (from `func (c *crawling) add`).
