// Package version records version information.
package version

import (
	"fmt"
	"strconv"
	"time"
)

// Variables replaced via -ldflags -X.
var (
	tag    string
	commit string
	date   string
	dirty  string
)

// Version records version information.
type Version struct {
	Version string    `json:"version"`
	Tag     string    `json:"tag"`
	Commit  string    `json:"commit"`
	Date    time.Time `json:"date"`
	Dirty   bool      `json:"dirty"`
}

func (v Version) String() string {
	return v.Version
}

// Get returns version information.
func Get() (v Version) {
	dt, e := strconv.ParseInt(date, 10, 64)
	if e != nil || len(commit) != 40 {
		v.Tag = "0.0.0"
		v.Version = "development"
		v.Commit = "unknown"
		v.Date = time.Now()
		v.Dirty = true
		return
	}

	v.Tag = tag
	v.Commit = commit
	v.Date = time.Unix(dt, 0)
	v.Dirty = dirty != ""
	dirtySuffix := ""
	if v.Dirty {
		dirtySuffix = "-dirty"
	}
	if v.Tag == "" {
		v.Tag = "0.0.0"
	}
	v.Version = fmt.Sprintf("%s-%s-%s%s", v.Tag, v.Date.Format("20060102150405"), commit[:12], dirtySuffix)
	return
}
