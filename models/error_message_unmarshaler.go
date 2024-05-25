package models

import (
	"encoding/json"

	"github.com/go-openapi/swag"
)

// UnmarshalJSON interface implementation
func (m *ErrorMessage) UnmarshalJSON(b []byte) error {
	type _ErrorMessage ErrorMessage
	res := make([]_ErrorMessage, 1)
	err := swag.ReadJSON(b, &res)
	if err != nil {
		err = json.Unmarshal(b, &res[0])
	}
	if err != nil {
		res[0].Code = "500"
		res[0].Description = string(b)
		res[0].Details = "Unknown Error Message"
	}
	*m = ErrorMessage(res[0])
	return nil
}

// UnmarshalText interface implementation
func (m *ErrorMessage) UnmarshalText(b []byte) error {
	m.Description = string(b)
	return nil
}
