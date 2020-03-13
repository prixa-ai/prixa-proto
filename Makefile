install:
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -u github.com/mwitkow/go-proto-validators/protoc-gen-govalidators

build: clean install
	./gen.sh

clean:
	rm -f $(shell find proto -name '*.pb.*') & rm -f $(shell find proto -name '*.swagger.json')

.PHONY: clean build
