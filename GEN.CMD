@Echo.

for /F %%d in (gen.prolog.lst) do go generate ./%%d...

@cd .generate.xx
@pwd
@call gen.cmd
@cd ..
@pwd
@Echo.

for /F %%d in (gen.epilog.lst) do go generate ./%%d...

@cd .generic
@pwd
@call gen.cmd
@cd ..
@pwd
@Echo.

go fmt  ./...
go vet  ./...
go test ./...
go doc -all	>godoc.md
@Echo.
@Echo finished :-)

@pause
