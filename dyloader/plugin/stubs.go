// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !cgo

package plugin

import "errors"
import "unsafe"


func open(name string) (unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, error) {
	return nil, nil, nil errors.New("plugin: not implemented")
}

func close(handle unsafe.Pointer) error {
	return nil, nil, nil, errors.New("plugin: not implemented")
}
