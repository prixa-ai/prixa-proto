// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/user/v1/User.proto

package prixa_user_v1

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	_ "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	regexp "regexp"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

var _regex_AuthData_Email = regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+.[a-z]{2,4}$`)

func (this *AuthData) Validate() error {
	if !_regex_AuthData_Email.MatchString(this.Email) {
		return github_com_mwitkow_go_proto_validators.FieldError("Email", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-z0-9._%+-]+@[a-z0-9.-]+.[a-z]{2,4}$"`, this.Email))
	}
	if this.Password == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Password", fmt.Errorf(`value '%v' must not be an empty string`, this.Password))
	}
	if this.PasswordConfirmation == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("PasswordConfirmation", fmt.Errorf(`value '%v' must not be an empty string`, this.PasswordConfirmation))
	}
	if this.DiagnosticSessionID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("DiagnosticSessionID", fmt.Errorf(`value '%v' must not be an empty string`, this.DiagnosticSessionID))
	}
	return nil
}
func (this *ChronicConditionData) Validate() error {
	return nil
}

var _regex_ProfileData_Email = regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+.[a-z]{2,4}$`)

func (this *ProfileData) Validate() error {
	if !_regex_ProfileData_Email.MatchString(this.Email) {
		return github_com_mwitkow_go_proto_validators.FieldError("Email", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-z0-9._%+-]+@[a-z0-9.-]+.[a-z]{2,4}$"`, this.Email))
	}
	if this.Name == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must not be an empty string`, this.Name))
	}
	if this.Gender == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Gender", fmt.Errorf(`value '%v' must not be an empty string`, this.Gender))
	}
	if this.Password == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Password", fmt.Errorf(`value '%v' must not be an empty string`, this.Password))
	}
	return nil
}
func (this *LoginRequest) Validate() error {
	if this.AuthData != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.AuthData); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("AuthData", err)
		}
	}
	return nil
}
func (this *LoginResponse) Validate() error {
	if this.LoginToken == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("LoginToken", fmt.Errorf(`value '%v' must not be an empty string`, this.LoginToken))
	}
	if this.PersonID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("PersonID", fmt.Errorf(`value '%v' must not be an empty string`, this.PersonID))
	}
	if this.SessionID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("SessionID", fmt.Errorf(`value '%v' must not be an empty string`, this.SessionID))
	}
	return nil
}
func (this *ConsentResponse) Validate() error {
	if this.Data == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Data", fmt.Errorf(`value '%v' must not be an empty string`, this.Data))
	}
	return nil
}
func (this *RegisterRequest) Validate() error {
	if this.AuthData != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.AuthData); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("AuthData", err)
		}
	}
	if this.ProfileData != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ProfileData); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ProfileData", err)
		}
	}
	return nil
}
func (this *RegisterResponse) Validate() error {
	return nil
}
func (this *ResendEmailVerificationResponse) Validate() error {
	return nil
}
func (this *VerifyRegisterRequest) Validate() error {
	if this.RegisterToken == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("RegisterToken", fmt.Errorf(`value '%v' must not be an empty string`, this.RegisterToken))
	}
	if this.DiagnosticSessionID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("DiagnosticSessionID", fmt.Errorf(`value '%v' must not be an empty string`, this.DiagnosticSessionID))
	}
	return nil
}
func (this *VerifyRegisterResponse) Validate() error {
	if this.LoginToken == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("LoginToken", fmt.Errorf(`value '%v' must not be an empty string`, this.LoginToken))
	}
	if this.SessionID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("SessionID", fmt.Errorf(`value '%v' must not be an empty string`, this.SessionID))
	}
	return nil
}

var _regex_ForgetPasswordRequest_Email = regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+.[a-z]{2,4}$`)

func (this *ForgetPasswordRequest) Validate() error {
	if !_regex_ForgetPasswordRequest_Email.MatchString(this.Email) {
		return github_com_mwitkow_go_proto_validators.FieldError("Email", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-z0-9._%+-]+@[a-z0-9.-]+.[a-z]{2,4}$"`, this.Email))
	}
	if this.DiagnosticSessionID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("DiagnosticSessionID", fmt.Errorf(`value '%v' must not be an empty string`, this.DiagnosticSessionID))
	}
	return nil
}
func (this *ForgetPasswordResponse) Validate() error {
	return nil
}
func (this *ForgetPasswordVerifRequest) Validate() error {
	if this.ForgetPwdToken == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ForgetPwdToken", fmt.Errorf(`value '%v' must not be an empty string`, this.ForgetPwdToken))
	}
	return nil
}
func (this *ForgetPasswordVerifResponse) Validate() error {
	return nil
}
func (this *UpdatePasswordRequest) Validate() error {
	if this.Password == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Password", fmt.Errorf(`value '%v' must not be an empty string`, this.Password))
	}
	if this.ConfirmPassword == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ConfirmPassword", fmt.Errorf(`value '%v' must not be an empty string`, this.ConfirmPassword))
	}
	if this.DiagnosticSessionID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("DiagnosticSessionID", fmt.Errorf(`value '%v' must not be an empty string`, this.DiagnosticSessionID))
	}
	return nil
}
func (this *UpdatePasswordResponse) Validate() error {
	if this.LoginToken == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("LoginToken", fmt.Errorf(`value '%v' must not be an empty string`, this.LoginToken))
	}
	if this.SessionID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("SessionID", fmt.Errorf(`value '%v' must not be an empty string`, this.SessionID))
	}
	return nil
}
func (this *UserInfoResponse) Validate() error {
	if this.ProfileData != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ProfileData); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ProfileData", err)
		}
	}
	return nil
}
func (this *ChangePasswordRequest) Validate() error {
	if this.Password == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Password", fmt.Errorf(`value '%v' must not be an empty string`, this.Password))
	}
	if this.ConfirmPassword == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ConfirmPassword", fmt.Errorf(`value '%v' must not be an empty string`, this.ConfirmPassword))
	}
	return nil
}
func (this *ChangePasswordResponse) Validate() error {
	return nil
}
func (this *LogoutResponse) Validate() error {
	return nil
}
