#include <vector>
#include <string>

using namespace std;

class Solution {
  public:
    int minDeletionSize(vector<string>& A) {
      int del_count {};
      size_t elem_size = A.at(0).size();
      
      for (size_t i = 0; i < elem_size; i++) {
        for (size_t j = 0; j < A.size() - 1; j++) {
          if (A.at(j).at(i) > A.at(j + 1).at(i)) {
            del_count++;
            break;
          }
        }
      }
      return del_count;
    }
};

