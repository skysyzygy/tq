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

// ListsGenerateReader is a Reader for the ListsGenerate structure.
type ListsGenerateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListsGenerateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListsGenerateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /Reporting/Lists/{listId}/Generate] Lists_Generate", response, response.Code())
	}
}

// NewListsGenerateOK creates a ListsGenerateOK with default headers values
func NewListsGenerateOK() *ListsGenerateOK {
	return &ListsGenerateOK{}
}

/*
ListsGenerateOK describes a response with status code 200, with default header values.

OK
*/
type ListsGenerateOK struct {
	Payload *models.List
}

// IsSuccess returns true when this lists generate o k response has a 2xx status code
func (o *ListsGenerateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this lists generate o k response has a 3xx status code
func (o *ListsGenerateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this lists generate o k response has a 4xx status code
func (o *ListsGenerateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this lists generate o k response has a 5xx status code
func (o *ListsGenerateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this lists generate o k response a status code equal to that given
func (o *ListsGenerateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the lists generate o k response
func (o *ListsGenerateOK) Code() int {
	return 200
}

func (o *ListsGenerateOK) Error() string {
	return fmt.Sprintf("[POST /Reporting/Lists/{listId}/Generate][%d] listsGenerateOK  %+v", 200, o.Payload)
}

func (o *ListsGenerateOK) String() string {
	return fmt.Sprintf("[POST /Reporting/Lists/{listId}/Generate][%d] listsGenerateOK  %+v", 200, o.Payload)
}

func (o *ListsGenerateOK) GetPayload() *models.List {
	return o.Payload
}

func (o *ListsGenerateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.List)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}