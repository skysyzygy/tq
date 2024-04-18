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

// SalutationsGetAllReader is a Reader for the SalutationsGetAll structure.
type SalutationsGetAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SalutationsGetAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSalutationsGetAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /CRM/Salutations] Salutations_GetAll", response, response.Code())
	}
}

// NewSalutationsGetAllOK creates a SalutationsGetAllOK with default headers values
func NewSalutationsGetAllOK() *SalutationsGetAllOK {
	return &SalutationsGetAllOK{}
}

/*
SalutationsGetAllOK describes a response with status code 200, with default header values.

OK
*/
type SalutationsGetAllOK struct {
	Payload []*models.Salutation
}

// IsSuccess returns true when this salutations get all o k response has a 2xx status code
func (o *SalutationsGetAllOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this salutations get all o k response has a 3xx status code
func (o *SalutationsGetAllOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this salutations get all o k response has a 4xx status code
func (o *SalutationsGetAllOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this salutations get all o k response has a 5xx status code
func (o *SalutationsGetAllOK) IsServerError() bool {
	return false
}

// IsCode returns true when this salutations get all o k response a status code equal to that given
func (o *SalutationsGetAllOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the salutations get all o k response
func (o *SalutationsGetAllOK) Code() int {
	return 200
}

func (o *SalutationsGetAllOK) Error() string {
	return fmt.Sprintf("[GET /CRM/Salutations][%d] salutationsGetAllOK  %+v", 200, o.Payload)
}

func (o *SalutationsGetAllOK) String() string {
	return fmt.Sprintf("[GET /CRM/Salutations][%d] salutationsGetAllOK  %+v", 200, o.Payload)
}

func (o *SalutationsGetAllOK) GetPayload() []*models.Salutation {
	return o.Payload
}

func (o *SalutationsGetAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}