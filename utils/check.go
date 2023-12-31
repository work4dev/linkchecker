package utils

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/work4dev/linkchecker/config"
)

// CheckLink Checks if the link is valid.
func CheckLink(url string, timeout time.Duration) (bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return false, err.Error()
	}
	if config.UA != "" {
		req.Header.Set("User-Agent", config.UA)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err.Error()
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return false, fmt.Sprintf("status code: %d", resp.StatusCode)
	}

	return true, "ok"
}
