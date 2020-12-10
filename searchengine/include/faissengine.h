#include <stdbool.h>
#include <stdio.h>
#include <unistd.h>
#ifdef __cplusplus
extern "C" {
#endif
  //初始化
  int InitFaissEngine(char* set_name, int feature_size);
  int LoadData(char* set_name, float* allFeatures, int featureNum);
  int Search(char* set_name, float* vfeat, int vfeat_size, const size_t top_n,
    long* I, float* D);
  void DeleteFaissEngine(char* set_name);
#ifdef __cplusplus
}
#endif
