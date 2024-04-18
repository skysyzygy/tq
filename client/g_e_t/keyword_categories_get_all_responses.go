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

// KeywordCategoriesGetAllReader is a Reader for the KeywordCategoriesGetAll structure.
type KeywordCategoriesGetAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KeywordCategoriesGetAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKeywordCategoriesGetAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /ReferenceData/KeywordCategories] KeywordCategories_GetAll", response, response.Code())
	}
}

// NewKeywordCategoriesGetAllOK creates a KeywordCategoriesGetAllOK with default headers values
func NewKeywordCategoriesGetAllOK() *KeywordCategoriesGetAllOK {
	return &KeywordCategoriesGetAllOK{}
}

/*
KeywordCategoriesGetAllOK describes a response with status code 200, with default header values.

OK
*/
type KeywordCategoriesGetAllOK struct {
	Payload []*models.KeywordCategory
}

// IsSuccess returns true when this keyword categories get all o k response has a 2xx status code
func (o *KeywordCategoriesGetAllOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this keyword categories get all o k response has a 3xx status code
func (o *KeywordCategoriesGetAllOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this keyword categories get all o k response has a 4xx status code
func (o *KeywordCategoriesGetAllOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this keyword categories get all o k response has a 5xx status code
func (o *KeywordCategoriesGetAllOK) IsServerError() bool {
	return false
}

// IsCode returns true when this keyword categories get all o k response a status code equal to that given
func (o *KeywordCategoriesGetAllOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the keyword categories get all o k response
func (o *KeywordCategoriesGetAllOK) Code() int {
	return 200
}

func (o *KeywordCategoriesGetAllOK) Error() string {
	return fmt.Sprintf("[GET /ReferenceData/KeywordCategories][%d] keywordCategoriesGetAllOK  %+v", 200, o.Payload)
}

func (o *KeywordCategoriesGetAllOK) String() string {
	return fmt.Sprintf("[GET /ReferenceData/KeywordCategories][%d] keywordCategoriesGetAllOK  %+v", 200, o.Payload)
}

func (o *KeywordCategoriesGetAllOK) GetPayload() []*models.KeywordCategory {
	return o.Payload
}

func (o *KeywordCategoriesGetAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}