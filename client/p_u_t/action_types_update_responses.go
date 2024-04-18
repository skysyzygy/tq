// Code generated by go-swagger; DO NOT EDIT.

package p_u_t

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skysyzygy/tq/models"
)

// ActionTypesUpdateReader is a Reader for the ActionTypesUpdate structure.
type ActionTypesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ActionTypesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewActionTypesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /ReferenceData/ActionTypes/{id}] ActionTypes_Update", response, response.Code())
	}
}

// NewActionTypesUpdateOK creates a ActionTypesUpdateOK with default headers values
func NewActionTypesUpdateOK() *ActionTypesUpdateOK {
	return &ActionTypesUpdateOK{}
}

/*
ActionTypesUpdateOK describes a response with status code 200, with default header values.

OK
*/
type ActionTypesUpdateOK struct {
	Payload *models.ActionType
}

// IsSuccess returns true when this action types update o k response has a 2xx status code
func (o *ActionTypesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this action types update o k response has a 3xx status code
func (o *ActionTypesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this action types update o k response has a 4xx status code
func (o *ActionTypesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this action types update o k response has a 5xx status code
func (o *ActionTypesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this action types update o k response a status code equal to that given
func (o *ActionTypesUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the action types update o k response
func (o *ActionTypesUpdateOK) Code() int {
	return 200
}

func (o *ActionTypesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /ReferenceData/ActionTypes/{id}][%d] actionTypesUpdateOK  %+v", 200, o.Payload)
}

func (o *ActionTypesUpdateOK) String() string {
	return fmt.Sprintf("[PUT /ReferenceData/ActionTypes/{id}][%d] actionTypesUpdateOK  %+v", 200, o.Payload)
}

func (o *ActionTypesUpdateOK) GetPayload() *models.ActionType {
	return o.Payload
}

func (o *ActionTypesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ActionType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}