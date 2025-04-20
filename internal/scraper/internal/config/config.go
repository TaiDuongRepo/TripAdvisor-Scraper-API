package config

import (
	"fmt"
	"os"
	"strings"
)

// Config is a struct that represents the configuration for the scraper
type Config struct {
	LocationURL string
	Languages   []string
	FileType    string
	ProxyHost   string
}

// NewConfig is a function that returns a new Config struct
// Returns an error if the LOCATION_URL is not set
func NewConfig(url string, langs string, file string) (*Config, error) {
	// Default languages
	defaultLanguages := []string{"en"}

	// Get location URL
	locationURL := url
	if locationURL == "" {
		return nil, fmt.Errorf("LOCATION_URL not set")
	}

	// Get languages
	languages := defaultLanguages
	if envLang := langs; envLang != "" {
		languages = strings.Split(envLang, "|")
	}

	// Get file type
	fileType := file
	if fileType == "" {
		fileType = "csv"
	}

	if fileType != "csv" && fileType != "json" {
		return nil, fmt.Errorf("invalid file type. Use csv or json")
	}

	// Get proxy host
	proxyHost := os.Getenv("PROXY_HOST")

	return &Config{
		LocationURL: locationURL,
		Languages:   languages,
		FileType:    fileType,
		ProxyHost:   proxyHost,
	}, nil
}
