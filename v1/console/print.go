package console

import (
	"encoding/json"
	"fmt"
	"log"
)

func Log(emp interface{}) {
	empJSON, err := json.MarshalIndent(emp, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Printf(string(empJSON))
	fmt.Println()
}
