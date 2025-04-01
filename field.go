package brass

import "github.com/mikerybka/english"

type Field struct {
	Name    *english.Name
	TypeID  string
	IsArray bool
	IsMap   bool
}
