#include "faissengine.h"

#include <malloc.h>
#include <stdio.h>

#include <iostream>
#include <map>
#include <mutex>
#include <sstream>
#include <utility>
#include <vector>

#include "faiss/IndexFlat.h"
#include "faiss/gpu/GpuAutoTune.h"
#include "faiss/gpu/GpuCloner.h"
#include "faiss/gpu/GpuClonerOptions.h"
#include "faiss/gpu/GpuIndexFlat.h"
#include "faiss/gpu/StandardGpuResources.h"
#include "faiss/gpu/utils/DeviceUtils.h"

class FaissEngine {
 public:
  FaissEngine(int feature_size) { feature_size_ = feature_size; }

 public:
  bool Init() {
    std::lock_guard<std::mutex> lk(mutex_);
    try {
      int is_use_float16 = 1;                    // 默认使用float16
      int reserved_mem_size = 10 * 1024 * 1024;  // 默认大小10M
      gpu_nums_ = faiss::gpu::getNumDevices();
      std::cout << "FaissEngine init, is_use_float16 =" << is_use_float16
                << ", reserved_mem_size = " << reserved_mem_size
                << ", all gpu nums = " << gpu_nums_
                << ", feature_size_ = " << feature_size_ << std::endl;
      for (int i = 0; i < gpu_nums_; i++) {
        auto res = new faiss::gpu::StandardGpuResources;
        res->setTempMemory(reserved_mem_size);
        gpu_resource_.push_back(res);
        devs_.push_back(i);
      }

      cpu_index_ = new faiss::IndexFlatL2(feature_size_);
      co_.shard = true;
      if (is_use_float16 == 1) {
        co_.useFloat16 = true;
      }
      gpu_index_ = faiss::gpu::index_cpu_to_gpu_multiple(gpu_resource_, devs_,
                                                         cpu_index_, &co_);
      std::cout << "FaissEngine init successed!\n";
    } catch (faiss::FaissException &e) {
      std::cout << "FaissEngine Init Execption,e = " << e.what() << std::endl;
      return false;
    }
    return true;
  }

  void Close() {
    std::lock_guard<std::mutex> lk(mutex_);
    if (gpu_index_) {
      delete gpu_index_;
      gpu_index_ = nullptr;
    }
    if (cpu_index_) {
      delete cpu_index_;
      cpu_index_ = nullptr;
    }

    for (size_t i = 0; i < gpu_resource_.size(); ++i) {
      if (gpu_resource_[i]) {
        delete gpu_resource_[i];
        gpu_resource_[i] = nullptr;
      }
    }
    gpu_resource_.clear();
  }

  bool Add(const float *feats, int nums) {
    std::lock_guard<std::mutex> lk(mutex_);
    std::cout << "gpu_index_ add index,num = " << nums << std::endl;
    if (nums <= 0) {
      return true;
    }
    feature_nums_ = nums;
    gpu_index_->reset();
    //每次全量加载
    if (gpu_index_) {
      delete gpu_index_;
      gpu_index_ = nullptr;
    }
    try {
      gpu_index_ = faiss::gpu::index_cpu_to_gpu_multiple(gpu_resource_, devs_,
                                                         cpu_index_, &co_);
      gpu_index_->add(nums, feats);
    } catch (faiss::FaissException &e) {
      std::cout << "FaissEngine Add Execption,e = " << e.what() << std::endl;
      return false;
    }
    return true;
  }

  int Search(const float *feat, const int feat_num, const int top_N, int64_t *I,
             float *D) {
    size_t top_k = std::min(top_N, feature_nums_);
    I = new int64_t[top_k * feat_num];
    D = new float[top_k * feat_num];
    // gpu_resource_非线程安全
    std::lock_guard<std::mutex> lk(mutex_);
    try {
      gpu_index_->search(feat_num, feat, top_k, D, I);
    } catch (faiss::FaissException &e) {
      std::cout << "faiss search exception, e = " << e.what() << std::endl;
      delete[] I;
      delete[] D;
      return -1;
    }
    return 0;
  }

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

std::map<std::string, std::shared_ptr<FaissEngine>> mapSet_Faissengine;

int InitFaissEngine(char *set_name, int feature_size) {
  if (set_name != nullptr) {
    auto iter = mapSet_Faissengine.find(set_name);
    if (mapSet_Faissengine.end() == iter) {
      auto engine = std::make_shared<FaissEngine>(feature_size);
      if (engine->Init() != true) {
        return -1;
      }
      mapSet_Faissengine.insert(std::make_pair(set_name, engine));
      std::cout << "init engine " << set_name << " successed\n";
    } else {
      std::cout << "we have " << set_name << "already\n";
    }
    return 0;
  } else {
    std::cout << "set_name is empty\n";
    return -1;
  }
}

int LoadData(char *set_name, void *allFeatures, int featureNum) {
  float *ptrAllFeatures = (float *)allFeatures;
  auto iter = mapSet_Faissengine.find(set_name);
  if (mapSet_Faissengine.end() != iter) {
    iter->second->Add(ptrAllFeatures, featureNum);
    std::cout << "AddData set_name= " << set_name << " successed\n";
  } else {
    std::cout << "we have not " << set_name << "  engine, may init first\n";
  }
  return 0;
}

int Search(char *set_name, const float *vfeat, int vfeat_size,
           const size_t top_n, int64_t *I, float *D) {
  auto iter = mapSet_Faissengine.find(set_name);
  if (mapSet_Faissengine.end() != iter) {
    if (0 == iter->second->Search(vfeat, vfeat_size, top_n, I, D)) {
      std::cout << "Search set_name= " << set_name << " successed\n";
    }
  } else {
    std::cout << "we have not " << set_name << "  engine, may init first\n";
  }
  return 0;
}

void DeleteFaissEngine(char *set_name) {
  auto iter = mapSet_Faissengine.find(set_name);
  if (iter != mapSet_Faissengine.end()) {
    iter->second->Close();
    mapSet_Faissengine.erase(iter);
    std::cout << "Delete faissengine name = " << set_name << std::endl;
  }
}
