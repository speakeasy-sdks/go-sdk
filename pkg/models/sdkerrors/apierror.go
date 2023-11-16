// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"fmt"
)

// Type - api_error
type Type string

const (
	TypeAPIError Type = "api_error"
)

func (e Type) ToPointer() *Type {
	return &e
}

func (e *Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "api_error":
		*e = Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Type: %v", v)
	}
}

type APIError struct {
	Code    *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	// api_error
	Type *Type `json:"type,omitempty"`
}

var _ error = &APIError{}

func (e *APIError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
