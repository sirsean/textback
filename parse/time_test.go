package parse

import (
	. "gopkg.in/check.v1"
	"time"
)

// start: afternoon, end: later afternoon, minutes: no, ampm: pm
func (s *ParseSuite) TestStartAfternoonEndAfternoonPm(c *C) {
	start := date(time.June, 1, 13, 30)
	parsed, when := Time(start, "do things at 5pm")
	c.Assert(parsed, Equals, "at 5pm")
	c.Assert(when, Equals, date(time.June, 1, 17, 0).Truncate(time.Minute))
}

// start: afternoon, end: later afternoon, minutes: yes, ampm: pm
func (s *ParseSuite) TestStartAfternoonEndAfternoonMinutesPm(c *C) {
	start := date(time.June, 1, 13, 30)
	parsed, when := Time(start, "do things at 5:30pm")
	c.Assert(parsed, Equals, "at 5:30pm")
	c.Assert(when, Equals, date(time.June, 1, 17, 30).Truncate(time.Minute))
}

// start: afternoon, end: next morning, minutes: no, ampm: am
func (s *ParseSuite) TestStartAfternoonEndMorningAm(c *C) {
	parsed, when := Time(date(time.June, 1, 13, 30), "do things at 10am")
	c.Assert(parsed, Equals, "at 10am")
	c.Assert(when, Equals, date(time.June, 2, 10, 0).Truncate(time.Minute))
}

// start: afternoon, end: later afternoon, minutes: no, ampm: no
func (s *ParseSuite) TestStartAfternoonEndAfternoon(c *C) {
	parsed, when := Time(date(time.June, 1, 13, 30), "do things at 4")
	c.Assert(parsed, Equals, "at 4")
	c.Assert(when, Equals, date(time.June, 1, 16, 0).Truncate(time.Minute))
}

// start: afternoon, end: next morning, minutes: no, ampm: no
func (s *ParseSuite) TestStartAfternoonEndMorning(c *C) {
	parsed, when := Time(date(time.June, 1, 14, 30), "do things at 1")
	c.Assert(parsed, Equals, "at 1")
	c.Assert(when, Equals, date(time.June, 2, 1, 0).Truncate(time.Minute))
}

// start: afternoon, end: next morning, minutes: no, ampm: no (starting earlier that hour)
func (s *ParseSuite) TestStartAfternoonEndMorningEarlierThatHour(c *C) {
	parsed, when := Time(date(time.June, 1, 13, 30), "do things at 1")
	c.Assert(parsed, Equals, "at 1")
	c.Assert(when, Equals, date(time.June, 2, 1, 0).Truncate(time.Minute))
}

// start: morning, end: later morning, minutes: no, ampm: no
func (s *ParseSuite) TestStartMorningEndMorning(c *C) {
	parsed, when := Time(date(time.June, 1, 8, 30), "do things at 10")
	c.Assert(parsed, Equals, "at 10")
	c.Assert(when, Equals, date(time.June, 1, 10, 0).Truncate(time.Minute))
}

// start: morning, end: afternoon, minutes: no, ampm: no
func (s *ParseSuite) TestStartMorningEndAfternoon(c *C) {
	parsed, when := Time(date(time.June, 1, 8, 30), "do things at 4")
	c.Assert(parsed, Equals, "at 4")
	c.Assert(when, Equals, date(time.June, 1, 16, 0).Truncate(time.Minute))
}

// start: afternoon, end: afternoon within the hour, minutes: yes, ampm: no
func (s *ParseSuite) TestWithinTheHour(c *C) {
	parsed, when := Time(date(time.June, 1, 14, 25), "do things at 2:29")
	c.Assert(parsed, Equals, "at 2:29")
	c.Assert(when, Equals, date(time.June, 1, 14, 29).Truncate(time.Minute))
}

// start: morning, end: morning within the hour, minutes: yes, ampm: no
func (s *ParseSuite) TestWithinTheHourMorning(c *C) {
	parsed, when := Time(date(time.June, 1, 8, 25), "do things at 8:29")
	c.Assert(parsed, Equals, "at 8:29")
	c.Assert(when, Equals, date(time.June, 1, 8, 29).Truncate(time.Minute))
}

// start: morning, end: morning within the hour earlier minutes, minutes: yes, ampm: no
func (s *ParseSuite) TestWithinTheHourMorningEarlierMinutes(c *C) {
	parsed, when := Time(date(time.June, 1, 8, 25), "do things at 8:20")
	c.Assert(parsed, Equals, "at 8:20")
	c.Assert(when, Equals, date(time.June, 1, 20, 20).Truncate(time.Minute))
}

// TODO: noon/midnight tests
// TODO: invalid inputs
// TODO: inputs without time
// TODO: different cases for AT and AM/PM
// TODO: support for @ symbol instead of "at"
