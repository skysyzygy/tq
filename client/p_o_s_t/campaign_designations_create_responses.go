// Code generated by go-swagger; DO NOT EDIT.

package p_o_s_t

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skysyzygy/tq/models"
)

// CampaignDesignationsCreateReader is a Reader for the CampaignDesignationsCreate structure.
type CampaignDesignationsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CampaignDesignationsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCampaignDesignationsCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /Finance/CampaignDesignations] CampaignDesignations_Create", response, response.Code())
	}
}

// NewCampaignDesignationsCreateOK creates a CampaignDesignationsCreateOK with default headers values
func NewCampaignDesignationsCreateOK() *CampaignDesignationsCreateOK {
	return &CampaignDesignationsCreateOK{}
}

/*
CampaignDesignationsCreateOK describes a response with status code 200, with default header values.

OK
*/
type CampaignDesignationsCreateOK struct {
	Payload *models.CampaignDesignation
}

// IsSuccess returns true when this campaign designations create o k response has a 2xx status code
func (o *CampaignDesignationsCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this campaign designations create o k response has a 3xx status code
func (o *CampaignDesignationsCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this campaign designations create o k response has a 4xx status code
func (o *CampaignDesignationsCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this campaign designations create o k response has a 5xx status code
func (o *CampaignDesignationsCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this campaign designations create o k response a status code equal to that given
func (o *CampaignDesignationsCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the campaign designations create o k response
func (o *CampaignDesignationsCreateOK) Code() int {
	return 200
}

func (o *CampaignDesignationsCreateOK) Error() string {
	return fmt.Sprintf("[POST /Finance/CampaignDesignations][%d] campaignDesignationsCreateOK  %+v", 200, o.Payload)
}

func (o *CampaignDesignationsCreateOK) String() string {
	return fmt.Sprintf("[POST /Finance/CampaignDesignations][%d] campaignDesignationsCreateOK  %+v", 200, o.Payload)
}

func (o *CampaignDesignationsCreateOK) GetPayload() *models.CampaignDesignation {
	return o.Payload
}

func (o *CampaignDesignationsCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CampaignDesignation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
