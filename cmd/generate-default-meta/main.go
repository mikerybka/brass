package main

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/mikerybka/brass"
	"github.com/mikerybka/english"
	"github.com/mikerybka/util"
)

func genMeta() *brass.Metadata {
	structTypes := []brass.Type{
		{
			Name:       english.NewName("Root"),
			PluralName: english.NewName("Roots"),
			IsStruct:   true,
			Fields: []brass.Field{
				{
					Name:   english.NewName("Title"),
					TypeID: "string",
				},
			},
		},
	}
	var meta = &brass.Metadata{
		Types:    map[string]*brass.Type{},
		RootType: "root",
	}
	for _, t := range structTypes {
		id := t.Name.SnakeCase()
		meta.Types[id] = &t

		arrayType := &brass.Type{
			Name:             english.NewName(t.Name.String() + " Array"),
			PluralName:       english.NewName(t.Name.String() + " Arrays"),
			IsArray:          true,
			UnderlyingTypeID: id,
		}
		id = arrayType.Name.SnakeCase()
		meta.Types[id] = arrayType

		mapType := &brass.Type{
			Name:             english.NewName(t.Name.String() + " Map"),
			PluralName:       english.NewName(t.Name.String() + " Maps"),
			IsMap:            true,
			UnderlyingTypeID: id,
		}
		id = mapType.Name.SnakeCase()
		meta.Types[id] = mapType
	}
	return meta
}

func main() {
	path := filepath.Join(util.RequireEnvVar("SRC_DIR"), "github.com/mikerybka/brass/default_meta.json")
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	err = json.NewEncoder(f).Encode(genMeta())
	if err != nil {
		panic(err)
	}
	err = f.Close()
	if err != nil {
		panic(err)
	}
}
