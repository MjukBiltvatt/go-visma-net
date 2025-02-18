package vismanet

import (
	"encoding/json"
	"reflect"
	"slices"
	"strings"
	"time"
)

// Determine if the specified struct field is required
func jsonFieldIsRequired(field reflect.StructField) bool {
	tag := field.Tag.Get("json")
	values := strings.Split(tag, ",")
	return slices.Contains(values, "required")
}

type TypeValue interface {
	IsEmpty() bool
	MarshalJSON() ([]byte, error)
	UnmarshalJSON([]byte) error
}

// Value is a wrapper for any type, implementing the TypeValue interface
type Value struct {
	Value interface{} `json:"value"`
}

// TimeValue is a wrapper for the Time type, implementing the TypeValue interface and
// when marshaled to JSON will be wrapped in a Value struct
type TimeValue Time

// MarshalJSON marshals the TimeValue into a JSON byte slice
func (v TimeValue) MarshalJSON() ([]byte, error) {
	// if v.IsEmpty() {
	// 	return []byte{}, nil
	// }
	return json.Marshal(Value{Time(v)})
}

// UnmarshalJSON unmarshals a JSON byte slice into the Time
func (v *TimeValue) UnmarshalJSON(data []byte) error {
	//Unmarshal the data into a string
	var value string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	} else if value == "" {
		return nil
	}

	//Parse the time
	pt, err := parseTime(value)
	if err != nil {
		return err
	}
	*v = TimeValue{pt}

	return nil
}

// StringValue is a wrapper for the string type, implementing the TypeValue interface and
// when marshaled to JSON will be wrapped in a Value struct
type StringValue string

// IsEmpty returns true if the StringValue is empty
func (v StringValue) IsEmpty() bool {
	return v == ""
}

// NullableStringValue is a wrapper for the string type that can be null, implementing the TypeValue
// interface and when marshaled to JSON will be wrapped in a Value struct
type NullableStringValue struct {
	*StringValue
}

// MarshalJSON marshals the NullableStringValue into a JSON byte slice
func (v NullableStringValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(Value{v.StringValue})
}

// UnmarshalJSON unmarshals a JSON byte slice into the NullableStringValue
func (v *NullableStringValue) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &v.StringValue)
}

// IsEmpty returns true if the NullableStringValue is empty
func (v NullableStringValue) IsEmpty() bool {
	return v.StringValue == nil || *v.StringValue == ""
}

// BoolValue is a wrapper for the bool type, implementing the TypeValue interface and
// when marshaled to JSON will be wrapped in a Value struct
type BoolValue bool

// MarshalJSON marshals the BoolValue into a JSON byte slice
func (v BoolValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(Value{bool(v)})
}

// IntValue is a wrapper for the int type, implementing the TypeValue interface and
// when marshaled to JSON will be wrapped in a Value struct
type IntValue int

// MarshalJSON marshals the IntValue into a JSON byte slice
func (v IntValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(Value{int(v)})
}

// IsEmpty returns true if the IntValue is empty
func (v IntValue) IsEmpty() bool {
	return v == 0
}

// FloatValue is a wrapper for the float64 type, implementing the TypeValue interface and
// when marshaled to JSON will be wrapped in a Value struct
type FloatValue float64

// MarshalJSON marshals the FloatValue into a JSON byte slice
func (v FloatValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(Value{float64(v)})
}

// IsEmpty returns true if the FloatValue is empty
func (v FloatValue) IsEmpty() bool {
	return v == 0.0
}

// TODO: Implement MarshalJSON method? currently when time unmarshals and marshals back a Z is added after string
// Time is a wrapper for the time.Time type with custom JSON marshaling to handle
// the different time formats used by Visma
type Time struct {
	time.Time
}

// IsEmpty returns true if the Time is empty
func (t Time) IsEmpty() bool {
	return reflect.ValueOf(t).IsNil() || t.IsZero()
}

// UnmarshalJSON unmarshals a JSON byte slice into the Time
func (t *Time) UnmarshalJSON(data []byte) error {
	//Unmarshal the data into a string
	var value string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	} else if value == "" {
		return nil
	}

	//Parse the time
	pt, err := parseTime(value)
	if err != nil {
		return err
	}
	*t = Time{pt}

	return nil
}

// AddressValue is a wrapper for the Address type, implementing the TypeValue interface and
// when marshaled to JSON will be wrapped in a Value struct
type AddressValue Address

// MarshalJSON marshals the Address into a JSON byte slice
func (v AddressValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(Value{Address(v)})
}

// IsEmpty returns true if the Address is empty
func (a AddressValue) IsEmpty() bool {
	return a == AddressValue{}
}

// ContactValue is a wrapper for the Contact type, implementing the TypeValue interface and
// when marshaled to JSON will be wrapped in a Value struct
type ContactValue Contact

// MarshalJSON marshals the Contact into a JSON byte slice
func (v ContactValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(Value{Contact(v)})
}

// IsEmpty returns true if the Contact is empty
func (a ContactValue) IsEmpty() bool {
	return a == ContactValue{}
}
