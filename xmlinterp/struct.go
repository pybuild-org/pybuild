package xmlinterp

type Interpreter struct {
	stack      []*Node
	onTagOpen  func()
	onTagClose func()
}

type Node struct {
	Name   string
	Attrs  map[string]string
	Value  string
	Childs []*Node
}
