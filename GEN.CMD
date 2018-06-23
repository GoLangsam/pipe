@Echo.

go generate .\...

@cd .generate.xx
@pwd
@call gen.cmd
@cd ..
@pwd
@Echo.

@Echo.
@xcopy	/Y	xxs\internal\01-any-mode.go	xxsl\internal\01-any-supply.go
@xcopy	/Y	xxl\internal\01-any-mode.go	xxsl\internal\01-any-demand.go
@Echo.

go generate .\...

@cd .generic
@pwd
@call gen.cmd
@cd ..
@pwd
@Echo.

go generate .\...
go fmt  .\...
go test .\...

@Echo.
@Echo finished :-)

@pause
