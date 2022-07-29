package server

import (
	"context"
	"fmt"
	apiv2beta1 "github.com/kubeflow/pipelines/backend/api/v2beta1/go_client"
	"github.com/kubeflow/pipelines/backend/src/apiserver/resource"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// Metric variables. Please prefix the metric names with experiment_server_.
var (
	// Used to calculate the request rate.
	createExperimentRequests = promauto.NewCounter(prometheus.CounterOpts{
		Name: "experiment_server_v2_create_requests",
		Help: "The total number of CreateExperiment requests",
	})

	listExperimentRequests = promauto.NewCounter(prometheus.CounterOpts{
		Name: "experiment_server_v2_list_requests",
		Help: "The total number of ListExperiments requests",
	})
)

type ExperimentServerOptions struct {
	CollectMetrics bool
}

type ExperimentServer struct {
	apiv2beta1.UnimplementedExperimentServiceServer
	resourceManager *resource.ResourceManager
	options         *ExperimentServerOptions
}

func NewExperimentServer(resourceManager *resource.ResourceManager, options *ExperimentServerOptions) *ExperimentServer {
	return &ExperimentServer{resourceManager: resourceManager, options: options}
}

func (*ExperimentServer) ListExperiment(context.Context, *apiv2beta1.ListExperimentsRequest) (*apiv2beta1.ListExperimentsResponse, error) {
	fmt.Println("v2: List Experiment has been called")
	return &apiv2beta1.ListExperimentsResponse{}, nil
}
