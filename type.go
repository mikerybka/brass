package brass

import (
	"github.com/mikerybka/english"
)

type Type struct {
	Name             *english.Name `json:"name"`
	PluralName       *english.Name `json:"plural_name"`
	IsArray          bool          `json:"is_array"`
	IsMap            bool          `json:"is_map"`
	UnderlyingTypeID string        `json:"underlying_type_id"`
	IsStruct         bool          `json:"is_struct"`
	Fields           []Field       `json:"fields"`
	HTMLTemplate     string        `json:"html_template"`
}
