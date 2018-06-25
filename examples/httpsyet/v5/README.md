# `v5` - Just for the fun of it

## Overview
As You see in [`genny.go`](v5/sites/genny.go) now `pipe/m` is used -
and `m` generate methods on `traffic` (instead of package functions).

Further, using `anyThing=Site` (watch the change to the public type!)
generated stuff becomes public (where intended) and can be seen in godoc.

----

Some remarks regarding changes to source files compared with the [previous](../v4) version:

## [`traffic.go`](traffic/traffic.go)
Functions become methods.

## [`genny.go`](traffic/genny.go) in `traffic/`
Now take from `s` insted `m´.

## [`site.go`](site.go)
No change.

## [`crawling.go`](crawling.go)
./.

## [`crawler_test.go`](crawler_test.go)
./.

## Changes to [`crawler.go`](crawler.go)
./.

----
[Back to Overview](../)