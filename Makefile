.PHONY: build
build: vendor-proto .generate .build

PHONY: .generate
.generate:
		mkdir -p pkg/ocp-project-api
		protoc -I vendor.protogen \
				--go_out=pkg/ocp-project-api --go_opt=paths=import \
				--go-grpc_out=pkg/ocp-project-api --go-grpc_opt=paths=import \
				--grpc-gateway_out=pkg/ocp-project-api \
				--grpc-gateway_opt=logtostderr=true \
				--grpc-gateway_opt=paths=import \
				--validate_out lang=go:pkg/ocp-project-api \
				--swagger_out=allow_merge=true,merge_file_name=ocp-project-api:swagger \
				api/ocp-project-api/ocp-project-api.proto
		mv pkg/ocp-project-api/github.com/ozoncp/ocp-project-api/pkg/ocp-project-api/* pkg/ocp-project-api/
		rm -rf pkg/ocp-project-api/github.com
		mkdir -p cmd/ocp-project-api

		mkdir -p pkg/ocp-repo-api
		protoc -I vendor.protogen \
				--go_out=pkg/ocp-repo-api --go_opt=paths=import \
				--go-grpc_out=pkg/ocp-repo-api --go-grpc_opt=paths=import \
				--grpc-gateway_out=pkg/ocp-repo-api \
				--grpc-gateway_opt=logtostderr=true \
				--grpc-gateway_opt=paths=import \
				--validate_out lang=go:pkg/ocp-repo-api \
				--swagger_out=allow_merge=true,merge_file_name=ocp-repo-api:swagger \
				api/ocp-repo-api/ocp-repo-api.proto
		mv pkg/ocp-repo-api/github.com/ozoncp/ocp-project-api/pkg/ocp-repo-api/* pkg/ocp-repo-api/
		rm -rf pkg/ocp-repo-api/github.com
		mkdir -p cmd/ocp-repo-api

PHONY: .build
.build:
		CGO_ENABLED=0 GOOS=linux go build -o bin/ocp-project-api cmd/ocp-project-api/main.go cmd/ocp-project-api/runner.go
		CGO_ENABLED=0 GOOS=linux go build -o bin/ocp-repo-api cmd/ocp-repo-api/main.go cmd/ocp-repo-api/runner.go

PHONY: install
install: build .install

PHONY: .install
install:
		go install cmd/grpc-server/main.go

PHONY: vendor-proto
vendor-proto: .vendor-proto

PHONY: .vendor-proto
.vendor-proto:
		mkdir -p vendor.protogen
		mkdir -p vendor.protogen/api/ocp-project-api
		cp api/ocp-project-api/ocp-project-api.proto vendor.protogen/api/ocp-project-api
		mkdir -p vendor.protogen/api/ocp-repo-api
		cp api/ocp-repo-api/ocp-repo-api.proto vendor.protogen/api/ocp-repo-api
		@if [ ! -d vendor.protogen/google ]; then \
			git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
			mkdir -p  vendor.protogen/google/ &&\
			mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
			rm -rf vendor.protogen/googleapis ;\
		fi
		@if [ ! -d vendor.protogen/github.com/envoyproxy ]; then \
			mkdir -p vendor.protogen/github.com/envoyproxy &&\
			git clone https://github.com/envoyproxy/protoc-gen-validate vendor.protogen/github.com/envoyproxy/protoc-gen-validate ;\
		fi


.PHONY: deps
deps: install-go-deps

.PHONY: install-go-deps
install-go-deps: .install-go-deps

.PHONY: .install-go-deps
.install-go-deps:
		ls go.mod || go mod init
		go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
		go get -u github.com/golang/protobuf/proto
		go get -u github.com/golang/protobuf/protoc-gen-go
		go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
		go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
		go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
		go install github.com/envoyproxy/protoc-gen-validate
