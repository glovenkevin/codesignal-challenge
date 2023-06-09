package main

// source: https://leetcode.com/problems/implement-stack-using-queues
func main() {
}

type MyStack struct {
	data []int
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.data = append(this.data, x)
}

func (this *MyStack) Pop() int {
	ret := this.data[len(this.data)-1]
	newData := make([]int, len(this.data)-1)
	copy(newData, this.data)
	this.data = newData
	return ret
}

func (this *MyStack) Top() int {
	return this.data[len(this.data)-1]
}

func (this *MyStack) Empty() bool {
	return len(this.data) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
