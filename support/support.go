package support

import (
	"regexp"
)

var detections []string

func Crawler(html string) []string {
	detections = []string{}

	for _, patternGroup := range allPatterns {
		for name, expr := range patternGroup {
			re := regexp.MustCompile(`(?i)` + expr)
			if re.MatchString(html) {
				detections = append(detections, name)
			}
		}
	}
	return detections
}
