// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/telemedicine/v1/Telemedicine.proto

package prixa_telemedicine_v1

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/mwitkow/go-proto-validators"
	regexp "regexp"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

var _regex_HospitalData_Id = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)

func (this *HospitalData) Validate() error {
	if !_regex_HospitalData_Id.MatchString(this.Id) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.Id))
	}
	if this.Id == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must not be an empty string`, this.Id))
	}
	if this.Address != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Address); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Address", err)
		}
	}
	if this.Contact != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Contact); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Contact", err)
		}
	}
	if this.Coordinate != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Coordinate); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Coordinate", err)
		}
	}
	for _, item := range this.SpecialitiesList {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("SpecialitiesList", err)
			}
		}
	}
	if this.LastUpdated != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.LastUpdated); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("LastUpdated", err)
		}
	}
	if this.Source != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Source); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Source", err)
		}
	}
	return nil
}

var _regex_DoctorData_Id = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)

func (this *DoctorData) Validate() error {
	if !_regex_DoctorData_Id.MatchString(this.Id) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.Id))
	}
	if this.Id == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must not be an empty string`, this.Id))
	}
	if this.Speciality != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Speciality); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Speciality", err)
		}
	}
	if this.Contact != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Contact); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Contact", err)
		}
	}
	for _, item := range this.HospitalsList {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("HospitalsList", err)
			}
		}
	}
	return nil
}

var _regex_SpecialityData_Id = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)

func (this *SpecialityData) Validate() error {
	if !_regex_SpecialityData_Id.MatchString(this.Id) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.Id))
	}
	if this.Id == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must not be an empty string`, this.Id))
	}
	if this.LastUpdated != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.LastUpdated); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("LastUpdated", err)
		}
	}
	if this.Source != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Source); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Source", err)
		}
	}
	return nil
}

var _regex_AreaData_Id = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)

func (this *AreaData) Validate() error {
	if !_regex_AreaData_Id.MatchString(this.Id) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.Id))
	}
	if this.Id == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must not be an empty string`, this.Id))
	}
	if this.LastUpdated != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.LastUpdated); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("LastUpdated", err)
		}
	}
	if this.Source != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Source); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Source", err)
		}
	}
	return nil
}
func (this *CoordinateData) Validate() error {
	return nil
}
func (this *AddressData) Validate() error {
	return nil
}

var _regex_SourceData_SourceId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)
var _regex_SourceData_OriginId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)

func (this *SourceData) Validate() error {
	if !_regex_SourceData_SourceId.MatchString(this.SourceId) {
		return github_com_mwitkow_go_proto_validators.FieldError("SourceId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.SourceId))
	}
	if this.SourceId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("SourceId", fmt.Errorf(`value '%v' must not be an empty string`, this.SourceId))
	}
	if !_regex_SourceData_OriginId.MatchString(this.OriginId) {
		return github_com_mwitkow_go_proto_validators.FieldError("OriginId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.OriginId))
	}
	if this.OriginId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("OriginId", fmt.Errorf(`value '%v' must not be an empty string`, this.OriginId))
	}
	return nil
}

var _regex_ContactData_PhoneNumber = regexp.MustCompile(`^[+]?[0-9]{10,13}$`)
var _regex_ContactData_Email = regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+.[a-z]{2,4}$`)

func (this *ContactData) Validate() error {
	if !_regex_ContactData_PhoneNumber.MatchString(this.PhoneNumber) {
		return github_com_mwitkow_go_proto_validators.FieldError("PhoneNumber", fmt.Errorf(`value '%v' must be a string conforming to regex "^[+]?[0-9]{10,13}$"`, this.PhoneNumber))
	}
	if !_regex_ContactData_Email.MatchString(this.Email) {
		return github_com_mwitkow_go_proto_validators.FieldError("Email", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-z0-9._%+-]+@[a-z0-9.-]+.[a-z]{2,4}$"`, this.Email))
	}
	return nil
}

var _regex_GetDoctorRequestData_DoctorId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)

func (this *GetDoctorRequestData) Validate() error {
	if !_regex_GetDoctorRequestData_DoctorId.MatchString(this.DoctorId) {
		return github_com_mwitkow_go_proto_validators.FieldError("DoctorId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.DoctorId))
	}
	if this.DoctorId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("DoctorId", fmt.Errorf(`value '%v' must not be an empty string`, this.DoctorId))
	}
	return nil
}

var _regex_FindDoctorsRequestData_AreaId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)
var _regex_FindDoctorsRequestData_SpecialityId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)
var _regex_FindDoctorsRequestData_HospitalId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)

func (this *FindDoctorsRequestData) Validate() error {
	if !_regex_FindDoctorsRequestData_AreaId.MatchString(this.AreaId) {
		return github_com_mwitkow_go_proto_validators.FieldError("AreaId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.AreaId))
	}
	if !_regex_FindDoctorsRequestData_SpecialityId.MatchString(this.SpecialityId) {
		return github_com_mwitkow_go_proto_validators.FieldError("SpecialityId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.SpecialityId))
	}
	if !_regex_FindDoctorsRequestData_HospitalId.MatchString(this.HospitalId) {
		return github_com_mwitkow_go_proto_validators.FieldError("HospitalId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.HospitalId))
	}
	return nil
}

var _regex_GetHospitalRequestData_HospitalId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)

func (this *GetHospitalRequestData) Validate() error {
	if !_regex_GetHospitalRequestData_HospitalId.MatchString(this.HospitalId) {
		return github_com_mwitkow_go_proto_validators.FieldError("HospitalId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.HospitalId))
	}
	if this.HospitalId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("HospitalId", fmt.Errorf(`value '%v' must not be an empty string`, this.HospitalId))
	}
	return nil
}

var _regex_FindHospitalsRequestData_AreaId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)
var _regex_FindHospitalsRequestData_SpecialityId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)

func (this *FindHospitalsRequestData) Validate() error {
	if !_regex_FindHospitalsRequestData_AreaId.MatchString(this.AreaId) {
		return github_com_mwitkow_go_proto_validators.FieldError("AreaId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.AreaId))
	}
	if !_regex_FindHospitalsRequestData_SpecialityId.MatchString(this.SpecialityId) {
		return github_com_mwitkow_go_proto_validators.FieldError("SpecialityId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.SpecialityId))
	}
	if this.Coordinate != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Coordinate); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Coordinate", err)
		}
	}
	return nil
}

var _regex_GetAreasRequestData_AreaId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)

func (this *GetAreasRequestData) Validate() error {
	if !_regex_GetAreasRequestData_AreaId.MatchString(this.AreaId) {
		return github_com_mwitkow_go_proto_validators.FieldError("AreaId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.AreaId))
	}
	return nil
}

var _regex_GetSpecialitiesRequestData_SpecialityId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)

func (this *GetSpecialitiesRequestData) Validate() error {
	if !_regex_GetSpecialitiesRequestData_SpecialityId.MatchString(this.SpecialityId) {
		return github_com_mwitkow_go_proto_validators.FieldError("SpecialityId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.SpecialityId))
	}
	return nil
}

var _regex_FindSpecialitiesRequestData_AreaId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)
var _regex_FindSpecialitiesRequestData_HospitalId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)

func (this *FindSpecialitiesRequestData) Validate() error {
	if !_regex_FindSpecialitiesRequestData_AreaId.MatchString(this.AreaId) {
		return github_com_mwitkow_go_proto_validators.FieldError("AreaId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.AreaId))
	}
	if !_regex_FindSpecialitiesRequestData_HospitalId.MatchString(this.HospitalId) {
		return github_com_mwitkow_go_proto_validators.FieldError("HospitalId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.HospitalId))
	}
	return nil
}
func (this *HospitalsResponseData) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *DoctorsResponseData) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *SpecialitiesResponseData) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *AreasResponseData) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
