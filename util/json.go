package util

import (
	"encoding/json"
	"fmt"

	"github.com/wildanie12/region-api/util/color"
)

func DebugJSON(obj any, placeholder ...string) {
	if len(placeholder) > 0 {
		fmt.Println("------------------------------------------")
		fmt.Println(" ", placeholder[0])
		fmt.Println("------------------------------------------")
	}
	jsonD, _ := json.MarshalIndent(obj, "", "  ")
	fmt.Println(color.This(color.CYAN, string(jsonD)))
	if len(placeholder) > 0 {
		fmt.Println("------------------------------------------")
	}
}
