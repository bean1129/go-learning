/* type SliceHeader struct {
    Data uintptr
    Len  int
    Cap  int
} */

package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"regexp"
	"unsafe"
)

//切片内存技巧
func TrimSpace(s []byte) []byte {
	b := s[:0]
	for _, v := range s {
		if v != ' ' {
			b = append(b, v)
		}
	}
	return b
}

//切片不会改变底层数据存储，如果一个很大的切片，我们只需要其中一段数据，那就应该考虑重新拷贝一块内存
func FindPhoneNumber(filename string) []byte {
	b, _ := ioutil.ReadFile(filename)
	b = regexp.MustCompile("[0-9]+").Find(b)
	//return regexp.MustCompile("[0-9]+").Find(b) //不可取
	return append([]byte{}, b...)
}

func main() {
	var g = make([]int, 3)
	var e = g[:2]
	fmt.Println(cap(g), cap(e), len(g), len(e), (*reflect.SliceHeader)(unsafe.Pointer(&g)), (*reflect.SliceHeader)(unsafe.Pointer(&e)))
	e = append(e, 8, 8, 8)
	/*如果被调用函数中修改了Len或Cap信息的话，就无法反映到调用参数的切片中，
	这时候我们一般会通过返回修改后的切片来更新之前的切片。这也是为何内置的append必须要返回一个切片的原因。*/
	fmt.Println(cap(g), cap(e), len(g), len(e), (*reflect.SliceHeader)(unsafe.Pointer(&g)), (*reflect.SliceHeader)(unsafe.Pointer(&e)))
	/* 	输出结果：
	   	3 3 3 2 &{824633787424 3 3} &{824633787424 2 3}
	   	3 6 3 5 &{824633787424 3 3} &{824633762800 5 6} */
	//------------------
	s := "a bfdafsd 88"
	fmt.Println(string(TrimSpace([]byte(s))))
}
