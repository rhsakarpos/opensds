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
package lvm

import (
	"fmt"
	"github.com/opensds/opensds/pkg/model"
)


type MetricDriver struct {

	cli  *MetricsCli
}

func (d *MetricDriver) CollectMetrics(metricsList []string,instanceID string) (metricArray []*model.Metric, err error) {


	metricMap,err := d.cli.CollectMetrics(metricsList,instanceID)
	fmt.Println(metricMap)

	return
}

func (d *MetricDriver) Setup() error {

	return nil
}

func (*MetricDriver) Teardown() error { return nil }









