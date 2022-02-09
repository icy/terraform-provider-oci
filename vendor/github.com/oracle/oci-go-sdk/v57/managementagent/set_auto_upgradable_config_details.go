// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Management Agent API
//
// API for Management Agent Cloud Service
//

package managementagent

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v57/common"
	"strings"
)

// SetAutoUpgradableConfigDetails Details for configuring tenancy-level agent AutoUpgradable configuration.
type SetAutoUpgradableConfigDetails struct {

	// Tenancy identifier i.e, Root compartment identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// true if the agents can be upgraded automatically; false if they must be upgraded manually.
	IsAgentAutoUpgradable *bool `mandatory:"true" json:"isAgentAutoUpgradable"`
}

func (m SetAutoUpgradableConfigDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SetAutoUpgradableConfigDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}