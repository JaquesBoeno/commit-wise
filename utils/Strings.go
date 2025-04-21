package utils

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

func WrapText(str string, maxWidth int) string {
	builder := strings.Builder{}
	wrapRecursive(str, maxWidth, &builder)
	return builder.String()
}

func wrapRecursive(str string, maxWidth int, builder *strings.Builder) {
	for _, line := range strings.Split(str, "\n") {
		if utf8.RuneCountInString(removeANSIEscapeCodes(line)) <= maxWidth {
			builder.WriteString(line)
			builder.WriteString("\n")
		} else {
			breakingPoint := strings.LastIndex(line[:maxWidth], " ")
			if breakingPoint != -1 {
				builder.WriteString(line[:breakingPoint])
				builder.WriteString("\n")
				wrapRecursive(line[breakingPoint+1:], maxWidth, builder)
			} else {
				builder.WriteString(line[:maxWidth])
				builder.WriteString("\n")
				wrapRecursive(line[maxWidth:], maxWidth, builder)
			}
		}
	}
}

func removeANSIEscapeCodes(input string) string {
	re := regexp.MustCompile(`\x1b\[[0-9;]*[A-Za-z]`)
	return re.ReplaceAllString(input, "")
}

func NormalizeNewlines(s string) string {
	for strings.Contains(s, "\n\n\n") {
		s = strings.ReplaceAll(s, "\n\n\n", "\n\n")
	}

	return s
}

func PadEnd(str string, length int, pad rune) string {
	for len(str) < length {
		str += string(pad)
	}

	return str
}
