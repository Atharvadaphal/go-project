package checker

import (
	"fmt"
	"net/http"
	"time"
)

// Result holds the status of a URL
type Result struct {
	URL    string `json:"url"`
	Status string `json:"status"`
	Up     bool   `json:"up"`
}

// Config holds settings for our HTTP checks.
// We will pass this as a pointer to avoid duplicating data in memory.
type Config struct {
	Timeout time.Duration
}

// Implement the Stringer interface for custom terminal output
func (r Result) String() string {
	if r.Up {
		return fmt.Sprintf("[UP] %s", r.URL)
	}
	return fmt.Sprintf("[DOWN] %s - %s", r.URL, r.Status)
}

// CheckURL performs the HTTP GET request.
// Notice the *Config - this accepts a pointer rather than a copied value.
func CheckURL(url string, cfg *Config, ch chan<- Result) {
	client := http.Client{Timeout: cfg.Timeout}
	resp, err := client.Get(url)

	if err != nil {
		ch <- Result{URL: url, Status: err.Error(), Up: false}
		return
	}
	defer resp.Body.Close()

	ch <- Result{URL: url, Status: resp.Status, Up: resp.StatusCode == 200}
}
