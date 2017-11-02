package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <math.h>
*/
import "C"

import (
	"fmt"
)

func main() {
	// Just an example of using CGO
	n := C.sqrt(4)
	fmt.Println("Started compression tool with options", n)
}
