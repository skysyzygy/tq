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
		if err != nil {
			return err
		}
	}
	*m = ErrorMessage(res[0])
	return nil
}
