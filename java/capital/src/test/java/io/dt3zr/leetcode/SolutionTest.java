package io.dt3zr.leetcode;

import org.junit.Test;

public class SolutionTest {
    @Test
    public void TestDetectCapitalUse() {
        Solution sln = new Solution();
        assert sln.detectCapitalUse("USA") == true;
        assert sln.detectCapitalUse("Usa") == true;
        assert sln.detectCapitalUse("usa") == true;
        assert sln.detectCapitalUse("uSa") == false;
        assert sln.detectCapitalUse("leetcode") == true;
        assert sln.detectCapitalUse("Leetcode") == true;
        assert sln.detectCapitalUse("LEETcode") == false;
    }

}
