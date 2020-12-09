#ifndef ___FAISSENGINE___
#define ___FAISSENGINE___
#include <iostream>
#include <mutex>
#include <utility>
#include "faiss/IndexFlat.h"
#include "faiss/gpu/GpuAutoTune.h"
#include "faiss/gpu/GpuCloner.h"
#include "faiss/gpu/GpuClonerOptions.h"
#include "faiss/gpu/GpuIndexFlat.h"
#include "faiss/gpu/StandardGpuResources.h"
#include "faiss/gpu/GpuCloner.h"
#include "faiss/gpu/GpuClonerOptions.h"
#include "faiss/gpu/utils/DeviceUtils.h"
#include <malloc.h>
class FaissEngine {
 public:
 FaissEngine(const int feature_size) { feature_size_ = feature_size; }

  ~FaissEngine() { Close(); }

  //初始化
  bool Init();

  
  void Close();

  //添加带检索数据
  bool Add(const float *feats, int nums);

  //检索
  bool Search(const float *feat,int num, int top_N,
                         std::vector<std::vector<std::pair<int64_t, float>>> *result);

 private:
  int feature_size_;
  int feature_nums_;
  int gpu_nums_;
  std::mutex mutex_;
  std::vector<int> devs_;
  faiss::Index *gpu_index_;
  faiss::IndexFlatL2 *cpu_index_;
  faiss::gpu::GpuMultipleClonerOptions co_;
  std::vector<faiss::gpu::GpuResources *> gpu_resource_;
};

#endif

