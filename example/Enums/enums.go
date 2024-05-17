package main

import "fmt"

// An enum is a type that has a fixed number of possible values,
// each with a distinct name.
// Go doesn’t have an enum type as a distinct language feature,
// but enums are simple to implement using existing language idioms.

type ServerState int

// The possible values for ServerState are defined as constants.
// The special keyword iota generates successive constant values automatically;
// in this case 0, 1, 2 and so on.
const (
	StateIdle = iota
	StateConnected
	StateError
	StateRetrying
)

var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

func (ss ServerState) String() string {
	return stateName[ss]
}

func main() {
	// If we have a value of type int, we cannot pass it to transition - the compiler will complain about type mismatch.
	// This provides some degree of compile-time type safety for enums.
	ns := transition(StateIdle)
	fmt.Println(ns)

	ns2 := transition(StateConnected)
	fmt.Println(ns2)
}

func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unwknown state: %s", s))
	}
	return StateConnected
}
