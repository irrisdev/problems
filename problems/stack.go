package problems

type value struct {
    val int
    min int
}

type MinStack struct {
    stack map[int]value
    size int
}


func Constructor() MinStack {
    return MinStack{
        stack: make(map[int]value),
        size: 0,
    }
}


func (this *MinStack) Push(val int)  {

    if this.size == 0 {
        this.stack[this.size] = value{
        min: val,
        val: val,
        }
    } else {
        localMin := this.stack[this.size-1].min
        if val < localMin {
            localMin = val
        }
        this.stack[this.size] = value{
            min: localMin,
            val: val,
        }    
    }
    
    this.size++
}


func (this *MinStack) Pop()  {

    delete(this.stack, this.size)
    this.size--
    
}


func (this *MinStack) Top() int {
    return this.stack[this.size].val
}


func (this *MinStack) GetMin() int {
    return this.stack[this.size].min
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */