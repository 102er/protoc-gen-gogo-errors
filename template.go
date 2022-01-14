package main

import (
	"bytes"
	"text/template"
)

var i18n = `
const (
{{ range .Errors }}
	{{.Name}}_{{.Value}}.String()_en="{{.MessageEn}}"
	{{.Name}}_{{.Value}}.String()_zh="{{.MessageZn}}"
{{- end }}
)
`

var errorsTemplate = `
{{ range .Errors }}
func Is{{.CamelValue}}(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == {{.Name}}_{{.Value}}.String() && e.Code == {{.HTTPCode}} 

}
func Error{{.CamelValue}}(format string, args ...interface{}) *errors.Error {
	 return errors.New({{.HTTPCode}}, {{.Name}}_{{.Value}}.String(), fmt.Sprintf(format, args...))
}
{{- end }}
`

type errorInfo struct {
	Name       string
	Value      string
	HTTPCode   int
	CamelValue string
	ErrCode    int
	MessageEn  string
	MessageZh  string
}

type errorWrapper struct {
	Errors []*errorInfo
}

func (e *errorWrapper) i18n() string {
	buf := new(bytes.Buffer)
	tmpl, err := template.New("errors").Parse(i18n)
	if err != nil {
		panic(err)
	}
	if err := tmpl.Execute(buf, e); err != nil {
		panic(err)
	}
	return buf.String()
}

func (e *errorWrapper) execute() string {
	buf := new(bytes.Buffer)
	tmpl, err := template.New("errors").Parse(errorsTemplate)
	if err != nil {
		panic(err)
	}
	if err := tmpl.Execute(buf, e); err != nil {
		panic(err)
	}
	return buf.String()
}
