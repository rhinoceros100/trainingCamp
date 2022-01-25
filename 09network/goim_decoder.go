package main

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

const (
	PackageLength   int = 4
	HeaderLength    int = 2
	VerLength       int = 2
	OperationLength int = 4
	SeqLength       int = 4
)

type Header struct {
	PackageLen int32
	HeaderLen  int16
	Ver        int16
	Operation  int32
	Seq        int32
}

//解析goim协议 https://github.com/Terry-Mao/goim/blob/master/docs/proto.md
func GoimDecoder(ReceivedMsg *[]byte) (header *Header, body string, err error) {
	//校验协议
	recvMsgLen := len(*ReceivedMsg)
	if recvMsgLen < PackageLength+HeaderLength {
		fmt.Printf("Received msg len too short:%d", recvMsgLen)
		return nil, "", errors.New("Package format error")
	}

	//PackageLength
	var parsedPackageLength = BytesToInt32((*ReceivedMsg)[0:PackageLength])
	if parsedPackageLength != int32(recvMsgLen) {
		fmt.Printf("ParsedPackageLength != recvMsgLen, parsedPackageLength: %d, , recvMsgLen: %d", parsedPackageLength, recvMsgLen)
		return nil, "", errors.New("Package format error")
	}

	//HeaderLength
	var parsedHeaderLength = BytesToInt16((*ReceivedMsg)[PackageLength : PackageLength+HeaderLength])
	if parsedHeaderLength <= 0 || int16(recvMsgLen) < parsedHeaderLength {
		fmt.Printf("parsedHeaderLength <= 0 || int16(recvMsgLen) < parsedHeaderLength, parsedHeaderLength: %d, , recvMsgLen: %d", parsedHeaderLength, recvMsgLen)
		return nil, "", errors.New("Package format error")
	}

	ver := BytesToInt16((*ReceivedMsg)[PackageLength+HeaderLength : PackageLength+HeaderLength+VerLength])
	operation := BytesToInt32((*ReceivedMsg)[PackageLength+HeaderLength+VerLength : PackageLength+HeaderLength+VerLength+OperationLength])
	seq := BytesToInt32((*ReceivedMsg)[PackageLength+HeaderLength+VerLength+OperationLength : PackageLength+HeaderLength+VerLength+OperationLength+SeqLength])
	header = &Header{
		PackageLen: parsedPackageLength,
		HeaderLen:  parsedHeaderLength,
		Ver:        ver,
		Operation:  operation,
		Seq:        seq,
	}
	ParsedBody := (*ReceivedMsg)[parsedHeaderLength:parsedPackageLength]

	return header, string(ParsedBody), nil
}

//字节转换成整形int32
func BytesToInt32(b []byte) int32 {
	bytesBuffer := bytes.NewBuffer(b)
	var tmp int32
	binary.Read(bytesBuffer, binary.BigEndian, &tmp)
	return tmp
}

//字节转换成整形int16
func BytesToInt16(b []byte) int16 {
	bytesBuffer := bytes.NewBuffer(b)
	var tmp int16
	binary.Read(bytesBuffer, binary.BigEndian, &tmp)
	return tmp
}

//拼接字符切片
func CombineBytes(bytes ...[]byte) []byte {
	if len(bytes) < 2 {
		return nil
	}
	cb := append(bytes[0], bytes[1]...)
	for i := 2; i < len(bytes); i++ {
		cb = append(cb, bytes[i]...)
	}
	return cb
}

//整形int32转换成字节
func Int32ToBytes(n int32) []byte {
	tmp := n
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, tmp)
	return bytesBuffer.Bytes()
}

//整形int16转换成字节
func Int16ToBytes(n int16) []byte {
	tmp := n
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, tmp)
	return bytesBuffer.Bytes()
}

func genPackage(header *Header, body string) *[]byte {
	totalLenByte := Int32ToBytes(header.PackageLen)
	headerLenByte := Int16ToBytes(header.HeaderLen)
	verByte := Int16ToBytes(header.Ver)
	operationByte := Int32ToBytes(header.Operation)
	seqByte := Int32ToBytes(header.Seq)
	TransferData := CombineBytes(totalLenByte, headerLenByte, verByte, operationByte, seqByte, []byte(body))

	return &TransferData
}
