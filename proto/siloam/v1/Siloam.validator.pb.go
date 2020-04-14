// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/Siloam/v1/Siloam.proto

package prixa_siloam_v1

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/mwitkow/go-proto-validators"
	_ "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/golang/protobuf/ptypes/wrappers"
	regexp "regexp"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *GetAreasResponseData) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *AreaData) Validate() error {
	return nil
}

var _regex_GetHospitalsRequestData_AreaId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)

func (this *GetHospitalsRequestData) Validate() error {
	if this.OnImplementation != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.OnImplementation); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("OnImplementation", err)
		}
	}
	if this.AreaId != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.AreaId); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("AreaId", err)
		}
	}
	return nil
}
func (this *GetHospitalsResponseData) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *HospitalData) Validate() error {
	return nil
}

var _regex_GetSpecialitiesRequestData_HospitalId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)

func (this *GetSpecialitiesRequestData) Validate() error {
	if this.HospitalId != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.HospitalId); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("HospitalId", err)
		}
	}
	return nil
}
func (this *GetSpecialitiesResponseData) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *SpecialityData) Validate() error {
	return nil
}

var _regex_GetDoctorsRequestData_DoctorId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)
var _regex_GetDoctorsRequestData_HospitalId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)
var _regex_GetDoctorsRequestData_SpecialityId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)

func (this *GetDoctorsRequestData) Validate() error {
	if this.DoctorName != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.DoctorName); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("DoctorName", err)
		}
	}
	if this.DoctorId != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.DoctorId); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("DoctorId", err)
		}
	}
	if this.HospitalId != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.HospitalId); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("HospitalId", err)
		}
	}
	if this.SpecialityId != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.SpecialityId); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("SpecialityId", err)
		}
	}
	return nil
}
func (this *GetDoctorsResponseData) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *DoctorData) Validate() error {
	return nil
}

var _regex_GetDoctorsLeavesHospitalRequestData_HospitalId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)
var _regex_GetDoctorsLeavesHospitalRequestData_DoctorId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)

func (this *GetDoctorsLeavesHospitalRequestData) Validate() error {
	if !_regex_GetDoctorsLeavesHospitalRequestData_HospitalId.MatchString(this.HospitalId) {
		return github_com_mwitkow_go_proto_validators.FieldError("HospitalId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.HospitalId))
	}
	if this.HospitalId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("HospitalId", fmt.Errorf(`value '%v' must not be an empty string`, this.HospitalId))
	}
	if this.DoctorId != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.DoctorId); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("DoctorId", err)
		}
	}
	return nil
}
func (this *GetDoctorsLeavesHospitalResponseData) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *DoctorLeaveHospitalData) Validate() error {
	return nil
}

var _regex_GetSchedulesPerWeeksRequestData_HospitalId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)
var _regex_GetSchedulesPerWeeksRequestData_DoctorId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)

func (this *GetSchedulesPerWeeksRequestData) Validate() error {
	if !_regex_GetSchedulesPerWeeksRequestData_HospitalId.MatchString(this.HospitalId) {
		return github_com_mwitkow_go_proto_validators.FieldError("HospitalId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.HospitalId))
	}
	if this.HospitalId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("HospitalId", fmt.Errorf(`value '%v' must not be an empty string`, this.HospitalId))
	}
	if !_regex_GetSchedulesPerWeeksRequestData_DoctorId.MatchString(this.DoctorId) {
		return github_com_mwitkow_go_proto_validators.FieldError("DoctorId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.DoctorId))
	}
	if this.DoctorId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("DoctorId", fmt.Errorf(`value '%v' must not be an empty string`, this.DoctorId))
	}
	return nil
}
func (this *GetSchedulesPerWeeksResponseData) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *SchedulesPerWeeksData) Validate() error {
	if this.FromTime != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.FromTime); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("FromTime", err)
		}
	}
	if this.ToTime != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ToTime); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ToTime", err)
		}
	}
	if this.FromDate != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.FromDate); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("FromDate", err)
		}
	}
	if this.ToDate != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ToDate); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ToDate", err)
		}
	}
	return nil
}

var _regex_GetTimeSlotsRequestData_HospitalId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)
var _regex_GetTimeSlotsRequestData_DoctorId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)
var _regex_GetTimeSlotsRequestData_AppointmentDate = regexp.MustCompile(`^\d{4}\-(0[1-9]|1[012])\-(0[1-9]|[12][0-9]|3[01])$`)

func (this *GetTimeSlotsRequestData) Validate() error {
	if !_regex_GetTimeSlotsRequestData_HospitalId.MatchString(this.HospitalId) {
		return github_com_mwitkow_go_proto_validators.FieldError("HospitalId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.HospitalId))
	}
	if this.HospitalId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("HospitalId", fmt.Errorf(`value '%v' must not be an empty string`, this.HospitalId))
	}
	if !_regex_GetTimeSlotsRequestData_DoctorId.MatchString(this.DoctorId) {
		return github_com_mwitkow_go_proto_validators.FieldError("DoctorId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.DoctorId))
	}
	if this.DoctorId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("DoctorId", fmt.Errorf(`value '%v' must not be an empty string`, this.DoctorId))
	}
	if !_regex_GetTimeSlotsRequestData_AppointmentDate.MatchString(this.AppointmentDate) {
		return github_com_mwitkow_go_proto_validators.FieldError("AppointmentDate", fmt.Errorf(`value '%v' must be a string conforming to regex "^\\d{4}\\-(0[1-9]|1[012])\\-(0[1-9]|[12][0-9]|3[01])$"`, this.AppointmentDate))
	}
	if this.AppointmentDate == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("AppointmentDate", fmt.Errorf(`value '%v' must not be an empty string`, this.AppointmentDate))
	}
	return nil
}
func (this *GetTimeSlotsResponseData) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *TimeSlotData) Validate() error {
	return nil
}

var _regex_PostAppointmentRequestData_AppointmentDate = regexp.MustCompile(`^\d{4}\-(0[1-9]|1[012])\-(0[1-9]|[12][0-9]|3[01])$`)
var _regex_PostAppointmentRequestData_AppointmentFromTime = regexp.MustCompile(`^([0-1]?[0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9]$`)
var _regex_PostAppointmentRequestData_ScheduleId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)
var _regex_PostAppointmentRequestData_HospitalId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)
var _regex_PostAppointmentRequestData_DoctorId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)
var _regex_PostAppointmentRequestData_BirthDate = regexp.MustCompile(`^\d{4}\-(0[1-9]|1[012])\-(0[1-9]|[12][0-9]|3[01])$`)
var _regex_PostAppointmentRequestData_PhoneNumber1 = regexp.MustCompile(`^[+]?[0-9]{10,13}$`)
var _regex_PostAppointmentRequestData_EmailAddress = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-z0-9.-]+.[a-z]{2,4}$`)

func (this *PostAppointmentRequestData) Validate() error {
	if this.ChannelId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ChannelId", fmt.Errorf(`value '%v' must not be an empty string`, this.ChannelId))
	}
	if !_regex_PostAppointmentRequestData_AppointmentDate.MatchString(this.AppointmentDate) {
		return github_com_mwitkow_go_proto_validators.FieldError("AppointmentDate", fmt.Errorf(`value '%v' must be a string conforming to regex "^\\d{4}\\-(0[1-9]|1[012])\\-(0[1-9]|[12][0-9]|3[01])$"`, this.AppointmentDate))
	}
	if this.AppointmentDate == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("AppointmentDate", fmt.Errorf(`value '%v' must not be an empty string`, this.AppointmentDate))
	}
	if !_regex_PostAppointmentRequestData_AppointmentFromTime.MatchString(this.AppointmentFromTime) {
		return github_com_mwitkow_go_proto_validators.FieldError("AppointmentFromTime", fmt.Errorf(`value '%v' must be a string conforming to regex "^([0-1]?[0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9]$"`, this.AppointmentFromTime))
	}
	if this.AppointmentFromTime == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("AppointmentFromTime", fmt.Errorf(`value '%v' must not be an empty string`, this.AppointmentFromTime))
	}
	if !_regex_PostAppointmentRequestData_ScheduleId.MatchString(this.ScheduleId) {
		return github_com_mwitkow_go_proto_validators.FieldError("ScheduleId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.ScheduleId))
	}
	if this.ScheduleId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ScheduleId", fmt.Errorf(`value '%v' must not be an empty string`, this.ScheduleId))
	}
	if !_regex_PostAppointmentRequestData_HospitalId.MatchString(this.HospitalId) {
		return github_com_mwitkow_go_proto_validators.FieldError("HospitalId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.HospitalId))
	}
	if this.HospitalId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("HospitalId", fmt.Errorf(`value '%v' must not be an empty string`, this.HospitalId))
	}
	if !_regex_PostAppointmentRequestData_DoctorId.MatchString(this.DoctorId) {
		return github_com_mwitkow_go_proto_validators.FieldError("DoctorId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.DoctorId))
	}
	if this.DoctorId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("DoctorId", fmt.Errorf(`value '%v' must not be an empty string`, this.DoctorId))
	}
	if this.Name == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must not be an empty string`, this.Name))
	}
	if !_regex_PostAppointmentRequestData_BirthDate.MatchString(this.BirthDate) {
		return github_com_mwitkow_go_proto_validators.FieldError("BirthDate", fmt.Errorf(`value '%v' must be a string conforming to regex "^\\d{4}\\-(0[1-9]|1[012])\\-(0[1-9]|[12][0-9]|3[01])$"`, this.BirthDate))
	}
	if this.BirthDate == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("BirthDate", fmt.Errorf(`value '%v' must not be an empty string`, this.BirthDate))
	}
	if !_regex_PostAppointmentRequestData_PhoneNumber1.MatchString(this.PhoneNumber1) {
		return github_com_mwitkow_go_proto_validators.FieldError("PhoneNumber1", fmt.Errorf(`value '%v' must be a string conforming to regex "^[+]?[0-9]{10,13}$"`, this.PhoneNumber1))
	}
	if this.PhoneNumber1 == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("PhoneNumber1", fmt.Errorf(`value '%v' must not be an empty string`, this.PhoneNumber1))
	}
	if !_regex_PostAppointmentRequestData_EmailAddress.MatchString(this.EmailAddress) {
		return github_com_mwitkow_go_proto_validators.FieldError("EmailAddress", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-zA-Z0-9._%+-]+@[a-z0-9.-]+.[a-z]{2,4}$"`, this.EmailAddress))
	}
	if this.UserId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("UserId", fmt.Errorf(`value '%v' must not be an empty string`, this.UserId))
	}
	return nil
}
func (this *PostAppointmentResponseData) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *PostAppointmentData) Validate() error {
	return nil
}

var _regex_CancelAppointmentRequestData_UserId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)

func (this *CancelAppointmentRequestData) Validate() error {
	if this.AppointmentId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("AppointmentId", fmt.Errorf(`value '%v' must not be an empty string`, this.AppointmentId))
	}
	if !_regex_CancelAppointmentRequestData_UserId.MatchString(this.UserId) {
		return github_com_mwitkow_go_proto_validators.FieldError("UserId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.UserId))
	}
	if this.UserId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("UserId", fmt.Errorf(`value '%v' must not be an empty string`, this.UserId))
	}
	return nil
}
func (this *CancelAppointmentResponseData) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CancelAppointmentData) Validate() error {
	return nil
}

var _regex_PostRescheduleAppointmentRequestData_AppointmentId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)
var _regex_PostRescheduleAppointmentRequestData_ScheduleId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)
var _regex_PostRescheduleAppointmentRequestData_HospitalId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)
var _regex_PostRescheduleAppointmentRequestData_UserId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)

func (this *PostRescheduleAppointmentRequestData) Validate() error {
	if !_regex_PostRescheduleAppointmentRequestData_AppointmentId.MatchString(this.AppointmentId) {
		return github_com_mwitkow_go_proto_validators.FieldError("AppointmentId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.AppointmentId))
	}
	if this.AppointmentId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("AppointmentId", fmt.Errorf(`value '%v' must not be an empty string`, this.AppointmentId))
	}
	if !_regex_PostRescheduleAppointmentRequestData_ScheduleId.MatchString(this.ScheduleId) {
		return github_com_mwitkow_go_proto_validators.FieldError("ScheduleId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.ScheduleId))
	}
	if this.ScheduleId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ScheduleId", fmt.Errorf(`value '%v' must not be an empty string`, this.ScheduleId))
	}
	if nil == this.AppointmentDate {
		return github_com_mwitkow_go_proto_validators.FieldError("AppointmentDate", fmt.Errorf("message must exist"))
	}
	if this.AppointmentDate != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.AppointmentDate); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("AppointmentDate", err)
		}
	}
	if nil == this.AppointmentFromTime {
		return github_com_mwitkow_go_proto_validators.FieldError("AppointmentFromTime", fmt.Errorf("message must exist"))
	}
	if this.AppointmentFromTime != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.AppointmentFromTime); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("AppointmentFromTime", err)
		}
	}
	if !_regex_PostRescheduleAppointmentRequestData_HospitalId.MatchString(this.HospitalId) {
		return github_com_mwitkow_go_proto_validators.FieldError("HospitalId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.HospitalId))
	}
	if this.HospitalId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("HospitalId", fmt.Errorf(`value '%v' must not be an empty string`, this.HospitalId))
	}
	if !_regex_PostRescheduleAppointmentRequestData_UserId.MatchString(this.UserId) {
		return github_com_mwitkow_go_proto_validators.FieldError("UserId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.UserId))
	}
	if this.UserId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("UserId", fmt.Errorf(`value '%v' must not be an empty string`, this.UserId))
	}
	return nil
}
func (this *PostRescheduleAppointmentResponseData) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *RescheduleAppointmentData) Validate() error {
	if this.AppointmentDate != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.AppointmentDate); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("AppointmentDate", err)
		}
	}
	return nil
}

var _regex_GetAppointmentsRequestData_ContactId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)
var _regex_GetAppointmentsRequestData_HospitalId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)

func (this *GetAppointmentsRequestData) Validate() error {
	if !_regex_GetAppointmentsRequestData_ContactId.MatchString(this.ContactId) {
		return github_com_mwitkow_go_proto_validators.FieldError("ContactId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.ContactId))
	}
	if this.ContactId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ContactId", fmt.Errorf(`value '%v' must not be an empty string`, this.ContactId))
	}
	if !_regex_GetAppointmentsRequestData_HospitalId.MatchString(this.HospitalId) {
		return github_com_mwitkow_go_proto_validators.FieldError("HospitalId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.HospitalId))
	}
	if this.HospitalId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("HospitalId", fmt.Errorf(`value '%v' must not be an empty string`, this.HospitalId))
	}
	if this.FutureOnly != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.FutureOnly); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("FutureOnly", err)
		}
	}
	if this.Limit != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Limit); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Limit", err)
		}
	}
	if this.Offset != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Offset); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Offset", err)
		}
	}
	return nil
}
func (this *GetAppointmentsResponseData) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *GetAppointmentData) Validate() error {
	if this.AppointmentFromTime != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.AppointmentFromTime); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("AppointmentFromTime", err)
		}
	}
	if this.AppointmentToTime != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.AppointmentToTime); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("AppointmentToTime", err)
		}
	}
	if this.AppointmentDate != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.AppointmentDate); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("AppointmentDate", err)
		}
	}
	return nil
}
