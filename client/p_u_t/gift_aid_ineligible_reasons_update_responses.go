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

// GiftAidIneligibleReasonsUpdateReader is a Reader for the GiftAidIneligibleReasonsUpdate structure.
type GiftAidIneligibleReasonsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GiftAidIneligibleReasonsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGiftAidIneligibleReasonsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /ReferenceData/GiftAidIneligibleReasons/{id}] GiftAidIneligibleReasons_Update", response, response.Code())
	}
}

// NewGiftAidIneligibleReasonsUpdateOK creates a GiftAidIneligibleReasonsUpdateOK with default headers values
func NewGiftAidIneligibleReasonsUpdateOK() *GiftAidIneligibleReasonsUpdateOK {
	return &GiftAidIneligibleReasonsUpdateOK{}
}

/*
GiftAidIneligibleReasonsUpdateOK describes a response with status code 200, with default header values.

OK
*/
type GiftAidIneligibleReasonsUpdateOK struct {
	Payload *models.GiftAidIneligibleReason
}

// IsSuccess returns true when this gift aid ineligible reasons update o k response has a 2xx status code
func (o *GiftAidIneligibleReasonsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this gift aid ineligible reasons update o k response has a 3xx status code
func (o *GiftAidIneligibleReasonsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this gift aid ineligible reasons update o k response has a 4xx status code
func (o *GiftAidIneligibleReasonsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this gift aid ineligible reasons update o k response has a 5xx status code
func (o *GiftAidIneligibleReasonsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this gift aid ineligible reasons update o k response a status code equal to that given
func (o *GiftAidIneligibleReasonsUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the gift aid ineligible reasons update o k response
func (o *GiftAidIneligibleReasonsUpdateOK) Code() int {
	return 200
}

func (o *GiftAidIneligibleReasonsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /ReferenceData/GiftAidIneligibleReasons/{id}][%d] giftAidIneligibleReasonsUpdateOK  %+v", 200, o.Payload)
}

func (o *GiftAidIneligibleReasonsUpdateOK) String() string {
	return fmt.Sprintf("[PUT /ReferenceData/GiftAidIneligibleReasons/{id}][%d] giftAidIneligibleReasonsUpdateOK  %+v", 200, o.Payload)
}

func (o *GiftAidIneligibleReasonsUpdateOK) GetPayload() *models.GiftAidIneligibleReason {
	return o.Payload
}

func (o *GiftAidIneligibleReasonsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GiftAidIneligibleReason)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}