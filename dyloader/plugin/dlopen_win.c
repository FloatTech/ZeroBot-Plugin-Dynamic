// +build windows,cgo

#include <limits.h>
#include <stdlib.h>
#include <stdint.h>
#include <stdio.h>
#include <windows.h>
#include "dlopen.h"
#define ERRBUF 20
static char *dlerror() {
	char *err=(char *) malloc(ERRBUF);
	sprintf(err, "error %i", GetLastError());
	return err;
}
uintptr_t pluginOpen(const char* path, char** err) {
	void* h = LoadLibrary(path);
	if (h == NULL) {
		*err = (char*)dlerror();
	}
	return (uintptr_t)h;
}
void* pluginLookup(uintptr_t h, const char* name, char** err) {
	void* r = GetProcAddress((void*)h, name);
	if (r == NULL) {
		*err = (char*)dlerror();
	}
	return r;
}
int pluginClose(void* handle, char** err) {
	int res = FreeLibrary(handle);
	if (res != 0) {
		*err = (char*)dlerror();
	}
	return res;
}
