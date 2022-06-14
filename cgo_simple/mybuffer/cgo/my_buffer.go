package cgo

import "unsafe"

type MyBuffer struct {
	cptr CgoMyBufferHandle
}

func NewMyBuffer(size int) *MyBuffer {
	return &MyBuffer{
		cptr: cgo_NewMyBuffer(size),
	}
}

func (this *MyBuffer) Delete() {
	cgo_DeleteMyBuffer(this.cptr)
}

func (this *MyBuffer) Data() []byte {
	data := cgo_MyBuffer_Data(this.cptr)
	len := cgo_MyBuffer_Size(this.cptr)
	return ((*[1 << 31]byte)(unsafe.Pointer(data)))[0:int(len):int(len)]
}
