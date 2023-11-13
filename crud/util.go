package crud

import "encoding/json"

func ToJson(v any) (string, error){
	b, err := json.MarshalIndent(v, "", "  ")
    if err != nil {
        return "", err
    }
	return string(b), nil
}