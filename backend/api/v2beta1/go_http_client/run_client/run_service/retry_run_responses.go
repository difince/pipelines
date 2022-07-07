// Code generated by go-swagger; DO NOT EDIT.

package run_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	run_model "github.com/kubeflow/pipelines/backend/api/v2beta1/go_http_client/run_model"
)

// RetryRunReader is a Reader for the RetryRun structure.
type RetryRunReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RetryRunReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRetryRunOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewRetryRunDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRetryRunOK creates a RetryRunOK with default headers values
func NewRetryRunOK() *RetryRunOK {
	return &RetryRunOK{}
}

/*RetryRunOK handles this case with default header values.

A successful response.
*/
type RetryRunOK struct {
	Payload interface{}
}

func (o *RetryRunOK) Error() string {
	return fmt.Sprintf("[POST /apis/v2beta1/runs/{run_id}/retry][%d] retryRunOK  %+v", 200, o.Payload)
}

func (o *RetryRunOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRetryRunDefault creates a RetryRunDefault with default headers values
func NewRetryRunDefault(code int) *RetryRunDefault {
	return &RetryRunDefault{
		_statusCode: code,
	}
}

/*RetryRunDefault handles this case with default header values.

RetryRunDefault retry run default
*/
type RetryRunDefault struct {
	_statusCode int

	Payload *run_model.Apiv2beta1Status
}

// Code gets the status code for the retry run default response
func (o *RetryRunDefault) Code() int {
	return o._statusCode
}

func (o *RetryRunDefault) Error() string {
	return fmt.Sprintf("[POST /apis/v2beta1/runs/{run_id}/retry][%d] RetryRun default  %+v", o._statusCode, o.Payload)
}

func (o *RetryRunDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(run_model.Apiv2beta1Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
