package twmerge

import "strings"

// TwJoin joins multiple class strings into a single space-separated string,
// filtering out empty strings.
func TwJoin(classLists ...string) string {
	var sb strings.Builder
	for _, classList := range classLists {
		if classList == "" {
			continue
		}
		if sb.Len() > 0 {
			sb.WriteByte(' ')
		}
		sb.WriteString(classList)
	}
	return sb.String()
}
