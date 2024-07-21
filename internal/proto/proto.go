// Package proto contains protobuf source and compiled files for communication with the frontend.
package proto

//go:generate go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.34.2
//go:generate go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.4.0
//go:generate protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative *.proto
