# `v4` - Teach `traffic` to be more gentle and easy to use

## Overview
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

----

Some remarks regarding changes to source files compared with the [previous](../v3) version:

## [`traffic.go`](traffic/traffic.go)
Implement a.m. improvements straight-forward. Note: The network itself remains as is.

## [`genny.go`](traffic/genny.go) in `traffic/`
Just change to private `site`.

## [`site.go`](site.go)
Just make it's (previously public) methods (`Attr` & `Print`) private.

## [`crawling.go`](crawling.go)
Much more focused and compact now.

## [`crawler_test.go`](crawler_test.go)
Just the import path.

## Changes to [`crawler.go`](crawler.go)
No need to touch.

----
[Back to Overview](../)