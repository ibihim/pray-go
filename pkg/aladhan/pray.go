package aladhan

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/ibihim/pray-go/pkg/api"
	"k8s.io/klog/v2"
)

// GetPrayers returns the prayers for the given city and nation.
func GetPrayers(time time.Time, city, nation string, method int) ([]*api.Prayer, error) {
	data, err := GetCalendarByCity(time, city, nation, method)
	if err != nil {
		return nil, fmt.Errorf("could not get calendar: %v", err)
	}

	prayers, err := DataToPrayers(data)
	if err != nil {
		return nil, fmt.Errorf("could not parse timings: %v", err)
	}

	return prayers, nil
}

// DataToPrayers converts the data slice from the API to a Prayer slice.
func DataToPrayers(data []*Data) ([]*api.Prayer, error) {
	prayers := []*api.Prayer{}

	for _, d := range data {
		p, err := dataToPrayers(d)
		if err != nil {
			return nil, fmt.Errorf("could not parse timings: %v", err)
		}

		prayers = append(prayers, p...)
	}

	sort.Slice(prayers, func(i, j int) bool {
		return prayers[i].Time.Before(prayers[j].Time)
	})

	return prayers, nil
}

// dataToPrayers converts the timings from the API to a Prayer slice.
func dataToPrayers(data *Data) ([]*api.Prayer, error) {
	klog.V(8).Infof("Got timestamp: %s", data.Date.Timestamp)

	// Get the day from the timestamp
	parseTimestamp, err := strconv.ParseInt(data.Date.Timestamp, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("could not parse timestamp: %v", err)
	}
	day := time.Unix(parseTimestamp, 0)
	klog.V(8).Infof("Made it day: %s", day)

	rawPrayers := []*struct {
		Name    string
		TimeStr string
	}{
		{"Fajr", data.Timings.Fajr},
		{"Sunrise", data.Timings.Sunrise},
		{"Dhuhr", data.Timings.Dhuhr},
		{"Asr", data.Timings.Asr},
		{"Sunset", data.Timings.Sunset},
		{"Maghrib", data.Timings.Maghrib},
		{"Isha", data.Timings.Isha},
		{"Midnight", data.Timings.Midnight},
		{"Imsak", data.Timings.Imsak},
	}

	prayers := make([]*api.Prayer, 0, len(rawPrayers))
	for _, rawPrayer := range rawPrayers {
		// API changed "21:14 (CEST)"?
		c, err := time.Parse("15:04", strings.SplitN(rawPrayer.TimeStr, " ", 2)[0])
		if err != nil {
			return nil, fmt.Errorf("could not parse clock: %v", err)
		}

		prayers = append(prayers, &api.Prayer{
			Name: rawPrayer.Name,
			Time: time.Date(
				day.Year(),
				day.Month(),
				day.Day(),
				c.Hour(),
				c.Minute(),
				c.Second(),
				c.Nanosecond(),
				day.Location(),
			),
		})
	}

	return prayers, nil
}
