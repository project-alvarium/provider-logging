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
package logging

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

type LogLevel string
type LogKey string

const (
	TraceLevel LogLevel = "trace"
	InfoLevel           = "info"
	DebugLevel          = "debug"
	WarnLevel           = "warn"
	ErrorLevel          = "error"

	CorrelationKey LogKey = "correlation-id"
)

type LogEntry struct {
	Timestamp     string   `json:"timestamp"`
	Hostname      string   `json:"hostname,omitempty"`
	Application   string   `json:"application,omitempty"`
	CorrelationId string   `json:"correlation-id,omitempty"`
	LineNumber    string   `json:"line-number,omitempty"`
	LogLevel      LogLevel `json:"log-level,omitempty"`
	Message       string   `json:"message,omitempty"`
}

func NewLogEntry(level LogLevel, message string) LogEntry {
	le := LogEntry{}
	le.Timestamp = time.Now().UTC().Format(time.RFC3339)
	le.Hostname, _ = os.Hostname()
	le.Application = os.Args[0]

	// Obtaining the line number in the same manner as the stdlib
	// https://github.com/golang/go/blob/54b251f542c97cf58a2ae800d3ed86cf14d0feed/src/log/log.go#L171
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "???"
		line = 0
	}
	le.LineNumber = fmt.Sprintf("%s:%v", filepath.Base(file), line)
	le.LogLevel = level
	le.Message = message

	return le
}
