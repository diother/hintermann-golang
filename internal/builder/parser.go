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
		srcSlug := entry.Name()
		var project Project

		project.Meta, err = loadMetadata(srcSlug)
		if err != nil {
			return fmt.Errorf("project %s metadata parsing failure: %w", srcSlug, err)
		}
		slug := project.Meta.Slug
		project.HTMLBody, err = parseMarkdown(srcSlug, slug)
		if err != nil {
			return fmt.Errorf("parsing project %s: %w", srcSlug, err)
		}
		if err := renderProject(project, slug, tmpl); err != nil {
			return fmt.Errorf("rendering project '%s': %w", srcSlug, err)
		}
		if err := copyProjectAssets(srcSlug, slug); err != nil {
			return fmt.Errorf("failed to copy assets for '%s': %w", srcSlug, err)
		}
	}
	return nil
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

func parseMarkdown(srcSlug, slug string) (string, error) {
	path := filepath.Join(projectDir, srcSlug, "body.md")
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

			imgSrc := fmt.Sprintf("/projects/%s/%s", slug, src)
			output.WriteString(fmt.Sprintf(`
				<img
					class="prose__img"
					src="%s-960.webp"
					srcset="%s-960.webp 960w, %s-1280.webp 1280w"
					sizes="(max-width: 1024px) 100vw, 1280px"
					alt="%s"
				>
				`+"\n", imgSrc, imgSrc, imgSrc, alt))
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
