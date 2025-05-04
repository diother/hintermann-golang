package fs

import "os"

func CleanDist() error {
	err := os.RemoveAll("dist")
	if err != nil {
		return err
	}
	return os.MkdirAll("dist", 0755)
}
