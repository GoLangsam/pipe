@Echo.
@cd internal
@pwd
@Echo.
go generate -run "genny.*" .
go generate -run "bundle.*" .

@Echo.
@cd ..

go fmt  ./...
go vet  ./...
go test ./...

genny -in pipe.go -out ../pipe.gen -pkg pipe gen "Thing=Any"
@Echo.
@Echo compare ../pipe.* - 2 Differences @ top: generic.Any ...
@Echo.

@pause
