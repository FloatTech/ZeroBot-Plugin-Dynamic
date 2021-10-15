#include "hook.h"
#include <stdio.h>
void* hook(void* ptrf, char* botconf, char* apicallers, char* hooknew, char* matlist, char* matlock, char* defen, char* reg, char* del, char* sndgrpmsg, char* sndprivmsg, char* getmsg, char* parsectx, char* custnode, char* pasemsg, char* parsemsgfromarr) {
	void*(*p)(void*, void*, void*, void*, void*, void*, void*, void*, void*, void*, void*, void*, void*, void*, void*);
	p = ptrf;
	return p(botconf, apicallers, hooknew, matlist, matlock, defen, reg, del, sndgrpmsg, sndprivmsg, getmsg, parsectx, custnode, pasemsg, parsemsgfromarr);
}
void init(void* ptrf) {
	void(*p)();
	p = ptrf;
	p();
}