package model

import (
	"github.com/sirsean/textback/parse"
	"log"
	"strings"
	"time"
)

var startTime = func(loc *time.Location) time.Time {
	return time.Now().In(loc)
}

type Command struct {
	Message string
	When    time.Time
	Phone   string
}

func NewCommandParse(original string, loc *time.Location) Command {
	remaining, when := parseAt(original, loc)
	log.Printf(remaining)
	log.Printf("%v", when)
	return Command{
		Message: remaining,
		When:    when,
	}
}

func NewCommandParseWithPhone(original, phone string, loc *time.Location) Command {
	c := NewCommandParse(original, loc)
	c.Phone = phone
	return c
}

func FormatKey(t time.Time) string {
	return t.In(time.UTC).Format("1/2/2006 15:04")
}

func (c Command) WhenKey() string {
	return FormatKey(c.When)
}

func parseAt(in string, loc *time.Location) (remaining string, when time.Time) {
	remaining = in
	when = startTime(loc)
	parsed, when := parse.Date(when, remaining)
	if parsed != "" {
		remaining = strings.Replace(remaining, parsed, "", 1)
	}
	parsed, when = parse.Time(when, remaining)
	remaining = strings.Replace(remaining, parsed, "", 1)
	remaining = strings.TrimSpace(remaining)
	return
}
