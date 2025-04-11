package vismanet

import (
	"encoding/json"
	"time"
)

// Value is a wrapper for any type
type Value struct {
	Value interface{} `json:"value"`
}

// Time is a wrapper for the time.Time type that parses the time in the formats used by Visma.net
type Time time.Time

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
	*t = Time(pt)

	return nil
}

// TimeValue is a wrapper for the time.Time type that when marshaled to JSON will be wrapped in a Value struct
type TimeValue time.Time

// MarshalJSON wraps the Time in a Value struct and marshals it into a JSON byte slice
func (v TimeValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(Value{time.Time(v)})
}

func (t *TimeValue) UnmarshalJSON(data []byte) error {
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
	*t = TimeValue(pt)

	return nil
}

// NewStringValue creates a new StringValue and returns a pointer to it
func NewStringValue(value string) *StringValue {
	v := new(StringValue)
	*v = StringValue(value)
	return v
}

// StringValue is a wrapper for the string type and
// when marshaled to JSON will be wrapped in a Value struct
type StringValue string

// MarshalJSON marshals the StringValue into a JSON byte slice
func (v StringValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(Value{string(v)})
}

// NewBoolValue creates a new BoolValue and returns a pointer to it
func NewBoolValue(value bool) *BoolValue {
	v := new(BoolValue)
	*v = BoolValue(value)
	return v
}

// BoolValue is a wrapper for the bool type and
// when marshaled to JSON will be wrapped in a Value struct
type BoolValue bool

// MarshalJSON marshals the BoolValue into a JSON byte slice
func (v BoolValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(Value{bool(v)})
}

// NewIntValue creates a new IntValue and returns a pointer to it
func NewIntValue(value int) *IntValue {
	v := new(IntValue)
	*v = IntValue(value)
	return v
}

// IntValue is a wrapper for the int type and
// when marshaled to JSON will be wrapped in a Value struct
type IntValue int

// MarshalJSON marshals the IntValue into a JSON byte slice
func (v IntValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(Value{int(v)})
}

// NewFloatValue creates a new FloatValue and returns a pointer to it
func NewFloatValue(value float64) *FloatValue {
	v := new(FloatValue)
	*v = FloatValue(value)
	return v
}

// FloatValue is a wrapper for the float64 type and
// when marshaled to JSON will be wrapped in a Value struct
type FloatValue float64

// MarshalJSON marshals the FloatValue into a JSON byte slice
func (v FloatValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(Value{float64(v)})
}
