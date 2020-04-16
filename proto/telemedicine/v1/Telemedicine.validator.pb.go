// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/telemedicine/v1/telemedicine.proto

package prixa_telemedicine_v1

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "github.com/mwitkow/go-proto-validators"
	regexp "regexp"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *GetAppointmentBookingsRequestData) Validate() error {
	if nil == this.Param {
		return github_com_mwitkow_go_proto_validators.FieldError("Param", fmt.Errorf("message must exist"))
	}
	if this.Param != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Param); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Param", err)
		}
	}
	if nil == this.FromDate {
		return github_com_mwitkow_go_proto_validators.FieldError("FromDate", fmt.Errorf("message must exist"))
	}
	if this.FromDate != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.FromDate); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("FromDate", err)
		}
	}
	if nil == this.ToDate {
		return github_com_mwitkow_go_proto_validators.FieldError("ToDate", fmt.Errorf("message must exist"))
	}
	if this.ToDate != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ToDate); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ToDate", err)
		}
	}
	return nil
}

var _regex_FindAppointmentBookingParamData_HospitalId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)
var _regex_FindAppointmentBookingParamData_DoctorId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)

func (this *FindAppointmentBookingParamData) Validate() error {
	if !_regex_FindAppointmentBookingParamData_HospitalId.MatchString(this.HospitalId) {
		return github_com_mwitkow_go_proto_validators.FieldError("HospitalId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.HospitalId))
	}
	if !_regex_FindAppointmentBookingParamData_DoctorId.MatchString(this.DoctorId) {
		return github_com_mwitkow_go_proto_validators.FieldError("DoctorId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.DoctorId))
	}
	return nil
}

var _regex_GetAppointmentBookingRequestData_BookingId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)

func (this *GetAppointmentBookingRequestData) Validate() error {
	if !_regex_GetAppointmentBookingRequestData_BookingId.MatchString(this.BookingId) {
		return github_com_mwitkow_go_proto_validators.FieldError("BookingId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.BookingId))
	}
	if this.BookingId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("BookingId", fmt.Errorf(`value '%v' must not be an empty string`, this.BookingId))
	}
	return nil
}

var _regex_PostAppointmentBookingRequestData_Nik = regexp.MustCompile(`^[0-9]{16}$`)
var _regex_PostAppointmentBookingRequestData_Email = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-z0-9.-]+.[a-z]{2,4}$`)
var _regex_PostAppointmentBookingRequestData_PhoneNumber = regexp.MustCompile(`^[+]?[0-9]{10,13}$`)
var _regex_PostAppointmentBookingRequestData_AppointmentDate = regexp.MustCompile(`^\d{4}\-(0[1-9]|1[012])\-(0[1-9]|[12][0-9]|3[01])$`)
var _regex_PostAppointmentBookingRequestData_AppointmentTime = regexp.MustCompile(`^([0-1]?[0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9]$`)
var _regex_PostAppointmentBookingRequestData_DoctorId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)
var _regex_PostAppointmentBookingRequestData_HospitalId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)
var _regex_PostAppointmentBookingRequestData_ScheduleId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)
var _regex_PostAppointmentBookingRequestData_BirthDate = regexp.MustCompile(`^\d{4}\-(0[1-9]|1[012])\-(0[1-9]|[12][0-9]|3[01])$`)

func (this *PostAppointmentBookingRequestData) Validate() error {
	if this.Name == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must not be an empty string`, this.Name))
	}
	if !_regex_PostAppointmentBookingRequestData_Nik.MatchString(this.Nik) {
		return github_com_mwitkow_go_proto_validators.FieldError("Nik", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9]{16}$"`, this.Nik))
	}
	if this.Nik == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Nik", fmt.Errorf(`value '%v' must not be an empty string`, this.Nik))
	}
	if !_regex_PostAppointmentBookingRequestData_Email.MatchString(this.Email) {
		return github_com_mwitkow_go_proto_validators.FieldError("Email", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-zA-Z0-9._%+-]+@[a-z0-9.-]+.[a-z]{2,4}$"`, this.Email))
	}
	if this.Email == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Email", fmt.Errorf(`value '%v' must not be an empty string`, this.Email))
	}
	if !_regex_PostAppointmentBookingRequestData_PhoneNumber.MatchString(this.PhoneNumber) {
		return github_com_mwitkow_go_proto_validators.FieldError("PhoneNumber", fmt.Errorf(`value '%v' must be a string conforming to regex "^[+]?[0-9]{10,13}$"`, this.PhoneNumber))
	}
	if this.PhoneNumber == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("PhoneNumber", fmt.Errorf(`value '%v' must not be an empty string`, this.PhoneNumber))
	}
	if this.PaymentMethod == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("PaymentMethod", fmt.Errorf(`value '%v' must not be an empty string`, this.PaymentMethod))
	}
	if !_regex_PostAppointmentBookingRequestData_AppointmentDate.MatchString(this.AppointmentDate) {
		return github_com_mwitkow_go_proto_validators.FieldError("AppointmentDate", fmt.Errorf(`value '%v' must be a string conforming to regex "^\\d{4}\\-(0[1-9]|1[012])\\-(0[1-9]|[12][0-9]|3[01])$"`, this.AppointmentDate))
	}
	if this.AppointmentDate == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("AppointmentDate", fmt.Errorf(`value '%v' must not be an empty string`, this.AppointmentDate))
	}
	if !_regex_PostAppointmentBookingRequestData_AppointmentTime.MatchString(this.AppointmentTime) {
		return github_com_mwitkow_go_proto_validators.FieldError("AppointmentTime", fmt.Errorf(`value '%v' must be a string conforming to regex "^([0-1]?[0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9]$"`, this.AppointmentTime))
	}
	if this.AppointmentTime == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("AppointmentTime", fmt.Errorf(`value '%v' must not be an empty string`, this.AppointmentTime))
	}
	if !_regex_PostAppointmentBookingRequestData_DoctorId.MatchString(this.DoctorId) {
		return github_com_mwitkow_go_proto_validators.FieldError("DoctorId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.DoctorId))
	}
	if this.DoctorId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("DoctorId", fmt.Errorf(`value '%v' must not be an empty string`, this.DoctorId))
	}
	if !_regex_PostAppointmentBookingRequestData_HospitalId.MatchString(this.HospitalId) {
		return github_com_mwitkow_go_proto_validators.FieldError("HospitalId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.HospitalId))
	}
	if this.HospitalId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("HospitalId", fmt.Errorf(`value '%v' must not be an empty string`, this.HospitalId))
	}
	if !_regex_PostAppointmentBookingRequestData_ScheduleId.MatchString(this.ScheduleId) {
		return github_com_mwitkow_go_proto_validators.FieldError("ScheduleId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.ScheduleId))
	}
	if this.ScheduleId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ScheduleId", fmt.Errorf(`value '%v' must not be an empty string`, this.ScheduleId))
	}
	if !_regex_PostAppointmentBookingRequestData_BirthDate.MatchString(this.BirthDate) {
		return github_com_mwitkow_go_proto_validators.FieldError("BirthDate", fmt.Errorf(`value '%v' must be a string conforming to regex "^\\d{4}\\-(0[1-9]|1[012])\\-(0[1-9]|[12][0-9]|3[01])$"`, this.BirthDate))
	}
	if this.BirthDate == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("BirthDate", fmt.Errorf(`value '%v' must not be an empty string`, this.BirthDate))
	}
	return nil
}
func (this *AppointmentBookingsResponseData) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *PaymentMethodResponseData) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *PaymentMethodData) Validate() error {
	if this.Name == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must not be an empty string`, this.Name))
	}
	return nil
}
func (this *AppointmentBookingResponseData) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}

var _regex_AppointmentBookingData_Id = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)
var _regex_AppointmentBookingData_AppointmentDate = regexp.MustCompile(`^\d{4}\-(0[1-9]|1[012])\-(0[1-9]|[12][0-9]|3[01])$`)
var _regex_AppointmentBookingData_AppointmentTime = regexp.MustCompile(`^([0-1]?[0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9]$`)

func (this *AppointmentBookingData) Validate() error {
	if !_regex_AppointmentBookingData_Id.MatchString(this.Id) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.Id))
	}
	if this.Id == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must not be an empty string`, this.Id))
	}
	if !_regex_AppointmentBookingData_AppointmentDate.MatchString(this.AppointmentDate) {
		return github_com_mwitkow_go_proto_validators.FieldError("AppointmentDate", fmt.Errorf(`value '%v' must be a string conforming to regex "^\\d{4}\\-(0[1-9]|1[012])\\-(0[1-9]|[12][0-9]|3[01])$"`, this.AppointmentDate))
	}
	if this.AppointmentDate == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("AppointmentDate", fmt.Errorf(`value '%v' must not be an empty string`, this.AppointmentDate))
	}
	if !_regex_AppointmentBookingData_AppointmentTime.MatchString(this.AppointmentTime) {
		return github_com_mwitkow_go_proto_validators.FieldError("AppointmentTime", fmt.Errorf(`value '%v' must be a string conforming to regex "^([0-1]?[0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9]$"`, this.AppointmentTime))
	}
	if this.AppointmentTime == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("AppointmentTime", fmt.Errorf(`value '%v' must not be an empty string`, this.AppointmentTime))
	}
	if this.Doctor != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Doctor); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Doctor", err)
		}
	}
	if this.Hospital != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Hospital); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Hospital", err)
		}
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
func (this *GetHospitalsRequestData) Validate() error {
	if this.Param != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Param); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Param", err)
		}
	}
	return nil
}
func (this *GetHospitalsByCoordinateRequestData) Validate() error {
	return nil
}
func (this *HospitalsByCoordinateResponseData) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *HospitalByCoordinateData) Validate() error {
	if this.Hospital != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Hospital); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Hospital", err)
		}
	}
	return nil
}

var _regex_GetHospitalParamData_AreaId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)

func (this *GetHospitalParamData) Validate() error {
	if !_regex_GetHospitalParamData_AreaId.MatchString(this.AreaId) {
		return github_com_mwitkow_go_proto_validators.FieldError("AreaId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.AreaId))
	}
	return nil
}

var _regex_HospitalData_Id = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)
var _regex_HospitalData_AreaId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)

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
	if !_regex_HospitalData_AreaId.MatchString(this.AreaId) {
		return github_com_mwitkow_go_proto_validators.FieldError("AreaId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.AreaId))
	}
	for _, item := range this.Specialities {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Specialities", err)
			}
		}
	}
	if this.Source != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Source); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Source", err)
		}
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
func (this *HospitalResponseData) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
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
func (this *GetDoctorsRequestData) Validate() error {
	if this.Param != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Param); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Param", err)
		}
	}
	return nil
}

var _regex_GetDoctorParamData_AreaId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)

func (this *GetDoctorParamData) Validate() error {
	if !_regex_GetDoctorParamData_AreaId.MatchString(this.AreaId) {
		return github_com_mwitkow_go_proto_validators.FieldError("AreaId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.AreaId))
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
func (this *DoctorResponseData) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
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
	for _, item := range this.Hospitals {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Hospitals", err)
			}
		}
	}
	if this.Source != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Source); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Source", err)
		}
	}
	return nil
}

var _regex_GetSpecialityRequestData_SpecialityId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)

func (this *GetSpecialityRequestData) Validate() error {
	if !_regex_GetSpecialityRequestData_SpecialityId.MatchString(this.SpecialityId) {
		return github_com_mwitkow_go_proto_validators.FieldError("SpecialityId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.SpecialityId))
	}
	if this.SpecialityId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("SpecialityId", fmt.Errorf(`value '%v' must not be an empty string`, this.SpecialityId))
	}
	return nil
}

var _regex_GetSpecialitiesRequestData_AreaId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)
var _regex_GetSpecialitiesRequestData_HospitalId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)

func (this *GetSpecialitiesRequestData) Validate() error {
	if !_regex_GetSpecialitiesRequestData_AreaId.MatchString(this.AreaId) {
		return github_com_mwitkow_go_proto_validators.FieldError("AreaId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.AreaId))
	}
	if !_regex_GetSpecialitiesRequestData_HospitalId.MatchString(this.HospitalId) {
		return github_com_mwitkow_go_proto_validators.FieldError("HospitalId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.HospitalId))
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
	if this.Source != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Source); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Source", err)
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
func (this *SpecialityResponseData) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}

var _regex_GetAreaRequestData_AreaId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)

func (this *GetAreaRequestData) Validate() error {
	if !_regex_GetAreaRequestData_AreaId.MatchString(this.AreaId) {
		return github_com_mwitkow_go_proto_validators.FieldError("AreaId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.AreaId))
	}
	if this.AreaId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("AreaId", fmt.Errorf(`value '%v' must not be an empty string`, this.AreaId))
	}
	return nil
}
func (this *GetAreasRequestData) Validate() error {
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
func (this *AreaResponseData) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
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
	if this.Source != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Source); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Source", err)
		}
	}
	return nil
}

var _regex_ScheduleData_ScheduleId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)

func (this *ScheduleData) Validate() error {
	if !_regex_ScheduleData_ScheduleId.MatchString(this.ScheduleId) {
		return github_com_mwitkow_go_proto_validators.FieldError("ScheduleId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.ScheduleId))
	}
	if this.Doctor != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Doctor); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Doctor", err)
		}
	}
	if this.Hospital != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Hospital); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Hospital", err)
		}
	}
	if this.Source != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Source); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Source", err)
		}
	}
	return nil
}
func (this *DoctorSchedulesResponseData) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *DoctorScheduleData) Validate() error {
	for _, item := range this.Schedules {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Schedules", err)
			}
		}
	}
	for _, item := range this.Leaves {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Leaves", err)
			}
		}
	}
	return nil
}
func (this *DoctorLeaveHospitalData) Validate() error {
	return nil
}

var _regex_GetDoctorSchedulesRequestData_DoctorId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)
var _regex_GetDoctorSchedulesRequestData_HospitalId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)

func (this *GetDoctorSchedulesRequestData) Validate() error {
	if !_regex_GetDoctorSchedulesRequestData_DoctorId.MatchString(this.DoctorId) {
		return github_com_mwitkow_go_proto_validators.FieldError("DoctorId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.DoctorId))
	}
	if this.DoctorId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("DoctorId", fmt.Errorf(`value '%v' must not be an empty string`, this.DoctorId))
	}
	if !_regex_GetDoctorSchedulesRequestData_HospitalId.MatchString(this.HospitalId) {
		return github_com_mwitkow_go_proto_validators.FieldError("HospitalId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.HospitalId))
	}
	if this.HospitalId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("HospitalId", fmt.Errorf(`value '%v' must not be an empty string`, this.HospitalId))
	}
	return nil
}

var _regex_GetDoctorTimeSlotRequestData_HospitalId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)
var _regex_GetDoctorTimeSlotRequestData_DoctorId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)
var _regex_GetDoctorTimeSlotRequestData_AppointmentDate = regexp.MustCompile(`^\d{4}\-(0[1-9]|1[012])\-(0[1-9]|[12][0-9]|3[01])$`)

func (this *GetDoctorTimeSlotRequestData) Validate() error {
	if !_regex_GetDoctorTimeSlotRequestData_HospitalId.MatchString(this.HospitalId) {
		return github_com_mwitkow_go_proto_validators.FieldError("HospitalId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.HospitalId))
	}
	if this.HospitalId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("HospitalId", fmt.Errorf(`value '%v' must not be an empty string`, this.HospitalId))
	}
	if !_regex_GetDoctorTimeSlotRequestData_DoctorId.MatchString(this.DoctorId) {
		return github_com_mwitkow_go_proto_validators.FieldError("DoctorId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.DoctorId))
	}
	if this.DoctorId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("DoctorId", fmt.Errorf(`value '%v' must not be an empty string`, this.DoctorId))
	}
	if !_regex_GetDoctorTimeSlotRequestData_AppointmentDate.MatchString(this.AppointmentDate) {
		return github_com_mwitkow_go_proto_validators.FieldError("AppointmentDate", fmt.Errorf(`value '%v' must be a string conforming to regex "^\\d{4}\\-(0[1-9]|1[012])\\-(0[1-9]|[12][0-9]|3[01])$"`, this.AppointmentDate))
	}
	if this.AppointmentDate == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("AppointmentDate", fmt.Errorf(`value '%v' must not be an empty string`, this.AppointmentDate))
	}
	return nil
}
func (this *DoctorTimeSlotsResponseData) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *DoctorTimeSlotData) Validate() error {
	if this.FromTime == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("FromTime", fmt.Errorf(`value '%v' must not be an empty string`, this.FromTime))
	}
	if this.ToTime == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ToTime", fmt.Errorf(`value '%v' must not be an empty string`, this.ToTime))
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
var _regex_ContactData_Email = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-z0-9.-]+.[a-z]{2,4}$`)

func (this *ContactData) Validate() error {
	if !_regex_ContactData_PhoneNumber.MatchString(this.PhoneNumber) {
		return github_com_mwitkow_go_proto_validators.FieldError("PhoneNumber", fmt.Errorf(`value '%v' must be a string conforming to regex "^[+]?[0-9]{10,13}$"`, this.PhoneNumber))
	}
	if !_regex_ContactData_Email.MatchString(this.Email) {
		return github_com_mwitkow_go_proto_validators.FieldError("Email", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-zA-Z0-9._%+-]+@[a-z0-9.-]+.[a-z]{2,4}$"`, this.Email))
	}
	return nil
}
func (this *InitConversationRequest) Validate() error {
	if this.DiagnosticSessionID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("DiagnosticSessionID", fmt.Errorf(`value '%v' must not be an empty string`, this.DiagnosticSessionID))
	}
	if this.SessionID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("SessionID", fmt.Errorf(`value '%v' must not be an empty string`, this.SessionID))
	}
	return nil
}
func (this *InitConversationResponse) Validate() error {
	if this.WebsiteID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("WebsiteID", fmt.Errorf(`value '%v' must not be an empty string`, this.WebsiteID))
	}
	return nil
}
func (this *CancelAppointmentBookingRequestData) Validate() error {
	if this.BookingId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("BookingId", fmt.Errorf(`value '%v' must not be an empty string`, this.BookingId))
	}
	if this.Reason == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Reason", fmt.Errorf(`value '%v' must not be an empty string`, this.Reason))
	}
	return nil
}
func (this *CancelAppointmentBookingResponseData) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}

var _regex_CancelAppointmentBookingData_BookingId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)

func (this *CancelAppointmentBookingData) Validate() error {
	if !_regex_CancelAppointmentBookingData_BookingId.MatchString(this.BookingId) {
		return github_com_mwitkow_go_proto_validators.FieldError("BookingId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.BookingId))
	}
	if this.BookingId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("BookingId", fmt.Errorf(`value '%v' must not be an empty string`, this.BookingId))
	}
	return nil
}
