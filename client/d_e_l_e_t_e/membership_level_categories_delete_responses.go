// Code generated by go-swagger; DO NOT EDIT.

package d_e_l_e_t_e

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// MembershipLevelCategoriesDeleteReader is a Reader for the MembershipLevelCategoriesDelete structure.
type MembershipLevelCategoriesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MembershipLevelCategoriesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewMembershipLevelCategoriesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /ReferenceData/MembershipLevelCategories/{id}] MembershipLevelCategories_Delete", response, response.Code())
	}
}

// NewMembershipLevelCategoriesDeleteNoContent creates a MembershipLevelCategoriesDeleteNoContent with default headers values
func NewMembershipLevelCategoriesDeleteNoContent() *MembershipLevelCategoriesDeleteNoContent {
	return &MembershipLevelCategoriesDeleteNoContent{}
}

/*
MembershipLevelCategoriesDeleteNoContent describes a response with status code 204, with default header values.

No Content
*/
type MembershipLevelCategoriesDeleteNoContent struct {
}

// IsSuccess returns true when this membership level categories delete no content response has a 2xx status code
func (o *MembershipLevelCategoriesDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this membership level categories delete no content response has a 3xx status code
func (o *MembershipLevelCategoriesDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this membership level categories delete no content response has a 4xx status code
func (o *MembershipLevelCategoriesDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this membership level categories delete no content response has a 5xx status code
func (o *MembershipLevelCategoriesDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this membership level categories delete no content response a status code equal to that given
func (o *MembershipLevelCategoriesDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the membership level categories delete no content response
func (o *MembershipLevelCategoriesDeleteNoContent) Code() int {
	return 204
}

func (o *MembershipLevelCategoriesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ReferenceData/MembershipLevelCategories/{id}][%d] membershipLevelCategoriesDeleteNoContent ", 204)
}

func (o *MembershipLevelCategoriesDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /ReferenceData/MembershipLevelCategories/{id}][%d] membershipLevelCategoriesDeleteNoContent ", 204)
}

func (o *MembershipLevelCategoriesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}