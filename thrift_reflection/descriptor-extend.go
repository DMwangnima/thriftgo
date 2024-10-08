// Copyright 2023 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package thrift_reflection

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/cloudwego/thriftgo/utils"
)

func (f *FileDescriptor) GetIncludeFD(alias string) *FileDescriptor {
	if f == nil {
		return nil
	}
	if alias == "" {
		return f
	}
	includePath := f.Includes[alias]
	if includePath != "" {
		return GetGlobalDescriptor(f).LookupFD(includePath)
	}
	return nil
}

func (f *FileDescriptor) getDescriptor(name string, lookupFunc func(fd *FileDescriptor, name string) interface{}) interface{} {
	if f == nil {
		return nil
	}
	if name == "" {
		return nil
	}
	prefix, name := utils.ParseAlias(name)
	if prefix == "" {
		return lookupFunc(f, name)
	} else {
		fromFd := f.GetIncludeFD(prefix)
		if fromFd != nil {
			return fromFd.getDescriptor(name, lookupFunc)
		}
	}
	return nil
}

func (f *FileDescriptor) GetStructDescriptor(name string) *StructDescriptor {
	des := f.getDescriptor(name, func(fd *FileDescriptor, name string) interface{} {
		for _, s := range fd.Structs {
			if s.Name == name {
				return s
			}
		}
		return nil
	})
	if des != nil {
		return des.(*StructDescriptor)
	}
	return nil
}

func (f *FileDescriptor) GetTypedefDescriptor(name string) *TypedefDescriptor {
	des := f.getDescriptor(name, func(fd *FileDescriptor, name string) interface{} {
		for _, s := range fd.Typedefs {
			if s.Alias == name {
				return s
			}
		}
		return nil
	})
	if des != nil {
		return des.(*TypedefDescriptor)
	}
	return nil
}

func (f *FileDescriptor) GetConstDescriptor(name string) *ConstDescriptor {
	des := f.getDescriptor(name, func(fd *FileDescriptor, name string) interface{} {
		for _, s := range fd.Consts {
			if s.Name == name {
				return s
			}
		}
		return nil
	})
	if des != nil {
		return des.(*ConstDescriptor)
	}
	return nil
}

func (f *FileDescriptor) GetEnumDescriptor(name string) *EnumDescriptor {
	des := f.getDescriptor(name, func(fd *FileDescriptor, name string) interface{} {
		for _, s := range fd.Enums {
			if s.Name == name {
				return s
			}
		}
		return nil
	})
	if des != nil {
		return des.(*EnumDescriptor)
	}
	return nil
}

func (f *FileDescriptor) GetExceptionDescriptor(name string) *StructDescriptor {
	des := f.getDescriptor(name, func(fd *FileDescriptor, name string) interface{} {
		for _, s := range fd.Exceptions {
			if s.Name == name {
				return s
			}
		}
		return nil
	})
	if des != nil {
		return des.(*StructDescriptor)
	}
	return nil
}

func (f *FileDescriptor) GetUnionDescriptor(name string) *StructDescriptor {
	des := f.getDescriptor(name, func(fd *FileDescriptor, name string) interface{} {
		for _, s := range fd.Unions {
			if s.Name == name {
				return s
			}
		}
		return nil
	})
	if des != nil {
		return des.(*StructDescriptor)
	}
	return nil
}

func (f *FileDescriptor) GetServiceDescriptor(name string) *ServiceDescriptor {
	des := f.getDescriptor(name, func(fd *FileDescriptor, name string) interface{} {
		for _, s := range fd.Services {
			if s.Name == name {
				return s
			}
		}
		return nil
	})
	if des != nil {
		return des.(*ServiceDescriptor)
	}
	return nil
}

func (f *FileDescriptor) GetMethodDescriptor(service, method string) *MethodDescriptor {
	if f == nil {
		return nil
	}
	if service == "" {
		for _, svc := range f.Services {
			for _, m := range svc.Methods {
				if m.Name == method {
					return m
				}
			}
		}
		return nil
	}
	svc := f.GetServiceDescriptor(service)
	if svc != nil {
		for _, m := range svc.Methods {
			if m.Name == method {
				return m
			}
		}
	}
	return nil
}

func (p *FieldDescriptor) IsOptional() bool {
	return p.GetRequiredness() == "Optional"
}

func (p *FieldDescriptor) IsDefault() bool {
	return p.GetRequiredness() == "Default"
}

func (p *FieldDescriptor) IsRequired() bool {
	return p.GetRequiredness() == "Required"
}

func (p *FieldDescriptor) GetGoType() (reflect.Type, error) {
	return p.GetType().GetGoType()
}

type thriftReflectStruct interface {
	GetDescriptor() *StructDescriptor
}

// GetInstanceValue
// Notice: This API is EXPERIMENTAL and may be changed or removed in a later release.
func (p *FieldDescriptor) GetInstanceValue(instance thriftReflectStruct) (val interface{}, err error) {
	structDesc := instance.GetDescriptor()
	for idx, fieldDesc := range structDesc.GetFields() {
		if fieldDesc == p {
			defer func() {
				if r := recover(); r != nil {
					err = fmt.Errorf("Recovered from panic: %v", r)
				}
			}()
			rv := reflect.ValueOf(instance).Elem()
			rf := rv.Field(idx)
			return rf.Interface(), nil
		}
	}
	return nil, errors.New("field not exists for current struct")
}

// SetInstanceValue
// Notice: This API is EXPERIMENTAL and may be changed or removed in a later release.
func (p *FieldDescriptor) SetInstanceValue(instance thriftReflectStruct, value interface{}) (err error) {
	structDesc := instance.GetDescriptor()
	for idx, fieldDesc := range structDesc.GetFields() {
		if fieldDesc == p {
			defer func() {
				if r := recover(); r != nil {
					err = fmt.Errorf("Recovered from panic: %v", r)
				}
			}()
			rv := reflect.ValueOf(instance).Elem()
			rf := rv.Field(idx)
			rf.Set(reflect.ValueOf(value))
			return nil
		}
	}
	return errors.New("field not exists for current struct")
}

func (sd *ServiceDescriptor) GetMethodByName(name string) *MethodDescriptor {
	for _, m := range sd.GetMethods() {
		if m.GetName() == name {
			return m
		}
	}
	return nil
}

func (s *ServiceDescriptor) GetParent() *ServiceDescriptor {
	return GetGlobalDescriptor(s).LookupFD(s.Filepath).GetServiceDescriptor(s.Base)
}

func (s *ServiceDescriptor) GetAllMethods() []*MethodDescriptor {
	allMethods := []*MethodDescriptor{}

	svc := s
	for svc != nil {
		allMethods = append(allMethods, svc.GetMethods()...)
		svc = svc.GetParent()
	}
	return allMethods
}

func (s *ServiceDescriptor) GetMethodByNameFromAll(name string) *MethodDescriptor {
	for _, m := range s.GetAllMethods() {
		if m.GetName() == name {
			return m
		}
	}
	return nil
}

func (ed *EnumDescriptor) GetGoType() reflect.Type {
	gd := GetGlobalDescriptor(ed)
	if gd != nil && gd.enumDes2goType != nil {
		return gd.enumDes2goType[ed]
	}
	return nil
}

func (td *TypeDescriptor) IsBasic() bool {
	return utils.IsBasic(td.GetName())
}

func (td *TypeDescriptor) IsContainer() bool {
	return utils.IsContainer(td.GetName())
}

func (td *TypeDescriptor) IsMap() bool {
	return td.IsContainer() && td.GetName() == "map"
}

func (td *TypeDescriptor) IsList() bool {
	return td.IsContainer() && (td.GetName() == "set" || td.GetName() == "list")
}

func (td *TypeDescriptor) GetExceptionDescriptor() (*StructDescriptor, error) {
	if td.IsContainer() || td.IsBasic() {
		return nil, errors.New("not exception")
	}
	stDesc := GetGlobalDescriptor(td).LookupFD(td.Filepath).GetExceptionDescriptor(td.GetName())
	if stDesc != nil {
		return stDesc, nil
	}
	return nil, errors.New("exception not found")
}

func (td *TypeDescriptor) GetUnionDescriptor() (*StructDescriptor, error) {
	if td.IsContainer() || td.IsBasic() {
		return nil, errors.New("not union")
	}
	stDesc := GetGlobalDescriptor(td).LookupFD(td.Filepath).GetUnionDescriptor(td.GetName())
	if stDesc != nil {
		return stDesc, nil
	}
	return nil, errors.New("union not found")
}

func (td *TypeDescriptor) GetStructDescriptor() (*StructDescriptor, error) {
	if td.IsContainer() || td.IsBasic() {
		return nil, errors.New("not struct")
	}
	stDesc := GetGlobalDescriptor(td).LookupFD(td.Filepath).GetStructDescriptor(td.GetName())
	if stDesc != nil {
		return stDesc, nil
	}
	return nil, errors.New("struct not found")
}

func (td *TypeDescriptor) IsException() bool {
	sd, err := td.GetExceptionDescriptor()
	isException := err == nil && sd != nil
	return isException
}

func (td *TypeDescriptor) IsUnion() bool {
	sd, err := td.GetUnionDescriptor()
	isUnion := err == nil && sd != nil
	return isUnion
}

func (td *TypeDescriptor) IsStruct() bool {
	sd, err := td.GetStructDescriptor()
	isStruct := err == nil && sd != nil
	return isStruct
}

func (td *TypeDescriptor) IsEnum() bool {
	sd, err := td.GetEnumDescriptor()
	isEnum := err == nil && sd != nil
	return isEnum
}

func (td *TypeDescriptor) GetEnumDescriptor() (*EnumDescriptor, error) {
	if td.IsContainer() || td.IsBasic() {
		return nil, errors.New("not enum")
	}
	prefix, name := utils.ParseAlias(td.GetName())
	fd := GetGlobalDescriptor(td).LookupFD(td.Filepath)
	if fd != nil {
		targetFd := fd.GetIncludeFD(prefix)
		if targetFd != nil {
			return targetFd.GetEnumDescriptor(name), nil
		}
	}
	return nil, errors.New("enum not found")
}

func (td *TypeDescriptor) IsTypedef() bool {
	sd, err := td.GetTypedefDescriptor()
	isTypedef := err == nil && sd != nil
	return isTypedef
}

func (td *TypeDescriptor) GetTypedefDescriptor() (*TypedefDescriptor, error) {
	if td.IsContainer() || td.IsBasic() {
		return nil, errors.New("not typedef")
	}
	prefix, name := utils.ParseAlias(td.GetName())
	fd := GetGlobalDescriptor(td).LookupFD(td.Filepath)
	if fd != nil {
		targetFd := fd.GetIncludeFD(prefix)
		if targetFd != nil {
			return targetFd.GetTypedefDescriptor(name), nil
		}
	}
	return nil, errors.New("typedef not found")
}

func (td *TypeDescriptor) GetGoType() (reflect.Type, error) {
	if td.IsBasic() {
		switch td.GetName() {
		case "bool":
			return reflect.TypeOf(bool(true)), nil
		case "byte":
			return reflect.TypeOf(byte(0)), nil
		case "i8":
			return reflect.TypeOf(int8(0)), nil
		case "i16":
			return reflect.TypeOf(int16(0)), nil
		case "i32":
			return reflect.TypeOf(int32(0)), nil
		case "i64":
			return reflect.TypeOf(int64(0)), nil
		case "double":
			return reflect.TypeOf(float64(0)), nil
		case "binary":
			return reflect.TypeOf([]byte{0}), nil
		case "string":
			return reflect.TypeOf(string("")), nil
		default:
			return nil, errors.New("unknown basic type")
		}
	}
	if td.IsStruct() {
		sd, err := td.GetStructDescriptor()
		if err != nil {
			return nil, err
		}
		return sd.GetGoType(), nil
	}
	if td.IsEnum() {
		ed, err := td.GetEnumDescriptor()
		if err != nil {
			return nil, err
		}
		return ed.GetGoType(), nil
	}
	if td.IsMap() {
		krt, err := td.GetKeyType().GetGoType()
		if td.GetKeyType().IsStruct() {
			krt = reflect.PtrTo(krt)
		}
		if err != nil {
			return nil, err
		}
		vrt, err := td.GetValueType().GetGoType()
		if td.GetValueType().IsStruct() {
			vrt = reflect.PtrTo(vrt)
		}
		if err != nil {
			return nil, err
		}
		mapType := reflect.MapOf(krt, vrt)
		return mapType, nil
	}
	if td.IsList() {
		if td.GetName() == "list" || td.GetName() == "set" {
			vrt, err := td.GetValueType().GetGoType()
			if td.GetValueType().IsStruct() {
				vrt = reflect.PtrTo(vrt)
			}
			if err != nil {
				return nil, err
			}
			listType := reflect.SliceOf(vrt)
			return listType, nil
		}
	}
	return nil, errors.New("not basic type:" + td.GetName())
}

func (s *StructDescriptor) GetGoType() reflect.Type {
	gd := GetGlobalDescriptor(s)
	if gd != nil && gd.structDes2goType != nil {
		return gd.structDes2goType[s]
	}
	return nil
}

func (s *TypedefDescriptor) GetGoType() reflect.Type {
	gd := GetGlobalDescriptor(s)
	if gd != nil && gd.typedefDes2goType != nil {
		return gd.typedefDes2goType[s]
	}
	return nil
}

func (s *StructDescriptor) GetFieldByName(name string) *FieldDescriptor {
	if s == nil {
		return nil
	}
	for _, f := range s.Fields {
		if f.Name == name {
			return f
		}
	}
	return nil
}

func (s *StructDescriptor) GetFieldById(id int32) *FieldDescriptor {
	if s == nil {
		return nil
	}
	for _, f := range s.Fields {
		if f.ID == id {
			return f
		}
	}
	return nil
}

func (s *ConstValueDescriptor) GetValueAsString() string {
	t := s.GetType()
	if t == ConstValueType_INT {
		return fmt.Sprintf("%v", s.GetValueInt())
	}
	if t == ConstValueType_DOUBLE {
		return fmt.Sprintf("%v", s.GetValueDouble())
	}
	if t == ConstValueType_BOOL {
		return fmt.Sprintf("%v", s.GetValueBool())
	}
	if t == ConstValueType_STRING {
		return s.GetValueString()
	}
	if t == ConstValueType_MAP {
		valueMapArr := []string{}
		for k, v := range s.GetValueMap() {
			valueMapArr = append(valueMapArr, k.GetValueAsString()+":"+v.GetValueAsString())
		}
		return fmt.Sprintf("{%v}", strings.Join(valueMapArr, ","))
	}
	if t == ConstValueType_LIST {
		valueList := []string{}
		for _, v := range s.GetValueList() {
			valueList = append(valueList, v.GetValueAsString())
		}
		return fmt.Sprintf("[%v]", strings.Join(valueList, ","))
	}
	if t == ConstValueType_IDENTIFIER {
		targetConst := GetGlobalDescriptor(s).LookupConst(s.GetValueIdentifier(), "")
		if targetConst != nil {
			return targetConst.GetValue().GetValueAsString()
		}
	}
	return ""
}

func (d *ServiceDescriptor) setExtra(m map[string]string) {
	d.Extra = m
}

func (d *MethodDescriptor) setExtra(m map[string]string) {
	d.Extra = m
}

func (d *FieldDescriptor) setExtra(m map[string]string) {
	d.Extra = m
}

func (d *StructDescriptor) setExtra(m map[string]string) {
	d.Extra = m
}

func (d *EnumDescriptor) setExtra(m map[string]string) {
	d.Extra = m
}

func (d *FileDescriptor) setExtra(m map[string]string) {
	d.Extra = m
}

func (d *TypedefDescriptor) setExtra(m map[string]string) {
	d.Extra = m
}

func (d *EnumValueDescriptor) setExtra(m map[string]string) {
	d.Extra = m
}

func (d *ConstDescriptor) setExtra(m map[string]string) {
	d.Extra = m
}

func (d *ConstValueDescriptor) setExtra(m map[string]string) {
	d.Extra = m
}

func (d *TypeDescriptor) setExtra(m map[string]string) {
	d.Extra = m
}
