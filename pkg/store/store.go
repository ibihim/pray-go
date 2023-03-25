package store

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/adrg/xdg"

	"github.com/ibihim/pray-go/pkg/api"
)

var (
	// cacheFilePath is the path to the cache file.
	cacheFilePath = filepath.Join(xdg.CacheHome, "pray-go", "data.json")
)

// GetAll returns the prayer times from the cache file. If the
// file doesn't exist, it returns nil, nil instead of an error.
func GetAll() (api.Prayers, error) {
	f, err := os.Stat(cacheFilePath)
	if err != nil {
		return nil, fmt.Errorf("could not stat cache directory: %w", err)
	}
	if f.IsDir() {
		return nil, fmt.Errorf("cache directory is a directory")
	}

	b, err := ioutil.ReadFile(cacheFilePath)
	if err != nil {
		return nil, fmt.Errorf("could not read data file: %w", err)
	}
	if len(b) == 0 {
		return nil, nil
	}

	var prayers api.Prayers
	if err := json.Unmarshal(b, &prayers); err != nil {
		return nil, fmt.Errorf("could not unmarshal data file: %w", err)
	}

	return prayers, nil
}

// StoreAll stores the prayer times in the cache file.
func StoreAll(prayerTimes api.Prayers) error {
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

// CreateFile creates the file, if it doesn't exist. Does nothing if it does.
// Creates also all the directories in the path, if they don't exist.
func CreateFile(filePath string) error {
	fileDir := filepath.Dir(filePath)
	if err := os.MkdirAll(fileDir, 0755); err != nil {
		return fmt.Errorf("could not create file directory: %w", err)
	}

	_, err := os.Stat(filePath)
	switch {
	case os.IsNotExist(err):
		file, err := os.Create(filePath)
		if err != nil {
			return fmt.Errorf("could not create file: %w", err)
		}
		defer file.Close()
	case err != nil:
		return fmt.Errorf("could not stat file: %w", err)
	}
	return nil
}
