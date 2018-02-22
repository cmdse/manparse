package normalize

import (
	"regexp"
	"strings"
)

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
