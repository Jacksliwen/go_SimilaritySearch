#include <stdbool.h>
#include <stdio.h>
#include <unistd.h>
#ifdef __cplusplus
extern "C" {
#endif
//初始化
int InitFaissEngine(char *set_name, int feature_size);
void Search();
void DeleteFaissEngine(char *set_name);
#ifdef __cplusplus
}
#endif
