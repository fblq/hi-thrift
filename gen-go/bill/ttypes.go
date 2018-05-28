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

var GoUnusedProtection__ int

// Attributes:
//  - BillID
//  - BillMonth
//  - Amount
//  - UserID
type BillInfo struct {
	BillID    string `thrift:"billID,1" json:"billID"`
	BillMonth string `thrift:"billMonth,2" json:"billMonth"`
	Amount    int64  `thrift:"amount,3" json:"amount"`
	UserID    string `thrift:"userID,4" json:"userID"`
}

func NewBillInfo() *BillInfo {
	return &BillInfo{}
}

func (p *BillInfo) GetBillID() string {
	return p.BillID
}

func (p *BillInfo) GetBillMonth() string {
	return p.BillMonth
}

func (p *BillInfo) GetAmount() int64 {
	return p.Amount
}

func (p *BillInfo) GetUserID() string {
	return p.UserID
}
func (p *BillInfo) Read(iprot thrift.TProtocol) error {
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
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
		case 4:
			if err := p.readField4(iprot); err != nil {
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

func (p *BillInfo) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.BillID = v
	}
	return nil
}

func (p *BillInfo) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.BillMonth = v
	}
	return nil
}

func (p *BillInfo) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.Amount = v
	}
	return nil
}

func (p *BillInfo) readField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 4: ", err)
	} else {
		p.UserID = v
	}
	return nil
}

func (p *BillInfo) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("BillInfo"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := p.writeField4(oprot); err != nil {
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

func (p *BillInfo) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("billID", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:billID: ", p), err)
	}
	if err := oprot.WriteString(string(p.BillID)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.billID (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:billID: ", p), err)
	}
	return err
}

func (p *BillInfo) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("billMonth", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:billMonth: ", p), err)
	}
	if err := oprot.WriteString(string(p.BillMonth)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.billMonth (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:billMonth: ", p), err)
	}
	return err
}

func (p *BillInfo) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("amount", thrift.I64, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:amount: ", p), err)
	}
	if err := oprot.WriteI64(int64(p.Amount)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.amount (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:amount: ", p), err)
	}
	return err
}

func (p *BillInfo) writeField4(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("userID", thrift.STRING, 4); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:userID: ", p), err)
	}
	if err := oprot.WriteString(string(p.UserID)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.userID (4) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 4:userID: ", p), err)
	}
	return err
}

func (p *BillInfo) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("BillInfo(%+v)", *p)
}
