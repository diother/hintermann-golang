package helpers

import (
	"fmt"
	"html/template"
	"maps"
	"time"
)

func SliceHelper(args ...any) []any {
	return args
}

func AttrHelper(s string) template.HTMLAttr {
	return template.HTMLAttr(s)
}

func PropsHelper(values ...any) (map[string]any, error) {
	if len(values)%2 != 0 {
		return nil, fmt.Errorf("invalid props call: uneven number of args")
	}
	props := make(map[string]any, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)
		if !ok {
			return nil, fmt.Errorf("props keys must be strings")
		}
		props[key] = values[i+1]
	}
	return props, nil
}

func MergePropsHelper(props map[string]any, defaults map[string]any) map[string]any {
	merged := make(map[string]any, len(defaults))
	maps.Copy(merged, defaults)
	maps.Copy(merged, props)
	return merged
}

func SafeHTMLHelper(s string) template.HTML {
	return template.HTML(s)
}

func AddHelper(a, b int) int {
	return a + b
}

func MulHelper(a, b int) int {
	return a * b
}

func MulFloatHelper(a, b float64) float64 {
	return a * b
}

func FormatDateHelper(dateStr string) string {
	t, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return dateStr
	}
	return t.Format("2 Jan 2006")
}
