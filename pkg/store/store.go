package store

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
	"k8s.io/klog/v2"

	"github.com/ibihim/pray-go/pkg/api"
)

var (
	// cacheFilePath is the path to the cache file.
	cacheFilePath = filepath.Join(xdg.CacheHome, "pray-go", "data.json")
)

// GetCacheFilePath returns the path to the cache file.
func GetCacheFilePath() string {
	return cacheFilePath
}

// Get returns the prayer times from the cache file. If the
// file doesn't exist, it returns nil, nil instead of an error.
func Get() ([]*api.Prayer, error) {
	klog.V(4).Infof("Cache file path: %s", cacheFilePath)

	info, err := os.Stat(cacheFilePath)
	if err != nil {
		return nil, fmt.Errorf("could not stat cache directory: %w", err)
	}
	if info.IsDir() {
		return nil, fmt.Errorf("cache directory is a directory")
	}

	f, err := os.Open(cacheFilePath)
	if err != nil {
		return nil, fmt.Errorf("could not open cache file: %w", err)
	}
	defer f.Close()

	var prayers []*api.Prayer
	if err := json.NewDecoder(f).Decode(&prayers); err != nil {
		return nil, fmt.Errorf("could not decode cache file: %w", err)
	}

	if len(prayers) == 0 {
		return nil, fmt.Errorf("cache file is empty")
	}

	return prayers, nil
}

// Store stores the prayer times in the cache file.
func Store(prayerTimes []*api.Prayer) error {
	klog.V(4).Infof("Cache file path: %s", cacheFilePath)

	cacheDirPath := filepath.Dir(cacheFilePath)

	if err := os.MkdirAll(cacheDirPath, 0755); err != nil {
		return fmt.Errorf("could not create file directory: %w", err)
	}

	var file *os.File
	_, err := os.Stat(cacheFilePath)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("could not stat file: %w", err)
	}

	file, err = os.Create(cacheFilePath)
	if err != nil {
		return fmt.Errorf("could not create file: %w", err)
	}
	defer file.Close()

	if err := json.NewEncoder(file).Encode(prayerTimes); err != nil {
		return fmt.Errorf("could not encode prayer times: %w", err)
	}

	return nil
}
