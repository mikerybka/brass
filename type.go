package brass

import (
	"github.com/mikerybka/english"
)

type Type struct {
	Name             *english.Name `json:"name"`
	PluralName       *english.Name `json:"pluralName"`
	IsArray          bool          `json:"isArray"`
	IsMap            bool          `json:"isMap"`
	UnderlyingTypeID string        `json:"underlyingTypeID"`
	IsStruct         bool          `json:"isStruct"`
	Fields           []Field       `json:"fields"`
}
