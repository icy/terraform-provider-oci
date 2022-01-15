// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools
//
// Database Tools APIs to manage Connections and Private Endpoints.
//

package databasetools

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v55/common"
)

// DatabaseToolsKeyStoreContentSecretId The key store content.
type DatabaseToolsKeyStoreContentSecretId struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the secret containing the key store.
	SecretId *string `mandatory:"false" json:"secretId"`
}

func (m DatabaseToolsKeyStoreContentSecretId) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m DatabaseToolsKeyStoreContentSecretId) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDatabaseToolsKeyStoreContentSecretId DatabaseToolsKeyStoreContentSecretId
	s := struct {
		DiscriminatorParam string `json:"valueType"`
		MarshalTypeDatabaseToolsKeyStoreContentSecretId
	}{
		"SECRETID",
		(MarshalTypeDatabaseToolsKeyStoreContentSecretId)(m),
	}

	return json.Marshal(&s)
}