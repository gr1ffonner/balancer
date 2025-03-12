package utils

import (
	"fmt"
	"net/url"
	"strings"
)

func ParseURL(inputURL string) (string, string, error) {
	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		return "", "", err
	}

	s := strings.Split(parsedURL.Hostname(), ".")[0]
	videoPath := strings.TrimPrefix(parsedURL.Path, "/")

	return s, videoPath, nil
}

func ConstructCDNURL(CDNHost, sStorage, videoPath string) string {
	return fmt.Sprintf("http://%s/%s/%s", CDNHost, sStorage, videoPath)
}
