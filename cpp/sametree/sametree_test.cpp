#include "common/catch.hpp"
#include "sametree.hpp"

TEST_CASE("Identical tree", "[isSameTree]") {
  Solution sln {};

  SECTION("2 equal 3-node trees") {
    TreeNode pl(2);
    TreeNode pr(3);
    TreeNode p(1, &pl, &pr);

    TreeNode ql(2);
    TreeNode qr(3);
    TreeNode q(1, &ql, &qr);

    REQUIRE(sln.isSameTree(&p, &q) == true);
  };

  SECTION("2 unequal 5-node trees") {
    TreeNode a(1);
    TreeNode b(1);
    TreeNode c(2, &a, nullptr);
    TreeNode d(2, nullptr, &b);
    TreeNode p(3, &c, &d);

    TreeNode r(1);
    TreeNode s(1);
    TreeNode t(2, &r, nullptr);
    TreeNode u(2, &s, nullptr);
    TreeNode q(3, &t, &u);

    REQUIRE(sln.isSameTree(&p, &q) == false);

  }
}
