package extractor

import (
	"regexp"
)

var regexGroupDelimiter = regexp.MustCompile(`,\s+`)

func splitSynopsis(synopsis string) (groups []string) {
	return regexGroupDelimiter.Split(synopsis, -1)
}
