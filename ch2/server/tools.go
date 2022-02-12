//go:build tools
// +build tools

package tools

import (
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go@v1.26"
)
