package main

/*

#cgo CFLAGS: -I${SRCDIR}/foo
#cgo LDFLAGS: -lstdc++ -L${SRCDIR}/foo -lfoo
#include "test.h"
*/
import "C"

import(
	"fmt"
)

func main()  {
	fmt.Println("cpp to go")
	C.bar()
}