#include "common/catch.hpp"
#include "delcol.hpp"

TEST_CASE("Column Deletion", "[minDeletionSize]") {
  Solution sln{};

  SECTION("test case 1", "[test1]") {
    vector<string> testData{"cba", "daf", "ghi"};
    REQUIRE(sln.minDeletionSize(testData) == 1);
  }

  SECTION("test case 2", "[test2]") {
    vector<string> testData{"a", "b"};
    REQUIRE(sln.minDeletionSize(testData) == 0);
  }

  SECTION("test case 3", "[test3]") {
    vector<string> testData{"zyx", "wvu", "tsr"};
    REQUIRE(sln.minDeletionSize(testData) == 3);
  }
}
