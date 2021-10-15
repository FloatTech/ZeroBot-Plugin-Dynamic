#include <stdint.h>

uintptr_t pluginOpen(const char* path, char** err);
void* pluginLookup(uintptr_t h, const char* name, char** err);
int pluginClose(void* handle, char** err);
