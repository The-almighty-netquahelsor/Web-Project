package asciiart

import "strings"

func Render(input string, banner []string) string {
	var sb strings.Builder

	input = strings.ReplaceAll(input, "\r\n", "\n")
	lines := strings.Split(input, "\n")

	for i, line := range lines {
		if line == "" {
			for row := 0; row < 8; row++ {
				sb.WriteString("\n")
			}
			continue
		}
		for row := 0; row < 8; row++ {
			for _, ch := range line {
				if ch < 32 || ch > 126 {
					continue
				}
				chartStart := (int(ch) - 32) * 9
				if chartStart+row >= len(banner) {
					continue
				}
				sb.WriteString(banner[chartStart+row])
			}
			sb.WriteString("\n")
		}
		if i < len(lines)-1 {
			sb.WriteString("\n")
		}
	}
	return sb.String()
}
