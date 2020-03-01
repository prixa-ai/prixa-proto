#!/bin/sh

function gen() {
	protoc -I/usr/local/include -I. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/mwitkow/go-proto-validators \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:. \
		--grpc-gateway_out=logtostderr=true:. \
		--govalidators_out=. \
		$1
	echo "Generated source code for: \""$1"\""
}

for protoFile in $(find proto -name '*.proto'); do
	gen $protoFile
done
