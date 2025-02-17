// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Analytics API
//
// Analytics API.
//

package analytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v57/common"
	"strings"
)

// Capacity Service instance capacity metadata (e.g.: OLPU count, number of users, ...etc...).
type Capacity struct {

	// The capacity model to use.
	CapacityType CapacityTypeEnum `mandatory:"true" json:"capacityType"`

	// The capacity value selected (OLPU count, number of users, ...etc...). This parameter affects the
	// number of CPUs, amount of memory or other resources allocated to the instance.
	CapacityValue *int `mandatory:"true" json:"capacityValue"`
}

func (m Capacity) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Capacity) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := mappingCapacityTypeEnum[string(m.CapacityType)]; !ok && m.CapacityType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CapacityType: %s. Supported values are: %s.", m.CapacityType, strings.Join(GetCapacityTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
