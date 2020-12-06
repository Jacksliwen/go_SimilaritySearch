package searchengine

/*
#cgo  CFLAGS:  -I${SRCDIR}/include
#cgo  LDFLAGS:  -L${SRCDIR}/lib -lfaissengine
#include "faissengine.h"
*/
import "C"

func SayHello() {
	C.hello()
}
