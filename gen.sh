#!/bin/sh

function gen() {
	params=()
	if [[ $2 == "proto" ]]; then
			params+=(--go_out=plugins=grpc:. --grpc-gateway_out=logtostderr=true:. --govalidators_out=.)
	fi
	if [[ $2 == "swagger" ]]; then
			params+=(--swagger_out=logtostderr=true:.)
	fi

	protoc -I/usr/local/include -I. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/mwitkow/go-proto-validators \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/ \
		"${params[@]}" \
		$1
}

for protoFile in $(find proto -name '*.proto'); do
	gen $protoFile proto
	echo "Generated source code for: \""$protoFile"\""
done

WithSwagger=(
	proto/user/v1/User.proto
	proto/partner/v1/Partner.proto 
	proto/partnerapp/v1/PartnerApplication.proto 
	proto/email/v1/EmailNotification.proto 
	proto/diagnostic/v1/Diagnostic.proto 
	proto/telemedicine/v1/Telemedicine.proto 
)
for protoSwaggerFile in ${WithSwagger[*]}; do
	gen $protoSwaggerFile swagger
	echo "Generated swagger for: \""$protoSwaggerFile"\""
done
