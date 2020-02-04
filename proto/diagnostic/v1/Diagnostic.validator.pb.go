// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/diagnostic/v1/Diagnostic.proto

package prixa_diagnostic_v1

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/golang/protobuf/ptypes/any"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *BotConversationRequest) Validate() error {
	if this.Reply != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Reply); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Reply", err)
		}
	}
	return nil
}
func (this *ReplyDatas) Validate() error {
	for _, item := range this.Preconditions {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Preconditions", err)
			}
		}
	}
	return nil
}
func (this *PreconditionsDatasRequest) Validate() error {
	for _, item := range this.Datas {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Datas", err)
			}
		}
	}
	return nil
}
func (this *PreconditionsDatasRequestProp) Validate() error {
	return nil
}
func (this *BotConversationResponse) Validate() error {
	if this.Result != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Result); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Result", err)
		}
	}
	for _, item := range this.LogEvents {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("LogEvents", err)
			}
		}
	}
	return nil
}
func (this *ResultDatas) Validate() error {
	if this.Messages != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Messages); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Messages", err)
		}
	}
	if this.Actions != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Actions); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Actions", err)
		}
	}
	for _, item := range this.Preconditions {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Preconditions", err)
			}
		}
	}
	return nil
}
func (this *MessagesData) Validate() error {
	return nil
}
func (this *ActionDatas) Validate() error {
	for _, item := range this.Value {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Value", err)
			}
		}
	}
	return nil
}
func (this *ValueDatas) Validate() error {
	return nil
}
func (this *PreconditionsDatas) Validate() error {
	return nil
}
func (this *LogEvents) Validate() error {
	return nil
}
func (this *SendEmailRequest) Validate() error {
	return nil
}
func (this *SendEmailResponse) Validate() error {
	if this.Message == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Message", fmt.Errorf(`value '%v' must not be an empty string`, this.Message))
	}
	return nil
}
func (this *SendSurveyRequest) Validate() error {
	if this.ApplicationId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ApplicationId", fmt.Errorf(`value '%v' must not be an empty string`, this.ApplicationId))
	}
	if this.Feedback == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Feedback", fmt.Errorf(`value '%v' must not be an empty string`, this.Feedback))
	}
	if this.PartnerId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("PartnerId", fmt.Errorf(`value '%v' must not be an empty string`, this.PartnerId))
	}
	if this.SessionId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("SessionId", fmt.Errorf(`value '%v' must not be an empty string`, this.SessionId))
	}
	if this.UserId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("UserId", fmt.Errorf(`value '%v' must not be an empty string`, this.UserId))
	}
	return nil
}
func (this *FeedbackContentResponse) Validate() error {
	if this.Question == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Question", fmt.Errorf(`value '%v' must not be an empty string`, this.Question))
	}
	if this.Instruction == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Instruction", fmt.Errorf(`value '%v' must not be an empty string`, this.Instruction))
	}
	return nil
}
func (this *SendFeedbackRequest) Validate() error {
	if this.UserId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("UserId", fmt.Errorf(`value '%v' must not be an empty string`, this.UserId))
	}
	if this.PartnerId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("PartnerId", fmt.Errorf(`value '%v' must not be an empty string`, this.PartnerId))
	}
	if this.ApplicationId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ApplicationId", fmt.Errorf(`value '%v' must not be an empty string`, this.ApplicationId))
	}
	if this.SessionId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("SessionId", fmt.Errorf(`value '%v' must not be an empty string`, this.SessionId))
	}
	if this.SymptomId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("SymptomId", fmt.Errorf(`value '%v' must not be an empty string`, this.SymptomId))
	}
	if this.Question == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Question", fmt.Errorf(`value '%v' must not be an empty string`, this.Question))
	}
	if this.Detail == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Detail", fmt.Errorf(`value '%v' must not be an empty string`, this.Detail))
	}
	return nil
}
func (this *SendFeedbackResponse) Validate() error {
	return nil
}