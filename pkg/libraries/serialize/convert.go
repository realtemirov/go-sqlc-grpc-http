package serialize

import (
	"encoding/json"
)

func MarshalUnMarshal(input any, output any) error {
	rawBytes, err := json.Marshal(input)
	if err != nil {
		return err
	}

	return json.Unmarshal(rawBytes, output)
}
