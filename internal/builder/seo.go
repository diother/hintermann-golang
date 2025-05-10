package builder

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

func RenderSeoFiles(projectMetaList []ProjectMeta, staticPageList []StaticPage) error {
	tmpl, err := loadTextTemplates()
	if err != nil {
		return fmt.Errorf("text templates weren't parsed: %w", err)
	}
	if err := renderSitemap(projectMetaList, staticPageList, tmpl); err != nil {
		return fmt.Errorf("sitemap render failed: %w", err)
	}
	if err := renderRobots(tmpl); err != nil {
		return fmt.Errorf("robots.txt render failed: %w", err)
	}
	return nil
}

func loadTextTemplates() (*template.Template, error) {
	tmpl := template.New("base")
	tmpl, err := tmpl.ParseGlob("internal/views/*.xml")
	if err != nil {
		return nil, fmt.Errorf("xml templates failed: %w", err)
	}
	tmpl, err = tmpl.ParseGlob("internal/views/*.txt")
	if err != nil {
		return nil, fmt.Errorf("txt templates failed: %w", err)
	}
	return tmpl, nil
}

func renderSitemap(projectMetaList []ProjectMeta, staticPageList []StaticPage, tmpl *template.Template) error {
	var sitemap []SitemapEntry
	for _, page := range staticPageList {
		sitemap = append(sitemap, SitemapEntry{
			Loc:        "https://hintermann.ro" + page.Path,
			LastMod:    page.LastMod,
			ChangeFreq: "monthly",
			Priority:   "0.8",
		})
	}
	for _, project := range projectMetaList {
		sitemap = append(sitemap, SitemapEntry{
			Loc:        "https://hintermann.ro/projects/" + project.Slug,
			LastMod:    project.Date,
			ChangeFreq: "monthly",
			Priority:   "0.6",
		})
	}
	distPath := filepath.Join(distDir, "sitemap.xml")
	f, err := os.Create(distPath)
	if err != nil {
		return err
	}
	defer f.Close()
	return tmpl.ExecuteTemplate(f, "sitemap", sitemap)
}

func renderRobots(tmpl *template.Template) error {
	distPath := filepath.Join(distDir, "robots.txt")
	f, err := os.Create(distPath)
	if err != nil {
		return err
	}
	defer f.Close()
	return tmpl.ExecuteTemplate(f, "robots", nil)
}
