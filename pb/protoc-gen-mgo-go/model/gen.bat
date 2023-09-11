@setlocal enabledelayedexpansion enableextensions
@set PROTO_FILES=
@for %%x in (*.proto) do @set PROTO_FILES=!PROTO_FILES! %%x
@set PROTO_FILES=%PROTO_FILES:~1%

protoc -I=. -I=%GOPATH%/src  --gogofaster_out=. %PROTO_FILES%
protoc -I=. -I=%GOPATH%/src  --mgo-go_out=extra_import=mlgs:. %PROTO_FILES%

goimports -local mlgs -w .
