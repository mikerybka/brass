package brass

import (
	_ "embed"
	"encoding/json"
)

//go:embed default_types.json
var defaultTypesBytes []byte

var DefaultTypes = map[string]Type{}

func init() {
	err := json.Unmarshal(defaultTypesBytes, &DefaultTypes)
	if err != nil {
		panic(err)
	}
}
