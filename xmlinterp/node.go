package xmlinterp

func (i *Interpreter) Node(idx int) *Node {
	size := i.StackSize()

	if idx >= 0 {
		if idx+1 > size {
			return nil
		}

		return i.stack[idx]
	}

	if idx < -size {
		return nil
	}

	return i.stack[size+idx]
}

func (i *Interpreter) CurrentNode() *Node {
	return i.Node(-1)
}

func (i *Interpreter) ParentNode() *Node {
	return i.Node(-2)
}

func (i *Interpreter) SaveChild() {
	pn := i.ParentNode()
	if pn != nil {
		pn.Childs = append(pn.Childs, i.CurrentNode())
	}
}
