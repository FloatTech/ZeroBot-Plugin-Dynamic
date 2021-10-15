// +build !windows,cgo

package plugin

// #cgo linux LDFLAGS: -ldl
// #include "dlopen.h"
import "C"

import (
	"errors"
	"unsafe"
)

func open(path string) (unsafe.Pointer, unsafe.Pointer, error) {
	cPath := path + "\000"
	var cErr *C.char
	h := C.pluginOpen(cChar(&cPath), &cErr)
	if h == 0 {
		return nil, nil, errors.New(C.GoString(cErr))
	}
	cname := "Hook\000"
	hook := C.pluginLookup(h, cChar(&cname), &cErr)
	if hook == nil {
		return *(*unsafe.Pointer)(unsafe.Pointer(&h)), nil, errors.New(C.GoString(cErr))
	}
	return *(*unsafe.Pointer)(unsafe.Pointer(&h)), hook, nil
}

func close(handle unsafe.Pointer) error {
	var cErr *C.char
	res := C.pluginClose(handle, &cErr)
	if res != 0 {
		return errors.New(`plugin.Close: ` + C.GoString(cErr))
	}
	return nil
}
