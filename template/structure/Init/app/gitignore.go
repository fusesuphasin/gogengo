package app

import "html/template"

var GitignoreTemplate = template.Must(template.New("").Parse(
	`.env
`))
