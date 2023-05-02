package _3_generator_v3

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

type templateData struct {
	Interface *Interface
}

var stubTemplate = `
package {{ .Interface.PackageName }}

type Stub{{ .Interface.Name }} struct {}

func (s *Stub{{ .Interface.Name }}) LoadByID(ctx context.Context, userID int64) (*user.User, error) { 
	return &user.User{}, nil
}
`
