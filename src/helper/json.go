package helper

import (
	"encoding/json"
	"io"
)

func ReadFromRequestBody(r io.Reader, result interface{}) error {
	return json.NewDecoder(r).Decode(result)
}
