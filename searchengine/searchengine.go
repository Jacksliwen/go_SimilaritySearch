package searchengine

/*
#cgo  CFLAGS:  -I./include
#cgo  LDFLAGS:  -L./lib -lfaissengine
#include "faissengine.h"
*/
import "C"
import "unsafe"

//InitEngine InitEngine
func InitEngine(Setname string) {
	C.InitFaissEngine(C.CString(Setname), 256)
}

//LoadData LoadData
func LoadData(Setname string, allFeatures *float32, featureNum int) {
	C.LoadData(C.CString(Setname), (*C.float)(unsafe.Pointer(allFeatures)), C.int(featureNum))
}

//Search Search
func Search() {
	//C.Search()
}

//DeleteFaissEngine DeleteFaissEngine
func DeleteFaissEngine() {

}
