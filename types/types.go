package types

import (
	"encoding/json"
	"fmt"
)

type EWBNumber uint64

const DefaultDateFormat = "02/01/2006"

type AmbiguousString string

func (as *AmbiguousString) UnmarshalJSON(b []byte) error {
	var item any
	if err := json.Unmarshal(b, &item); err != nil {
		return err
	}
	*as = AmbiguousString(fmt.Sprintf("%v", item))
	return nil
}
