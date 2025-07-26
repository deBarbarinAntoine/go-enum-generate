package internal

import (
	"embed"
)

const enumTemplateFileName = "enum.go.tmpl"
const enumTemplateName = "enum"

//go:embed "enum.go.tmpl"
var enumTemplate embed.FS
