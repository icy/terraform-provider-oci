// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Monitoring API
//
// Use the Monitoring API to manage metric queries and alarms for assessing the health, capacity, and performance of your cloud resources.
// Endpoints vary by operation. For PostMetric, use the `telemetry-ingestion` endpoints; for all other operations, use the `telemetry` endpoints.
// For information about monitoring, see Monitoring Overview (https://docs.cloud.oracle.com/iaas/Content/Monitoring/Concepts/monitoringoverview.htm).
//

package monitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v57/common"
	"strings"
)

// AlarmStatusSummary A summary of properties for the specified alarm and its current evaluation status.
// For information about alarms, see Alarms Overview (https://docs.cloud.oracle.com/iaas/Content/Monitoring/Concepts/monitoringoverview.htm#AlarmsOverview).
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/policygetstarted.htm).
// For information about endpoints and signing API requests, see
// About the API (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm). For information about available SDKs and tools, see
// SDKS and Other Tools (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/sdks.htm).
type AlarmStatusSummary struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the alarm.
	Id *string `mandatory:"true" json:"id"`

	// The configured name of the alarm.
	// Example: `High CPU Utilization`
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The configured severity of the alarm.
	// Example: `CRITICAL`
	Severity AlarmStatusSummarySeverityEnum `mandatory:"true" json:"severity"`

	// Timestamp for the transition of the alarm state. For example, the time when the alarm transitioned from OK to Firing.
	// Note: A three-minute lag for this value accounts for any late-arriving metrics.
	// Example: `2019-02-01T01:02:29.600Z`
	TimestampTriggered *common.SDKTime `mandatory:"true" json:"timestampTriggered"`

	// The status of this alarm.
	// Example: `FIRING`
	Status AlarmStatusSummaryStatusEnum `mandatory:"true" json:"status"`

	// The configuration details for suppressing an alarm.
	Suppression *Suppression `mandatory:"false" json:"suppression"`
}

func (m AlarmStatusSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AlarmStatusSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := mappingAlarmStatusSummarySeverityEnum[string(m.Severity)]; !ok && m.Severity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Severity: %s. Supported values are: %s.", m.Severity, strings.Join(GetAlarmStatusSummarySeverityEnumStringValues(), ",")))
	}
	if _, ok := mappingAlarmStatusSummaryStatusEnum[string(m.Status)]; !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetAlarmStatusSummaryStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AlarmStatusSummarySeverityEnum Enum with underlying type: string
type AlarmStatusSummarySeverityEnum string

// Set of constants representing the allowable values for AlarmStatusSummarySeverityEnum
const (
	AlarmStatusSummarySeverityCritical AlarmStatusSummarySeverityEnum = "CRITICAL"
	AlarmStatusSummarySeverityError    AlarmStatusSummarySeverityEnum = "ERROR"
	AlarmStatusSummarySeverityWarning  AlarmStatusSummarySeverityEnum = "WARNING"
	AlarmStatusSummarySeverityInfo     AlarmStatusSummarySeverityEnum = "INFO"
)

var mappingAlarmStatusSummarySeverityEnum = map[string]AlarmStatusSummarySeverityEnum{
	"CRITICAL": AlarmStatusSummarySeverityCritical,
	"ERROR":    AlarmStatusSummarySeverityError,
	"WARNING":  AlarmStatusSummarySeverityWarning,
	"INFO":     AlarmStatusSummarySeverityInfo,
}

// GetAlarmStatusSummarySeverityEnumValues Enumerates the set of values for AlarmStatusSummarySeverityEnum
func GetAlarmStatusSummarySeverityEnumValues() []AlarmStatusSummarySeverityEnum {
	values := make([]AlarmStatusSummarySeverityEnum, 0)
	for _, v := range mappingAlarmStatusSummarySeverityEnum {
		values = append(values, v)
	}
	return values
}

// GetAlarmStatusSummarySeverityEnumStringValues Enumerates the set of values in String for AlarmStatusSummarySeverityEnum
func GetAlarmStatusSummarySeverityEnumStringValues() []string {
	return []string{
		"CRITICAL",
		"ERROR",
		"WARNING",
		"INFO",
	}
}

// AlarmStatusSummaryStatusEnum Enum with underlying type: string
type AlarmStatusSummaryStatusEnum string

// Set of constants representing the allowable values for AlarmStatusSummaryStatusEnum
const (
	AlarmStatusSummaryStatusFiring    AlarmStatusSummaryStatusEnum = "FIRING"
	AlarmStatusSummaryStatusOk        AlarmStatusSummaryStatusEnum = "OK"
	AlarmStatusSummaryStatusSuspended AlarmStatusSummaryStatusEnum = "SUSPENDED"
)

var mappingAlarmStatusSummaryStatusEnum = map[string]AlarmStatusSummaryStatusEnum{
	"FIRING":    AlarmStatusSummaryStatusFiring,
	"OK":        AlarmStatusSummaryStatusOk,
	"SUSPENDED": AlarmStatusSummaryStatusSuspended,
}

// GetAlarmStatusSummaryStatusEnumValues Enumerates the set of values for AlarmStatusSummaryStatusEnum
func GetAlarmStatusSummaryStatusEnumValues() []AlarmStatusSummaryStatusEnum {
	values := make([]AlarmStatusSummaryStatusEnum, 0)
	for _, v := range mappingAlarmStatusSummaryStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetAlarmStatusSummaryStatusEnumStringValues Enumerates the set of values in String for AlarmStatusSummaryStatusEnum
func GetAlarmStatusSummaryStatusEnumStringValues() []string {
	return []string{
		"FIRING",
		"OK",
		"SUSPENDED",
	}
}
