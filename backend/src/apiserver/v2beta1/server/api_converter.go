// Copyright 2022 The Kubeflow Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	apiv2beta1 "github.com/kubeflow/pipelines/backend/api/v2beta1/go_client"
	"github.com/kubeflow/pipelines/backend/src/apiserver/model"
)

func ToApiExperiment(experiment *model.Experiment) *apiv2beta1.Experiment {
	return &apiv2beta1.Experiment{
		Id:           experiment.UUID,
		Name:         experiment.Name,
		Description:  experiment.Description,
		CreatedAt:    &timestamp.Timestamp{Seconds: experiment.CreatedAtInSec},
		Namespace:    experiment.Namespace,
		StorageState: apiv2beta1.Experiment_StorageState(apiv2beta1.Experiment_StorageState_value[experiment.StorageState]),
	}
}

func ToApiExperiments(experiments []*model.Experiment) []*apiv2beta1.Experiment {
	apiExperiments := make([]*apiv2beta1.Experiment, 0)
	for _, experiment := range experiments {
		apiExperiments = append(apiExperiments, ToApiExperiment(experiment))
	}
	return apiExperiments
}
