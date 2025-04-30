package client

import (
	"net/http"
	"os"
	"time"
)

type JenkinsBasicAuthRoundTripper struct {
	Username string
	Password string

	RoundTripper http.RoundTripper
}

// RoundTrip executes a single HTTP transaction
func (j *JenkinsBasicAuthRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	// Set the basic auth header
	req.SetBasicAuth(j.Username, j.Password)

	// Use the provided RoundTripper or the default one
	if j.RoundTripper == nil {
		j.RoundTripper = http.DefaultTransport
	}

	// Perform the HTTP request
	return j.RoundTripper.RoundTrip(req)
}

// GetJenkinsClient returns a new Jenkins HTTP client
func GetJenkinsClient() (*http.Client, error) {
	c := &http.Client{
		Transport: &JenkinsBasicAuthRoundTripper{
			Username:     os.Getenv("JENKINS_USERNAME"),
			Password:     os.Getenv("JENKINS_PASSWORD"),
			RoundTripper: http.DefaultTransport,
		},
	}

	// Set a timeout for the HTTP client
	c.Timeout = 10 * time.Second

	return c, nil
}
