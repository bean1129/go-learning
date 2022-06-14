package main

import "C"

import (
	"fmt"

	qsort "cgo_simple/qsort/capi"
)

func main() {

	values := []int64{42, 9, 101, 95, 27, 25}

	qsort.Slice(values, func(i, j int) bool {
		return values[i] < values[j]
	})

	fmt.Println(values)
}
