// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/firestore/v1beta1/document.proto

package firestore

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	latlng "google.golang.org/genproto/googleapis/type/latlng"
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

// A Firestore document.
//
// Must not exceed 1 MiB - 4 bytes.
type Document struct {
	// The resource name of the document, for example
	// `projects/{project_id}/databases/{database_id}/documents/{document_path}`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The document's fields.
	//
	// The map keys represent field names.
	//
	// A simple field name contains only characters `a` to `z`, `A` to `Z`,
	// `0` to `9`, or `_`, and must not start with `0` to `9`. For example,
	// `foo_bar_17`.
	//
	// Field names matching the regular expression `__.*__` are reserved. Reserved
	// field names are forbidden except in certain documented contexts. The map
	// keys, represented as UTF-8, must not exceed 1,500 bytes and cannot be
	// empty.
	//
	// Field paths may be used in other contexts to refer to structured fields
	// defined here. For `map_value`, the field path is represented by the simple
	// or quoted field names of the containing fields, delimited by `.`. For
	// example, the structured field
	// `"foo" : { map_value: { "x&y" : { string_value: "hello" }}}` would be
	// represented by the field path `foo.x&y`.
	//
	// Within a field path, a quoted field name starts and ends with `` ` `` and
	// may contain any character. Some characters, including `` ` ``, must be
	// escaped using a `\`. For example, `` `x&y` `` represents `x&y` and
	// `` `bak\`tik` `` represents `` bak`tik ``.
	Fields map[string]*Value `protobuf:"bytes,2,rep,name=fields,proto3" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Output only. The time at which the document was created.
	//
	// This value increases monotonically when a document is deleted then
	// recreated. It can also be compared to values from other documents and
	// the `read_time` of a query.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. The time at which the document was last changed.
	//
	// This value is initially set to the `create_time` then increases
	// monotonically with each change to the document. It can also be
	// compared to values from other documents and the `read_time` of a query.
	UpdateTime           *timestamp.Timestamp `protobuf:"bytes,4,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Document) Reset()         { *m = Document{} }
func (m *Document) String() string { return proto.CompactTextString(m) }
func (*Document) ProtoMessage()    {}
func (*Document) Descriptor() ([]byte, []int) {
	return fileDescriptor_1522b475188e04d0, []int{0}
}

func (m *Document) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Document.Unmarshal(m, b)
}
func (m *Document) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Document.Marshal(b, m, deterministic)
}
func (m *Document) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Document.Merge(m, src)
}
func (m *Document) XXX_Size() int {
	return xxx_messageInfo_Document.Size(m)
}
func (m *Document) XXX_DiscardUnknown() {
	xxx_messageInfo_Document.DiscardUnknown(m)
}

var xxx_messageInfo_Document proto.InternalMessageInfo

func (m *Document) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Document) GetFields() map[string]*Value {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *Document) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Document) GetUpdateTime() *timestamp.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

// A message that can hold any of the supported value types.
type Value struct {
	// Must have a value set.
	//
	// Types that are valid to be assigned to ValueType:
	//	*Value_NullValue
	//	*Value_BooleanValue
	//	*Value_IntegerValue
	//	*Value_DoubleValue
	//	*Value_TimestampValue
	//	*Value_StringValue
	//	*Value_BytesValue
	//	*Value_ReferenceValue
	//	*Value_GeoPointValue
	//	*Value_ArrayValue
	//	*Value_MapValue
	ValueType            isValue_ValueType `protobuf_oneof:"value_type"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Value) Reset()         { *m = Value{} }
func (m *Value) String() string { return proto.CompactTextString(m) }
func (*Value) ProtoMessage()    {}
func (*Value) Descriptor() ([]byte, []int) {
	return fileDescriptor_1522b475188e04d0, []int{1}
}

func (m *Value) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Value.Unmarshal(m, b)
}
func (m *Value) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Value.Marshal(b, m, deterministic)
}
func (m *Value) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Value.Merge(m, src)
}
func (m *Value) XXX_Size() int {
	return xxx_messageInfo_Value.Size(m)
}
func (m *Value) XXX_DiscardUnknown() {
	xxx_messageInfo_Value.DiscardUnknown(m)
}

var xxx_messageInfo_Value proto.InternalMessageInfo

type isValue_ValueType interface {
	isValue_ValueType()
}

type Value_NullValue struct {
	NullValue _struct.NullValue `protobuf:"varint,11,opt,name=null_value,json=nullValue,proto3,enum=google.protobuf.NullValue,oneof"`
}

type Value_BooleanValue struct {
	BooleanValue bool `protobuf:"varint,1,opt,name=boolean_value,json=booleanValue,proto3,oneof"`
}

type Value_IntegerValue struct {
	IntegerValue int64 `protobuf:"varint,2,opt,name=integer_value,json=integerValue,proto3,oneof"`
}

type Value_DoubleValue struct {
	DoubleValue float64 `protobuf:"fixed64,3,opt,name=double_value,json=doubleValue,proto3,oneof"`
}

type Value_TimestampValue struct {
	TimestampValue *timestamp.Timestamp `protobuf:"bytes,10,opt,name=timestamp_value,json=timestampValue,proto3,oneof"`
}

type Value_StringValue struct {
	StringValue string `protobuf:"bytes,17,opt,name=string_value,json=stringValue,proto3,oneof"`
}

type Value_BytesValue struct {
	BytesValue []byte `protobuf:"bytes,18,opt,name=bytes_value,json=bytesValue,proto3,oneof"`
}

type Value_ReferenceValue struct {
	ReferenceValue string `protobuf:"bytes,5,opt,name=reference_value,json=referenceValue,proto3,oneof"`
}

type Value_GeoPointValue struct {
	GeoPointValue *latlng.LatLng `protobuf:"bytes,8,opt,name=geo_point_value,json=geoPointValue,proto3,oneof"`
}

type Value_ArrayValue struct {
	ArrayValue *ArrayValue `protobuf:"bytes,9,opt,name=array_value,json=arrayValue,proto3,oneof"`
}

type Value_MapValue struct {
	MapValue *MapValue `protobuf:"bytes,6,opt,name=map_value,json=mapValue,proto3,oneof"`
}

func (*Value_NullValue) isValue_ValueType() {}

func (*Value_BooleanValue) isValue_ValueType() {}

func (*Value_IntegerValue) isValue_ValueType() {}

func (*Value_DoubleValue) isValue_ValueType() {}

func (*Value_TimestampValue) isValue_ValueType() {}

func (*Value_StringValue) isValue_ValueType() {}

func (*Value_BytesValue) isValue_ValueType() {}

func (*Value_ReferenceValue) isValue_ValueType() {}

func (*Value_GeoPointValue) isValue_ValueType() {}

func (*Value_ArrayValue) isValue_ValueType() {}

func (*Value_MapValue) isValue_ValueType() {}

func (m *Value) GetValueType() isValue_ValueType {
	if m != nil {
		return m.ValueType
	}
	return nil
}

func (m *Value) GetNullValue() _struct.NullValue {
	if x, ok := m.GetValueType().(*Value_NullValue); ok {
		return x.NullValue
	}
	return _struct.NullValue_NULL_VALUE
}

func (m *Value) GetBooleanValue() bool {
	if x, ok := m.GetValueType().(*Value_BooleanValue); ok {
		return x.BooleanValue
	}
	return false
}

func (m *Value) GetIntegerValue() int64 {
	if x, ok := m.GetValueType().(*Value_IntegerValue); ok {
		return x.IntegerValue
	}
	return 0
}

func (m *Value) GetDoubleValue() float64 {
	if x, ok := m.GetValueType().(*Value_DoubleValue); ok {
		return x.DoubleValue
	}
	return 0
}

func (m *Value) GetTimestampValue() *timestamp.Timestamp {
	if x, ok := m.GetValueType().(*Value_TimestampValue); ok {
		return x.TimestampValue
	}
	return nil
}

func (m *Value) GetStringValue() string {
	if x, ok := m.GetValueType().(*Value_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (m *Value) GetBytesValue() []byte {
	if x, ok := m.GetValueType().(*Value_BytesValue); ok {
		return x.BytesValue
	}
	return nil
}

func (m *Value) GetReferenceValue() string {
	if x, ok := m.GetValueType().(*Value_ReferenceValue); ok {
		return x.ReferenceValue
	}
	return ""
}

func (m *Value) GetGeoPointValue() *latlng.LatLng {
	if x, ok := m.GetValueType().(*Value_GeoPointValue); ok {
		return x.GeoPointValue
	}
	return nil
}

func (m *Value) GetArrayValue() *ArrayValue {
	if x, ok := m.GetValueType().(*Value_ArrayValue); ok {
		return x.ArrayValue
	}
	return nil
}

func (m *Value) GetMapValue() *MapValue {
	if x, ok := m.GetValueType().(*Value_MapValue); ok {
		return x.MapValue
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Value) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Value_NullValue)(nil),
		(*Value_BooleanValue)(nil),
		(*Value_IntegerValue)(nil),
		(*Value_DoubleValue)(nil),
		(*Value_TimestampValue)(nil),
		(*Value_StringValue)(nil),
		(*Value_BytesValue)(nil),
		(*Value_ReferenceValue)(nil),
		(*Value_GeoPointValue)(nil),
		(*Value_ArrayValue)(nil),
		(*Value_MapValue)(nil),
	}
}

// An array value.
type ArrayValue struct {
	// Values in the array.
	Values               []*Value `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArrayValue) Reset()         { *m = ArrayValue{} }
func (m *ArrayValue) String() string { return proto.CompactTextString(m) }
func (*ArrayValue) ProtoMessage()    {}
func (*ArrayValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_1522b475188e04d0, []int{2}
}

func (m *ArrayValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArrayValue.Unmarshal(m, b)
}
func (m *ArrayValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArrayValue.Marshal(b, m, deterministic)
}
func (m *ArrayValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArrayValue.Merge(m, src)
}
func (m *ArrayValue) XXX_Size() int {
	return xxx_messageInfo_ArrayValue.Size(m)
}
func (m *ArrayValue) XXX_DiscardUnknown() {
	xxx_messageInfo_ArrayValue.DiscardUnknown(m)
}

var xxx_messageInfo_ArrayValue proto.InternalMessageInfo

func (m *ArrayValue) GetValues() []*Value {
	if m != nil {
		return m.Values
	}
	return nil
}

// A map value.
type MapValue struct {
	// The map's fields.
	//
	// The map keys represent field names. Field names matching the regular
	// expression `__.*__` are reserved. Reserved field names are forbidden except
	// in certain documented contexts. The map keys, represented as UTF-8, must
	// not exceed 1,500 bytes and cannot be empty.
	Fields               map[string]*Value `protobuf:"bytes,1,rep,name=fields,proto3" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *MapValue) Reset()         { *m = MapValue{} }
func (m *MapValue) String() string { return proto.CompactTextString(m) }
func (*MapValue) ProtoMessage()    {}
func (*MapValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_1522b475188e04d0, []int{3}
}

func (m *MapValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MapValue.Unmarshal(m, b)
}
func (m *MapValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MapValue.Marshal(b, m, deterministic)
}
func (m *MapValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MapValue.Merge(m, src)
}
func (m *MapValue) XXX_Size() int {
	return xxx_messageInfo_MapValue.Size(m)
}
func (m *MapValue) XXX_DiscardUnknown() {
	xxx_messageInfo_MapValue.DiscardUnknown(m)
}

var xxx_messageInfo_MapValue proto.InternalMessageInfo

func (m *MapValue) GetFields() map[string]*Value {
	if m != nil {
		return m.Fields
	}
	return nil
}

func init() {
	proto.RegisterType((*Document)(nil), "google.firestore.v1beta1.Document")
	proto.RegisterMapType((map[string]*Value)(nil), "google.firestore.v1beta1.Document.FieldsEntry")
	proto.RegisterType((*Value)(nil), "google.firestore.v1beta1.Value")
	proto.RegisterType((*ArrayValue)(nil), "google.firestore.v1beta1.ArrayValue")
	proto.RegisterType((*MapValue)(nil), "google.firestore.v1beta1.MapValue")
	proto.RegisterMapType((map[string]*Value)(nil), "google.firestore.v1beta1.MapValue.FieldsEntry")
}

func init() {
	proto.RegisterFile("google/firestore/v1beta1/document.proto", fileDescriptor_1522b475188e04d0)
}

var fileDescriptor_1522b475188e04d0 = []byte{
	// 655 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0xc7, 0xe3, 0x24, 0x8d, 0x92, 0x71, 0xda, 0xfe, 0x7e, 0xe6, 0x12, 0x45, 0x15, 0x0d, 0x01,
	0x44, 0xb8, 0xd8, 0x6a, 0x11, 0x02, 0x51, 0x71, 0x68, 0x4a, 0xd3, 0x1c, 0x0a, 0xaa, 0x0c, 0xea,
	0xa1, 0xaa, 0x14, 0xad, 0x93, 0xcd, 0xca, 0x62, 0xbd, 0x6b, 0xad, 0xd7, 0x95, 0xf2, 0x3a, 0x1c,
	0x39, 0xf0, 0x02, 0xf0, 0x04, 0x7d, 0x2a, 0xb4, 0xff, 0xdc, 0x0a, 0x1a, 0xe5, 0xc4, 0xcd, 0x9e,
	0xf9, 0x7c, 0xbf, 0x33, 0xb3, 0xde, 0x31, 0xbc, 0x20, 0x9c, 0x13, 0x8a, 0xa3, 0x65, 0x2a, 0x70,
	0x21, 0xb9, 0xc0, 0xd1, 0xcd, 0x41, 0x82, 0x25, 0x3a, 0x88, 0x16, 0x7c, 0x5e, 0x66, 0x98, 0xc9,
	0x30, 0x17, 0x5c, 0xf2, 0xa0, 0x67, 0xc0, 0xb0, 0x02, 0x43, 0x0b, 0xf6, 0xf7, 0xac, 0x85, 0xe6,
	0x92, 0x72, 0x19, 0x15, 0x52, 0x94, 0x73, 0xab, 0xeb, 0xef, 0xff, 0x99, 0x95, 0x69, 0x86, 0x0b,
	0x89, 0xb2, 0xdc, 0x02, 0xd6, 0x38, 0x92, 0xab, 0x1c, 0x47, 0x14, 0x49, 0xca, 0x88, 0xcd, 0x38,
	0x63, 0x94, 0xa7, 0x11, 0x62, 0x8c, 0x4b, 0x24, 0x53, 0xce, 0x0a, 0x93, 0x1d, 0xfe, 0xaa, 0x43,
	0xfb, 0x83, 0xed, 0x31, 0x08, 0xa0, 0xc9, 0x50, 0x86, 0x7b, 0xde, 0xc0, 0x1b, 0x75, 0x62, 0xfd,
	0x1c, 0x4c, 0xa0, 0xb5, 0x4c, 0x31, 0x5d, 0x14, 0xbd, 0xfa, 0xa0, 0x31, 0xf2, 0x0f, 0xc3, 0x70,
	0xdd, 0x08, 0xa1, 0xf3, 0x09, 0x27, 0x5a, 0x70, 0xca, 0xa4, 0x58, 0xc5, 0x56, 0x1d, 0x1c, 0x81,
	0x3f, 0x17, 0x18, 0x49, 0x3c, 0x53, 0xad, 0xf7, 0x1a, 0x03, 0x6f, 0xe4, 0x1f, 0xf6, 0x9d, 0x99,
	0x9b, 0x2b, 0xfc, 0xe2, 0xe6, 0x8a, 0xc1, 0xe0, 0x2a, 0xa0, 0xc4, 0x65, 0xbe, 0xa8, 0xc4, 0xcd,
	0xcd, 0x62, 0x83, 0xab, 0x40, 0xff, 0x0a, 0xfc, 0x7b, 0x0d, 0x05, 0xff, 0x41, 0xe3, 0x2b, 0x5e,
	0xd9, 0x19, 0xd5, 0x63, 0xf0, 0x1a, 0xb6, 0x6e, 0x10, 0x2d, 0x71, 0xaf, 0xae, 0x7d, 0xf7, 0xd7,
	0x4f, 0x78, 0xa9, 0xb0, 0xd8, 0xd0, 0xef, 0xea, 0x6f, 0xbd, 0xe1, 0x6d, 0x13, 0xb6, 0x74, 0x30,
	0x38, 0x02, 0x60, 0x25, 0xa5, 0x33, 0xe3, 0xe4, 0x0f, 0xbc, 0xd1, 0xce, 0x03, 0x1d, 0x7e, 0x2a,
	0x29, 0xd5, 0xfc, 0xb4, 0x16, 0x77, 0x98, 0x7b, 0x09, 0x9e, 0xc3, 0x76, 0xc2, 0x39, 0xc5, 0x88,
	0x59, 0xbd, 0xea, 0xae, 0x3d, 0xad, 0xc5, 0x5d, 0x1b, 0xae, 0xb0, 0x94, 0x49, 0x4c, 0xb0, 0x98,
	0xdd, 0x35, 0xdc, 0x50, 0x98, 0x0d, 0x1b, 0xec, 0x29, 0x74, 0x17, 0xbc, 0x4c, 0x28, 0xb6, 0x94,
	0x3a, 0x6b, 0x6f, 0x5a, 0x8b, 0x7d, 0x13, 0x35, 0xd0, 0x29, 0xec, 0x56, 0x77, 0xc8, 0x72, 0xb0,
	0xe9, 0x58, 0xa7, 0xb5, 0x78, 0xa7, 0x12, 0x55, 0xb5, 0x0a, 0x29, 0x52, 0x46, 0xac, 0xc7, 0xff,
	0xea, 0x58, 0x55, 0x2d, 0x13, 0x35, 0xd0, 0x13, 0xf0, 0x93, 0x95, 0xc4, 0x85, 0x65, 0x82, 0x81,
	0x37, 0xea, 0x4e, 0x6b, 0x31, 0xe8, 0xa0, 0x41, 0x5e, 0xc2, 0xae, 0xc0, 0x4b, 0x2c, 0x30, 0x9b,
	0xbb, 0xb6, 0xb7, 0xac, 0xd5, 0x4e, 0x95, 0x30, 0xe8, 0x7b, 0xd8, 0x25, 0x98, 0xcf, 0x72, 0x9e,
	0x32, 0x69, 0xd1, 0xb6, 0xee, 0xfc, 0x91, 0xeb, 0x5c, 0x2d, 0x41, 0x78, 0x8e, 0xe4, 0x39, 0x23,
	0xd3, 0x5a, 0xbc, 0x4d, 0x30, 0xbf, 0x50, 0xb0, 0x91, 0x9f, 0x81, 0x8f, 0x84, 0x40, 0x2b, 0x2b,
	0xed, 0x68, 0xe9, 0xb3, 0xf5, 0xdf, 0xfc, 0x58, 0xc1, 0xee, 0x9b, 0x01, 0xaa, 0xde, 0x82, 0x63,
	0xe8, 0x64, 0xc8, 0x9d, 0x5d, 0x4b, 0xdb, 0x0c, 0xd7, 0xdb, 0x7c, 0x44, 0xb9, 0x33, 0x69, 0x67,
	0xf6, 0x79, 0xdc, 0x05, 0xd0, 0xf2, 0x99, 0xea, 0x78, 0x78, 0x0a, 0x70, 0x57, 0x2c, 0x78, 0x03,
	0x2d, 0x9d, 0x2b, 0x7a, 0x9e, 0x5e, 0xbc, 0x8d, 0xd7, 0xd2, 0xe2, 0xc3, 0x1f, 0x1e, 0xb4, 0x5d,
	0xb5, 0x7b, 0xeb, 0xeb, 0x6d, 0x5a, 0x5f, 0xa7, 0x79, 0x68, 0x7d, 0xff, 0xe5, 0x12, 0x8d, 0x7f,
	0x7a, 0xb0, 0x37, 0xe7, 0xd9, 0x5a, 0xc5, 0x78, 0xdb, 0xfd, 0x59, 0x2e, 0xd4, 0x95, 0xbc, 0xf0,
	0xae, 0x8e, 0x2d, 0x4a, 0x38, 0x45, 0x8c, 0x84, 0x5c, 0x90, 0x88, 0x60, 0xa6, 0x2f, 0x6c, 0x64,
	0x52, 0x28, 0x4f, 0x8b, 0xbf, 0x7f, 0xc7, 0x47, 0x55, 0xe4, 0x5b, 0xbd, 0x79, 0x76, 0x32, 0xf9,
	0xfc, 0xbd, 0xfe, 0xf8, 0xcc, 0x58, 0x9d, 0x50, 0x5e, 0x2e, 0xc2, 0x49, 0x55, 0xfb, 0xf2, 0x60,
	0xac, 0x14, 0xb7, 0x0e, 0xb8, 0xd6, 0xc0, 0x75, 0x05, 0x5c, 0x5f, 0x1a, 0xcb, 0xa4, 0xa5, 0xcb,
	0xbe, 0xfa, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x00, 0x4b, 0xd9, 0xd6, 0x04, 0x06, 0x00, 0x00,
}
