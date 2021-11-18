package app

import (
	"os"

	"github.com/mitchellh/go-homedir"
)

const (
	// SystemBinariesDir is the default binaries directory used in system mode
	SystemBinariesDir string = "/var/lib/binenv"
	// SystemCacheDir is the default cache directory used in system mode
	SystemCacheDir string = "/var/cache/binenv"
	// SystemDistributionsDir is the default distributions directory used in system mode
	SystemDistributionsDir string = "/var/cache/binenv"
)

// GetDefaultBinDir returns the bin directory
func GetDefaultBinDir() string {
	d, err := homedir.Dir()
	if err != nil {
		d = "~"
	}
	d += "/.binenv/"

	return d
}

// GetDefaultDistDir returns the default path for distribution
func GetDefaultDistDir() string {
	d, err := getDistDir()
	if err != nil {
		d = "~"
	}

	return d
}

// GetDefaultCacheDir returns the default path for distribution
func GetDefaultCacheDir() string {
	d, err := getCacheDir()
	if err != nil {
		d = "~"
	}

	return d
}

func getDistDir() (string, error) {
	var err error

	dir := os.Getenv("XDG_CONFIG_HOME")
	if dir == "" {
		dir, err = homedir.Dir()
		if err != nil {
			return "", err
		}
		dir += "/.config/binenv"
	}

	return dir, nil
}

func getCacheDir() (string, error) {
	dir := os.Getenv("XDG_CACHE_HOME")
	if dir == "" {
		dir, err := homedir.Dir()
		if err != nil {
			return "", err
		}
		dir += "/.cache/binenv"
	}

	return dir, nil
}

func stringInSlice(st string, sl []string) bool {
	for _, v := range sl {
		if v == st {
			return true
		}
	}

	return false
}
