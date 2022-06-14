package cgo

/*
#cgo CXXFLAGS: -std=c++11
#cgo LDFLAGS: -L ../capi -lMyBuffer
#include "../capi/MyBuffer_T.h"
#include <stdio.h>
*/
import "C"

type CgoMyBufferHandle = C.MyBufferHandle

func cgo_NewMyBuffer(size int) CgoMyBufferHandle {
	p := C.NewMyBuffer(C.int(size))
	return (CgoMyBufferHandle)(p)
}

func cgo_DeleteMyBuffer(p CgoMyBufferHandle) {
	C.DeleteMyBuffer((CgoMyBufferHandle)(p))
}

func cgo_MyBuffer_Data(p CgoMyBufferHandle) *C.char {
	return C.MyBufferData((CgoMyBufferHandle)(p))
}

func cgo_MyBuffer_Size(p CgoMyBufferHandle) C.int {
	return C.MyBufferLen((CgoMyBufferHandle)(p))
}
