@Echo.
@cd internal\
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

@pause