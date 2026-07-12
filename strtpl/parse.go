package strtpl

import "github.com/valyala/fasttemplate"

func Parse(target string) string {
	t := fasttemplate.New(
		target,
		TemplateConfig.StartTag,
		TemplateConfig.EndTag,
	)

	return t.ExecuteString(vars)
}
