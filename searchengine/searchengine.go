package searchengine

/*
#cgo  CFLAGS:  -I./include
#cgo  LDFLAGS:  -L./lib -lfaissengine
#include "faissengine.h"
*/
import "C"

func Init() {
	C.InitFaissEngine()
}
