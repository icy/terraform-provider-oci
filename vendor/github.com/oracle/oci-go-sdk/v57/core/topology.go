// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
//

package core

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v57/common"
	"strings"
)

// Topology Defines the representation of a virtual network topology.
type Topology interface {

	// Lists entities comprising the virtual network topology.
	GetEntities() []interface{}

	// Lists relationships between entities in the virtual network topology.
	GetRelationships() []TopologyEntityRelationship

	// Records when the virtual network topology was created, in RFC3339 (https://tools.ietf.org/html/rfc3339) format for date and time.
	GetTimeCreated() *common.SDKTime
}

type topology struct {
	JsonData      []byte
	Entities      []interface{}   `mandatory:"true" json:"entities"`
	Relationships json.RawMessage `mandatory:"true" json:"relationships"`
	TimeCreated   *common.SDKTime `mandatory:"true" json:"timeCreated"`
	Type          string          `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *topology) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalertopology topology
	s := struct {
		Model Unmarshalertopology
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Entities = s.Model.Entities
	m.Relationships = s.Model.Relationships
	m.TimeCreated = s.Model.TimeCreated
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *topology) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "VCN":
		mm := VcnTopology{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "NETWORKING":
		mm := NetworkingTopology{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SUBNET":
		mm := SubnetTopology{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetEntities returns Entities
func (m topology) GetEntities() []interface{} {
	return m.Entities
}

//GetRelationships returns Relationships
func (m topology) GetRelationships() json.RawMessage {
	return m.Relationships
}

//GetTimeCreated returns TimeCreated
func (m topology) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

func (m topology) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m topology) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TopologyTypeEnum Enum with underlying type: string
type TopologyTypeEnum string

// Set of constants representing the allowable values for TopologyTypeEnum
const (
	TopologyTypeNetworking TopologyTypeEnum = "NETWORKING"
	TopologyTypeVcn        TopologyTypeEnum = "VCN"
	TopologyTypeSubnet     TopologyTypeEnum = "SUBNET"
)

var mappingTopologyTypeEnum = map[string]TopologyTypeEnum{
	"NETWORKING": TopologyTypeNetworking,
	"VCN":        TopologyTypeVcn,
	"SUBNET":     TopologyTypeSubnet,
}

// GetTopologyTypeEnumValues Enumerates the set of values for TopologyTypeEnum
func GetTopologyTypeEnumValues() []TopologyTypeEnum {
	values := make([]TopologyTypeEnum, 0)
	for _, v := range mappingTopologyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTopologyTypeEnumStringValues Enumerates the set of values in String for TopologyTypeEnum
func GetTopologyTypeEnumStringValues() []string {
	return []string{
		"NETWORKING",
		"VCN",
		"SUBNET",
	}
}
