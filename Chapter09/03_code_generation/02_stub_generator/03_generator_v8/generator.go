package _3_generator_v8

import (
	"bytes"
	"fmt"
	"strings"
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
		"stubValue": stubValue,
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

func stubValue(typeName string) string {
	switch {
	case strings.HasPrefix(typeName, "int"):
		return "0"

	case strings.HasPrefix(typeName, "float"):
		return "0"

	case typeName == "bool":
		return "false"

	case typeName == "string":
		return `""`

	case typeName == "error":
		return "nil"

	case strings.HasPrefix(typeName, "*"):
		return "&" + strings.TrimPrefix(typeName, "*") + "{}"

	default:
		return typeName + "{}"
	}
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
) (
	{{- $len := len .Inputs }}
	{{- range $index, $result := .Outputs }}
		{{- $result.Name }} {{- $result.Type }} {{- isNotLast $len $index ", " }}
	{{- end -}}
) { 
	return {{ $len := len .Inputs }}
	{{- range $index, $result := .Outputs }}
		{{- stubValue $result.Type }} {{- isNotLast $len $index ", " }}
	{{- end }}
}
{{ end -}}
`
