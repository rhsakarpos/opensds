// Copyright (c) 2017 Huawei Technologies Co., Ltd. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
This module implements the common data structure.

*/

package model

// CollectMetricSpec is used to map the inputs from the POST call to collect metrics
type CollectMetricSpec struct {
	*BaseModel

	// the instance on which the metrics are to be collected
	InstanceId string `json:"instanceId,omitempty"`

	// the list of metrics to be collected
	Metrics []string `json:"metrics,omitempty"`
}

type Metric struct {
	/*Following are the labels associated with Metric, same as Prometheus labels
	Example: {device="dm-0",instance="121.244.95.60:12419",job="prometheus"}
	Instance ID -\> volumeID/NodeID*/
	instanceID string

	// instance name -\> volume name / node name etc.
	instanceName string

	// job -\> Prometheus/openSDS
	job string

	/*associator - Some metric would need specific fields to relate components.
	Use case could be to query volumes of a particular pool. Attaching the related
	components as labels would help us to form promQl query efficiently.

	Example: node_disk_read_bytes_total{instance="121.244.95.60"}

	Above query will respond with all disks associated with node 121.244.95.60
	Since associated components vary, we will keep a map in metric struct to denote
	the associated component type as key and component name as value

	Example: associator[pool]=pool1*/

	associator map[string]string

	// Following fields can be used to form a unique metric name source -\> Node/Dock
	source string

	// component -\> disk/logicalVolume/VG etc
	component string

	// name -\> metric name -\> readRequests/WriteRequests/Latency etc
	name string

	// unit -\> seconds/bytes/MBs etc
	unit string

	/*If isAggregated ='True' then type of aggregation can be set in this field
	ie:- if collector is aggregating some metrics and producing a new metric of
	higher level constructs, then this field can be set as 'Total' to indicate it is
	aggregated/derived from other metrics.*/
	isAggregated bool

	// metric timestamp
	timestamp int64

	// metric value
	value float64
}
