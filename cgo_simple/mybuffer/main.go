package main

/*
#include <stdio.h>
*/
import "C"

import (
	"cgo_simple/mybuffer/cgo"
	"unsafe"
)

func main() {
	buf := cgo.NewMyBuffer(1024)
	defer buf.Delete()

	copy(buf.Data(), []byte("hello,cgo\x00"))
	C.puts((*C.char)(unsafe.Pointer(&(buf.Data()[0]))))
}
