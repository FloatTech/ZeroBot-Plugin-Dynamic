// +build !windows,cgo

// #cgo linux LDFLAGS: -ldl
#include <dlfcn.h>
#include <limits.h>
#include <stdlib.h>
#include <stdint.h>
#include <stdio.h>
#include "dlopen.h"
uintptr_t pluginOpen(const char* path, char** err) {
	void* h = dlopen(path, RTLD_LAZY);
	if (h == NULL) {
		*err = (char*)dlerror();
	}
	return (uintptr_t)h;
}
void* pluginLookup(uintptr_t h, const char* name, char** err) {
	void* r = dlsym((void*)h, name);
	if (r == NULL) {
		*err = (char*)dlerror();
	}
	return r;
}
int pluginClose(void* handle, char** err) {
	int res = dlclose(handle);
	if (res != 0) {
		*err = (char*)dlerror();
	}
	return res;
}
