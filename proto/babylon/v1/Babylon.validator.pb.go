// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/babylon/v1/babylon.proto

package prixa_babylon_v1

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	regexp "regexp"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *CreateDoctorRequest) Validate() error {
	if this.Name == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must not be an empty string`, this.Name))
	}
	return nil
}
func (this *CreateDoctorResponse) Validate() error {
	return nil
}
func (this *DoctorData) Validate() error {
	if this.Status == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Status", fmt.Errorf(`value '%v' must not be an empty string`, this.Status))
	}
	return nil
}

var _regex_CreateChatInboxRequest_PartnerId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)
var _regex_CreateChatInboxRequest_SpecialityId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)

func (this *CreateChatInboxRequest) Validate() error {
	if !_regex_CreateChatInboxRequest_PartnerId.MatchString(this.PartnerId) {
		return github_com_mwitkow_go_proto_validators.FieldError("PartnerId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.PartnerId))
	}
	if this.PartnerId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("PartnerId", fmt.Errorf(`value '%v' must not be an empty string`, this.PartnerId))
	}
	if !_regex_CreateChatInboxRequest_SpecialityId.MatchString(this.SpecialityId) {
		return github_com_mwitkow_go_proto_validators.FieldError("SpecialityId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.SpecialityId))
	}
	if this.SpecialityId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("SpecialityId", fmt.Errorf(`value '%v' must not be an empty string`, this.SpecialityId))
	}
	return nil
}
func (this *CreateChatInboxResponse) Validate() error {
	return nil
}
func (this *ChatInboxData) Validate() error {
	if this.PartnerId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("PartnerId", fmt.Errorf(`value '%v' must not be an empty string`, this.PartnerId))
	}
	if this.SpecialityId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("SpecialityId", fmt.Errorf(`value '%v' must not be an empty string`, this.SpecialityId))
	}
	return nil
}
func (this *GetDoctorRequest) Validate() error {
	return nil
}
func (this *GetDoctorResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *ListDoctorsRequest) Validate() error {
	return nil
}
func (this *ListDoctorsResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *GetChatInboxRequest) Validate() error {
	return nil
}
func (this *GetChatInboxResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *ListChatInboxesRequest) Validate() error {
	return nil
}
func (this *ListChatInboxesResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}