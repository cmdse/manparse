package normalize

import (
	"regexp"
	"strings"
)

var separatorRegex = regexp.MustCompile(`(=|\s+)`)

func splitSynopsisToParts(synopsis string) []string {
	return separatorRegex.Split(synopsis, 2)
}

func findSeparator(synopsis string) string {
	sep := separatorRegex.FindStringSubmatch(synopsis)
	if len(sep) > 1 {
		return sep[1]
	}
	return ""
}

// normalizeOptSynopsis normalize the assignment part when containing special meta-characters
// such as ellipsis, square brackets, quoted strings...
// It returns a slice of expressions, typically of length 1 but occasionally length 2 when matching
// optional assignment value.
func normalizeOptSynopsis(synopsis string) string {
	parts := splitSynopsisToParts(synopsis)
	if len(parts) == 1 {
		// no assignment part, no-op
		return synopsis
	} else {
		sep := findSeparator(synopsis)
		// must be of len 2
		parts[1] = normalizeValuePart(parts[1])
		return strings.Join(parts, sep)
	}

}

// NormalizeOptDescriptions normalize the assignment part when containing special meta-character
// such as ellipsis, square brackets, quoted strings for given expressions.
// Length of the returned slice might be higher then expressions', when special edge-cases are encountered.
func NormalizeOptDescriptions(expressions []string) []string {
	normalizedExpressions := make([]string, 0, len(expressions))
	for i := range expressions {
		normalizedSynopsis := normalizeOptSynopsis(expressions[i])
		normalizedExpressions = append(normalizedExpressions, normalizedSynopsis)
	}
	return normalizedExpressions
}
