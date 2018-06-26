# `v5` - Just for the fun of it

## Overview
As You see in [`genny.go`](v5/sites/genny.go) now `pipe/m` is used -
and `m` generate methods on `traffic` (instead of package functions).

Further, using `anyThing=Site` (watch the change to the public type!)
generated stuff becomes public (where intended) and can be seen in godoc.

Last, but not least, the idea to have `Results` in a separate package has been abandoned,
as this adds more complications than benefit.

Now `result` lives inside the `httpsyet` package again - in `crawling.go`..

----

Some remarks regarding changes to source files compared with the [previous](../v4) version:

## [`traffic.go`](traffic/traffic.go)
Functions become methods.

## [`genny.go`](traffic/genny.go) in `traffic/`
Now take from `s` insted `m´.

## [`site.go`](sites/site.go)
No change.

## [`crawling.go`](crawling.go)
Adjust for `result` becoming a local type again.

## [`crawler_test.go`](crawler_test.go)
Just the import path.

## Changes to [`crawler.go`](crawler.go)
No need to touch.

----
[Back to Overview](../)