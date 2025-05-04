package brass

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

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
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}
	errs := util.RunInParallel([]func() error{
		func() error { return a.generateGoMod(dir) },
		func() error { return a.generateTypes(dir) },
		func() error { return a.generateServer(dir) },
		func() error { return a.generateClient(dir) },
		func() error { return a.generateFrontend(dir) },
		func() error { return a.generateFavicon(dir) },
		func() error { return a.generateDockerfile(dir) },
	})
	for i, err := range errs {
		if err != nil {
			return fmt.Errorf("%d: %s", i, err)
		}
	}
	cmd := exec.Command("go", "fmt", "./...")
	cmd.Dir = dir
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("go fmt: %s: %s", err, out)
	}
	cmd = exec.Command("go", "mod", "tidy")
	cmd.Dir = dir
	out, err = cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("go mod tidy: %s: %s", err, out)
	}
	return nil
}

func (a *App) generateGoMod(dir string) error {
	os.Remove(filepath.Join(dir, "go.mod"))
	os.Remove(filepath.Join(dir, "go.sum"))
	cmd := exec.Command("go", "mod", "init", a.Repo)
	cmd.Dir = dir
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s: %s", err, out)
	}
	return nil
}

func (a *App) generateTypes(dir string) (err error) {
	for _, t := range a.Types {
		err = a.generateType(dir, t)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *App) packageName() string {
	return filepath.Base(a.Repo)
}
func (a *App) generateType(dir string, t Type) error {
	path := filepath.Join(dir, t.Name.SnakeCase()+".go")
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = fmt.Fprintf(f, "package %s\n\ntype %s ", a.packageName(), t.Name.PascalCase())
	if err != nil {
		return err
	}
	if t.IsArray {
		if isBuiltinType(t.UnderlyingTypeID) {
			_, err = fmt.Fprintf(f, "[]%s\n", t.UnderlyingTypeID)
			return err
		}
		underlyingType, ok := a.Types[t.UnderlyingTypeID]
		if !ok {
			return fmt.Errorf("no type %s", t.UnderlyingTypeID)
		}
		_, err = fmt.Fprintf(f, "[]%s\n", underlyingType.Name.PascalCase())
		return err
	}
	if t.IsMap {
		if isBuiltinType(t.UnderlyingTypeID) {
			_, err = fmt.Fprintf(f, "map[string]%s\n", t.UnderlyingTypeID)
			return err
		}
		underlyingType, ok := a.Types[t.UnderlyingTypeID]
		if !ok {
			return fmt.Errorf("no type %s", t.UnderlyingTypeID)
		}
		_, err = fmt.Fprintf(f, "map[string]%s\n", underlyingType.Name.PascalCase())
		return err
	}
	if t.IsStruct {
		_, err = fmt.Fprintf(f, "struct {\n")
		if err != nil {
			return err
		}
		for _, field := range t.Fields {
			_, err = fmt.Fprintf(f, "\t%s ", field.Name.PascalCase())
			if err != nil {
				return err
			}
			if field.IsArray {
				_, err = fmt.Fprintf(f, "[]")
				if err != nil {
					return err
				}
			} else if field.IsMap {
				_, err = fmt.Fprintf(f, "map[string]")
				if err != nil {
					return err
				}
			} else if !isBuiltinType(field.TypeID) {
				_, err = fmt.Fprintf(f, "*")
				if err != nil {
					return err
				}
			}
			if isBuiltinType(field.TypeID) {
				_, err = fmt.Fprintf(f, "%s", field.TypeID)
				if err != nil {
					return err
				}
			} else {
				typ := a.Types[field.TypeID]
				_, err = fmt.Fprintf(f, "%s", typ.Name.PascalCase())
				if err != nil {
					return err
				}
			}
			_, err = fmt.Fprintf(f, " `json:\"%s\"`\n", field.Name.SnakeCase())
			if err != nil {
				return err
			}
		}
		_, err = fmt.Fprintf(f, "}\n")
		if err != nil {
			return err
		}
		return err
	}
	panic("bad type: no kind")
}
func (a *App) generateServer(dir string) error     { return nil }
func (a *App) generateServerCmd(dir string) error  { return nil }
func (a *App) generateClient(dir string) error     { return nil }
func (a *App) generateFrontend(dir string) error   { return nil }
func (a *App) generateFavicon(dir string) error    { return nil }
func (a *App) generateDockerfile(dir string) error { return nil }
