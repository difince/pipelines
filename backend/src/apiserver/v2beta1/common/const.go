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

package common

import (
	apiv2beta1 "github.com/kubeflow/pipelines/backend/api/v2beta1/go_client"
	"github.com/kubeflow/pipelines/backend/src/apiserver/v2beta1/model"
	"github.com/kubeflow/pipelines/backend/src/common/util"
)

const (
	Experiment      model.ResourceType = "Experiment"
	Job             model.ResourceType = "Job"
	Run             model.ResourceType = "Run"
	Pipeline        model.ResourceType = "pipeline"
	PipelineVersion model.ResourceType = "PipelineVersion"
	Namespace       model.ResourceType = "Namespace"
)

const (
	RbacKubeflowGroup    = "kubeflow.org"
	RbacPipelinesGroup   = "pipelines.kubeflow.org"
	RbacPipelinesVersion = "v1beta1"

	RbacResourceTypePipelines      = "pipelines"
	RbacResourceTypeExperiments    = "experiments"
	RbacResourceTypeRuns           = "runs"
	RbacResourceTypeJobs           = "jobs"
	RbacResourceTypeViewers        = "viewers"
	RbacResourceTypeVisualizations = "visualizations"

	RbacResourceVerbArchive   = "archive"
	RbacResourceVerbUpdate    = "update"
	RbacResourceVerbCreate    = "create"
	RbacResourceVerbDelete    = "delete"
	RbacResourceVerbDisable   = "disable"
	RbacResourceVerbEnable    = "enable"
	RbacResourceVerbGet       = "get"
	RbacResourceVerbList      = "list"
	RbacResourceVerbRetry     = "retry"
	RbacResourceVerbTerminate = "terminate"
	RbacResourceVerbUnarchive = "unarchive"
)

const (
	Owner   model.Relationship = "Owner"
	Creator model.Relationship = "Creator"
)

const (
	GoogleIAPUserIdentityHeader    string = "x-goog-authenticated-user-email"
	GoogleIAPUserIdentityPrefix    string = "accounts.google.com:"
	AuthorizationBearerTokenHeader string = "Authorization"
	AuthorizationBearerTokenPrefix string = "Bearer "
)

const DefaultTokenReviewAudience string = "pipelines.kubeflow.org"

func ToModelResourceType(apiType apiv2beta1.ResourceType) (model.ResourceType, error) {
	switch apiType {
	case apiv2beta1.ResourceType_EXPERIMENT:
		return Experiment, nil
	case apiv2beta1.ResourceType_JOB:
		return Job, nil
	case apiv2beta1.ResourceType_PIPELINE_VERSION:
		return PipelineVersion, nil
	case apiv2beta1.ResourceType_NAMESPACE:
		return Namespace, nil
	default:
		return "", util.NewInvalidInputError("Unsupported resource type: %s", apiv2beta1.ResourceType_name[int32(apiType)])
	}
}

func ToModelRelationship(r apiv2beta1.Relationship) (model.Relationship, error) {
	switch r {
	case apiv2beta1.Relationship_CREATOR:
		return Creator, nil
	case apiv2beta1.Relationship_OWNER:
		return Owner, nil
	default:
		return "", util.NewInvalidInputError("Unsupported resource relationship: %s", apiv2beta1.Relationship_name[int32(r)])
	}
}
