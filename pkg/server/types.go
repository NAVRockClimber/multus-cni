// Copyright (c) 2022 Multus Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"net/http"

	"github.com/containernetworking/cni/pkg/invoke"

	"github.com/prometheus/client_golang/prometheus"

	"gopkg.in/k8snetworkplumbingwg/multus-cni.v4/pkg/k8sclient"
)

const (
	// const block for multus-daemon configs

	// DefaultMultusDaemonConfigFile is the Default path of the config file
	DefaultMultusDaemonConfigFile = "/etc/cni/net.d/multus.d/daemon-config.json"
	// DefaultMultusRunDir specifies default RunDir for multus
	DefaultMultusRunDir = "/run/multus/"
)

// Metrics represents server's metrics.
type Metrics struct {
	requestCounter *prometheus.CounterVec
}

// Server represents an HTTP server listening to a unix socket. It will handle
// the CNI shim requests issued when a pod is added / removed.
type Server struct {
	http.Server
	rundir       string
	kubeclient   *k8sclient.ClientInfo
	exec         invoke.Exec
	serverConfig []byte
	metrics      *Metrics
}

// ControllerNetConf for the controller cni configuration
type ControllerNetConf struct {
	ChrootDir   string `json:"chrootDir,omitempty"`
	LogFile     string `json:"logFile"`
	LogLevel    string `json:"logLevel"`
	LogToStderr bool   `json:"logToStderr,omitempty"`

	MetricsPort *int `json:"metricsPort,omitempty"`

	// Option to point to the path of the unix domain socket through which the
	// multus client / server communicate.
	SocketDir string `json:"socketDir"`

	ConfigFileContents []byte `json:"-"`
}
