package brass

import (
	"os"
	"path/filepath"

	"github.com/mikerybka/util"
)

type FS struct {
	RootDir string
}

func (fs *FS) WriteFile(path, data string) error {
	return os.WriteFile(filepath.Join(fs.RootDir, path), []byte(data), os.ModePerm)
}

func (fs *FS) ReadDir(path string) ([]string, error) {
	return util.ReadDir(filepath.Join(fs.RootDir, path))
}

func (fs *FS) ReadFile(path string) (string, error) {
	b, err := os.ReadFile(filepath.Join(fs.RootDir, path))
	if err != nil {
		return "", err
	}
	return string(b), nil
}
