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
	// dayInHours is the number of hours a day has.
	dayInHours = 24.0
)

// NextPrayer returns the next prayer time.
// It is expected that the prayers are sorted by time.
func NextPrayer(prayers []*api.Prayer) (*api.Prayer, error) {
	nextPrayers, err := NextPrayers(prayers)
	if err != nil {
		return nil, err
	}

	return nextPrayers[0], nil
}

// NextPrayers returns the prayer times after now.
// It is expected that the prayers are sorted by time.
func NextPrayers(prayers []*api.Prayer) ([]*api.Prayer, error) {
	now := time.Now()

	klog.V(8).Infof("Now: %s", now.Format("02.01.2006 15:04"))
	klog.V(8).Infof("PrayerTimes: %v", prayers)

	for i, prayer := range prayers {
		if now.Before(prayer.Time) {
			d := prayer.Time.Sub(now)
			if d.Hours() > dayInHours {
				return nil, fmt.Errorf(
					"difference is more than 24 hours, now: %s, next: %s",
					now, prayer.Time,
				)
			}

			return prayers[i:], nil
		}
	}

	return nil, ErrNoPrayerFound
}
