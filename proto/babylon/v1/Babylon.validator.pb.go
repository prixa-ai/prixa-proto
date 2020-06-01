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
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *CreateDoctorAgentRequest) Validate() error {
	if this.PartnerId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("PartnerId", fmt.Errorf(`value '%v' must not be an empty string`, this.PartnerId))
	}
	if this.SpecialityId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("SpecialityId", fmt.Errorf(`value '%v' must not be an empty string`, this.SpecialityId))
	}
	return nil
}
func (this *CreateDoctorAgentResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *DoctorAgentData) Validate() error {
	if this.PartnerId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("PartnerId", fmt.Errorf(`value '%v' must not be an empty string`, this.PartnerId))
	}
	if this.SpecialityId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("SpecialityId", fmt.Errorf(`value '%v' must not be an empty string`, this.SpecialityId))
	}
	if this.Status == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Status", fmt.Errorf(`value '%v' must not be an empty string`, this.Status))
	}
	return nil
}
func (this *CreateChatInboxRequest) Validate() error {
	return nil
}
func (this *CreateChatInboxResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
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
func (this *GetDoctorAgentRequest) Validate() error {
	return nil
}
func (this *GetDoctorAgentResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *ListDoctorAgentsRequest) Validate() error {
	return nil
}
func (this *ListDoctorAgentsResponse) Validate() error {
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
