# `v2` - A new home for the processing network

## The new struct `traffic`

## Overview

As the processing network is mainly about the site channel and it's observing WaitGroup,
Thus, it's natural to factor it out into a new struct `traffic` and into a new sourcefile.
Anonymous embedding allows seamless use in `crawling` - and it's sourcefile gets a little more compact.

Last, but not least there is also this string called `result` passing through a channel.
In order to improve clarity we give it a named type - go is a type safe language.

----

Some remarks regarding changes to source files compared with the [previous](../v1) version:

## [`traffic.go`](traffic.go)
New home for guarded traffic. Move `Processor` and `Feed` (formerly `add`) into here.

## [`site.go`](site.go)
Moved `func queueURLs` from [`crawler.go`](crawler.go) into here.

I will regret this later.

## [`crawling.go`](crawling.go)
Use type aliases:
```go
type site = Site
type traffic = Traffic
```

- Make `result` an explicit type now.
- Move methods `processor` and `add` into new [`traffic.go`](traffic.go).
- Make use of new `traffic`.

## [`crawler_test.go`](crawler_test.go)
Just the import path.

## [`genny.go`](genny.go)
Adjust to having `result` as explicit type now.

## Changes to [`crawler.go`](crawler.go)

No need here to touch the previosly refactured `crawler.go` 
except for the one line where the result is sent and we simply need to cast it to the new explicit type now

----
[Back to Overview](../)
