package client

import (
	"net/http"
	"os"
	"time"
)

type JFrogBasicAuthRoundTripper struct {
	Username string
	Password string

	RoundTripper http.RoundTripper
}

// RoundTrip executes a single HTTP transaction
func (j *JFrogBasicAuthRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	// Set the basic auth header
	req.SetBasicAuth(j.Username, j.Password)

	// Use the provided RoundTripper or the default one
	if j.RoundTripper == nil {
		j.RoundTripper = http.DefaultTransport
	}

	// Perform the HTTP request
	return j.RoundTripper.RoundTrip(req)
}

// GetJFrogClient returns a new Jenkins HTTP client
func GetJFrogClient() (*http.Client, error) {
	c := &http.Client{
		Transport: &JFrogBasicAuthRoundTripper{
			Username:     os.Getenv("JENKINS_USERNAME"),
			Password:     os.Getenv("JENKINS_PASSWORD"),
			RoundTripper: http.DefaultTransport,
		},
	}

	// Set a timeout for the HTTP client
	c.Timeout = 10 * time.Second

	return c, nil
}
