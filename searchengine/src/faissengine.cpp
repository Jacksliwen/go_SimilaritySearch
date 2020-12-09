#include "faissengine.h"
#include <sstream>
bool FaissEngine::Init() {
  try {
    int is_use_float16 = 1;                    // 默认使用float16
    int reserved_mem_size = 10 * 1024 * 1024;  // 默认大小10M
    gpu_nums_ = faiss::gpu::getNumDevices();
    std::cout << "FaissEngine init, is_use_float16 =" << is_use_float16
              << ", reserved_mem_size = " << reserved_mem_size
              << ", all gpu nums = " << gpu_nums_
              << ", feature_size = " << feature_size_ << std::endl;
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

void FaissEngine::Close() {
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

bool FaissEngine::Add(const float *feats, int nums) {
  std::lock_guard<std::mutex> lk(mutex_);
  std::cout << "gpu_index_ add index,num = " << nums << std::endl;
  if (nums <= 0) {
    return true;
  }
  feature_nums_ =nums ;
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


bool FaissEngine::Search(
    const float *feat, int num, int top_N,
    std::vector<std::vector<std::pair<int64_t, float>>> *result) {
  int top_k = top_N > feature_nums_ ? feature_nums_ : top_N;
  int64_t *I = new int64_t[top_k * num];
  float *D = new float[top_k * num];
  // gpu_resource_非线程安全
  std::lock_guard<std::mutex> lk(mutex_);
  try {
    gpu_index_->search(num, feat, top_k, D, I);
  } catch (faiss::FaissException &e) {
    std::cout << "faiss search exception, e = " << e.what() << std::endl;
    delete[] I;
    delete[] D;
    return false;
  }
  for (int i = 0; i < top_k * num;) {
    std::vector<std::pair<int64_t, float>> each_result;
    for (int j = 0; j < top_k; j++) {
      each_result.push_back(std::make_pair(I[i], D[i]));
      i++;
    }
    result->push_back(each_result);
  }

  delete[] I;
  delete[] D;
  return true;
}

// size_t FaissEngine::FindClosetRecord(
//     const std::unordered_map<std::string, IdObject> &map,
//     const feat_t &all_cls_center, const std::vector<std::string> &id_index,
//     const std::unordered_map<std::string, size_t> &id_reverted,
//     const vfeat_t &vfeat, const size_t feat_size, const float vote_threshold,
//     size_t top_n, std::vector<std::vector<IdObject>> *res) {
//   res->clear();
//   if (map.empty() || top_n < 1) {
//     return 0;
//   }
//   size_t close_num = std::min(top_n, map.size());
//   std::vector<std::vector<std::pair<int64_t, float>>> computer_result;
//   // 1.检索
//   std::vector<float> vec_points(feature_size_ * vfeat.size());
//   float *pointer_to_points = &(vec_points[0]);  // 100W: 1G
//   for (size_t i = 0; i < vfeat.size(); ++i) {
//     for (size_t j = 0; j < feature_size_; ++j) {
//       pointer_to_points[i * feature_size_ + j] = vfeat[i][j];
//     }
//   }
//   std::cout << "vfeat size = " << vfeat.size() << " ,close_num = " << close_num << std::endl;

//   if (!Search(pointer_to_points, vfeat.size(), close_num, &computer_result)) {
//     std::cout << "faiss_wapper search error!" << std::endl;
//     return 0;
//   }
//   try {
//     // 2.组装信息
//     for (int j = 0; j < computer_result.size(); j++) {
//       std::vector<IdObject> vecTemp;
//       for (size_t i = 0; i < computer_result[j].size(); ++i) {
//         IdObject temp;
//         auto id_tt = id_index[computer_result[j][i].first];
//         std::cout << "id = " << id_tt << "computer_result = " << computer_result[j][i].second
//                   << " ,similar = " << (2 - computer_result[j][i].second) / 2 << std::endl;
//         auto iter = map.find(id_tt);
//         if (iter == map.end()) {
//           std::cout << "id_index cannot find in map. id_index = "
//                    << id_index[computer_result[j][i].first] << std::endl;
//           continue;
//         }
//         auto idtemp = iter->second;
//         temp.id = idtemp.id;
//         temp.group_id = idtemp.group_id;
//         temp.features = idtemp.features;
//         temp.similar = computer_result[j][i].second;
//         temp.IPC_id = idtemp.IPC_id;
//         temp.snap_time = idtemp.snap_time;
//         temp.face_id = idtemp.face_id;

//         vecTemp.emplace_back(temp);
//       }
//       res->push_back(vecTemp);
//     }
//   } catch (std::exception &e) {
//     std::cout << "FindClosetRecord Execption,e = " << e.what() <<std::endl;
//     return false;
//   }
//   return 1;
// }
