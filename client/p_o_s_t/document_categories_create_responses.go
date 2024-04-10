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

// DocumentCategoriesCreateReader is a Reader for the DocumentCategoriesCreate structure.
type DocumentCategoriesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DocumentCategoriesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDocumentCategoriesCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /ReferenceData/DocumentCategories] DocumentCategories_Create", response, response.Code())
	}
}

// NewDocumentCategoriesCreateOK creates a DocumentCategoriesCreateOK with default headers values
func NewDocumentCategoriesCreateOK() *DocumentCategoriesCreateOK {
	return &DocumentCategoriesCreateOK{}
}

/*
DocumentCategoriesCreateOK describes a response with status code 200, with default header values.

OK
*/
type DocumentCategoriesCreateOK struct {
	Payload *models.DocumentCategory
}

// IsSuccess returns true when this document categories create o k response has a 2xx status code
func (o *DocumentCategoriesCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this document categories create o k response has a 3xx status code
func (o *DocumentCategoriesCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this document categories create o k response has a 4xx status code
func (o *DocumentCategoriesCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this document categories create o k response has a 5xx status code
func (o *DocumentCategoriesCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this document categories create o k response a status code equal to that given
func (o *DocumentCategoriesCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the document categories create o k response
func (o *DocumentCategoriesCreateOK) Code() int {
	return 200
}

func (o *DocumentCategoriesCreateOK) Error() string {
	return fmt.Sprintf("[POST /ReferenceData/DocumentCategories][%d] documentCategoriesCreateOK  %+v", 200, o.Payload)
}

func (o *DocumentCategoriesCreateOK) String() string {
	return fmt.Sprintf("[POST /ReferenceData/DocumentCategories][%d] documentCategoriesCreateOK  %+v", 200, o.Payload)
}

func (o *DocumentCategoriesCreateOK) GetPayload() *models.DocumentCategory {
	return o.Payload
}

func (o *DocumentCategoriesCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DocumentCategory)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
