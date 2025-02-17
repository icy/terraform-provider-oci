// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

// MigrationTypesEnum Enum with underlying type: string
type MigrationTypesEnum string

// Set of constants representing the allowable values for MigrationTypesEnum
const (
	MigrationTypesOnline  MigrationTypesEnum = "ONLINE"
	MigrationTypesOffline MigrationTypesEnum = "OFFLINE"
)

var mappingMigrationTypesEnum = map[string]MigrationTypesEnum{
	"ONLINE":  MigrationTypesOnline,
	"OFFLINE": MigrationTypesOffline,
}

// GetMigrationTypesEnumValues Enumerates the set of values for MigrationTypesEnum
func GetMigrationTypesEnumValues() []MigrationTypesEnum {
	values := make([]MigrationTypesEnum, 0)
	for _, v := range mappingMigrationTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetMigrationTypesEnumStringValues Enumerates the set of values in String for MigrationTypesEnum
func GetMigrationTypesEnumStringValues() []string {
	return []string{
		"ONLINE",
		"OFFLINE",
	}
}
