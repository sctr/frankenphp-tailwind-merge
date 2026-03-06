package tailwindmerge

// #include <string.h>
//
// // memrchr requires _GNU_SOURCE on glibc systems. If the #cgo CFLAGS
// // directive in tailwind_merge.go ever stops defining _GNU_SOURCE this
// // test will fail to compile, catching the regression early.
// static const void *test_memrchr(const void *s, int c, size_t n) {
//     return memrchr(s, c, n);
// }
import "C"

import (
	"testing"
	"unsafe"
)

func TestMemrchrAvailable(t *testing.T) {
	s := "hello world"
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))

	result := C.test_memrchr(unsafe.Pointer(cs), C.int('o'), C.size_t(len(s)))
	if result == nil {
		t.Fatal("memrchr returned nil, expected a pointer to the last 'o' in 'hello world'")
	}
}
