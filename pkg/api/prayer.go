package api

import (
	"fmt"
	"time"
)

// Prayer represents a prayer time.
type Prayer struct {
	Name string
	Time time.Time
}

// String returns a string representation of the prayer time in the format of
// "Prayer Name: 15:04 02.01.2006".
func (pt *Prayer) String() string {
	return fmt.Sprintf("%s: %s", pt.Name, pt.Time.Format("15:04 02.01.2006"))
}

// Clock returns a string representation of the clock that the prayer happens.
func (pt *Prayer) ClockString() string {
	return fmt.Sprintf("%s: %s", pt.Name, pt.Time.Format("15:04"))
}
