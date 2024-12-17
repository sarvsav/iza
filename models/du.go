package models

import "log/slog"

type DuOptions struct {
	Args   []string
	Logger *slog.Logger
}
