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

// SourceGroupsCreateReader is a Reader for the SourceGroupsCreate structure.
type SourceGroupsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SourceGroupsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSourceGroupsCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /ReferenceData/SourceGroups] SourceGroups_Create", response, response.Code())
	}
}

// NewSourceGroupsCreateOK creates a SourceGroupsCreateOK with default headers values
func NewSourceGroupsCreateOK() *SourceGroupsCreateOK {
	return &SourceGroupsCreateOK{}
}

/*
SourceGroupsCreateOK describes a response with status code 200, with default header values.

OK
*/
type SourceGroupsCreateOK struct {
	Payload *models.SourceGroup
}

// IsSuccess returns true when this source groups create o k response has a 2xx status code
func (o *SourceGroupsCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this source groups create o k response has a 3xx status code
func (o *SourceGroupsCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this source groups create o k response has a 4xx status code
func (o *SourceGroupsCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this source groups create o k response has a 5xx status code
func (o *SourceGroupsCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this source groups create o k response a status code equal to that given
func (o *SourceGroupsCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the source groups create o k response
func (o *SourceGroupsCreateOK) Code() int {
	return 200
}

func (o *SourceGroupsCreateOK) Error() string {
	return fmt.Sprintf("[POST /ReferenceData/SourceGroups][%d] sourceGroupsCreateOK  %+v", 200, o.Payload)
}

func (o *SourceGroupsCreateOK) String() string {
	return fmt.Sprintf("[POST /ReferenceData/SourceGroups][%d] sourceGroupsCreateOK  %+v", 200, o.Payload)
}

func (o *SourceGroupsCreateOK) GetPayload() *models.SourceGroup {
	return o.Payload
}

func (o *SourceGroupsCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SourceGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}