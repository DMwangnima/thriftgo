// Code generated by thriftgo (0.1.1). DO NOT EDIT.

package plugin

import (
	"fmt"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/cloudwego/thriftgo/parser"
)

type Request struct {
	Version             string         `thrift:"Version,1,required" json:"Version"`
	GeneratorParameters []string       `thrift:"GeneratorParameters,2,required" json:"GeneratorParameters"`
	PluginParameters    []string       `thrift:"PluginParameters,3,required" json:"PluginParameters"`
	Language            string         `thrift:"Language,4,required" json:"Language"`
	OutputPath          string         `thrift:"OutputPath,5,required" json:"OutputPath"`
	Recursive           bool           `thrift:"Recursive,6,required" json:"Recursive"`
	AST                 *parser.Thrift `thrift:"AST,7,required" json:"AST"`
}

func NewRequest() *Request {
	return &Request{}
}

func (p *Request) GetVersion() string {
	return p.Version
}

func (p *Request) GetGeneratorParameters() []string {
	return p.GeneratorParameters
}

func (p *Request) GetPluginParameters() []string {
	return p.PluginParameters
}

func (p *Request) GetLanguage() string {
	return p.Language
}

func (p *Request) GetOutputPath() string {
	return p.OutputPath
}

func (p *Request) GetRecursive() bool {
	return p.Recursive
}

var Request_AST_DEFAULT *parser.Thrift

func (p *Request) GetAST() *parser.Thrift {
	if !p.IsSetAST() {
		return Request_AST_DEFAULT
	}
	return p.AST
}

var fieldIDToName_Request = map[int16]string{
	1: "Version",
	2: "GeneratorParameters",
	3: "PluginParameters",
	4: "Language",
	5: "OutputPath",
	6: "Recursive",
	7: "AST",
}

func (p *Request) IsSetAST() bool {
	return p.AST != nil
}

func (p *Request) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16
	var issetVersion bool = false
	var issetGeneratorParameters bool = false
	var issetPluginParameters bool = false
	var issetLanguage bool = false
	var issetOutputPath bool = false
	var issetRecursive bool = false
	var issetAST bool = false

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
				issetVersion = true
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.LIST {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
				issetGeneratorParameters = true
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 3:
			if fieldTypeId == thrift.LIST {
				if err = p.ReadField3(iprot); err != nil {
					goto ReadFieldError
				}
				issetPluginParameters = true
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 4:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField4(iprot); err != nil {
					goto ReadFieldError
				}
				issetLanguage = true
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 5:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField5(iprot); err != nil {
					goto ReadFieldError
				}
				issetOutputPath = true
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 6:
			if fieldTypeId == thrift.BOOL {
				if err = p.ReadField6(iprot); err != nil {
					goto ReadFieldError
				}
				issetRecursive = true
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 7:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField7(iprot); err != nil {
					goto ReadFieldError
				}
				issetAST = true
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	if !issetVersion {
		fieldId = 1
		goto RequiredFieldNotSetError
	}

	if !issetGeneratorParameters {
		fieldId = 2
		goto RequiredFieldNotSetError
	}

	if !issetPluginParameters {
		fieldId = 3
		goto RequiredFieldNotSetError
	}

	if !issetLanguage {
		fieldId = 4
		goto RequiredFieldNotSetError
	}

	if !issetOutputPath {
		fieldId = 5
		goto RequiredFieldNotSetError
	}

	if !issetRecursive {
		fieldId = 6
		goto RequiredFieldNotSetError
	}

	if !issetAST {
		fieldId = 7
		goto RequiredFieldNotSetError
	}
	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_Request[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
RequiredFieldNotSetError:
	return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field %s is not set", fieldIDToName_Request[fieldId]))
}

func (p *Request) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.Version = v
	}
	return nil
}

func (p *Request) ReadField2(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return err
	}
	p.GeneratorParameters = make([]string, 0, size)
	for i := 0; i < size; i++ {
		var _elem string
		if v, err := iprot.ReadString(); err != nil {
			return err
		} else {
			_elem = v
		}

		p.GeneratorParameters = append(p.GeneratorParameters, _elem)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return err
	}
	return nil
}

func (p *Request) ReadField3(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return err
	}
	p.PluginParameters = make([]string, 0, size)
	for i := 0; i < size; i++ {
		var _elem string
		if v, err := iprot.ReadString(); err != nil {
			return err
		} else {
			_elem = v
		}

		p.PluginParameters = append(p.PluginParameters, _elem)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return err
	}
	return nil
}

func (p *Request) ReadField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.Language = v
	}
	return nil
}

func (p *Request) ReadField5(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.OutputPath = v
	}
	return nil
}

func (p *Request) ReadField6(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBool(); err != nil {
		return err
	} else {
		p.Recursive = v
	}
	return nil
}

func (p *Request) ReadField7(iprot thrift.TProtocol) error {
	p.AST = parser.NewThrift()
	if err := p.AST.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *Request) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("Request"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}
		if err = p.writeField3(oprot); err != nil {
			fieldId = 3
			goto WriteFieldError
		}
		if err = p.writeField4(oprot); err != nil {
			fieldId = 4
			goto WriteFieldError
		}
		if err = p.writeField5(oprot); err != nil {
			fieldId = 5
			goto WriteFieldError
		}
		if err = p.writeField6(oprot); err != nil {
			fieldId = 6
			goto WriteFieldError
		}
		if err = p.writeField7(oprot); err != nil {
			fieldId = 7
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *Request) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("Version", thrift.STRING, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Version); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *Request) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("GeneratorParameters", thrift.LIST, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteListBegin(thrift.STRING, len(p.GeneratorParameters)); err != nil {
		return err
	}
	for _, v := range p.GeneratorParameters {
		if err := oprot.WriteString(v); err != nil {
			return err
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *Request) writeField3(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("PluginParameters", thrift.LIST, 3); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteListBegin(thrift.STRING, len(p.PluginParameters)); err != nil {
		return err
	}
	for _, v := range p.PluginParameters {
		if err := oprot.WriteString(v); err != nil {
			return err
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 end error: ", p), err)
}

func (p *Request) writeField4(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("Language", thrift.STRING, 4); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Language); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 4 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 4 end error: ", p), err)
}

func (p *Request) writeField5(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("OutputPath", thrift.STRING, 5); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.OutputPath); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 5 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 5 end error: ", p), err)
}

func (p *Request) writeField6(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("Recursive", thrift.BOOL, 6); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteBool(p.Recursive); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 6 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 6 end error: ", p), err)
}

func (p *Request) writeField7(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("AST", thrift.STRUCT, 7); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.AST.Write(oprot); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 7 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 7 end error: ", p), err)
}

func (p *Request) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Request(%+v)", *p)
}

type Generated struct {
	Content        string  `thrift:"Content,1,required" json:"Content"`
	Name           *string `thrift:"Name,2" json:"Name,omitempty"`
	InsertionPoint *string `thrift:"InsertionPoint,3" json:"InsertionPoint,omitempty"`
}

func NewGenerated() *Generated {
	return &Generated{}
}

func (p *Generated) GetContent() string {
	return p.Content
}

var Generated_Name_DEFAULT string

func (p *Generated) GetName() string {
	if !p.IsSetName() {
		return Generated_Name_DEFAULT
	}
	return *p.Name
}

var Generated_InsertionPoint_DEFAULT string

func (p *Generated) GetInsertionPoint() string {
	if !p.IsSetInsertionPoint() {
		return Generated_InsertionPoint_DEFAULT
	}
	return *p.InsertionPoint
}

var fieldIDToName_Generated = map[int16]string{
	1: "Content",
	2: "Name",
	3: "InsertionPoint",
}

func (p *Generated) IsSetName() bool {
	return p.Name != nil
}

func (p *Generated) IsSetInsertionPoint() bool {
	return p.InsertionPoint != nil
}

func (p *Generated) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16
	var issetContent bool = false

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
				issetContent = true
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 3:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField3(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	if !issetContent {
		fieldId = 1
		goto RequiredFieldNotSetError
	}
	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_Generated[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
RequiredFieldNotSetError:
	return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field %s is not set", fieldIDToName_Generated[fieldId]))
}

func (p *Generated) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.Content = v
	}
	return nil
}

func (p *Generated) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.Name = &v
	}
	return nil
}

func (p *Generated) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.InsertionPoint = &v
	}
	return nil
}

func (p *Generated) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("Generated"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}
		if err = p.writeField3(oprot); err != nil {
			fieldId = 3
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *Generated) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("Content", thrift.STRING, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Content); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *Generated) writeField2(oprot thrift.TProtocol) (err error) {
	if p.IsSetName() {
		if err = oprot.WriteFieldBegin("Name", thrift.STRING, 2); err != nil {
			goto WriteFieldBeginError
		}
		if err := oprot.WriteString(*p.Name); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *Generated) writeField3(oprot thrift.TProtocol) (err error) {
	if p.IsSetInsertionPoint() {
		if err = oprot.WriteFieldBegin("InsertionPoint", thrift.STRING, 3); err != nil {
			goto WriteFieldBeginError
		}
		if err := oprot.WriteString(*p.InsertionPoint); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 end error: ", p), err)
}

func (p *Generated) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Generated(%+v)", *p)
}

type Response struct {
	Error    *string      `thrift:"Error,1" json:"Error,omitempty"`
	Contents []*Generated `thrift:"Contents,2" json:"Contents,omitempty"`
	Warnings []string     `thrift:"Warnings,3" json:"Warnings,omitempty"`
}

func NewResponse() *Response {
	return &Response{}
}

var Response_Error_DEFAULT string

func (p *Response) GetError() string {
	if !p.IsSetError() {
		return Response_Error_DEFAULT
	}
	return *p.Error
}

var Response_Contents_DEFAULT []*Generated

func (p *Response) GetContents() []*Generated {
	if !p.IsSetContents() {
		return Response_Contents_DEFAULT
	}
	return p.Contents
}

var Response_Warnings_DEFAULT []string

func (p *Response) GetWarnings() []string {
	if !p.IsSetWarnings() {
		return Response_Warnings_DEFAULT
	}
	return p.Warnings
}

var fieldIDToName_Response = map[int16]string{
	1: "Error",
	2: "Contents",
	3: "Warnings",
}

func (p *Response) IsSetError() bool {
	return p.Error != nil
}

func (p *Response) IsSetContents() bool {
	return p.Contents != nil
}

func (p *Response) IsSetWarnings() bool {
	return p.Warnings != nil
}

func (p *Response) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.LIST {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 3:
			if fieldTypeId == thrift.LIST {
				if err = p.ReadField3(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_Response[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *Response) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.Error = &v
	}
	return nil
}

func (p *Response) ReadField2(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return err
	}
	p.Contents = make([]*Generated, 0, size)
	for i := 0; i < size; i++ {
		_elem := NewGenerated()
		if err := _elem.Read(iprot); err != nil {
			return err
		}

		p.Contents = append(p.Contents, _elem)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return err
	}
	return nil
}

func (p *Response) ReadField3(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return err
	}
	p.Warnings = make([]string, 0, size)
	for i := 0; i < size; i++ {
		var _elem string
		if v, err := iprot.ReadString(); err != nil {
			return err
		} else {
			_elem = v
		}

		p.Warnings = append(p.Warnings, _elem)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return err
	}
	return nil
}

func (p *Response) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("Response"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}
		if err = p.writeField3(oprot); err != nil {
			fieldId = 3
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *Response) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetError() {
		if err = oprot.WriteFieldBegin("Error", thrift.STRING, 1); err != nil {
			goto WriteFieldBeginError
		}
		if err := oprot.WriteString(*p.Error); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *Response) writeField2(oprot thrift.TProtocol) (err error) {
	if p.IsSetContents() {
		if err = oprot.WriteFieldBegin("Contents", thrift.LIST, 2); err != nil {
			goto WriteFieldBeginError
		}
		if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Contents)); err != nil {
			return err
		}
		for _, v := range p.Contents {
			if err := v.Write(oprot); err != nil {
				return err
			}
		}
		if err := oprot.WriteListEnd(); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *Response) writeField3(oprot thrift.TProtocol) (err error) {
	if p.IsSetWarnings() {
		if err = oprot.WriteFieldBegin("Warnings", thrift.LIST, 3); err != nil {
			goto WriteFieldBeginError
		}
		if err := oprot.WriteListBegin(thrift.STRING, len(p.Warnings)); err != nil {
			return err
		}
		for _, v := range p.Warnings {
			if err := oprot.WriteString(v); err != nil {
				return err
			}
		}
		if err := oprot.WriteListEnd(); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 end error: ", p), err)
}

func (p *Response) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Response(%+v)", *p)
}
