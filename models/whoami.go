package models

import "log/slog"

type WhoAmIOptions struct {
	Args   []string
	Logger *slog.Logger
}
