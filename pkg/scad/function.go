package scad

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

// Function describes a SCAD function call.
type Function struct {
	// Name is the function name, such as "cube", "cylinder", "translate".
	Name string

	// Parameters is a map of parameter names to parameter values in string form.
	Parameters map[string]string

	// Children is a slice of Function objects that are children of the Function.
	Children []Function
}

// SetParameter sets the parameter with the given key to the given value. A boolean
// is returned indicating if the parameter was replaced (already had a value).
func (fn *Function) SetParameter(key string, value string) bool {
	if fn.Parameters == nil {
		fn.Parameters = map[string]string{}
	}

	_, replaced := fn.Parameters[key]

	fn.Parameters[key] = value

	return replaced
}

// parametersString returns the Function's Parameters as a string suitable
// to use when calling the function in .scad.
func (fn Function) parametersString() string {
	paramKeys := make([]string, 0, len(fn.Parameters))
	for key := range fn.Parameters {
		paramKeys = append(paramKeys, key)
	}
	sort.Strings(paramKeys)

	params := make([]string, len(fn.Parameters))
	for i, key := range paramKeys {
		params[i] = fmt.Sprintf("%s=%s", key, fn.Parameters[key])
	}

	return strings.Join(params, ", ")
}
// FunctionNameGetter is the interface for types that implement GetFunctionName.
type FunctionNameGetter interface {
	// GetFunctionName returns a string representing the function name. The returned
	// string may be empty.
	GetFunctionName() string
}

// ParameterValueGetter is the interface for types that implement GetParameterValue.
type ParameterValueGetter interface {
	// GetParameterValue returns a string representing the value, and a boolean indicating
	// if the value is set.
	GetParameterValue() (string, bool)
}

// EncodeFunction encodes an interface into a Function. The given interface must be a struct.
//
// The resulting Function will have values applied based on the struct fields implementing
// one or more of these interfaces:
//
// FunctionNameGetter fields will determine the Function's Name value as the first non-empty
// value of these options:
//
// •Return value of the field's GetFunctionName method
//
// •The value of the field's "scad" tag
//
// •The lowercased field name
//
// ParameterValueGetter fields will set Parameter values for the Function. The Parameter's
// key will be the last non-empty value of these options:
//
// •The value of the field's "scad" tag
//
// •The lowercased field name
//
// Slice fields set the Children values for the Function.
//
// An error will be returned if:
//
// • Exactly one FunctionNameGetter field is not found
//
// • Multiple field set a value for the same Parameter key (multiple fields can have the same key name,
//   as long no more than one sets a value)
//
// • Multiple Children fields are found
func EncodeFunction(i interface{}) (Function, error) {
	var fn Function

	if i == nil {
		return Function{}, fmt.Errorf("scad: attempted to encode null value to Function")
	}

	// be nice and dereference pointers
	iV := reflect.ValueOf(i)
	if iV.Kind() == reflect.Ptr {
		iV = iV.Elem()
	}
	iT := iV.Type()

	if iV.Kind() != reflect.Struct {
		return Function{}, fmt.Errorf("scad: attempted to encode non-struct (%T) to Function", i)
	}

	for i := 0; i < iT.NumField(); i++ {
		field := iT.Field(i)

		scadName := field.Tag.Get("scad")
		if scadName == "" {
			scadName = strings.ToLower(field.Name)
		}

		fieldV := iV.Field(i)
		fieldT := fieldV.Type()

		// Name
		if fieldT.Implements(reflect.TypeOf((*FunctionNameGetter)(nil)).Elem()) {
			if fn.Name != "" {
				return Function{}, fmt.Errorf("scad: attempted to encode type (%T) with multiple FunctionNameGetter fields", i)
			}

			fnName := scadName

			if field.IsExported() {
				gotFnName := fieldV.Interface().(FunctionNameGetter).GetFunctionName()
				if gotFnName != "" {
					fnName = gotFnName
				}
			}

			fn.Name = fnName
		}

		// Parameters
		if fieldT.Implements(reflect.TypeOf((*ParameterValueGetter)(nil)).Elem()) {
			// silently ignore unexported ParameterValueGetter fields
			if !field.IsExported() {
				continue
			}

			if gotValue, ok := fieldV.Interface().(ParameterValueGetter).GetParameterValue(); ok {
				if replaced := fn.SetParameter(scadName, gotValue); replaced {
					return Function{}, fmt.Errorf("scad: attempted to encode type (%T) with multiple ParameterValueGetter fields with the same name: %s", i, scadName)
				}
			}
		}

		// Children
		if fieldV.Kind() == reflect.Slice {
			// silently ignore unexported slices
			if !field.IsExported() {
				continue
			}

			if fn.Children != nil {
				return Function{}, fmt.Errorf("scad: attempted to encode type (%T) with multiple slice fields", i)
			}

			children := make([]Function, fieldV.Len())

			for i := 0; i < fieldV.Len(); i++ {
				child, err := EncodeFunction(fieldV.Index(i).Interface())
				if err != nil {
					return Function{}, err
				}

				children[i] = child
			}
			fn.Children = children
		}
	}

	if fn.Name == "" {
		return Function{}, fmt.Errorf("scad: attempted to encode type (%T) without FunctionNameGetter field", i)
	}

	return fn, nil
}
