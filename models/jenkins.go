package models

import "time"

type JenkinsJob struct {
	Name         string
	Size         int64     // In bytes
	Perms        string    // Permissions (read/stop/start/create/configure)
	Owner        string    // Owner
	Group        string    // Group
	LastModified time.Time // Last run
}

type JenkinsJobFolder struct {
	Name         string
	Size         int64     // In bytes
	Perms        string    // Permissions (read/stop/start/create/configure)
	Owner        string    // Owner
	Group        string    // Group
	LastModified time.Time // Last run
}

type JenkinsLsResponse struct {
	JenkinsJobs       []JenkinsJob
	JenkinsJobFolders []JenkinsJobFolder
}

type JenkinsApiResponse []struct {
	Class                     string      `json:"_class"`
	Links                     Links       `json:"_links"`
	Actions0                  []any       `json:"actions"`
	Disabled                  bool        `json:"disabled"`
	DisplayName               string      `json:"displayName"`
	EstimatedDurationInMillis int         `json:"estimatedDurationInMillis,omitempty"`
	FullDisplayName           string      `json:"fullDisplayName"`
	FullName                  string      `json:"fullName"`
	LatestRun                 any         `json:"latestRun,omitempty"`
	Name                      string      `json:"name"`
	Organization              string      `json:"organization"`
	Parameters                any         `json:"parameters"`
	Permissions               Permissions `json:"permissions"`
	WeatherScore              int         `json:"weatherScore,omitempty"`
	NumberOfFolders           int         `json:"numberOfFolders,omitempty"`
	NumberOfPipelines         int         `json:"numberOfPipelines,omitempty"`
	PipelineFolderNames       []string    `json:"pipelineFolderNames,omitempty"`
	StartTime                 string      `json:"startTime,omitempty"`
}

type Self struct {
	Class string `json:"_class"`
	Href  string `json:"href"`
}

type Scm struct {
	Class string `json:"_class"`
	Href  string `json:"href"`
}

type Actions struct {
	Class string `json:"_class"`
	Href  string `json:"href"`
}

type Runs struct {
	Class string `json:"_class"`
	Href  string `json:"href"`
}

type Trends struct {
	Class string `json:"_class"`
	Href  string `json:"href"`
}

type Queue struct {
	Class string `json:"_class"`
	Href  string `json:"href"`
}

type Links struct {
	Self    Self    `json:"self"`
	Scm     Scm     `json:"scm"`
	Actions Actions `json:"actions"`
	Runs    Runs    `json:"runs"`
	Trends  Trends  `json:"trends"`
	Queue   Queue   `json:"queue"`
}

type Permissions struct {
	Read      bool `json:"read"`
	Stop      bool `json:"stop"`
	Start     bool `json:"start"`
	Create    bool `json:"create"`
	Configure bool `json:"configure"`
}
