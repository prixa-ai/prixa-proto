// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/analytic/v1/Analytic.proto

package prixa_analytics_v1

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/golang/protobuf/ptypes/any"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *EventLog) Validate() error {
	if this.Date != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Date); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Date", err)
		}
	}
	if this.Metadata != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Metadata); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Metadata", err)
		}
	}
	return nil
}
func (this *NalarEvent) Validate() error {
	if this.EventLog != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.EventLog); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("EventLog", err)
		}
	}
	if this.Reply != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Reply); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Reply", err)
		}
	}
	if this.Result != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Result); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Result", err)
		}
	}
	return nil
}
