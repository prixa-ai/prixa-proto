install:
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -u github.com/mwitkow/go-proto-validators/protoc-gen-govalidators

build: install
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

	protoc -I/usr/local/include -I. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/mwitkow/go-proto-validators \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:. \
		--grpc-gateway_out=logtostderr=true:. \
		--govalidators_out=. \
		proto/siloam/v1/Siloam.proto


	protoc -I/usr/local/include -I. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/mwitkow/go-proto-validators \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:. \
		--grpc-gateway_out=logtostderr=true:. \
		--govalidators_out=. \
		proto/myprotection/v1/MyProtection.proto		

clean:
	rm $(find proto -name '*.pb.*')

.PHONY: install clean build
