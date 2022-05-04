// Code generated by github.com/kamilsk/egg. DO NOT EDIT.

//go:build tools

package tools

import (
	_ "github.com/gogo/protobuf/protoc-gen-gofast"
	_ "github.com/golang/mock/mockgen"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"
	_ "github.com/twitchtv/twirp/protoc-gen-twirp"
	_ "golang.org/x/exp/cmd/gorelease"
	_ "golang.org/x/tools/cmd/goimports"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)

//go:generate go install github.com/gogo/protobuf/protoc-gen-gofast
//go:generate go install github.com/golang/mock/mockgen
//go:generate go install github.com/golangci/golangci-lint/cmd/golangci-lint
//go:generate go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
//go:generate go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
//go:generate go install github.com/twitchtv/twirp/protoc-gen-twirp
//go:generate go install golang.org/x/exp/cmd/gorelease
//go:generate go install golang.org/x/tools/cmd/goimports
//go:generate go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
//go:generate go install google.golang.org/protobuf/cmd/protoc-gen-go
