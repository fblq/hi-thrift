package main

import (
	"github.com/fblq/hi-thrift/gen-go/bill"
	"git.apache.org/thrift.git/lib/go/thrift"
	"log"
	"fmt"
	"os"
	"context"
)

const (
	NetworkAddr = "127.0.0.1:9090"
)

type BillImpl struct {

}
func (b *BillImpl) GetBillList(ctx context.Context, userID string) ([]*bill.BillInfo, error) {
	bills := []*bill.BillInfo{{"1", "201803", 100, userID}, {"2", "201804", 200, userID}}
	log.Printf("GetBillList(%s)", userID)
	return bills, nil
}

func main() {
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	serverTransport, err := thrift.NewTServerSocket(NetworkAddr)
	if err != nil {
		fmt.Println("Error!", err)
		os.Exit(1)
	}
	handler := &BillImpl{}
	processor := bill.NewBillServiceProcessor(handler)
	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	fmt.Println("thrift server in", NetworkAddr)
	server.Serve()
}
