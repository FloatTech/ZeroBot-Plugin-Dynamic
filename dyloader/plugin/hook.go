package plugin

// #include "hook.h"
import "C"
import (
	"errors"
	"unsafe"
)

// Plugin is a loaded Go plugin.
type Plugin struct {
	handle unsafe.Pointer
	hook   unsafe.Pointer
	inita  unsafe.Pointer
}

// Open opens a Go plugin.
// If a path has already been opened, then the existing *Plugin is returned.
// It is safe for concurrent use by multiple goroutines.
func Open(path string) (*Plugin, error) {
	handle, hook, inita, err := open(path)
	return &Plugin{handle, hook, inita}, err
}

// Close closes a Go plugin.
// If a path is noth opened, it is ignored.
// It is safe for concurrent use by multiple goroutines.
func Close(p *Plugin) error {
	if p.handle != nil {
		return close(p.handle)
	}
	return errors.New("plugin.Close: handle is nil.")
}

// Hook 改变本插件的环境变量以加载插件
func (p *Plugin) Hook(botconf interface{}, apicallers interface{}, hooknew interface{},
	matlist interface{}, matlock interface{}, defen interface{},
	reg interface{}, del interface{},
	sndgrpmsg interface{}, sndprivmsg interface{}, getmsg interface{},
	parsectx interface{},
	custnode interface{}, pasemsg interface{}, parsemsgfromarr interface{},
) {
	C.hook(p.hook, getdata(&botconf), getdata(&apicallers), getdata(&hooknew), getdata(&matlist), getdata(&matlock), getdata(&defen), getdata(&reg), getdata(&del), getdata(&sndgrpmsg), getdata(&sndprivmsg), getdata(&getmsg), getdata(&parsectx), getdata(&custnode), getdata(&pasemsg), getdata(&parsemsgfromarr))
}

// Init 插件加载
func (p *Plugin) Init() {
	C.init(p.inita)
}

func getdata(ptr *interface{}) *C.char {
	return (*C.char)((*eface)(unsafe.Pointer(ptr)).data)
}
