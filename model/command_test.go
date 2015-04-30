package model

import (
	. "gopkg.in/check.v1"
	"log"
	"testing"
	"time"
)

func Test(t *testing.T) { TestingT(t) }

type CommandModelSuite struct{}

var _ = Suite(&CommandModelSuite{})

var start time.Time

func init() {
	start = time.Now()
	startTime = func(loc *time.Location) time.Time {
		return start
	}
}

func (s *CommandModelSuite) TestLaterAfternoonNoMinutes(c *C) {
	start = time.Date(2015, time.June, 1, 13, 30, 30, 30, time.UTC)
	command := NewCommandParse("pick up food at 5pm", time.UTC)
	log.Printf("%v", command)
	c.Assert(command.Message, Equals, "pick up food")
	c.Assert(command.When, Equals, time.Date(2015, time.June, 1, 17, 0, 0, 0, time.UTC))
	c.Assert(command.Phone, Equals, "")
}

func (s *CommandModelSuite) TestWithMinutes(c *C) {
	start = time.Date(2015, time.June, 1, 13, 30, 30, 30, time.UTC)
	command := NewCommandParse("pick up food at 5:15pm", time.UTC)
	log.Printf("%v", command)
	c.Assert(command.Message, Equals, "pick up food")
	c.Assert(command.When, Equals, time.Date(2015, time.June, 1, 17, 15, 0, 0, time.UTC))
	c.Assert(command.Phone, Equals, "")
}

func (s *CommandModelSuite) TestNextMorning(c *C) {
	start = time.Date(2015, time.June, 1, 14, 30, 30, 30, time.UTC)
	command := NewCommandParse("pick up food at 10am", time.UTC)
	c.Assert(command.Message, Equals, "pick up food")
	c.Assert(command.When, Equals, time.Date(2015, time.June, 2, 10, 0, 0, 0, time.UTC))
}

func (s *CommandModelSuite) TestWithoutAmpm(c *C) {
	start = time.Date(2015, time.June, 1, 14, 30, 30, 30, time.UTC) // 2:30pm
	command := NewCommandParse("pick up food at 8", time.UTC)       // should guess 8pm
	c.Assert(command.Message, Equals, "pick up food")
	c.Assert(command.When, Equals, time.Date(2015, time.June, 1, 20, 0, 0, 0, time.UTC))
}

func (s *CommandModelSuite) TestWithoutAmpmNextMorning(c *C) {
	start = time.Date(2015, time.June, 1, 14, 30, 30, 30, time.UTC) // 2:30pm
	command := NewCommandParse("pick up food at 1", time.UTC)       // should guess 1am
	c.Assert(command.Message, Equals, "pick up food")
	c.Assert(command.When, Equals, time.Date(2015, time.June, 2, 1, 0, 0, 0, time.UTC))
}

func (s *CommandModelSuite) TestWithoutAmpmStartMorningLaterMorning(c *C) {
	start = time.Date(2015, time.June, 1, 9, 30, 30, 30, time.UTC) // 9:30am
	command := NewCommandParse("pick up food at 11", time.UTC)     // should guess 11am
	c.Assert(command.Message, Equals, "pick up food")
	c.Assert(command.When, Equals, time.Date(2015, time.June, 1, 11, 0, 0, 0, time.UTC))
}

func (s *CommandModelSuite) TestWithoutAmpmStartMorningLaterAfternoon(c *C) {
	start = time.Date(2015, time.June, 1, 9, 30, 30, 30, time.UTC) // 9:30am
	command := NewCommandParse("pick up food at 5", time.UTC)      // should guess 5pm
	c.Assert(command.Message, Equals, "pick up food")
	c.Assert(command.When, Equals, time.Date(2015, time.June, 1, 17, 0, 0, 0, time.UTC))
}

func (s *CommandModelSuite) TestWithDateAndTime(c *C) {
	start = time.Date(2015, time.June, 1, 13, 0, 0, 0, time.UTC)
	command := NewCommandParse("do things at 5 on 6/15", time.UTC)
	c.Assert(command.Message, Equals, "do things")
	c.Assert(command.When, Equals, time.Date(2015, time.June, 15, 17, 0, 0, 0, time.UTC))
}

// TODO: this failed...
/*
Mar 13 17:25:47 cricket textback.linux: 2015/03/13 17:25:47 x [at 4:27 4 27 ]
Mar 13 17:25:47 cricket textback.linux: 2015/03/13 17:25:47 start 2015-03-13 16:25:47.762843741 -0500 CDT
Mar 13 17:25:47 cricket textback.linux: 2015/03/13 17:25:47 h 4, m 27, ampm , hour 4, min 27
Mar 13 17:25:47 cricket textback.linux: 2015/03/13 17:25:47 when hour 16
Mar 13 17:25:47 cricket textback.linux: 2015/03/13 17:25:47 when 2015-03-13 16:25:47.762843741 -0500 CDT
Mar 13 17:25:47 cricket textback.linux: 2015/03/13 17:25:47 when 2015-03-13 17:25:47.762843741 -0500 CDT
Mar 13 17:25:47 cricket textback.linux: 2015/03/13 17:25:47 when 2015-03-13 18:25:47.762843741 -0500 CDT
Mar 13 17:25:47 cricket textback.linux: 2015/03/13 17:25:47 when 2015-03-13 19:25:47.762843741 -0500 CDT
Mar 13 17:25:47 cricket textback.linux: 2015/03/13 17:25:47 when 2015-03-13 20:25:47.762843741 -0500 CDT
Mar 13 17:25:47 cricket textback.linux: 2015/03/13 17:25:47 when 2015-03-13 21:25:47.762843741 -0500 CDT
Mar 13 17:25:47 cricket textback.linux: 2015/03/13 17:25:47 when 2015-03-13 22:25:47.762843741 -0500 CDT
Mar 13 17:25:47 cricket textback.linux: 2015/03/13 17:25:47 when 2015-03-13 23:25:47.762843741 -0500 CDT
Mar 13 17:25:47 cricket textback.linux: 2015/03/13 17:25:47 when 2015-03-14 00:25:47.762843741 -0500 CDT
Mar 13 17:25:47 cricket textback.linux: 2015/03/13 17:25:47 when 2015-03-14 01:25:47.762843741 -0500 CDT
Mar 13 17:25:47 cricket textback.linux: 2015/03/13 17:25:47 when 2015-03-14 02:25:47.762843741 -0500 CDT
Mar 13 17:25:47 cricket textback.linux: 2015/03/13 17:25:47 when 2015-03-14 03:25:47.762843741 -0500 CDT
Mar 13 17:25:47 cricket textback.linux: 2015/03/13 17:25:47 s 2015-03-14 04:27:00 -0500 CDT
Mar 13 17:25:47 cricket textback.linux: 2015/03/13 17:25:47 return at 4:27, 2015-03-14 04:27:00 -0500 CDT
Mar 13 17:25:47 cricket textback.linux: 2015/03/13 17:25:47 Test things
Mar 13 17:25:47 cricket textback.linux: 2015/03/13 17:25:47 2015-03-14 04:27:00 -0500 CDT
Mar 13 17:25:47 cricket textback.linux: 2015/03/13 17:25:47 {Test things 2015-03-14 04:27:00 -0500 CDT +17735048753}
*/
