package main

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/fblq/hi-thrift/gen-go/bill"
	"context"
	"log"
)

func main() {
	sock, err := thrift.NewTSocket("localhost:9090")
	if err != nil {
		panic(err)
	}
	defer sock.Close()
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	transport, err := transportFactory.GetTransport(sock)
	if err != nil {
		panic(err)
	}
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	err = transport.Open()
	if err != nil {
		panic(err)
	}
	defer transport.Close()
	protocol := protocolFactory.GetProtocol(transport)
	client := thrift.NewTStandardClient(protocol, protocol)
	billClient := bill.NewBillServiceClient(client)
	ctx := context.Background()
	billInfoList, err := billClient.GetBillList(ctx, "123")
	if err != nil {
		log.Print(err.Error())
	}
	log.Printf("GetBillList = %+v", billInfoList)

	billInfoList, err = billClient.GetBillList(ctx, "124")
	if err != nil {
		log.Print(err.Error())
	}
	log.Printf("GetBillList = %+v", billInfoList)

}
