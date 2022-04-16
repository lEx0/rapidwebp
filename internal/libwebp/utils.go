package libwebp

/*
#cgo CFLAGS: -I../source
#include <stdlib.h>
#include <decode.h>
*/
import "C"

// CGO is not supported in test files
// use CGO in utils.go

func bool2CInt(v bool) C.int {
	if v {
		return toCInt(1)
	}

	return toCInt(0)
}

func toCInt(v int) C.int {
	return C.int(v)
}
