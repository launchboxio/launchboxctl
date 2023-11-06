package printer

import (
	"encoding/json"
	"fmt"
)

func Print(in interface{}) error {
	bytes, err := json.MarshalIndent(in, "", "\t")
	if err != nil {
		return err
	}
	fmt.Println(string(bytes))
	return nil
}
