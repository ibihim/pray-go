package prayer

import (
	"errors"
	"fmt"
	"time"

	"k8s.io/klog/v2"

	"github.com/ibihim/pray-go/pkg/api"
)

var (
	// ErrNoPrayerFound is returned when no prayer was found.
	ErrNoPrayerFound = errors.New("no prayer found")
)

// Next returns the next prayer time in the format "Prayername: HH:MM".
// Returns an error if no prayer was found.
func Next(prayers api.Prayers) (string, error) {
	now := time.Now()

	klog.V(4).Infof("Now: %s", now.Format("02.01.2006 15:04"))
	klog.V(4).Infof("PrayerTimes: %v", prayers)

	todayPrayers, err := prayers.Get(api.ToDay(now))
	if err != nil {
		return "", fmt.Errorf("Couldn't get today's prayer times: %v", err)
	}

	if prayer := findLater(todayPrayers, now); prayer != nil {
		return prayer.String(), nil
	}

	tomorrowPrayers, err := prayers.Get(api.ToDay(now.AddDate(0, 0, 1)))
	if err != nil {
		return "", fmt.Errorf("Couldn't get tomorrow's prayer times: %v", err)
	}

	if prayer := findLater(tomorrowPrayers, now); prayer != nil {
		return prayer.String(), nil
	}

	klog.Errorf("Couldn't find a prayer. PrayerTimes: %v", todayPrayers)

	return "", ErrNoPrayerFound
}

func findLater(prayers []*api.Prayer, now time.Time) *api.Prayer {
	for _, prayer := range prayers {
		if now.Before(prayer.Time) {
			return prayer
		}
	}

	return nil
}
