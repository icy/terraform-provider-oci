// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

// InfrastructureTypeEnum Enum with underlying type: string
type InfrastructureTypeEnum string

// Set of constants representing the allowable values for InfrastructureTypeEnum
const (
	InfrastructureTypeOracleCloud     InfrastructureTypeEnum = "ORACLE_CLOUD"
	InfrastructureTypeCloudAtCustomer InfrastructureTypeEnum = "CLOUD_AT_CUSTOMER"
	InfrastructureTypeOnPremises      InfrastructureTypeEnum = "ON_PREMISES"
	InfrastructureTypeNonOracleCloud  InfrastructureTypeEnum = "NON_ORACLE_CLOUD"
)

var mappingInfrastructureTypeEnum = map[string]InfrastructureTypeEnum{
	"ORACLE_CLOUD":      InfrastructureTypeOracleCloud,
	"CLOUD_AT_CUSTOMER": InfrastructureTypeCloudAtCustomer,
	"ON_PREMISES":       InfrastructureTypeOnPremises,
	"NON_ORACLE_CLOUD":  InfrastructureTypeNonOracleCloud,
}

// GetInfrastructureTypeEnumValues Enumerates the set of values for InfrastructureTypeEnum
func GetInfrastructureTypeEnumValues() []InfrastructureTypeEnum {
	values := make([]InfrastructureTypeEnum, 0)
	for _, v := range mappingInfrastructureTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetInfrastructureTypeEnumStringValues Enumerates the set of values in String for InfrastructureTypeEnum
func GetInfrastructureTypeEnumStringValues() []string {
	return []string{
		"ORACLE_CLOUD",
		"CLOUD_AT_CUSTOMER",
		"ON_PREMISES",
		"NON_ORACLE_CLOUD",
	}
}
