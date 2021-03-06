// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2020 Intel Corporation

package logutils

import (
	"github.com/open-ness/EMCO/src/orchestrator/pkg/infra/config"
	log "github.com/sirupsen/logrus"
	"strings"
)

//Fields is type that will be used by the calling function
type Fields map[string]interface{}

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})
	if strings.EqualFold(config.GetConfiguration().LogLevel, "warn") {
		log.SetLevel(log.WarnLevel)

	}
	if strings.EqualFold(config.GetConfiguration().LogLevel, "info") {
		log.SetLevel(log.InfoLevel)
	}
}

// Error uses the fields provided and logs
func Error(msg string, fields Fields) {
	log.WithFields(log.Fields(fields)).Error(msg)
}

// Warn uses the fields provided and logs
func Warn(msg string, fields Fields) {
	log.WithFields(log.Fields(fields)).Warn(msg)
}

// Info uses the fields provided and logs
func Info(msg string, fields Fields) {
	log.WithFields(log.Fields(fields)).Info(msg)
}
