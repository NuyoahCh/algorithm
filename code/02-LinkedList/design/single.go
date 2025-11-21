package main

import (
	"errors"
)

// Node 节点结构
type Node01[E any] struct {
	val  E
	next *Node01[E]
}

// MyLinkedList2 链表结构
type MyLinkedList2[E any] struct {
	head *Node01[E]
	// 实际的尾部节点引用
	tail  *Node01[E]
	size_ int
}

// NewMyLinkedList2 创建一个新的链表
func NewMyLinkedList2[E any]() *MyLinkedList2[E] {
	head := &Node01[E]{}
	return &MyLinkedList2[E]{head: head, tail: head, size_: 0}
}

// addFirst 在头部添加元素
func (list *MyLinkedList2[E]) AddFirst(e E) {
	newNode := &Node01[E]{val: e}
	newNode.next = list.head.next
	list.head.next = newNode
	if list.size_ == 0 {
		list.tail = newNode
	}
	list.size_++
}

// addLast 在尾部添加元素
func (list *MyLinkedList2[E]) AddLast(e E) {
	newNode := &Node01[E]{val: e}
	list.tail.next = newNode
	list.tail = newNode
	list.size_++
}

// add 在指定索引处添加元素
func (list *MyLinkedList2[E]) Add(index int, element E) error {
	if index < 0 || index > list.size_ {
		return errors.New("index out of bounds")
	}

	if index == list.size_ {
		list.AddLast(element)
		return nil
	}

	prev := list.head
	for i := 0; i < index; i++ {
		prev = prev.next
	}
	newNode := &Node01[E]{val: element}
	newNode.next = prev.next
	prev.next = newNode
	list.size_++
	return nil
}

// removeFirst 移除头部元素
func (list *MyLinkedList2[E]) RemoveFirst() (E, error) {
	if list.IsEmpty() {
		return *new(E), errors.New("no elements to remove")
	}
	first := list.head.next
	list.head.next = first.next
	if list.size_ == 1 {
		list.tail = list.head
	}
	list.size_--
	return first.val, nil
}

// removeLast 移除尾部元素
func (list *MyLinkedList2[E]) RemoveLast() (E, error) {
	if list.IsEmpty() {
		return *new(E), errors.New("no elements to remove")
	}

	prev := list.head
	for prev.next != list.tail {
		prev = prev.next
	}
	val := list.tail.val
	prev.next = nil
	list.tail = prev
	list.size_--
	return val, nil
}

// remove 在指定索引处移除元素
func (list *MyLinkedList2[E]) Remove(index int) (E, error) {
	if index < 0 || index >= list.size_ {
		return *new(E), errors.New("index out of bounds")
	}

	prev := list.head
	for i := 0; i < index; i++ {
		prev = prev.next
	}

	nodeToRemove := prev.next
	prev.next = nodeToRemove.next
	// 删除的是最后一个元素
	if index == list.size_-1 {
		list.tail = prev
	}
	list.size_--
	return nodeToRemove.val, nil
}

// GetFirst 获取头部元素
func (list *MyLinkedList2[E]) GetFirst() (E, error) {
	if list.IsEmpty() {
		return *new(E), errors.New("no elements in the list")
	}
	return list.head.next.val, nil
}

// GetLast 获取尾部元素
func (list *MyLinkedList2[E]) GetLast() (E, error) {
	if list.IsEmpty() {
		return *new(E), errors.New("no elements in the list")
	}
	return list.tail.val, nil
}

// Get 获取指定索引的元素
func (list *MyLinkedList2[E]) Get(index int) (E, error) {
	if index < 0 || index >= list.size_ {
		return *new(E), errors.New("index out of bounds")
	}
	return list.getNode(index).val, nil
}

// Set 更新指定索引的元素
func (list *MyLinkedList2[E]) Set(index int, element E) (E, error) {
	if index < 0 || index >= list.size_ {
		return *new(E), errors.New("index out of bounds")
	}
	node := list.getNode(index)
	oldVal := node.val
	node.val = element
	return oldVal, nil
}

// Size 获取链表大小
func (list *MyLinkedList2[E]) Size() int {
	return list.size_
}

// IsEmpty 检查链表是否为空
func (list *MyLinkedList2[E]) IsEmpty() bool {
	return list.size_ == 0
}

// getNode 返回指定索引的节点
func (list *MyLinkedList2[E]) getNode(index int) *Node01[E] {
	p := list.head.next
	for i := 0; i < index; i++ {
		p = p.next
	}
	return p
}

// func main() {
// 	list := NewMyLinkedList2[int]()
// 	list.addFirst(1)
// 	list.addFirst(2)
// 	list.addLast(3)
// 	list.addLast(4)
// 	list.add(2, 5)

// 	if val, err := list.removeFirst(); err == nil {
// 		fmt.Println(val) // 2
// 	}
// 	if val, err := list.removeLast(); err == nil {
// 		fmt.Println(val) // 4
// 	}
// 	if val, err := list.remove(1); err == nil {
// 		fmt.Println(val) // 5
// 	}

// 	if val, err := list.getFirst(); err == nil {
// 		fmt.Println(val) // 1
// 	}
// 	if val, err := list.getLast(); err == nil {
// 		fmt.Println(val) // 3
// 	}
// 	if val, err := list.get(1); err == nil {
// 		fmt.Println(val) // 3
// 	}
// }
