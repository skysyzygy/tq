// Code generated by go-swagger; DO NOT EDIT.

package d_e_l_e_t_e

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PriceTemplatesDeleteReader is a Reader for the PriceTemplatesDelete structure.
type PriceTemplatesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PriceTemplatesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPriceTemplatesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /TXN/PriceTemplates/{priceTemplateId}] PriceTemplates_Delete", response, response.Code())
	}
}

// NewPriceTemplatesDeleteNoContent creates a PriceTemplatesDeleteNoContent with default headers values
func NewPriceTemplatesDeleteNoContent() *PriceTemplatesDeleteNoContent {
	return &PriceTemplatesDeleteNoContent{}
}

/*
PriceTemplatesDeleteNoContent describes a response with status code 204, with default header values.

No Content
*/
type PriceTemplatesDeleteNoContent struct {
}

// IsSuccess returns true when this price templates delete no content response has a 2xx status code
func (o *PriceTemplatesDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this price templates delete no content response has a 3xx status code
func (o *PriceTemplatesDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this price templates delete no content response has a 4xx status code
func (o *PriceTemplatesDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this price templates delete no content response has a 5xx status code
func (o *PriceTemplatesDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this price templates delete no content response a status code equal to that given
func (o *PriceTemplatesDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the price templates delete no content response
func (o *PriceTemplatesDeleteNoContent) Code() int {
	return 204
}

func (o *PriceTemplatesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /TXN/PriceTemplates/{priceTemplateId}][%d] priceTemplatesDeleteNoContent ", 204)
}

func (o *PriceTemplatesDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /TXN/PriceTemplates/{priceTemplateId}][%d] priceTemplatesDeleteNoContent ", 204)
}

func (o *PriceTemplatesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
