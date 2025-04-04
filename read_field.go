package brass

import "path/filepath"

func ReadField(path ...string) (*Field, error) {
	var err error
	f := &Field{}
	p := filepath.Join(path...)

	f.Name, err = ReadName(p, "name")
	if err != nil {
		return nil, err
	}

	f.TypeID, err = ReadString(p, "type_id")
	if err != nil {
		return nil, err
	}

	return f, nil
}
