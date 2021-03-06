// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/nalar/v1/nalar.proto

package nalar

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type RecordSymptomsData struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Fields               *Symptom `protobuf:"bytes,2,opt,name=fields,proto3" json:"fields,omitempty"`
	CreatedTime          string   `protobuf:"bytes,3,opt,name=createdTime,proto3" json:"createdTime,omitempty"`
	LastUpdated          string   `protobuf:"bytes,4,opt,name=lastUpdated,proto3" json:"lastUpdated,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecordSymptomsData) Reset()         { *m = RecordSymptomsData{} }
func (m *RecordSymptomsData) String() string { return proto.CompactTextString(m) }
func (*RecordSymptomsData) ProtoMessage()    {}
func (*RecordSymptomsData) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6abc415730e04a0, []int{0}
}

func (m *RecordSymptomsData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecordSymptomsData.Unmarshal(m, b)
}
func (m *RecordSymptomsData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecordSymptomsData.Marshal(b, m, deterministic)
}
func (m *RecordSymptomsData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecordSymptomsData.Merge(m, src)
}
func (m *RecordSymptomsData) XXX_Size() int {
	return xxx_messageInfo_RecordSymptomsData.Size(m)
}
func (m *RecordSymptomsData) XXX_DiscardUnknown() {
	xxx_messageInfo_RecordSymptomsData.DiscardUnknown(m)
}

var xxx_messageInfo_RecordSymptomsData proto.InternalMessageInfo

func (m *RecordSymptomsData) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *RecordSymptomsData) GetFields() *Symptom {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *RecordSymptomsData) GetCreatedTime() string {
	if m != nil {
		return m.CreatedTime
	}
	return ""
}

func (m *RecordSymptomsData) GetLastUpdated() string {
	if m != nil {
		return m.LastUpdated
	}
	return ""
}

type Symptom struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,json=id,proto3" json:"ID,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	NameIndo             string   `protobuf:"bytes,3,opt,name=NameIndo,json=Name Indo,proto3" json:"NameIndo,omitempty"`
	Description          string   `protobuf:"bytes,4,opt,name=Description,json=Symptoms Description,proto3" json:"Description,omitempty"`
	Explanation          string   `protobuf:"bytes,5,opt,name=Explanation,json=Symptoms Explanation,proto3" json:"Explanation,omitempty"`
	Question             string   `protobuf:"bytes,6,opt,name=Question,proto3" json:"Question,omitempty"`
	Characteristic       string   `protobuf:"bytes,7,opt,name=Characteristic,proto3" json:"Characteristic,omitempty"`
	MediaFilesLink       string   `protobuf:"bytes,8,opt,name=MediaFilesLink,proto3" json:"MediaFilesLink,omitempty"`
	SourceURL            string   `protobuf:"bytes,9,opt,name=SourceURL,proto3" json:"SourceURL,omitempty"`
	AssociatedDiseases   []string `protobuf:"bytes,10,rep,name=AssociatedDiseases,json=Diseases,proto3" json:"AssociatedDiseases,omitempty"`
	Keywords             string   `protobuf:"bytes,11,opt,name=Keywords,json=Keyword,proto3" json:"Keywords,omitempty"`
	Differentials        []string `protobuf:"bytes,12,rep,name=Differentials,json=Differential,proto3" json:"Differentials,omitempty"`
	Triage               string   `protobuf:"bytes,13,opt,name=Triage,proto3" json:"Triage,omitempty"`
	NiceToKnows          []string `protobuf:"bytes,14,rep,name=NiceToKnows,json=Nice To Know,proto3" json:"NiceToKnows,omitempty"`
	MustKnows            []string `protobuf:"bytes,15,rep,name=MustKnows,json=Must Know,proto3" json:"MustKnows,omitempty"`
	PreConditions        []string `protobuf:"bytes,16,rep,name=PreConditions,json=Precondition,proto3" json:"PreConditions,omitempty"`
	SymptomsTriage       []string `protobuf:"bytes,17,rep,name=SymptomsTriage,json=Symptoms-Triage,proto3" json:"SymptomsTriage,omitempty"`
	SID                  string   `protobuf:"bytes,18,opt,name=SID,json=Symptom_ID,proto3" json:"SID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Symptom) Reset()         { *m = Symptom{} }
func (m *Symptom) String() string { return proto.CompactTextString(m) }
func (*Symptom) ProtoMessage()    {}
func (*Symptom) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6abc415730e04a0, []int{1}
}

func (m *Symptom) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Symptom.Unmarshal(m, b)
}
func (m *Symptom) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Symptom.Marshal(b, m, deterministic)
}
func (m *Symptom) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Symptom.Merge(m, src)
}
func (m *Symptom) XXX_Size() int {
	return xxx_messageInfo_Symptom.Size(m)
}
func (m *Symptom) XXX_DiscardUnknown() {
	xxx_messageInfo_Symptom.DiscardUnknown(m)
}

var xxx_messageInfo_Symptom proto.InternalMessageInfo

func (m *Symptom) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Symptom) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Symptom) GetNameIndo() string {
	if m != nil {
		return m.NameIndo
	}
	return ""
}

func (m *Symptom) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Symptom) GetExplanation() string {
	if m != nil {
		return m.Explanation
	}
	return ""
}

func (m *Symptom) GetQuestion() string {
	if m != nil {
		return m.Question
	}
	return ""
}

func (m *Symptom) GetCharacteristic() string {
	if m != nil {
		return m.Characteristic
	}
	return ""
}

func (m *Symptom) GetMediaFilesLink() string {
	if m != nil {
		return m.MediaFilesLink
	}
	return ""
}

func (m *Symptom) GetSourceURL() string {
	if m != nil {
		return m.SourceURL
	}
	return ""
}

func (m *Symptom) GetAssociatedDiseases() []string {
	if m != nil {
		return m.AssociatedDiseases
	}
	return nil
}

func (m *Symptom) GetKeywords() string {
	if m != nil {
		return m.Keywords
	}
	return ""
}

func (m *Symptom) GetDifferentials() []string {
	if m != nil {
		return m.Differentials
	}
	return nil
}

func (m *Symptom) GetTriage() string {
	if m != nil {
		return m.Triage
	}
	return ""
}

func (m *Symptom) GetNiceToKnows() []string {
	if m != nil {
		return m.NiceToKnows
	}
	return nil
}

func (m *Symptom) GetMustKnows() []string {
	if m != nil {
		return m.MustKnows
	}
	return nil
}

func (m *Symptom) GetPreConditions() []string {
	if m != nil {
		return m.PreConditions
	}
	return nil
}

func (m *Symptom) GetSymptomsTriage() []string {
	if m != nil {
		return m.SymptomsTriage
	}
	return nil
}

func (m *Symptom) GetSID() string {
	if m != nil {
		return m.SID
	}
	return ""
}

type RecordDiseaseData struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Fields               *Disease `protobuf:"bytes,2,opt,name=fields,proto3" json:"fields,omitempty"`
	CreatedTime          string   `protobuf:"bytes,3,opt,name=createdTime,proto3" json:"createdTime,omitempty"`
	LastUpdated          string   `protobuf:"bytes,4,opt,name=lastUpdated,proto3" json:"lastUpdated,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecordDiseaseData) Reset()         { *m = RecordDiseaseData{} }
func (m *RecordDiseaseData) String() string { return proto.CompactTextString(m) }
func (*RecordDiseaseData) ProtoMessage()    {}
func (*RecordDiseaseData) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6abc415730e04a0, []int{2}
}

func (m *RecordDiseaseData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecordDiseaseData.Unmarshal(m, b)
}
func (m *RecordDiseaseData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecordDiseaseData.Marshal(b, m, deterministic)
}
func (m *RecordDiseaseData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecordDiseaseData.Merge(m, src)
}
func (m *RecordDiseaseData) XXX_Size() int {
	return xxx_messageInfo_RecordDiseaseData.Size(m)
}
func (m *RecordDiseaseData) XXX_DiscardUnknown() {
	xxx_messageInfo_RecordDiseaseData.DiscardUnknown(m)
}

var xxx_messageInfo_RecordDiseaseData proto.InternalMessageInfo

func (m *RecordDiseaseData) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *RecordDiseaseData) GetFields() *Disease {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *RecordDiseaseData) GetCreatedTime() string {
	if m != nil {
		return m.CreatedTime
	}
	return ""
}

func (m *RecordDiseaseData) GetLastUpdated() string {
	if m != nil {
		return m.LastUpdated
	}
	return ""
}

type Disease struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,json=id,proto3" json:"ID,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	NameIndo             string   `protobuf:"bytes,3,opt,name=NameIndo,json=Name Indo,proto3" json:"NameIndo,omitempty"`
	Description          string   `protobuf:"bytes,4,opt,name=Description,proto3" json:"Description,omitempty"`
	AssociatedSymptoms   []string `protobuf:"bytes,5,rep,name=AssociatedSymptoms,json=Symptoms,proto3" json:"AssociatedSymptoms,omitempty"`
	Differentials        []string `protobuf:"bytes,6,rep,name=Differentials,json=Differential,proto3" json:"Differentials,omitempty"`
	RiskFactors          []string `protobuf:"bytes,7,rep,name=RiskFactors,json=Risk Factor,proto3" json:"RiskFactors,omitempty"`
	Likeliness           string   `protobuf:"bytes,8,opt,name=Likeliness,proto3" json:"Likeliness,omitempty"`
	Triage               int32    `protobuf:"varint,9,opt,name=Triage,proto3" json:"Triage,omitempty"`
	PreConditions        []string `protobuf:"bytes,10,rep,name=PreConditions,json=Pre Condition,proto3" json:"PreConditions,omitempty"`
	MustHave             string   `protobuf:"bytes,11,opt,name=MustHave,json=Must Have,proto3" json:"MustHave,omitempty"`
	URL                  string   `protobuf:"bytes,12,opt,name=URL,json=Article link,proto3" json:"URL,omitempty"`
	Labs                 string   `protobuf:"bytes,13,opt,name=Labs,proto3" json:"Labs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Disease) Reset()         { *m = Disease{} }
func (m *Disease) String() string { return proto.CompactTextString(m) }
func (*Disease) ProtoMessage()    {}
func (*Disease) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6abc415730e04a0, []int{3}
}

func (m *Disease) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Disease.Unmarshal(m, b)
}
func (m *Disease) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Disease.Marshal(b, m, deterministic)
}
func (m *Disease) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Disease.Merge(m, src)
}
func (m *Disease) XXX_Size() int {
	return xxx_messageInfo_Disease.Size(m)
}
func (m *Disease) XXX_DiscardUnknown() {
	xxx_messageInfo_Disease.DiscardUnknown(m)
}

var xxx_messageInfo_Disease proto.InternalMessageInfo

func (m *Disease) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Disease) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Disease) GetNameIndo() string {
	if m != nil {
		return m.NameIndo
	}
	return ""
}

func (m *Disease) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Disease) GetAssociatedSymptoms() []string {
	if m != nil {
		return m.AssociatedSymptoms
	}
	return nil
}

func (m *Disease) GetDifferentials() []string {
	if m != nil {
		return m.Differentials
	}
	return nil
}

func (m *Disease) GetRiskFactors() []string {
	if m != nil {
		return m.RiskFactors
	}
	return nil
}

func (m *Disease) GetLikeliness() string {
	if m != nil {
		return m.Likeliness
	}
	return ""
}

func (m *Disease) GetTriage() int32 {
	if m != nil {
		return m.Triage
	}
	return 0
}

func (m *Disease) GetPreConditions() []string {
	if m != nil {
		return m.PreConditions
	}
	return nil
}

func (m *Disease) GetMustHave() string {
	if m != nil {
		return m.MustHave
	}
	return ""
}

func (m *Disease) GetURL() string {
	if m != nil {
		return m.URL
	}
	return ""
}

func (m *Disease) GetLabs() string {
	if m != nil {
		return m.Labs
	}
	return ""
}

type RecordLabData struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Fields               *Lab     `protobuf:"bytes,2,opt,name=fields,proto3" json:"fields,omitempty"`
	CreatedTime          string   `protobuf:"bytes,3,opt,name=createdTime,proto3" json:"createdTime,omitempty"`
	LastUpdated          string   `protobuf:"bytes,4,opt,name=lastUpdated,proto3" json:"lastUpdated,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecordLabData) Reset()         { *m = RecordLabData{} }
func (m *RecordLabData) String() string { return proto.CompactTextString(m) }
func (*RecordLabData) ProtoMessage()    {}
func (*RecordLabData) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6abc415730e04a0, []int{4}
}

func (m *RecordLabData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecordLabData.Unmarshal(m, b)
}
func (m *RecordLabData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecordLabData.Marshal(b, m, deterministic)
}
func (m *RecordLabData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecordLabData.Merge(m, src)
}
func (m *RecordLabData) XXX_Size() int {
	return xxx_messageInfo_RecordLabData.Size(m)
}
func (m *RecordLabData) XXX_DiscardUnknown() {
	xxx_messageInfo_RecordLabData.DiscardUnknown(m)
}

var xxx_messageInfo_RecordLabData proto.InternalMessageInfo

func (m *RecordLabData) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *RecordLabData) GetFields() *Lab {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *RecordLabData) GetCreatedTime() string {
	if m != nil {
		return m.CreatedTime
	}
	return ""
}

func (m *RecordLabData) GetLastUpdated() string {
	if m != nil {
		return m.LastUpdated
	}
	return ""
}

type Lab struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,json=id,proto3" json:"ID,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	SKU                  string   `protobuf:"bytes,3,opt,name=SKU,proto3" json:"SKU,omitempty"`
	Diseases             []string `protobuf:"bytes,4,rep,name=Diseases,proto3" json:"Diseases,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Lab) Reset()         { *m = Lab{} }
func (m *Lab) String() string { return proto.CompactTextString(m) }
func (*Lab) ProtoMessage()    {}
func (*Lab) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6abc415730e04a0, []int{5}
}

func (m *Lab) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Lab.Unmarshal(m, b)
}
func (m *Lab) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Lab.Marshal(b, m, deterministic)
}
func (m *Lab) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Lab.Merge(m, src)
}
func (m *Lab) XXX_Size() int {
	return xxx_messageInfo_Lab.Size(m)
}
func (m *Lab) XXX_DiscardUnknown() {
	xxx_messageInfo_Lab.DiscardUnknown(m)
}

var xxx_messageInfo_Lab proto.InternalMessageInfo

func (m *Lab) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Lab) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Lab) GetSKU() string {
	if m != nil {
		return m.SKU
	}
	return ""
}

func (m *Lab) GetDiseases() []string {
	if m != nil {
		return m.Diseases
	}
	return nil
}

type RecordDifferentialData struct {
	Id                   string        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Fields               *Differential `protobuf:"bytes,2,opt,name=fields,proto3" json:"fields,omitempty"`
	CreatedTime          string        `protobuf:"bytes,3,opt,name=createdTime,proto3" json:"createdTime,omitempty"`
	LastUpdated          string        `protobuf:"bytes,4,opt,name=lastUpdated,proto3" json:"lastUpdated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *RecordDifferentialData) Reset()         { *m = RecordDifferentialData{} }
func (m *RecordDifferentialData) String() string { return proto.CompactTextString(m) }
func (*RecordDifferentialData) ProtoMessage()    {}
func (*RecordDifferentialData) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6abc415730e04a0, []int{6}
}

func (m *RecordDifferentialData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecordDifferentialData.Unmarshal(m, b)
}
func (m *RecordDifferentialData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecordDifferentialData.Marshal(b, m, deterministic)
}
func (m *RecordDifferentialData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecordDifferentialData.Merge(m, src)
}
func (m *RecordDifferentialData) XXX_Size() int {
	return xxx_messageInfo_RecordDifferentialData.Size(m)
}
func (m *RecordDifferentialData) XXX_DiscardUnknown() {
	xxx_messageInfo_RecordDifferentialData.DiscardUnknown(m)
}

var xxx_messageInfo_RecordDifferentialData proto.InternalMessageInfo

func (m *RecordDifferentialData) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *RecordDifferentialData) GetFields() *Differential {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *RecordDifferentialData) GetCreatedTime() string {
	if m != nil {
		return m.CreatedTime
	}
	return ""
}

func (m *RecordDifferentialData) GetLastUpdated() string {
	if m != nil {
		return m.LastUpdated
	}
	return ""
}

type Differential struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,json=id,proto3" json:"ID,omitempty"`
	Disease              []string `protobuf:"bytes,2,rep,name=Disease,proto3" json:"Disease,omitempty"`
	Symptom              []string `protobuf:"bytes,3,rep,name=Symptom,proto3" json:"Symptom,omitempty"`
	Site                 []string `protobuf:"bytes,4,rep,name=Site,proto3" json:"Site,omitempty"`
	Onset                []string `protobuf:"bytes,5,rep,name=Onset,proto3" json:"Onset,omitempty"`
	Characteristic       []string `protobuf:"bytes,6,rep,name=Characteristic,proto3" json:"Characteristic,omitempty"`
	Timing               []string `protobuf:"bytes,7,rep,name=Timing,proto3" json:"Timing,omitempty"`
	Progressivity        []string `protobuf:"bytes,8,rep,name=Progressivity,proto3" json:"Progressivity,omitempty"`
	Severity             []string `protobuf:"bytes,9,rep,name=Severity,proto3" json:"Severity,omitempty"`
	WorseningFactors     []string `protobuf:"bytes,10,rep,name=WorseningFactors,json=Worsening Factors,proto3" json:"WorseningFactors,omitempty"`
	Radiation            []string `protobuf:"bytes,11,rep,name=Radiation,proto3" json:"Radiation,omitempty"`
	PregnancyState       []string `protobuf:"bytes,12,rep,name=PregnancyState,json=Pregnancy State,proto3" json:"PregnancyState,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Differential) Reset()         { *m = Differential{} }
func (m *Differential) String() string { return proto.CompactTextString(m) }
func (*Differential) ProtoMessage()    {}
func (*Differential) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6abc415730e04a0, []int{7}
}

func (m *Differential) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Differential.Unmarshal(m, b)
}
func (m *Differential) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Differential.Marshal(b, m, deterministic)
}
func (m *Differential) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Differential.Merge(m, src)
}
func (m *Differential) XXX_Size() int {
	return xxx_messageInfo_Differential.Size(m)
}
func (m *Differential) XXX_DiscardUnknown() {
	xxx_messageInfo_Differential.DiscardUnknown(m)
}

var xxx_messageInfo_Differential proto.InternalMessageInfo

func (m *Differential) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Differential) GetDisease() []string {
	if m != nil {
		return m.Disease
	}
	return nil
}

func (m *Differential) GetSymptom() []string {
	if m != nil {
		return m.Symptom
	}
	return nil
}

func (m *Differential) GetSite() []string {
	if m != nil {
		return m.Site
	}
	return nil
}

func (m *Differential) GetOnset() []string {
	if m != nil {
		return m.Onset
	}
	return nil
}

func (m *Differential) GetCharacteristic() []string {
	if m != nil {
		return m.Characteristic
	}
	return nil
}

func (m *Differential) GetTiming() []string {
	if m != nil {
		return m.Timing
	}
	return nil
}

func (m *Differential) GetProgressivity() []string {
	if m != nil {
		return m.Progressivity
	}
	return nil
}

func (m *Differential) GetSeverity() []string {
	if m != nil {
		return m.Severity
	}
	return nil
}

func (m *Differential) GetWorseningFactors() []string {
	if m != nil {
		return m.WorseningFactors
	}
	return nil
}

func (m *Differential) GetRadiation() []string {
	if m != nil {
		return m.Radiation
	}
	return nil
}

func (m *Differential) GetPregnancyState() []string {
	if m != nil {
		return m.PregnancyState
	}
	return nil
}

type RecordTriageData struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Fields               *Triage  `protobuf:"bytes,2,opt,name=fields,proto3" json:"fields,omitempty"`
	CreatedTime          string   `protobuf:"bytes,3,opt,name=createdTime,proto3" json:"createdTime,omitempty"`
	LastUpdated          string   `protobuf:"bytes,4,opt,name=lastUpdated,proto3" json:"lastUpdated,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecordTriageData) Reset()         { *m = RecordTriageData{} }
func (m *RecordTriageData) String() string { return proto.CompactTextString(m) }
func (*RecordTriageData) ProtoMessage()    {}
func (*RecordTriageData) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6abc415730e04a0, []int{8}
}

func (m *RecordTriageData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecordTriageData.Unmarshal(m, b)
}
func (m *RecordTriageData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecordTriageData.Marshal(b, m, deterministic)
}
func (m *RecordTriageData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecordTriageData.Merge(m, src)
}
func (m *RecordTriageData) XXX_Size() int {
	return xxx_messageInfo_RecordTriageData.Size(m)
}
func (m *RecordTriageData) XXX_DiscardUnknown() {
	xxx_messageInfo_RecordTriageData.DiscardUnknown(m)
}

var xxx_messageInfo_RecordTriageData proto.InternalMessageInfo

func (m *RecordTriageData) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *RecordTriageData) GetFields() *Triage {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *RecordTriageData) GetCreatedTime() string {
	if m != nil {
		return m.CreatedTime
	}
	return ""
}

func (m *RecordTriageData) GetLastUpdated() string {
	if m != nil {
		return m.LastUpdated
	}
	return ""
}

type Triage struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,json=id,proto3" json:"ID,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	NameIndo             string   `protobuf:"bytes,3,opt,name=NameIndo,json=Name Indo,proto3" json:"NameIndo,omitempty"`
	Description          string   `protobuf:"bytes,4,opt,name=Description,proto3" json:"Description,omitempty"`
	Symptoms             []string `protobuf:"bytes,5,rep,name=Symptoms,proto3" json:"Symptoms,omitempty"`
	Properties           []string `protobuf:"bytes,6,rep,name=Properties,proto3" json:"Properties,omitempty"`
	Diseases             []string `protobuf:"bytes,7,rep,name=Diseases,proto3" json:"Diseases,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Triage) Reset()         { *m = Triage{} }
func (m *Triage) String() string { return proto.CompactTextString(m) }
func (*Triage) ProtoMessage()    {}
func (*Triage) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6abc415730e04a0, []int{9}
}

func (m *Triage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Triage.Unmarshal(m, b)
}
func (m *Triage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Triage.Marshal(b, m, deterministic)
}
func (m *Triage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Triage.Merge(m, src)
}
func (m *Triage) XXX_Size() int {
	return xxx_messageInfo_Triage.Size(m)
}
func (m *Triage) XXX_DiscardUnknown() {
	xxx_messageInfo_Triage.DiscardUnknown(m)
}

var xxx_messageInfo_Triage proto.InternalMessageInfo

func (m *Triage) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Triage) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Triage) GetNameIndo() string {
	if m != nil {
		return m.NameIndo
	}
	return ""
}

func (m *Triage) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Triage) GetSymptoms() []string {
	if m != nil {
		return m.Symptoms
	}
	return nil
}

func (m *Triage) GetProperties() []string {
	if m != nil {
		return m.Properties
	}
	return nil
}

func (m *Triage) GetDiseases() []string {
	if m != nil {
		return m.Diseases
	}
	return nil
}

type RecordProfileData struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Fields               *Profile `protobuf:"bytes,2,opt,name=fields,proto3" json:"fields,omitempty"`
	CreatedTime          string   `protobuf:"bytes,3,opt,name=createdTime,proto3" json:"createdTime,omitempty"`
	LastUpdated          string   `protobuf:"bytes,4,opt,name=lastUpdated,proto3" json:"lastUpdated,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecordProfileData) Reset()         { *m = RecordProfileData{} }
func (m *RecordProfileData) String() string { return proto.CompactTextString(m) }
func (*RecordProfileData) ProtoMessage()    {}
func (*RecordProfileData) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6abc415730e04a0, []int{10}
}

func (m *RecordProfileData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecordProfileData.Unmarshal(m, b)
}
func (m *RecordProfileData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecordProfileData.Marshal(b, m, deterministic)
}
func (m *RecordProfileData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecordProfileData.Merge(m, src)
}
func (m *RecordProfileData) XXX_Size() int {
	return xxx_messageInfo_RecordProfileData.Size(m)
}
func (m *RecordProfileData) XXX_DiscardUnknown() {
	xxx_messageInfo_RecordProfileData.DiscardUnknown(m)
}

var xxx_messageInfo_RecordProfileData proto.InternalMessageInfo

func (m *RecordProfileData) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *RecordProfileData) GetFields() *Profile {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *RecordProfileData) GetCreatedTime() string {
	if m != nil {
		return m.CreatedTime
	}
	return ""
}

func (m *RecordProfileData) GetLastUpdated() string {
	if m != nil {
		return m.LastUpdated
	}
	return ""
}

type Profile struct {
	ID                          string   `protobuf:"bytes,1,opt,name=ID,json=id,proto3" json:"ID,omitempty"`
	Name                        string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	NameIndo                    string   `protobuf:"bytes,3,opt,name=NameIndo,json=Name Indo,proto3" json:"NameIndo,omitempty"`
	Description                 string   `protobuf:"bytes,4,opt,name=Description,proto3" json:"Description,omitempty"`
	PreconditionDescription     string   `protobuf:"bytes,5,opt,name=PreconditionDescription,json=preconditionDescription,proto3" json:"PreconditionDescription,omitempty"`
	PreconditionDescriptionCopy string   `protobuf:"bytes,6,opt,name=PreconditionDescriptionCopy,json=preconditionDescriptionCopy,proto3" json:"PreconditionDescriptionCopy,omitempty"`
	Type                        string   `protobuf:"bytes,7,opt,name=Type,proto3" json:"Type,omitempty"`
	Order                       int32    `protobuf:"varint,8,opt,name=Order,proto3" json:"Order,omitempty"`
	CustomValue1                float32  `protobuf:"fixed32,9,opt,name=CustomValue1,json=customValue1,proto3" json:"CustomValue1,omitempty"`
	CustomValue2                float32  `protobuf:"fixed32,10,opt,name=CustomValue2,json=customValue2,proto3" json:"CustomValue2,omitempty"`
	XXX_NoUnkeyedLiteral        struct{} `json:"-"`
	XXX_unrecognized            []byte   `json:"-"`
	XXX_sizecache               int32    `json:"-"`
}

func (m *Profile) Reset()         { *m = Profile{} }
func (m *Profile) String() string { return proto.CompactTextString(m) }
func (*Profile) ProtoMessage()    {}
func (*Profile) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6abc415730e04a0, []int{11}
}

func (m *Profile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Profile.Unmarshal(m, b)
}
func (m *Profile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Profile.Marshal(b, m, deterministic)
}
func (m *Profile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Profile.Merge(m, src)
}
func (m *Profile) XXX_Size() int {
	return xxx_messageInfo_Profile.Size(m)
}
func (m *Profile) XXX_DiscardUnknown() {
	xxx_messageInfo_Profile.DiscardUnknown(m)
}

var xxx_messageInfo_Profile proto.InternalMessageInfo

func (m *Profile) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Profile) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Profile) GetNameIndo() string {
	if m != nil {
		return m.NameIndo
	}
	return ""
}

func (m *Profile) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Profile) GetPreconditionDescription() string {
	if m != nil {
		return m.PreconditionDescription
	}
	return ""
}

func (m *Profile) GetPreconditionDescriptionCopy() string {
	if m != nil {
		return m.PreconditionDescriptionCopy
	}
	return ""
}

func (m *Profile) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Profile) GetOrder() int32 {
	if m != nil {
		return m.Order
	}
	return 0
}

func (m *Profile) GetCustomValue1() float32 {
	if m != nil {
		return m.CustomValue1
	}
	return 0
}

func (m *Profile) GetCustomValue2() float32 {
	if m != nil {
		return m.CustomValue2
	}
	return 0
}

func init() {
	proto.RegisterType((*RecordSymptomsData)(nil), "RecordSymptomsData")
	proto.RegisterType((*Symptom)(nil), "Symptom")
	proto.RegisterType((*RecordDiseaseData)(nil), "RecordDiseaseData")
	proto.RegisterType((*Disease)(nil), "Disease")
	proto.RegisterType((*RecordLabData)(nil), "RecordLabData")
	proto.RegisterType((*Lab)(nil), "Lab")
	proto.RegisterType((*RecordDifferentialData)(nil), "RecordDifferentialData")
	proto.RegisterType((*Differential)(nil), "Differential")
	proto.RegisterType((*RecordTriageData)(nil), "RecordTriageData")
	proto.RegisterType((*Triage)(nil), "Triage")
	proto.RegisterType((*RecordProfileData)(nil), "RecordProfileData")
	proto.RegisterType((*Profile)(nil), "Profile")
}

func init() { proto.RegisterFile("proto/nalar/v1/nalar.proto", fileDescriptor_a6abc415730e04a0) }

var fileDescriptor_a6abc415730e04a0 = []byte{
	// 962 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x96, 0xdd, 0x6e, 0x1b, 0x45,
	0x14, 0xc7, 0xe5, 0x8f, 0xc4, 0xf6, 0xb1, 0x9d, 0x8f, 0x51, 0xd5, 0x4c, 0x9b, 0xaa, 0x2c, 0x6e,
	0x81, 0x22, 0x84, 0xab, 0x86, 0x1b, 0x2e, 0xa9, 0x62, 0x2a, 0xa2, 0xb8, 0xc5, 0xd8, 0x0e, 0xbd,
	0x44, 0xe3, 0xdd, 0x13, 0x33, 0xca, 0x7a, 0x67, 0x35, 0x33, 0x4e, 0xf1, 0x25, 0xe2, 0x02, 0xc1,
	0x2d, 0x8f, 0xc4, 0x0b, 0x70, 0xcd, 0x63, 0xf0, 0x04, 0x68, 0x3e, 0x76, 0xbb, 0xde, 0x3a, 0x52,
	0xa4, 0x5a, 0xbd, 0xf2, 0x9c, 0xff, 0xfc, 0x67, 0x3c, 0x1f, 0xbf, 0x73, 0x66, 0xe1, 0x7e, 0x2a,
	0x85, 0x16, 0x4f, 0x13, 0x16, 0x33, 0xf9, 0xf4, 0xfa, 0x99, 0x6b, 0xf4, 0xad, 0xd8, 0xfb, 0xb3,
	0x02, 0x64, 0x8c, 0xa1, 0x90, 0xd1, 0x64, 0xb5, 0x48, 0xb5, 0x58, 0xa8, 0x01, 0xd3, 0x8c, 0xec,
	0x41, 0x95, 0x47, 0xb4, 0x12, 0x54, 0x9e, 0xb4, 0xc6, 0x55, 0x1e, 0x91, 0x00, 0x76, 0x2f, 0x39,
	0xc6, 0x91, 0xa2, 0xd5, 0xa0, 0xf2, 0xa4, 0x7d, 0xd2, 0xec, 0x7b, 0xfb, 0xd8, 0xeb, 0x24, 0x80,
	0x76, 0x28, 0x91, 0x69, 0x8c, 0xa6, 0x7c, 0x81, 0xb4, 0x66, 0x87, 0x16, 0x25, 0xe3, 0x88, 0x99,
	0xd2, 0x17, 0x69, 0x64, 0x24, 0x5a, 0x77, 0x8e, 0x82, 0xd4, 0xfb, 0xa7, 0x0e, 0x0d, 0x3f, 0xaf,
	0x59, 0xc1, 0xd9, 0xa0, 0xb0, 0x02, 0x02, 0xf5, 0x57, 0x6c, 0x81, 0xf6, 0xff, 0x5b, 0x63, 0xdb,
	0x26, 0xc7, 0xd0, 0x34, 0xbf, 0x67, 0x49, 0x24, 0xfc, 0x1f, 0xb6, 0x4c, 0x1c, 0x18, 0x81, 0x7c,
	0x0e, 0xed, 0x01, 0xaa, 0x50, 0xf2, 0x54, 0x73, 0x91, 0xf8, 0xbf, 0xbb, 0x93, 0xed, 0x32, 0x28,
	0xf4, 0x19, 0xeb, 0xb7, 0xbf, 0xa4, 0x31, 0x4b, 0x98, 0xb5, 0xee, 0x94, 0xac, 0x85, 0x3e, 0x72,
	0x1f, 0x9a, 0x3f, 0x2c, 0x51, 0x59, 0xdf, 0xae, 0xf5, 0xe5, 0x31, 0xf9, 0x14, 0xf6, 0x4e, 0x7f,
	0x66, 0x92, 0x85, 0x1a, 0x25, 0x57, 0x9a, 0x87, 0xb4, 0x61, 0x1d, 0x25, 0xd5, 0xf8, 0x5e, 0x62,
	0xc4, 0xd9, 0x0b, 0x1e, 0xa3, 0x1a, 0xf2, 0xe4, 0x8a, 0x36, 0x9d, 0x6f, 0x5d, 0x25, 0x0f, 0xa0,
	0x35, 0x11, 0x4b, 0x19, 0xe2, 0xc5, 0x78, 0x48, 0x5b, 0x6e, 0x7f, 0xb9, 0x40, 0x1e, 0x03, 0x79,
	0xae, 0x94, 0x08, 0xb9, 0x39, 0xba, 0x01, 0x57, 0xc8, 0x14, 0x2a, 0x0a, 0x41, 0xcd, 0xac, 0x29,
	0x8b, 0xc9, 0x3d, 0x68, 0x9e, 0xe3, 0xea, 0x8d, 0x90, 0x91, 0xa2, 0x6d, 0x3b, 0x45, 0xc3, 0xc7,
	0xe4, 0x11, 0x74, 0x07, 0xfc, 0xf2, 0x12, 0x25, 0x26, 0x9a, 0xb3, 0x58, 0xd1, 0x8e, 0x1d, 0xdb,
	0x29, 0x8a, 0xe4, 0x2e, 0xec, 0x4e, 0x25, 0x67, 0x73, 0xa4, 0x5d, 0x3b, 0xda, 0x47, 0xe4, 0x63,
	0x68, 0xbf, 0xe2, 0x21, 0x4e, 0xc5, 0x79, 0x22, 0xde, 0x28, 0xba, 0xe7, 0x86, 0x1a, 0x29, 0x98,
	0x8a, 0xc0, 0x88, 0x66, 0xf9, 0x2f, 0x97, 0x4a, 0x3b, 0xc3, 0xbe, 0x35, 0x58, 0xc1, 0xf5, 0x3e,
	0x82, 0xee, 0x48, 0xe2, 0xa9, 0x48, 0x22, 0x6e, 0x0e, 0x4f, 0xd1, 0x03, 0x37, 0xc5, 0x48, 0x62,
	0x98, 0x89, 0xe4, 0x33, 0xd8, 0xcb, 0x6e, 0xc1, 0xaf, 0xe2, 0xd0, 0xba, 0xf6, 0x33, 0xf5, 0x4b,
	0xbf, 0x9c, 0x23, 0xa8, 0x4d, 0xce, 0x06, 0x94, 0xd8, 0x35, 0x82, 0xef, 0xfd, 0xe9, 0x6c, 0xd0,
	0xfb, 0xa3, 0x02, 0x87, 0x8e, 0x6f, 0x7f, 0x24, 0xb7, 0xc4, 0xdb, 0xbb, 0xb7, 0x8a, 0xf7, 0x6f,
	0x35, 0x68, 0xf8, 0x79, 0xdf, 0x1f, 0xef, 0x60, 0x13, 0xde, 0x45, 0x69, 0x1d, 0x90, 0xec, 0xc0,
	0xe8, 0x8e, 0x03, 0x24, 0x8b, 0xdf, 0xa5, 0x60, 0x77, 0x03, 0x05, 0x01, 0xb4, 0xc7, 0x5c, 0x5d,
	0xbd, 0x60, 0xa1, 0x16, 0x52, 0xd1, 0x86, 0xb5, 0x58, 0x29, 0x70, 0x1a, 0x79, 0x08, 0x30, 0xe4,
	0x57, 0x18, 0xf3, 0x04, 0x95, 0xf2, 0x3c, 0x17, 0x94, 0x02, 0x47, 0x06, 0xe4, 0x9d, 0x9c, 0xa3,
	0xc7, 0x65, 0x0c, 0x1c, 0xc0, 0x46, 0x0c, 0x72, 0xd5, 0x9c, 0x84, 0x21, 0xe7, 0x3b, 0x76, 0x8d,
	0x9e, 0x62, 0x47, 0x92, 0x11, 0xc8, 0x3d, 0xa8, 0x99, 0x04, 0xe9, 0x58, 0xbd, 0xf3, 0x5c, 0x6a,
	0x1e, 0xc6, 0x18, 0xc4, 0x26, 0x83, 0x08, 0xd4, 0x87, 0x6c, 0xa6, 0x3c, 0xbb, 0xb6, 0xdd, 0xfb,
	0xb5, 0x02, 0x5d, 0x47, 0xc4, 0x90, 0xcd, 0x36, 0xd2, 0xf0, 0xa0, 0x44, 0x43, 0xbd, 0x3f, 0x64,
	0xb3, 0xad, 0x92, 0xf0, 0x1a, 0x6a, 0x43, 0x36, 0xbb, 0x15, 0x04, 0x07, 0x50, 0x9b, 0x9c, 0x5f,
	0xf8, 0xbf, 0x31, 0x4d, 0x53, 0x82, 0xf2, 0x74, 0xaf, 0xaf, 0xa7, 0x7b, 0xef, 0xaf, 0x0a, 0xdc,
	0xcd, 0x70, 0x7f, 0x7b, 0x7f, 0x1b, 0x77, 0xf9, 0x49, 0x69, 0x97, 0xdd, 0x7e, 0x71, 0xc8, 0x56,
	0xb7, 0xfb, 0x5f, 0x15, 0xd6, 0x79, 0x2a, 0x6f, 0x9c, 0xe6, 0x89, 0x41, 0xab, 0x76, 0x47, 0x79,
	0x9e, 0xd0, 0xfc, 0x45, 0xa0, 0x35, 0xd7, 0x93, 0x3d, 0x10, 0x04, 0xea, 0x13, 0xae, 0xd1, 0x1f,
	0x81, 0x6d, 0x93, 0x3b, 0xb0, 0xf3, 0x7d, 0xa2, 0x50, 0x7b, 0xca, 0x5d, 0xb0, 0xa1, 0x2e, 0x3b,
	0xc6, 0xcb, 0x75, 0xd9, 0x30, 0xca, 0x17, 0x3c, 0x99, 0x7b, 0xc0, 0x7d, 0xe4, 0x18, 0x15, 0x73,
	0x89, 0x4a, 0xf1, 0x6b, 0xae, 0x57, 0xb4, 0x99, 0x31, 0x5a, 0x10, 0xcd, 0xb5, 0x4c, 0xf0, 0x1a,
	0xa5, 0x31, 0xb4, 0x7c, 0x92, 0xf9, 0x98, 0x7c, 0x01, 0x07, 0xaf, 0x85, 0x54, 0x98, 0xf0, 0x64,
	0x9e, 0x25, 0x91, 0x03, 0xfd, 0x30, 0xd7, 0x7d, 0x26, 0x29, 0x53, 0x37, 0xc7, 0x2c, 0xe2, 0xee,
	0x2d, 0x6a, 0xbb, 0xba, 0x99, 0x0b, 0xa6, 0x24, 0x8e, 0x24, 0xce, 0x13, 0x96, 0x84, 0xab, 0x89,
	0x66, 0x1a, 0x7d, 0xd9, 0xde, 0xcf, 0xd5, 0xc0, 0xca, 0xbd, 0xdf, 0x2b, 0x70, 0xe0, 0x50, 0x70,
	0xa9, 0xb6, 0x11, 0x82, 0x8f, 0x4a, 0x10, 0x34, 0xfa, 0xce, 0xbc, 0xd5, 0xeb, 0xff, 0xbb, 0x92,
	0x25, 0xff, 0x87, 0x28, 0x7b, 0xe6, 0x1e, 0x6e, 0x2a, 0x76, 0x0f, 0x01, 0x46, 0x52, 0xa4, 0x28,
	0x35, 0xc7, 0xac, 0xd2, 0x15, 0x94, 0xb5, 0xd4, 0x6a, 0x94, 0x52, 0xeb, 0xed, 0x4b, 0x32, 0x92,
	0xe2, 0x92, 0xc7, 0xb7, 0x7d, 0x49, 0xbc, 0x7b, 0xab, 0x27, 0xfa, 0x6f, 0x15, 0x1a, 0x7e, 0xde,
	0x0f, 0x71, 0xa4, 0x5f, 0xc3, 0x51, 0xf1, 0x59, 0x2e, 0xba, 0xdd, 0xb7, 0xd2, 0x51, 0xba, 0xb9,
	0x9b, 0x7c, 0x03, 0xc7, 0x37, 0x8c, 0x3c, 0x15, 0xe9, 0xca, 0x7f, 0x41, 0x1d, 0xa7, 0x37, 0x5b,
	0xcc, 0x76, 0xa6, 0xab, 0x14, 0xfd, 0xa7, 0x94, 0x6d, 0xdb, 0x34, 0x97, 0x11, 0x4a, 0xfb, 0xce,
	0xec, 0x8c, 0x5d, 0x40, 0x7a, 0xd0, 0x39, 0x5d, 0x2a, 0x2d, 0x16, 0x3f, 0xb2, 0x78, 0x89, 0xcf,
	0xec, 0x43, 0x53, 0x1d, 0x77, 0xc2, 0x82, 0x56, 0xf2, 0x9c, 0x50, 0x78, 0xc7, 0x73, 0x32, 0xdb,
	0xb5, 0x5f, 0xc6, 0x5f, 0xfd, 0x1f, 0x00, 0x00, 0xff, 0xff, 0x0d, 0xa6, 0xe8, 0xea, 0x37, 0x0b,
	0x00, 0x00,
}
