package _3_generator_v7

import (
	"bytes"
	"fmt"
	"text/template"
)

type StubGenerator struct{}

func (s *StubGenerator) Generate(filename, targetInterface string) (string, error) {
	interfaces, err := s.parseSource(filename)
	if err != nil {
		return "", err
	}

	for _, thisInterface := range interfaces {
		if thisInterface.Name == targetInterface {
			return s.generate(thisInterface)
		}
	}

	return "", fmt.Errorf("supplied file does not include the interface %s", targetInterface)
}

func (s *StubGenerator) generate(thisInterface *Interface) (string, error) {
	output := &bytes.Buffer{}

	data := &templateData{
		Interface: thisInterface,
	}

	tmpl := template.New("generator")

	tmpl.Funcs(template.FuncMap{
		"isNotLast": isNotLast,
	})
	_, err := tmpl.Parse(stubTemplate)
	if err != nil {
		return "", err
	}

	err = tmpl.Execute(output, data)
	if err != nil {
		return "", err
	}

	return output.String(), nil
}

func isNotLast(len int, index int, insert string) string {
	if (len - 1) == index {
		return ""
	}
	return insert
}

type templateData struct {
	Interface *Interface
}

var stubTemplate = `
package {{ .Interface.PackageName }}

type Stub{{ .Interface.Name }} struct {}

{{ range .Interface.Methods -}}
func (s *Stub{{ $.Interface.Name }}) {{ .Name }}(
	{{- $len := len .Inputs }}
	{{- range $index, $param := .Inputs }}
		{{- $param.Name }} {{ $param.Type }} {{- isNotLast $len $index ", " }}
	{{- end -}}
) (*user.User, error) { 
	return &user.User{}, nil
}
{{ end -}}
`
