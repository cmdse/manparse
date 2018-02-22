package normalize

import (
	"fmt"
	"regexp"
)

// match patterns like [=optional-value]
var matchOptionalAssignmentValue = regexp.MustCompile(`^(\S+)\[=(\S+)\]$`)

// normalizeFlag tries to match an optional value assignment pattern
// it returns an array of those expressions and true if matched
// an array with the given flag if not matched
func normalizeFlag(flag string) ([]string, bool) {
	matches := matchOptionalAssignmentValue.FindStringSubmatch(flag)
	if len(matches) == 3 {
		expressions := make([]string, 2)
		expressions[0] = matches[1]
		expressions[1] = fmt.Sprintf("%v=%v", matches[1], normalizeValuePart(matches[2]))
		return expressions, true
	}
	return []string{flag}, false
}
