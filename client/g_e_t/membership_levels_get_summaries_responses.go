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

// MembershipLevelsGetSummariesReader is a Reader for the MembershipLevelsGetSummaries structure.
type MembershipLevelsGetSummariesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MembershipLevelsGetSummariesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMembershipLevelsGetSummariesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /Finance/MembershipLevels/Summary] MembershipLevels_GetSummaries", response, response.Code())
	}
}

// NewMembershipLevelsGetSummariesOK creates a MembershipLevelsGetSummariesOK with default headers values
func NewMembershipLevelsGetSummariesOK() *MembershipLevelsGetSummariesOK {
	return &MembershipLevelsGetSummariesOK{}
}

/*
MembershipLevelsGetSummariesOK describes a response with status code 200, with default header values.

OK
*/
type MembershipLevelsGetSummariesOK struct {
	Payload []*models.MembershipLevelSummary
}

// IsSuccess returns true when this membership levels get summaries o k response has a 2xx status code
func (o *MembershipLevelsGetSummariesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this membership levels get summaries o k response has a 3xx status code
func (o *MembershipLevelsGetSummariesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this membership levels get summaries o k response has a 4xx status code
func (o *MembershipLevelsGetSummariesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this membership levels get summaries o k response has a 5xx status code
func (o *MembershipLevelsGetSummariesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this membership levels get summaries o k response a status code equal to that given
func (o *MembershipLevelsGetSummariesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the membership levels get summaries o k response
func (o *MembershipLevelsGetSummariesOK) Code() int {
	return 200
}

func (o *MembershipLevelsGetSummariesOK) Error() string {
	return fmt.Sprintf("[GET /Finance/MembershipLevels/Summary][%d] membershipLevelsGetSummariesOK  %+v", 200, o.Payload)
}

func (o *MembershipLevelsGetSummariesOK) String() string {
	return fmt.Sprintf("[GET /Finance/MembershipLevels/Summary][%d] membershipLevelsGetSummariesOK  %+v", 200, o.Payload)
}

func (o *MembershipLevelsGetSummariesOK) GetPayload() []*models.MembershipLevelSummary {
	return o.Payload
}

func (o *MembershipLevelsGetSummariesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
