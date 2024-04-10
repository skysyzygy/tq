// Code generated by go-swagger; DO NOT EDIT.

package d_e_l_e_t_e

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CampaignDesignationsDeleteReader is a Reader for the CampaignDesignationsDelete structure.
type CampaignDesignationsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CampaignDesignationsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewCampaignDesignationsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /Finance/CampaignDesignations/{campaignDesignationId}] CampaignDesignations_Delete", response, response.Code())
	}
}

// NewCampaignDesignationsDeleteNoContent creates a CampaignDesignationsDeleteNoContent with default headers values
func NewCampaignDesignationsDeleteNoContent() *CampaignDesignationsDeleteNoContent {
	return &CampaignDesignationsDeleteNoContent{}
}

/*
CampaignDesignationsDeleteNoContent describes a response with status code 204, with default header values.

No Content
*/
type CampaignDesignationsDeleteNoContent struct {
}

// IsSuccess returns true when this campaign designations delete no content response has a 2xx status code
func (o *CampaignDesignationsDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this campaign designations delete no content response has a 3xx status code
func (o *CampaignDesignationsDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this campaign designations delete no content response has a 4xx status code
func (o *CampaignDesignationsDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this campaign designations delete no content response has a 5xx status code
func (o *CampaignDesignationsDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this campaign designations delete no content response a status code equal to that given
func (o *CampaignDesignationsDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the campaign designations delete no content response
func (o *CampaignDesignationsDeleteNoContent) Code() int {
	return 204
}

func (o *CampaignDesignationsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /Finance/CampaignDesignations/{campaignDesignationId}][%d] campaignDesignationsDeleteNoContent ", 204)
}

func (o *CampaignDesignationsDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /Finance/CampaignDesignations/{campaignDesignationId}][%d] campaignDesignationsDeleteNoContent ", 204)
}

func (o *CampaignDesignationsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
