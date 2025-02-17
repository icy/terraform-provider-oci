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

// RoleSummary A summary of each role.
type RoleSummary struct {

	// The name of the role granted to the user.
	Name *string `mandatory:"false" json:"name"`

	// Indicates whether the role is granted with the ADMIN OPTION (YES) or not (NO).
	AdminOption RoleSummaryAdminOptionEnum `mandatory:"false" json:"adminOption,omitempty"`

	// Indicates whether the role is granted with the DELEGATE OPTION (YES) or not (NO).
	DelegateOption RoleSummaryDelegateOptionEnum `mandatory:"false" json:"delegateOption,omitempty"`

	// Indicates whether the role is designated as a DEFAULT ROLE for the user (YES) or not (NO).
	DefaultRole RoleSummaryDefaultRoleEnum `mandatory:"false" json:"defaultRole,omitempty"`

	// Indicates how the role was granted. Possible values:
	// YES if the role is granted commonly (CONTAINER=ALL is used)
	// NO if the role is granted locally (CONTAINER=ALL is not used)
	Common RoleSummaryCommonEnum `mandatory:"false" json:"common,omitempty"`

	// Indicates whether the granted role is inherited from another container (YES) or not (NO).
	Inherited RoleSummaryInheritedEnum `mandatory:"false" json:"inherited,omitempty"`
}

func (m RoleSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RoleSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := mappingRoleSummaryAdminOptionEnum[string(m.AdminOption)]; !ok && m.AdminOption != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AdminOption: %s. Supported values are: %s.", m.AdminOption, strings.Join(GetRoleSummaryAdminOptionEnumStringValues(), ",")))
	}
	if _, ok := mappingRoleSummaryDelegateOptionEnum[string(m.DelegateOption)]; !ok && m.DelegateOption != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DelegateOption: %s. Supported values are: %s.", m.DelegateOption, strings.Join(GetRoleSummaryDelegateOptionEnumStringValues(), ",")))
	}
	if _, ok := mappingRoleSummaryDefaultRoleEnum[string(m.DefaultRole)]; !ok && m.DefaultRole != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DefaultRole: %s. Supported values are: %s.", m.DefaultRole, strings.Join(GetRoleSummaryDefaultRoleEnumStringValues(), ",")))
	}
	if _, ok := mappingRoleSummaryCommonEnum[string(m.Common)]; !ok && m.Common != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Common: %s. Supported values are: %s.", m.Common, strings.Join(GetRoleSummaryCommonEnumStringValues(), ",")))
	}
	if _, ok := mappingRoleSummaryInheritedEnum[string(m.Inherited)]; !ok && m.Inherited != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Inherited: %s. Supported values are: %s.", m.Inherited, strings.Join(GetRoleSummaryInheritedEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RoleSummaryAdminOptionEnum Enum with underlying type: string
type RoleSummaryAdminOptionEnum string

// Set of constants representing the allowable values for RoleSummaryAdminOptionEnum
const (
	RoleSummaryAdminOptionYes RoleSummaryAdminOptionEnum = "YES"
	RoleSummaryAdminOptionNo  RoleSummaryAdminOptionEnum = "NO"
)

var mappingRoleSummaryAdminOptionEnum = map[string]RoleSummaryAdminOptionEnum{
	"YES": RoleSummaryAdminOptionYes,
	"NO":  RoleSummaryAdminOptionNo,
}

// GetRoleSummaryAdminOptionEnumValues Enumerates the set of values for RoleSummaryAdminOptionEnum
func GetRoleSummaryAdminOptionEnumValues() []RoleSummaryAdminOptionEnum {
	values := make([]RoleSummaryAdminOptionEnum, 0)
	for _, v := range mappingRoleSummaryAdminOptionEnum {
		values = append(values, v)
	}
	return values
}

// GetRoleSummaryAdminOptionEnumStringValues Enumerates the set of values in String for RoleSummaryAdminOptionEnum
func GetRoleSummaryAdminOptionEnumStringValues() []string {
	return []string{
		"YES",
		"NO",
	}
}

// RoleSummaryDelegateOptionEnum Enum with underlying type: string
type RoleSummaryDelegateOptionEnum string

// Set of constants representing the allowable values for RoleSummaryDelegateOptionEnum
const (
	RoleSummaryDelegateOptionYes RoleSummaryDelegateOptionEnum = "YES"
	RoleSummaryDelegateOptionNo  RoleSummaryDelegateOptionEnum = "NO"
)

var mappingRoleSummaryDelegateOptionEnum = map[string]RoleSummaryDelegateOptionEnum{
	"YES": RoleSummaryDelegateOptionYes,
	"NO":  RoleSummaryDelegateOptionNo,
}

// GetRoleSummaryDelegateOptionEnumValues Enumerates the set of values for RoleSummaryDelegateOptionEnum
func GetRoleSummaryDelegateOptionEnumValues() []RoleSummaryDelegateOptionEnum {
	values := make([]RoleSummaryDelegateOptionEnum, 0)
	for _, v := range mappingRoleSummaryDelegateOptionEnum {
		values = append(values, v)
	}
	return values
}

// GetRoleSummaryDelegateOptionEnumStringValues Enumerates the set of values in String for RoleSummaryDelegateOptionEnum
func GetRoleSummaryDelegateOptionEnumStringValues() []string {
	return []string{
		"YES",
		"NO",
	}
}

// RoleSummaryDefaultRoleEnum Enum with underlying type: string
type RoleSummaryDefaultRoleEnum string

// Set of constants representing the allowable values for RoleSummaryDefaultRoleEnum
const (
	RoleSummaryDefaultRoleYes RoleSummaryDefaultRoleEnum = "YES"
	RoleSummaryDefaultRoleNo  RoleSummaryDefaultRoleEnum = "NO"
)

var mappingRoleSummaryDefaultRoleEnum = map[string]RoleSummaryDefaultRoleEnum{
	"YES": RoleSummaryDefaultRoleYes,
	"NO":  RoleSummaryDefaultRoleNo,
}

// GetRoleSummaryDefaultRoleEnumValues Enumerates the set of values for RoleSummaryDefaultRoleEnum
func GetRoleSummaryDefaultRoleEnumValues() []RoleSummaryDefaultRoleEnum {
	values := make([]RoleSummaryDefaultRoleEnum, 0)
	for _, v := range mappingRoleSummaryDefaultRoleEnum {
		values = append(values, v)
	}
	return values
}

// GetRoleSummaryDefaultRoleEnumStringValues Enumerates the set of values in String for RoleSummaryDefaultRoleEnum
func GetRoleSummaryDefaultRoleEnumStringValues() []string {
	return []string{
		"YES",
		"NO",
	}
}

// RoleSummaryCommonEnum Enum with underlying type: string
type RoleSummaryCommonEnum string

// Set of constants representing the allowable values for RoleSummaryCommonEnum
const (
	RoleSummaryCommonYes RoleSummaryCommonEnum = "YES"
	RoleSummaryCommonNo  RoleSummaryCommonEnum = "NO"
)

var mappingRoleSummaryCommonEnum = map[string]RoleSummaryCommonEnum{
	"YES": RoleSummaryCommonYes,
	"NO":  RoleSummaryCommonNo,
}

// GetRoleSummaryCommonEnumValues Enumerates the set of values for RoleSummaryCommonEnum
func GetRoleSummaryCommonEnumValues() []RoleSummaryCommonEnum {
	values := make([]RoleSummaryCommonEnum, 0)
	for _, v := range mappingRoleSummaryCommonEnum {
		values = append(values, v)
	}
	return values
}

// GetRoleSummaryCommonEnumStringValues Enumerates the set of values in String for RoleSummaryCommonEnum
func GetRoleSummaryCommonEnumStringValues() []string {
	return []string{
		"YES",
		"NO",
	}
}

// RoleSummaryInheritedEnum Enum with underlying type: string
type RoleSummaryInheritedEnum string

// Set of constants representing the allowable values for RoleSummaryInheritedEnum
const (
	RoleSummaryInheritedYes RoleSummaryInheritedEnum = "YES"
	RoleSummaryInheritedNo  RoleSummaryInheritedEnum = "NO"
)

var mappingRoleSummaryInheritedEnum = map[string]RoleSummaryInheritedEnum{
	"YES": RoleSummaryInheritedYes,
	"NO":  RoleSummaryInheritedNo,
}

// GetRoleSummaryInheritedEnumValues Enumerates the set of values for RoleSummaryInheritedEnum
func GetRoleSummaryInheritedEnumValues() []RoleSummaryInheritedEnum {
	values := make([]RoleSummaryInheritedEnum, 0)
	for _, v := range mappingRoleSummaryInheritedEnum {
		values = append(values, v)
	}
	return values
}

// GetRoleSummaryInheritedEnumStringValues Enumerates the set of values in String for RoleSummaryInheritedEnum
func GetRoleSummaryInheritedEnumStringValues() []string {
	return []string{
		"YES",
		"NO",
	}
}
