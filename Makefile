BUILDPATH=proto/diagnostic/v1
build:
	protoc -I $(BUILDPATH) $(BUILDPATH)/Diagnostic.proto --go_out=$(BUILDPATH)
	protoc -I $(BUILDPATH) $(BUILDPATH)/EmailNotification.proto --go_out=$(BUILDPATH)
	protoc -I $(BUILDPATH) $(BUILDPATH)/Partner.proto --go_out=$(BUILDPATH)
	protoc -I $(BUILDPATH) $(BUILDPATH)/PartnerApplication.proto --go_out=$(BUILDPATH)