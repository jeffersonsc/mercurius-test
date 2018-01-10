package template

import (
	"html/template"

	"github.com/jeffersonsc/mercurius-test/lib/context"
)

// FuncMaps to view
func FuncMaps() []template.FuncMap {
	return []template.FuncMap{
		map[string]interface{}{
			"Tr": context.I18n,
		}}
}
