package main

import (
	"fmt"
)

type Time struct {
	unix uint64
}

type Interval struct {
	start Time
	end Time
}

type Event struct {
	interval Interval
}

func main() {
	start := Time{unix: 0}
	end := Time{unix: 1000}
	interval := Interval{start, end}
	event := Event{interval}
	fmt.Println(event)
}
