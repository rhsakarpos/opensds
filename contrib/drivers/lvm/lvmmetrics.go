package lvm

import (
	"fmt"
	"github.com/opensds/opensds/pkg/model"
	"strconv"
)

// Copyright (c) 2017 Huawei Technologies Co., Ltd. All Rights Reserved.
//
//    Licensed under the Apache License, Version 2.0 (the "License"); you may
//    not use this file except in compliance with the License. You may obtain
//    a copy of the License at
//
//         http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
//    WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
//    License for the specific language governing permissions and limitations
//    under the License.

type MetricDriver struct {

	cli  *MetricsCli
}





func (d *MetricDriver) CollectMetrics(metricsList []string,instanceID string) (metricArray []model.MetricSpec, err error) {


	metricMap,err := d.cli.CollectMetrics(metricsList,instanceID)
	fmt.Println(metricMap)

	var tempmetricArray []model.MetricSpec
	for _,element := range metricsList {
		val,_:=strconv.ParseFloat(metricMap[element],64)
		m := make(map[string]string)

		metricValue := model.Metric{
			Timestamp:000,
			Value:val,
		}
		metricValues := make([]model.Metric,1)
		metricValues = append(metricValues, metricValue)

		metric := model.MetricSpec{
			InstanceID: instanceID,

			InstanceName:metricMap["InstanceName"],

			Job: "OpenSDS",

			Associator: m,

			Source: "Node",

			Component: "Volume",

			Name: element,

			Unit: "KB/s",

			IsAggregated: false,

			MetricValues: metricValues,
		}
		tempmetricArray = append(tempmetricArray, metric)
	}
	metricArray=tempmetricArray
	fmt.Println(metricArray)
	return
}

func (d *MetricDriver) Setup() error {



	return nil
}

func (*MetricDriver) Unset() error { return nil }









