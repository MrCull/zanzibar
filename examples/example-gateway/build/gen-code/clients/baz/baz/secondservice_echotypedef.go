// Code generated by thriftrw v1.6.0. DO NOT EDIT.
// @generated

package baz

import (
	"errors"
	"fmt"
	"github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/baz/base"
	"go.uber.org/thriftrw/wire"
	"strings"
)

type SecondService_EchoTypedef_Args struct {
	Arg base.UUID `json:"arg,required"`
}

func (v *SecondService_EchoTypedef_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	w, err = v.Arg.ToWire()
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _UUID_Read(w wire.Value) (base.UUID, error) {
	var x base.UUID
	err := x.FromWire(w)
	return x, err
}

func (v *SecondService_EchoTypedef_Args) FromWire(w wire.Value) error {
	var err error
	argIsSet := false
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.Arg, err = _UUID_Read(field.Value)
				if err != nil {
					return err
				}
				argIsSet = true
			}
		}
	}
	if !argIsSet {
		return errors.New("field Arg of SecondService_EchoTypedef_Args is required")
	}
	return nil
}

func (v *SecondService_EchoTypedef_Args) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [1]string
	i := 0
	fields[i] = fmt.Sprintf("Arg: %v", v.Arg)
	i++
	return fmt.Sprintf("SecondService_EchoTypedef_Args{%v}", strings.Join(fields[:i], ", "))
}

func (v *SecondService_EchoTypedef_Args) Equals(rhs *SecondService_EchoTypedef_Args) bool {
	if !(v.Arg == rhs.Arg) {
		return false
	}
	return true
}

func (v *SecondService_EchoTypedef_Args) MethodName() string {
	return "echoTypedef"
}

func (v *SecondService_EchoTypedef_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

var SecondService_EchoTypedef_Helper = struct {
	Args           func(arg base.UUID) *SecondService_EchoTypedef_Args
	IsException    func(error) bool
	WrapResponse   func(base.UUID, error) (*SecondService_EchoTypedef_Result, error)
	UnwrapResponse func(*SecondService_EchoTypedef_Result) (base.UUID, error)
}{}

func init() {
	SecondService_EchoTypedef_Helper.Args = func(arg base.UUID) *SecondService_EchoTypedef_Args {
		return &SecondService_EchoTypedef_Args{Arg: arg}
	}
	SecondService_EchoTypedef_Helper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}
	SecondService_EchoTypedef_Helper.WrapResponse = func(success base.UUID, err error) (*SecondService_EchoTypedef_Result, error) {
		if err == nil {
			return &SecondService_EchoTypedef_Result{Success: &success}, nil
		}
		return nil, err
	}
	SecondService_EchoTypedef_Helper.UnwrapResponse = func(result *SecondService_EchoTypedef_Result) (success base.UUID, err error) {
		if result.Success != nil {
			success = *result.Success
			return
		}
		err = errors.New("expected a non-void result")
		return
	}
}

type SecondService_EchoTypedef_Result struct {
	Success *base.UUID `json:"success,omitempty"`
}

func (v *SecondService_EchoTypedef_Result) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Success != nil {
		w, err = v.Success.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}
	if i != 1 {
		return wire.Value{}, fmt.Errorf("SecondService_EchoTypedef_Result should have exactly one field: got %v fields", i)
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *SecondService_EchoTypedef_Result) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TBinary {
				var x base.UUID
				x, err = _UUID_Read(field.Value)
				v.Success = &x
				if err != nil {
					return err
				}
			}
		}
	}
	count := 0
	if v.Success != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("SecondService_EchoTypedef_Result should have exactly one field: got %v fields", count)
	}
	return nil
}

func (v *SecondService_EchoTypedef_Result) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [1]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", *(v.Success))
		i++
	}
	return fmt.Sprintf("SecondService_EchoTypedef_Result{%v}", strings.Join(fields[:i], ", "))
}

func _UUID_EqualsPtr(lhs, rhs *base.UUID) bool {
	if lhs != nil && rhs != nil {
		x := *lhs
		y := *rhs
		return (x == y)
	}
	return lhs == nil && rhs == nil
}

func (v *SecondService_EchoTypedef_Result) Equals(rhs *SecondService_EchoTypedef_Result) bool {
	if !_UUID_EqualsPtr(v.Success, rhs.Success) {
		return false
	}
	return true
}

func (v *SecondService_EchoTypedef_Result) GetSuccess() (o base.UUID) {
	if v.Success != nil {
		return *v.Success
	}
	return
}

func (v *SecondService_EchoTypedef_Result) MethodName() string {
	return "echoTypedef"
}

func (v *SecondService_EchoTypedef_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
