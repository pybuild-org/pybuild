package xmlinterp

import (
	"encoding/xml"
	"io"
	"strings"
)

func (i *Interpreter) Run(r io.Reader) {
	decoder := xml.NewDecoder(r)

	for {
		token, err := decoder.Token()
		if err == io.EOF {
			break
		}

		switch t := token.(type) {

		case xml.StartElement:
			n := &Node{
				Name:  t.Name.Local,
				Attrs: make(map[string]string),
				Value: "",
			}

			for _, attr := range t.Attr {
				n.Attrs[attr.Name.Local] = strings.TrimSpace(attr.Value)
			}

			i.PushStack(n)
			i.onTagOpen()

		case xml.CharData:
			n := i.CurrentNode()
			if n != nil {
				n.Value += strings.TrimSpace(string(t))
			}

		case xml.EndElement:
			i.onTagClose()

		}
	}
}
