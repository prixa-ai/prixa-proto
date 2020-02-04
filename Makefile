install:
	go get github.com/mwitkow/go-proto-validators/protoc-gen-govalidators

build:
	protoc -I/usr/local/include -I. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/mwitkow/go-proto-validators \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:. \
		--grpc-gateway_out=logtostderr=true:. \
		--govalidators_out=. \
		proto/partner/v1/Partner.proto

	protoc -I/usr/local/include -I. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/mwitkow/go-proto-validators \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:. \
		--grpc-gateway_out=logtostderr=true:. \
		--govalidators_out=. \
		proto/partnerapp/v1/PartnerApplication.proto

	protoc -I/usr/local/include -I. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/mwitkow/go-proto-validators \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:. \
		--grpc-gateway_out=logtostderr=true:. \
		--govalidators_out=. \
		proto/diagnostic/v1/Diagnostic.proto

	protoc -I/usr/local/include -I. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/mwitkow/go-proto-validators \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:. \
		--grpc-gateway_out=logtostderr=true:. \
		--govalidators_out=. \
		proto/email/v1/EmailNotification.proto

