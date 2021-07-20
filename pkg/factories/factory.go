package factories

import (
	"github.com/project-alvarium/provider-logging/internal/console"
	"github.com/project-alvarium/provider-logging/pkg/config"
	"github.com/project-alvarium/provider-logging/pkg/interfaces"
)

func NewLogger(cfg config.LoggingInfo) interfaces.Logger {
	return console.NewConsoleLogger(cfg)
}
