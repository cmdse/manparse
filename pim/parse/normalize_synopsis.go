package parse

import (
	"regexp"
	"strings"
)

var separatorRegex = regexp.MustCompile(`(=|\s+)`)
var quotedStringRegex = regexp.MustCompile(`^(".+"|'.+')$`)

func normalizeQuotedValue(valuePart string) string {
	matchQuoted := quotedStringRegex.FindStringSubmatch(valuePart)
	if len(matchQuoted) < 1 {
		// noop
		return valuePart
	} else {
		trimed := strings.Trim(valuePart, `'"`)
		slugged := strings.Replace(trimed, " ", "-", -1)
		return slugged
	}
}

func normalizeValuePart(valuePart string) string {
	return normalizeQuotedValue(valuePart)
}

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

// This function normalize the assignment part when containing special meta-character
// such as ellipsis, square brackets, quoted strings...
func normalizeOptSynopsis(synopsis string) string {
	parts := splitSynopsisToParts(synopsis)
	if len(parts) == 1 {
		// no assignment part, no-op
		return parts[0]
	} else {
		sep := findSeparator(synopsis)
		// must be of len 2
		parts[1] = normalizeValuePart(parts[1])
		return strings.Join(parts, sep)
	}

}

func normalizeOptSynopsises(expressions []string) []string {
	normalizedExpressions := make([]string, len(expressions))
	for i := range expressions {
		normalizedExpressions[i] = normalizeOptSynopsis(expressions[i])
	}
	return normalizedExpressions
}
