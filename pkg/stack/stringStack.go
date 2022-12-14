package stack

type StringStack struct {
	elements        []string
	currNumElements int
}

func New() StringStack {
	return StringStack{elements: make([]string, 0), currNumElements: 0}
}

func (stack *StringStack) PopList(n int) []string {
	lastElementList := stack.elements[stack.currNumElements - n:]
	stack.elements = stack.elements[:stack.currNumElements - n]
	stack.currNumElements -= n
	return lastElementList
}

func (stack *StringStack) PushList(strings []string){
	stack.elements = append(stack.elements, strings...)
	stack.currNumElements += len(strings)

}

func (stack *StringStack) Pop() string {
	lastElement := stack.elements[stack.currNumElements-1]
	stack.elements = stack.elements[:stack.currNumElements-1]
	stack.currNumElements -= 1
	return lastElement
}

func (stack *StringStack) Push(element string) {
	stack.elements = append(stack.elements, element)
	stack.currNumElements += 1
}

func (stack *StringStack) IsEmpty() bool {
	return stack.currNumElements == 0
}