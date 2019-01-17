// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/day_of_week.proto

package enums // import "google.golang.org/genproto/googleapis/ads/googleads/v0/enums"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Enumerates days of the week, e.g., "Monday".
type DayOfWeekEnum_DayOfWeek int32

const (
	// Not specified.
	DayOfWeekEnum_UNSPECIFIED DayOfWeekEnum_DayOfWeek = 0
	// The value is unknown in this version.
	DayOfWeekEnum_UNKNOWN DayOfWeekEnum_DayOfWeek = 1
	// Monday.
	DayOfWeekEnum_MONDAY DayOfWeekEnum_DayOfWeek = 2
	// Tuesday.
	DayOfWeekEnum_TUESDAY DayOfWeekEnum_DayOfWeek = 3
	// Wednesday.
	DayOfWeekEnum_WEDNESDAY DayOfWeekEnum_DayOfWeek = 4
	// Thursday.
	DayOfWeekEnum_THURSDAY DayOfWeekEnum_DayOfWeek = 5
	// Friday.
	DayOfWeekEnum_FRIDAY DayOfWeekEnum_DayOfWeek = 6
	// Saturday.
	DayOfWeekEnum_SATURDAY DayOfWeekEnum_DayOfWeek = 7
	// Sunday.
	DayOfWeekEnum_SUNDAY DayOfWeekEnum_DayOfWeek = 8
)

var DayOfWeekEnum_DayOfWeek_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "MONDAY",
	3: "TUESDAY",
	4: "WEDNESDAY",
	5: "THURSDAY",
	6: "FRIDAY",
	7: "SATURDAY",
	8: "SUNDAY",
}
var DayOfWeekEnum_DayOfWeek_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"MONDAY":      2,
	"TUESDAY":     3,
	"WEDNESDAY":   4,
	"THURSDAY":    5,
	"FRIDAY":      6,
	"SATURDAY":    7,
	"SUNDAY":      8,
}

func (x DayOfWeekEnum_DayOfWeek) String() string {
	return proto.EnumName(DayOfWeekEnum_DayOfWeek_name, int32(x))
}
func (DayOfWeekEnum_DayOfWeek) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_day_of_week_8b74b2709672516d, []int{0, 0}
}

// Container for enumeration of days of the week, e.g., "Monday".
type DayOfWeekEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DayOfWeekEnum) Reset()         { *m = DayOfWeekEnum{} }
func (m *DayOfWeekEnum) String() string { return proto.CompactTextString(m) }
func (*DayOfWeekEnum) ProtoMessage()    {}
func (*DayOfWeekEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_day_of_week_8b74b2709672516d, []int{0}
}
func (m *DayOfWeekEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DayOfWeekEnum.Unmarshal(m, b)
}
func (m *DayOfWeekEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DayOfWeekEnum.Marshal(b, m, deterministic)
}
func (dst *DayOfWeekEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DayOfWeekEnum.Merge(dst, src)
}
func (m *DayOfWeekEnum) XXX_Size() int {
	return xxx_messageInfo_DayOfWeekEnum.Size(m)
}
func (m *DayOfWeekEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_DayOfWeekEnum.DiscardUnknown(m)
}

var xxx_messageInfo_DayOfWeekEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*DayOfWeekEnum)(nil), "google.ads.googleads.v0.enums.DayOfWeekEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.DayOfWeekEnum_DayOfWeek", DayOfWeekEnum_DayOfWeek_name, DayOfWeekEnum_DayOfWeek_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/day_of_week.proto", fileDescriptor_day_of_week_8b74b2709672516d)
}

var fileDescriptor_day_of_week_8b74b2709672516d = []byte{
	// 303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4f, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x4c, 0x29, 0x86, 0x32, 0x41, 0xac, 0x32, 0x03, 0xfd, 0xd4, 0xbc, 0xd2,
	0xdc, 0x62, 0xfd, 0x94, 0xc4, 0xca, 0xf8, 0xfc, 0xb4, 0xf8, 0xf2, 0xd4, 0xd4, 0x6c, 0xbd, 0x82,
	0xa2, 0xfc, 0x92, 0x7c, 0x21, 0x59, 0x88, 0x2a, 0xbd, 0xc4, 0x94, 0x62, 0x3d, 0xb8, 0x06, 0xbd,
	0x32, 0x03, 0x3d, 0xb0, 0x06, 0xa5, 0xe9, 0x8c, 0x5c, 0xbc, 0x2e, 0x89, 0x95, 0xfe, 0x69, 0xe1,
	0xa9, 0xa9, 0xd9, 0xae, 0x79, 0xa5, 0xb9, 0x4a, 0xad, 0x8c, 0x5c, 0x9c, 0x70, 0x11, 0x21, 0x7e,
	0x2e, 0xee, 0x50, 0xbf, 0xe0, 0x00, 0x57, 0x67, 0x4f, 0x37, 0x4f, 0x57, 0x17, 0x01, 0x06, 0x21,
	0x6e, 0x2e, 0xf6, 0x50, 0x3f, 0x6f, 0x3f, 0xff, 0x70, 0x3f, 0x01, 0x46, 0x21, 0x2e, 0x2e, 0x36,
	0x5f, 0x7f, 0x3f, 0x17, 0xc7, 0x48, 0x01, 0x26, 0x90, 0x44, 0x48, 0xa8, 0x6b, 0x30, 0x88, 0xc3,
	0x2c, 0xc4, 0xcb, 0xc5, 0x19, 0xee, 0xea, 0xe2, 0x07, 0xe1, 0xb2, 0x08, 0xf1, 0x70, 0x71, 0x84,
	0x78, 0x84, 0x06, 0x81, 0x79, 0xac, 0x20, 0x5d, 0x6e, 0x41, 0x9e, 0x20, 0x36, 0x1b, 0x48, 0x26,
	0xd8, 0x31, 0x24, 0x34, 0x08, 0xc4, 0x63, 0x07, 0xc9, 0x04, 0x87, 0x82, 0xcd, 0xe3, 0x70, 0xda,
	0xcf, 0xc8, 0xa5, 0x98, 0x9c, 0x9f, 0xab, 0x87, 0xd7, 0xfd, 0x4e, 0x7c, 0x70, 0xa7, 0x06, 0x80,
	0xbc, 0x1b, 0xc0, 0x18, 0xe5, 0x04, 0xd5, 0x90, 0x9e, 0x9f, 0x93, 0x98, 0x97, 0xae, 0x97, 0x5f,
	0x94, 0xae, 0x9f, 0x9e, 0x9a, 0x07, 0x0e, 0x0c, 0x58, 0x88, 0x15, 0x64, 0x16, 0xe3, 0x08, 0x40,
	0x6b, 0x30, 0xb9, 0x88, 0x89, 0xd9, 0xdd, 0xd1, 0x71, 0x15, 0x93, 0xac, 0x3b, 0xc4, 0x28, 0xc7,
	0x94, 0x62, 0x3d, 0x08, 0x13, 0xc4, 0x0a, 0x33, 0xd0, 0x03, 0x05, 0x54, 0xf1, 0x29, 0x98, 0x7c,
	0x8c, 0x63, 0x4a, 0x71, 0x0c, 0x5c, 0x3e, 0x26, 0xcc, 0x20, 0x06, 0x2c, 0x9f, 0xc4, 0x06, 0xb6,
	0xd4, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xeb, 0xea, 0x27, 0x3e, 0xb4, 0x01, 0x00, 0x00,
}