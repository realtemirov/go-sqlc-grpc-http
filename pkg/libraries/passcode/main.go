package passcode

import "fmt"

func Passcode(login, code, authType string) string {
	return fmt.Sprintf("%s-%s-%s", login, authType, code)
}
