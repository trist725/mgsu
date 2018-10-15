//go:generate protoc --enum-go_out=. common.proto login.proto version.proto
//go:generate protoc --myjson-go_out=. common.proto login.proto version.proto
//go:generate protoc --doc_out=html,index.html:. common.proto login.proto version.proto
package msg
