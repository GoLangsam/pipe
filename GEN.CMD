@Echo.

for /F %%d in (gen.prolog.lst) do go generate ./%%d...

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

for /F %%d in (gen.epilog.lst) do go generate ./%%d...

@cd .generic
@pwd
@call gen.cmd
@cd ..
@pwd
@Echo.

@call test.cmd

@Echo.
@Echo finished :-)

@pause
