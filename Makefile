GOPATH:=$(shell go env GOPATH)

.PHONY: errors
# generate go protoc code
# 官方带的proto-gen
errors:
	protoc  --go_out=paths=source_relative:. \
			 errors/errors.proto