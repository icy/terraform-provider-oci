// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v57/common"
	"strings"
)

// TargetResponderRecipeResponderRule Details of ResponderRule.
type TargetResponderRecipeResponderRule struct {

	// Identifier for ResponderRule.
	ResponderRuleId *string `mandatory:"true" json:"responderRuleId"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// ResponderRule Display Name
	DisplayName *string `mandatory:"false" json:"displayName"`

	// ResponderRule Description
	Description *string `mandatory:"false" json:"description"`

	// Type of Responder
	Type ResponderTypeEnum `mandatory:"false" json:"type,omitempty"`

	// List of Policy
	Policies []string `mandatory:"false" json:"policies"`

	// Supported Execution Modes
	SupportedModes []TargetResponderRecipeResponderRuleSupportedModesEnum `mandatory:"false" json:"supportedModes,omitempty"`

	Details *ResponderRuleDetails `mandatory:"false" json:"details"`

	// The date and time the target responder recipe rule was created. Format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the target responder recipe rule was updated. Format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current state of the ResponderRule.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`
}

func (m TargetResponderRecipeResponderRule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TargetResponderRecipeResponderRule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := mappingResponderTypeEnum[string(m.Type)]; !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetResponderTypeEnumStringValues(), ",")))
	}
	for _, val := range m.SupportedModes {
		if _, ok := mappingTargetResponderRecipeResponderRuleSupportedModesEnum[string(val)]; !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SupportedModes: %s. Supported values are: %s.", val, strings.Join(GetTargetResponderRecipeResponderRuleSupportedModesEnumStringValues(), ",")))
		}
	}

	if _, ok := mappingLifecycleStateEnum[string(m.LifecycleState)]; !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TargetResponderRecipeResponderRuleSupportedModesEnum Enum with underlying type: string
type TargetResponderRecipeResponderRuleSupportedModesEnum string

// Set of constants representing the allowable values for TargetResponderRecipeResponderRuleSupportedModesEnum
const (
	TargetResponderRecipeResponderRuleSupportedModesAutoaction TargetResponderRecipeResponderRuleSupportedModesEnum = "AUTOACTION"
	TargetResponderRecipeResponderRuleSupportedModesUseraction TargetResponderRecipeResponderRuleSupportedModesEnum = "USERACTION"
)

var mappingTargetResponderRecipeResponderRuleSupportedModesEnum = map[string]TargetResponderRecipeResponderRuleSupportedModesEnum{
	"AUTOACTION": TargetResponderRecipeResponderRuleSupportedModesAutoaction,
	"USERACTION": TargetResponderRecipeResponderRuleSupportedModesUseraction,
}

// GetTargetResponderRecipeResponderRuleSupportedModesEnumValues Enumerates the set of values for TargetResponderRecipeResponderRuleSupportedModesEnum
func GetTargetResponderRecipeResponderRuleSupportedModesEnumValues() []TargetResponderRecipeResponderRuleSupportedModesEnum {
	values := make([]TargetResponderRecipeResponderRuleSupportedModesEnum, 0)
	for _, v := range mappingTargetResponderRecipeResponderRuleSupportedModesEnum {
		values = append(values, v)
	}
	return values
}

// GetTargetResponderRecipeResponderRuleSupportedModesEnumStringValues Enumerates the set of values in String for TargetResponderRecipeResponderRuleSupportedModesEnum
func GetTargetResponderRecipeResponderRuleSupportedModesEnumStringValues() []string {
	return []string{
		"AUTOACTION",
		"USERACTION",
	}
}
