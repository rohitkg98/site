package helpers

import (
	"strings"
)

func BreakAndSanitize(content []byte) []string {
	return strings.Split(string(content), "\n")
}
