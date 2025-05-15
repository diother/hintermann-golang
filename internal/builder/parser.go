package builder

import (
	"bufio"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func RenderProjects(projectMetaList []ProjectMeta, tmpl *template.Template) error {
	entries, err := os.ReadDir(projectDir)
	if err != nil {
		return fmt.Errorf("reading project directory: %w", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		slug := entry.Name()
		var project Project

		project.Meta, err = loadProjectMetaBySlug(projectMetaList, slug)
		if err != nil {
			return fmt.Errorf("loading metadata for project %s: %w", slug, err)
		}
		project.HTMLBody, err = parseMarkdown(slug)
		if err != nil {
			return fmt.Errorf("parsing project %s: %w", slug, err)
		}
		if err := renderProject(project, slug, tmpl); err != nil {
			return fmt.Errorf("rendering project '%s': %w", slug, err)
		}
		if err := copyProjectAssets(slug); err != nil {
			return fmt.Errorf("failed to copy assets for '%s': %w", slug, err)
		}
	}
	return nil
}

func loadProjectMetaBySlug(projectMetaList []ProjectMeta, slug string) (*ProjectMeta, error) {
	for i := range projectMetaList {
		if projectMetaList[i].Slug == slug {
			return &projectMetaList[i], nil
		}
	}
	return nil, fmt.Errorf("no entry in the array found")
}

func renderProject(project Project, slug string, tmpl *template.Template) error {
	distPath := filepath.Join(distDir, "projects", slug, "index.html")

	if err := os.MkdirAll(filepath.Dir(distPath), 0755); err != nil {
		return fmt.Errorf("creating dist dir: %w", err)
	}
	f, err := os.Create(distPath)
	if err != nil {
		return err
	}
	defer f.Close()

	return tmpl.ExecuteTemplate(f, "project-single", project)
}

func parseMarkdown(slug string) (string, error) {
	path := filepath.Join(projectDir, slug, "body.md")
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var output strings.Builder
	scanner := bufio.NewScanner(file)

	headingRegex := regexp.MustCompile(`^(#{1,3}) (.+)$`)
	imageRegex := regexp.MustCompile(`^!\[(.*?)\]\((.*?)\)$`)
	linkRegex := regexp.MustCompile(`^\[(.*?)\]\((.*?)\)$`)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		if headingMatch := headingRegex.FindStringSubmatch(line); headingMatch != nil {
			level := len(headingMatch[1])
			content := headingMatch[2]
			output.WriteString(fmt.Sprintf(`<h%d class="prose__heading">%s</h%d>`+"\n", level, content, level))
			continue
		}
		if imageMatch := imageRegex.FindStringSubmatch(line); imageMatch != nil {
			alt := imageMatch[1]
			src := imageMatch[2]
			output.WriteString(fmt.Sprintf(`<img class="prose__img" src="/projects/%s/%s-1280.webp" alt="%s"/>`+"\n", slug, src, alt))
			continue
		}
		if linkMatch := linkRegex.FindStringSubmatch(line); linkMatch != nil {
			text := linkMatch[1]
			href := linkMatch[2]
			output.WriteString(fmt.Sprintf(`<a href="%s" class="prose__link">%s</a>`+"\n", href, text))
			continue
		}
		output.WriteString(fmt.Sprintf(`<p class="prose__text">%s</p>`+"\n", line))
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}
	return output.String(), nil
}
