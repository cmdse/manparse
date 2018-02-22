package extractor

import (
	"fmt"
	"regexp"
)

// match patterns like [=optional-value]
var matchOptionalAssignmentValue = regexp.MustCompile(`^(\S+)\[=(\S+)\]$`)

// splitOptExpression tries to match an optional value assignment pattern
// it returns an array of those expressions and true if matched
// an array with the given flag if not matched
func splitOptExpression(expression string) ([]string, bool) {
	matches := matchOptionalAssignmentValue.FindStringSubmatch(expression)
	if len(matches) == 3 {
		expressions := make([]string, 2)
		expressions[0] = matches[1]
		expressions[1] = fmt.Sprintf("%v=%v", matches[1], matches[2])
		return expressions, true
	}
	return []string{expression}, false
}
