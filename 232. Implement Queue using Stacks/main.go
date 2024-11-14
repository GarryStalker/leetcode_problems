package main

type MyQueue struct {
	s []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.s = append(this.s, x)
}

func (this *MyQueue) Pop() int {
	result := this.s[0]
	this.s = this.s[1:]
	return result
}

func (this *MyQueue) Peek() int {
	return this.s[0]
}

func (this *MyQueue) Empty() bool {
	return len(this.s) == 0
}
