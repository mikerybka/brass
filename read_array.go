package brass

import (
	"errors"
	"os"
	"strconv"
)

func ReadArray[T any](readElem func(path ...string) (T, error), path ...string) ([]T, error) {
	i := 1001
	res := []T{}
	for {
		v, err := readElem(append(path, strconv.Itoa(i))...)
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				return res, nil
			}
			return nil, err
		}
		res = append(res, v)
		i++
	}
}
