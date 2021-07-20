package config

import "github.com/project-alvarium/provider-logging/pkg/logging"

// LoggingInfo provides properties for configuration the behavior of logging in your application
type LoggingInfo struct {
	MinLogLevel logging.LogLevel `json:"minLogLevel,omitempty"`
}
