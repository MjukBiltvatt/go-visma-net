package vismanet

import (
	"encoding/json"
	"errors"
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

// MarshalJSON marshals the Value into a JSON byte slice, omitting empty values if the struct field has the omitempty tag
func (v Value) MarshalJSON() ([]byte, error) {
	//Reflect the value and type of the value
	val := reflect.ValueOf(v)
	vt := reflect.TypeOf(v)

	//Get the fields of the struct
	fields := []reflect.StructField{}
	for i := 0; i < vt.NumField(); i++ {
		fields = append(fields, vt.Field(i))
	}

	//Iterate over the fields of the struct
	for i := range fields {
		//Get values of the json tag
		var jsonTagValues []string
		if tagValue, ok := fields[i].Tag.Lookup("json"); ok {
			jsonTagValues = strings.Split(tagValue, ",")
		}

		//Skip if the json tag does not contain omitempty
		if !slices.Contains(jsonTagValues, "omitempty") {
			continue
		}

		//Replace json struct tag if the value is empty
		if tv, ok := val.Field(i).Interface().(TypeValue); ok && tv.IsEmpty() {
			old := `json:"` + fields[i].Tag.Get("json") + `"`
			s := strings.ReplaceAll(string(fields[i].Tag), old, `json:"-"`)
			fields[i].Tag = reflect.StructTag(s)
		} else {
			return nil, errors.New("struct field with omitempty tag must implement TypeValue interface")
		}
	}

	//Create and marshal a new struct with the modified fields
	return json.Marshal(val.Convert(reflect.StructOf(fields)).Interface())
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

// TODO: Implemented MarshalJSON method? currently when time unmarshals and marshals back a Z is added after string
// Time is a wrapper for the time.Time type with custom JSON marshaling to handle
// the different time formats used by Visma
type Time struct {
	time.Time
}

// IsEmpty returns true if the TimeValue is empty
func (t Time) IsEmpty() bool {
	return reflect.ValueOf(t).IsNil() || t.IsZero()
}

// UnmarshalJSON unmarshals a JSON byte slice into the TimeValue
func (t *Time) UnmarshalJSON(data []byte) error {
	//Unmarshal the data into a string
	var value string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	} else if value == "" {
		return nil
	}

	//For some reason Visma uses different time formats, so we need to try multiple formats
	formats := []string{
		time.RFC3339,
		"2006-01-02T15:04:05",
		"2006-01-02T15:04:05-07:00",
	}

	//Try to parse the time with the different formats
	var err error
	for _, format := range formats {
		var pt time.Time
		pt, err = time.Parse(format, value)
		if err == nil {
			//If we successfully parsed the time, set the TimeValue to the parsed time
			*t = Time{pt}
			return nil
		}
	}

	//If we couldn't parse the time, return the last error
	return err
}
