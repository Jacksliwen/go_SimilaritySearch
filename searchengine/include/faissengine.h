#include <stdbool.h>
#include <stdio.h>
#include <unistd.h>
#include <string.h>
#ifdef __cplusplus
extern "C" {
#endif
  typedef struct {
    char* set_name_;   //集合名称
    int feature_size_;  //特征向量大小
    int feature_num_;  //加载的特征数量
  } EngineLoadInfo;

  //初始化
  int InitFaissEngine(const char* set_name, const int feature_size);

  int LoadData(const char* set_name, float* allFeatures, const int featureNum);

  int Search(const char* set_name, const float* vfeat, const int vfeat_size,
    const int top_n, long* I, float* D);

  int DeleteFaissEngine(const char* set_name);
  int GetAllEngineNum();

  void GetAllEngineStatus(EngineLoadInfo *info);
#ifdef __cplusplus
}
#endif
