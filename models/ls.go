package models

import "log/slog"

type LsOptions struct {
	LongListing bool
	Color       bool
	Args        []string
	Logger      *slog.Logger
}
