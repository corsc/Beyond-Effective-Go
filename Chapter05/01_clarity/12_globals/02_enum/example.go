package _2_enum

import (
	"encoding/json"
	"fmt"
)

const (
	North Direction = "north"
	East  Direction = "east"
	South Direction = "south"
	West  Direction = "west"
)

var all = map[Direction]struct{}{
	North: {}, East: {}, South: {}, West: {},
}

type Direction string

func (d *Direction) UnmarshalJSON(bytes []byte) error {
	var asString string

	err := json.Unmarshal(bytes, &asString)
	if err != nil {
		return err
	}

	val := Direction(asString)

	_, found := all[val]
	if !found {
		return fmt.Errorf("expected value '%s'", val)
	}

	*d = val

	return nil
}

func (d Direction) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(d))
}
