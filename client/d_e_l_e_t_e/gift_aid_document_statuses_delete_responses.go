// Code generated by go-swagger; DO NOT EDIT.

package d_e_l_e_t_e

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GiftAidDocumentStatusesDeleteReader is a Reader for the GiftAidDocumentStatusesDelete structure.
type GiftAidDocumentStatusesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GiftAidDocumentStatusesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewGiftAidDocumentStatusesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /ReferenceData/GiftAidDocumentStatuses/{id}] GiftAidDocumentStatuses_Delete", response, response.Code())
	}
}

// NewGiftAidDocumentStatusesDeleteNoContent creates a GiftAidDocumentStatusesDeleteNoContent with default headers values
func NewGiftAidDocumentStatusesDeleteNoContent() *GiftAidDocumentStatusesDeleteNoContent {
	return &GiftAidDocumentStatusesDeleteNoContent{}
}

/*
GiftAidDocumentStatusesDeleteNoContent describes a response with status code 204, with default header values.

No Content
*/
type GiftAidDocumentStatusesDeleteNoContent struct {
}

// IsSuccess returns true when this gift aid document statuses delete no content response has a 2xx status code
func (o *GiftAidDocumentStatusesDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this gift aid document statuses delete no content response has a 3xx status code
func (o *GiftAidDocumentStatusesDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this gift aid document statuses delete no content response has a 4xx status code
func (o *GiftAidDocumentStatusesDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this gift aid document statuses delete no content response has a 5xx status code
func (o *GiftAidDocumentStatusesDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this gift aid document statuses delete no content response a status code equal to that given
func (o *GiftAidDocumentStatusesDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the gift aid document statuses delete no content response
func (o *GiftAidDocumentStatusesDeleteNoContent) Code() int {
	return 204
}

func (o *GiftAidDocumentStatusesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ReferenceData/GiftAidDocumentStatuses/{id}][%d] giftAidDocumentStatusesDeleteNoContent ", 204)
}

func (o *GiftAidDocumentStatusesDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /ReferenceData/GiftAidDocumentStatuses/{id}][%d] giftAidDocumentStatusesDeleteNoContent ", 204)
}

func (o *GiftAidDocumentStatusesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
