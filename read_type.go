package brass

import "path/filepath"

func ReadType(path ...string) (*Type, error) {
	var err error
	t := &Type{}
	p := filepath.Join(path...)

	t.Name, err = ReadName(p, "name")
	if err != nil {
		return nil, err
	}

	t.PluralName, err = ReadName(p, "plural_name")
	if err != nil {
		return nil, err
	}

	t.IsArray, err = ReadBool(p, "is_array")
	if err != nil {
		return nil, err
	}

	t.IsMap, err = ReadBool(p, "is_map")
	if err != nil {
		return nil, err
	}

	t.UnderlyingTypeID, err = ReadString(p, "underlying_type_id")
	if err != nil {
		return nil, err
	}

	t.IsStruct, err = ReadBool(p, "is_struct")
	if err != nil {
		return nil, err
	}

	t.Fields, err = ReadArray(ReadField, p, "fields")
	if err != nil {
		return nil, err
	}

	t.HTMLTemplate, err = ReadString(p, "html_template")
	if err != nil {
		return nil, err
	}

	return t, nil
}
