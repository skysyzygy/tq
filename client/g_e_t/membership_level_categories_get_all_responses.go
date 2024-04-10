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

// MembershipLevelCategoriesGetAllReader is a Reader for the MembershipLevelCategoriesGetAll structure.
type MembershipLevelCategoriesGetAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MembershipLevelCategoriesGetAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMembershipLevelCategoriesGetAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /ReferenceData/MembershipLevelCategories] MembershipLevelCategories_GetAll", response, response.Code())
	}
}

// NewMembershipLevelCategoriesGetAllOK creates a MembershipLevelCategoriesGetAllOK with default headers values
func NewMembershipLevelCategoriesGetAllOK() *MembershipLevelCategoriesGetAllOK {
	return &MembershipLevelCategoriesGetAllOK{}
}

/*
MembershipLevelCategoriesGetAllOK describes a response with status code 200, with default header values.

OK
*/
type MembershipLevelCategoriesGetAllOK struct {
	Payload []*models.MembershipLevelCategory
}

// IsSuccess returns true when this membership level categories get all o k response has a 2xx status code
func (o *MembershipLevelCategoriesGetAllOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this membership level categories get all o k response has a 3xx status code
func (o *MembershipLevelCategoriesGetAllOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this membership level categories get all o k response has a 4xx status code
func (o *MembershipLevelCategoriesGetAllOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this membership level categories get all o k response has a 5xx status code
func (o *MembershipLevelCategoriesGetAllOK) IsServerError() bool {
	return false
}

// IsCode returns true when this membership level categories get all o k response a status code equal to that given
func (o *MembershipLevelCategoriesGetAllOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the membership level categories get all o k response
func (o *MembershipLevelCategoriesGetAllOK) Code() int {
	return 200
}

func (o *MembershipLevelCategoriesGetAllOK) Error() string {
	return fmt.Sprintf("[GET /ReferenceData/MembershipLevelCategories][%d] membershipLevelCategoriesGetAllOK  %+v", 200, o.Payload)
}

func (o *MembershipLevelCategoriesGetAllOK) String() string {
	return fmt.Sprintf("[GET /ReferenceData/MembershipLevelCategories][%d] membershipLevelCategoriesGetAllOK  %+v", 200, o.Payload)
}

func (o *MembershipLevelCategoriesGetAllOK) GetPayload() []*models.MembershipLevelCategory {
	return o.Payload
}

func (o *MembershipLevelCategoriesGetAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
