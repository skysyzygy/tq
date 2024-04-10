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

// HoldCodesGetAllReader is a Reader for the HoldCodesGetAll structure.
type HoldCodesGetAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HoldCodesGetAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHoldCodesGetAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /TXN/HoldCodes] HoldCodes_GetAll", response, response.Code())
	}
}

// NewHoldCodesGetAllOK creates a HoldCodesGetAllOK with default headers values
func NewHoldCodesGetAllOK() *HoldCodesGetAllOK {
	return &HoldCodesGetAllOK{}
}

/*
HoldCodesGetAllOK describes a response with status code 200, with default header values.

OK
*/
type HoldCodesGetAllOK struct {
	Payload []*models.HoldCode
}

// IsSuccess returns true when this hold codes get all o k response has a 2xx status code
func (o *HoldCodesGetAllOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this hold codes get all o k response has a 3xx status code
func (o *HoldCodesGetAllOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hold codes get all o k response has a 4xx status code
func (o *HoldCodesGetAllOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this hold codes get all o k response has a 5xx status code
func (o *HoldCodesGetAllOK) IsServerError() bool {
	return false
}

// IsCode returns true when this hold codes get all o k response a status code equal to that given
func (o *HoldCodesGetAllOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the hold codes get all o k response
func (o *HoldCodesGetAllOK) Code() int {
	return 200
}

func (o *HoldCodesGetAllOK) Error() string {
	return fmt.Sprintf("[GET /TXN/HoldCodes][%d] holdCodesGetAllOK  %+v", 200, o.Payload)
}

func (o *HoldCodesGetAllOK) String() string {
	return fmt.Sprintf("[GET /TXN/HoldCodes][%d] holdCodesGetAllOK  %+v", 200, o.Payload)
}

func (o *HoldCodesGetAllOK) GetPayload() []*models.HoldCode {
	return o.Payload
}

func (o *HoldCodesGetAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
