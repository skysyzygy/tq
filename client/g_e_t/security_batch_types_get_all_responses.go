// Code generated by go-swagger; DO NOT EDIT.

package g_e_t

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skysyzygy/tq/models"
)

// SecurityBatchTypesGetAllReader is a Reader for the SecurityBatchTypesGetAll structure.
type SecurityBatchTypesGetAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecurityBatchTypesGetAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSecurityBatchTypesGetAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /Security/BatchTypes] SecurityBatchTypes_GetAll", response, response.Code())
	}
}

// NewSecurityBatchTypesGetAllOK creates a SecurityBatchTypesGetAllOK with default headers values
func NewSecurityBatchTypesGetAllOK() *SecurityBatchTypesGetAllOK {
	return &SecurityBatchTypesGetAllOK{}
}

/*
SecurityBatchTypesGetAllOK describes a response with status code 200, with default header values.

OK
*/
type SecurityBatchTypesGetAllOK struct {
	Payload []*models.BatchTypeUserGroup
}

// IsSuccess returns true when this security batch types get all o k response has a 2xx status code
func (o *SecurityBatchTypesGetAllOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this security batch types get all o k response has a 3xx status code
func (o *SecurityBatchTypesGetAllOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this security batch types get all o k response has a 4xx status code
func (o *SecurityBatchTypesGetAllOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this security batch types get all o k response has a 5xx status code
func (o *SecurityBatchTypesGetAllOK) IsServerError() bool {
	return false
}

// IsCode returns true when this security batch types get all o k response a status code equal to that given
func (o *SecurityBatchTypesGetAllOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the security batch types get all o k response
func (o *SecurityBatchTypesGetAllOK) Code() int {
	return 200
}

func (o *SecurityBatchTypesGetAllOK) Error() string {
	return fmt.Sprintf("[GET /Security/BatchTypes][%d] securityBatchTypesGetAllOK  %+v", 200, o.Payload)
}

func (o *SecurityBatchTypesGetAllOK) String() string {
	return fmt.Sprintf("[GET /Security/BatchTypes][%d] securityBatchTypesGetAllOK  %+v", 200, o.Payload)
}

func (o *SecurityBatchTypesGetAllOK) GetPayload() []*models.BatchTypeUserGroup {
	return o.Payload
}

func (o *SecurityBatchTypesGetAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}