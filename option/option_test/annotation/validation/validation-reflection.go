// Code generated by thriftgo (0.3.2-option-exp). DO NOT EDIT.

package validation

import (
	"reflect"

	"github.com/cloudwego/thriftgo/thrift_reflection"
)

// IDL Name: validation
// IDL Path: option_idl/annotations/validation/validation.thrift

var file_validation_thrift_go_types = []interface{}{
	(*_StructOptions)(nil),         // Struct 0: validation._StructOptions
	(*MyStructWithDefaultVal)(nil), // Struct 1: validation.MyStructWithDefaultVal
	(*_FieldOptions)(nil),          // Struct 2: validation._FieldOptions
	(*TestInfo)(nil),               // Struct 3: validation.TestInfo
	(*_ServiceOptions)(nil),        // Struct 4: validation._ServiceOptions
	(*_MethodOptions)(nil),         // Struct 5: validation._MethodOptions
	(*_EnumOptions)(nil),           // Struct 6: validation._EnumOptions
	(*_EnumValueOptions)(nil),      // Struct 7: validation._EnumValueOptions
	(*MyEnum)(nil),                 // Enum 0: validation.MyEnum
	(*MyBasicTypedef)(nil),         // Enum 0: validation.MyBasicTypedef
	(*MyStructTypedef)(nil),        // Enum 1: validation.MyStructTypedef
}
var file_validation_thrift *thrift_reflection.FileDescriptor
var file_idl_validation_rawDesc = []byte{
	0x1f, 0x8b, 0x8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xff, 0xe4, 0x98, 0xcd, 0x4e, 0xdb, 0x40,
	0x10, 0xc7, 0xc7, 0xeb, 0xf, 0x1c, 0xec, 0x60, 0xda, 0x43, 0xa5, 0xf6, 0x4, 0x2f, 0x0, 0x4d,
	0xa0, 0x94, 0xde, 0x50, 0x55, 0xaa, 0x56, 0x22, 0x45, 0x2a, 0x88, 0x72, 0xb3, 0x9c, 0x78, 0x43,
	0x56, 0xf8, 0x23, 0xb5, 0xd7, 0x96, 0xf2, 0x6c, 0x7d, 0xb4, 0x5e, 0xaa, 0xb5, 0x1d, 0x27, 0x51,
	0x1c, 0x82, 0xea, 0x71, 0x10, 0x2d, 0x97, 0x98, 0x28, 0xfe, 0xdb, 0xbf, 0xff, 0xcc, 0xce, 0xec,
	0x8e, 0x1, 0x12, 0x0, 0x1c, 0x85, 0x63, 0xce, 0xc2, 0xc0, 0x66, 0xae, 0x77, 0xe8, 0x4, 0x41,
	0xc8, 0x1d, 0xf1, 0x6f, 0x7c, 0x98, 0x3a, 0x1e, 0x73, 0xb3, 0xeb, 0xb9, 0xcb, 0x3, 0x3e, 0x8a,
	0xd8, 0x90, 0xb7, 0x81, 0x18, 0x6, 0x0, 0x40, 0x1b, 0xe4, 0xec, 0x42, 0x8, 0x91, 0xbb, 0x10,
	0x0, 0xf6, 0xb, 0x39, 0x4e, 0x63, 0x7e, 0x30, 0xd3, 0x3b, 0x98, 0x69, 0x58, 0xa0, 0x98, 0xe2,
	0x5e, 0xb, 0x54, 0xf1, 0xa9, 0xff, 0xfd, 0x6b, 0x18, 0x40, 0x0, 0x60, 0xc7, 0xbe, 0xe2, 0x51,
	0x32, 0xe0, 0x97, 0x99, 0x40, 0x6c, 0x81, 0x2c, 0x64, 0xb5, 0xba, 0xb2, 0x2f, 0xc7, 0x34, 0x8a,
	0xc3, 0xc0, 0x8e, 0x79, 0xc4, 0x82, 0x3b, 0x9b, 0x5, 0xc3, 0xd0, 0x4, 0xb9, 0xae, 0xaa, 0x96,
	0xcb, 0x81, 0x1, 0x8a, 0x60, 0xff, 0x4e, 0x7f, 0x26, 0x2c, 0xa2, 0xae, 0xe, 0xaa, 0x70, 0xb1,
	0xd, 0x5b, 0x86, 0x25, 0xbc, 0x31, 0x40, 0x87, 0xfc, 0xb3, 0xde, 0xe3, 0xac, 0x2, 0xc2, 0x77,
	0xc6, 0x58, 0x4, 0xb2, 0xef, 0x8c, 0x11, 0x8d, 0x30, 0x41, 0xc1, 0x92, 0xaa, 0x36, 0x95, 0xa0,
	0x9b, 0xba, 0x5b, 0x98, 0x4a, 0x83, 0xc4, 0x47, 0xcb, 0x8b, 0xde, 0xe4, 0x3c, 0x48, 0xfc, 0x6a,
	0x4, 0x19, 0x1d, 0xe1, 0x75, 0x81, 0xd0, 0x77, 0x62, 0x36, 0xb0, 0xf9, 0x64, 0x4c, 0x5d, 0x3a,
	0xc4, 0x62, 0xd9, 0xe9, 0x4d, 0x3e, 0xa, 0xdd, 0xeb, 0x5c, 0xb6, 0x9a, 0x49, 0x41, 0x67, 0x7a,
	0x33, 0x5b, 0xb0, 0xc9, 0x80, 0x63, 0x43, 0x59, 0xbd, 0x49, 0x5e, 0x66, 0x1e, 0xa4, 0x52, 0xd1,
	0xa9, 0xf6, 0x16, 0xa9, 0x5c, 0x3a, 0x74, 0x12, 0x8f, 0xdb, 0xa9, 0xe3, 0x25, 0x14, 0x8b, 0xed,
	0xd5, 0x94, 0xed, 0x7, 0xe3, 0xa3, 0x4f, 0xf9, 0x13, 0x6e, 0x1c, 0xaf, 0x1a, 0x51, 0x5b, 0x42,
	0x6c, 0x83, 0x32, 0xfd, 0x42, 0x45, 0x61, 0x5e, 0xf1, 0x3e, 0x45, 0x65, 0x37, 0xea, 0xca, 0x93,
	0xb4, 0xf3, 0xfc, 0x2a, 0x39, 0x49, 0xbb, 0x4d, 0xbf, 0x34, 0x31, 0x41, 0xd3, 0xf3, 0x56, 0xae,
	0x64, 0xbf, 0xce, 0xfe, 0xb6, 0x41, 0x9e, 0x5e, 0xe6, 0x77, 0x91, 0xb4, 0x4b, 0x40, 0x9d, 0x8b,
	0x3e, 0x3a, 0xe9, 0x11, 0x2, 0x29, 0x61, 0xa7, 0x2b, 0x8a, 0x69, 0x49, 0x29, 0x55, 0x52, 0xea,
	0xf9, 0x5d, 0xd0, 0x30, 0xe3, 0x31, 0x46, 0x2b, 0x66, 0x9d, 0x93, 0x15, 0xd5, 0x75, 0xd, 0xe4,
	0xee, 0x66, 0x20, 0xdf, 0xa1, 0x40, 0x1e, 0x75, 0x57, 0x14, 0xdb, 0x35, 0x90, 0x7b, 0x9b, 0x81,
	0x3c, 0x41, 0x81, 0x3c, 0x39, 0x5e, 0x51, 0x6e, 0xd7, 0x40, 0x9e, 0x6d, 0x6, 0xf2, 0x3d, 0x2,
	0xa4, 0xd2, 0xf, 0xc3, 0x15, 0x4d, 0x65, 0xab, 0xa4, 0x94, 0x1f, 0x2a, 0x3d, 0x82, 0x52, 0x6a,
	0x90, 0xf2, 0x14, 0xa3, 0xc4, 0xba, 0x61, 0xd2, 0xf7, 0x68, 0x35, 0xa7, 0x5e, 0x72, 0x82, 0x2,
	0xe4, 0xac, 0xb5, 0xff, 0xfb, 0xdb, 0xaf, 0xd6, 0x4d, 0x25, 0x67, 0x93, 0xd1, 0xfc, 0xf0, 0x1f,
	0x9e, 0x3, 0x5a, 0xa5, 0xf5, 0xea, 0xba, 0x14, 0x13, 0x96, 0x9b, 0x62, 0xab, 0x21, 0x3d, 0xa6,
	0x1d, 0xde, 0x77, 0xe6, 0xa3, 0xf5, 0xa8, 0x6, 0xba, 0x70, 0x47, 0x63, 0x71, 0x96, 0xd3, 0xce,
	0x5b, 0x8c, 0x65, 0xeb, 0xb1, 0x98, 0x37, 0x1e, 0x9e, 0xed, 0x32, 0x3c, 0xca, 0xda, 0xf0, 0x58,
	0xa2, 0x2a, 0x2, 0x90, 0x66, 0xa2, 0x73, 0xdf, 0xdd, 0x54, 0x74, 0x1a, 0xdf, 0x86, 0x1a, 0xa5,
	0xa9, 0xda, 0x5a, 0x53, 0x73, 0x2e, 0xf5, 0xcb, 0xf9, 0xc5, 0xc5, 0xe5, 0x32, 0x2d, 0xfa, 0x1e,
	0xbf, 0x6d, 0x7f, 0x66, 0xd4, 0x73, 0x17, 0x87, 0x36, 0x52, 0xed, 0x53, 0xda, 0xc0, 0x89, 0x5c,
	0x7b, 0x28, 0x94, 0x9f, 0x68, 0x62, 0x83, 0x6e, 0x94, 0x7e, 0x4d, 0x63, 0xfe, 0x35, 0x18, 0x86,
	0x85, 0x47, 0xa4, 0xf6, 0x7a, 0xe, 0x1c, 0x9f, 0x3e, 0xbf, 0x3, 0x90, 0x16, 0x24, 0x7e, 0x9f,
	0x46, 0x8d, 0x6e, 0x9b, 0x97, 0x67, 0x45, 0xe8, 0xe1, 0xb4, 0xec, 0x2b, 0x1a, 0xa5, 0x6c, 0x40,
	0x91, 0x33, 0x5f, 0x8f, 0xd3, 0x1, 0x56, 0xca, 0x97, 0x29, 0xf7, 0x64, 0x49, 0xbf, 0x63, 0xf7,
	0x28, 0x1f, 0x85, 0xd8, 0xe5, 0xc1, 0xf0, 0x33, 0xd5, 0x7f, 0xc7, 0x27, 0xd3, 0x3e, 0xf, 0x12,
	0x1f, 0xd9, 0xa5, 0x16, 0xe6, 0x60, 0xf3, 0xe9, 0x3d, 0x7a, 0x91, 0x79, 0x74, 0xe3, 0x78, 0x9,
	0xf6, 0x9a, 0xb3, 0x32, 0xa3, 0x50, 0x47, 0x71, 0xf5, 0xed, 0x2a, 0x36, 0x47, 0x60, 0x89, 0x43,
	0x15, 0x2, 0x64, 0x31, 0x99, 0x2e, 0x3c, 0xab, 0xd, 0x28, 0xdd, 0xce, 0x6d, 0x3f, 0xd0, 0x63,
	0x4d, 0x6e, 0x2f, 0x66, 0xf2, 0x12, 0xba, 0xbc, 0x7c, 0x3b, 0xaf, 0x4f, 0x96, 0xf4, 0x2b, 0x82,
	0xa1, 0xd7, 0xec, 0xd9, 0x66, 0xfd, 0x86, 0x3f, 0x6b, 0xd4, 0xf2, 0xf2, 0x74, 0x1e, 0xd1, 0x23,
	0x84, 0x57, 0x9d, 0x4f, 0x7f, 0xb9, 0x62, 0xea, 0x5e, 0x61, 0x70, 0xab, 0xc8, 0xf6, 0x6d, 0x94,
	0x6c, 0xcf, 0xf7, 0xbf, 0xb8, 0xe7, 0xda, 0x47, 0x1c, 0x3a, 0x8c, 0x11, 0xf5, 0xbc, 0x70, 0x8f,
	0x8f, 0x68, 0x44, 0x17, 0xcf, 0xfe, 0xea, 0x94, 0x57, 0xcb, 0x7f, 0xff, 0x27, 0x0, 0x0, 0xff,
	0xff, 0x10, 0x34, 0x4b, 0x4e, 0xb5, 0x1d, 0x0, 0x0,
}

func init() {
	if file_validation_thrift != nil {
		return
	}
	type x struct{}
	builder := &thrift_reflection.FileDescriptorBuilder{
		Bytes:         file_idl_validation_rawDesc,
		GoTypes:       file_validation_thrift_go_types,
		GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
	}
	file_validation_thrift = thrift_reflection.BuildFileDescriptor(builder)
}

func GetFileDescriptorForValidation() *thrift_reflection.FileDescriptor {
	return file_validation_thrift
}
func (p *_StructOptions) GetDescriptor() *thrift_reflection.StructDescriptor {
	return file_validation_thrift.GetStructDescriptor("_StructOptions")
}
func (p *MyStructWithDefaultVal) GetDescriptor() *thrift_reflection.StructDescriptor {
	return file_validation_thrift.GetStructDescriptor("MyStructWithDefaultVal")
}
func (p *_FieldOptions) GetDescriptor() *thrift_reflection.StructDescriptor {
	return file_validation_thrift.GetStructDescriptor("_FieldOptions")
}
func (p *TestInfo) GetDescriptor() *thrift_reflection.StructDescriptor {
	return file_validation_thrift.GetStructDescriptor("TestInfo")
}
func (p *_ServiceOptions) GetDescriptor() *thrift_reflection.StructDescriptor {
	return file_validation_thrift.GetStructDescriptor("_ServiceOptions")
}
func (p *_MethodOptions) GetDescriptor() *thrift_reflection.StructDescriptor {
	return file_validation_thrift.GetStructDescriptor("_MethodOptions")
}
func (p *_EnumOptions) GetDescriptor() *thrift_reflection.StructDescriptor {
	return file_validation_thrift.GetStructDescriptor("_EnumOptions")
}
func (p *_EnumValueOptions) GetDescriptor() *thrift_reflection.StructDescriptor {
	return file_validation_thrift.GetStructDescriptor("_EnumValueOptions")
}
func (p MyEnum) GetDescriptor() *thrift_reflection.EnumDescriptor {
	return file_validation_thrift.GetEnumDescriptor("MyEnum")
}
