package brass

import (
	"github.com/mikerybka/english"
)

type Type struct {
	Name       english.Name
	PluralName english.Name
	Fields     []Field
}
