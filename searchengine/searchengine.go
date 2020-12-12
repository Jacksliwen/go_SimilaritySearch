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

var Sets = make(map[string]*common.FeatureSet)

// AllFeature     [][]float32
// IDInfo         []string

//InitEngine InitEngine
func InitEngine(Setname string) int {
	set := new(common.FeatureSet)
	set.SetName = Setname
	// set.AllFeature = make([][]float32, 0)
	// set.IDInfo = make([]string, 0)
	Sets[Setname] = set
	return int(C.InitFaissEngine(C.CString(Setname), 256))
}

//Addid Addid
func Addid(Setname string, Features []float32, featureInfo string) int {
	set := Sets[Setname]
	set.AllFeature = append(set.AllFeature, Features)
	set.IDInfo = append(set.IDInfo, featureInfo)

	return int(C.LoadData(C.CString(Setname), (*C.float)(unsafe.Pointer((*float32)(&(set.AllFeature[0][0])))), C.int(len(set.IDInfo))))
}

//Search Search
func Search(Setname string, vfeat []float32, vfeatSize int32, topN int32) (searchRets []common.SearchRet, ret int) {
	Cinfos := make([]C.SearchRetInfo, topN*vfeatSize)
	num := int(C.Search(C.CString(Setname), (*C.float)(unsafe.Pointer((*float32)(&(vfeat[0])))), C.int(vfeatSize), C.int(topN), &(Cinfos[0])))
	if 0 >= num {
		return nil, -1
	}

	for i, info := range Cinfos {
		if i == num {
			break
		}
		if (int64)(len(Sets[Setname].IDInfo)) < (int64)(info.id) {
			continue
		}
		searchret := &common.SearchRet{
			SetName:  Setname,
			ID:       Sets[Setname].IDInfo[(int64)(info.id)],
			Distance: (float32)(info.dis),
		}
		searchRets = append(searchRets, *searchret)
	}
	return searchRets, 0
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
