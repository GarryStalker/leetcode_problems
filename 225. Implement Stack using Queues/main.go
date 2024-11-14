package main

type MyStack struct {
	q []int
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.q = append(this.q, x)
}

func (this *MyStack) Pop() int {
	result := this.q[len(this.q)-1]
	this.q = this.q[:len(this.q)-1]
	return result
}

func (this *MyStack) Top() int {
	return this.q[len(this.q)-1]
}

func (this *MyStack) Empty() bool {
	return len(this.q) == 0
}
