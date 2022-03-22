/*隐式实现的鸭子类型,如果一个对象只要看起来像是某种接口类型的实现，那么它就可以作为该接口类型使用*/

package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

/*func Fprintf(w io.Writer, format string, args ...interface{}) (int, error)
type io.Writer interface {
    Write(p []byte) (n int, err error)
}

type error interface {
    Error() string
}*/

type UpperWriter struct {
	io.Writer
}

type UpperString string

func (s UpperString) String() string {
	return strings.ToUpper(string(s))
}

func (p *UpperWriter) Write(data []byte) (n int, err error) {
	return p.Writer.Write(bytes.ToUpper(data))
}

func main() {
	fmt.Fprintf(&UpperWriter{os.Stdout}, "hello")
	fmt.Fprintln(os.Stdout, UpperString("world"))

}
