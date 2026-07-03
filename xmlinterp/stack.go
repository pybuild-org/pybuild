package xmlinterp

import "log"

func (i *Interpreter) StackSize() int {
	return len(i.stack)
}

func (i *Interpreter) PushStack(n *Node) {
	if i.StackSize() > 255 {
		log.Fatalln("stack overflow")
	}

	i.stack = append(i.stack, n)
}

func (i *Interpreter) PopStack() {
	size := i.StackSize()
	if size < 1 {
		return
	}

	i.stack = i.stack[:size-1]
}
