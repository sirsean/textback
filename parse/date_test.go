package parse

import (
	. "gopkg.in/check.v1"
	"time"
)

// today
func (s *ParseSuite) TestToday(c *C) {
	start := date(time.June, 1, 10, 10)
	parsed, when := Date(start, "do things today")
	c.Assert(parsed, Equals, "today")
	assertDate(c, when, start)
}

// tomorrow
func (s *ParseSuite) TestTomorrow(c *C) {
	start := date(time.June, 1, 10, 10)
	parsed, when := Date(start, "do things tomorrow")
	c.Assert(parsed, Equals, "tomorrow")
	assertDate(c, when, start.Add(24*time.Hour))
}

// 4/10/15
func (s *ParseSuite) TestSlash1(c *C) {
	start := date(time.March, 1, 10, 10)
	parsed, when := Date(start, "do things on 4/10/15")
	c.Assert(parsed, Equals, "on 4/10/15")
	assertDate(c, when, date(time.April, 10, 0, 0))
}

// 4/1/15
func (s *ParseSuite) TestSlash2(c *C) {
	start := date(time.March, 1, 10, 10)
	parsed, when := Date(start, "do things on 4/1/15")
	c.Assert(parsed, Equals, "on 4/1/15")
	assertDate(c, when, date(time.April, 1, 0, 0))
}

// 04/10/15
func (s *ParseSuite) TestSlash3(c *C) {
	start := date(time.March, 1, 10, 10)
	parsed, when := Date(start, "do things on 04/10/15")
	c.Assert(parsed, Equals, "on 04/10/15")
	assertDate(c, when, date(time.April, 10, 0, 0))
}

// 04/1/15
func (s *ParseSuite) TestSlash4(c *C) {
	start := date(time.March, 1, 10, 10)
	parsed, when := Date(start, "do things on 04/1/15")
	c.Assert(parsed, Equals, "on 04/1/15")
	assertDate(c, when, date(time.April, 1, 0, 0))
}

// 04/01/15
func (s *ParseSuite) TestSlash5(c *C) {
	start := date(time.March, 1, 10, 10)
	parsed, when := Date(start, "do things on 04/01/15")
	c.Assert(parsed, Equals, "on 04/01/15")
	assertDate(c, when, date(time.April, 1, 0, 0))
}

// 4/10/2015
func (s *ParseSuite) TestSlash6(c *C) {
	start := date(time.March, 1, 10, 10)
	parsed, when := Date(start, "do things on 4/10/2015")
	c.Assert(parsed, Equals, "on 4/10/2015")
	assertDate(c, when, date(time.April, 10, 0, 0))
}

// 5/20
func (s *ParseSuite) TestSlash7(c *C) {
	start := date(time.March, 1, 10, 10)
	parsed, when := Date(start, "do things on 5/20")
	c.Assert(parsed, Equals, "on 5/20")
	assertDate(c, when, date(time.May, 20, 0, 0))
}

// 04/1
func (s *ParseSuite) TestSlash8(c *C) {
	start := date(time.March, 1, 10, 10)
	parsed, when := Date(start, "do things on 04/1")
	c.Assert(parsed, Equals, "on 04/1")
	assertDate(c, when, date(time.April, 1, 0, 0))
}

// 04/02
func (s *ParseSuite) TestSlash9(c *C) {
	start := date(time.March, 1, 10, 10)
	parsed, when := Date(start, "do things on 04/02")
	c.Assert(parsed, Equals, "on 04/02")
	assertDate(c, when, date(time.April, 2, 0, 0))
}

// 4/03
func (s *ParseSuite) TestSlash10(c *C) {
	start := date(time.March, 1, 10, 10)
	parsed, when := Date(start, "do things on 4/03")
	c.Assert(parsed, Equals, "on 4/03")
	assertDate(c, when, date(time.April, 3, 0, 0))
}

// 4-10-15
func (s *ParseSuite) TestSlash11(c *C) {
	start := date(time.March, 1, 10, 10)
	parsed, when := Date(start, "do things on 4-10-15")
	c.Assert(parsed, Equals, "on 4-10-15")
	assertDate(c, when, date(time.April, 10, 0, 0))
}

// 4-10-2015
func (s *ParseSuite) TestSlash12(c *C) {
	start := date(time.March, 1, 10, 10)
	parsed, when := Date(start, "do things on 4-10-2015")
	c.Assert(parsed, Equals, "on 4-10-2015")
	assertDate(c, when, date(time.April, 10, 0, 0))
}

// 5-20
func (s *ParseSuite) TestSlash13(c *C) {
	start := date(time.March, 1, 10, 10)
	parsed, when := Date(start, "do things on 5-20")
	c.Assert(parsed, Equals, "on 5-20")
	assertDate(c, when, date(time.May, 20, 0, 0))
}

// 04-1
// 04-02
// 4-03

// TODO: case insensitive
