// Copyright 2018 The Kubeflow Authors
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
	// api "github.com/kubeflow/pipelines/backend/api/go_client"
	apiv2beta1 "github.com/kubeflow/pipelines/backend/api/go_client"
	"github.com/kubeflow/pipelines/backend/src/apiserver/common"
	"github.com/kubeflow/pipelines/backend/src/apiserver/model"
	"github.com/kubeflow/pipelines/backend/src/common/util"
)

func ToApiExperiment(experiment *model.Experiment) *apiv2beta1.Experiment {
	resourceReferences := []*apiv2beta1.ResourceReference(nil)
	if common.IsMultiUserMode() {
		resourceReferences = []*apiv2beta1.ResourceReference{
			{
				Key: &apiv2beta1.ResourceKey{
					Type: apiv2beta1.ResourceType_NAMESPACE,
					Id:   experiment.Namespace,
				},
				Relationship: apiv2beta1.Relationship_OWNER,
			},
		}
	}
	return &apiv2beta1.Experiment{
		Id:                 experiment.UUID,
		Name:               experiment.Name,
		Description:        experiment.Description,
		CreatedAt:          &timestamp.Timestamp{Seconds: experiment.CreatedAtInSec},
		ResourceReferences: resourceReferences,
		StorageState:       apiv2beta1.Experiment_StorageState(apiv2beta1.Experiment_StorageState_value[experiment.StorageState]),
	}
}

func ToApiExperiments(experiments []*model.Experiment) []*apiv2beta1.Experiment {
	apiExperiments := make([]*apiv2beta1.Experiment, 0)
	for _, experiment := range experiments {
		apiExperiments = append(apiExperiments, ToApiExperiment(experiment))
	}
	return apiExperiments
}

func ToApiPipeline(pipeline *model.Pipeline) *apiv2beta1.Pipeline {
	params, err := toApiParameters(pipeline.Parameters)
	if err != nil {
		return &apiv2beta1.Pipeline{
			Id:    pipeline.UUID,
			Error: err.Error(),
		}
	}

	defaultVersion, err := ToApiPipelineVersion(pipeline.DefaultVersion)
	if err != nil {
		return &apiv2beta1.Pipeline{
			Id:    pipeline.UUID,
			Error: err.Error(),
		}
	}

	return &apiv2beta1.Pipeline{
		Id:             pipeline.UUID,
		CreatedAt:      &timestamp.Timestamp{Seconds: pipeline.CreatedAtInSec},
		Name:           pipeline.Name,
		Description:    pipeline.Description,
		Parameters:     params,
		DefaultVersion: defaultVersion,
	}
}

func ToApiPipelineVersion(version *model.PipelineVersion) (*apiv2beta1.PipelineVersion, error) {
	if version == nil {
		return nil, nil
	}
	params, err := toApiParameters(version.Parameters)
	if err != nil {
		return nil, err
	}

	return &apiv2beta1.PipelineVersion{
		Id:            version.UUID,
		Name:          version.Name,
		CreatedAt:     &timestamp.Timestamp{Seconds: version.CreatedAtInSec},
		Parameters:    params,
		Description:   version.Description,
		CodeSourceUrl: version.CodeSourceUrl,
		ResourceReferences: []*apiv2beta1.ResourceReference{
			{
				Key: &apiv2beta1.ResourceKey{
					Id:   version.PipelineId,
					Type: apiv2beta1.ResourceType_PIPELINE,
				},
				Relationship: apiv2beta1.Relationship_OWNER,
			},
		},
	}, nil
}

func ToApiPipelineVersions(versions []*model.PipelineVersion) ([]*apiv2beta1.PipelineVersion, error) {
	apiVersions := make([]*apiv2beta1.PipelineVersion, 0)
	for _, version := range versions {
		v, _ := ToApiPipelineVersion(version)
		apiVersions = append(apiVersions, v)
	}
	return apiVersions, nil
}

func ToApiPipelines(pipelines []*model.Pipeline) []*apiv2beta1.Pipeline {
	apiPipelines := make([]*apiv2beta1.Pipeline, 0)
	for _, pipeline := range pipelines {
		apiPipelines = append(apiPipelines, ToApiPipeline(pipeline))
	}
	return apiPipelines
}

func toApiParameters(paramsString string) ([]*apiv2beta1.Parameter, error) {
	if paramsString == "" {
		return nil, nil
	}
	params, err := util.UnmarshalParameters(util.ArgoWorkflow, paramsString)
	if err != nil {
		return nil, util.NewInternalServerError(err, "Parameter with wrong format is stored")
	}
	apiParams := make([]*apiv2beta1.Parameter, 0)
	for _, param := range params {
		var value string
		if param.Value != nil {
			value = *param.Value
		}
		apiParam := apiv2beta1.Parameter{
			Name:  param.Name,
			Value: value,
		}
		apiParams = append(apiParams, &apiParam)
	}
	return apiParams, nil
}

func toApiRun(run *model.Run) *apiv2beta1.Run {
	params, err := toApiParameters(run.Parameters)
	if err != nil {
		return &apiv2beta1.Run{
			Id:    run.UUID,
			Error: err.Error(),
		}
	}
	var metrics []*apiv2beta1.RunMetric
	if run.Metrics != nil {
		for _, metric := range run.Metrics {
			metrics = append(metrics, ToApiRunMetric(metric))
		}
	}
	return &apiv2beta1.Run{
		CreatedAt:      &timestamp.Timestamp{Seconds: run.CreatedAtInSec},
		Id:             run.UUID,
		Metrics:        metrics,
		Name:           run.DisplayName,
		ServiceAccount: run.ServiceAccount,
		StorageState:   apiv2beta1.Run_StorageState(apiv2beta1.Run_StorageState_value[run.StorageState]),
		Description:    run.Description,
		ScheduledAt:    &timestamp.Timestamp{Seconds: run.ScheduledAtInSec},
		FinishedAt:     &timestamp.Timestamp{Seconds: run.FinishedAtInSec},
		Status:         run.Conditions,
		PipelineSpec: &apiv2beta1.PipelineSpec{
			PipelineId:       run.PipelineId,
			PipelineName:     run.PipelineName,
			WorkflowManifest: run.WorkflowSpecManifest,
			PipelineManifest: run.PipelineSpecManifest,
			Parameters:       params,
		},
		ResourceReferences: toApiResourceReferences(run.ResourceReferences),
	}
}

func ToApiRuns(runs []*model.Run) []*apiv2beta1.Run {
	apiRuns := make([]*apiv2beta1.Run, 0)
	for _, run := range runs {
		apiRuns = append(apiRuns, toApiRun(run))
	}
	return apiRuns
}

func ToApiRunDetail(run *model.RunDetail) *apiv2beta1.RunDetail {
	return &apiv2beta1.RunDetail{
		Run: toApiRun(&run.Run),
		PipelineRuntime: &apiv2beta1.PipelineRuntime{
			WorkflowManifest: run.WorkflowRuntimeManifest,
			PipelineManifest: run.PipelineRuntimeManifest,
		},
	}
}

func ToApiTask(task *model.Task) *apiv2beta1.Task {
	return &apiv2beta1.Task{
		Id:              task.UUID,
		Namespace:       task.Namespace,
		PipelineName:    task.PipelineName,
		RunId:           task.RunUUID,
		MlmdExecutionID: task.MLMDExecutionID,
		CreatedAt:       &timestamp.Timestamp{Seconds: task.CreatedTimestamp},
		FinishedAt:      &timestamp.Timestamp{Seconds: task.FinishedTimestamp},
		Fingerprint:     task.Fingerprint,
	}
}

func ToApiTasks(tasks []*model.Task) []*apiv2beta1.Task {
	apiTasks := make([]*apiv2beta1.Task, 0)
	for _, task := range tasks {
		apiTasks = append(apiTasks, ToApiTask(task))
	}
	return apiTasks
}
func ToApiJob(job *model.Job) *apiv2beta1.Job {
	params, err := toApiParameters(job.Parameters)
	if err != nil {
		return &apiv2beta1.Job{
			Id:    job.UUID,
			Error: err.Error(),
		}
	}
	return &apiv2beta1.Job{
		Id:             job.UUID,
		Name:           job.DisplayName,
		ServiceAccount: job.ServiceAccount,
		Description:    job.Description,
		Enabled:        job.Enabled,
		CreatedAt:      &timestamp.Timestamp{Seconds: job.CreatedAtInSec},
		UpdatedAt:      &timestamp.Timestamp{Seconds: job.UpdatedAtInSec},
		Status:         job.Conditions,
		MaxConcurrency: job.MaxConcurrency,
		NoCatchup:      job.NoCatchup,
		Trigger:        toApiTrigger(job.Trigger),
		PipelineSpec: &apiv2beta1.PipelineSpec{
			PipelineId:       job.PipelineId,
			PipelineName:     job.PipelineName,
			WorkflowManifest: job.WorkflowSpecManifest,
			PipelineManifest: job.PipelineSpecManifest,
			Parameters:       params,
		},
		ResourceReferences: toApiResourceReferences(job.ResourceReferences),
	}
}

func ToApiJobs(jobs []*model.Job) []*apiv2beta1.Job {
	apiJobs := make([]*apiv2beta1.Job, 0)
	for _, job := range jobs {
		apiJobs = append(apiJobs, ToApiJob(job))
	}
	return apiJobs
}

func ToApiRunMetric(metric *model.RunMetric) *apiv2beta1.RunMetric {
	return &apiv2beta1.RunMetric{
		Name:   metric.Name,
		NodeId: metric.NodeID,
		Value: &apiv2beta1.RunMetric_NumberValue{
			NumberValue: metric.NumberValue,
		},
		Format: apiv2beta1.RunMetric_Format(apiv2beta1.RunMetric_Format_value[metric.Format]),
	}
}

func toApiResourceReferences(references []*model.ResourceReference) []*apiv2beta1.ResourceReference {
	var apiReferences []*apiv2beta1.ResourceReference
	for _, ref := range references {
		apiReferences = append(apiReferences, &apiv2beta1.ResourceReference{
			Key: &apiv2beta1.ResourceKey{
				Type: toApiResourceType(ref.ReferenceType),
				Id:   ref.ReferenceUUID,
			},
			Name:         ref.ReferenceName,
			Relationship: toApiRelationship(ref.Relationship),
		})
	}
	return apiReferences
}

func toApiResourceType(modelType model.ResourceType) apiv2beta1.ResourceType {
	switch modelType {
	case common.Experiment:
		return apiv2beta1.ResourceType_EXPERIMENT
	case common.Job:
		return apiv2beta1.ResourceType_JOB
	case common.PipelineVersion:
		return apiv2beta1.ResourceType_PIPELINE_VERSION
	case common.Namespace:
		return apiv2beta1.ResourceType_NAMESPACE
	default:
		return apiv2beta1.ResourceType_UNKNOWN_RESOURCE_TYPE
	}
}

func toApiRelationship(r model.Relationship) apiv2beta1.Relationship {
	switch r {
	case common.Creator:
		return apiv2beta1.Relationship_CREATOR
	case common.Owner:
		return apiv2beta1.Relationship_OWNER
	default:
		return apiv2beta1.Relationship_UNKNOWN_RELATIONSHIP
	}
}

func toApiTrigger(trigger model.Trigger) *apiv2beta1.Trigger {
	if trigger.Cron != nil && *trigger.Cron != "" {
		var cronSchedule apiv2beta1.CronSchedule
		cronSchedule.Cron = *trigger.Cron
		if trigger.CronScheduleStartTimeInSec != nil {
			cronSchedule.StartTime = &timestamp.Timestamp{
				Seconds: *trigger.CronScheduleStartTimeInSec}
		}
		if trigger.CronScheduleEndTimeInSec != nil {
			cronSchedule.EndTime = &timestamp.Timestamp{
				Seconds: *trigger.CronScheduleEndTimeInSec}
		}
		return &apiv2beta1.Trigger{Trigger: &apiv2beta1.Trigger_CronSchedule{CronSchedule: &cronSchedule}}
	}

	if trigger.IntervalSecond != nil && *trigger.IntervalSecond != 0 {
		var periodicSchedule apiv2beta1.PeriodicSchedule
		periodicSchedule.IntervalSecond = *trigger.IntervalSecond
		if trigger.PeriodicScheduleStartTimeInSec != nil {
			periodicSchedule.StartTime = &timestamp.Timestamp{
				Seconds: *trigger.PeriodicScheduleStartTimeInSec}
		}
		if trigger.PeriodicScheduleEndTimeInSec != nil {
			periodicSchedule.EndTime = &timestamp.Timestamp{
				Seconds: *trigger.PeriodicScheduleEndTimeInSec}
		}
		return &apiv2beta1.Trigger{Trigger: &apiv2beta1.Trigger_PeriodicSchedule{PeriodicSchedule: &periodicSchedule}}
	}
	return &apiv2beta1.Trigger{}
}
