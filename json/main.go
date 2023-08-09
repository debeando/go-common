package json

import (
	"encoding/json"
)

func Escape(i string) (string, error) {
	b, err := json.Marshal(i)
	if err != nil {
		return "", err
	}
	s := string(b)
	return s[1 : len(s)-1], nil
}

func StructToJSON(i any) (string, error) {
	e, err := json.Marshal(i)
	if err != nil {
		return "", err
	}
	return string(e), nil
}

func StringToStruct(data string, v any) error {
	return json.Unmarshal([]byte(data), &v)
}
