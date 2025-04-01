package brass

import (
	"github.com/mikerybka/english"
)

type Type struct {
	Name             *english.Name
	PluralName       *english.Name
	IsArray          bool
	IsMap            bool
	UnderlyingTypeID string
	IsStruct         bool
	Fields           []Field
}
