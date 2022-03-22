/*
字符串,是一个不可改变的字节，字符串的元素不能修改,底层结构是
type StringHeader struct {
    Data uintptr
    Len  int
} */

//[]byte(s)转换模拟实现
func bytes2str(s []byte) (p string)){
	data : = make([]byte,len(s))
	for i,v := range s{
		data[i] = v
	}
	 hdr := (*reflect.StringHeader)(unsafe.Pointer(p))
	 hdr.Data = uintptr(unsafe.Pointer(&data))
	 hdr.Len = len(s)
	 return p
}

//string(bytes)转换模拟实现
func str2bytes(s string) []byte {
    p := make([]byte, len(s))
    for i := 0; i < len(s); i++ {
        c := s[i]
        p[i] = c
    }
    return p
}

package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	a := "hello,world"
	hello := a[:5]
	world := a[6:]
	fmt.Printf("%s,%s,%v\n", hello, world, uintptr(unsafe.Pointer(&hello[0])))

	b := a
	fmt.Printf("%v\n", (*reflect.StringHeader)(unsafe.Pointer(&a)).Data)
	fmt.Printf("%v\n", (*reflect.StringHeader)(unsafe.Pointer(&b)).Data)
	//输出结果都是5063801,地址一样

	a = "hello,new world"
	fmt.Printf("%v\n", (*reflect.StringHeader)(unsafe.Pointer(&a)).Data)
	//输出结果是5066797，地址改变

}
