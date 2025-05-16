package builder

import (
	"html/template"

	"github.com/diother/hintermann-golang/internal/helpers"
)

func LoadTemplates() (tmpl *template.Template, err error) {
	tmpl = template.New("base").Funcs(template.FuncMap{
		"slice":      helpers.SliceHelper,
		"props":      helpers.PropsHelper,
		"merge":      helpers.MergePropsHelper,
		"safe":       helpers.SafeHTMLHelper,
		"attr":       helpers.AttrHelper,
		"add":        helpers.AddHelper,
		"mul":        helpers.MulHelper,
		"mulFloat":   helpers.MulFloatHelper,
		"formatDate": helpers.FormatDateHelper,
	})
	tmpl, err = tmpl.ParseGlob("internal/views/*.html")
	if err != nil {
		return nil, err
	}
	tmpl, err = tmpl.ParseGlob("internal/views/components/*.html")
	if err != nil {
		return nil, err
	}
	return tmpl, nil
}
