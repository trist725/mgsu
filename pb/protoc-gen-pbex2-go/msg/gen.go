//go:generate protoc --gogofaster_out=. common.proto login.proto version.proto
//go:generate protoc --pbex2-go_out=. common.proto login.proto version.proto
//go:generate protoc --doc_out=html,index.html:. common.proto login.proto version.proto
package msg
