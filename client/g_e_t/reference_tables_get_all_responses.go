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

// ReferenceTablesGetAllReader is a Reader for the ReferenceTablesGetAll structure.
type ReferenceTablesGetAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReferenceTablesGetAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReferenceTablesGetAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /ReferenceData/ReferenceTables] ReferenceTables_GetAll", response, response.Code())
	}
}

// NewReferenceTablesGetAllOK creates a ReferenceTablesGetAllOK with default headers values
func NewReferenceTablesGetAllOK() *ReferenceTablesGetAllOK {
	return &ReferenceTablesGetAllOK{}
}

/*
ReferenceTablesGetAllOK describes a response with status code 200, with default header values.

OK
*/
type ReferenceTablesGetAllOK struct {
	Payload []*models.ReferenceTable
}

// IsSuccess returns true when this reference tables get all o k response has a 2xx status code
func (o *ReferenceTablesGetAllOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this reference tables get all o k response has a 3xx status code
func (o *ReferenceTablesGetAllOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this reference tables get all o k response has a 4xx status code
func (o *ReferenceTablesGetAllOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this reference tables get all o k response has a 5xx status code
func (o *ReferenceTablesGetAllOK) IsServerError() bool {
	return false
}

// IsCode returns true when this reference tables get all o k response a status code equal to that given
func (o *ReferenceTablesGetAllOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the reference tables get all o k response
func (o *ReferenceTablesGetAllOK) Code() int {
	return 200
}

func (o *ReferenceTablesGetAllOK) Error() string {
	return fmt.Sprintf("[GET /ReferenceData/ReferenceTables][%d] referenceTablesGetAllOK  %+v", 200, o.Payload)
}

func (o *ReferenceTablesGetAllOK) String() string {
	return fmt.Sprintf("[GET /ReferenceData/ReferenceTables][%d] referenceTablesGetAllOK  %+v", 200, o.Payload)
}

func (o *ReferenceTablesGetAllOK) GetPayload() []*models.ReferenceTable {
	return o.Payload
}

func (o *ReferenceTablesGetAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}