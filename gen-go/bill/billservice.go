// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package bill

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

type BillService interface {
	// Parameters:
	//  - UserID
	GetBillList(userID string) (r []*BillInfo, err error)
}

type BillServiceClient struct {
	Transport       thrift.TTransport
	ProtocolFactory thrift.TProtocolFactory
	InputProtocol   thrift.TProtocol
	OutputProtocol  thrift.TProtocol
	SeqId           int32
}

func NewBillServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *BillServiceClient {
	return &BillServiceClient{Transport: t,
		ProtocolFactory: f,
		InputProtocol:   f.GetProtocol(t),
		OutputProtocol:  f.GetProtocol(t),
		SeqId:           0,
	}
}

func NewBillServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *BillServiceClient {
	return &BillServiceClient{Transport: t,
		ProtocolFactory: nil,
		InputProtocol:   iprot,
		OutputProtocol:  oprot,
		SeqId:           0,
	}
}

// Parameters:
//  - UserID
func (p *BillServiceClient) GetBillList(userID string) (r []*BillInfo, err error) {
	if err = p.sendGetBillList(userID); err != nil {
		return
	}
	return p.recvGetBillList()
}

func (p *BillServiceClient) sendGetBillList(userID string) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("GetBillList", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := BillServiceGetBillListArgs{
		UserID: userID,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *BillServiceClient) recvGetBillList() (value []*BillInfo, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "GetBillList" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "GetBillList failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "GetBillList failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error0 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error1 error
		error1, err = error0.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error1
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "GetBillList failed: invalid message type")
		return
	}
	result := BillServiceGetBillListResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

type BillServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      BillService
}

func (p *BillServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *BillServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *BillServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewBillServiceProcessor(handler BillService) *BillServiceProcessor {

	self2 := &BillServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self2.processorMap["GetBillList"] = &billServiceProcessorGetBillList{handler: handler}
	return self2
}

func (p *BillServiceProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x3 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x3.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return false, x3

}

type billServiceProcessorGetBillList struct {
	handler BillService
}

func (p *billServiceProcessorGetBillList) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := BillServiceGetBillListArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("GetBillList", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := BillServiceGetBillListResult{}
	var retval []*BillInfo
	var err2 error
	if retval, err2 = p.handler.GetBillList(args.UserID); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing GetBillList: "+err2.Error())
		oprot.WriteMessageBegin("GetBillList", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("GetBillList", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - UserID
type BillServiceGetBillListArgs struct {
	UserID string `thrift:"userID,1" json:"userID"`
}

func NewBillServiceGetBillListArgs() *BillServiceGetBillListArgs {
	return &BillServiceGetBillListArgs{}
}

func (p *BillServiceGetBillListArgs) GetUserID() string {
	return p.UserID
}
func (p *BillServiceGetBillListArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *BillServiceGetBillListArgs) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.UserID = v
	}
	return nil
}

func (p *BillServiceGetBillListArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("GetBillList_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *BillServiceGetBillListArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("userID", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:userID: ", p), err)
	}
	if err := oprot.WriteString(string(p.UserID)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.userID (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:userID: ", p), err)
	}
	return err
}

func (p *BillServiceGetBillListArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("BillServiceGetBillListArgs(%+v)", *p)
}

// Attributes:
//  - Success
type BillServiceGetBillListResult struct {
	Success []*BillInfo `thrift:"success,0" json:"success,omitempty"`
}

func NewBillServiceGetBillListResult() *BillServiceGetBillListResult {
	return &BillServiceGetBillListResult{}
}

var BillServiceGetBillListResult_Success_DEFAULT []*BillInfo

func (p *BillServiceGetBillListResult) GetSuccess() []*BillInfo {
	return p.Success
}
func (p *BillServiceGetBillListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *BillServiceGetBillListResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *BillServiceGetBillListResult) readField0(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]*BillInfo, 0, size)
	p.Success = tSlice
	for i := 0; i < size; i++ {
		_elem4 := &BillInfo{}
		if err := _elem4.Read(iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem4), err)
		}
		p.Success = append(p.Success, _elem4)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *BillServiceGetBillListResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("GetBillList_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *BillServiceGetBillListResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.LIST, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Success)); err != nil {
			return thrift.PrependError("error writing list begin: ", err)
		}
		for _, v := range p.Success {
			if err := v.Write(oprot); err != nil {
				return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
			}
		}
		if err := oprot.WriteListEnd(); err != nil {
			return thrift.PrependError("error writing list end: ", err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *BillServiceGetBillListResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("BillServiceGetBillListResult(%+v)", *p)
}
