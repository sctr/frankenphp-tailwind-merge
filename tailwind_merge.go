package tailwindmerge

//#include <zend.h>
//#include "tailwind_merge.h"
import "C"
import (
	"unsafe"

	"github.com/sctr/frankenphp-tailwind-merge/pkg/twmerge"
)

func init() {
	C.register_extension()
}

//export go_tailwind_merge
func go_tailwind_merge(strings **C.zend_string, count C.int) *C.char {
	n := int(count)
	if n == 0 {
		return nil
	}

	classes := make([]string, n)
	cStrings := unsafe.Slice(strings, n)
	for i := 0; i < n; i++ {
		classes[i] = zendStringToGoString(cStrings[i])
	}

	merged := twmerge.TwMerge(classes...)
	if merged == "" {
		return nil
	}

	return C.CString(merged)
}
