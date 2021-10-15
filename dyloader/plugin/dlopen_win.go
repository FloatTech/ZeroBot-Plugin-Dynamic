// +build windows,cgo

package plugin

// #include "dlopen.h"
import "C"

import (
	"errors"
	"fmt"
	"unsafe"
)

func open(path string) (unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, error) {
	cPath := path + "\000"
	var cErr *C.char
	h := C.pluginOpen(cChar(&cPath), &cErr)
	if h == 0 {
		return nil, nil, nil, errors.New(C.GoString(cErr))
	}
	cname := "Hook\000"
	hook := C.pluginLookup(h, cChar(&cname), &cErr)
	if hook == nil {
		return *(*unsafe.Pointer)(unsafe.Pointer(&h)), nil, nil, errors.New(C.GoString(cErr))
	}
	cname2 := "Inita\000"
	inita := C.pluginLookup(h, cChar(&cname2), &cErr)
	if inita == nil {
		fmt.Println(C.GoString(cErr))
		return *(*unsafe.Pointer)(unsafe.Pointer(&h)), hook, nil, errors.New(C.GoString(cErr))
	}
	return *(*unsafe.Pointer)(unsafe.Pointer(&h)), hook, inita, nil
}

func close(handle unsafe.Pointer) error {
	var cErr *C.char
	res := C.pluginClose(handle, &cErr)
	if res != 0 {
		return errors.New(`plugin.Close: ` + C.GoString(cErr))
	}
	return nil
}
