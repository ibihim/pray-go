package aladhan

type Time string

type Timings struct {
	Fajr     Time `json:"Fajr"`
	Sunrise  Time `json:"Sunrise"`
	Dhuhr    Time `json:"Dhuhr"`
	Asr      Time `json:"Asr"`
	Sunset   Time `json:"Sunset"`
	Maghrib  Time `json:"Maghrib"`
	Isha     Time `json:"Isha"`
	Midnight Time `json:"Midnight"`
}

type DetailedDate struct {
	English string `json:"en"`
	Number  int    `json:"number"`
}

type Date struct {
	Date    string       `json:"date"`
	Format  string       `json:"format"`
	Day     string       `json:"day"`
	Weekday DetailedDate `json:"weekday"`
	Month   DetailedDate `json:"month"`
	Year    string       `json:"year"`
}

type Method struct {
	ID     int            `json:"id"`
	Name   string         `json:"name"`
	Params map[string]int `json:"params"`
}

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

type Data struct {
	Readable  string `json:"readable"`
	Timestamp string `json:"timestamp"`
	Gregorian Date   `json:"gregorian"`
	Hijri     Date   `json:"hijri"`
	Meta      Meta   `json:"timestamp"`
}

type Data struct {
	Timings Timings `json:"Timings"`
	Data    Data    `json:"Date"`
}

type Response struct {
	HTTPCode int    `json:"code"`
	Payload  []Data `json:"data"`
}
