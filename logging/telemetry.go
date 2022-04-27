// Copyright (C) 2019-2022 Algorand, Inc.
// This file is part of go-algorand
//
// go-algorand is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// go-algorand is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with go-algorand.  If not, see <https://www.gnu.org/licenses/>.

package logging

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"

	"github.com/Orca18/novarand/config"
	"github.com/Orca18/novarand/logging/telemetryspec"
)

const telemetryPrefix = "/"
const telemetrySeparator = "/"
const logBufferDepth = 2

// EnableTelemetry configures and enables telemetry based on the config provided
func EnableTelemetry(cfg TelemetryConfig, l *logger) (err error) {
	telemetry, err := makeTelemetryState(cfg, createElasticHook)
	if err != nil {
		return
	}
	enableTelemetryState(telemetry, l)
	return
}

func enableTelemetryState(telemetry *telemetryState, l *logger) {
	l.loggerState.telemetry = telemetry
	// Hook our normal logging to send desired types to telemetry
	l.AddHook(telemetry.hook)
	// Wrap current logger Output writer to capture history
	l.setOutput(telemetry.wrapOutput(l.getOutput()))
}

func makeLevels(min logrus.Level) []logrus.Level {
	levels := []logrus.Level{}
	for _, l := range []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
	} {
		if l <= min {
			levels = append(levels, l)
		}
	}
	return levels
}

func makeTelemetryState(cfg TelemetryConfig, hookFactory hookFactory) (*telemetryState, error) {
	telemetry := &telemetryState{}
	telemetry.history = createLogBuffer(logBufferDepth)
	if cfg.Enable {
		if cfg.SessionGUID == "" {
			cfg.SessionGUID = uuid.NewV4().String()
		}
		hook, err := createTelemetryHook(cfg, telemetry.history, hookFactory)
		if err != nil {
			return nil, err
		}
		telemetry.hook = createAsyncHookLevels(hook, 32, 100, makeLevels(cfg.MinLogLevel))
	} else {
		telemetry.hook = new(dummyHook)
	}
	telemetry.telemetryConfig = cfg
	return telemetry, nil
}

// ReadTelemetryConfigOrDefault reads telemetry config from file or defaults if no config file found.
func ReadTelemetryConfigOrDefault(dataDir string, genesisID string) (cfg TelemetryConfig, err error) {
	err = nil
	dataDirProvided := dataDir != ""
	var configPath string

	// If we have a data directory, then load the config
	if dataDirProvided {
		configPath = filepath.Join(dataDir, TelemetryConfigFilename)
		// Load the config, if the GUID is there then we are all set
		// However if it isn't there then we must create it, save the file and load it.
		cfg, err = LoadTelemetryConfig(configPath)
	}

	// We couldn't load the telemetry config for some reason
	// If the reason is because the directory doesn't exist or we didn't provide a data directory then...
	if (err != nil && os.IsNotExist(err)) || !dataDirProvided {

		configPath, err = config.GetConfigFilePath(TelemetryConfigFilename)
		if err != nil {
			// In this case we don't know what to do since we couldn't
			// create the directory.  Just create an ephemeral config.
			cfg = createTelemetryConfig()
			return
		}

		// Load the telemetry from the default config path
		cfg, err = LoadTelemetryConfig(configPath)
	}

	// If there was some error loading the configuration from the config path...
	if err != nil {
		// Create an ephemeral config
		cfg = createTelemetryConfig()

		// If the error was that the the config wasn't there then it wasn't really an error
		if os.IsNotExist(err) {
			err = nil
		} else {
			// The error was actually due to a malformed config file...just return
			return
		}
	}
	ver := config.GetCurrentVersion()
	ch := ver.Channel
	// Should not happen, but default to "dev" if channel is unspecified.
	if ch == "" {
		ch = "dev"
	}
	cfg.ChainID = fmt.Sprintf("%s-%s", ch, genesisID)
	cfg.Version = ver.String()
	return cfg, err
}

// EnsureTelemetryConfig creates a new TelemetryConfig structure with a generated GUID and the appropriate Telemetry endpoint
// Err will be non-nil if the file doesn't exist, or if error loading.
// Cfg will always be valid.
func EnsureTelemetryConfig(dataDir *string, genesisID string) (TelemetryConfig, error) {
	cfg, _, err := EnsureTelemetryConfigCreated(dataDir, genesisID)
	return cfg, err
}

// EnsureTelemetryConfigCreated is the same as EnsureTelemetryConfig but it also returns a bool indicating
// whether EnsureTelemetryConfig had to create the config.
func EnsureTelemetryConfigCreated(dataDir *string, genesisID string) (TelemetryConfig, bool, error) {
	/*
		Our logic should be as follows:
			- We first look inside the provided data-directory.  If a config file is there, load it
			  and return it
			- Otherwise, look in the global directory.  If a config file is there, load it and return it.
			- Otherwise, if a data-directory was provided then save the config file there.
			- Otherwise, save the config file in the global directory

	*/

	configPath := ""
	var cfg TelemetryConfig
	var err error

	if dataDir != nil && *dataDir != "" {
		configPath = filepath.Join(*dataDir, TelemetryConfigFilename)
		cfg, err = LoadTelemetryConfig(configPath)
		if err != nil && os.IsNotExist(err) {
			// if it just didn't exist, try again at the other path
			configPath = ""
		}
	}
	if configPath == "" {
		configPath, err = config.GetConfigFilePath(TelemetryConfigFilename)
		if err != nil {
			cfg := createTelemetryConfig()
			// Since GetConfigFilePath failed, there is no chance that we
			// can save the next config files
			return cfg, true, err
		}
		cfg, err = LoadTelemetryConfig(configPath)
	}
	created := false
	if err != nil {
		created = true
		cfg = createTelemetryConfig()

		if dataDir != nil && *dataDir != "" {

			/*
				There could be a scenario where a data directory was supplied that doesn't exist.
				In that case, we don't want to create the directory, just save in the global one
			*/

			// If the directory exists...
			if _, err := os.Stat(*dataDir); err == nil {

				// Remember, if we had a data directory supplied we want to save the config there
				configPath = filepath.Join(*dataDir, TelemetryConfigFilename)
			}

		}

		cfg.FilePath = configPath // Initialize our desired cfg.FilePath

		// There was no config file, create it.
		err = cfg.Save(configPath)
	}

	ch := config.GetCurrentVersion().Channel
	// Should not happen, but default to "dev" if channel is unspecified.
	if ch == "" {
		ch = "dev"
	}
	cfg.ChainID = fmt.Sprintf("%s-%s", ch, genesisID)

	return cfg, created, err
}

// wrapOutput wraps the log writer so we can keep a history of
// the tail of the file to send with critical telemetry events when logged.
func (t *telemetryState) wrapOutput(out io.Writer) io.Writer {
	return t.history.wrapOutput(out)
}

func (t *telemetryState) logMetrics(l logger, category telemetryspec.Category, metrics telemetryspec.MetricDetails, details interface{}) {
	if metrics == nil {
		return
	}
	l = l.WithFields(logrus.Fields{
		"metrics": metrics,
	}).(logger)

	t.logTelemetry(l, buildMessage(string(category), string(metrics.Identifier())), details)
}

func (t *telemetryState) logEvent(l logger, category telemetryspec.Category, identifier telemetryspec.Event, details interface{}) {
	t.logTelemetry(l, buildMessage(string(category), string(identifier)), details)
}

func (t *telemetryState) logStartOperation(l logger, category telemetryspec.Category, identifier telemetryspec.Operation) TelemetryOperation {
	op := makeTelemetryOperation(t, category, identifier)
	t.logTelemetry(l, buildMessage(string(category), string(identifier), "Start"), nil)
	return op
}

func buildMessage(args ...string) string {
	message := telemetryPrefix + strings.Join(args, telemetrySeparator)
	return message
}

// logTelemetry explicitly only sends telemetry events to the cloud.
func (t *telemetryState) logTelemetry(l logger, message string, details interface{}) {
	if details != nil {
		l = l.WithFields(logrus.Fields{
			"details": details,
		}).(logger)
	}

	entry := l.entry.WithFields(Fields{
		"session":      l.GetTelemetrySession(),
		"instanceName": l.GetInstanceName(),
		"v":            l.GetTelemetryVersion(),
	})
	// Populate entry like logrus.entry.log() does
	entry.Time = time.Now()
	entry.Level = logrus.InfoLevel
	entry.Message = message

	if t.telemetryConfig.SendToLog {
		entry.Info(message)
	}
	t.hook.Fire(entry)
}

func (t *telemetryState) Close() {
	if t.hook != nil {
		t.hook.Close()
	}
}

func (t *telemetryState) Flush() {
	t.hook.Flush()
}
