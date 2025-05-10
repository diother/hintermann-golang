package builder

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"time"
)

func RenderStaticPages(projectMetaList []ProjectMeta, pages []StaticPage, tmpl *template.Template) error {
	for _, page := range pages {
		var buf bytes.Buffer

		if err := tmpl.ExecuteTemplate(&buf, page.Slug, projectMetaList); err != nil {
			return fmt.Errorf("rendering template %s: %w", page.Slug, err)
		}

		outputPath := filepath.Join(distDir, page.Path, "index.html")
		if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
			return fmt.Errorf("creating dir for %s: %w", outputPath, err)
		}

		if err := os.WriteFile(outputPath, buf.Bytes(), 0644); err != nil {
			return fmt.Errorf("writing file %s: %w", outputPath, err)
		}
	}
	return nil
}

func LoadStaticPageList() ([]StaticPage, error) {
	data, err := os.ReadFile("internal/views/static_pages.json")
	if err != nil {
		return nil, err
	}
	var pages []StaticPage
	if err = json.Unmarshal(data, &pages); err != nil {
		return nil, err
	}
	for _, page := range pages {
		if err := page.validate(); err != nil {
			return nil, err
		}
	}
	return pages, err
}

func (pm *StaticPage) validate() error {
	if pm.Slug == "" {
		return fmt.Errorf("missing slug")
	}
	if pm.Path == "" {
		return fmt.Errorf("missing path")
	}
	if pm.LastMod == "" {
		return fmt.Errorf("missing path")
	}
	if _, err := time.Parse("2006-01-02", pm.LastMod); err != nil {
		return fmt.Errorf("invalid date format (expected YYYY-MM-DD): %s", pm.LastMod)
	}
	return nil
}
