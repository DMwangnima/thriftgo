// Code generated by thriftgo (0.3.2-option-exp). DO NOT EDIT.

package option_gen

import (
	"reflect"

	"github.com/cloudwego/thriftgo/thrift_reflection"
)

// IDL Name: test
// IDL Path: ../option_idl/test.thrift

var file_test_thrift_go_types = []interface{}{
	(*IDCard)(nil),         // Struct 0: option_gen.IDCard
	(*Person)(nil),         // Struct 1: option_gen.Person
	(*_FieldOptions)(nil),  // Struct 2: option_gen._FieldOptions
	(*_StructOptions)(nil), // Struct 3: option_gen._StructOptions
	(*TinyStruct)(nil),     // Struct 4: option_gen.TinyStruct
	(*MyEnum)(nil),         // Enum 0: option_gen.MyEnum
}
var file_test_thrift *thrift_reflection.FileDescriptor
var file_idl_test_rawDesc = []byte{
	0x1f, 0x8b, 0x8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xff, 0xcc, 0x57, 0x5d, 0x8f, 0xda, 0x46,
	0x14, 0xbd, 0x98, 0x8f, 0x6c, 0x60, 0x29, 0x89, 0xd2, 0x8f, 0xa4, 0x1f, 0x92, 0x65, 0xb5, 0xca,
	0x43, 0x11, 0x2b, 0x1b, 0x16, 0x6d, 0xe6, 0xa9, 0xdd, 0x6e, 0x5b, 0xad, 0xd4, 0x6d, 0xab, 0xa6,
	0xaa, 0x22, 0xa1, 0x15, 0x1a, 0xe3, 0x1, 0x46, 0xb1, 0xc7, 0xd4, 0x1e, 0x50, 0x11, 0xe2, 0xbf,
	0x57, 0x33, 0x1e, 0x83, 0x8d, 0x87, 0xad, 0x89, 0x12, 0xa9, 0xbc, 0xd8, 0x8c, 0xcf, 0x9d, 0x7b,
	0xce, 0x9d, 0xe3, 0x19, 0xdf, 0x16, 0x54, 0x0, 0xe0, 0x45, 0xaf, 0x77, 0x11, 0x2e, 0x38, 0xd,
	0xd9, 0x98, 0x7a, 0xfe, 0x5, 0x27, 0x31, 0xef, 0xf1, 0x79, 0x44, 0xa7, 0xbc, 0xd, 0x46, 0xab,
	0x5, 0x0, 0x6, 0x0, 0x34, 0x8, 0xe3, 0x94, 0xaf, 0x1, 0xa0, 0x97, 0xc7, 0x63, 0xc6, 0x42,
	0x8e, 0xc5, 0xdf, 0xf8, 0x22, 0xc1, 0xa8, 0x8b, 0x9a, 0x4, 0x0, 0x9a, 0x2b, 0xec, 0x53, 0x4f,
	0x62, 0x0, 0x60, 0x78, 0x3c, 0x7e, 0x8f, 0xcb, 0xdc, 0xee, 0xc8, 0x54, 0x25, 0x19, 0x41, 0xd9,
	0x98, 0x85, 0x62, 0x5a, 0x35, 0xcb, 0x8c, 0xb0, 0xe, 0xd4, 0xce, 0xc5, 0xb3, 0xff, 0x52, 0xd4,
	0x92, 0x5a, 0x1e, 0xdf, 0xad, 0x5f, 0x93, 0x68, 0x45, 0x27, 0xa4, 0x3, 0x55, 0x11, 0x67, 0x94,
	0x8b, 0x33, 0xee, 0xec, 0x73, 0xa8, 0x96, 0xc3, 0x36, 0x62, 0x1e, 0x51, 0x36, 0x3, 0xc5, 0xc,
	0xda, 0x50, 0x6f, 0x75, 0x14, 0xfd, 0x4f, 0x33, 0xe2, 0x2, 0xc2, 0xe7, 0xa1, 0x37, 0xa6, 0x6c,
	0x1a, 0xa6, 0xea, 0x6e, 0x36, 0x4d, 0x33, 0xf3, 0x63, 0x38, 0x20, 0xc8, 0xbc, 0x93, 0xb8, 0x5b,
	0x36, 0xd, 0x7f, 0xc5, 0x1, 0xc9, 0x3, 0x96, 0x81, 0x4b, 0x22, 0x64, 0x5e, 0x5e, 0x5e, 0xee,
	0xc6, 0xb7, 0x2d, 0x68, 0x88, 0xb4, 0x1d, 0x78, 0x24, 0xd3, 0x1b, 0x70, 0x6, 0x50, 0x56, 0xa5,
	0xf3, 0x7f, 0x56, 0x39, 0x18, 0xc, 0x1e, 0x56, 0xd9, 0x86, 0xda, 0x8e, 0xc3, 0xb3, 0xc, 0x87,
	0x78, 0x35, 0xc9, 0x11, 0x78, 0xb5, 0x27, 0x90, 0x24, 0x57, 0x9e, 0x28, 0x64, 0x4f, 0x33, 0xf,
	0x87, 0xc3, 0xa6, 0xca, 0x5a, 0x87, 0x24, 0x6d, 0x5d, 0xa4, 0xad, 0x97, 0xac, 0xd5, 0xed, 0xcd,
	0xf, 0x38, 0xf2, 0x4e, 0xb3, 0x5c, 0x23, 0x49, 0x7e, 0xf2, 0x82, 0xb4, 0xa0, 0x6, 0x0, 0x67,
	0x7f, 0x90, 0xbf, 0x97, 0x34, 0x22, 0xde, 0x99, 0x64, 0x5c, 0x69, 0xc3, 0x23, 0x59, 0x1a, 0x68,
	0x89, 0x42, 0x41, 0x59, 0x47, 0x54, 0xf1, 0x8c, 0x94, 0x66, 0x60, 0xd0, 0x2b, 0x7d, 0x76, 0xa3,
	0x90, 0x3d, 0x5d, 0x29, 0x48, 0xb, 0x5a, 0x52, 0xe3, 0xef, 0x24, 0x8a, 0x43, 0x76, 0x5a, 0x21,
	0x6b, 0x62, 0x8d, 0xdf, 0x73, 0x19, 0x5, 0xf2, 0xb9, 0xda, 0xec, 0x16, 0x92, 0xd3, 0x78, 0x4a,
	0x89, 0x9f, 0xf7, 0xf9, 0x67, 0x7c, 0x4e, 0xa4, 0xc1, 0xcc, 0x70, 0x6a, 0xf2, 0x39, 0x8d, 0xcd,
	0x4, 0xa, 0x0, 0x4f, 0xfc, 0x70, 0x82, 0x7d, 0x4d, 0xcc, 0x27, 0x22, 0xe6, 0xf6, 0xe6, 0x20,
	0xe2, 0xb4, 0x45, 0x33, 0xa8, 0x57, 0x5e, 0x6e, 0x62, 0xcd, 0x53, 0xd7, 0x4d, 0xd0, 0x7d, 0x84,
	0x31, 0xee, 0xb9, 0xae, 0x9b, 0x52, 0xaf, 0xcf, 0x89, 0xef, 0x8b, 0xdd, 0xf9, 0x45, 0xbe, 0x30,
	0x31, 0x8f, 0x96, 0x13, 0x9e, 0x55, 0x59, 0xb9, 0xce, 0xef, 0x0, 0x2b, 0xec, 0x2f, 0x49, 0x2,
	0x43, 0x1b, 0x12, 0x60, 0xea, 0x23, 0x8b, 0x4, 0xb, 0xbe, 0x36, 0xe5, 0x1f, 0x6b, 0x5b, 0x44,
	0xb, 0x21, 0x69, 0x44, 0xee, 0xa9, 0x7a, 0xa3, 0x2d, 0x9f, 0x10, 0x2b, 0xff, 0x80, 0x32, 0x46,
	0xa2, 0xd7, 0xda, 0x18, 0xf1, 0x53, 0x79, 0x59, 0xa8, 0x92, 0x1e, 0x20, 0xf2, 0x1c, 0x34, 0x8c,
	0x8, 0x5b, 0x6, 0xc8, 0xbc, 0x3e, 0x26, 0x8c, 0xaf, 0x17, 0xc4, 0x23, 0xd3, 0xb2, 0xfa, 0x5c,
	0x1c, 0xd3, 0x49, 0x1a, 0x63, 0x5a, 0xb2, 0xb4, 0x26, 0x9f, 0x93, 0x48, 0xa9, 0xda, 0x2, 0xc0,
	0x97, 0xf9, 0x3a, 0x4f, 0x42, 0xc6, 0x31, 0x65, 0x24, 0xca, 0x95, 0xda, 0xd1, 0x94, 0x3a, 0xc0,
	0xb, 0xb4, 0xb1, 0xe6, 0x64, 0x6d, 0x5b, 0xc8, 0x92, 0x23, 0xb6, 0x8e, 0x83, 0x4f, 0x63, 0x8e,
	0x46, 0x96, 0xb8, 0xd8, 0x56, 0x57, 0x5e, 0x1d, 0xeb, 0x5e, 0x23, 0x90, 0xa4, 0xb0, 0xbe, 0x82,
	0xd, 0x74, 0x30, 0xf1, 0x40, 0x42, 0x47, 0xb8, 0xeb, 0x76, 0x27, 0xf7, 0xdd, 0x91, 0xd7, 0x25,
	0xdd, 0xe9, 0xfd, 0x71, 0xa8, 0x5a, 0xe0, 0xd1, 0x48, 0x15, 0x8d, 0xd8, 0xdb, 0x6e, 0x7a, 0xeb,
	0x6c, 0xef, 0xbb, 0xbb, 0xf1, 0xfe, 0x7e, 0x7c, 0xb0, 0xd5, 0x4d, 0x18, 0xe0, 0x45, 0xea, 0x96,
	0xb7, 0x36, 0xda, 0x4f, 0x67, 0xbe, 0x75, 0xd0, 0x7e, 0xc6, 0xed, 0xae, 0xb4, 0x5f, 0x65, 0x4e,
	0x8f, 0xbd, 0x8d, 0x29, 0x9b, 0xe5, 0x5e, 0xd6, 0x9d, 0xe3, 0x3f, 0x2f, 0xc2, 0x3, 0xbc, 0xc8,
	0x61, 0x9f, 0x16, 0xa, 0xe, 0x0, 0x5f, 0x14, 0xe3, 0x84, 0x8d, 0x72, 0x81, 0xd5, 0x37, 0x6f,
	0x7e, 0x1, 0x80, 0xaf, 0x8b, 0x50, 0xe9, 0x91, 0xb1, 0x32, 0x49, 0x2e, 0xa6, 0x9d, 0x33, 0xc,
	0x0, 0x7c, 0xa3, 0xd5, 0x23, 0x5e, 0x4b, 0x5d, 0xf8, 0xb3, 0x8d, 0x7c, 0x89, 0x94, 0x47, 0xc5,
	0xbd, 0x64, 0xfb, 0xed, 0xd1, 0x49, 0x3c, 0x32, 0xc5, 0x4b, 0x9f, 0x8f, 0xa5, 0xb4, 0xdc, 0x54,
	0x4f, 0x36, 0x2b, 0x1b, 0x59, 0x2b, 0xdb, 0x4c, 0xca, 0x27, 0xe7, 0x79, 0x9a, 0x6c, 0x7e, 0xc5,
	0x7d, 0x1, 0x9e, 0xef, 0xcd, 0xea, 0xda, 0xc8, 0xe4, 0xd1, 0x92, 0xec, 0x96, 0xe4, 0x60, 0xbb,
	0x4d, 0xc4, 0x67, 0x9d, 0x6e, 0x68, 0x9c, 0x4e, 0xaf, 0xd0, 0x95, 0x66, 0xd4, 0x1e, 0x22, 0x7b,
	0xa8, 0x19, 0xef, 0x3b, 0xa8, 0xef, 0x68, 0xc6, 0x87, 0x3, 0x34, 0x1c, 0x68, 0xdf, 0x6c, 0xca,
	0x66, 0xe8, 0x25, 0xf9, 0x7, 0x7, 0xb, 0x9f, 0x7c, 0x27, 0x8d, 0xd4, 0x9b, 0x84, 0xc1, 0x4b,
	0xcd, 0xb, 0xbd, 0xe6, 0x4, 0x99, 0xb6, 0xe6, 0x1, 0x65, 0x38, 0x5a, 0x23, 0xd3, 0xd6, 0x24,
	0xf6, 0xc2, 0xa5, 0xeb, 0x13, 0xd4, 0xef, 0xd9, 0x3, 0xfb, 0xf2, 0x95, 0x26, 0x36, 0xc, 0xfd,
	0x6c, 0x99, 0x4e, 0x3b, 0x4d, 0xdb, 0xe3, 0x9f, 0xc4, 0xf9, 0xf3, 0x9b, 0x84, 0xc4, 0xea, 0x50,
	0x2d, 0xf9, 0x21, 0x5d, 0x38, 0xc0, 0x3e, 0xd4, 0x77, 0xca, 0x3b, 0x7e, 0x29, 0x7c, 0x34, 0x4e,
	0x76, 0xfa, 0x77, 0x12, 0x57, 0x34, 0x68, 0x69, 0x75, 0xcd, 0x3f, 0x29, 0x5b, 0x27, 0xa9, 0x3f,
	0xb0, 0xc2, 0x4c, 0xa6, 0x13, 0x7b, 0x19, 0xb7, 0x7c, 0x2f, 0x53, 0x13, 0xe, 0x7b, 0xbf, 0x9f,
	0x94, 0x86, 0x5b, 0xbe, 0xc9, 0x78, 0x20, 0x7d, 0x89, 0x6f, 0xca, 0xe, 0x34, 0xce, 0x33, 0x3d,
	0x42, 0xc9, 0xc5, 0x6f, 0xdc, 0xad, 0x7f, 0x64, 0xcb, 0xe0, 0xb4, 0x9a, 0x56, 0xbe, 0x6f, 0x42,
	0x15, 0xd4, 0x2f, 0xdb, 0x86, 0x64, 0x4f, 0x6, 0xb9, 0xb5, 0x17, 0xb7, 0xc8, 0x9f, 0x75, 0xed,
	0x90, 0xa0, 0xf0, 0x97, 0x80, 0x3e, 0xd8, 0x11, 0x39, 0x8e, 0x93, 0xed, 0x88, 0x4e, 0xb1, 0x4f,
	0xe5, 0x7a, 0xcf, 0xb8, 0x52, 0x28, 0x5d, 0x56, 0xc2, 0xc7, 0x87, 0x12, 0xb2, 0xe4, 0x87, 0x87,
	0xad, 0x94, 0x20, 0x7e, 0xb4, 0x8f, 0xea, 0xf7, 0xfb, 0x87, 0x7d, 0xd4, 0x99, 0x5a, 0xa2, 0xc7,
	0xea, 0xda, 0x94, 0x57, 0xf8, 0x37, 0x0, 0x0, 0xff, 0xff, 0xc4, 0x12, 0x16, 0x6, 0xa1, 0x10,
	0x0, 0x0,
}

func init() {
	if file_test_thrift != nil {
		return
	}
	type x struct{}
	builder := &thrift_reflection.FileDescriptorBuilder{
		Bytes:         file_idl_test_rawDesc,
		GoTypes:       file_test_thrift_go_types,
		GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
	}
	file_test_thrift = thrift_reflection.BuildFileDescriptor(builder)
}

func GetFileDescriptorForTest() *thrift_reflection.FileDescriptor {
	return file_test_thrift
}
func (p *IDCard) GetDescriptor() *thrift_reflection.StructDescriptor {
	return file_test_thrift.GetStructDescriptor("IDCard")
}
func (p *Person) GetDescriptor() *thrift_reflection.StructDescriptor {
	return file_test_thrift.GetStructDescriptor("Person")
}
func (p *_FieldOptions) GetDescriptor() *thrift_reflection.StructDescriptor {
	return file_test_thrift.GetStructDescriptor("_FieldOptions")
}
func (p *_StructOptions) GetDescriptor() *thrift_reflection.StructDescriptor {
	return file_test_thrift.GetStructDescriptor("_StructOptions")
}
func (p *TinyStruct) GetDescriptor() *thrift_reflection.StructDescriptor {
	return file_test_thrift.GetStructDescriptor("TinyStruct")
}
func (p MyEnum) GetDescriptor() *thrift_reflection.EnumDescriptor {
	return file_test_thrift.GetEnumDescriptor("MyEnum")
}
