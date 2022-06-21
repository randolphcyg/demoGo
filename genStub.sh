#!/bin/bash

goctl rpc protoc demo.proto --go_out=./tmp --go-grpc_out=./tmp --zrpc_out=./tmp  --style goZero

rm -r ./javastub/democlient/
rm -r ./javastub/demopb/

mv ./tmp/democlient/ ./javastub/democlient/
mv ./tmp/demopb/ ./javastub/demopb/

rm -r ./tmp