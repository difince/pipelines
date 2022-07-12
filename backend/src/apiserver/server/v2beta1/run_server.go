package v2beta1

import (
	"context"
	"errors"
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	authorizationv1 "k8s.io/api/authorization/v1"

	apiv2beta1 "github.com/kubeflow/pipelines/backend/api/v2beta1/go_client"
	"github.com/kubeflow/pipelines/backend/src/apiserver/common"
	"github.com/kubeflow/pipelines/backend/src/apiserver/model"
	"github.com/kubeflow/pipelines/backend/src/apiserver/resource"
	"github.com/kubeflow/pipelines/backend/src/apiserver/server"
	"github.com/kubeflow/pipelines/backend/src/common/util"
)

type RunServerOptions struct {
	CollectMetrics bool
}

var (
	listRunRequests = promauto.NewCounter(prometheus.CounterOpts{
		Name: "run_server_v2beta1_list_requests",
		Help: "The total number of v2beta1.ListRuns requests",
	})
)

type RunServer struct {
	// TODO Helper interface - to make our current partial implementation of RunServiceInterface
	// still compatible. To be removed once all methods are implemented.
	apiv2beta1.UnimplementedRunServiceServer

	resourceManager *resource.ResourceManager
	options         *RunServerOptions
}

func (s *RunServer) GetRun(context.Context, *apiv2beta1.GetRunRequest) (*apiv2beta1.RunDetail, error) {
	fmt.Println("V2: GetRun called")
	return &apiv2beta1.RunDetail{}, nil
}

func (s *RunServer) ListRuns(ctx context.Context, request *apiv2beta1.ListRunsRequest) (*apiv2beta1.ListRunsResponse, error) {
	if s.options.CollectMetrics {
		listRunRequests.Inc()
	}

	opts, err := server.ValidatedListOptions(&model.Run{}, request.PageToken, int(request.PageSize), request.SortBy, request.Filter)
	if err != nil {
		return nil, util.Wrap(err, "Failed to create list options")
	}

	if common.IsMultiUserMode() {
		namespace, err := s.resourceManager.GetNamespaceFromExperimentID(request.ExperimentId)
		if err != nil {
			return nil, util.Wrap(err, "Run's experiment has no namespace.")
		}
		resourceAttributes := &authorizationv1.ResourceAttributes{
			Namespace: namespace,
			Verb:      common.RbacResourceVerbList,
		}
		err = s.canAccessRun(ctx, "", resourceAttributes)
		if err != nil {
			return nil, util.Wrap(err, "Failed to authorize with namespace in experiment resource reference.")
		}
	}
	filterContext := &common.FilterContext{ReferenceKey: &common.ReferenceKey{Type: common.Experiment, ID: request.GetExperimentId()}}
	runs, total_size, nextPageToken, err := s.resourceManager.ListRuns(filterContext, opts)
	if err != nil {
		return nil, util.Wrap(err, "Failed to list runs.")
	}
	return &apiv2beta1.ListRunsResponse{Runs: ToApiRuns(runs), TotalSize: int32(total_size), NextPageToken: nextPageToken}, nil
}

//This is a copy of v1.run_server method!
func (s *RunServer) canAccessRun(ctx context.Context, runId string, resourceAttributes *authorizationv1.ResourceAttributes) error {
	if common.IsMultiUserMode() == false {
		// Skip authz if not multi-user mode.
		return nil
	}

	if len(runId) > 0 {
		runDetail, err := s.resourceManager.GetRun(runId)
		if err != nil {
			return util.Wrap(err, "Failed to authorize with the experiment ID.")
		}
		if len(resourceAttributes.Namespace) == 0 {
			if len(runDetail.Namespace) == 0 {
				return util.NewInternalServerError(
					errors.New("empty namespace"),
					"The run doesn't have a valid namespace.",
				)
			}
			resourceAttributes.Namespace = runDetail.Namespace
		}
		if len(resourceAttributes.Name) == 0 {
			resourceAttributes.Name = runDetail.Name
		}
	}

	resourceAttributes.Group = common.RbacPipelinesGroup
	resourceAttributes.Version = common.RbacPipelinesVersion
	resourceAttributes.Resource = common.RbacResourceTypeRuns

	err := server.IsAuthorized(s.resourceManager, ctx, resourceAttributes)
	if err != nil {
		return util.Wrap(err, "Failed to authorize with API resource references")
	}
	return nil
}

func NewRunServer(resourceManager *resource.ResourceManager, options *RunServerOptions) *RunServer {
	return &RunServer{resourceManager: resourceManager, options: options}
}
