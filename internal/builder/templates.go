package builder

import (
	"html/template"

	"github.com/diother/hintermann-golang/internal/helpers"
)

func LoadTemplates() (tmpl *template.Template, err error) {
	tmpl = template.New("base").Funcs(template.FuncMap{
		"slice": helpers.SliceHelper,
		"props": helpers.PropsHelper,
		"merge": helpers.MergePropsHelper,
		"safe":  helpers.SafeHTML,
	})
	tmpl, err = tmpl.ParseGlob("internal/views/*.html")
	tmpl, err = tmpl.ParseGlob("internal/views/components/*.html")
	tmpl, err = tmpl.ParseGlob("internal/views/project/*.html")

	if err != nil {
		return nil, err
	}
	return tmpl, nil
}
