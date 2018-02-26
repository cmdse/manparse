package reporter

import (
	"bytes"
	"fmt"
	"strings"
)

func writeOffset(offset int, str string) string {
	var buffer bytes.Buffer
	buffer.Grow(100)
	lines := strings.Split(str, "\n")
	for _, line := range lines {
		buffer.WriteString(fmt.Sprintf("%v%v\n", strings.Repeat("  ", offset), line))
	}
	return buffer.String()
}
