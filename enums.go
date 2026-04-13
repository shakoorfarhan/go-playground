package main

import "fmt"

type ServerState int

const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateDisconnected
	StateRetrying
)

var stateNames = map[ServerState]string{
	StateIdle:         "Idle",
	StateConnected:    "Connected",
	StateDisconnected: "Disconnected",
	StateRetrying:     "Retrying",
}

func (s ServerState) String() string {
	return stateNames[s]
}
func main() {
	ns := transition(StateIdle)
	fmt.Println("Current state:", ns)

	ns2 := transition(ns)
	fmt.Println("Current state:", ns2)
}

func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected:
		return StateDisconnected
	case StateDisconnected:
		return StateRetrying
	case StateRetrying:
		return StateIdle
	default:
		panic("Unknown state")
	}
}
