#!/usr/bin/env bash
mkdir output
go build -o ./output/thrift-client github.com/fblq/hi-thrift/client
go build -o ./output/thrift-server github.com/fblq/hi-thrift/server
