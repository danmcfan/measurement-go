// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.0
// source: internal/measurement.proto

package grpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Measurement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude    float64 `protobuf:"fixed64,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude   float64 `protobuf:"fixed64,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Altitude    float64 `protobuf:"fixed64,3,opt,name=altitude,proto3" json:"altitude,omitempty"`
	Temperature float64 `protobuf:"fixed64,4,opt,name=temperature,proto3" json:"temperature,omitempty"`
	Pressure    float64 `protobuf:"fixed64,5,opt,name=pressure,proto3" json:"pressure,omitempty"`
	Humidity    float64 `protobuf:"fixed64,6,opt,name=humidity,proto3" json:"humidity,omitempty"`
	Timestamp   int64   `protobuf:"varint,7,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	DeviceId    string  `protobuf:"bytes,8,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
}

func (x *Measurement) Reset() {
	*x = Measurement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_measurement_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Measurement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Measurement) ProtoMessage() {}

func (x *Measurement) ProtoReflect() protoreflect.Message {
	mi := &file_internal_measurement_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Measurement.ProtoReflect.Descriptor instead.
func (*Measurement) Descriptor() ([]byte, []int) {
	return file_internal_measurement_proto_rawDescGZIP(), []int{0}
}

func (x *Measurement) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Measurement) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *Measurement) GetAltitude() float64 {
	if x != nil {
		return x.Altitude
	}
	return 0
}

func (x *Measurement) GetTemperature() float64 {
	if x != nil {
		return x.Temperature
	}
	return 0
}

func (x *Measurement) GetPressure() float64 {
	if x != nil {
		return x.Pressure
	}
	return 0
}

func (x *Measurement) GetHumidity() float64 {
	if x != nil {
		return x.Humidity
	}
	return 0
}

func (x *Measurement) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Measurement) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

type MeasurementResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *MeasurementResponse) Reset() {
	*x = MeasurementResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_measurement_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MeasurementResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MeasurementResponse) ProtoMessage() {}

func (x *MeasurementResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_measurement_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MeasurementResponse.ProtoReflect.Descriptor instead.
func (*MeasurementResponse) Descriptor() ([]byte, []int) {
	return file_internal_measurement_proto_rawDescGZIP(), []int{1}
}

func (x *MeasurementResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *MeasurementResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_internal_measurement_proto protoreflect.FileDescriptor

var file_internal_measurement_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6d, 0x65, 0x61, 0x73, 0x75,
	0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x72,
	0x70, 0x63, 0x22, 0xf8, 0x01, 0x0a, 0x0b, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x61, 0x6c, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08,
	0x61, 0x6c, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x74,
	0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x75, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x75, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x75, 0x6d, 0x69, 0x64, 0x69,
	0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x68, 0x75, 0x6d, 0x69, 0x64, 0x69,
	0x74, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x22, 0x49, 0x0a,
	0x13, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x57, 0x0a, 0x12, 0x4d, 0x65, 0x61, 0x73,
	0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x41,
	0x0a, 0x0f, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x11, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x19, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4d, 0x65, 0x61, 0x73,
	0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_measurement_proto_rawDescOnce sync.Once
	file_internal_measurement_proto_rawDescData = file_internal_measurement_proto_rawDesc
)

func file_internal_measurement_proto_rawDescGZIP() []byte {
	file_internal_measurement_proto_rawDescOnce.Do(func() {
		file_internal_measurement_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_measurement_proto_rawDescData)
	})
	return file_internal_measurement_proto_rawDescData
}

var file_internal_measurement_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_internal_measurement_proto_goTypes = []any{
	(*Measurement)(nil),         // 0: grpc.Measurement
	(*MeasurementResponse)(nil), // 1: grpc.MeasurementResponse
}
var file_internal_measurement_proto_depIdxs = []int32{
	0, // 0: grpc.MeasurementService.SendMeasurement:input_type -> grpc.Measurement
	1, // 1: grpc.MeasurementService.SendMeasurement:output_type -> grpc.MeasurementResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_measurement_proto_init() }
func file_internal_measurement_proto_init() {
	if File_internal_measurement_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_measurement_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Measurement); i {
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
		file_internal_measurement_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*MeasurementResponse); i {
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
			RawDescriptor: file_internal_measurement_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_measurement_proto_goTypes,
		DependencyIndexes: file_internal_measurement_proto_depIdxs,
		MessageInfos:      file_internal_measurement_proto_msgTypes,
	}.Build()
	File_internal_measurement_proto = out.File
	file_internal_measurement_proto_rawDesc = nil
	file_internal_measurement_proto_goTypes = nil
	file_internal_measurement_proto_depIdxs = nil
}
