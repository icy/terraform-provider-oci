// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud VMware Solution API
//
// Use the Oracle Cloud VMware API to create SDDCs and manage ESXi hosts and software.
// For more information, see Oracle Cloud VMware Solution (https://docs.cloud.oracle.com/iaas/Content/VMware/Concepts/ocvsoverview.htm).
//

package ocvp

// SortOrdersEnum Enum with underlying type: string
type SortOrdersEnum string

// Set of constants representing the allowable values for SortOrdersEnum
const (
	SortOrdersAsc  SortOrdersEnum = "ASC"
	SortOrdersDesc SortOrdersEnum = "DESC"
)

var mappingSortOrdersEnum = map[string]SortOrdersEnum{
	"ASC":  SortOrdersAsc,
	"DESC": SortOrdersDesc,
}

// GetSortOrdersEnumValues Enumerates the set of values for SortOrdersEnum
func GetSortOrdersEnumValues() []SortOrdersEnum {
	values := make([]SortOrdersEnum, 0)
	for _, v := range mappingSortOrdersEnum {
		values = append(values, v)
	}
	return values
}

// GetSortOrdersEnumStringValues Enumerates the set of values in String for SortOrdersEnum
func GetSortOrdersEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}
