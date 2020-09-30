package _2_template

import (
	"bytes"
	"html/template"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTemplate(t *testing.T) {
	welcomeSource := `
<html>
	<body>
		Hello {{ .Name }}.<br />
		It has been {{ .Days }} since your last login.<br />
		You have {{ .Messages }} unread messages.<br />
	</body>
</html>
`
	results := &bytes.Buffer{}

	data := struct {
		Name     string
		Days     int
		Messages int
	}{
		Name:     "John",
		Days:     7,
		Messages: 2,
	}

	welcomeTemplate, err := template.New("welcome").Parse(welcomeSource)
	require.NoError(t, err)

	err = welcomeTemplate.Execute(results, data)
	require.NoError(t, err)

	assert.Equal(t, expected, results.String())
}

var expected = `
<html>
	<body>
		Hello John.<br />
		It has been 7 since your last login.<br />
		You have 2 unread messages.<br />
	</body>
</html>
`
