/*
单链表, 线程不安全
 */
package algorithm

type ListNode struct {
	value interface{}
	next  *ListNode
}

func (l *ListNode) Push(value interface{}) {

}

func (l *ListNode) Pop() (value interface{}, err error) {
	return nil, nil
}

func (l *ListNode) Peek() (value interface{}, err error) {
	return nil, nil
}

func (l *ListNode) Reverse() *ListNode {
	return nil
}

func (l *ListNode) Insert(n int) {

}
