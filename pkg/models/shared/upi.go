// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// UpiChannel - Specify the channel through which the payment must be processed. Can be one of ["link", "collect", "qrcode"]
type UpiChannel string

const (
	UpiChannelLink    UpiChannel = "link"
	UpiChannelCollect UpiChannel = "collect"
	UpiChannelQrcode  UpiChannel = "qrcode"
)

func (e UpiChannel) ToPointer() *UpiChannel {
	return &e
}

func (e *UpiChannel) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "link":
		fallthrough
	case "collect":
		fallthrough
	case "qrcode":
		*e = UpiChannel(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpiChannel: %v", v)
	}
}

type Upi struct {
	Authorization *UPIAuthorizeDetails `json:"authorization,omitempty"`
	// For one time mandate on UPI. Set this as authorize_only = true. Please note that you can only use the "collect" channel if you are sending a one time mandate request
	AuthorizeOnly *bool `json:"authorize_only,omitempty"`
	// Specify the channel through which the payment must be processed. Can be one of ["link", "collect", "qrcode"]
	Channel UpiChannel `json:"channel"`
	// The UPI request will be valid for this expiry minutes. This parameter is only applicable for a UPI collect payment. The default value is 5 minutes. You should keep the minimum as 5 minutes, and maximum as 15 minutes
	UpiExpiryMinutes *float64 `json:"upi_expiry_minutes,omitempty"`
	// Customer UPI VPA to process payment.
	UpiID *string `json:"upi_id,omitempty"`
}

func (o *Upi) GetAuthorization() *UPIAuthorizeDetails {
	if o == nil {
		return nil
	}
	return o.Authorization
}

func (o *Upi) GetAuthorizeOnly() *bool {
	if o == nil {
		return nil
	}
	return o.AuthorizeOnly
}

func (o *Upi) GetChannel() UpiChannel {
	if o == nil {
		return UpiChannel("")
	}
	return o.Channel
}

func (o *Upi) GetUpiExpiryMinutes() *float64 {
	if o == nil {
		return nil
	}
	return o.UpiExpiryMinutes
}

func (o *Upi) GetUpiID() *string {
	if o == nil {
		return nil
	}
	return o.UpiID
}
