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

// AssociationsGetAllReader is a Reader for the AssociationsGetAll structure.
type AssociationsGetAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AssociationsGetAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAssociationsGetAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /CRM/Associations] Associations_GetAll", response, response.Code())
	}
}

// NewAssociationsGetAllOK creates a AssociationsGetAllOK with default headers values
func NewAssociationsGetAllOK() *AssociationsGetAllOK {
	return &AssociationsGetAllOK{}
}

/*
AssociationsGetAllOK describes a response with status code 200, with default header values.

OK
*/
type AssociationsGetAllOK struct {
	Payload []*models.Association
}

// IsSuccess returns true when this associations get all o k response has a 2xx status code
func (o *AssociationsGetAllOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this associations get all o k response has a 3xx status code
func (o *AssociationsGetAllOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this associations get all o k response has a 4xx status code
func (o *AssociationsGetAllOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this associations get all o k response has a 5xx status code
func (o *AssociationsGetAllOK) IsServerError() bool {
	return false
}

// IsCode returns true when this associations get all o k response a status code equal to that given
func (o *AssociationsGetAllOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the associations get all o k response
func (o *AssociationsGetAllOK) Code() int {
	return 200
}

func (o *AssociationsGetAllOK) Error() string {
	return fmt.Sprintf("[GET /CRM/Associations][%d] associationsGetAllOK  %+v", 200, o.Payload)
}

func (o *AssociationsGetAllOK) String() string {
	return fmt.Sprintf("[GET /CRM/Associations][%d] associationsGetAllOK  %+v", 200, o.Payload)
}

func (o *AssociationsGetAllOK) GetPayload() []*models.Association {
	return o.Payload
}

func (o *AssociationsGetAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
