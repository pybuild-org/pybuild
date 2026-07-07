package strtpl

import "github.com/valyala/fasttemplate"

func Parse(target string) string {
	t := fasttemplate.New(target, "{", "}")
	return t.ExecuteString(vars)
}
