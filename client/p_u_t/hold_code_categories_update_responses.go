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

// HoldCodeCategoriesUpdateReader is a Reader for the HoldCodeCategoriesUpdate structure.
type HoldCodeCategoriesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HoldCodeCategoriesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHoldCodeCategoriesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /ReferenceData/HoldCodeCategories/{id}] HoldCodeCategories_Update", response, response.Code())
	}
}

// NewHoldCodeCategoriesUpdateOK creates a HoldCodeCategoriesUpdateOK with default headers values
func NewHoldCodeCategoriesUpdateOK() *HoldCodeCategoriesUpdateOK {
	return &HoldCodeCategoriesUpdateOK{}
}

/*
HoldCodeCategoriesUpdateOK describes a response with status code 200, with default header values.

OK
*/
type HoldCodeCategoriesUpdateOK struct {
	Payload *models.HoldCodeCategory
}

// IsSuccess returns true when this hold code categories update o k response has a 2xx status code
func (o *HoldCodeCategoriesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this hold code categories update o k response has a 3xx status code
func (o *HoldCodeCategoriesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hold code categories update o k response has a 4xx status code
func (o *HoldCodeCategoriesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this hold code categories update o k response has a 5xx status code
func (o *HoldCodeCategoriesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this hold code categories update o k response a status code equal to that given
func (o *HoldCodeCategoriesUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the hold code categories update o k response
func (o *HoldCodeCategoriesUpdateOK) Code() int {
	return 200
}

func (o *HoldCodeCategoriesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /ReferenceData/HoldCodeCategories/{id}][%d] holdCodeCategoriesUpdateOK  %+v", 200, o.Payload)
}

func (o *HoldCodeCategoriesUpdateOK) String() string {
	return fmt.Sprintf("[PUT /ReferenceData/HoldCodeCategories/{id}][%d] holdCodeCategoriesUpdateOK  %+v", 200, o.Payload)
}

func (o *HoldCodeCategoriesUpdateOK) GetPayload() *models.HoldCodeCategory {
	return o.Payload
}

func (o *HoldCodeCategoriesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HoldCodeCategory)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
