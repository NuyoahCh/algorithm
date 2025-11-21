package main

// 这里因为前面细致的讲解了链表的原理，那么就以手撕 lru 缓存为例，来综合运用链表和哈希表完成一个经典的数据结构设计题目。

/*
请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
实现 LRUCache 类：
LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组 key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。

示例：

输入
["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
输出
[null, null, null, 1, null, -1, null, -1, 3, 4]

解释
LRUCache lRUCache = new LRUCache(2);
lRUCache.put(1, 1); // 缓存是 {1=1}
lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
lRUCache.get(1);    // 返回 1
lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
lRUCache.get(2);    // 返回 -1 (未找到)
lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
lRUCache.get(1);    // 返回 -1 (未找到)
lRUCache.get(3);    // 返回 3
lRUCache.get(4);    // 返回 4
*/

type Node struct {
	k, v      int
	pre, next *Node
}

type LRUCache struct {
	cap        int
	head, tail *Node
	m          map[int]*Node
}

func Constructor(capacity int) LRUCache {
	h, t := &Node{}, &Node{}
	h.next = t
	t.pre = h
	return LRUCache{capacity, h, t, make(map[int]*Node)}
}

func (this *LRUCache) Get(key int) int {
	cache := this.m
	if node, ok := cache[key]; ok {
		this.moveToHead(node)
		return node.v
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	cache := this.m
	if node, ok := cache[key]; ok {
		this.moveToHead(node)
		node.v = value
	} else {
		if len(cache) >= this.cap {
			deleteNode := this.tail.pre
			delete(cache, deleteNode.k)
			this.deleteNodeFromList(deleteNode)
		}
		newNode := &Node{key, value, nil, nil}
		cache[key] = newNode
		this.addNodeToHead(newNode)
	}
	return
}

func (this *LRUCache) moveToHead(node *Node) {
	this.deleteNodeFromList(node)
	this.addNodeToHead(node)
}

func (this *LRUCache) deleteNodeFromList(node *Node) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (this *LRUCache) addNodeToHead(node *Node) {
	node.next = this.head.next
	node.pre = this.head
	this.head.next.pre = node
	this.head.next = node
}
