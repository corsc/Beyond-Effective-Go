package _1_example

import (
	"fmt"
	"time"
)

func Example() {
	carrot := struct {
		Variety         string
		WhenToPlant     []time.Month
		Spacing         int
		WeeksTilHarvest int
	}{
		Variety:         "Atomic Red",
		WhenToPlant:     []time.Month{time.January, time.February, time.November, time.December},
		Spacing:         10,
		WeeksTilHarvest: 18,
	}

	// Output: struct { Variety string; WhenToPlant []time.Month; Spacing int; WeeksTilHarvest int }{Variety:"Atomic Red", WhenToPlant:[]time.Month{1, 2, 11, 12}, Spacing:10, WeeksTilHarvest:18}
	fmt.Printf("%#v", carrot)
}
