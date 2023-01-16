package main

import (
	"errors"
	"fmt"
)

type Object interface {
}

// 链表结构体参数：1.该节点的值 2.下个节点的指针
type Node struct {
	Data Object
	Next *Node
}

type List struct {
	head *Node
}

// 获取链表长度
func (l *List) Len() int {
	var len int
	if l.IsEmpty() {
		return 0
	}
	pre := l.head
	for {
		len++
		pre = pre.Next
		if pre == nil {
			break
		}
	}
	return len
}

// 批量插入节点
func (l *List) BatchAppend(values []int) {
	for _, v := range values {
		l.Append(v)
	}
}

// 从链表头部插入节点
func (l *List) HeadAppend(value interface{}) {
	newHead := &Node{
		Data: value,
		Next: l.head,
	}
	l.head = newHead
}

// 从链表尾部插入节点
func (l *List) Append(value interface{}) {
	if l.IsEmpty() {
		l.head = &Node{
			Data: value,
		}
		return
	}
	cur := l.head
	for {
		if cur.Next == nil {
			cur.Next = &Node{
				Data: value,
			}
			break
		} else {
			cur = cur.Next
		}
	}
}

// 获取第index节点的值
func (l *List) GetIndex(index int) (value interface{}) {
	if l.IsEmpty() {
		fmt.Println(errors.New("链表为空"))
		return 0
	}
	if l.Len() < index+1 {
		fmt.Println(errors.New("index超出链表长度"))
		return 0
	}
	if index < 0 {
		fmt.Println(errors.New("index不能为负数"))
		return 0
	}
	n := l.head
	for i := 0; i < index; i++ {
		n = n.Next
	}
	return n.Data
}

// 获取第一个值为value的位置
func (l *List) GetValue(value interface{}) (index int) {
	if l.IsEmpty() {
		return 0
	}
	n := l.head
	for {
		if n.Data == value {
			return index
		}
		index++
		n = n.Next
		if n == nil {
			break
		}
	}
	fmt.Println(errors.New("不存在该值"))
	return -1
}

// 获取所有值为value的位置
//func (l *List) GetAllByValue(value interface{}){}

// 删除第index节点
func (l *List) DeleteIndex(index int) {
	if l.IsEmpty() {
		fmt.Println(errors.New("链表为空"))
		return
	}
	if l.Len() < index+1 {
		fmt.Println(errors.New("index超出链表长度"))
		return
	}
	if index < 0 {
		fmt.Println(errors.New("index不能为负数"))
		return
	}
	if index == 0 {
		l.head = l.head.Next
		return
	}
	n := l.head
	for i := 0; i < index-1; i++ {
		n = n.Next
	}
	n = n.Next.Next
	return
}

// 删除值为value的第一个值
func (l *List) DeletVaule() {

}

// 删除链表中的重复值

// 排序链表
func (l *List) Sort() {

}

// 打印链表
func (l *List) PrintList() {
	//头节点为空，打印空
	if l.IsEmpty() {
		fmt.Println("")
		return
	}
	fmt.Print(l.head.Data)
	cur := l.head
	for i := 0; i < l.Len(); i++ {
		cur = cur.Next
		if cur != nil {
			fmt.Print("->")
			fmt.Print(cur.Data)
		}
	}
	fmt.Println("")
}

// 判空
func (l *List) IsEmpty() bool {
	if l.head == nil {
		return true
	} else {
		return false
	}
}

func main() {
	var l List
	l.HeadAppend(1)
	l.Append(2)
	l.HeadAppend(1)
	l.Append(2)
	l.HeadAppend(3)
	fmt.Println(l.GetValue(3))
	fmt.Println(l.GetValue(1))
	fmt.Println(l.GetValue(2))
	fmt.Println(l.GetValue(5))
	fmt.Println(l.Len())
	l.DeleteIndex(1)
	l.PrintList()
}
