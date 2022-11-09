package b64

import "encoding/base64"

func Decode(text string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(text)
}
