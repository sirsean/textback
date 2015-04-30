package parse

import (
	. "gopkg.in/check.v1"
	"testing"
	"time"
)

func Test(t *testing.T) { TestingT(t) }

type ParseSuite struct{}

var _ = Suite(&ParseSuite{})

func date(month time.Month, day, hour, minute int) time.Time {
	return time.Date(2015, month, day, hour, minute, 30, 30, time.UTC)
}

func assertDate(c *C, obtained, expected time.Time) {
	c.Assert(obtained.Year(), Equals, expected.Year())
	c.Assert(obtained.Month(), Equals, expected.Month())
	c.Assert(obtained.Day(), Equals, expected.Day())
}
