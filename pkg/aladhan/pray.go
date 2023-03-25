package aladhan

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/ibihim/pray-go/pkg/api"
)

// CalendarToPrayerTimes converts the calendar from the API to a PrayerTimes
// struct.
func CalendarToPrayerTimes(data []*Data) (api.Prayers, error) {
	prayers := api.Prayers{}

	for _, d := range data {
		times, day, err := TimingsToPrayerTimes(d)
		if err != nil {
			return nil, fmt.Errorf("could not parse timings: %v", err)
		}

		if err := prayers.Add(day, times); err != nil {
			return nil, fmt.Errorf("could not add prayer times: %v", err)
		}
	}

	return prayers, nil
}

// TimingsToPrayerTimes converts the timings from the API to a PrayerTimes struct.
func TimingsToPrayerTimes(data *Data) ([]*api.Prayer, api.Day, error) {
	// Get the day from the timestamp
	parseTimestamp, err := strconv.ParseInt(data.Date.Timestamp, 10, 64)
	if err != nil {
		return nil, "", fmt.Errorf("could not parse timestamp: %v", err)
	}
	day := time.Unix(parseTimestamp, 0)

	prayersMap := map[string]string{
		"Fajr":     data.Timings.Fajr,
		"Sunrise":  data.Timings.Sunrise,
		"Dhuhr":    data.Timings.Dhuhr,
		"Asr":      data.Timings.Asr,
		"Sunset":   data.Timings.Sunset,
		"Maghrib":  data.Timings.Maghrib,
		"Isha":     data.Timings.Isha,
		"Midnight": data.Timings.Midnight,
		"Imsak":    data.Timings.Imsak,
	}

	prayers := make([]*api.Prayer, 0, len(prayersMap))
	for name, clock := range prayersMap {
		// API changed "21:14 (CEST)"?
		t, err := time.Parse("15:04", strings.SplitN(clock, " ", 2)[0])
		if err != nil {
			return nil, "", fmt.Errorf("could not parse clock: %v", err)
		}

		prayers = append(prayers, &api.Prayer{
			Name: name,
			Time: t.AddDate(day.Year(), int(day.Month()), day.Day()),
		})
	}

	sort.Slice(prayers, func(i, j int) bool {
		return prayers[i].Time.Before(prayers[j].Time)
	})

	return prayers, api.ToDay(day), nil
}
