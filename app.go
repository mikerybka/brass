package brass

import (
	"fmt"

	"github.com/mikerybka/util"
)

type App struct {
	Repo     string
	Name     string
	Icon     []byte // 1024x1024 pixel .png file
	CoreType string
	Types    map[string]Type
}

func (a *App) GenerateSourceCode(dir string) error {
	errs := util.RunInParallel([]func() error{})
	for i, err := range errs {
		if err != nil {
			return fmt.Errorf("%d: %s", i, err)
		}
	}
	return nil
}
