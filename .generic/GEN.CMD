@Echo.
@cd internal
@pwd
@Echo.
go generate .

@Echo.
@cd ..
genny -in pipe.go -out ../pipe.gen -pkg pipe gen "Thing=Any"
@Echo.
@Echo compare ../pipe.* - 2 Differences @ top: generic.Any ...
@Echo.

@pause
