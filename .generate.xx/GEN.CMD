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
@Echo compare xxs\internal\01-any-mode.go* - 11 Differences
@Echo.

@pause

@Echo.
@xcopy	/Y	..\xxl\internal\01-any-mode.go	..\xxsl\internal\01-any-demand.go
@xcopy	/Y	..\xxs\internal\01-any-mode.go	..\xxsl\internal\01-any-supply.go
@Echo.

@pause
