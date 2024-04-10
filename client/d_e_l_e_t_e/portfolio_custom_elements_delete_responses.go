// Code generated by go-swagger; DO NOT EDIT.

package d_e_l_e_t_e

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PortfolioCustomElementsDeleteReader is a Reader for the PortfolioCustomElementsDelete structure.
type PortfolioCustomElementsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PortfolioCustomElementsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPortfolioCustomElementsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /ReferenceData/PortfolioCustomElements/{id}] PortfolioCustomElements_Delete", response, response.Code())
	}
}

// NewPortfolioCustomElementsDeleteNoContent creates a PortfolioCustomElementsDeleteNoContent with default headers values
func NewPortfolioCustomElementsDeleteNoContent() *PortfolioCustomElementsDeleteNoContent {
	return &PortfolioCustomElementsDeleteNoContent{}
}

/*
PortfolioCustomElementsDeleteNoContent describes a response with status code 204, with default header values.

No Content
*/
type PortfolioCustomElementsDeleteNoContent struct {
}

// IsSuccess returns true when this portfolio custom elements delete no content response has a 2xx status code
func (o *PortfolioCustomElementsDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this portfolio custom elements delete no content response has a 3xx status code
func (o *PortfolioCustomElementsDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this portfolio custom elements delete no content response has a 4xx status code
func (o *PortfolioCustomElementsDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this portfolio custom elements delete no content response has a 5xx status code
func (o *PortfolioCustomElementsDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this portfolio custom elements delete no content response a status code equal to that given
func (o *PortfolioCustomElementsDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the portfolio custom elements delete no content response
func (o *PortfolioCustomElementsDeleteNoContent) Code() int {
	return 204
}

func (o *PortfolioCustomElementsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ReferenceData/PortfolioCustomElements/{id}][%d] portfolioCustomElementsDeleteNoContent ", 204)
}

func (o *PortfolioCustomElementsDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /ReferenceData/PortfolioCustomElements/{id}][%d] portfolioCustomElementsDeleteNoContent ", 204)
}

func (o *PortfolioCustomElementsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
