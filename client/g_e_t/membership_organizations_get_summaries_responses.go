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

// MembershipOrganizationsGetSummariesReader is a Reader for the MembershipOrganizationsGetSummaries structure.
type MembershipOrganizationsGetSummariesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MembershipOrganizationsGetSummariesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMembershipOrganizationsGetSummariesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /Finance/MembershipOrganizations/Summary] MembershipOrganizations_GetSummaries", response, response.Code())
	}
}

// NewMembershipOrganizationsGetSummariesOK creates a MembershipOrganizationsGetSummariesOK with default headers values
func NewMembershipOrganizationsGetSummariesOK() *MembershipOrganizationsGetSummariesOK {
	return &MembershipOrganizationsGetSummariesOK{}
}

/*
MembershipOrganizationsGetSummariesOK describes a response with status code 200, with default header values.

OK
*/
type MembershipOrganizationsGetSummariesOK struct {
	Payload []*models.MembershipOrganizationSummary
}

// IsSuccess returns true when this membership organizations get summaries o k response has a 2xx status code
func (o *MembershipOrganizationsGetSummariesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this membership organizations get summaries o k response has a 3xx status code
func (o *MembershipOrganizationsGetSummariesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this membership organizations get summaries o k response has a 4xx status code
func (o *MembershipOrganizationsGetSummariesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this membership organizations get summaries o k response has a 5xx status code
func (o *MembershipOrganizationsGetSummariesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this membership organizations get summaries o k response a status code equal to that given
func (o *MembershipOrganizationsGetSummariesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the membership organizations get summaries o k response
func (o *MembershipOrganizationsGetSummariesOK) Code() int {
	return 200
}

func (o *MembershipOrganizationsGetSummariesOK) Error() string {
	return fmt.Sprintf("[GET /Finance/MembershipOrganizations/Summary][%d] membershipOrganizationsGetSummariesOK  %+v", 200, o.Payload)
}

func (o *MembershipOrganizationsGetSummariesOK) String() string {
	return fmt.Sprintf("[GET /Finance/MembershipOrganizations/Summary][%d] membershipOrganizationsGetSummariesOK  %+v", 200, o.Payload)
}

func (o *MembershipOrganizationsGetSummariesOK) GetPayload() []*models.MembershipOrganizationSummary {
	return o.Payload
}

func (o *MembershipOrganizationsGetSummariesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
