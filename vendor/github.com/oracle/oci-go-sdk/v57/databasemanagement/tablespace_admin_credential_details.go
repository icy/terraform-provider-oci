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
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v57/common"
	"strings"
)

// TablespaceAdminCredentialDetails The credential to connect to the database to perform tablespace administration tasks.
type TablespaceAdminCredentialDetails interface {

	// The user to connect to the database.
	GetUsername() *string

	// The role of the database user.
	GetRole() TablespaceAdminCredentialDetailsRoleEnum
}

type tablespaceadmincredentialdetails struct {
	JsonData                      []byte
	Username                      *string                                  `mandatory:"true" json:"username"`
	Role                          TablespaceAdminCredentialDetailsRoleEnum `mandatory:"true" json:"role"`
	TablespaceAdminCredentialType string                                   `json:"tablespaceAdminCredentialType"`
}

// UnmarshalJSON unmarshals json
func (m *tablespaceadmincredentialdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalertablespaceadmincredentialdetails tablespaceadmincredentialdetails
	s := struct {
		Model Unmarshalertablespaceadmincredentialdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Username = s.Model.Username
	m.Role = s.Model.Role
	m.TablespaceAdminCredentialType = s.Model.TablespaceAdminCredentialType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *tablespaceadmincredentialdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.TablespaceAdminCredentialType {
	case "PASSWORD":
		mm := TablespaceAdminPasswordCredentialDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SECRET":
		mm := TablespaceAdminSecretCredentialDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetUsername returns Username
func (m tablespaceadmincredentialdetails) GetUsername() *string {
	return m.Username
}

//GetRole returns Role
func (m tablespaceadmincredentialdetails) GetRole() TablespaceAdminCredentialDetailsRoleEnum {
	return m.Role
}

func (m tablespaceadmincredentialdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m tablespaceadmincredentialdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := mappingTablespaceAdminCredentialDetailsRoleEnum[string(m.Role)]; !ok && m.Role != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Role: %s. Supported values are: %s.", m.Role, strings.Join(GetTablespaceAdminCredentialDetailsRoleEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TablespaceAdminCredentialDetailsRoleEnum Enum with underlying type: string
type TablespaceAdminCredentialDetailsRoleEnum string

// Set of constants representing the allowable values for TablespaceAdminCredentialDetailsRoleEnum
const (
	TablespaceAdminCredentialDetailsRoleNormal TablespaceAdminCredentialDetailsRoleEnum = "NORMAL"
	TablespaceAdminCredentialDetailsRoleSysdba TablespaceAdminCredentialDetailsRoleEnum = "SYSDBA"
)

var mappingTablespaceAdminCredentialDetailsRoleEnum = map[string]TablespaceAdminCredentialDetailsRoleEnum{
	"NORMAL": TablespaceAdminCredentialDetailsRoleNormal,
	"SYSDBA": TablespaceAdminCredentialDetailsRoleSysdba,
}

// GetTablespaceAdminCredentialDetailsRoleEnumValues Enumerates the set of values for TablespaceAdminCredentialDetailsRoleEnum
func GetTablespaceAdminCredentialDetailsRoleEnumValues() []TablespaceAdminCredentialDetailsRoleEnum {
	values := make([]TablespaceAdminCredentialDetailsRoleEnum, 0)
	for _, v := range mappingTablespaceAdminCredentialDetailsRoleEnum {
		values = append(values, v)
	}
	return values
}

// GetTablespaceAdminCredentialDetailsRoleEnumStringValues Enumerates the set of values in String for TablespaceAdminCredentialDetailsRoleEnum
func GetTablespaceAdminCredentialDetailsRoleEnumStringValues() []string {
	return []string{
		"NORMAL",
		"SYSDBA",
	}
}

// TablespaceAdminCredentialDetailsTablespaceAdminCredentialTypeEnum Enum with underlying type: string
type TablespaceAdminCredentialDetailsTablespaceAdminCredentialTypeEnum string

// Set of constants representing the allowable values for TablespaceAdminCredentialDetailsTablespaceAdminCredentialTypeEnum
const (
	TablespaceAdminCredentialDetailsTablespaceAdminCredentialTypeSecret   TablespaceAdminCredentialDetailsTablespaceAdminCredentialTypeEnum = "SECRET"
	TablespaceAdminCredentialDetailsTablespaceAdminCredentialTypePassword TablespaceAdminCredentialDetailsTablespaceAdminCredentialTypeEnum = "PASSWORD"
)

var mappingTablespaceAdminCredentialDetailsTablespaceAdminCredentialTypeEnum = map[string]TablespaceAdminCredentialDetailsTablespaceAdminCredentialTypeEnum{
	"SECRET":   TablespaceAdminCredentialDetailsTablespaceAdminCredentialTypeSecret,
	"PASSWORD": TablespaceAdminCredentialDetailsTablespaceAdminCredentialTypePassword,
}

// GetTablespaceAdminCredentialDetailsTablespaceAdminCredentialTypeEnumValues Enumerates the set of values for TablespaceAdminCredentialDetailsTablespaceAdminCredentialTypeEnum
func GetTablespaceAdminCredentialDetailsTablespaceAdminCredentialTypeEnumValues() []TablespaceAdminCredentialDetailsTablespaceAdminCredentialTypeEnum {
	values := make([]TablespaceAdminCredentialDetailsTablespaceAdminCredentialTypeEnum, 0)
	for _, v := range mappingTablespaceAdminCredentialDetailsTablespaceAdminCredentialTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTablespaceAdminCredentialDetailsTablespaceAdminCredentialTypeEnumStringValues Enumerates the set of values in String for TablespaceAdminCredentialDetailsTablespaceAdminCredentialTypeEnum
func GetTablespaceAdminCredentialDetailsTablespaceAdminCredentialTypeEnumStringValues() []string {
	return []string{
		"SECRET",
		"PASSWORD",
	}
}
