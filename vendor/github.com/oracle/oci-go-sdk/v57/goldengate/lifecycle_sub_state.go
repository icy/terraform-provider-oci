// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

// LifecycleSubStateEnum Enum with underlying type: string
type LifecycleSubStateEnum string

// Set of constants representing the allowable values for LifecycleSubStateEnum
const (
	LifecycleSubStateRecovering       LifecycleSubStateEnum = "RECOVERING"
	LifecycleSubStateStarting         LifecycleSubStateEnum = "STARTING"
	LifecycleSubStateStopping         LifecycleSubStateEnum = "STOPPING"
	LifecycleSubStateMoving           LifecycleSubStateEnum = "MOVING"
	LifecycleSubStateUpgrading        LifecycleSubStateEnum = "UPGRADING"
	LifecycleSubStateRestoring        LifecycleSubStateEnum = "RESTORING"
	LifecycleSubStateBackupInProgress LifecycleSubStateEnum = "BACKUP_IN_PROGRESS"
)

var mappingLifecycleSubStateEnum = map[string]LifecycleSubStateEnum{
	"RECOVERING":         LifecycleSubStateRecovering,
	"STARTING":           LifecycleSubStateStarting,
	"STOPPING":           LifecycleSubStateStopping,
	"MOVING":             LifecycleSubStateMoving,
	"UPGRADING":          LifecycleSubStateUpgrading,
	"RESTORING":          LifecycleSubStateRestoring,
	"BACKUP_IN_PROGRESS": LifecycleSubStateBackupInProgress,
}

// GetLifecycleSubStateEnumValues Enumerates the set of values for LifecycleSubStateEnum
func GetLifecycleSubStateEnumValues() []LifecycleSubStateEnum {
	values := make([]LifecycleSubStateEnum, 0)
	for _, v := range mappingLifecycleSubStateEnum {
		values = append(values, v)
	}
	return values
}

// GetLifecycleSubStateEnumStringValues Enumerates the set of values in String for LifecycleSubStateEnum
func GetLifecycleSubStateEnumStringValues() []string {
	return []string{
		"RECOVERING",
		"STARTING",
		"STOPPING",
		"MOVING",
		"UPGRADING",
		"RESTORING",
		"BACKUP_IN_PROGRESS",
	}
}
