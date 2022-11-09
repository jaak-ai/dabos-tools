package b64

import "encoding/base64"

func Encode(text []byte) string {
	return base64.StdEncoding.EncodeToString(text)
}
