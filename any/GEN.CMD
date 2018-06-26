@Echo.

go generate -run "genny.*" ./...
go generate -run "bundle.*" ./...

@Echo.

go fmt  ./...
go test ./...

@Echo.
@Echo finished :-)

@pause
