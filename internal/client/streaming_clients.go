// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_streaming "github.com/oracle/oci-go-sdk/v57/streaming"

	oci_common "github.com/oracle/oci-go-sdk/v57/common"
)

func init() {
	RegisterOracleClient("oci_streaming.StreamAdminClient", &OracleClient{InitClientFn: initStreamingStreamAdminClient})
}

func initStreamingStreamAdminClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_streaming.NewStreamAdminClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}

	if serviceClientOverrides.HostUrlOverride != "" {
		client.Host = serviceClientOverrides.HostUrlOverride
	}
	return &client, nil
}

func (m *OracleClients) StreamAdminClient() *oci_streaming.StreamAdminClient {
	return m.GetClient("oci_streaming.StreamAdminClient").(*oci_streaming.StreamAdminClient)
}
