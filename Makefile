build:
	protoc \
		-I proto/diagnostic/v1 \
		-I proto/support/v1:. \
		proto/diagnostic/v1/Partner.proto --go_out=proto/diagnostic/v1
	protoc \
		-I proto/diagnostic/v1 \
		-I proto/support/v1:. \
		proto/diagnostic/v1/PartnerApplication.proto --go_out=proto/diagnostic/v1
	protoc \
		-I proto/diagnostic/v1 \
		proto/diagnostic/v1/Diagnostic.proto --go_out=proto/diagnostic/v1
	protoc \
		-I proto/diagnostic/v1 \
		proto/diagnostic/v1/EmailNotification.proto --go_out=proto/diagnostic/v1

	protoc \
		-I proto/support/v1 \
		proto/support/v1/Timestamp.proto --go_out=proto/support/v1