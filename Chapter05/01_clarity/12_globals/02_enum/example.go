package _2_enum

import (
	"encoding/json"
	"fmt"
)

const (
	NORTH Direction = "north"
	EAST  Direction = "east"
	SOUTH Direction = "south"
	WEST  Direction = "west"
)

var all = map[Direction]struct{}{
	NORTH: {}, EAST: {}, SOUTH: {}, WEST: {},
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
		return fmt.Errorf("failed to convert '%s'", bytes)
	}

	*d = val

	return nil
}

func (d Direction) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(d))
}
