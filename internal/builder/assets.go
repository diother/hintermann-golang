package builder

import (
	"io/fs"
	"os"
	"path/filepath"
)

func CopyGlobalAssets() error {
	const (
		src  = "static"
		dest = "dist/static"
	)

	return filepath.WalkDir(src, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		targetPath := filepath.Join(dest, relPath)

		if d.IsDir() {
			return os.MkdirAll(targetPath, 0755)
		}

		input, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		return os.WriteFile(targetPath, input, 0644)
	})
}

func copyProjectAssets(slug string) error {
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
