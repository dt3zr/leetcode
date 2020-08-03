package io.dt3zr.leetcode;

public class Solution {
    enum CharState {
        Start,
        FirstUpper,
        Upper,
        Lower,
        Reject
    }
    //            | Start      | Lower  | FirstUpper | Upper
    // -------------------------------------------------------
    // isLower(c) | Lower      | Lower  | Lower      | Reject
    // isUpper(c) | FirstUpper | Reject | Upper      | Upper

    CharState parseTable[][] = {
            {CharState.Lower, CharState.Lower, CharState.Lower, CharState.Reject},
            {CharState.FirstUpper, CharState.Reject, CharState.Upper, CharState.Upper}
    };

    public boolean detectCapitalUse(String word) {
        char chars[] = word.toCharArray();
        CharState curState = CharState.Start;

        for (int i = 0; i < word.length(); i++) {
            curState = parseTable[indexFromCase(chars[i])][indexFromState(curState)];
            if (curState == CharState.Reject) {
                return false;
            }
        }
        return true;
    }

    private int indexFromCase(char c) {
        if (Character.isLowerCase(c)) {
            return 0;
        }
        return 1;
    }

    private int indexFromState(CharState s) {
        switch (s) {
            case Start:
                return 0;
            case Lower:
                return 1;
            case FirstUpper:
                return 2;
            case Upper:
                return 3;
        }
        return -1;
    }
}
