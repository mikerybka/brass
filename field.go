package brass

import "github.com/mikerybka/english"

type Field struct {
	Name    *english.Name `json:"name"`
	IsArray bool          `json:"is_array"`
	IsMap   bool          `json:"is_map"`
	TypeID  string        `json:"type_id"`
}
