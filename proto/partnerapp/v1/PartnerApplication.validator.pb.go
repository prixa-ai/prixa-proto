// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/partnerapp/v1/PartnerApplication.proto

package prixa_partnerapplication_v1

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	regexp "regexp"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

var _regex_PartnerAppResponseData_Id = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)

func (this *PartnerAppResponseData) Validate() error {
	if !_regex_PartnerAppResponseData_Id.MatchString(this.Id) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.Id))
	}
	if this.Id == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must not be an empty string`, this.Id))
	}
	if this.Name == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must not be an empty string`, this.Name))
	}
	if this.Status == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Status", fmt.Errorf(`value '%v' must not be an empty string`, this.Status))
	}
	if this.SecretKey == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("SecretKey", fmt.Errorf(`value '%v' must not be an empty string`, this.SecretKey))
	}
	if this.CreatedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.CreatedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("CreatedAt", err)
		}
	}
	if this.UpdatedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.UpdatedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("UpdatedAt", err)
		}
	}
	if this.Ancillary != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Ancillary); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Ancillary", err)
		}
	}
	if this.Theme != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Theme); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Theme", err)
		}
	}
	return nil
}
func (this *Colors) Validate() error {
	return nil
}
func (this *Theme) Validate() error {
	if this.Colors != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Colors); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Colors", err)
		}
	}
	return nil
}

var _regex_CreatePartnerApplicationRequest_PartnerId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)

func (this *CreatePartnerApplicationRequest) Validate() error {
	if this.Name == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must not be an empty string`, this.Name))
	}
	if !_regex_CreatePartnerApplicationRequest_PartnerId.MatchString(this.PartnerId) {
		return github_com_mwitkow_go_proto_validators.FieldError("PartnerId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.PartnerId))
	}
	if this.PartnerId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("PartnerId", fmt.Errorf(`value '%v' must not be an empty string`, this.PartnerId))
	}
	if this.Theme != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Theme); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Theme", err)
		}
	}
	return nil
}
func (this *CreatePartnerApplicationResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}

var _regex_DeletePartnerApplicationRequest_PartnerId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)
var _regex_DeletePartnerApplicationRequest_ApplicationId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)

func (this *DeletePartnerApplicationRequest) Validate() error {
	if !_regex_DeletePartnerApplicationRequest_PartnerId.MatchString(this.PartnerId) {
		return github_com_mwitkow_go_proto_validators.FieldError("PartnerId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.PartnerId))
	}
	if this.PartnerId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("PartnerId", fmt.Errorf(`value '%v' must not be an empty string`, this.PartnerId))
	}
	if !_regex_DeletePartnerApplicationRequest_ApplicationId.MatchString(this.ApplicationId) {
		return github_com_mwitkow_go_proto_validators.FieldError("ApplicationId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.ApplicationId))
	}
	if this.ApplicationId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ApplicationId", fmt.Errorf(`value '%v' must not be an empty string`, this.ApplicationId))
	}
	return nil
}
func (this *DeletePartnerApplicationResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}

var _regex_UpdatePartnerApplicationRequest_PartnerId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)
var _regex_UpdatePartnerApplicationRequest_ApplicationId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)

func (this *UpdatePartnerApplicationRequest) Validate() error {
	if this.Name == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must not be an empty string`, this.Name))
	}
	if !_regex_UpdatePartnerApplicationRequest_PartnerId.MatchString(this.PartnerId) {
		return github_com_mwitkow_go_proto_validators.FieldError("PartnerId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.PartnerId))
	}
	if this.PartnerId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("PartnerId", fmt.Errorf(`value '%v' must not be an empty string`, this.PartnerId))
	}
	if !_regex_UpdatePartnerApplicationRequest_ApplicationId.MatchString(this.ApplicationId) {
		return github_com_mwitkow_go_proto_validators.FieldError("ApplicationId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.ApplicationId))
	}
	if this.ApplicationId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ApplicationId", fmt.Errorf(`value '%v' must not be an empty string`, this.ApplicationId))
	}
	if this.Theme != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Theme); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Theme", err)
		}
	}
	return nil
}
func (this *UpdatePartnerApplicationResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}

var _regex_GetPartnerApplicationRequest_PartnerId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)
var _regex_GetPartnerApplicationRequest_ApplicationId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)

func (this *GetPartnerApplicationRequest) Validate() error {
	if !_regex_GetPartnerApplicationRequest_PartnerId.MatchString(this.PartnerId) {
		return github_com_mwitkow_go_proto_validators.FieldError("PartnerId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.PartnerId))
	}
	if this.PartnerId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("PartnerId", fmt.Errorf(`value '%v' must not be an empty string`, this.PartnerId))
	}
	if !_regex_GetPartnerApplicationRequest_ApplicationId.MatchString(this.ApplicationId) {
		return github_com_mwitkow_go_proto_validators.FieldError("ApplicationId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.ApplicationId))
	}
	if this.ApplicationId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ApplicationId", fmt.Errorf(`value '%v' must not be an empty string`, this.ApplicationId))
	}
	return nil
}
func (this *GetPartnerApplicationResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}

var _regex_ListPartnerApplicationsRequest_PartnerId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)

func (this *ListPartnerApplicationsRequest) Validate() error {
	if !_regex_ListPartnerApplicationsRequest_PartnerId.MatchString(this.PartnerId) {
		return github_com_mwitkow_go_proto_validators.FieldError("PartnerId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.PartnerId))
	}
	if this.PartnerId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("PartnerId", fmt.Errorf(`value '%v' must not be an empty string`, this.PartnerId))
	}
	return nil
}

var _regex_ListPartnersRequest_PartnerId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)

func (this *ListPartnersRequest) Validate() error {
	if !_regex_ListPartnersRequest_PartnerId.MatchString(this.PartnerId) {
		return github_com_mwitkow_go_proto_validators.FieldError("PartnerId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.PartnerId))
	}
	if this.PartnerId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("PartnerId", fmt.Errorf(`value '%v' must not be an empty string`, this.PartnerId))
	}
	return nil
}
func (this *ListPartnerApplicationsResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}

var _regex_GetAppMetadataRequest_PartnerId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)
var _regex_GetAppMetadataRequest_ApplicationId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)

func (this *GetAppMetadataRequest) Validate() error {
	if !_regex_GetAppMetadataRequest_PartnerId.MatchString(this.PartnerId) {
		return github_com_mwitkow_go_proto_validators.FieldError("PartnerId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.PartnerId))
	}
	if this.PartnerId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("PartnerId", fmt.Errorf(`value '%v' must not be an empty string`, this.PartnerId))
	}
	if !_regex_GetAppMetadataRequest_ApplicationId.MatchString(this.ApplicationId) {
		return github_com_mwitkow_go_proto_validators.FieldError("ApplicationId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.ApplicationId))
	}
	if this.ApplicationId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ApplicationId", fmt.Errorf(`value '%v' must not be an empty string`, this.ApplicationId))
	}
	return nil
}
func (this *AppMetadata) Validate() error {
	return nil
}
func (this *GetAppMetadataResponse) Validate() error {
	if this.Metadata != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Metadata); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Metadata", err)
		}
	}
	return nil
}

var _regex_UpdateAppMetadataRequest_PartnerId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)
var _regex_UpdateAppMetadataRequest_ApplicationId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)

func (this *UpdateAppMetadataRequest) Validate() error {
	if !_regex_UpdateAppMetadataRequest_PartnerId.MatchString(this.PartnerId) {
		return github_com_mwitkow_go_proto_validators.FieldError("PartnerId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.PartnerId))
	}
	if this.PartnerId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("PartnerId", fmt.Errorf(`value '%v' must not be an empty string`, this.PartnerId))
	}
	if !_regex_UpdateAppMetadataRequest_ApplicationId.MatchString(this.ApplicationId) {
		return github_com_mwitkow_go_proto_validators.FieldError("ApplicationId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.ApplicationId))
	}
	if this.ApplicationId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ApplicationId", fmt.Errorf(`value '%v' must not be an empty string`, this.ApplicationId))
	}
	if this.Metadata != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Metadata); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Metadata", err)
		}
	}
	return nil
}
func (this *UpdateAppMetadataResponse) Validate() error {
	return nil
}
func (this *AncillaryControl) Validate() error {
	return nil
}
