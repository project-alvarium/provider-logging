package interfaces

import "github.com/project-alvarium/provider-logging/pkg/logging"

// Logger provides the abstraction through which logging capabilities can be obtained.
type Logger interface {
	// Write facilitates creation and writing of a LogEntry of a specified LogLevel. The client application
	// can also supply a message and a flexible list of additional arguments. These additional arguments are
	// optional. If provided, they should be treated as a key/value pair where the key is of type LogKey.
	//
	// Write flushes the LogEntry to StdOut in JSON format.
	Write(level logging.LogLevel, message string, args ...interface{})
	// Error facilitates creation and writing of a LogEntry at the Error LogLevel. The client application
	// can also supply a message and a flexible list of additional arguments. These additional arguments are
	// optional. If provided, they should be treated as a key/value pair where the key is of type LogKey.
	//
	// Write flushes the LogEntry to StdErr in JSON format.
	Error(message string, args ...interface{})
}
