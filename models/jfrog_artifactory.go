package models

import "time"

type JFrogArtifactoryJob struct {
	Name         string
	Size         int64     // In bytes
	Perms        string    // Permissions (read/stop/start/create/configure)
	Owner        string    // Owner
	Group        string    // Group
	LastModified time.Time // Last run
}

type JFrogArtifactoryApiResponse struct {
	Repo         string     `json:"repo"`
	Path         string     `json:"path"`
	Created      time.Time  `json:"created"`
	LastModified time.Time  `json:"lastModified"`
	LastUpdated  time.Time  `json:"lastUpdated"`
	Children     []Children `json:"children"`
	URI          string     `json:"uri"`
	Errors       []Errors   `json:"errors"`
}

type Children struct {
	URI    string `json:"uri"`
	Folder bool   `json:"folder"`
}

type Errors struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type JFrogResult struct {
	JFrogLsResponse ArtifactoryLsResponseData
}

func (jfr JFrogResult) isArtifactoryLsResponse() {
	// marker function
}

func (jfr JFrogResult) GetLsResult() (ArtifactoryLsResponseData, error) {
	return jfr.JFrogLsResponse, nil
}
