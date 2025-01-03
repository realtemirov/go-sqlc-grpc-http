package middleware

import (
	"encoding/json"

	"github.com/realtemirov/go-sqlc-grpc-http/utils"
)

func convertToMap(data any) (map[string]any, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	var jsonMap map[string]any
	if errMarshaling := json.Unmarshal(jsonData, &jsonMap); errMarshaling != nil {
		return nil, errMarshaling
	}

	return jsonMap, nil
}

func redactSensitiveInfo(data map[string]any) {
	var sensitiveInfo = []string{
		"password", "email", "phone",
		"passport", "pin", "otp",
		"card_number", "cvv",
		"card_holder", "card_expiry_date",
		"card_token", "login",
	}

	for _, key := range sensitiveInfo {
		redactKey(data, key)
	}
}

func redactKey(data map[string]any, key string) {
	for k, v := range data {
		if k == key {
			data[k] = utils.SanitizeLogin(v.(string))
		}
		if nestedMap, ok := v.(map[string]any); ok {
			redactKey(nestedMap, key)
		}
	}
}
