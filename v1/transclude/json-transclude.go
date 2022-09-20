package transclude

import (
	"encoding/json"
)

func ByJSON(origin interface{}, destine interface{}) {
	bData, err := json.Marshal(origin)
	if err != nil {
		return
	}

	err = json.Unmarshal(bData, destine)
	if err != nil {
		return
	}
}
