package parse

import (
	"log"
	"regexp"
	"strconv"
	"time"
)

func Time(start time.Time, input string) (parsed string, when time.Time) {
	log.Printf("YO")
	r := regexp.MustCompile(`at\s+([0-9]+):?([0-9]+)?(am|pm)?`)
	x := r.FindStringSubmatch(input)
	log.Printf("x %v", x)
	if len(x) == 0 {
		// TODO there is no time for this command
	} else {
		parsed = x[0]
		if len(x) > 1 {
			when = start.Add(0)
			log.Printf("start %v", when)
			h := x[1]
			m := x[2]
			ampm := x[3]
			hour, _ := strconv.ParseInt(h, 10, 64)
			minute, _ := strconv.ParseInt(m, 10, 64)
			if ampm == "" {
				if when.Hour() > 12 {
					if int64(when.Hour()) < hour+12 {
						hour += 12
					} else if int64(when.Hour()) == hour+12 {
						if int64(when.Minute()) <= minute {
							hour += 12
						}
					}
				} else {
					if int64(when.Hour()) > hour {
						hour += 12
					} else if int64(when.Hour()) == hour {
						if int64(when.Minute()) >= minute {
							hour += 12
						}
					}
				}
			}
			if ampm == "pm" {
				hour += 12
			}
			log.Printf("h %v, m %v, ampm %v, hour %v, min %v", h, m, ampm, hour, minute)
			log.Printf("when hour %v (%v)", when.Hour(), hour)
			for int64(when.Hour()) != hour {
				log.Printf("when %v", when)
				when = when.Add(time.Hour)
			}
			when = when.Truncate(time.Hour)
			when = when.Add(time.Duration(minute) * time.Minute)
			log.Printf("s %v", when)
		} else {
			// TODO there is no time for this command
		}
	}
	log.Printf("return %v, %v", parsed, when)
	return
}
