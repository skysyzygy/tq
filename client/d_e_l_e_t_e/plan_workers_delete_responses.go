// Code generated by go-swagger; DO NOT EDIT.

package d_e_l_e_t_e

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PlanWorkersDeleteReader is a Reader for the PlanWorkersDelete structure.
type PlanWorkersDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PlanWorkersDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPlanWorkersDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /Finance/PlanWorkers/{planWorkerId}] PlanWorkers_Delete", response, response.Code())
	}
}

// NewPlanWorkersDeleteNoContent creates a PlanWorkersDeleteNoContent with default headers values
func NewPlanWorkersDeleteNoContent() *PlanWorkersDeleteNoContent {
	return &PlanWorkersDeleteNoContent{}
}

/*
PlanWorkersDeleteNoContent describes a response with status code 204, with default header values.

No Content
*/
type PlanWorkersDeleteNoContent struct {
}

// IsSuccess returns true when this plan workers delete no content response has a 2xx status code
func (o *PlanWorkersDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this plan workers delete no content response has a 3xx status code
func (o *PlanWorkersDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this plan workers delete no content response has a 4xx status code
func (o *PlanWorkersDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this plan workers delete no content response has a 5xx status code
func (o *PlanWorkersDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this plan workers delete no content response a status code equal to that given
func (o *PlanWorkersDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the plan workers delete no content response
func (o *PlanWorkersDeleteNoContent) Code() int {
	return 204
}

func (o *PlanWorkersDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /Finance/PlanWorkers/{planWorkerId}][%d] planWorkersDeleteNoContent ", 204)
}

func (o *PlanWorkersDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /Finance/PlanWorkers/{planWorkerId}][%d] planWorkersDeleteNoContent ", 204)
}

func (o *PlanWorkersDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
