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

var allDirections = map[Direction]struct{}{
	North: {}, East: {}, South: {}, West: {},
}

type Direction string

func (d *Direction) UnmarshalJSON(bytes []byte) error {
	var directionString string

	err := json.Unmarshal(bytes, &directionString)
	if err != nil {
		return err
	}

	direction := Direction(directionString)

	_, found := allDirections[direction]
	if !found {
		return fmt.Errorf("invalid direction: '%s'", direction)
	}

	*d = direction

	return nil
}

func (d Direction) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(d))
}
