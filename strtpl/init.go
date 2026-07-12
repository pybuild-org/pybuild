package strtpl

var vars map[string]any

var TemplateConfig = struct {
	StartTag string `prop:"start-tag"`
	EndTag   string `prop:"end-tag"`
}{
	StartTag: "{",
	EndTag:   "}",
}

func init() {
	vars = make(map[string]any)
}
