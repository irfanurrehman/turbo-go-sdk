// Code generated by protoc-gen-go. DO NOT EDIT.
// source: NonMarketEntityDTO.proto

package proto

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

// Enumerate non market entities
type NonMarketEntityDTO_NonMarketEntityType int32

const (
	NonMarketEntityDTO_CLOUD_SERVICE    NonMarketEntityDTO_NonMarketEntityType = 0
	NonMarketEntityDTO_ACCOUNT          NonMarketEntityDTO_NonMarketEntityType = 2
	NonMarketEntityDTO_PLAN_DESTINATION NonMarketEntityDTO_NonMarketEntityType = 3
)

var NonMarketEntityDTO_NonMarketEntityType_name = map[int32]string{
	0: "CLOUD_SERVICE",
	2: "ACCOUNT",
	3: "PLAN_DESTINATION",
}

var NonMarketEntityDTO_NonMarketEntityType_value = map[string]int32{
	"CLOUD_SERVICE":    0,
	"ACCOUNT":          2,
	"PLAN_DESTINATION": 3,
}

func (x NonMarketEntityDTO_NonMarketEntityType) Enum() *NonMarketEntityDTO_NonMarketEntityType {
	p := new(NonMarketEntityDTO_NonMarketEntityType)
	*p = x
	return p
}

func (x NonMarketEntityDTO_NonMarketEntityType) String() string {
	return proto.EnumName(NonMarketEntityDTO_NonMarketEntityType_name, int32(x))
}

func (x *NonMarketEntityDTO_NonMarketEntityType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(NonMarketEntityDTO_NonMarketEntityType_value, data, "NonMarketEntityDTO_NonMarketEntityType")
	if err != nil {
		return err
	}
	*x = NonMarketEntityDTO_NonMarketEntityType(value)
	return nil
}

func (NonMarketEntityDTO_NonMarketEntityType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a7ef4e88daee47ae, []int{0, 0}
}

type NonMarketEntityDTO_CloudServiceData_PriceModel int32

const (
	// Pay for compute capacity by per hour or per second depending
	// on which instances are run.
	NonMarketEntityDTO_CloudServiceData_ON_DEMAND NonMarketEntityDTO_CloudServiceData_PriceModel = 0
	// Pay a discounted price for reserved instances with a significant
	// discount (up to 75%) compared to On-Demand instance pricing.
	NonMarketEntityDTO_CloudServiceData_RESERVED NonMarketEntityDTO_CloudServiceData_PriceModel = 1
	// Pay the Spot price that is in effect for the time period the
	// instances are running and prices are set by Amazon EC2 and
	// adjusted gradually based on long-term trends in supply and demand.
	NonMarketEntityDTO_CloudServiceData_SPOT NonMarketEntityDTO_CloudServiceData_PriceModel = 2
)

var NonMarketEntityDTO_CloudServiceData_PriceModel_name = map[int32]string{
	0: "ON_DEMAND",
	1: "RESERVED",
	2: "SPOT",
}

var NonMarketEntityDTO_CloudServiceData_PriceModel_value = map[string]int32{
	"ON_DEMAND": 0,
	"RESERVED":  1,
	"SPOT":      2,
}

func (x NonMarketEntityDTO_CloudServiceData_PriceModel) Enum() *NonMarketEntityDTO_CloudServiceData_PriceModel {
	p := new(NonMarketEntityDTO_CloudServiceData_PriceModel)
	*p = x
	return p
}

func (x NonMarketEntityDTO_CloudServiceData_PriceModel) String() string {
	return proto.EnumName(NonMarketEntityDTO_CloudServiceData_PriceModel_name, int32(x))
}

func (x *NonMarketEntityDTO_CloudServiceData_PriceModel) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(NonMarketEntityDTO_CloudServiceData_PriceModel_value, data, "NonMarketEntityDTO_CloudServiceData_PriceModel")
	if err != nil {
		return err
	}
	*x = NonMarketEntityDTO_CloudServiceData_PriceModel(value)
	return nil
}

func (NonMarketEntityDTO_CloudServiceData_PriceModel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a7ef4e88daee47ae, []int{0, 0, 0}
}

// Types of Billing Data, based on the whether the discovering probe is also
// creating the main Entities or only BillingData or both.
type NonMarketEntityDTO_CloudServiceData_BillingData_Type int32

const (
	// Default implies that the discovering probe is capable of discovering BillingData
	// as well as the Entities being billed. For example, Probe 'X' discovers EntityDTOs
	// *and* BillingData for those EntityDTOs marks the BillingData as DEFAULT.
	NonMarketEntityDTO_CloudServiceData_BillingData_DEFAULT NonMarketEntityDTO_CloudServiceData_BillingData_Type = 0
	// Proxy implies that the discovering probe fetches BillingData but is not capable
	// of discovering Entities for which the BillingData is being discovered. For this
	// type of BillingData, DTOs in virtual_machines need to be marked proxy.
	// Optionally, mainTargetId can be set to the target id for which this discovery
	// probe is the acting proxy.
	// For example, Probe 'X' discovers EntityDTOs, Probe 'Y' discovers BillingData,
	// Probe 'Y' marks BillingData as PROXY.
	NonMarketEntityDTO_CloudServiceData_BillingData_PROXY NonMarketEntityDTO_CloudServiceData_BillingData_Type = 1
	// Stitching implies that the discovering probe is capable of discovering Entities
	// but not the BillingData for the Entities. This type of BillingData would contain
	// stitching information in the virtual_machines list.
	// For example, Probe 'X' discovers EntityDTOs, Probe 'Y' discovers BillingData,
	// Probe 'X' marks BillingData as STITCHING. virtual_machines in this BillingData
	// contain VM Billing Id in the entityProperties list.
	NonMarketEntityDTO_CloudServiceData_BillingData_STITCHING NonMarketEntityDTO_CloudServiceData_BillingData_Type = 2
)

var NonMarketEntityDTO_CloudServiceData_BillingData_Type_name = map[int32]string{
	0: "DEFAULT",
	1: "PROXY",
	2: "STITCHING",
}

var NonMarketEntityDTO_CloudServiceData_BillingData_Type_value = map[string]int32{
	"DEFAULT":   0,
	"PROXY":     1,
	"STITCHING": 2,
}

func (x NonMarketEntityDTO_CloudServiceData_BillingData_Type) Enum() *NonMarketEntityDTO_CloudServiceData_BillingData_Type {
	p := new(NonMarketEntityDTO_CloudServiceData_BillingData_Type)
	*p = x
	return p
}

func (x NonMarketEntityDTO_CloudServiceData_BillingData_Type) String() string {
	return proto.EnumName(NonMarketEntityDTO_CloudServiceData_BillingData_Type_name, int32(x))
}

func (x *NonMarketEntityDTO_CloudServiceData_BillingData_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(NonMarketEntityDTO_CloudServiceData_BillingData_Type_value, data, "NonMarketEntityDTO_CloudServiceData_BillingData_Type")
	if err != nil {
		return err
	}
	*x = NonMarketEntityDTO_CloudServiceData_BillingData_Type(value)
	return nil
}

func (NonMarketEntityDTO_CloudServiceData_BillingData_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a7ef4e88daee47ae, []int{0, 0, 0, 0}
}

// The NonMarketEntityDTO message represents an Entity discovered in the target that your probe is
// monitoring and this entity does not participate in the market.
//
// Each entity must have a unique ID to identify it in the Operations Manager repository.
// Many targets provide unique IDs for their entities, or you can generate your own.
// To guarantee that it's unique, you can give the ID a prefix that identifies your
// probe and the given target.
//
// Specify entity type by setting a 'NonMarketEntityType' value to the 'entityType' field.
//
// The 'displayName' value appears in the product GUI and in reports to identify the entity.
//
type NonMarketEntityDTO struct {
	EntityType  *NonMarketEntityDTO_NonMarketEntityType `protobuf:"varint,1,req,name=entityType,enum=common_dto.NonMarketEntityDTO_NonMarketEntityType" json:"entityType,omitempty"`
	Id          *string                                 `protobuf:"bytes,2,req,name=id" json:"id,omitempty"`
	DisplayName *string                                 `protobuf:"bytes,3,opt,name=displayName" json:"displayName,omitempty"`
	Description *string                                 `protobuf:"bytes,4,opt,name=description" json:"description,omitempty"`
	// Notifications associated with the entity
	Notification []*NotificationDTO `protobuf:"bytes,5,rep,name=notification" json:"notification,omitempty"`
	// list of <string, string, string> namespace, key, value triplets
	EntityProperties []*EntityDTO_EntityProperty `protobuf:"bytes,6,rep,name=entityProperties" json:"entityProperties,omitempty"`
	// Collection of entity type's specific data.
	//
	// Types that are valid to be assigned to EntityData:
	//	*NonMarketEntityDTO_CloudServiceData_
	//	*NonMarketEntityDTO_PlanDestinationData_
	EntityData           isNonMarketEntityDTO_EntityData `protobuf_oneof:"entity_data"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *NonMarketEntityDTO) Reset()         { *m = NonMarketEntityDTO{} }
func (m *NonMarketEntityDTO) String() string { return proto.CompactTextString(m) }
func (*NonMarketEntityDTO) ProtoMessage()    {}
func (*NonMarketEntityDTO) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7ef4e88daee47ae, []int{0}
}

func (m *NonMarketEntityDTO) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NonMarketEntityDTO.Unmarshal(m, b)
}
func (m *NonMarketEntityDTO) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NonMarketEntityDTO.Marshal(b, m, deterministic)
}
func (m *NonMarketEntityDTO) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NonMarketEntityDTO.Merge(m, src)
}
func (m *NonMarketEntityDTO) XXX_Size() int {
	return xxx_messageInfo_NonMarketEntityDTO.Size(m)
}
func (m *NonMarketEntityDTO) XXX_DiscardUnknown() {
	xxx_messageInfo_NonMarketEntityDTO.DiscardUnknown(m)
}

var xxx_messageInfo_NonMarketEntityDTO proto.InternalMessageInfo

func (m *NonMarketEntityDTO) GetEntityType() NonMarketEntityDTO_NonMarketEntityType {
	if m != nil && m.EntityType != nil {
		return *m.EntityType
	}
	return NonMarketEntityDTO_CLOUD_SERVICE
}

func (m *NonMarketEntityDTO) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *NonMarketEntityDTO) GetDisplayName() string {
	if m != nil && m.DisplayName != nil {
		return *m.DisplayName
	}
	return ""
}

func (m *NonMarketEntityDTO) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return ""
}

func (m *NonMarketEntityDTO) GetNotification() []*NotificationDTO {
	if m != nil {
		return m.Notification
	}
	return nil
}

func (m *NonMarketEntityDTO) GetEntityProperties() []*EntityDTO_EntityProperty {
	if m != nil {
		return m.EntityProperties
	}
	return nil
}

type isNonMarketEntityDTO_EntityData interface {
	isNonMarketEntityDTO_EntityData()
}

type NonMarketEntityDTO_CloudServiceData_ struct {
	CloudServiceData *NonMarketEntityDTO_CloudServiceData `protobuf:"bytes,500,opt,name=cloud_service_data,json=cloudServiceData,oneof"`
}

type NonMarketEntityDTO_PlanDestinationData_ struct {
	PlanDestinationData *NonMarketEntityDTO_PlanDestinationData `protobuf:"bytes,502,opt,name=plan_destination_data,json=planDestinationData,oneof"`
}

func (*NonMarketEntityDTO_CloudServiceData_) isNonMarketEntityDTO_EntityData() {}

func (*NonMarketEntityDTO_PlanDestinationData_) isNonMarketEntityDTO_EntityData() {}

func (m *NonMarketEntityDTO) GetEntityData() isNonMarketEntityDTO_EntityData {
	if m != nil {
		return m.EntityData
	}
	return nil
}

func (m *NonMarketEntityDTO) GetCloudServiceData() *NonMarketEntityDTO_CloudServiceData {
	if x, ok := m.GetEntityData().(*NonMarketEntityDTO_CloudServiceData_); ok {
		return x.CloudServiceData
	}
	return nil
}

func (m *NonMarketEntityDTO) GetPlanDestinationData() *NonMarketEntityDTO_PlanDestinationData {
	if x, ok := m.GetEntityData().(*NonMarketEntityDTO_PlanDestinationData_); ok {
		return x.PlanDestinationData
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*NonMarketEntityDTO) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*NonMarketEntityDTO_CloudServiceData_)(nil),
		(*NonMarketEntityDTO_PlanDestinationData_)(nil),
	}
}

type NonMarketEntityDTO_CloudServiceData struct {
	CloudProvider        *string                                          `protobuf:"bytes,1,req,name=cloud_provider,json=cloudProvider" json:"cloud_provider,omitempty"`
	BillingData          *NonMarketEntityDTO_CloudServiceData_BillingData `protobuf:"bytes,3,opt,name=billing_data,json=billingData" json:"billing_data,omitempty"`
	AccountId            *string                                          `protobuf:"bytes,4,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	PriceModels          []NonMarketEntityDTO_CloudServiceData_PriceModel `protobuf:"varint,5,rep,name=price_models,json=priceModels,enum=common_dto.NonMarketEntityDTO_CloudServiceData_PriceModel" json:"price_models,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                         `json:"-"`
	XXX_unrecognized     []byte                                           `json:"-"`
	XXX_sizecache        int32                                            `json:"-"`
}

func (m *NonMarketEntityDTO_CloudServiceData) Reset()         { *m = NonMarketEntityDTO_CloudServiceData{} }
func (m *NonMarketEntityDTO_CloudServiceData) String() string { return proto.CompactTextString(m) }
func (*NonMarketEntityDTO_CloudServiceData) ProtoMessage()    {}
func (*NonMarketEntityDTO_CloudServiceData) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7ef4e88daee47ae, []int{0, 0}
}

func (m *NonMarketEntityDTO_CloudServiceData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NonMarketEntityDTO_CloudServiceData.Unmarshal(m, b)
}
func (m *NonMarketEntityDTO_CloudServiceData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NonMarketEntityDTO_CloudServiceData.Marshal(b, m, deterministic)
}
func (m *NonMarketEntityDTO_CloudServiceData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NonMarketEntityDTO_CloudServiceData.Merge(m, src)
}
func (m *NonMarketEntityDTO_CloudServiceData) XXX_Size() int {
	return xxx_messageInfo_NonMarketEntityDTO_CloudServiceData.Size(m)
}
func (m *NonMarketEntityDTO_CloudServiceData) XXX_DiscardUnknown() {
	xxx_messageInfo_NonMarketEntityDTO_CloudServiceData.DiscardUnknown(m)
}

var xxx_messageInfo_NonMarketEntityDTO_CloudServiceData proto.InternalMessageInfo

func (m *NonMarketEntityDTO_CloudServiceData) GetCloudProvider() string {
	if m != nil && m.CloudProvider != nil {
		return *m.CloudProvider
	}
	return ""
}

func (m *NonMarketEntityDTO_CloudServiceData) GetBillingData() *NonMarketEntityDTO_CloudServiceData_BillingData {
	if m != nil {
		return m.BillingData
	}
	return nil
}

func (m *NonMarketEntityDTO_CloudServiceData) GetAccountId() string {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return ""
}

func (m *NonMarketEntityDTO_CloudServiceData) GetPriceModels() []NonMarketEntityDTO_CloudServiceData_PriceModel {
	if m != nil {
		return m.PriceModels
	}
	return nil
}

type NonMarketEntityDTO_CloudServiceData_BillingData struct {
	VirtualMachines   []*EntityDTO                                          `protobuf:"bytes,1,rep,name=virtual_machines,json=virtualMachines" json:"virtual_machines,omitempty"`
	ReservedInstances []*EntityDTO                                          `protobuf:"bytes,2,rep,name=reserved_instances,json=reservedInstances" json:"reserved_instances,omitempty"`
	BillingDataUuid   *string                                               `protobuf:"bytes,3,opt,name=billingDataUuid" json:"billingDataUuid,omitempty"`
	Type              *NonMarketEntityDTO_CloudServiceData_BillingData_Type `protobuf:"varint,4,opt,name=type,enum=common_dto.NonMarketEntityDTO_CloudServiceData_BillingData_Type,def=0" json:"type,omitempty"`
	// Optionally set for Proxy Type BillingData. Represents the target id which is the
	// source of Entities whose BillingData is being fetched by the discovering probe.
	MainTargetId         *string  `protobuf:"bytes,5,opt,name=mainTargetId" json:"mainTargetId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NonMarketEntityDTO_CloudServiceData_BillingData) Reset() {
	*m = NonMarketEntityDTO_CloudServiceData_BillingData{}
}
func (m *NonMarketEntityDTO_CloudServiceData_BillingData) String() string {
	return proto.CompactTextString(m)
}
func (*NonMarketEntityDTO_CloudServiceData_BillingData) ProtoMessage() {}
func (*NonMarketEntityDTO_CloudServiceData_BillingData) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7ef4e88daee47ae, []int{0, 0, 0}
}

func (m *NonMarketEntityDTO_CloudServiceData_BillingData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NonMarketEntityDTO_CloudServiceData_BillingData.Unmarshal(m, b)
}
func (m *NonMarketEntityDTO_CloudServiceData_BillingData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NonMarketEntityDTO_CloudServiceData_BillingData.Marshal(b, m, deterministic)
}
func (m *NonMarketEntityDTO_CloudServiceData_BillingData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NonMarketEntityDTO_CloudServiceData_BillingData.Merge(m, src)
}
func (m *NonMarketEntityDTO_CloudServiceData_BillingData) XXX_Size() int {
	return xxx_messageInfo_NonMarketEntityDTO_CloudServiceData_BillingData.Size(m)
}
func (m *NonMarketEntityDTO_CloudServiceData_BillingData) XXX_DiscardUnknown() {
	xxx_messageInfo_NonMarketEntityDTO_CloudServiceData_BillingData.DiscardUnknown(m)
}

var xxx_messageInfo_NonMarketEntityDTO_CloudServiceData_BillingData proto.InternalMessageInfo

const Default_NonMarketEntityDTO_CloudServiceData_BillingData_Type NonMarketEntityDTO_CloudServiceData_BillingData_Type = NonMarketEntityDTO_CloudServiceData_BillingData_DEFAULT

func (m *NonMarketEntityDTO_CloudServiceData_BillingData) GetVirtualMachines() []*EntityDTO {
	if m != nil {
		return m.VirtualMachines
	}
	return nil
}

func (m *NonMarketEntityDTO_CloudServiceData_BillingData) GetReservedInstances() []*EntityDTO {
	if m != nil {
		return m.ReservedInstances
	}
	return nil
}

func (m *NonMarketEntityDTO_CloudServiceData_BillingData) GetBillingDataUuid() string {
	if m != nil && m.BillingDataUuid != nil {
		return *m.BillingDataUuid
	}
	return ""
}

func (m *NonMarketEntityDTO_CloudServiceData_BillingData) GetType() NonMarketEntityDTO_CloudServiceData_BillingData_Type {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Default_NonMarketEntityDTO_CloudServiceData_BillingData_Type
}

func (m *NonMarketEntityDTO_CloudServiceData_BillingData) GetMainTargetId() string {
	if m != nil && m.MainTargetId != nil {
		return *m.MainTargetId
	}
	return ""
}

type NonMarketEntityDTO_PlanDestinationData struct {
	// Set to true if Plan destination has exported data to Azure Migrate Project.
	HasExportedData      *bool    `protobuf:"varint,1,req,name=has_exported_data,json=hasExportedData" json:"has_exported_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NonMarketEntityDTO_PlanDestinationData) Reset() {
	*m = NonMarketEntityDTO_PlanDestinationData{}
}
func (m *NonMarketEntityDTO_PlanDestinationData) String() string { return proto.CompactTextString(m) }
func (*NonMarketEntityDTO_PlanDestinationData) ProtoMessage()    {}
func (*NonMarketEntityDTO_PlanDestinationData) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7ef4e88daee47ae, []int{0, 1}
}

func (m *NonMarketEntityDTO_PlanDestinationData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NonMarketEntityDTO_PlanDestinationData.Unmarshal(m, b)
}
func (m *NonMarketEntityDTO_PlanDestinationData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NonMarketEntityDTO_PlanDestinationData.Marshal(b, m, deterministic)
}
func (m *NonMarketEntityDTO_PlanDestinationData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NonMarketEntityDTO_PlanDestinationData.Merge(m, src)
}
func (m *NonMarketEntityDTO_PlanDestinationData) XXX_Size() int {
	return xxx_messageInfo_NonMarketEntityDTO_PlanDestinationData.Size(m)
}
func (m *NonMarketEntityDTO_PlanDestinationData) XXX_DiscardUnknown() {
	xxx_messageInfo_NonMarketEntityDTO_PlanDestinationData.DiscardUnknown(m)
}

var xxx_messageInfo_NonMarketEntityDTO_PlanDestinationData proto.InternalMessageInfo

func (m *NonMarketEntityDTO_PlanDestinationData) GetHasExportedData() bool {
	if m != nil && m.HasExportedData != nil {
		return *m.HasExportedData
	}
	return false
}

// Generic cost data struct used for sending cost/spend from probe to the platform.
type CostDataDTO struct {
	// UUID of the entity (Account/CloudService) or profile (VMTemplate/DBTemplate)
	Id *string `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	// Whether the cost applies to an entity (default), or a profile.
	AppliesProfile *bool `protobuf:"varint,2,req,name=applies_profile,json=appliesProfile" json:"applies_profile,omitempty"`
	// Type of the cost data entity or the entity that the profile applies to.
	EntityType *EntityDTO_EntityType `protobuf:"varint,3,req,name=entity_type,json=entityType,enum=common_dto.EntityDTO_EntityType" json:"entity_type,omitempty"`
	// Business account id
	AccountId *string `protobuf:"bytes,4,req,name=account_id,json=accountId" json:"account_id,omitempty"`
	// Cost of the associated entity (e.g template/VM etc.)
	// If representing spend (top-down), then how much was spent for this account/service/template,
	// parsed from the bill.
	Cost *float32 `protobuf:"fixed32,5,req,name=cost" json:"cost,omitempty"`
	// The date for which the spent captured in this costDTO is relevant
	// for example a 'Backup' service that was used for a specific date in the cloud
	UsageDate            *int64   `protobuf:"varint,6,opt,name=usage_date,json=usageDate" json:"usage_date,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CostDataDTO) Reset()         { *m = CostDataDTO{} }
func (m *CostDataDTO) String() string { return proto.CompactTextString(m) }
func (*CostDataDTO) ProtoMessage()    {}
func (*CostDataDTO) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7ef4e88daee47ae, []int{1}
}

func (m *CostDataDTO) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CostDataDTO.Unmarshal(m, b)
}
func (m *CostDataDTO) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CostDataDTO.Marshal(b, m, deterministic)
}
func (m *CostDataDTO) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CostDataDTO.Merge(m, src)
}
func (m *CostDataDTO) XXX_Size() int {
	return xxx_messageInfo_CostDataDTO.Size(m)
}
func (m *CostDataDTO) XXX_DiscardUnknown() {
	xxx_messageInfo_CostDataDTO.DiscardUnknown(m)
}

var xxx_messageInfo_CostDataDTO proto.InternalMessageInfo

func (m *CostDataDTO) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *CostDataDTO) GetAppliesProfile() bool {
	if m != nil && m.AppliesProfile != nil {
		return *m.AppliesProfile
	}
	return false
}

func (m *CostDataDTO) GetEntityType() EntityDTO_EntityType {
	if m != nil && m.EntityType != nil {
		return *m.EntityType
	}
	return EntityDTO_SWITCH
}

func (m *CostDataDTO) GetAccountId() string {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return ""
}

func (m *CostDataDTO) GetCost() float32 {
	if m != nil && m.Cost != nil {
		return *m.Cost
	}
	return 0
}

func (m *CostDataDTO) GetUsageDate() int64 {
	if m != nil && m.UsageDate != nil {
		return *m.UsageDate
	}
	return 0
}

func init() {
	proto.RegisterEnum("common_dto.NonMarketEntityDTO_NonMarketEntityType", NonMarketEntityDTO_NonMarketEntityType_name, NonMarketEntityDTO_NonMarketEntityType_value)
	proto.RegisterEnum("common_dto.NonMarketEntityDTO_CloudServiceData_PriceModel", NonMarketEntityDTO_CloudServiceData_PriceModel_name, NonMarketEntityDTO_CloudServiceData_PriceModel_value)
	proto.RegisterEnum("common_dto.NonMarketEntityDTO_CloudServiceData_BillingData_Type", NonMarketEntityDTO_CloudServiceData_BillingData_Type_name, NonMarketEntityDTO_CloudServiceData_BillingData_Type_value)
	proto.RegisterType((*NonMarketEntityDTO)(nil), "common_dto.NonMarketEntityDTO")
	proto.RegisterType((*NonMarketEntityDTO_CloudServiceData)(nil), "common_dto.NonMarketEntityDTO.CloudServiceData")
	proto.RegisterType((*NonMarketEntityDTO_CloudServiceData_BillingData)(nil), "common_dto.NonMarketEntityDTO.CloudServiceData.BillingData")
	proto.RegisterType((*NonMarketEntityDTO_PlanDestinationData)(nil), "common_dto.NonMarketEntityDTO.PlanDestinationData")
	proto.RegisterType((*CostDataDTO)(nil), "common_dto.CostDataDTO")
}

func init() { proto.RegisterFile("NonMarketEntityDTO.proto", fileDescriptor_a7ef4e88daee47ae) }

var fileDescriptor_a7ef4e88daee47ae = []byte{
	// 851 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0xae, 0x9d, 0xa4, 0x24, 0xc7, 0xf9, 0x71, 0xa7, 0xac, 0x64, 0x15, 0x21, 0xa2, 0x08, 0x44,
	0xb4, 0x12, 0x59, 0x29, 0xdc, 0x2d, 0x17, 0x6c, 0x1a, 0x7b, 0xd9, 0x40, 0x6b, 0x5b, 0x13, 0x97,
	0x05, 0xa1, 0xc5, 0x4c, 0xed, 0x69, 0x3b, 0x5a, 0xdb, 0x63, 0xd9, 0x93, 0x2e, 0x7d, 0x20, 0xde,
	0x83, 0xf7, 0xe0, 0x29, 0x90, 0x60, 0xaf, 0xd1, 0x8c, 0xdd, 0x24, 0x4d, 0xab, 0x5d, 0xad, 0xb8,
	0x1b, 0x7f, 0x3e, 0xdf, 0x39, 0xdf, 0xcc, 0x39, 0xe7, 0x03, 0xcb, 0xe5, 0xd9, 0x29, 0x29, 0x5e,
	0x53, 0xe1, 0x64, 0x82, 0x89, 0x1b, 0x3b, 0xf0, 0x26, 0x79, 0xc1, 0x05, 0x47, 0x10, 0xf1, 0x34,
	0xe5, 0x59, 0x18, 0x0b, 0x7e, 0x34, 0x98, 0xab, 0xf3, 0xfa, 0xe7, 0xe8, 0x0f, 0x03, 0xd0, 0x7d,
	0x26, 0xc2, 0x00, 0x54, 0x7d, 0x04, 0x37, 0x39, 0xb5, 0xb4, 0xa1, 0x3e, 0xee, 0x4f, 0xa7, 0x93,
	0x4d, 0xa2, 0xc9, 0x03, 0xd5, 0x76, 0x20, 0xc9, 0xc4, 0x5b, 0x59, 0x50, 0x1f, 0x74, 0x16, 0x5b,
	0xfa, 0x50, 0x1f, 0x77, 0xb0, 0xce, 0x62, 0x34, 0x04, 0x23, 0x66, 0x65, 0x9e, 0x90, 0x1b, 0x97,
	0xa4, 0xd4, 0x6a, 0x0c, 0xb5, 0x71, 0x07, 0x6f, 0x43, 0x2a, 0x82, 0x96, 0x51, 0xc1, 0x72, 0xc1,
	0x78, 0x66, 0x35, 0xeb, 0x88, 0x0d, 0x84, 0xbe, 0x85, 0x6e, 0xc6, 0x05, 0xbb, 0x60, 0x11, 0x51,
	0x21, 0xad, 0x61, 0x63, 0x6c, 0x4c, 0x3f, 0xb9, 0xab, 0x74, 0xf3, 0xdf, 0x0e, 0x3c, 0x7c, 0x87,
	0x80, 0x7c, 0x30, 0x2b, 0x89, 0x7e, 0xc1, 0x73, 0x5a, 0x08, 0x46, 0x4b, 0x6b, 0x5f, 0x25, 0xf9,
	0x7c, 0x3b, 0xc9, 0xe6, 0x96, 0xce, 0x76, 0xf4, 0x0d, 0xbe, 0xc7, 0x46, 0xbf, 0x01, 0x8a, 0x12,
	0xbe, 0x8a, 0xc3, 0x92, 0x16, 0xd7, 0x2c, 0xa2, 0x61, 0x4c, 0x04, 0xb1, 0xfe, 0x91, 0xd7, 0x33,
	0xa6, 0x4f, 0xde, 0xf3, 0x86, 0x73, 0xc9, 0x5c, 0x56, 0x44, 0x9b, 0x08, 0xf2, 0x62, 0x0f, 0x9b,
	0xd1, 0x0e, 0x86, 0x18, 0x3c, 0xca, 0x13, 0x92, 0x85, 0x31, 0x2d, 0x05, 0xcb, 0xd4, 0x3d, 0xaa,
	0x22, 0x6f, 0xab, 0x22, 0xef, 0x6b, 0x94, 0x9f, 0x90, 0xcc, 0xde, 0x70, 0xeb, 0x3a, 0x87, 0xf9,
	0x7d, 0xf8, 0xe8, 0xcf, 0x16, 0x98, 0xbb, 0x9a, 0xd0, 0x17, 0xd0, 0xaf, 0x6e, 0x98, 0x17, 0xfc,
	0x9a, 0xc5, 0xb4, 0x50, 0x03, 0xd2, 0xc1, 0x3d, 0x85, 0xfa, 0x35, 0x88, 0x7e, 0x85, 0xee, 0x39,
	0x4b, 0x12, 0x96, 0x5d, 0x56, 0xea, 0x2a, 0x71, 0xdf, 0x7c, 0xe0, 0x0b, 0x4c, 0x8e, 0xab, 0x1c,
	0xf2, 0x8c, 0x8d, 0xf3, 0xcd, 0x07, 0xfa, 0x14, 0x80, 0x44, 0x11, 0x5f, 0x65, 0x22, 0x64, 0x71,
	0x3d, 0x1c, 0x9d, 0x1a, 0x59, 0xc4, 0xe8, 0x15, 0x74, 0xf3, 0x42, 0xbe, 0x7f, 0xca, 0x63, 0x9a,
	0x94, 0x6a, 0x34, 0xfa, 0xd3, 0xa7, 0x1f, 0x5a, 0xde, 0x97, 0x39, 0x4e, 0x65, 0x0a, 0x6c, 0xe4,
	0xeb, 0x73, 0x79, 0xf4, 0xb7, 0x0e, 0xc6, 0x96, 0x34, 0xf4, 0x0c, 0xcc, 0x6b, 0x56, 0x88, 0x15,
	0x49, 0xc2, 0x94, 0x44, 0x57, 0x2c, 0xa3, 0xa5, 0xa5, 0xa9, 0x41, 0x7a, 0xf4, 0xe0, 0x20, 0xe1,
	0x41, 0x1d, 0x7e, 0x5a, 0x47, 0x23, 0x1b, 0x50, 0x41, 0xe5, 0xd0, 0xd0, 0x38, 0x64, 0x59, 0x29,
	0x48, 0x16, 0xd1, 0xd2, 0xd2, 0xdf, 0x95, 0xe3, 0xe0, 0x96, 0xb0, 0xb8, 0x8d, 0x47, 0x63, 0x18,
	0x6c, 0x3d, 0xd2, 0xd9, 0x8a, 0xc5, 0xf5, 0x66, 0xed, 0xc2, 0xe8, 0x15, 0x34, 0x85, 0xdc, 0x6e,
	0xf9, 0x72, 0xfd, 0xe9, 0xb3, 0xff, 0xd1, 0x97, 0x89, 0xdc, 0xef, 0xa7, 0x1f, 0xd9, 0xce, 0xf3,
	0xd9, 0xd9, 0x49, 0x80, 0x55, 0x5a, 0x34, 0x82, 0x6e, 0x4a, 0x58, 0x16, 0x90, 0xe2, 0x92, 0x8a,
	0x45, 0x6c, 0xb5, 0x94, 0x8a, 0x3b, 0xd8, 0xe8, 0x2b, 0x68, 0x2a, 0x6b, 0x30, 0xe0, 0x96, 0x6c,
	0xee, 0xa1, 0x0e, 0xb4, 0x7c, 0xec, 0xfd, 0xf4, 0xb3, 0xa9, 0xa1, 0x1e, 0x74, 0x96, 0xc1, 0x22,
	0x98, 0xbf, 0x58, 0xb8, 0xdf, 0x99, 0xfa, 0xe8, 0x6b, 0x80, 0x4d, 0x3b, 0xe4, 0x4f, 0xcf, 0x0d,
	0x6d, 0xe7, 0x74, 0xe6, 0xda, 0xe6, 0x1e, 0xea, 0x42, 0x1b, 0x3b, 0x4b, 0x07, 0xff, 0xe8, 0xd8,
	0xa6, 0x86, 0xda, 0xd0, 0x5c, 0xfa, 0x5e, 0x60, 0xea, 0x47, 0x33, 0x38, 0x7c, 0x60, 0xe0, 0xd1,
	0x63, 0x38, 0xb8, 0x22, 0x65, 0x48, 0x7f, 0xcf, 0x79, 0x21, 0x68, 0x5c, 0x8d, 0xa8, 0x9c, 0xe3,
	0x36, 0x1e, 0x5c, 0x91, 0xd2, 0xa9, 0x71, 0x19, 0x3b, 0xfa, 0x05, 0x0e, 0x1f, 0x30, 0x37, 0x74,
	0x00, 0xbd, 0xf9, 0x89, 0x77, 0x66, 0x87, 0xb2, 0xea, 0x62, 0xee, 0x98, 0x7b, 0xf2, 0x22, 0xb3,
	0xf9, 0xdc, 0x3b, 0x73, 0x03, 0x53, 0x47, 0x1f, 0x83, 0xe9, 0x9f, 0xcc, 0xa4, 0xc4, 0x65, 0xb0,
	0x70, 0x67, 0xc1, 0xc2, 0x73, 0xcd, 0xc6, 0xa8, 0xd9, 0xd6, 0x4c, 0xed, 0x71, 0xfb, 0xa5, 0x87,
	0x7f, 0x78, 0x7e, 0xe2, 0xbd, 0x3c, 0xee, 0x81, 0x51, 0x79, 0x88, 0x92, 0xf0, 0xfd, 0x7e, 0xfb,
	0xdf, 0x86, 0xf9, 0xb6, 0x81, 0x7b, 0x6f, 0x78, 0xf1, 0xfa, 0x22, 0xe1, 0x6f, 0x14, 0x3c, 0xfa,
	0x4b, 0x03, 0x63, 0xce, 0x4b, 0x21, 0xf5, 0x48, 0x83, 0xae, 0xcc, 0x54, 0x5b, 0x9b, 0xe9, 0x97,
	0x30, 0x20, 0x79, 0x9e, 0x30, 0x5a, 0xca, 0xad, 0xbc, 0x60, 0x09, 0x55, 0x4e, 0xdb, 0xc6, 0xfd,
	0x1a, 0xf6, 0x2b, 0x14, 0xcd, 0xd6, 0xe5, 0x54, 0xf3, 0x1b, 0xca, 0xda, 0x87, 0xef, 0xf2, 0xba,
	0x7b, 0x46, 0xbe, 0xbb, 0x78, 0xfa, 0xdd, 0xc5, 0x43, 0xd0, 0x8c, 0x78, 0x29, 0xac, 0xd6, 0x50,
	0x1f, 0xeb, 0x58, 0x9d, 0x25, 0x65, 0x55, 0x92, 0x4b, 0x65, 0x86, 0xd4, 0xda, 0x1f, 0x6a, 0xe3,
	0x06, 0xee, 0x28, 0xc4, 0x26, 0x82, 0x1e, 0x3f, 0x81, 0xcf, 0x22, 0x9e, 0x4e, 0xae, 0x53, 0xb1,
	0x2a, 0xce, 0xf9, 0x24, 0x4f, 0x88, 0xb8, 0xe0, 0x45, 0x5a, 0xab, 0x9a, 0xc4, 0x82, 0x1f, 0x77,
	0xd7, 0x1d, 0xb0, 0x03, 0xef, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xec, 0x39, 0xe1, 0x8b, 0xee,
	0x06, 0x00, 0x00,
}
