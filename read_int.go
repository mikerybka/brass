package brass

import (
	"errors"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func ReadInt(path ...string) (int, error) {
	b, err := os.ReadFile(filepath.Join(path...))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return 0, nil
		}
		return 0, err
	}
	i, err := strconv.Atoi(strings.TrimSpace(string(b)))
	if err != nil {
		return 0, err
	}
	return i, nil
}
