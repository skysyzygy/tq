// Code generated by go-swagger; DO NOT EDIT.

package p_o_s_t

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skysyzygy/tq/models"
)

// QueryElementGroupsCreateReader is a Reader for the QueryElementGroupsCreate structure.
type QueryElementGroupsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryElementGroupsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryElementGroupsCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /Reporting/QueryElementGroups] QueryElementGroups_Create", response, response.Code())
	}
}

// NewQueryElementGroupsCreateOK creates a QueryElementGroupsCreateOK with default headers values
func NewQueryElementGroupsCreateOK() *QueryElementGroupsCreateOK {
	return &QueryElementGroupsCreateOK{}
}

/*
QueryElementGroupsCreateOK describes a response with status code 200, with default header values.

OK
*/
type QueryElementGroupsCreateOK struct {
	Payload *models.QueryElementGroup
}

// IsSuccess returns true when this query element groups create o k response has a 2xx status code
func (o *QueryElementGroupsCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this query element groups create o k response has a 3xx status code
func (o *QueryElementGroupsCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query element groups create o k response has a 4xx status code
func (o *QueryElementGroupsCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this query element groups create o k response has a 5xx status code
func (o *QueryElementGroupsCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this query element groups create o k response a status code equal to that given
func (o *QueryElementGroupsCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the query element groups create o k response
func (o *QueryElementGroupsCreateOK) Code() int {
	return 200
}

func (o *QueryElementGroupsCreateOK) Error() string {
	return fmt.Sprintf("[POST /Reporting/QueryElementGroups][%d] queryElementGroupsCreateOK  %+v", 200, o.Payload)
}

func (o *QueryElementGroupsCreateOK) String() string {
	return fmt.Sprintf("[POST /Reporting/QueryElementGroups][%d] queryElementGroupsCreateOK  %+v", 200, o.Payload)
}

func (o *QueryElementGroupsCreateOK) GetPayload() *models.QueryElementGroup {
	return o.Payload
}

func (o *QueryElementGroupsCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.QueryElementGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
