package v2beta1

import (
	"context"
	"fmt"
	apiv2beta1 "github.com/kubeflow/pipelines/backend/api/v2beta1/go_client"
)

type ExperimentServer struct {
	apiv2beta1.UnimplementedExperimentServiceServer
}

func (e *ExperimentServer) ListExperiments(ctx context.Context, request *apiv2beta1.ListExperimentsRequest) (*apiv2beta1.ListExperimentsResponse, error) {
	fmt.Println("v2: List Experiment has been called")
	return &apiv2beta1.ListExperimentsResponse{}, nil
}

func NewExperimentServer() *ExperimentServer {
	return &ExperimentServer{}
}
