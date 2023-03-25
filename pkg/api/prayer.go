package api

import (
	"fmt"
	"time"
)

// Day is a string in the format "YYYY-MM-DD".
type Day string

// ToDay returns the Day for the given time.
func ToDay(t time.Time) Day {
	return Day(t.Format("2006-01-02"))
}

// ValidateDay returns an error if the given day is not in a date in the format
// "YYYY-MM-DD".
func ValidateDay(day string) error {
	_, err := time.Parse("2006-01-02", day)
	return err
}

// Prayers maps days to prayer times.
type Prayers map[Day][]*Prayer

// Add adds prayers to a given day.
func (t Prayers) Add(day Day, times []*Prayer) error {
	if err := ValidateDay(string(day)); err != nil {
		return fmt.Errorf("invalid day: %v", err)
	}

	t[day] = times

	return nil
}

// Get returns the prayer times for a given day.
func (t Prayers) Get(day Day) ([]*Prayer, error) {
	if err := ValidateDay(string(day)); err != nil {
		return nil, fmt.Errorf("invalid day: %v", err)
	}

	return t[day], nil
}

// Prayer represents a prayer time.
type Prayer struct {
	Name string
	Time time.Time
}

// String returns a string representation of the prayer time.
func (pt *Prayer) String() string {
	return fmt.Sprintf("%s: %s", pt.Name, pt.Time.Format("15:04"))
}
