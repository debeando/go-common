package template

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/debeando/go-common/file"
)

func Render(p string, o interface{}) string {
	var b bytes.Buffer

	f := template.FuncMap{
		"join":      strings.Join,
		"separator": Separator,
	}

	t, err := template.New(file.Name(p)).Funcs(f).Parse(file.ReadAsString(p))
	if err != nil {
		return ""
	}

	if err = t.Execute(&b, o); err != nil {
		return ""
	}

	return b.String()
}
