// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v57/common"
	"strings"
)

// IamWorkRequest A IAM work request object that allows users to track Asynchronous API status.
type IamWorkRequest struct {

	// The OCID of the work request.
	Id *string `mandatory:"true" json:"id"`

	// The asynchronous operation tracked by this IAM work request.
	OperationType IamWorkRequestOperationTypeEnum `mandatory:"true" json:"operationType"`

	// Status of the work request
	Status IamWorkRequestStatusEnum `mandatory:"true" json:"status"`

	// The OCID of the compartment containing this IAM work request.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The resources this work request affects.
	Resources []IamWorkRequestResource `mandatory:"false" json:"resources"`

	// How much progress the operation has made.
	PercentComplete *float32 `mandatory:"false" json:"percentComplete"`

	// Date and time the work was accepted, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeAccepted *common.SDKTime `mandatory:"false" json:"timeAccepted"`

	// Date and time the work started, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// Date and time the work completed, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`
}

func (m IamWorkRequest) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m IamWorkRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := mappingIamWorkRequestOperationTypeEnum[string(m.OperationType)]; !ok && m.OperationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationType: %s. Supported values are: %s.", m.OperationType, strings.Join(GetIamWorkRequestOperationTypeEnumStringValues(), ",")))
	}
	if _, ok := mappingIamWorkRequestStatusEnum[string(m.Status)]; !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetIamWorkRequestStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// IamWorkRequestOperationTypeEnum Enum with underlying type: string
type IamWorkRequestOperationTypeEnum string

// Set of constants representing the allowable values for IamWorkRequestOperationTypeEnum
const (
	IamWorkRequestOperationTypeCreateDomain               IamWorkRequestOperationTypeEnum = "CREATE_DOMAIN"
	IamWorkRequestOperationTypeReplicateDomainToRegion    IamWorkRequestOperationTypeEnum = "REPLICATE_DOMAIN_TO_REGION"
	IamWorkRequestOperationTypeUpdateDomain               IamWorkRequestOperationTypeEnum = "UPDATE_DOMAIN"
	IamWorkRequestOperationTypeActivateDomain             IamWorkRequestOperationTypeEnum = "ACTIVATE_DOMAIN"
	IamWorkRequestOperationTypeDeactivateDomain           IamWorkRequestOperationTypeEnum = "DEACTIVATE_DOMAIN"
	IamWorkRequestOperationTypeDeleteDomain               IamWorkRequestOperationTypeEnum = "DELETE_DOMAIN"
	IamWorkRequestOperationTypeChangeCompartmentForDomain IamWorkRequestOperationTypeEnum = "CHANGE_COMPARTMENT_FOR_DOMAIN"
	IamWorkRequestOperationTypeChangeLicenseTypeForDomain IamWorkRequestOperationTypeEnum = "CHANGE_LICENSE_TYPE_FOR_DOMAIN"
)

var mappingIamWorkRequestOperationTypeEnum = map[string]IamWorkRequestOperationTypeEnum{
	"CREATE_DOMAIN":                  IamWorkRequestOperationTypeCreateDomain,
	"REPLICATE_DOMAIN_TO_REGION":     IamWorkRequestOperationTypeReplicateDomainToRegion,
	"UPDATE_DOMAIN":                  IamWorkRequestOperationTypeUpdateDomain,
	"ACTIVATE_DOMAIN":                IamWorkRequestOperationTypeActivateDomain,
	"DEACTIVATE_DOMAIN":              IamWorkRequestOperationTypeDeactivateDomain,
	"DELETE_DOMAIN":                  IamWorkRequestOperationTypeDeleteDomain,
	"CHANGE_COMPARTMENT_FOR_DOMAIN":  IamWorkRequestOperationTypeChangeCompartmentForDomain,
	"CHANGE_LICENSE_TYPE_FOR_DOMAIN": IamWorkRequestOperationTypeChangeLicenseTypeForDomain,
}

// GetIamWorkRequestOperationTypeEnumValues Enumerates the set of values for IamWorkRequestOperationTypeEnum
func GetIamWorkRequestOperationTypeEnumValues() []IamWorkRequestOperationTypeEnum {
	values := make([]IamWorkRequestOperationTypeEnum, 0)
	for _, v := range mappingIamWorkRequestOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetIamWorkRequestOperationTypeEnumStringValues Enumerates the set of values in String for IamWorkRequestOperationTypeEnum
func GetIamWorkRequestOperationTypeEnumStringValues() []string {
	return []string{
		"CREATE_DOMAIN",
		"REPLICATE_DOMAIN_TO_REGION",
		"UPDATE_DOMAIN",
		"ACTIVATE_DOMAIN",
		"DEACTIVATE_DOMAIN",
		"DELETE_DOMAIN",
		"CHANGE_COMPARTMENT_FOR_DOMAIN",
		"CHANGE_LICENSE_TYPE_FOR_DOMAIN",
	}
}

// IamWorkRequestStatusEnum Enum with underlying type: string
type IamWorkRequestStatusEnum string

// Set of constants representing the allowable values for IamWorkRequestStatusEnum
const (
	IamWorkRequestStatusAccepted   IamWorkRequestStatusEnum = "ACCEPTED"
	IamWorkRequestStatusInProgress IamWorkRequestStatusEnum = "IN_PROGRESS"
	IamWorkRequestStatusFailed     IamWorkRequestStatusEnum = "FAILED"
	IamWorkRequestStatusSucceeded  IamWorkRequestStatusEnum = "SUCCEEDED"
	IamWorkRequestStatusCanceling  IamWorkRequestStatusEnum = "CANCELING"
	IamWorkRequestStatusCanceled   IamWorkRequestStatusEnum = "CANCELED"
)

var mappingIamWorkRequestStatusEnum = map[string]IamWorkRequestStatusEnum{
	"ACCEPTED":    IamWorkRequestStatusAccepted,
	"IN_PROGRESS": IamWorkRequestStatusInProgress,
	"FAILED":      IamWorkRequestStatusFailed,
	"SUCCEEDED":   IamWorkRequestStatusSucceeded,
	"CANCELING":   IamWorkRequestStatusCanceling,
	"CANCELED":    IamWorkRequestStatusCanceled,
}

// GetIamWorkRequestStatusEnumValues Enumerates the set of values for IamWorkRequestStatusEnum
func GetIamWorkRequestStatusEnumValues() []IamWorkRequestStatusEnum {
	values := make([]IamWorkRequestStatusEnum, 0)
	for _, v := range mappingIamWorkRequestStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetIamWorkRequestStatusEnumStringValues Enumerates the set of values in String for IamWorkRequestStatusEnum
func GetIamWorkRequestStatusEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"FAILED",
		"SUCCEEDED",
		"CANCELING",
		"CANCELED",
	}
}
