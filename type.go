package brass

import (
	"fmt"

	"github.com/mikerybka/english"
)

type Type struct {
	Name       english.Name
	PluralName english.Name
}

func (t *Type) mutatorCmd() string {
	return fmt.Sprintf(`package main

import (
	"github.com/mikerybka/brass"
	"types"
)

func main() {
	brass.NewMutator[types.%s]().Run()
}
`, t.Name.PascalCase())
}

func (t *Type) pkgFile(pkgName string) string {
	panic("todo")
	return fmt.Sprintf(`package %s`, pkgName)
}
