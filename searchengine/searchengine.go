package searchengine

/*
#cgo  CFLAGS:  -I./include
#cgo  LDFLAGS:  -L./lib -lfaissengine
#include "faissengine.h"
*/
import "C"
import (
	"SimilaritySearch/common"
	"unsafe"
)

//InitEngine InitEngine
func InitEngine(Setname string) int {
	return int(C.InitFaissEngine(C.CString(Setname), 256))
}

//LoadData LoadData
func LoadData(Setname string, allFeatures *float32, featureNum int) int {
	return int(C.LoadData(C.CString(Setname), (*C.float)(unsafe.Pointer(allFeatures)), C.int(featureNum)))
}

//Search Search
func Search(Setname string, vfeat *float32, vfeatSize int32, topN int32, I *int64, D *float32) int {
	return int(C.Search(C.CString(Setname), (*C.float)(unsafe.Pointer(vfeat)), C.int(vfeatSize), C.int(topN), (*C.long)(unsafe.Pointer(I)), (*C.float)(unsafe.Pointer(D))))
}

//DeleteFaissEngine DeleteFaissEngine
func DeleteFaissEngine(Setname string) int {
	return int(C.DeleteFaissEngine(C.CString(Setname)))
}

//GetAllEngineStatus GetAllEngineStatus
func GetAllEngineStatus() (Engines []common.EngineInfo, ret int) {
	num := (int)(C.GetAllEngineNum())
	if num < 0 {
		return nil, -1
	}
	Cinfos := make([]C.EngineLoadInfo, num)
	C.GetAllEngineStatus(&(Cinfos[0]))
	for _, info := range Cinfos {
		sLen := int(C.strlen(info.set_name_))
		s1 := string((*[31]byte)(unsafe.Pointer(info.set_name_))[:sLen:sLen])

		engine := &common.EngineInfo{
			SetName:     s1,
			FeatureSize: (int)(info.feature_size_),
			FeatureNum:  (int)(info.feature_num_),
		}
		Engines = append(Engines, *engine)
	}
	ret = 0
	return
}
