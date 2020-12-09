package searchengine

/*
#cgo  CFLAGS:  -I${SRCDIR}/include
#cgo  LDFLAGS:  -L${SRCDIR}/lib -lfaissengine
#include "faissengine.h"
#ifdef __cplusplus
extern "C" {
	#endif
	bool Init(){
		FaissEngine *engine = new FaissEngine(256);
		if(engine->Init()){
			printf("c++ init OK \n");
			return true;
		}
		printf("c++ init Failed \n");
		return false;
	}

	#ifdef __cplusplus
}
#endif
*/
import "C"
import (
	"fmt"
)

func Init() {
	if C.Init() == true {
		fmt.Println("go   init ok")
	}
}
