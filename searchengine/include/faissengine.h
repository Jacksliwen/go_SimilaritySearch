#include <stdio.h>
#include <unistd.h>
#include <stdbool.h>
#ifdef __cplusplus
extern "C" {
#endif
typedef struct  {
  int id;
  float Dis;
}stuResult_;

//初始化
bool Init(int feature_size);


void Close();

//添加带检索数据
bool Add(const float *feats, int nums);

//检索
bool Search(const float *feat,int num, int top_N, stuResult_ **results);

#ifdef __cplusplus
}
#endif
