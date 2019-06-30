@Echo.
@cd internal\
@pwd
@Echo.
go generate .
@Echo.
@Echo compare .generate.xx/pipe.* - 2 Differences
@Echo.

@pause

@Echo.
@cd ..
@pwd
@Echo.
go generate .
@Echo.
@Echo compare xxl\internal\01-any-mode.go* -  4 Differences
@Echo compare xxs\internal\01-any-mode.go* - 10 Differences
@Echo.

@pause

@Echo.
@xcopy	/Y	..\xxs\internal\01-any-mode.go	..\xxsl\internal\01-any-supply.go
@xcopy	/Y	..\xxl\internal\01-any-mode.go	..\xxsl\internal\01-any-demand.go
@Echo.

@pause
