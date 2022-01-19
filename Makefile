GOPATH:=$(shell go env GOPATH)

.PHONY: init
# init env
init:
	go get -u github.com/102er/protoc-gen-gogo-errors

.PHONY: errors
# generate go protoc code
# 官方带的proto-gen
errors:
	protoc  --go_out=paths=source_relative:. \
			 errors/errors.proto

.PHONY: ex-errors
# 测试
ex-errors:
	protoc  --proto_path=. \
			-I=./errors \
			--gogo_out=paths=source_relative:. \
			--gogo-errors_out=paths=source_relative:. \
		    examples/v1/hello_world.proto