package extractor

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/cmdse/core/schema"
)

// match patterns like: -tag[=optional-value]
var matchOptionalExplicitAssignment = regexp.MustCompile(`^(\S+)\[=(\S+)\]$`)

// match patterns like: -tag [optional-value]
var matchOptionalImplicitAssignmentValue = regexp.MustCompile(`^(\S+)\s+\[(\S+)\]$`)

// match option expressions delimiters in option synopses
var regexGroupDelimiter = regexp.MustCompile(`,\s+`)

func splitSynopsis(synopsis string) (groups []string) {
	return regexGroupDelimiter.Split(synopsis, -1)
}

func formatVariantNames(variants []*schema.OptExpressionVariant) string {
	names := make([]string, len(variants))
	for i, variant := range variants {
		names[i] = variant.Name()
	}
	return strings.Join(names, ", ")
}

func findAssignment(expression string, template string, matcher *regexp.Regexp) ([]string, bool) {
	matches := matcher.FindStringSubmatch(expression)
	if len(matches) == 3 {
		expressions := make([]string, 2)
		expressions[0] = matches[1]
		expressions[1] = fmt.Sprintf(template, matches[1], matches[2])
		return expressions, true
	}
	return []string{expression}, false
}

// findOptionalExplicitAssignment tries to match an optional value assignment pattern
// it returns an array of those expressions and true if matched
// an array with the given flag if not matched
func findOptionalExplicitAssignment(expression string) ([]string, bool) {
	return findAssignment(expression, "%v=%v", matchOptionalExplicitAssignment)
}

// findOptionalExplicitAssignment tries to match an optional value assignment pattern
// it returns an array of those expressions and true if matched
// an array with the given flag if not matched
func findOptionalImplicitAssignment(expression string) ([]string, bool) {
	return findAssignment(expression, "%v %v", matchOptionalImplicitAssignmentValue)
}
