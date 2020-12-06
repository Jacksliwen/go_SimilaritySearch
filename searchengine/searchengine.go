package searchengine

/*
#cgo  CPPFLAGS:  -I./include
#cgo  LDFLAGS:  -L./lib  -lfaissengine -Wl,-rpath=./lib
#include "faissengine.h"
*/
import "C"

func SayHello() {
	C.hello(C.CString("call C hello func"))
}
