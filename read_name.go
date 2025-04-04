package brass

import "github.com/mikerybka/english"

func ReadName(path ...string) (*english.Name, error) {
	s, err := ReadString(path...)
	if err != nil {
		return nil, err
	}
	return english.NewName(s), nil
}
