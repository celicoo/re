package re

import (
	"github.com/celicoo/re/internal/base"
	"github.com/celicoo/re/internal/collections"
	"github.com/celicoo/re/internal/re"
)

// Compile returns a type that implements RE and nil if expression can be
// parsed.
// Compile returns nil and an error if expression cannot be parsed.
func Compile(expression string) (RE, error) {{
	var (
		expression = collections.Slice[base.Character](expression)
		r, e       = re.Compile(&expression)
	)
	if e != nil {
		return nil, e
	}
	return r, nil
}}

// RE represents a compiled regular expression that can be used to match
// against text.
type RE interface {
	// Match returns whether text contains any match of RE.
	Match(text string) bool
}
