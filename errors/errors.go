package errors

import "errors"

var (
	// MissingClosingParenthesis error is returned when an invalid regular
	// expression is passed to re2postfix.Run
	MissingClosingParenthesis = errors.New("error parsing regex: missing closing )")
	// MissingOpeningParenthesis error is returned when an invalid regular
	// expression is passed to re2postfix.Run
	MissingOpeningParenthesis = errors.New("error parsing regex: missing opening (")
)
