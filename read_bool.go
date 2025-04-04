package brass

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func ReadBool(path ...string) (bool, error) {
	p := filepath.Join(path...)
	fi, err := os.Stat(p)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		}
		return false, err
	}
	if fi.IsDir() {
		return false, fmt.Errorf("%s is dir", p)
	}
	return true, nil
}
