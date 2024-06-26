/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package multizone

import (
	"github.com/pkg/errors"

	"go.uber.org/multierr"
)

import (
	"github.com/apache/dubbo-kubernetes/pkg/config"
	config_types "github.com/apache/dubbo-kubernetes/pkg/config/types"
)

type DdsServerConfig struct {
	config.BaseConfig

	// Port of a gRPC server that serves Dubbo Discovery Service (DDS).
	GrpcPort uint32 `json:"grpcPort" envconfig:"dubbo_multizone_global_dds_grpc_port"`
	// Interval for refreshing state of the world
	RefreshInterval config_types.Duration `json:"refreshInterval" envconfig:"dubbo_multizone_global_dds_refresh_interval"`
	// Interval for flushing Zone Insights (stats of multi-zone communication)
	ZoneInsightFlushInterval config_types.Duration `json:"zoneInsightFlushInterval" envconfig:"dubbo_multizone_global_dds_zone_insight_flush_interval"`
	// TlsEnabled turns on TLS for DDS
	TlsEnabled bool `json:"tlsEnabled" envconfig:"dubbo_multizone_global_dds_tls_enabled"`
	// TlsCertFile defines a path to a file with PEM-encoded TLS cert.
	TlsCertFile string `json:"tlsCertFile" envconfig:"dubbo_multizone_global_dds_tls_cert_file"`
	// TlsKeyFile defines a path to a file with PEM-encoded TLS key.
	TlsKeyFile string `json:"tlsKeyFile" envconfig:"dubbo_multizone_global_dds_tls_key_file"`
	// TlsMinVersion defines the minimum TLS version to be used
	TlsMinVersion string `json:"tlsMinVersion" envconfig:"dubbo_multizone_global_dds_tls_min_version"`
	// TlsMaxVersion defines the maximum TLS version to be used
	TlsMaxVersion string `json:"tlsMaxVersion" envconfig:"dubbo_multizone_global_dds_tls_max_version"`
	// TlsCipherSuites defines the list of ciphers to use
	TlsCipherSuites []string `json:"tlsCipherSuites" envconfig:"dubbo_multizone_global_dds_tls_cipher_suites"`
	// MaxMsgSize defines a maximum size of the message that is exchanged using DDS.
	// In practice this means a limit on full list of one resource type.
	MaxMsgSize uint32 `json:"maxMsgSize" envconfig:"dubbo_multizone_global_dds_max_msg_size"`
	// MsgSendTimeout defines a timeout on sending a single DDS message.
	// DDS stream between control planes is terminated if the control plane hits this timeout.
	MsgSendTimeout config_types.Duration `json:"msgSendTimeout" envconfig:"dubbo_multizone_global_dds_msg_send_timeout"`
	// Backoff that is executed when the global control plane is sending the response that was previously rejected by zone control plane.
	NackBackoff config_types.Duration `json:"nackBackoff" envconfig:"dubbo_multizone_global_dds_nack_backoff"`
	// DisableSOTW if true doesn't expose SOTW version of DDS. Default: false
	DisableSOTW bool `json:"disableSOTW" envconfig:"dubbo_multizone_global_dds_disable_sotw"`
	// ResponseBackoff is a time Global CP waits before sending ACK/NACK.
	// This is a way to slow down Zone CP from sending resources too often.
	ResponseBackoff config_types.Duration `json:"responseBackoff" envconfig:"dubbo_multizone_global_dds_response_backoff"`
	// ZoneHealthCheck holds config for ensuring zones are online
	ZoneHealthCheck ZoneHealthCheckConfig `json:"zoneHealthCheck"`
}

var _ config.Config = &DdsServerConfig{}

func (c *DdsServerConfig) PostProcess() error {
	return multierr.Combine(c.ZoneHealthCheck.PostProcess())
}

func (c *DdsServerConfig) Validate() error {
	var errs error
	if c.GrpcPort > 65535 {
		errs = multierr.Append(errs, errors.Errorf(".GrpcPort must be in the range [0, 65535]"))
	}
	if c.RefreshInterval.Duration <= 0 {
		errs = multierr.Append(errs, errors.New(".RefreshInterval must be positive"))
	}
	if c.ZoneInsightFlushInterval.Duration <= 0 {
		errs = multierr.Append(errs, errors.New(".ZoneInsightFlushInterval must be positive"))
	}
	if c.TlsCertFile == "" && c.TlsKeyFile != "" {
		errs = multierr.Append(errs, errors.New(".TlsCertFile cannot be empty if TlsKeyFile has been set"))
	}
	if c.TlsKeyFile == "" && c.TlsCertFile != "" {
		errs = multierr.Append(errs, errors.New(".TlsKeyFile cannot be empty if TlsCertFile has been set"))
	}
	if _, err := config_types.TLSVersion(c.TlsMinVersion); err != nil {
		errs = multierr.Append(errs, errors.New(".TlsMinVersion"+err.Error()))
	}
	if _, err := config_types.TLSVersion(c.TlsMaxVersion); err != nil {
		errs = multierr.Append(errs, errors.New(".TlsMaxVersion"+err.Error()))
	}
	if _, err := config_types.TLSCiphers(c.TlsCipherSuites); err != nil {
		errs = multierr.Append(errs, errors.New(".TlsCipherSuites"+err.Error()))
	}
	if err := c.ZoneHealthCheck.Validate(); err != nil {
		errs = multierr.Append(errs, errors.Wrap(err, "invalid zoneHealthCheck config"))
	}
	return errs
}

type DdsClientConfig struct {
	config.BaseConfig

	// Interval for refreshing state of the world
	RefreshInterval config_types.Duration `json:"refreshInterval" envconfig:"dubbo_multizone_zone_dds_refresh_interval"`
	// If true, TLS connection to the server won't be verified.
	TlsSkipVerify bool `json:"tlsSkipVerify" envconfig:"dubbo_multizone_zone_dds_tls_skip_verify"`
	// RootCAFile defines a path to a file with PEM-encoded Root CA. Client will verify the server by using it.
	RootCAFile string `json:"rootCaFile" envconfig:"dubbo_multizone_zone_dds_root_ca_file"`
	// MaxMsgSize defines a maximum size of the message that is exchanged using DDS.
	// In practice this means a limit on full list of one resource type.
	MaxMsgSize uint32 `json:"maxMsgSize" envconfig:"dubbo_multizone_zone_dds_max_msg_size"`
	// MsgSendTimeout defines a timeout on sending a single DDS message.
	// DDS stream between control planes is terminated if the control plane hits this timeout.
	MsgSendTimeout config_types.Duration `json:"msgSendTimeout" envconfig:"dubbo_multizone_zone_dds_msg_send_timeout"`
	// Backoff that is executed when the zone control plane is sending the response that was previously rejected by global control plane.
	NackBackoff config_types.Duration `json:"nackBackoff" envconfig:"dubbo_multizone_zone_dds_nack_backoff"`
	// ResponseBackoff is a time Zone CP waits before sending ACK/NACK.
	// This is a way to slow down Global CP from sending resources too often.
	ResponseBackoff config_types.Duration `json:"responseBackoff" envconfig:"dubbo_multizone_zone_dds_response_backoff"`
}

var _ config.Config = &DdsClientConfig{}

var _ config.Config = ZoneHealthCheckConfig{}

type ZoneHealthCheckConfig struct {
	config.BaseConfig

	// PollInterval is the interval between the global CP checking ZoneInsight for
	// health check pings and interval between zone CP sending health check pings
	PollInterval config_types.Duration `json:"pollInterval" envconfig:"dubbo_multizone_global_dds_zone_health_check_poll_interval"`
	// Timeout is the time after the last health check that a zone counts as
	// no longer online
	Timeout config_types.Duration `json:"timeout" envconfig:"dubbo_multizone_global_dds_zone_health_check_timeout"`
}

func (c ZoneHealthCheckConfig) Validate() error {
	if (c.Timeout.Duration > 0) != (c.PollInterval.Duration > 0) {
		return errors.New("timeout and pollInterval must both be either set or unset")
	}
	return nil
}
