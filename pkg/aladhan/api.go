package aladhan

import "encoding/json"

const (
	TimingsByCityURL  = "https://api.aladhan.com/v1/timingsByCity"
	CalendarByCityURL = "https://api.aladhan.com/v1/calendarByCity"
)

// CalendarByCity is the main struct for the CalendarByCity API response
type CalendarByCity struct {
	Code   int     `json:"code"`
	Status string  `json:"status"`
	Data   []*Data `json:"data"`
}

func (c *CalendarByCity) String() string {
	b, err := json.Marshal(c)
	if err != nil {
		return err.Error()
	}
	return string(b)
}

// TimingsByCity is the main struct for the TimingsByCity API response
type TimingsByCity struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   Data   `json:"data"`
}

// String returns a string representation of the TimingsByCity struct.
func (t *TimingsByCity) String() string {
	b, err := json.Marshal(t)
	if err != nil {
		return err.Error()
	}
	return string(b)
}

// Data is the part of the TimingsByCity API response that contains the actual
// data.
type Data struct {
	Timings Timings `json:"timings"`
	Date    Date    `json:"date"`
	Meta    Meta    `json:"meta"`
}

// Timings is the most interesting part of the TimingsByCity API response as it
// contains the actual prayer times for the day.
type Timings struct {
	Fajr     string `json:"Fajr"`
	Sunrise  string `json:"Sunrise"`
	Dhuhr    string `json:"Dhuhr"`
	Asr      string `json:"Asr"`
	Sunset   string `json:"Sunset"`
	Maghrib  string `json:"Maghrib"`
	Isha     string `json:"Isha"`
	Imsak    string `json:"Imsak"`
	Midnight string `json:"Midnight"`
}

// Date is the part of the TimingsByCity API response that contains the date in
// various formats.
type Date struct {
	Readable  string    `json:"readable"`
	Timestamp string    `json:"timestamp"`
	Gregorian Gregorian `json:"gregorian"`
	Hijri     Hijri     `json:"hijri"`
}

// Gregorian is the Western calendar date.
type Gregorian struct {
	Date        string      `json:"date"`
	Format      string      `json:"format"`
	Day         string      `json:"day"`
	Weekday     Weekday     `json:"weekday"`
	Month       Month       `json:"month"`
	Year        string      `json:"year"`
	Designation Designation `json:"designation"`
}

// Hijri is the Islamic calendar date.
type Hijri struct {
	Date        string      `json:"date"`
	Format      string      `json:"format"`
	Day         string      `json:"day"`
	Weekday     Weekday     `json:"weekday"`
	Month       Month       `json:"month"`
	Year        string      `json:"year"`
	Designation Designation `json:"designation"`
	Holidays    []string    `json:"holidays"`
}

// Weekday is the weekday for both calendars in english and arabic.
type Weekday struct {
	En string `json:"en"`
	Ar string `json:"ar,omitempty"`
}

// Month is the month for both calendars in english and arabic.
type Month struct {
	Number int    `json:"number"`
	En     string `json:"en"`
	Ar     string `json:"ar,omitempty"`
}

// Designation is Anno Domini / Anno Hegirae.
type Designation struct {
	Abbreviated string `json:"abbreviated"`
	Expanded    string `json:"expanded"`
}

// Meta contains the location, the method and the offsets used to calculate the
// prayer times.
type Meta struct {
	Latitude                 float64 `json:"latitude"`
	Longitude                float64 `json:"longitude"`
	Timezone                 string  `json:"timezone"`
	Method                   Method  `json:"method"`
	LatitudeAdjustmentMethod string  `json:"latitudeAdjustmentMethod"`
	MidnightMode             string  `json:"midnightMode"`
	School                   string  `json:"school"`
	Offset                   Offset  `json:"offset"`
}

// Method contains the method used to calculate the prayer times.
type Method struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Params struct {
		Fajr int `json:"Fajr"`
		Isha int `json:"Isha"`
	} `json:"params"`
}

// Offset contains the offsets used to calculate the prayer times.
type Offset struct {
	Imsak    int `json:"Imsak"`
	Fajr     int `json:"Fajr"`
	Sunrise  int `json:"Sunrise"`
	Dhuhr    int `json:"Dhuhr"`
	Asr      int `json:"Asr"`
	Maghrib  int `json:"Maghrib"`
	Sunset   int `json:"Sunset"`
	Isha     int `json:"Isha"`
	Midnight int `json:"Midnight"`
}
