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

// QueryElementsGetReader is a Reader for the QueryElementsGet structure.
type QueryElementsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryElementsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryElementsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /Reporting/QueryElements/{id}] QueryElements_Get", response, response.Code())
	}
}

// NewQueryElementsGetOK creates a QueryElementsGetOK with default headers values
func NewQueryElementsGetOK() *QueryElementsGetOK {
	return &QueryElementsGetOK{}
}

/*
QueryElementsGetOK describes a response with status code 200, with default header values.

OK
*/
type QueryElementsGetOK struct {
	Payload *models.QueryElement
}

// IsSuccess returns true when this query elements get o k response has a 2xx status code
func (o *QueryElementsGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this query elements get o k response has a 3xx status code
func (o *QueryElementsGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query elements get o k response has a 4xx status code
func (o *QueryElementsGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this query elements get o k response has a 5xx status code
func (o *QueryElementsGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this query elements get o k response a status code equal to that given
func (o *QueryElementsGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the query elements get o k response
func (o *QueryElementsGetOK) Code() int {
	return 200
}

func (o *QueryElementsGetOK) Error() string {
	return fmt.Sprintf("[GET /Reporting/QueryElements/{id}][%d] queryElementsGetOK  %+v", 200, o.Payload)
}

func (o *QueryElementsGetOK) String() string {
	return fmt.Sprintf("[GET /Reporting/QueryElements/{id}][%d] queryElementsGetOK  %+v", 200, o.Payload)
}

func (o *QueryElementsGetOK) GetPayload() *models.QueryElement {
	return o.Payload
}

func (o *QueryElementsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.QueryElement)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}