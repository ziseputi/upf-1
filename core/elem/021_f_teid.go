package elem

import (
	"bytes"
)

type FTEIDFlag byte

/**
const (
	FTEIDFlagIPv4    FTEIDFlag = 1 //bits:0000 0001
	FTEIDFlagIPv6    FTEIDFlag = 2 //bits:0000 0010
	FTEIDFlagIPv46   FTEIDFlag = 3 //bits:0000 0011
	FTEIDFlagCH      FTEIDFlag = 4 //bits:0000 0100
	FTEIDFlagCHIPv4  FTEIDFlag = 5 //bits:0000 0100
	FTEIDFlagCHIPv6  FTEIDFlag = 6 //bits:0000 0100
	FTEIDFlagCHIPv46 FTEIDFlag = 7 //bits:0000 0100
)*/

type FTEID struct {
	EType    IEType
	ELength  uint16
	Flag     FTEIDFlag
	TEID     []byte //4byte
	IPv4Addr []byte //4byte
	IPv6Addr []byte //16byte
	ChooseID byte
}

func DecodeFTEID(buf *bytes.Buffer, len uint16) *FTEID {
	f := FTEID{
		EType:   IETypeFTEID,
		ELength: len,
		Flag:    FTEIDFlag(getValue(buf, 1)[0]),
	}
	if f.Flag&0b00000100>>2 == 1 { //CH=1
		f.ChooseID = getValue(buf, 1)[0]
	} else {
		f.TEID = getValue(buf, 4)
		if f.Flag&0b00000001 == 1 { //V4=1
			f.IPv4Addr = getValue(buf, 4)
		}
		if f.Flag&0b00000010>>1 == 1 { //V6=1
			f.IPv6Addr = getValue(buf, 16)
		}
		if f.Flag&0b00001000>>3 == 1 { //CHID=1
			f.ChooseID = getValue(buf, 1)[0]
		}
	}
	//log.Println("FTEID: ", f)
	return &f
}

func EncodeFTEID(f FTEID) *bytes.Buffer {
	ret := SetValue(f.EType, f.ELength, f.Flag)
	if f.Flag&0b00000100>>2 == 0 {
		SetValue(ret, f.TEID)
		if f.Flag&0b00000001 == 1 {
			SetValue(ret, f.IPv4Addr)
		}
		if f.Flag&0b00000010>>1 == 1 {
			SetValue(ret, f.IPv6Addr)
		}
		if f.Flag&0b00001000>>3 == 1 {
			SetValue(ret, f.ChooseID)
		}
	}
	return ret
}

//判断是否含有FTEID
func HasFTEID(f FTEID) bool {
	if f.EType == 0 {
		return false
	}
	return true
}
