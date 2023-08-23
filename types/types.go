package types

import (
	"fmt"
	"encoding/json"
)

type EWBNumber uint64

const DefaultDateFormat = "02/01/2006"

type AmbiguousString string

func (as *AmbiguousString) UnmarshalJSON(b []byte) error {
	var item any
	if err := json.Unmarshal(b, &item); err != nil {
		return err
	}
	*as = fmt.Sprintf("%v", item)
	return nil
}