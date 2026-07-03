package xmlinterp

import (
	"encoding/xml"
	"io"
	"log"
	"strings"
)

func (i *Interpreter) Run(r io.Reader) {
	decoder := xml.NewDecoder(r)

	for {
		token, err := decoder.Token()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalln(err)
		}

		switch t := token.(type) {

		case xml.StartElement:
			n := &Node{
				Name:  t.Name.Local,
				Attrs: make(map[string]string),
				Value: "",
			}

			for _, attr := range t.Attr {
				n.Attrs[attr.Name.Local] = attr.Value
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
