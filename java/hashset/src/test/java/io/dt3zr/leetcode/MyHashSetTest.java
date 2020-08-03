package io.dt3zr.leetcode;

import org.junit.Test;

public class MyHashSetTest {
    @Test
    public void TestMyHashSet() {
        MyHashSet hashSet = new MyHashSet();
        hashSet.add(1);
        hashSet.add(2);
        assert hashSet.contains(1) == true;
        assert hashSet.contains(3) == false;
        hashSet.add(2);
        assert hashSet.contains(2) == true;
        hashSet.remove(2);
        assert hashSet.contains(2) == false;
    }
}
