/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

type DataProtectionStatus string

const (
	DataProtectionStatus_ACTIVATED DataProtectionStatus = "ACTIVATED"
	DataProtectionStatus_DELETED   DataProtectionStatus = "DELETED"
	DataProtectionStatus_ARCHIVED  DataProtectionStatus = "ARCHIVED"
	DataProtectionStatus_DISABLED  DataProtectionStatus = "DISABLED"
)

type DeliveryDestinationType string

const (
	DeliveryDestinationType_S3  DeliveryDestinationType = "S3"
	DeliveryDestinationType_CWL DeliveryDestinationType = "CWL"
	DeliveryDestinationType_FH  DeliveryDestinationType = "FH"
)

type Distribution string

const (
	Distribution_Random      Distribution = "Random"
	Distribution_ByLogStream Distribution = "ByLogStream"
)

type ExportTaskStatusCode string

const (
	ExportTaskStatusCode_CANCELLED      ExportTaskStatusCode = "CANCELLED"
	ExportTaskStatusCode_COMPLETED      ExportTaskStatusCode = "COMPLETED"
	ExportTaskStatusCode_FAILED         ExportTaskStatusCode = "FAILED"
	ExportTaskStatusCode_PENDING        ExportTaskStatusCode = "PENDING"
	ExportTaskStatusCode_PENDING_CANCEL ExportTaskStatusCode = "PENDING_CANCEL"
	ExportTaskStatusCode_RUNNING        ExportTaskStatusCode = "RUNNING"
)

type InheritedProperty string

const (
	InheritedProperty_ACCOUNT_DATA_PROTECTION InheritedProperty = "ACCOUNT_DATA_PROTECTION"
)

type OrderBy string

const (
	OrderBy_LogStreamName OrderBy = "LogStreamName"
	OrderBy_LastEventTime OrderBy = "LastEventTime"
)

type OutputFormat string

const (
	OutputFormat_json    OutputFormat = "json"
	OutputFormat_plain   OutputFormat = "plain"
	OutputFormat_w3c     OutputFormat = "w3c"
	OutputFormat_raw     OutputFormat = "raw"
	OutputFormat_parquet OutputFormat = "parquet"
)

type PolicyType string

const (
	PolicyType_DATA_PROTECTION_POLICY PolicyType = "DATA_PROTECTION_POLICY"
)

type QueryStatus string

const (
	QueryStatus_Scheduled QueryStatus = "Scheduled"
	QueryStatus_Running   QueryStatus = "Running"
	QueryStatus_Complete  QueryStatus = "Complete"
	QueryStatus_Failed    QueryStatus = "Failed"
	QueryStatus_Cancelled QueryStatus = "Cancelled"
	QueryStatus_Timeout   QueryStatus = "Timeout"
	QueryStatus_Unknown   QueryStatus = "Unknown"
)

type Scope string

const (
	Scope_ALL Scope = "ALL"
)

type StandardUnit string

const (
	StandardUnit_Seconds          StandardUnit = "Seconds"
	StandardUnit_Microseconds     StandardUnit = "Microseconds"
	StandardUnit_Milliseconds     StandardUnit = "Milliseconds"
	StandardUnit_Bytes            StandardUnit = "Bytes"
	StandardUnit_Kilobytes        StandardUnit = "Kilobytes"
	StandardUnit_Megabytes        StandardUnit = "Megabytes"
	StandardUnit_Gigabytes        StandardUnit = "Gigabytes"
	StandardUnit_Terabytes        StandardUnit = "Terabytes"
	StandardUnit_Bits             StandardUnit = "Bits"
	StandardUnit_Kilobits         StandardUnit = "Kilobits"
	StandardUnit_Megabits         StandardUnit = "Megabits"
	StandardUnit_Gigabits         StandardUnit = "Gigabits"
	StandardUnit_Terabits         StandardUnit = "Terabits"
	StandardUnit_Percent          StandardUnit = "Percent"
	StandardUnit_Count            StandardUnit = "Count"
	StandardUnit_Bytes_Second     StandardUnit = "Bytes/Second"
	StandardUnit_Kilobytes_Second StandardUnit = "Kilobytes/Second"
	StandardUnit_Megabytes_Second StandardUnit = "Megabytes/Second"
	StandardUnit_Gigabytes_Second StandardUnit = "Gigabytes/Second"
	StandardUnit_Terabytes_Second StandardUnit = "Terabytes/Second"
	StandardUnit_Bits_Second      StandardUnit = "Bits/Second"
	StandardUnit_Kilobits_Second  StandardUnit = "Kilobits/Second"
	StandardUnit_Megabits_Second  StandardUnit = "Megabits/Second"
	StandardUnit_Gigabits_Second  StandardUnit = "Gigabits/Second"
	StandardUnit_Terabits_Second  StandardUnit = "Terabits/Second"
	StandardUnit_Count_Second     StandardUnit = "Count/Second"
	StandardUnit_None             StandardUnit = "None"
)
