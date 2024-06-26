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

// DiagnosticsGetAPIPluginConfigurationReader is a Reader for the DiagnosticsGetAPIPluginConfiguration structure.
type DiagnosticsGetAPIPluginConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DiagnosticsGetAPIPluginConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDiagnosticsGetAPIPluginConfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDiagnosticsGetAPIPluginConfigurationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDiagnosticsGetAPIPluginConfigurationOK creates a DiagnosticsGetAPIPluginConfigurationOK with default headers values
func NewDiagnosticsGetAPIPluginConfigurationOK() *DiagnosticsGetAPIPluginConfigurationOK {
	return &DiagnosticsGetAPIPluginConfigurationOK{}
}

/*
DiagnosticsGetAPIPluginConfigurationOK describes a response with status code 200, with default header values.

OK
*/
type DiagnosticsGetAPIPluginConfigurationOK struct {
	Payload *models.APIPluginConfiguration
}

// IsSuccess returns true when this diagnostics get Api plugin configuration o k response has a 2xx status code
func (o *DiagnosticsGetAPIPluginConfigurationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this diagnostics get Api plugin configuration o k response has a 3xx status code
func (o *DiagnosticsGetAPIPluginConfigurationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this diagnostics get Api plugin configuration o k response has a 4xx status code
func (o *DiagnosticsGetAPIPluginConfigurationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this diagnostics get Api plugin configuration o k response has a 5xx status code
func (o *DiagnosticsGetAPIPluginConfigurationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this diagnostics get Api plugin configuration o k response a status code equal to that given
func (o *DiagnosticsGetAPIPluginConfigurationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the diagnostics get Api plugin configuration o k response
func (o *DiagnosticsGetAPIPluginConfigurationOK) Code() int {
	return 200
}

func (o *DiagnosticsGetAPIPluginConfigurationOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /Diagnostics/ApiPluginConfiguration][%d] diagnosticsGetApiPluginConfigurationOK %s", 200, payload)
}

func (o *DiagnosticsGetAPIPluginConfigurationOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /Diagnostics/ApiPluginConfiguration][%d] diagnosticsGetApiPluginConfigurationOK %s", 200, payload)
}

func (o *DiagnosticsGetAPIPluginConfigurationOK) GetPayload() *models.APIPluginConfiguration {
	return o.Payload
}

func (o *DiagnosticsGetAPIPluginConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIPluginConfiguration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDiagnosticsGetAPIPluginConfigurationDefault creates a DiagnosticsGetAPIPluginConfigurationDefault with default headers values
func NewDiagnosticsGetAPIPluginConfigurationDefault(code int) *DiagnosticsGetAPIPluginConfigurationDefault {
	return &DiagnosticsGetAPIPluginConfigurationDefault{
		_statusCode: code,
	}
}

/*
DiagnosticsGetAPIPluginConfigurationDefault describes a response with status code -1, with default header values.

Error
*/
type DiagnosticsGetAPIPluginConfigurationDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this diagnostics get Api plugin configuration default response has a 2xx status code
func (o *DiagnosticsGetAPIPluginConfigurationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this diagnostics get Api plugin configuration default response has a 3xx status code
func (o *DiagnosticsGetAPIPluginConfigurationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this diagnostics get Api plugin configuration default response has a 4xx status code
func (o *DiagnosticsGetAPIPluginConfigurationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this diagnostics get Api plugin configuration default response has a 5xx status code
func (o *DiagnosticsGetAPIPluginConfigurationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this diagnostics get Api plugin configuration default response a status code equal to that given
func (o *DiagnosticsGetAPIPluginConfigurationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the diagnostics get Api plugin configuration default response
func (o *DiagnosticsGetAPIPluginConfigurationDefault) Code() int {
	return o._statusCode
}

func (o *DiagnosticsGetAPIPluginConfigurationDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /Diagnostics/ApiPluginConfiguration][%d] Diagnostics_GetApiPluginConfiguration default %s", o._statusCode, payload)
}

func (o *DiagnosticsGetAPIPluginConfigurationDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /Diagnostics/ApiPluginConfiguration][%d] Diagnostics_GetApiPluginConfiguration default %s", o._statusCode, payload)
}

func (o *DiagnosticsGetAPIPluginConfigurationDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *DiagnosticsGetAPIPluginConfigurationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
