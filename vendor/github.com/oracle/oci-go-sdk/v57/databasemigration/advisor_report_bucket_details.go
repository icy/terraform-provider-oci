// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v57/common"
	"strings"
)

// AdvisorReportBucketDetails Details to access Pre-Migration Advisor report in the specified Object Storage bucket, if any.
type AdvisorReportBucketDetails struct {

	// Name of the bucket containing the Pre-Migration Advisor report.
	BucketName *string `mandatory:"true" json:"bucketName"`

	// Object Storage namespace.
	Namespace *string `mandatory:"true" json:"namespace"`

	// Pre-Migration Advisor report object name.
	ObjectName *string `mandatory:"true" json:"objectName"`
}

func (m AdvisorReportBucketDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AdvisorReportBucketDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
