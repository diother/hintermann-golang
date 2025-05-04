package builder

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
)

func RenderStaticPages(projectMetaList []ProjectMeta, tmpl *template.Template) error {
	pages, err := loadStaticPagesConfig()
	if err != nil {
		return fmt.Errorf("trouble reading static page config: %w", err)
	}
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

func loadStaticPagesConfig() ([]StaticPage, error) {
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
	return nil
}
