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

func (a *App) genAPIAuthPkg(path string) error {
	return nil
}

func (a *App) genAPIDataPkg(path string) error {
	return nil
}

func (a *App) genAPIPkg(path string) error {
	return nil
}

func (a *App) genDockerfile(path string) error {
	return nil
}

func (a *App) genMainGo(path string) error {
	return nil
}

func (a *App) Generate(dir string) error {
	// Create work dir
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}

	// Generate pkg/meta
	path := filepath.Join(dir, "pkg/meta")
	err = a.genMetaPkg(path)
	if err != nil {
		return fmt.Errorf("generating pkg/meta: %s", err)
	}

	// Generate pkg/api/auth
	path = filepath.Join(dir, "pkg/api/auth")
	err = a.genAPIAuthPkg(path)
	if err != nil {
		return fmt.Errorf("generating pkg/api/auth: %s", err)
	}

	// Generate pkg/api/data
	path = filepath.Join(dir, "pkg/api/data")
	err = a.genAPIDataPkg(path)
	if err != nil {
		return fmt.Errorf("generating pkg/api/data: %s", err)
	}

	// Generate pkg/api
	path = filepath.Join(dir, "pkg/api")
	err = a.genAPIPkg(path)
	if err != nil {
		return fmt.Errorf("generating pkg/api: %s", err)
	}

	// Generate pkg/www
	// TODO

	// Generate pkg/admin
	// TODO

	// Generate main.go
	path = filepath.Join(dir, "main.go")
	err = a.genMainGo(path)
	if err != nil {
		return fmt.Errorf("generating pkg/api: %s", err)
	}

	// Generate Dockerfile
	path = filepath.Join(dir, "Dockerfile")
	err = a.genDockerfile(path)
	if err != nil {
		return fmt.Errorf("generating pkg/api: %s", err)
	}

	return nil
}
