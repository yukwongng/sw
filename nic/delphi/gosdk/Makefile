.PHONY: default all pb test

default: all

all: pb build test

client_test.delphi.go: client_test.proto
	make -C ${GOPATH}/src/github.com/pensando/sw/nic/delphi/messenger/proto && \
	make -C ${GOPATH}/src/github.com/pensando/sw/nic/delphi/proto/delphi && \
	protoc --delphigo_out=Mdelphi.proto=github.com/pensando/sw/nic/delphi/proto/delphi:. \
		./client_test.proto -I /usr/src/github.com/pensando/sw -I . \
		-I ${GOPATH}/src/github.com/pensando/sw/nic/delphi/proto/delphi

build:
	go build

test: client_test.delphi.go
	go test -cover
