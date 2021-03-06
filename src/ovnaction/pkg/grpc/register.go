// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2020 Intel Corporation

package grpc

import (
	"os"
	"strconv"
	"strings"

	log "github.com/open-ness/EMCO/src/orchestrator/pkg/infra/logutils"
)

const default_host = "localhost"
const default_port = 9032
const default_ovnaction_name = "ovnaction"
const ENV_OVNACTION_NAME = "OVNACTION_NAME"

func GetServerHostPort() (string, int) {

	// expect name of this ncm program to be in env variable "OVNACTION_NAME" - e.g. OVNACTION_NAME="ncm"
	serviceName := os.Getenv(ENV_OVNACTION_NAME)
	if serviceName == "" {
		serviceName = default_ovnaction_name
		log.Info("Using default name for OVNACTION service name", log.Fields{
			"Name": serviceName,
		})
	}

	// expect service name to be in env variable - e.g. OVNACTION_SERVICE_HOST
	host := os.Getenv(strings.ToUpper(serviceName) + "_SERVICE_HOST")
	if host == "" {
		host = default_host
		log.Info("Using default host for ovnaction gRPC controller", log.Fields{
			"Host": host,
		})
	}

	// expect service port to be in env variable - e.g. OVNACTION_SERVICE_PORT
	port, err := strconv.Atoi(os.Getenv(strings.ToUpper(serviceName) + "_SERVICE_PORT"))
	if err != nil || port < 0 {
		port = default_port
		log.Info("Using default port for ovnaction gRPC controller", log.Fields{
			"Port": port,
		})
	}
	return host, port
}
