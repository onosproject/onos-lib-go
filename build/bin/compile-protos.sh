#!/bin/sh

proto_imports=".:${GOPATH}/src/github.com/gogo/protobuf/protobuf:${GOPATH}/src/github.com/gogo/protobuf:${GOPATH}/src"

protoc -I=$proto_imports --doc_out=docs/logging/api  --doc_opt=markdown,logging.md  --gogo_out=Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,import_path=github.com/onosproject/onos-lib-go/api/logging,plugins=grpc:. api/logging/*.proto

protoc -I=$proto_imports --go_out=../../.. api/asn1/v1/asn1.proto
# Remove the license header copied over by protoc
sed -i "1,15d" api/asn1/v1/asn1/asn1.pb.go

protoc -I=$proto_imports:api --go_out=. pkg/asn1/test/aper-test.proto
# Remove the license header copied over by protoc
sed -i "1,17d" pkg/asn1/test/aper-test.pb.go

protoc-go-inject-tag -input=pkg/asn1/test/aper-test.pb.go

protoc -I=$proto_imports:api --go_out=. pkg/asn1/testsm/test_sm.proto
# Remove the license header copied over by protoc
sed -i "1,19d" pkg/asn1/testsm/test_sm.pb.go

protoc-go-inject-tag -input=pkg/asn1/testsm/test_sm.pb.go
