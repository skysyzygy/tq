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

// AppealCategoriesGetAllReader is a Reader for the AppealCategoriesGetAll structure.
type AppealCategoriesGetAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AppealCategoriesGetAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAppealCategoriesGetAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /ReferenceData/AppealCategories] AppealCategories_GetAll", response, response.Code())
	}
}

// NewAppealCategoriesGetAllOK creates a AppealCategoriesGetAllOK with default headers values
func NewAppealCategoriesGetAllOK() *AppealCategoriesGetAllOK {
	return &AppealCategoriesGetAllOK{}
}

/*
AppealCategoriesGetAllOK describes a response with status code 200, with default header values.

OK
*/
type AppealCategoriesGetAllOK struct {
	Payload []*models.AppealCategory
}

// IsSuccess returns true when this appeal categories get all o k response has a 2xx status code
func (o *AppealCategoriesGetAllOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this appeal categories get all o k response has a 3xx status code
func (o *AppealCategoriesGetAllOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this appeal categories get all o k response has a 4xx status code
func (o *AppealCategoriesGetAllOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this appeal categories get all o k response has a 5xx status code
func (o *AppealCategoriesGetAllOK) IsServerError() bool {
	return false
}

// IsCode returns true when this appeal categories get all o k response a status code equal to that given
func (o *AppealCategoriesGetAllOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the appeal categories get all o k response
func (o *AppealCategoriesGetAllOK) Code() int {
	return 200
}

func (o *AppealCategoriesGetAllOK) Error() string {
	return fmt.Sprintf("[GET /ReferenceData/AppealCategories][%d] appealCategoriesGetAllOK  %+v", 200, o.Payload)
}

func (o *AppealCategoriesGetAllOK) String() string {
	return fmt.Sprintf("[GET /ReferenceData/AppealCategories][%d] appealCategoriesGetAllOK  %+v", 200, o.Payload)
}

func (o *AppealCategoriesGetAllOK) GetPayload() []*models.AppealCategory {
	return o.Payload
}

func (o *AppealCategoriesGetAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
