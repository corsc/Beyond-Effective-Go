package _2_template_example

import (
	"bytes"
	"text/template"
	"time"
)

func SimpleTemplate() (string, error) {
	templateContent := "Good morning {{ .Name }}.\nThe current time is: {{ .Time }}"

	tmpl, err := template.New("example").Parse(templateContent)
	if err != nil {
		return "", err
	}

	data := map[string]interface{}{
		"Name": "Craig",
		"Time": time.Date(2000, 2, 1, 2, 3, 4, 5, time.UTC).Format(time.RFC3339),
	}

	destination := &bytes.Buffer{}
	err = tmpl.Execute(destination, data)
	if err != nil {
		return "", err
	}

	return destination.String(), nil
}
