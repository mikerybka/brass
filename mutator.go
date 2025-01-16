package brass

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func NewMutator[T any]() *Mutator[T] {
	return &Mutator[T]{}
}

type Mutator[T any] struct{}

func newID() string {
	return strconv.FormatInt(time.Now().UnixNano(), 36)
}

func (m *Mutator[T]) Run() {
	method := os.Args[1]
	path := os.Args[2]
	switch method {
	case "POST":
		id := newID()
		path := filepath.Join(path, id)
		var v T
		json.NewDecoder(os.Stdin).Decode(&v)
		b, _ := json.MarshalIndent(v, "", "  ")
		err := os.WriteFile(path, b, os.ModePerm)
		if err != nil {
			panic(err)
		}
	case "PUT":
		var v T
		json.NewDecoder(os.Stdin).Decode(&v)
		b, _ := json.MarshalIndent(v, "", "  ")
		err := os.WriteFile(path, b, os.ModePerm)
		if err != nil {
			panic(err)
		}
	case "PATCH":
		var v T
		b, _ := os.ReadFile(path)
		json.Unmarshal(b, &v)
		json.NewDecoder(os.Stdin).Decode(&v)
		b, _ = json.MarshalIndent(v, "", "  ")
		err := os.WriteFile(path, b, os.ModePerm)
		if err != nil {
			panic(err)
		}
	case "DELETE":
		os.Remove(path)
	default:
		panic("unknown method " + method)
	}
}
