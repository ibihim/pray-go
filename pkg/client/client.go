package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"k8s.io/klog/v2"

	"github.com/ibihim/pray-go/api"
)

// GetCalendarByCity returns the prayer times for a given month and city (country).
func GetCalendarByCity(t time.Time, city, country string, method int) ([]api.Data, error) {
	return GetCalendarByCityWithClient(http.Client{}, t, city, country, method)
}

// GetCalendarByCityWithClient returns the prayer times for a given month and cit (country).
func GetCalendarByCityWithClient(client http.Client, t time.Time, city, country string, method int) ([]api.Data, error) {
	timeString := t.Format("2006/01")
	urlStr, err := url.JoinPath(api.CalendarByCityURL, timeString)
	if err != nil {
		return nil, fmt.Errorf("error joining URL: %w", err)
	}

	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, fmt.Errorf("error parsing URL: %w", err)
	}

	query := url.Values{}
	query.Set("city", city)
	query.Set("country", country)
	query.Set("method", strconv.Itoa(method))

	u.RawQuery = query.Encode()

	klog.V(4).Infof("URL: %s", u.String())

	res, err := client.Get(u.String())
	if err != nil {
		return nil, fmt.Errorf("error getting response: %w", err)
	}
	defer res.Body.Close()

	klog.V(4).Infof("Response: %v", res)

	var timings api.CalendarByCity
	if err := json.NewDecoder(res.Body).Decode(&timings); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	klog.V(4).Infof("Timings: %s", timings.String())

	return timings.Data, nil
}

// GetTimingsByCity returns the prayer times for a given day, city and country.
func GetTimingsByCity(t time.Time, city, country string, method int) ([]api.Data, error) {
	return GetTimingsByCityWithClient(http.Client{}, t, city, country, method)
}

// GetTimingsByCityWithClient returns the prayer times for a given day, city and
// country using a custom HTTP client.
func GetTimingsByCityWithClient(client http.Client, t time.Time, city, country string, method int) ([]api.Data, error) {
	timeString := t.Format("02-01-2006")
	urlStr, err := url.JoinPath(api.TimingsByCityURL, timeString)
	if err != nil {
		return nil, fmt.Errorf("error joining URL: %w", err)
	}

	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, fmt.Errorf("error parsing URL: %w", err)
	}

	query := url.Values{}
	query.Set("city", city)
	query.Set("country", country)
	query.Set("method", strconv.Itoa(method))

	u.RawQuery = query.Encode()

	klog.V(4).Infof("URL: %s", u.String())

	res, err := client.Get(u.String())
	if err != nil {
		return nil, fmt.Errorf("error getting response: %w", err)
	}
	defer res.Body.Close()

	klog.V(4).Infof("Response: %v", res)

	var timings api.TimingsByCity
	if err := json.NewDecoder(res.Body).Decode(&timings); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	klog.V(4).Infof("Timings: %s", timings.String())

	return []api.Data{timings.Data}, nil
}

// CreateURL creates a URL for the TimingsByCity API based on the given date,
// city, country and method.
func CreateURL(t time.Time, city, country string, method int) (*url.URL, error) {
	timeString := t.Format("02-01-2006")
	urlStr, err := url.JoinPath(api.TimingsByCityURL, timeString)
	if err != nil {
		return nil, fmt.Errorf("error joining URL: %w", err)
	}

	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, fmt.Errorf("error parsing URL: %w", err)
	}

	query := url.Values{}
	query.Set("city", city)
	query.Set("country", country)
	query.Set("method", strconv.Itoa(method))

	u.RawQuery = query.Encode()

	return u, nil
}
