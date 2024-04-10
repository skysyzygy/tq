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

// SectionsGetAllReader is a Reader for the SectionsGetAll structure.
type SectionsGetAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SectionsGetAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSectionsGetAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /ReferenceData/Sections] Sections_GetAll", response, response.Code())
	}
}

// NewSectionsGetAllOK creates a SectionsGetAllOK with default headers values
func NewSectionsGetAllOK() *SectionsGetAllOK {
	return &SectionsGetAllOK{}
}

/*
SectionsGetAllOK describes a response with status code 200, with default header values.

OK
*/
type SectionsGetAllOK struct {
	Payload []*models.Section
}

// IsSuccess returns true when this sections get all o k response has a 2xx status code
func (o *SectionsGetAllOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this sections get all o k response has a 3xx status code
func (o *SectionsGetAllOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this sections get all o k response has a 4xx status code
func (o *SectionsGetAllOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this sections get all o k response has a 5xx status code
func (o *SectionsGetAllOK) IsServerError() bool {
	return false
}

// IsCode returns true when this sections get all o k response a status code equal to that given
func (o *SectionsGetAllOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the sections get all o k response
func (o *SectionsGetAllOK) Code() int {
	return 200
}

func (o *SectionsGetAllOK) Error() string {
	return fmt.Sprintf("[GET /ReferenceData/Sections][%d] sectionsGetAllOK  %+v", 200, o.Payload)
}

func (o *SectionsGetAllOK) String() string {
	return fmt.Sprintf("[GET /ReferenceData/Sections][%d] sectionsGetAllOK  %+v", 200, o.Payload)
}

func (o *SectionsGetAllOK) GetPayload() []*models.Section {
	return o.Payload
}

func (o *SectionsGetAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
