//go:generate protoc --enum-go_out=. user.proto
//go:generate protoc --mgo-go_out=. user.proto
//go:generate protoc --doc_out=html,index.html:. user.proto
package model
