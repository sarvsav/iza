package models

import "log/slog"

type CatOptions struct {
	Args   []string
	Logger *slog.Logger
}
