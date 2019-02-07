include ../Makefile
BINARYNAME=udb-server
BINARYPATH=target
BINARY=${BINARYPATH}/${BINARYNAME}

SRCPATH=cmd/udb-server/main.go

build:
	go build -o ${BINARY} -v ${SRCPATH}

run:
	./${BINARY}
