package parse

import (
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func Date(start time.Time, input string) (parsed string, when time.Time) {
	slashes := regexp.MustCompile(`on\s+([0-9]+)[/-]([0-9]+)[/-]?([0-9]+)?`)
	if strings.Contains(input, "today") {
		parsed = "today"
		when = start.Add(0)
	} else if strings.Contains(input, "tomorrow") {
		parsed = "tomorrow"
		when = start.Add(24 * time.Hour)
	} else if slashMatch := slashes.FindStringSubmatch(input); len(slashMatch) > 0 {
		parsed = slashMatch[0]
		month, _ := strconv.ParseInt(slashMatch[1], 10, 64)
		day, _ := strconv.ParseInt(slashMatch[2], 10, 64)
		year := int64(start.Year()) // default to current year
		if slashMatch[3] != "" {
			year, _ = strconv.ParseInt(slashMatch[3], 10, 64)
		}
		log.Printf("year %v %v", slashMatch[3], year)
		if year < 100 {
			year += 2000
		}
		when = time.Date(int(year), time.Month(month), int(day), start.Hour(), start.Minute(), start.Second(), start.Nanosecond(), start.Location())
	} else {
		when = start.Add(0)
	}
	return
}
