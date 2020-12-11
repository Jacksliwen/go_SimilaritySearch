package searchengine

/*
#cgo  CFLAGS:  -I./include
#cgo  LDFLAGS:  -L./lib -lfaissengine
#include "faissengine.h"
*/
import "C"
import (
	"unsafe"
)

//InitEngine InitEngine
func InitEngine(Setname string) {
	C.InitFaissEngine(C.CString(Setname), 256)
}

//LoadData LoadData
func LoadData(Setname string, allFeatures *float32, featureNum int) {
	C.LoadData(C.CString(Setname), (*C.float)(unsafe.Pointer(allFeatures)), C.int(featureNum))
}

//Search Search
func Search(Setname string, vfeat *float32, vfeat_size int32, top_n int32, I *int64, D *float32) {
	C.Search(C.CString(Setname), (*C.float)(unsafe.Pointer(vfeat)), C.int(top_n), C.int(top_n), (*C.long)(unsafe.Pointer(I)), (*C.float)(unsafe.Pointer(D)))
}

//DeleteFaissEngine DeleteFaissEngine
func DeleteFaissEngine(Setname string) {
	C.DeleteFaissEngine(C.CString(Setname))
}

// func GetAllEngineStatus() (EngineNum int, mapEngine map[string]string) {
// 	//C.EngineLoadInfo* info
// 	return C.GetAllEngineStatus()
// }
