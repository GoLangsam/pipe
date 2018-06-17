@Echo.
@cd .generate.xx\internal\
@pwd
@Echo.
go generate .
@Echo.
@Echo compare .generate.xx/pipe.* - 3 Differences
@Echo.

@pause

@Echo.
@cd ..
@pwd
@Echo.
go generate .
@Echo.
@Echo compare xxs\internal\01-any-mode.go* - 8 Differences
@Echo compare xxl\internal\01-any-mode.go* - 2 Differences
@Echo.

@cd ..
@pwd
@Echo.

@pause

@Echo.

@cd .generic
@pwd
@call gen.cmd

@cd ..
@pwd
@Echo.
@xcopy	/Y	xxs\internal\01-any-mode.go	xxsl\internal\01-any-supply.go
@xcopy	/Y	xxl\internal\01-any-mode.go	xxsl\internal\01-any-demand.go
@Echo.

go generate .\...
go test .\...

@Echo.
@Echo finished :-)
