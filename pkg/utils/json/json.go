package josn

import "encoding/json"

func StructToJson(v interface{}) (string, error) {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "", err
	}

	return string(data), nil
}
