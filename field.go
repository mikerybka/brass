package brass

import "github.com/mikerybka/english"

type Field struct {
	Name   *english.Name `json:"name"`
	TypeID string        `json:"typeID"`
}
