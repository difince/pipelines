// Code generated by go-swagger; DO NOT EDIT.

package run_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewListRunsParams creates a new ListRunsParams object
// with the default values initialized.
func NewListRunsParams() *ListRunsParams {
	var ()
	return &ListRunsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListRunsParamsWithTimeout creates a new ListRunsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListRunsParamsWithTimeout(timeout time.Duration) *ListRunsParams {
	var ()
	return &ListRunsParams{

		timeout: timeout,
	}
}

// NewListRunsParamsWithContext creates a new ListRunsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListRunsParamsWithContext(ctx context.Context) *ListRunsParams {
	var ()
	return &ListRunsParams{

		Context: ctx,
	}
}

// NewListRunsParamsWithHTTPClient creates a new ListRunsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListRunsParamsWithHTTPClient(client *http.Client) *ListRunsParams {
	var ()
	return &ListRunsParams{
		HTTPClient: client,
	}
}

/*ListRunsParams contains all the parameters to send to the API endpoint
for the list runs operation typically these are written to a http.Request
*/
type ListRunsParams struct {

	/*ExperimentID
	  TODO Diana: to check this option
	Filter Lists based on the Experiment they belong to.
	Optional fields when single-user mode is enabled (Pipeline standalone installation).
	In this case, please provide "hyphen" instead ("dash",  "-")

	*/
	ExperimentID string
	/*Filter
	  A url-encoded, JSON-serialized Filter protocol buffer (see
	[filter.proto](https://github.com/kubeflow/pipelines/blob/master/backend/api/filter.proto)).

	*/
	Filter *string
	/*Namespace
	  Filter Lists based on Namespace they belong to.

	*/
	Namespace *string
	/*PageSize
	  The number of runs to be listed per page. If there are more runs than this
	number, the response message will contain a nextPageToken field you can use
	to fetch the next page.

	*/
	PageSize *int32
	/*PageToken
	  A page token to request the next page of results. The token is acquried
	from the nextPageToken field of the response from the previous
	ListRuns call or can be omitted when fetching the first page.

	*/
	PageToken *string
	/*PipelineVersionID
	  Optional input Parameter. Filter based on the Pipeline_version.

	*/
	PipelineVersionID *string
	/*SortBy
	  Can be format of "field_name", "field_name asc" or "field_name desc"
	(Example, "name asc" or "id desc"). Ascending by default.

	*/
	SortBy *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list runs params
func (o *ListRunsParams) WithTimeout(timeout time.Duration) *ListRunsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list runs params
func (o *ListRunsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list runs params
func (o *ListRunsParams) WithContext(ctx context.Context) *ListRunsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list runs params
func (o *ListRunsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list runs params
func (o *ListRunsParams) WithHTTPClient(client *http.Client) *ListRunsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list runs params
func (o *ListRunsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExperimentID adds the experimentID to the list runs params
func (o *ListRunsParams) WithExperimentID(experimentID string) *ListRunsParams {
	o.SetExperimentID(experimentID)
	return o
}

// SetExperimentID adds the experimentId to the list runs params
func (o *ListRunsParams) SetExperimentID(experimentID string) {
	o.ExperimentID = experimentID
}

// WithFilter adds the filter to the list runs params
func (o *ListRunsParams) WithFilter(filter *string) *ListRunsParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the list runs params
func (o *ListRunsParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithNamespace adds the namespace to the list runs params
func (o *ListRunsParams) WithNamespace(namespace *string) *ListRunsParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the list runs params
func (o *ListRunsParams) SetNamespace(namespace *string) {
	o.Namespace = namespace
}

// WithPageSize adds the pageSize to the list runs params
func (o *ListRunsParams) WithPageSize(pageSize *int32) *ListRunsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the list runs params
func (o *ListRunsParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithPageToken adds the pageToken to the list runs params
func (o *ListRunsParams) WithPageToken(pageToken *string) *ListRunsParams {
	o.SetPageToken(pageToken)
	return o
}

// SetPageToken adds the pageToken to the list runs params
func (o *ListRunsParams) SetPageToken(pageToken *string) {
	o.PageToken = pageToken
}

// WithPipelineVersionID adds the pipelineVersionID to the list runs params
func (o *ListRunsParams) WithPipelineVersionID(pipelineVersionID *string) *ListRunsParams {
	o.SetPipelineVersionID(pipelineVersionID)
	return o
}

// SetPipelineVersionID adds the pipelineVersionId to the list runs params
func (o *ListRunsParams) SetPipelineVersionID(pipelineVersionID *string) {
	o.PipelineVersionID = pipelineVersionID
}

// WithSortBy adds the sortBy to the list runs params
func (o *ListRunsParams) WithSortBy(sortBy *string) *ListRunsParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the list runs params
func (o *ListRunsParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WriteToRequest writes these params to a swagger request
func (o *ListRunsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param experiment_id
	if err := r.SetPathParam("experiment_id", o.ExperimentID); err != nil {
		return err
	}

	if o.Filter != nil {

		// query param filter
		var qrFilter string
		if o.Filter != nil {
			qrFilter = *o.Filter
		}
		qFilter := qrFilter
		if qFilter != "" {
			if err := r.SetQueryParam("filter", qFilter); err != nil {
				return err
			}
		}

	}

	if o.Namespace != nil {

		// query param namespace
		var qrNamespace string
		if o.Namespace != nil {
			qrNamespace = *o.Namespace
		}
		qNamespace := qrNamespace
		if qNamespace != "" {
			if err := r.SetQueryParam("namespace", qNamespace); err != nil {
				return err
			}
		}

	}

	if o.PageSize != nil {

		// query param page_size
		var qrPageSize int32
		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt32(qrPageSize)
		if qPageSize != "" {
			if err := r.SetQueryParam("page_size", qPageSize); err != nil {
				return err
			}
		}

	}

	if o.PageToken != nil {

		// query param page_token
		var qrPageToken string
		if o.PageToken != nil {
			qrPageToken = *o.PageToken
		}
		qPageToken := qrPageToken
		if qPageToken != "" {
			if err := r.SetQueryParam("page_token", qPageToken); err != nil {
				return err
			}
		}

	}

	if o.PipelineVersionID != nil {

		// query param pipeline_version_id
		var qrPipelineVersionID string
		if o.PipelineVersionID != nil {
			qrPipelineVersionID = *o.PipelineVersionID
		}
		qPipelineVersionID := qrPipelineVersionID
		if qPipelineVersionID != "" {
			if err := r.SetQueryParam("pipeline_version_id", qPipelineVersionID); err != nil {
				return err
			}
		}

	}

	if o.SortBy != nil {

		// query param sort_by
		var qrSortBy string
		if o.SortBy != nil {
			qrSortBy = *o.SortBy
		}
		qSortBy := qrSortBy
		if qSortBy != "" {
			if err := r.SetQueryParam("sort_by", qSortBy); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
