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

// ColorsGetAllReader is a Reader for the ColorsGetAll structure.
type ColorsGetAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColorsGetAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColorsGetAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /ReferenceData/Colors] Colors_GetAll", response, response.Code())
	}
}

// NewColorsGetAllOK creates a ColorsGetAllOK with default headers values
func NewColorsGetAllOK() *ColorsGetAllOK {
	return &ColorsGetAllOK{}
}

/*
ColorsGetAllOK describes a response with status code 200, with default header values.

OK
*/
type ColorsGetAllOK struct {
	Payload []*models.Color
}

// IsSuccess returns true when this colors get all o k response has a 2xx status code
func (o *ColorsGetAllOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this colors get all o k response has a 3xx status code
func (o *ColorsGetAllOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this colors get all o k response has a 4xx status code
func (o *ColorsGetAllOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this colors get all o k response has a 5xx status code
func (o *ColorsGetAllOK) IsServerError() bool {
	return false
}

// IsCode returns true when this colors get all o k response a status code equal to that given
func (o *ColorsGetAllOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the colors get all o k response
func (o *ColorsGetAllOK) Code() int {
	return 200
}

func (o *ColorsGetAllOK) Error() string {
	return fmt.Sprintf("[GET /ReferenceData/Colors][%d] colorsGetAllOK  %+v", 200, o.Payload)
}

func (o *ColorsGetAllOK) String() string {
	return fmt.Sprintf("[GET /ReferenceData/Colors][%d] colorsGetAllOK  %+v", 200, o.Payload)
}

func (o *ColorsGetAllOK) GetPayload() []*models.Color {
	return o.Payload
}

func (o *ColorsGetAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}