package brass

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mikerybka/english"
	"github.com/mikerybka/golang"
)

type App struct {
	Name     *english.Name
	BaseType *golang.Ident
}

func (a *App) genMetaPkg(path string) error {
	return nil
}

func (a *App) Build(tag string) error {
	// Create work dir
	workdir, err := os.MkdirTemp(os.TempDir(), "")
	if err != nil {
		return err
	}

	// Generate pkg/meta
	path := filepath.Join(workdir, "pkg/meta")
	err = a.genMetaPkg(path)
	if err != nil {
		return fmt.Errorf("generating pkg/meta: %s", err)
	}

	// Generate pkg/api/auth
	path = filepath.Join(workdir, "pkg/api/auth")
	err = a.genAPIAuthPkg(path)
	if err != nil {
		return fmt.Errorf("generating pkg/api/auth: %s", err)
	}

	// Generate pkg/api/data
	path = filepath.Join(workdir, "pkg/api/data")
	err = a.genAPIDataPkg(path)
	if err != nil {
		return fmt.Errorf("generating pkg/api/data: %s", err)
	}

	// Generate pkg/api
	path = filepath.Join(workdir, "pkg/api")
	err = a.genAPIPkg(path)
	if err != nil {
		return fmt.Errorf("generating pkg/api: %s", err)
	}

	// Generate pkg/www
	// TODO

	// Generate pkg/admin
	// TODO

	// Generate main.go

	// Generate Dockerfile
	// Run docker build

	// Delete workdir
}
