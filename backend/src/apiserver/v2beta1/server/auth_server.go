package server

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	commonapi "github.com/kubeflow/pipelines/backend/api/common/go_client"
	"github.com/kubeflow/pipelines/backend/src/apiserver/common"
	"github.com/kubeflow/pipelines/backend/src/apiserver/resource"
	"github.com/kubeflow/pipelines/backend/src/common/util"
)

var rbacResourceTypeToGroup = map[string]string{
	common.RbacResourceTypePipelines:      common.RbacPipelinesGroup,
	common.RbacResourceTypeExperiments:    common.RbacPipelinesGroup,
	common.RbacResourceTypeRuns:           common.RbacPipelinesGroup,
	common.RbacResourceTypeJobs:           common.RbacPipelinesGroup,
	common.RbacResourceTypeViewers:        common.RbacKubeflowGroup,
	common.RbacResourceTypeVisualizations: common.RbacPipelinesGroup,
}

type AuthServer struct {
	resourceManager *resource.ResourceManager
}

func (s *AuthServer) Authorize(ctx context.Context, request *commonapi.AuthorizeRequest) (
	*empty.Empty, error) {
	fmt.Println("v2: Authorize has been called")
	err := ValidateAuthorizeRequest(request)
	if err != nil {
		return nil, util.Wrap(err, "Authorize request is not valid")
	}

	//TODO V2 isAuthorized need to be moved in a common folder.

	//namespace := strings.ToLower(request.GetNamespace())
	//verb := strings.ToLower(request.GetVerb().String())
	//resource := strings.ToLower(request.GetResources().String())
	//resourceAttributes := &authorizationv1.ResourceAttributes{
	//	Namespace:   namespace,
	//	Verb:        verb,
	//	Group:       rbacResourceTypeToGroup[resource],
	//	Version:     common.RbacPipelinesVersion,
	//	Resource:    resource,
	//	Subresource: "",
	//	Name:        "",
	//}
	//err = isAuthorized(s.resourceManager, ctx, resourceAttributes)
	//if err != nil {
	//	return nil, util.Wrap(err, "Failed to authorize the request")
	//}

	return &empty.Empty{}, nil
}

func ValidateAuthorizeRequest(request *commonapi.AuthorizeRequest) error {
	if request == nil {
		return util.NewInvalidInputError("request object is empty.")
	}
	if len(request.Namespace) == 0 {
		return util.NewInvalidInputError("Namespace is empty. Please specify a valid namespace.")
	}
	if request.Resources == commonapi.AuthorizeRequest_UNASSIGNED_RESOURCES {
		return util.NewInvalidInputError("Resources not specified. Please specify a valid resources.")
	}
	if request.Verb == commonapi.AuthorizeRequest_UNASSIGNED_VERB {
		return util.NewInvalidInputError("Verb not specified. Please specify a valid verb.")
	}
	return nil
}

func NewAuthServer(resourceManager *resource.ResourceManager) *AuthServer {
	return &AuthServer{resourceManager: resourceManager}
}
