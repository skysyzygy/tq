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

// GiftAidDeclarationsGetAllReader is a Reader for the GiftAidDeclarationsGetAll structure.
type GiftAidDeclarationsGetAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GiftAidDeclarationsGetAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGiftAidDeclarationsGetAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /CRM/GiftAidDeclarations] GiftAidDeclarations_GetAll", response, response.Code())
	}
}

// NewGiftAidDeclarationsGetAllOK creates a GiftAidDeclarationsGetAllOK with default headers values
func NewGiftAidDeclarationsGetAllOK() *GiftAidDeclarationsGetAllOK {
	return &GiftAidDeclarationsGetAllOK{}
}

/*
GiftAidDeclarationsGetAllOK describes a response with status code 200, with default header values.

OK
*/
type GiftAidDeclarationsGetAllOK struct {
	Payload []*models.GiftAidDeclaration
}

// IsSuccess returns true when this gift aid declarations get all o k response has a 2xx status code
func (o *GiftAidDeclarationsGetAllOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this gift aid declarations get all o k response has a 3xx status code
func (o *GiftAidDeclarationsGetAllOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this gift aid declarations get all o k response has a 4xx status code
func (o *GiftAidDeclarationsGetAllOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this gift aid declarations get all o k response has a 5xx status code
func (o *GiftAidDeclarationsGetAllOK) IsServerError() bool {
	return false
}

// IsCode returns true when this gift aid declarations get all o k response a status code equal to that given
func (o *GiftAidDeclarationsGetAllOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the gift aid declarations get all o k response
func (o *GiftAidDeclarationsGetAllOK) Code() int {
	return 200
}

func (o *GiftAidDeclarationsGetAllOK) Error() string {
	return fmt.Sprintf("[GET /CRM/GiftAidDeclarations][%d] giftAidDeclarationsGetAllOK  %+v", 200, o.Payload)
}

func (o *GiftAidDeclarationsGetAllOK) String() string {
	return fmt.Sprintf("[GET /CRM/GiftAidDeclarations][%d] giftAidDeclarationsGetAllOK  %+v", 200, o.Payload)
}

func (o *GiftAidDeclarationsGetAllOK) GetPayload() []*models.GiftAidDeclaration {
	return o.Payload
}

func (o *GiftAidDeclarationsGetAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
