type MinStack struct {
	Stack []int
	Min_stack []int

}

func Constructor() MinStack {
	return MinStack {
		Stack: []int{},
		Min_stack: []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.Stack=append(this.Stack,val)
	n:=len(this.Min_stack)-1
	if n >=0 {
		if this.Min_stack[n] >= val {
			this.Min_stack=append(this.Min_stack,val)
		}
	} else {
		this.Min_stack=append(this.Min_stack,val)
	}

}

func (this *MinStack) Pop() {
	n:=len(this.Stack)-1
	n_min:=len(this.Min_stack)-1
	top:=this.Stack[n]
	this.Stack=this.Stack[:n] // we dont want the nth index
	if top == this.Min_stack[n_min] {
		this.Min_stack=this.Min_stack[:n_min]
	}
}

func (this *MinStack) Top() int {
	n:=len(this.Stack)-1
	return this.Stack[n]

}

func (this *MinStack) GetMin() int {
	n:=len(this.Min_stack)-1
	return this.Min_stack[n]

}
