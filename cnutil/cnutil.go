package cnutil

// #cgo CFLAGS: -std=c11 -D_GNU_SOURCE
// #cgo LDFLAGS: -L${SRCDIR} -lcnutil -lcnutilxx -lstdc++
// #include "cnutil.h"
import "C"
import "unsafe"

func ConvertBlob(blob []byte) []byte {
	output := make([]byte, 76)
	out := (*C.char)(unsafe.Pointer(&output[0]))

	input := C.CString(string(blob))
	defer C.free(unsafe.Pointer(input))

	size := (C.uint32_t)(len(blob))
	C.convert_blob(input, size, out)
	return output
}

func ValidateAddress(addr string) bool {
	input := C.CString(addr)
	defer C.free(unsafe.Pointer(input))

	size := (C.uint32_t)(len(addr))
	result := C.validate_address(input, size)
	return (bool)(result)
}
