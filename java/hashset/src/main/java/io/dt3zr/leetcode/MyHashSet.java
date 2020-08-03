package io.dt3zr.leetcode;

class MyHashSet {

    private class LinkedList {
        private class Node {
            private int value;
            private Node next;

            public Node(int value) {
                this.value = value;
                this.next = null;
            }
        }

        private Node head;

        public LinkedList() {
            head = new Node(-1);
        }

        public void put(int value) {
            Node node = get(value);
            if (node == null) {
                Node n = new Node(value);
                n.next = head.next;
                head.next = n;
            }
        }

        public Node get(int value) {
            for (Node cur = head.next; cur != null; cur = cur.next) {
                if (cur.value == value) {
                    return cur;
                }
            }
            return null;
        }

        public void remove(int value) {
            for (Node cur = head; cur != null && cur.next != null; cur = cur.next) {
                if (cur.next.value == value) {
                    cur.next = cur.next.next;
                }
            }
        }

        public boolean contain(int value) {
            return get(value) != null;
        }
    }

    private final int SET_SIZE = 1999;
    private LinkedList list[];

    /** Initialize your data structure here. */
    public MyHashSet() {
        list = new LinkedList[SET_SIZE];
    }

    public void add(int key) {
        int index = hash(key);
        if (list[index] == null) {
            list[index] = new LinkedList();
        }
        list[index].put(key);
    }

    public void remove(int key) {
        int index = hash(key);
        if (list[index] != null) {
            list[index].remove(key);
        }
    }

    /** Returns true if this set contains the specified element */
    public boolean contains(int key) {
        int index = hash(key);
        if (list[index] == null) {
            return false;
        }
        return list[index].contain(key);
    }

    private int hash(int value) {
        return value % SET_SIZE;
    }
}