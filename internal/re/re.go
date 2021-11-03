package re

import (
	"github.com/celicoo/re/errors"
	"github.com/celicoo/re/internal/base"
	"github.com/celicoo/re/internal/collections"
)

func Compile(expression *collections.Slice[base.Character]) (r *RE, e error) {
	var (
		s1 = NewState(MatchState.Label)
		s2 = s1
		s0 = NewState(StartState.Label, s1)
	)
	for len(*expression) > 0 {
		c := expression.Shift()
		switch c {
		default:
			s1 = NewState(c, s0.States...)
			s0 = NewState(s0.Label, s1)
		case '(':
			var subexpression collections.Slice[base.Character]
			for len(*expression) > 0 {
				c = expression.Shift()
				if c == ')' {
					break
				}
				subexpression.Push(c)
			}
			if c == ')' {
				s1, e = Compile(&subexpression)
				s0.Push(s1.States...)
				continue
			}
			fallthrough
		case ')':
			e = errors.Syntax("unmatched closing parenthesis")
		case '|':
			s1, e = Compile(expression)
			s0.Push(s1.States...)
		case '+':
			s2 = s1.Pop()
			if s2 == nil {
				e = errors.Syntax("nothing to repeat")
			}
			s1.Push(s2)
			s1.Push(s1)
		case '*':
			s2 = s1.Pop()
			if s2 == nil {
				e = errors.Syntax("nothing to repeat")
			}
			s1.Push(s2)
			s0.Push(s2)
			s1.Push(s1)
		case '?':
			s2 = s1.Pop()
			if s2 == nil {
				e = errors.Syntax("nothing to repeat")
			}
			s1.Push(s2)
			s0.Push(s2)
		}
		if e != nil {
			break
		}
	}
	s1 = s0.Pop()
	if s1 != nil && s1.Label != MatchState.Label {
		s0.Push(s1)
	}
	return s0, e
}

type RE = State

// Match returns whether text contains any match of s.
func (s *RE) Match(text string) bool {
	return false
}
