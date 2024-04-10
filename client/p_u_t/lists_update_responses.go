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

// ListsUpdateReader is a Reader for the ListsUpdate structure.
type ListsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /Reporting/Lists/{listId}] Lists_Update", response, response.Code())
	}
}

// NewListsUpdateOK creates a ListsUpdateOK with default headers values
func NewListsUpdateOK() *ListsUpdateOK {
	return &ListsUpdateOK{}
}

/*
ListsUpdateOK describes a response with status code 200, with default header values.

OK
*/
type ListsUpdateOK struct {
	Payload *models.List
}

// IsSuccess returns true when this lists update o k response has a 2xx status code
func (o *ListsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this lists update o k response has a 3xx status code
func (o *ListsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this lists update o k response has a 4xx status code
func (o *ListsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this lists update o k response has a 5xx status code
func (o *ListsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this lists update o k response a status code equal to that given
func (o *ListsUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the lists update o k response
func (o *ListsUpdateOK) Code() int {
	return 200
}

func (o *ListsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /Reporting/Lists/{listId}][%d] listsUpdateOK  %+v", 200, o.Payload)
}

func (o *ListsUpdateOK) String() string {
	return fmt.Sprintf("[PUT /Reporting/Lists/{listId}][%d] listsUpdateOK  %+v", 200, o.Payload)
}

func (o *ListsUpdateOK) GetPayload() *models.List {
	return o.Payload
}

func (o *ListsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.List)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
