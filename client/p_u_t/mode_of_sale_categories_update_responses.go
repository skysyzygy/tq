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

// ModeOfSaleCategoriesUpdateReader is a Reader for the ModeOfSaleCategoriesUpdate structure.
type ModeOfSaleCategoriesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ModeOfSaleCategoriesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewModeOfSaleCategoriesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /TXN/ModeOfSaleCategories/{modeOfSaleCategoryId}] ModeOfSaleCategories_Update", response, response.Code())
	}
}

// NewModeOfSaleCategoriesUpdateOK creates a ModeOfSaleCategoriesUpdateOK with default headers values
func NewModeOfSaleCategoriesUpdateOK() *ModeOfSaleCategoriesUpdateOK {
	return &ModeOfSaleCategoriesUpdateOK{}
}

/*
ModeOfSaleCategoriesUpdateOK describes a response with status code 200, with default header values.

OK
*/
type ModeOfSaleCategoriesUpdateOK struct {
	Payload *models.ModeOfSaleCategory
}

// IsSuccess returns true when this mode of sale categories update o k response has a 2xx status code
func (o *ModeOfSaleCategoriesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this mode of sale categories update o k response has a 3xx status code
func (o *ModeOfSaleCategoriesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this mode of sale categories update o k response has a 4xx status code
func (o *ModeOfSaleCategoriesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this mode of sale categories update o k response has a 5xx status code
func (o *ModeOfSaleCategoriesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this mode of sale categories update o k response a status code equal to that given
func (o *ModeOfSaleCategoriesUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the mode of sale categories update o k response
func (o *ModeOfSaleCategoriesUpdateOK) Code() int {
	return 200
}

func (o *ModeOfSaleCategoriesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /TXN/ModeOfSaleCategories/{modeOfSaleCategoryId}][%d] modeOfSaleCategoriesUpdateOK  %+v", 200, o.Payload)
}

func (o *ModeOfSaleCategoriesUpdateOK) String() string {
	return fmt.Sprintf("[PUT /TXN/ModeOfSaleCategories/{modeOfSaleCategoryId}][%d] modeOfSaleCategoriesUpdateOK  %+v", 200, o.Payload)
}

func (o *ModeOfSaleCategoriesUpdateOK) GetPayload() *models.ModeOfSaleCategory {
	return o.Payload
}

func (o *ModeOfSaleCategoriesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModeOfSaleCategory)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
