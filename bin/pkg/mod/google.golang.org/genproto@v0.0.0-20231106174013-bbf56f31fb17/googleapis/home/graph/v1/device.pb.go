// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.9
// source: google/home/graph/v1/device.proto

package graph

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Third-party device definition.
type Device struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Third-party device ID.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Hardware type of the device.
	// See [device
	// types](https://developers.home.google.com/cloud-to-cloud/guides).
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	// Traits supported by the device.
	// See [device
	// traits](https://developers.home.google.com/cloud-to-cloud/traits).
	Traits []string `protobuf:"bytes,3,rep,name=traits,proto3" json:"traits,omitempty"`
	// Names given to this device by your smart home Action.
	Name *DeviceNames `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Indicates whether your smart home Action will report state of this device
	// to Google via
	// [ReportStateAndNotification][google.home.graph.v1.HomeGraphApiService.ReportStateAndNotification].
	WillReportState bool `protobuf:"varint,5,opt,name=will_report_state,json=willReportState,proto3" json:"will_report_state,omitempty"`
	// Suggested name for the room where this device is installed.
	// Google attempts to use this value during user setup.
	RoomHint string `protobuf:"bytes,6,opt,name=room_hint,json=roomHint,proto3" json:"room_hint,omitempty"`
	// Suggested name for the structure where this device is installed.
	// Google attempts to use this value during user setup.
	StructureHint string `protobuf:"bytes,7,opt,name=structure_hint,json=structureHint,proto3" json:"structure_hint,omitempty"`
	// Device manufacturer, model, hardware version, and software version.
	DeviceInfo *DeviceInfo `protobuf:"bytes,8,opt,name=device_info,json=deviceInfo,proto3" json:"device_info,omitempty"`
	// Attributes for the traits supported by the device.
	Attributes *structpb.Struct `protobuf:"bytes,9,opt,name=attributes,proto3" json:"attributes,omitempty"`
	// Custom device attributes stored in Home Graph and provided to your
	// smart home Action in each
	// [QUERY](https://developers.home.google.com/cloud-to-cloud/intents/query)
	// and
	// [EXECUTE](https://developers.home.google.com/cloud-to-cloud/intents/execute)
	// intent.
	// Data in this object has a few constraints: No sensitive information,
	// including but not limited to Personally Identifiable Information.
	CustomData *structpb.Struct `protobuf:"bytes,10,opt,name=custom_data,json=customData,proto3" json:"custom_data,omitempty"`
	// Alternate IDs associated with this device.
	// This is used to identify cloud synced devices enabled for [local
	// fulfillment](https://developers.home.google.com/local-home/overview).
	OtherDeviceIds []*AgentOtherDeviceId `protobuf:"bytes,11,rep,name=other_device_ids,json=otherDeviceIds,proto3" json:"other_device_ids,omitempty"`
	// Indicates whether your smart home Action will report notifications
	// to Google for this device via
	// [ReportStateAndNotification][google.home.graph.v1.HomeGraphApiService.ReportStateAndNotification].
	//
	// If your smart home Action enables users to control device notifications,
	// you should update this field and call
	// [RequestSyncDevices][google.home.graph.v1.HomeGraphApiService.RequestSyncDevices].
	NotificationSupportedByAgent bool `protobuf:"varint,12,opt,name=notification_supported_by_agent,json=notificationSupportedByAgent,proto3" json:"notification_supported_by_agent,omitempty"`
}

func (x *Device) Reset() {
	*x = Device{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_home_graph_v1_device_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Device) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Device) ProtoMessage() {}

func (x *Device) ProtoReflect() protoreflect.Message {
	mi := &file_google_home_graph_v1_device_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Device.ProtoReflect.Descriptor instead.
func (*Device) Descriptor() ([]byte, []int) {
	return file_google_home_graph_v1_device_proto_rawDescGZIP(), []int{0}
}

func (x *Device) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Device) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Device) GetTraits() []string {
	if x != nil {
		return x.Traits
	}
	return nil
}

func (x *Device) GetName() *DeviceNames {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *Device) GetWillReportState() bool {
	if x != nil {
		return x.WillReportState
	}
	return false
}

func (x *Device) GetRoomHint() string {
	if x != nil {
		return x.RoomHint
	}
	return ""
}

func (x *Device) GetStructureHint() string {
	if x != nil {
		return x.StructureHint
	}
	return ""
}

func (x *Device) GetDeviceInfo() *DeviceInfo {
	if x != nil {
		return x.DeviceInfo
	}
	return nil
}

func (x *Device) GetAttributes() *structpb.Struct {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *Device) GetCustomData() *structpb.Struct {
	if x != nil {
		return x.CustomData
	}
	return nil
}

func (x *Device) GetOtherDeviceIds() []*AgentOtherDeviceId {
	if x != nil {
		return x.OtherDeviceIds
	}
	return nil
}

func (x *Device) GetNotificationSupportedByAgent() bool {
	if x != nil {
		return x.NotificationSupportedByAgent
	}
	return false
}

// Identifiers used to describe the device.
type DeviceNames struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Primary name of the device, generally provided by the user.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Additional names provided by the user for the device.
	Nicknames []string `protobuf:"bytes,2,rep,name=nicknames,proto3" json:"nicknames,omitempty"`
	// List of names provided by the manufacturer rather than the user, such as
	// serial numbers, SKUs, etc.
	DefaultNames []string `protobuf:"bytes,3,rep,name=default_names,json=defaultNames,proto3" json:"default_names,omitempty"`
}

func (x *DeviceNames) Reset() {
	*x = DeviceNames{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_home_graph_v1_device_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceNames) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceNames) ProtoMessage() {}

func (x *DeviceNames) ProtoReflect() protoreflect.Message {
	mi := &file_google_home_graph_v1_device_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceNames.ProtoReflect.Descriptor instead.
func (*DeviceNames) Descriptor() ([]byte, []int) {
	return file_google_home_graph_v1_device_proto_rawDescGZIP(), []int{1}
}

func (x *DeviceNames) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DeviceNames) GetNicknames() []string {
	if x != nil {
		return x.Nicknames
	}
	return nil
}

func (x *DeviceNames) GetDefaultNames() []string {
	if x != nil {
		return x.DefaultNames
	}
	return nil
}

// Device information.
type DeviceInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Device manufacturer.
	Manufacturer string `protobuf:"bytes,1,opt,name=manufacturer,proto3" json:"manufacturer,omitempty"`
	// Device model.
	Model string `protobuf:"bytes,2,opt,name=model,proto3" json:"model,omitempty"`
	// Device hardware version.
	HwVersion string `protobuf:"bytes,3,opt,name=hw_version,json=hwVersion,proto3" json:"hw_version,omitempty"`
	// Device software version.
	SwVersion string `protobuf:"bytes,4,opt,name=sw_version,json=swVersion,proto3" json:"sw_version,omitempty"`
}

func (x *DeviceInfo) Reset() {
	*x = DeviceInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_home_graph_v1_device_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceInfo) ProtoMessage() {}

func (x *DeviceInfo) ProtoReflect() protoreflect.Message {
	mi := &file_google_home_graph_v1_device_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceInfo.ProtoReflect.Descriptor instead.
func (*DeviceInfo) Descriptor() ([]byte, []int) {
	return file_google_home_graph_v1_device_proto_rawDescGZIP(), []int{2}
}

func (x *DeviceInfo) GetManufacturer() string {
	if x != nil {
		return x.Manufacturer
	}
	return ""
}

func (x *DeviceInfo) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *DeviceInfo) GetHwVersion() string {
	if x != nil {
		return x.HwVersion
	}
	return ""
}

func (x *DeviceInfo) GetSwVersion() string {
	if x != nil {
		return x.SwVersion
	}
	return ""
}

// Alternate third-party device ID.
type AgentOtherDeviceId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Project ID for your smart home Action.
	AgentId string `protobuf:"bytes,1,opt,name=agent_id,json=agentId,proto3" json:"agent_id,omitempty"`
	// Unique third-party device ID.
	DeviceId string `protobuf:"bytes,2,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
}

func (x *AgentOtherDeviceId) Reset() {
	*x = AgentOtherDeviceId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_home_graph_v1_device_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentOtherDeviceId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentOtherDeviceId) ProtoMessage() {}

func (x *AgentOtherDeviceId) ProtoReflect() protoreflect.Message {
	mi := &file_google_home_graph_v1_device_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentOtherDeviceId.ProtoReflect.Descriptor instead.
func (*AgentOtherDeviceId) Descriptor() ([]byte, []int) {
	return file_google_home_graph_v1_device_proto_rawDescGZIP(), []int{3}
}

func (x *AgentOtherDeviceId) GetAgentId() string {
	if x != nil {
		return x.AgentId
	}
	return ""
}

func (x *AgentOtherDeviceId) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

var File_google_home_graph_v1_device_proto protoreflect.FileDescriptor

var file_google_home_graph_v1_device_proto_rawDesc = []byte{
	0x0a, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x68, 0x6f, 0x6d, 0x65, 0x2f, 0x67, 0x72,
	0x61, 0x70, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x14, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x68, 0x6f, 0x6d, 0x65,
	0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbc, 0x04, 0x0a, 0x06, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x12, 0x35,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x77, 0x69, 0x6c, 0x6c, 0x5f, 0x72, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0f, 0x77, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x68, 0x69, 0x6e, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x6f, 0x6f, 0x6d, 0x48, 0x69, 0x6e, 0x74, 0x12, 0x25,
	0x0a, 0x0e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x68, 0x69, 0x6e, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x48, 0x69, 0x6e, 0x74, 0x12, 0x41, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x37, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x73, 0x12, 0x38, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52,
	0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x12, 0x52, 0x0a, 0x10, 0x6f,
	0x74, 0x68, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18,
	0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x68,
	0x6f, 0x6d, 0x65, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x52,
	0x0e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x73, 0x12,
	0x45, 0x0a, 0x1f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x5f, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1c, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x42,
	0x79, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x22, 0x64, 0x0a, 0x0b, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x69, 0x63,
	0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x69,
	0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c,
	0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x22, 0x84, 0x01, 0x0a,
	0x0a, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x22, 0x0a, 0x0c, 0x6d,
	0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x12,
	0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x68, 0x77, 0x5f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x68, 0x77, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x77, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x77, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x22, 0x4c, 0x0a, 0x12, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x4f, 0x74, 0x68, 0x65,
	0x72, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49,
	0x64, 0x42, 0x79, 0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x68, 0x6f, 0x6d, 0x65, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x39, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x68, 0x6f, 0x6d, 0x65, 0x2f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2f, 0x76, 0x31, 0x3b,
	0x67, 0x72, 0x61, 0x70, 0x68, 0xca, 0x02, 0x14, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x48,
	0x6f, 0x6d, 0x65, 0x5c, 0x47, 0x72, 0x61, 0x70, 0x68, 0x5c, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_home_graph_v1_device_proto_rawDescOnce sync.Once
	file_google_home_graph_v1_device_proto_rawDescData = file_google_home_graph_v1_device_proto_rawDesc
)

func file_google_home_graph_v1_device_proto_rawDescGZIP() []byte {
	file_google_home_graph_v1_device_proto_rawDescOnce.Do(func() {
		file_google_home_graph_v1_device_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_home_graph_v1_device_proto_rawDescData)
	})
	return file_google_home_graph_v1_device_proto_rawDescData
}

var file_google_home_graph_v1_device_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_home_graph_v1_device_proto_goTypes = []interface{}{
	(*Device)(nil),             // 0: google.home.graph.v1.Device
	(*DeviceNames)(nil),        // 1: google.home.graph.v1.DeviceNames
	(*DeviceInfo)(nil),         // 2: google.home.graph.v1.DeviceInfo
	(*AgentOtherDeviceId)(nil), // 3: google.home.graph.v1.AgentOtherDeviceId
	(*structpb.Struct)(nil),    // 4: google.protobuf.Struct
}
var file_google_home_graph_v1_device_proto_depIdxs = []int32{
	1, // 0: google.home.graph.v1.Device.name:type_name -> google.home.graph.v1.DeviceNames
	2, // 1: google.home.graph.v1.Device.device_info:type_name -> google.home.graph.v1.DeviceInfo
	4, // 2: google.home.graph.v1.Device.attributes:type_name -> google.protobuf.Struct
	4, // 3: google.home.graph.v1.Device.custom_data:type_name -> google.protobuf.Struct
	3, // 4: google.home.graph.v1.Device.other_device_ids:type_name -> google.home.graph.v1.AgentOtherDeviceId
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_google_home_graph_v1_device_proto_init() }
func file_google_home_graph_v1_device_proto_init() {
	if File_google_home_graph_v1_device_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_home_graph_v1_device_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Device); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_home_graph_v1_device_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceNames); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_home_graph_v1_device_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_home_graph_v1_device_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentOtherDeviceId); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_home_graph_v1_device_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_home_graph_v1_device_proto_goTypes,
		DependencyIndexes: file_google_home_graph_v1_device_proto_depIdxs,
		MessageInfos:      file_google_home_graph_v1_device_proto_msgTypes,
	}.Build()
	File_google_home_graph_v1_device_proto = out.File
	file_google_home_graph_v1_device_proto_rawDesc = nil
	file_google_home_graph_v1_device_proto_goTypes = nil
	file_google_home_graph_v1_device_proto_depIdxs = nil
}
