package server_v2beta1

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	apiv2beta1 "github.com/kubeflow/pipelines/backend/api/v2beta1/go_client"
	"github.com/kubeflow/pipelines/backend/src/apiserver/common"
	"github.com/kubeflow/pipelines/backend/src/apiserver/model"
	"github.com/kubeflow/pipelines/backend/src/apiserver/resource"
	"github.com/kubeflow/pipelines/backend/src/common/util"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	authorizationv1 "k8s.io/api/authorization/v1"
)

// Metric variables. Please prefix the metric names with experiment_server_.
// Updated for v2beta1 APIs
var (
	// Used to calculate the request rate.
	createExperimentRequests = promauto.NewCounter(prometheus.CounterOpts{
		Name: "experiment_server_v2beta1_create_requests",
		Help: "The total number of CreateExperiment requests for v2beta1 API",
	})

	getExperimentRequests = promauto.NewCounter(prometheus.CounterOpts{
		Name: "experiment_server_v2beta1_get_requests",
		Help: "The total number of GetExperiment requests for v2beta1 API",
	})

	listExperimentRequests = promauto.NewCounter(prometheus.CounterOpts{
		Name: "experiment_server_v2beta1_list_requests",
		Help: "The total number of ListExperiments requests for v2beta1 API",
	})

	deleteExperimentRequests = promauto.NewCounter(prometheus.CounterOpts{
		Name: "experiment_server_v2beta1_delete_requests",
		Help: "The total number of DeleteExperiment requests for v2beta1 API",
	})

	archiveExperimentRequests = promauto.NewCounter(prometheus.CounterOpts{
		Name: "experiment_server_v2beta1_archive_requests",
		Help: "The total number of ArchiveExperiment requests for v2beta1 API",
	})

	unarchiveExperimentRequests = promauto.NewCounter(prometheus.CounterOpts{
		Name: "experiment_server_v2beta1_unarchive_requests",
		Help: "The total number of UnarchiveExperiment requests for v2beta1 API",
	})

	// TODO(jingzhang36): error count and success count.

	experimentCount = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "experiment_server_v2beta1_run_count",
		Help: "The current number of experiments in Kubeflow Pipelines instance",
	})
)

type ExperimentServerOptions struct {
	CollectMetrics bool
}

type ExperimentServer struct {
	apiv2beta1.UnimplementedExperimentServiceServer
}

// CreateExperiment
func (s *ExperimentServer) CreateExperiment(ctx context.Context, request *apiv2beta1.CreateExperimentRequest) (
	*apiv2beta1.Experiment, error) {
	if s.options.CollectMetrics {
		createExperimentRequests.Inc()
	}

	err := ValidateCreateExperimentRequest(request)
	if err != nil {
		return nil, util.Wrap(err, "Validate experiment request failed.")
	}

	resourceAttributes := &authorizationv1.ResourceAttributes{
		Namespace: common.GetNamespaceFromAPIResourceReferences(request.Experiment.ResourceReferences),
		Verb:      common.RbacResourceVerbCreate,
		Name:      request.Experiment.Name,
	}
	err = s.canAccessExperiment(ctx, "", resourceAttributes)
	if err != nil {
		return nil, util.Wrap(err, "Failed to authorize the request")
	}

	newExperiment, err := s.resourceManager.CreateExperiment(request.Experiment)
	if err != nil {
		return nil, util.Wrap(err, "Create experiment failed.")
	}

	if s.options.CollectMetrics {
		experimentCount.Inc()
	}

	// ToApiExperiment needs change
	return ToApiExperiment(newExperiment), nil
}

// GetExperiment

// ListExperiment

// DeleteExperiment

// ValidateCreateExperimentRequest

// canAccessExperiment

// ArchiveExperiment

// UnarchiveExperiment

// NewExperimentServer
func NewExperimentServer(resourceManager *resource.ResourceManager, options *ExperimentServerOptions) *ExperimentServer {
	return &ExperimentServer{resourceManager: resourceManager, options: options}
}

func (*ExperimentServer) ListExperiment(context.Context, *apiv2beta1.ListExperimentsRequest) (*apiv2beta1.ListExperimentsResponse, error) {
	fmt.Println("v2: List Experiment has been called")
	return &apiv2beta1.ListExperimentsResponse{}, nil
}
