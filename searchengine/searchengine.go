package searchengine

/*
#cgo  CFLAGS:  -I./include
#cgo  LDFLAGS:  -L./lib -lfaissengine
#include "faissengine.h"
*/
import "C"
import (
	"fmt"
	"time"
)

func Init() {
	C.InitFaissEngine(C.CString("youlike"), 256)
	time.Sleep(time.Duration(2) * time.Second)
	C.DeleteFaissEngine(C.CString("youlike"))
	time.Sleep(time.Duration(5) * time.Second)
	var name string
	fmt.Scanln(&name)
}
