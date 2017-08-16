package clock

import "fmt"

const testVersion = 4

type Clock struct {
	Hours, Minutes int
}

func New(hour, minute int) Clock {
	return Clock{
		Hours:   hour,
		Minutes: minute,
	}.normalize()
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.Hours, c.Minutes)
}

func (c Clock) Add(minutes int) Clock {
	c.Minutes += minutes
	return c.normalize()
}

func (c Clock) normalize() Clock {
	// get rid of negative minutes and hours
	if c.Minutes < 0 {
		for {
			c.Hours -= 1
			c.Minutes += 60
			if c.Minutes >= 0 {
				break
			}
		}
	}
	if c.Hours < 0 {
		for {
			c.Hours += 24
			if c.Hours >= 0 {
				break
			}
		}
	}

	c.Hours = (c.Hours + c.Minutes/60) % 24 // normalize hours
	c.Minutes = c.Minutes % 60              // normalize minutes
	return c
}
