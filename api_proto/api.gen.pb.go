// Code generated by protoc-gen-go.
// source: api_proto/api.proto
// DO NOT EDIT!

/*
Package devtools_buildozer is a generated protocol buffer package.

It is generated from these files:
	api_proto/api.proto

It has these top-level messages:
	Output
	RepeatedString
*/
package devtools_buildozer

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

type Output_Record_Field_ERROR int32

const (
	Output_Record_Field_UNKNOWN Output_Record_Field_ERROR = 0
	Output_Record_Field_MISSING Output_Record_Field_ERROR = 1
)

var Output_Record_Field_ERROR_name = map[int32]string{
	0: "UNKNOWN",
	1: "MISSING",
}
var Output_Record_Field_ERROR_value = map[string]int32{
	"UNKNOWN": 0,
	"MISSING": 1,
}

func (x Output_Record_Field_ERROR) String() string {
	return proto.EnumName(Output_Record_Field_ERROR_name, int32(x))
}
func (Output_Record_Field_ERROR) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0, 0, 0}
}

type Output struct {
	Records []*Output_Record `protobuf:"bytes,1,rep,name=records" json:"records,omitempty"`
}

func (m *Output) Reset()                    { *m = Output{} }
func (m *Output) String() string            { return proto.CompactTextString(m) }
func (*Output) ProtoMessage()               {}
func (*Output) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Output) GetRecords() []*Output_Record {
	if m != nil {
		return m.Records
	}
	return nil
}

type Output_Record struct {
	Fields []*Output_Record_Field `protobuf:"bytes,1,rep,name=fields" json:"fields,omitempty"`
}

func (m *Output_Record) Reset()                    { *m = Output_Record{} }
func (m *Output_Record) String() string            { return proto.CompactTextString(m) }
func (*Output_Record) ProtoMessage()               {}
func (*Output_Record) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *Output_Record) GetFields() []*Output_Record_Field {
	if m != nil {
		return m.Fields
	}
	return nil
}

type Output_Record_Field struct {
	// Types that are valid to be assigned to Value:
	//	*Output_Record_Field_Text
	//	*Output_Record_Field_Number
	//	*Output_Record_Field_Error
	//	*Output_Record_Field_List
	Value isOutput_Record_Field_Value `protobuf_oneof:"value"`
	// Used internally by Buildozer to decide whether a field should be quoted
	// when printing. This does not affect the contents of 'value'.
	QuoteWhenPrinting bool `protobuf:"varint,7,opt,name=quote_when_printing,json=quoteWhenPrinting" json:"quote_when_printing,omitempty"`
}

func (m *Output_Record_Field) Reset()                    { *m = Output_Record_Field{} }
func (m *Output_Record_Field) String() string            { return proto.CompactTextString(m) }
func (*Output_Record_Field) ProtoMessage()               {}
func (*Output_Record_Field) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0, 0} }

type isOutput_Record_Field_Value interface {
	isOutput_Record_Field_Value()
}

type Output_Record_Field_Text struct {
	Text string `protobuf:"bytes,1,opt,name=text,oneof"`
}
type Output_Record_Field_Number struct {
	Number int32 `protobuf:"varint,2,opt,name=number,oneof"`
}
type Output_Record_Field_Error struct {
	Error Output_Record_Field_ERROR `protobuf:"varint,3,opt,name=error,enum=devtools.buildozer.Output_Record_Field_ERROR,oneof"`
}
type Output_Record_Field_List struct {
	List *RepeatedString `protobuf:"bytes,5,opt,name=list,oneof"`
}

func (*Output_Record_Field_Text) isOutput_Record_Field_Value()   {}
func (*Output_Record_Field_Number) isOutput_Record_Field_Value() {}
func (*Output_Record_Field_Error) isOutput_Record_Field_Value()  {}
func (*Output_Record_Field_List) isOutput_Record_Field_Value()   {}

func (m *Output_Record_Field) GetValue() isOutput_Record_Field_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Output_Record_Field) GetText() string {
	if x, ok := m.GetValue().(*Output_Record_Field_Text); ok {
		return x.Text
	}
	return ""
}

func (m *Output_Record_Field) GetNumber() int32 {
	if x, ok := m.GetValue().(*Output_Record_Field_Number); ok {
		return x.Number
	}
	return 0
}

func (m *Output_Record_Field) GetError() Output_Record_Field_ERROR {
	if x, ok := m.GetValue().(*Output_Record_Field_Error); ok {
		return x.Error
	}
	return Output_Record_Field_UNKNOWN
}

func (m *Output_Record_Field) GetList() *RepeatedString {
	if x, ok := m.GetValue().(*Output_Record_Field_List); ok {
		return x.List
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Output_Record_Field) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Output_Record_Field_OneofMarshaler, _Output_Record_Field_OneofUnmarshaler, _Output_Record_Field_OneofSizer, []interface{}{
		(*Output_Record_Field_Text)(nil),
		(*Output_Record_Field_Number)(nil),
		(*Output_Record_Field_Error)(nil),
		(*Output_Record_Field_List)(nil),
	}
}

func _Output_Record_Field_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Output_Record_Field)
	// value
	switch x := m.Value.(type) {
	case *Output_Record_Field_Text:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Text)
	case *Output_Record_Field_Number:
		b.EncodeVarint(2<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Number))
	case *Output_Record_Field_Error:
		b.EncodeVarint(3<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Error))
	case *Output_Record_Field_List:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.List); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Output_Record_Field.Value has unexpected type %T", x)
	}
	return nil
}

func _Output_Record_Field_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Output_Record_Field)
	switch tag {
	case 1: // value.text
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Value = &Output_Record_Field_Text{x}
		return true, err
	case 2: // value.number
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &Output_Record_Field_Number{int32(x)}
		return true, err
	case 3: // value.error
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &Output_Record_Field_Error{Output_Record_Field_ERROR(x)}
		return true, err
	case 5: // value.list
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RepeatedString)
		err := b.DecodeMessage(msg)
		m.Value = &Output_Record_Field_List{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Output_Record_Field_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Output_Record_Field)
	// value
	switch x := m.Value.(type) {
	case *Output_Record_Field_Text:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Text)))
		n += len(x.Text)
	case *Output_Record_Field_Number:
		n += proto.SizeVarint(2<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Number))
	case *Output_Record_Field_Error:
		n += proto.SizeVarint(3<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Error))
	case *Output_Record_Field_List:
		s := proto.Size(x.List)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type RepeatedString struct {
	Strings []string `protobuf:"bytes,1,rep,name=strings" json:"strings,omitempty"`
}

func (m *RepeatedString) Reset()                    { *m = RepeatedString{} }
func (m *RepeatedString) String() string            { return proto.CompactTextString(m) }
func (*RepeatedString) ProtoMessage()               {}
func (*RepeatedString) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*Output)(nil), "devtools.buildozer.Output")
	proto.RegisterType((*Output_Record)(nil), "devtools.buildozer.Output.Record")
	proto.RegisterType((*Output_Record_Field)(nil), "devtools.buildozer.Output.Record.Field")
	proto.RegisterType((*RepeatedString)(nil), "devtools.buildozer.RepeatedString")
	proto.RegisterEnum("devtools.buildozer.Output_Record_Field_ERROR", Output_Record_Field_ERROR_name, Output_Record_Field_ERROR_value)
}

func init() { proto.RegisterFile("api_proto/api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 336 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x91, 0xc1, 0x4a, 0xf3, 0x40,
	0x14, 0x85, 0x33, 0x6d, 0x93, 0xfc, 0xbd, 0x85, 0xd2, 0x7f, 0xea, 0x62, 0xe8, 0x2a, 0xcd, 0xc6,
	0x20, 0x38, 0x42, 0xdd, 0x08, 0x2e, 0x04, 0xa1, 0xda, 0x22, 0xa6, 0x32, 0x45, 0xba, 0x2c, 0xa9,
	0x19, 0xdb, 0x81, 0x98, 0x89, 0x93, 0x49, 0x15, 0x9f, 0xc3, 0x87, 0xf0, 0xb1, 0x7c, 0x14, 0xc9,
	0x24, 0x15, 0x44, 0x41, 0x77, 0xf7, 0xdc, 0x7b, 0xbe, 0x73, 0x42, 0x06, 0xfa, 0x51, 0x26, 0x96,
	0x99, 0x92, 0x5a, 0x1e, 0x45, 0x99, 0xa0, 0x66, 0xc2, 0x38, 0xe6, 0x5b, 0x2d, 0x65, 0x92, 0xd3,
	0x55, 0x21, 0x92, 0x58, 0xbe, 0x70, 0xe5, 0xbf, 0x35, 0xc1, 0x99, 0x15, 0x3a, 0x2b, 0x34, 0x3e,
	0x05, 0x57, 0xf1, 0x3b, 0xa9, 0xe2, 0x9c, 0x20, 0xaf, 0x19, 0x74, 0x46, 0x43, 0xfa, 0x1d, 0xa0,
	0x95, 0x99, 0x32, 0xe3, 0x64, 0x3b, 0x62, 0xf0, 0xde, 0x00, 0xa7, 0xda, 0xe1, 0x33, 0x70, 0xee,
	0x05, 0x4f, 0x3e, 0x63, 0xf6, 0x7f, 0x8d, 0xa1, 0x17, 0xa5, 0x9f, 0xd5, 0xd8, 0xe0, 0xb5, 0x01,
	0xb6, 0xd9, 0xe0, 0x3d, 0x68, 0x69, 0xfe, 0xac, 0x09, 0xf2, 0x50, 0xd0, 0x9e, 0x58, 0xcc, 0x28,
	0x4c, 0xc0, 0x49, 0x8b, 0x87, 0x15, 0x57, 0xa4, 0xe1, 0xa1, 0xc0, 0x9e, 0x58, 0xac, 0xd6, 0x78,
	0x0c, 0x36, 0x57, 0x4a, 0x2a, 0xd2, 0xf4, 0x50, 0xd0, 0x1d, 0x1d, 0xfe, 0xb1, 0x99, 0x8e, 0x19,
	0x9b, 0xb1, 0x89, 0xc5, 0x2a, 0x1a, 0x9f, 0x40, 0x2b, 0x11, 0xb9, 0x26, 0xb6, 0x87, 0x82, 0xce,
	0xc8, 0xff, 0x29, 0x85, 0xf1, 0x8c, 0x47, 0x9a, 0xc7, 0x73, 0xad, 0x44, 0xba, 0x2e, 0x3f, 0xad,
	0x24, 0x30, 0x85, 0xfe, 0x63, 0x21, 0x35, 0x5f, 0x3e, 0x6d, 0x78, 0xba, 0xcc, 0x94, 0x48, 0xb5,
	0x48, 0xd7, 0xc4, 0xf5, 0x50, 0xf0, 0x8f, 0xfd, 0x37, 0xa7, 0xc5, 0x86, 0xa7, 0x37, 0xf5, 0xc1,
	0x1f, 0x82, 0x6d, 0xba, 0x71, 0x07, 0xdc, 0xdb, 0xf0, 0x2a, 0x9c, 0x2d, 0xc2, 0x9e, 0x55, 0x8a,
	0xeb, 0xe9, 0x7c, 0x3e, 0x0d, 0x2f, 0x7b, 0xe8, 0xdc, 0x05, 0x7b, 0x1b, 0x25, 0x05, 0xf7, 0x0f,
	0xa0, 0xfb, 0xb5, 0x15, 0x13, 0x70, 0x73, 0x33, 0x55, 0xbf, 0xba, 0xcd, 0x76, 0x72, 0xe5, 0x98,
	0x17, 0x3f, 0xfe, 0x08, 0x00, 0x00, 0xff, 0xff, 0x8d, 0x62, 0x58, 0xc4, 0x08, 0x02, 0x00, 0x00,
}