// Code generated by go-swagger; DO NOT EDIT.

package pipeline_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeletePipelineByNameParams creates a new DeletePipelineByNameParams object
// with the default values initialized.
func NewDeletePipelineByNameParams() *DeletePipelineByNameParams {
	var ()
	return &DeletePipelineByNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePipelineByNameParamsWithTimeout creates a new DeletePipelineByNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeletePipelineByNameParamsWithTimeout(timeout time.Duration) *DeletePipelineByNameParams {
	var ()
	return &DeletePipelineByNameParams{

		timeout: timeout,
	}
}

// NewDeletePipelineByNameParamsWithContext creates a new DeletePipelineByNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeletePipelineByNameParamsWithContext(ctx context.Context) *DeletePipelineByNameParams {
	var ()
	return &DeletePipelineByNameParams{

		Context: ctx,
	}
}

// NewDeletePipelineByNameParamsWithHTTPClient creates a new DeletePipelineByNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeletePipelineByNameParamsWithHTTPClient(client *http.Client) *DeletePipelineByNameParams {
	var ()
	return &DeletePipelineByNameParams{
		HTTPClient: client,
	}
}

/*DeletePipelineByNameParams contains all the parameters to send to the API endpoint
for the delete pipeline by name operation typically these are written to a http.Request
*/
type DeletePipelineByNameParams struct {

	/*Name
	  The Name of the pipeline to be deleted.

	*/
	Name string
	/*Namespace
	  The Namespace the pipeline belongs to.
	In the case of shared pipelines and KFPipeline standalone installation,
	the pipeline name is the only needed field for unique resource lookup..
	In those cases, please provide a hyphen (dash character, "-") for the namespace.

	*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete pipeline by name params
func (o *DeletePipelineByNameParams) WithTimeout(timeout time.Duration) *DeletePipelineByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete pipeline by name params
func (o *DeletePipelineByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete pipeline by name params
func (o *DeletePipelineByNameParams) WithContext(ctx context.Context) *DeletePipelineByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete pipeline by name params
func (o *DeletePipelineByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete pipeline by name params
func (o *DeletePipelineByNameParams) WithHTTPClient(client *http.Client) *DeletePipelineByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete pipeline by name params
func (o *DeletePipelineByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the delete pipeline by name params
func (o *DeletePipelineByNameParams) WithName(name string) *DeletePipelineByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete pipeline by name params
func (o *DeletePipelineByNameParams) SetName(name string) {
	o.Name = name
}

// WithNamespace adds the namespace to the delete pipeline by name params
func (o *DeletePipelineByNameParams) WithNamespace(namespace string) *DeletePipelineByNameParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the delete pipeline by name params
func (o *DeletePipelineByNameParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePipelineByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
