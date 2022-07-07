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

// NewGetPipelineByNameParams creates a new GetPipelineByNameParams object
// with the default values initialized.
func NewGetPipelineByNameParams() *GetPipelineByNameParams {
	var ()
	return &GetPipelineByNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPipelineByNameParamsWithTimeout creates a new GetPipelineByNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPipelineByNameParamsWithTimeout(timeout time.Duration) *GetPipelineByNameParams {
	var ()
	return &GetPipelineByNameParams{

		timeout: timeout,
	}
}

// NewGetPipelineByNameParamsWithContext creates a new GetPipelineByNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPipelineByNameParamsWithContext(ctx context.Context) *GetPipelineByNameParams {
	var ()
	return &GetPipelineByNameParams{

		Context: ctx,
	}
}

// NewGetPipelineByNameParamsWithHTTPClient creates a new GetPipelineByNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPipelineByNameParamsWithHTTPClient(client *http.Client) *GetPipelineByNameParams {
	var ()
	return &GetPipelineByNameParams{
		HTTPClient: client,
	}
}

/*GetPipelineByNameParams contains all the parameters to send to the API endpoint
for the get pipeline by name operation typically these are written to a http.Request
*/
type GetPipelineByNameParams struct {

	/*Name
	  The Name of the pipeline to be retrieved.

	*/
	Name string
	/*Namespace
	  The Namespace the pipeline belongs to.
	In the case of shared pipelines and KFPipeline standalone installation `namespace` should be skipped.

	*/
	Namespace *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get pipeline by name params
func (o *GetPipelineByNameParams) WithTimeout(timeout time.Duration) *GetPipelineByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get pipeline by name params
func (o *GetPipelineByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get pipeline by name params
func (o *GetPipelineByNameParams) WithContext(ctx context.Context) *GetPipelineByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get pipeline by name params
func (o *GetPipelineByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get pipeline by name params
func (o *GetPipelineByNameParams) WithHTTPClient(client *http.Client) *GetPipelineByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get pipeline by name params
func (o *GetPipelineByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get pipeline by name params
func (o *GetPipelineByNameParams) WithName(name string) *GetPipelineByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get pipeline by name params
func (o *GetPipelineByNameParams) SetName(name string) {
	o.Name = name
}

// WithNamespace adds the namespace to the get pipeline by name params
func (o *GetPipelineByNameParams) WithNamespace(namespace *string) *GetPipelineByNameParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get pipeline by name params
func (o *GetPipelineByNameParams) SetNamespace(namespace *string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *GetPipelineByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
