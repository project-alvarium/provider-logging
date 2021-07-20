package factories

import (
	"github.com/project-alvarium/provider-logging/pkg/config"
	"github.com/project-alvarium/provider-logging/pkg/logging"
	"testing"
)

func TestWriteLogNoArgs(t *testing.T) {
	cfg := config.LoggingInfo{MinLogLevel: logging.InfoLevel}
	logger := NewLogger(cfg)
	logger.Write(logging.InfoLevel, "Test message with no args")
}

func TestWriteLogWithArgs(t *testing.T) {
	cfg := config.LoggingInfo{MinLogLevel: logging.InfoLevel}
	logger := NewLogger(cfg)
	logger.Write(logging.InfoLevel, "Test message with correlation", logging.CorrelationKey, "TEST_CORRELATION_ID")
}

func TestErrorNoArgs(t *testing.T) {
	cfg := config.LoggingInfo{MinLogLevel: logging.InfoLevel}
	logger := NewLogger(cfg)
	logger.Error("An error message with no args")
}

func TestErrorWithArgs(t *testing.T) {
	cfg := config.LoggingInfo{MinLogLevel: logging.InfoLevel}
	logger := NewLogger(cfg)
	logger.Error("An error message with correlation", logging.CorrelationKey, "TEST_CORRELATION_ID")
}
