package brass

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func ReadString(path ...string) (string, error) {
	b, err := os.ReadFile(filepath.Join(path...))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return "", nil
		}
		return "", err
	}
	return strings.TrimSpace(string(b)), nil
}
