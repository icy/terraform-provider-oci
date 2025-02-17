// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v57/common"
	"strings"
)

// TablespaceAdminStatus The status of a tablespace admin action.
type TablespaceAdminStatus struct {

	// The status of a tablespace admin action.
	Status TablespaceAdminStatusStatusEnum `mandatory:"true" json:"status"`

	// The error code that denotes failure if the tablespace admin action is not successful. The error code is "null" if the admin action is successful.
	ErrorCode *int `mandatory:"false" json:"errorCode"`

	// The error message that indicates the reason for failure if the tablespace admin action is not successful. The error message is "null" if the admin action is successful.
	ErrorMessage *string `mandatory:"false" json:"errorMessage"`
}

func (m TablespaceAdminStatus) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TablespaceAdminStatus) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := mappingTablespaceAdminStatusStatusEnum[string(m.Status)]; !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetTablespaceAdminStatusStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TablespaceAdminStatusStatusEnum Enum with underlying type: string
type TablespaceAdminStatusStatusEnum string

// Set of constants representing the allowable values for TablespaceAdminStatusStatusEnum
const (
	TablespaceAdminStatusStatusSucceeded TablespaceAdminStatusStatusEnum = "SUCCEEDED"
	TablespaceAdminStatusStatusFailed    TablespaceAdminStatusStatusEnum = "FAILED"
)

var mappingTablespaceAdminStatusStatusEnum = map[string]TablespaceAdminStatusStatusEnum{
	"SUCCEEDED": TablespaceAdminStatusStatusSucceeded,
	"FAILED":    TablespaceAdminStatusStatusFailed,
}

// GetTablespaceAdminStatusStatusEnumValues Enumerates the set of values for TablespaceAdminStatusStatusEnum
func GetTablespaceAdminStatusStatusEnumValues() []TablespaceAdminStatusStatusEnum {
	values := make([]TablespaceAdminStatusStatusEnum, 0)
	for _, v := range mappingTablespaceAdminStatusStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetTablespaceAdminStatusStatusEnumStringValues Enumerates the set of values in String for TablespaceAdminStatusStatusEnum
func GetTablespaceAdminStatusStatusEnumStringValues() []string {
	return []string{
		"SUCCEEDED",
		"FAILED",
	}
}
