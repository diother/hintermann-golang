package builder

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"time"
)

const (
	projectDir = "projects"
	distDir    = "dist"
)

func LoadProjectMetaList() (projectMetaList []ProjectMeta, err error) {
	entries, err := os.ReadDir(projectDir)
	if err != nil {
		return nil, fmt.Errorf("reading project directory: %w", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		srcSlug := entry.Name()
		metadata, err := loadMetadata(srcSlug)
		if err != nil {
			return nil, fmt.Errorf("project %s metadata parsing failure: %w", srcSlug, err)
		}
		projectMetaList = append(projectMetaList, *metadata)
	}

	sortMetadataSliceByDate(projectMetaList)

	return projectMetaList, nil
}

func loadMetadata(srcSlug string) (*ProjectMeta, error) {
	var meta ProjectMeta
	path := filepath.Join(projectDir, srcSlug, "meta.json")
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(file, &meta); err != nil {
		return nil, err
	}
	if err := meta.validate(); err != nil {
		return nil, err
	}
	return &meta, nil
}

func (pm *ProjectMeta) validate() error {
	if pm.Title == "" {
		return fmt.Errorf("missing title")
	}
	if pm.Slug == "" {
		return fmt.Errorf("missing slug")
	}
	if slugRegex := regexp.MustCompile(`^[a-z0-9]+(-[a-z0-9]+)*$`); !slugRegex.MatchString(pm.Slug) {
		return fmt.Errorf("invalid slug format: %s", pm.Slug)
	}
	if pm.Date == "" {
		return fmt.Errorf("missing date")
	}
	if _, err := time.Parse("2006-01-02", pm.Date); err != nil {
		return fmt.Errorf("invalid date format (expected YYYY-MM-DD): %s", pm.Date)
	}
	if pm.Description == "" {
		return fmt.Errorf("missing description")
	}
	if pm.Read == "" {
		return fmt.Errorf("missing read")
	}
	if len(pm.Sponsors) == 0 {
		return fmt.Errorf("missing sponsors")
	}
	return nil
}

func sortMetadataSliceByDate(projects []ProjectMeta) {
	sort.Slice(projects, func(i, j int) bool {
		ti, _ := time.Parse("2006-01-02", projects[i].Date)
		tj, _ := time.Parse("2006-01-02", projects[j].Date)
		return tj.Before(ti)
	})
}
