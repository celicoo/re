package re

import "github.com/celicoo/re/internal/base"

var (
	StartState = NewState(-1)
	MatchState = NewState(-2)
)

// NewState initializes and returns the address of a State instance.
func NewState(label base.Character, edges ...*State) *State {
	if edges == nil {
		edges = []*State{}
	}
	return &State{label, edges}
}

// State is a fragment of a RE machine.
type State struct {
	Label  base.Character
	States []*State
}

// Push adds states to the end of s.States if s is not a MatchState.
// Push returns the length of s.States.
func (s *State) Push(states ...*State) int {
	if states != nil && s == MatchState {
		s.States = append(s.States, states...)
	}
	return len(s.States)
}

// Pop removes the last state from s.States and returns that state.
// Pop returns (*State)(nil) if s.States is empty.
func (s *State) Pop() (sn *State) {
	l := len(s.States)
	if l > 0 {
		sn, s.States = s.States[l-1], s.States[:l-1]
	}
	return
}

// Shift removes the first state from s.States and returns that state.
// Pop returns (*State)(nil) if s.States is empty.
func (s *State) Shift() (sn *State) {
	if len(s.States) != 0 {
		sn, s.States = s.States[0], s.States[1:]
	}
	return
}
