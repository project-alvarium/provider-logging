/*******************************************************************************
 * Copyright 2021 Dell Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *******************************************************************************/
package console

import (
	"encoding/json"
	"fmt"
	"github.com/project-alvarium/provider-logging/pkg/config"
	"github.com/project-alvarium/provider-logging/pkg/interfaces"
	"github.com/project-alvarium/provider-logging/pkg/logging"
	"os"
)

type ConsoleLogger struct {
	minLogLevel logging.LogLevel
	minSeverity int
	logLevels   []logging.LogLevel
}

func NewConsoleLogger(cfg config.LoggingInfo) interfaces.Logger {
	cl := ConsoleLogger{minLogLevel: cfg.MinLogLevel}
	//NOTE: The order of these array elements is important and is used to filter log messages that don't meet
	// the configured MinLogLevel. The elements are arranged from least to highest severity.
	cl.logLevels = []logging.LogLevel{logging.TraceLevel, logging.DebugLevel, logging.InfoLevel, logging.WarnLevel, logging.ErrorLevel}

	//Validate the LogLevel provided by configuration is one of the valid types.
	//If not, set to InfoLevel.
	if !cl.isValidLogLevel(cl.minLogLevel) {
		cl.minLogLevel = logging.InfoLevel
	}
	cl.minSeverity = cl.getSeverity(cl.minLogLevel)
	return &cl
}

func (cl ConsoleLogger) Write(level logging.LogLevel, message string, args ...interface{}) {
	//First, validate the supplied LogLevel. If invalid, the level will be set to DebugLevel.
	if !cl.isValidLogLevel(level) {
		//Further, if the MinLogLevel is higher than InfoLevel this message will be ignored
		level = logging.InfoLevel
	}

	severity := cl.getSeverity(level)
	if severity >= cl.minSeverity {
		entry := logging.NewLogEntry(level, message)
		cl.parseArgs(&entry, args)
		b, err := json.Marshal(&entry)
		if err != nil {
			cl.Error("Failed to marshal LogEntry: "+err.Error(), args)
			return
		}
		os.Stdout.WriteString(fmt.Sprintln(string(b)))
	}
}

func (cl ConsoleLogger) Error(message string, args ...interface{}) {
	// Error being the highest level of severity is always written, so no need to check.
	entry := logging.NewLogEntry(logging.ErrorLevel, message)
	cl.parseArgs(&entry, args)
	b, err := json.Marshal(&entry)
	if err != nil {
		os.Stderr.WriteString("ERROR: Failed to marshal LogEntry: " + err.Error())
		return
	}
	os.Stderr.WriteString(fmt.Sprintln(string(b)))
}

func (cl ConsoleLogger) isValidLogLevel(l logging.LogLevel) bool {
	for _, name := range cl.logLevels {
		if name == l {
			return true
		}
	}
	return false
}

func (cl ConsoleLogger) getSeverity(l logging.LogLevel) int {
	severity := -1

	for i, v := range cl.logLevels {
		if l == v {
			severity = i
			break
		}
	}

	return severity
}

func (cl ConsoleLogger) parseArgs(entry *logging.LogEntry, args []interface{}) {
	if args == nil {
		return
	}
	//If the number of args is 1, check to see if the key was supplied without a value.
	// If the first arg is not a key, append it to the LogEntry's Message property.
	if len(args) == 1 {
		if args[0] != logging.CorrelationKey {
			entry.Message = fmt.Sprintf("%s args=%s", entry.Message, args[0])
		}
		return
	}

	// This loop is meant to provide extension capability for other optional information we may want to log.
	// For now, everything except for the K/V dealing with CorrelationKey will be ignored. If you're passing
	// other stuff you're using the logger improperly. Make it part of your message.
	for i, arg := range args {
		if arg == logging.CorrelationKey {
			if i+1 < len(args) {
				entry.CorrelationId = fmt.Sprintf("%v", args[i+1])
			}
		}
	}
}
