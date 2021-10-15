package plugin

import "C"
import "unsafe"

type String struct {
	Data unsafe.Pointer
	Len  int
}

func cChar(s *string) *C.char {
	a := (*String)(unsafe.Pointer(s))
	return (*C.char)(a.Data)
}

// 没有方法的interface
type eface struct {
	_type uintptr
	data  unsafe.Pointer
}
