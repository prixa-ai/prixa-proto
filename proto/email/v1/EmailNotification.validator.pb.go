// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/email/v1/EmailNotification.proto

package prixa_emailnotification_v1

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	_ "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *EmailDiagnosticResultRequest) Validate() error {
	if this.To == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("To", fmt.Errorf(`value '%v' must not be an empty string`, this.To))
	}
	if this.Name == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must not be an empty string`, this.Name))
	}
	if this.Gender == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Gender", fmt.Errorf(`value '%v' must not be an empty string`, this.Gender))
	}
	if this.AgeRange == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("AgeRange", fmt.Errorf(`value '%v' must not be an empty string`, this.AgeRange))
	}
	if !(this.WeightInKg >= 1) {
		return github_com_mwitkow_go_proto_validators.FieldError("WeightInKg", fmt.Errorf(`value '%v' must be greater than or equal to '1'`, this.WeightInKg))
	}
	if !(this.WeightInKg <= 999) {
		return github_com_mwitkow_go_proto_validators.FieldError("WeightInKg", fmt.Errorf(`value '%v' must be lower than or equal to '999'`, this.WeightInKg))
	}
	if !(this.HeightInCm >= 40) {
		return github_com_mwitkow_go_proto_validators.FieldError("HeightInCm", fmt.Errorf(`value '%v' must be greater than or equal to '40'`, this.HeightInCm))
	}
	if !(this.HeightInCm <= 300) {
		return github_com_mwitkow_go_proto_validators.FieldError("HeightInCm", fmt.Errorf(`value '%v' must be lower than or equal to '300'`, this.HeightInCm))
	}
	if !(this.BmiIndex >= 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("BmiIndex", fmt.Errorf(`value '%v' must be greater than or equal to '0'`, this.BmiIndex))
	}
	if !(this.BmiIndex <= 99) {
		return github_com_mwitkow_go_proto_validators.FieldError("BmiIndex", fmt.Errorf(`value '%v' must be lower than or equal to '99'`, this.BmiIndex))
	}
	if this.Subject == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Subject", fmt.Errorf(`value '%v' must not be an empty string`, this.Subject))
	}
	if this.Preconditions == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Preconditions", fmt.Errorf(`value '%v' must not be an empty string`, this.Preconditions))
	}
	if this.MainSymptoms == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("MainSymptoms", fmt.Errorf(`value '%v' must not be an empty string`, this.MainSymptoms))
	}
	if this.Symptoms == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Symptoms", fmt.Errorf(`value '%v' must not be an empty string`, this.Symptoms))
	}
	if this.NotKnown == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("NotKnown", fmt.Errorf(`value '%v' must not be an empty string`, this.NotKnown))
	}
	if this.Unknown == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Unknown", fmt.Errorf(`value '%v' must not be an empty string`, this.Unknown))
	}
	if this.PrixaResults == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("PrixaResults", fmt.Errorf(`value '%v' must not be an empty string`, this.PrixaResults))
	}
	return nil
}
