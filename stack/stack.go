package stack

type Stack struct {
	s []interface{}
}

func NewStack() *Stack {
	return &Stack{make([]interface{}, 0)}
}

func NewStackWithCapacity(capacity int) *Stack {
	return &Stack{make([]interface{}, 0, capacity)}
}

func (stack *Stack) Len() int {
	return len(stack.s)
}

func (stack *Stack) Empty() bool {
	return len(stack.s) == 0
}

func (stack *Stack) Push(item interface{}) {
	stack.s = append(stack.s, item)
}

func (stack *Stack) Pop() interface{} {
	item := stack.Peek()
	if item != nil {
		stack.s = stack.s[:len(stack.s)-1]
	}
	return item
}

func (stack *Stack) Peek() interface{} {
	if len(stack.s) > 0 {
		return stack.s[len(stack.s)-1]
	}
	return nil
}
