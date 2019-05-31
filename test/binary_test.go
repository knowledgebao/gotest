package test

import (
	"encoding/binary"
	"fmt"
	"testing"
)

func TestBytes(t *testing.T) {
	var i64 int64 = 2323
	var buf64 = make([]byte, 8)
	binary.BigEndian.PutUint64(buf64, uint64(i64))
	fmt.Println(buf64)
	fmt.Println(int64(binary.BigEndian.Uint64(buf64)))

	var i32 int32 = 800
	var buf32 = make([]byte, 4)
	binary.BigEndian.PutUint32(buf32, uint32(i32))
	fmt.Println(buf32)
	fmt.Println(int32(binary.BigEndian.Uint32(buf32)))

	var i16 int16 = 27
	var buf16 = make([]byte, 2)
	binary.BigEndian.PutUint16(buf16, uint16(i16))
	fmt.Println(buf16)
	fmt.Println(int16(binary.BigEndian.Uint16(buf16)))

	binary.LittleEndian.PutUint16(buf16, uint16(i16))
	fmt.Println(buf16)
	fmt.Println(int16(binary.LittleEndian.Uint16(buf16)))
}

type Rect struct {
	Left   int `json:"left"`
	Top    int `json:"top"`
	Width  int `json:"width"`
	Height int `json:"height"`
}

func TestEnlarge(t *testing.T) {
	r := Rect{1074, 546, 56, 56}
	leftRatio := 0.5
	topRatio := 1.0
	heightRatio := 2.5
	widthRatio := 2.0
	r.Left -= int(float64(r.Width) * leftRatio)
	r.Top -= int(float64(r.Height) * topRatio)
	r.Height = int(float64(r.Height) * heightRatio)
	r.Width = int(float64(r.Width) * widthRatio)
	fmt.Println(r)
}
