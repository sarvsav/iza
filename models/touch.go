package models

import "log/slog"

type TouchOptions struct {
	Args   []string
	Logger *slog.Logger
}
