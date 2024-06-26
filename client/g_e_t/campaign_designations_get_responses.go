// Code generated by go-swagger; DO NOT EDIT.

package g_e_t

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skysyzygy/tq/models"
)

// CampaignDesignationsGetReader is a Reader for the CampaignDesignationsGet structure.
type CampaignDesignationsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CampaignDesignationsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCampaignDesignationsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCampaignDesignationsGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCampaignDesignationsGetOK creates a CampaignDesignationsGetOK with default headers values
func NewCampaignDesignationsGetOK() *CampaignDesignationsGetOK {
	return &CampaignDesignationsGetOK{}
}

/*
CampaignDesignationsGetOK describes a response with status code 200, with default header values.

OK
*/
type CampaignDesignationsGetOK struct {
	Payload *models.CampaignDesignation
}

// IsSuccess returns true when this campaign designations get o k response has a 2xx status code
func (o *CampaignDesignationsGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this campaign designations get o k response has a 3xx status code
func (o *CampaignDesignationsGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this campaign designations get o k response has a 4xx status code
func (o *CampaignDesignationsGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this campaign designations get o k response has a 5xx status code
func (o *CampaignDesignationsGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this campaign designations get o k response a status code equal to that given
func (o *CampaignDesignationsGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the campaign designations get o k response
func (o *CampaignDesignationsGetOK) Code() int {
	return 200
}

func (o *CampaignDesignationsGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /Finance/CampaignDesignations/{campaignDesignationId}][%d] campaignDesignationsGetOK %s", 200, payload)
}

func (o *CampaignDesignationsGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /Finance/CampaignDesignations/{campaignDesignationId}][%d] campaignDesignationsGetOK %s", 200, payload)
}

func (o *CampaignDesignationsGetOK) GetPayload() *models.CampaignDesignation {
	return o.Payload
}

func (o *CampaignDesignationsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CampaignDesignation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCampaignDesignationsGetDefault creates a CampaignDesignationsGetDefault with default headers values
func NewCampaignDesignationsGetDefault(code int) *CampaignDesignationsGetDefault {
	return &CampaignDesignationsGetDefault{
		_statusCode: code,
	}
}

/*
CampaignDesignationsGetDefault describes a response with status code -1, with default header values.

Error
*/
type CampaignDesignationsGetDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this campaign designations get default response has a 2xx status code
func (o *CampaignDesignationsGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this campaign designations get default response has a 3xx status code
func (o *CampaignDesignationsGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this campaign designations get default response has a 4xx status code
func (o *CampaignDesignationsGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this campaign designations get default response has a 5xx status code
func (o *CampaignDesignationsGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this campaign designations get default response a status code equal to that given
func (o *CampaignDesignationsGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the campaign designations get default response
func (o *CampaignDesignationsGetDefault) Code() int {
	return o._statusCode
}

func (o *CampaignDesignationsGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /Finance/CampaignDesignations/{campaignDesignationId}][%d] CampaignDesignations_Get default %s", o._statusCode, payload)
}

func (o *CampaignDesignationsGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /Finance/CampaignDesignations/{campaignDesignationId}][%d] CampaignDesignations_Get default %s", o._statusCode, payload)
}

func (o *CampaignDesignationsGetDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *CampaignDesignationsGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
