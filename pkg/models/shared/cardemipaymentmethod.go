// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type CardEMIPaymentMethod struct {
	Emi *CardEMI `json:"emi,omitempty"`
}

func (o *CardEMIPaymentMethod) GetEmi() *CardEMI {
	if o == nil {
		return nil
	}
	return o.Emi
}
