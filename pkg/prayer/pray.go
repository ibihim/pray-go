package prayer

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"time"

	"k8s.io/klog/v2"

	"github.com/ibihim/pray-go/api"
)

type PrayTime struct {
	Name string
	Time time.Time
}

func (pt *PrayTime) String() string {
	return fmt.Sprintf("%s: %s", pt.Name, pt.Time.Format("15:04"))
}

// Next returns the next prayer time in the format "Prayer: HH:MM".
func Next(now time.Time, timings []api.Data) (string, error) {
	prayerTimes := []*PrayTime{}

	for _, timing := range timings {
		prayers, err := TimingsToPrayerTimes(timing)
		if err != nil {
			return "", fmt.Errorf("could not parse timings: %v", err)
		}

		prayerTimes = append(prayerTimes, prayers...)
	}

	klog.V(4).Infof("Now: %s", now.Format("02.01.2006 15:04"))

	for _, prayer := range prayerTimes {
		if now.Before(prayer.Time) {
			return prayer.String(), nil
		}
	}

	klog.Errorf("Couldn't find a prayer. PrayerTimes: %v", prayerTimes)

	return "", errors.New("Couldn't find a prayer!")
}

// TimingsToPrayerTimes converts the timings from the API to a PrayerTimes struct.
func TimingsToPrayerTimes(data api.Data) ([]*PrayTime, error) {
	parseTimestamp, err := strconv.ParseInt(data.Date.Timestamp, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("could not parse timestamp: %v", err)
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

	prayers := make([]*PrayTime, 0, len(prayersMap))
	for name, clock := range prayersMap {
		t, err := time.Parse("15:04", clock)
		if err != nil {
			return nil, fmt.Errorf("could not parse clock: %v", err)
		}

		prayers = append(prayers, &PrayTime{
			Name: name,
			Time: t.AddDate(day.Year(), int(day.Month()), day.Day()),
		})
	}

	sort.Slice(prayers, func(i, j int) bool {
		return prayers[i].Time.Before(prayers[j].Time)
	})

	return prayers, nil
}
