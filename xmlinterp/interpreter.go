package xmlinterp

func New(onTagOpen, onTagClose func()) *Interpreter {
	return &Interpreter{
		stack:      make([]*Node, 0),
		onTagOpen:  onTagOpen,
		onTagClose: onTagClose,
	}
}
