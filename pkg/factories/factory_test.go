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
