@setlocal enabledelayedexpansion enableextensions
@set PROTO_FILES=
@for %%x in (*.proto) do @set PROTO_FILES=!PROTO_FILES! %%x
@set PROTO_FILES=%PROTO_FILES:~1%

protoc -I=. -I=%GOPATH%\src --gogofaster_out=. %PROTO_FILES%
::protoc -I=. -I=%GOPATH%\src --pbex-go_out=. --pbex-go_opt=tpl:example.tpl %PROTO_FILES%
protoc -I=. -I=%GOPATH%\src --pbex-go_out=. %PROTO_FILES%
protoc -I=. -I=%GOPATH%\src --doc_out=html,index.html:. %PROTO_FILES%
