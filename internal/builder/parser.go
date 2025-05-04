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
		if err := copyAssets(slug); err != nil {
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
	return nil, fmt.Errorf("no meta found")
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

	return tmpl.ExecuteTemplate(f, "project", project)
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
			output.WriteString(fmt.Sprintf("<h%d>%s</h%d>\n", level, content, level))
			continue
		}
		if imageMatch := imageRegex.FindStringSubmatch(line); imageMatch != nil {
			alt := imageMatch[1]
			src := imageMatch[2]
			output.WriteString(fmt.Sprintf(`<img src="%s" alt="%s"/>`+"\n", src, alt))
			continue
		}
		if linkMatch := linkRegex.FindStringSubmatch(line); linkMatch != nil {
			text := linkMatch[1]
			href := linkMatch[2]
			output.WriteString(fmt.Sprintf(`<a href="%s">%s</a>`+"\n", href, text))
			continue
		}
		output.WriteString(fmt.Sprintf("<p>%s</p>\n", line))
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}
	return output.String(), nil
}

func copyAssets(slug string) error {
	distDirPath := filepath.Join(distDir, "projects", slug)
	srcMediaPath := filepath.Join(projectDir, slug, "media")

	entries, err := os.ReadDir(srcMediaPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	for _, entry := range entries {
		srcPath := filepath.Join(srcMediaPath, entry.Name())
		distPath := filepath.Join(distDirPath, entry.Name())

		input, err := os.ReadFile(srcPath)
		if err != nil {
			return err
		}
		if err := os.WriteFile(distPath, input, 0644); err != nil {
			return err
		}
	}
	return nil
}
