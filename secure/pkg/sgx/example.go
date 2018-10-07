package main

/*
#include "./Intel/App/TEE.h"
#cgo CFLAGS: -I./Intel/App
#cgo LDFLAGS: -L. -ltee
*/
import "C"

import (
	"fmt"
)


func test() {
	fmt.Printf("intel sgx hello\n")
}

func main() {
	C.testMain()
	// test()
}
