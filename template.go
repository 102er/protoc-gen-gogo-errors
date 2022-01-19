package main

import (
	"bytes"
	"text/template"
)

var i18n = `
var (
	zhMap map[string]string
	enMap map[string]string
)
func init(){
	zhMap=make(map[string]string)
	{{ range .Errors }}
		zhMap[{{.Name}}_{{.Value}}.String()]="{{.MessageZh}}"
	{{- end }}
}
func init(){
	enMap=make(map[string]string)
	{{ range .Errors }}
		enMap[{{.Name}}_{{.Value}}.String()]="{{.MessageEn}}"
	{{- end }}
}
func getMessage(ctx context.Context,key string)string{
	l:="en"
	if ll,ok:=errors.FromErrorLangCtx(ctx);ok && ll =="zh"{
		l = "zh"
	}
	if l == "zh" {
		return zhMap[key]
	}
	return enMap[key]
}
`

var errorsTemplate = `
{{ range .Errors }}
func Is{{.CamelValue}}(ctx context.Context,err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	
	return e.Message == getMessage(ctx, {{.Name}}_{{.Value}}.String()) && e.Code == {{.HTTPCode}}
}
func Error{{.CamelValue}}(ctx context.Context,format string, args ...interface{}) *errors.Error {
	 return errors.NewWithErrCode({{.HTTPCode}}, {{.ErrCode}}, getMessage(ctx, {{.Name}}_{{.Value}}.String()), fmt.Sprintf(format, args...))
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
