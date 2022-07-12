package v2beta1

import (
	"github.com/golang/protobuf/ptypes/timestamp"

	apiv2beta1 "github.com/kubeflow/pipelines/backend/api/v2beta1/go_client"

	"github.com/kubeflow/pipelines/backend/src/apiserver/common"
	"github.com/kubeflow/pipelines/backend/src/apiserver/model"
	"github.com/kubeflow/pipelines/backend/src/common/util"
)

func ToApiRuns(runs []*model.Run) []*apiv2beta1.Run {
	apiRuns := make([]*apiv2beta1.Run, 0)
	for _, run := range runs {
		apiRuns = append(apiRuns, toApiRun(run))
	}
	return apiRuns
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
	}
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
