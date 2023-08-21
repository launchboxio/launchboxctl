package reporter

import (
	"encoding/json"
	"fmt"
)

type JsonReporter struct {
	Indent bool
}

func NewJsonReporter() JsonReporter {
	return JsonReporter{
		Indent: true,
	}
}

func (r JsonReporter) Output(in interface{}) error {
	var out []byte
	var err error
	if r.Indent {
		out, err = json.MarshalIndent(in, "", "  ")
	} else {
		out, err = json.Marshal(in)
	}
	if err != nil {
		return err
	}
	fmt.Print(string(out))
	return nil
}
