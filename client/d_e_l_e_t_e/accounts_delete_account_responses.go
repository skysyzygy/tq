// Code generated by go-swagger; DO NOT EDIT.

package d_e_l_e_t_e

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

// AccountsDeleteAccountReader is a Reader for the AccountsDeleteAccount structure.
type AccountsDeleteAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AccountsDeleteAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewAccountsDeleteAccountNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAccountsDeleteAccountDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAccountsDeleteAccountNoContent creates a AccountsDeleteAccountNoContent with default headers values
func NewAccountsDeleteAccountNoContent() *AccountsDeleteAccountNoContent {
	return &AccountsDeleteAccountNoContent{}
}

/*
AccountsDeleteAccountNoContent describes a response with status code 204, with default header values.

No Content
*/
type AccountsDeleteAccountNoContent struct {
}

// IsSuccess returns true when this accounts delete account no content response has a 2xx status code
func (o *AccountsDeleteAccountNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this accounts delete account no content response has a 3xx status code
func (o *AccountsDeleteAccountNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this accounts delete account no content response has a 4xx status code
func (o *AccountsDeleteAccountNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this accounts delete account no content response has a 5xx status code
func (o *AccountsDeleteAccountNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this accounts delete account no content response a status code equal to that given
func (o *AccountsDeleteAccountNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the accounts delete account no content response
func (o *AccountsDeleteAccountNoContent) Code() int {
	return 204
}

func (o *AccountsDeleteAccountNoContent) Error() string {
	return fmt.Sprintf("[DELETE /CRM/Accounts/{accountId}][%d] accountsDeleteAccountNoContent", 204)
}

func (o *AccountsDeleteAccountNoContent) String() string {
	return fmt.Sprintf("[DELETE /CRM/Accounts/{accountId}][%d] accountsDeleteAccountNoContent", 204)
}

func (o *AccountsDeleteAccountNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAccountsDeleteAccountDefault creates a AccountsDeleteAccountDefault with default headers values
func NewAccountsDeleteAccountDefault(code int) *AccountsDeleteAccountDefault {
	return &AccountsDeleteAccountDefault{
		_statusCode: code,
	}
}

/*
AccountsDeleteAccountDefault describes a response with status code -1, with default header values.

Error
*/
type AccountsDeleteAccountDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this accounts delete account default response has a 2xx status code
func (o *AccountsDeleteAccountDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this accounts delete account default response has a 3xx status code
func (o *AccountsDeleteAccountDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this accounts delete account default response has a 4xx status code
func (o *AccountsDeleteAccountDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this accounts delete account default response has a 5xx status code
func (o *AccountsDeleteAccountDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this accounts delete account default response a status code equal to that given
func (o *AccountsDeleteAccountDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the accounts delete account default response
func (o *AccountsDeleteAccountDefault) Code() int {
	return o._statusCode
}

func (o *AccountsDeleteAccountDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /CRM/Accounts/{accountId}][%d] Accounts_DeleteAccount default %s", o._statusCode, payload)
}

func (o *AccountsDeleteAccountDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /CRM/Accounts/{accountId}][%d] Accounts_DeleteAccount default %s", o._statusCode, payload)
}

func (o *AccountsDeleteAccountDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *AccountsDeleteAccountDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
