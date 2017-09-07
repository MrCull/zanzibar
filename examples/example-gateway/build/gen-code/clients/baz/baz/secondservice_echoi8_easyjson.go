// Code generated by zanzibar
// @generated
// Checksum : 1zHHSA+gxega0XvVBsZfyA==
// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package baz

import (
	json "encoding/json"
	fmt "fmt"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson822f1558DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSecondServiceEchoI8(in *jlexer.Lexer, out *SecondService_EchoI8_Result) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "success":
			if in.IsNull() {
				in.Skip()
				out.Success = nil
			} else {
				if out.Success == nil {
					out.Success = new(int8)
				}
				*out.Success = int8(in.Int8())
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson822f1558EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSecondServiceEchoI8(out *jwriter.Writer, in SecondService_EchoI8_Result) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Success != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"success\":")
		if in.Success == nil {
			out.RawString("null")
		} else {
			out.Int8(int8(*in.Success))
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SecondService_EchoI8_Result) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson822f1558EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSecondServiceEchoI8(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SecondService_EchoI8_Result) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson822f1558EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSecondServiceEchoI8(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SecondService_EchoI8_Result) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson822f1558DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSecondServiceEchoI8(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SecondService_EchoI8_Result) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson822f1558DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSecondServiceEchoI8(l, v)
}
func easyjson822f1558DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSecondServiceEchoI81(in *jlexer.Lexer, out *SecondService_EchoI8_Args) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var ArgSet bool
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "arg":
			out.Arg = int8(in.Int8())
			ArgSet = true
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
	if !ArgSet {
		in.AddError(fmt.Errorf("key 'arg' is required"))
	}
}
func easyjson822f1558EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSecondServiceEchoI81(out *jwriter.Writer, in SecondService_EchoI8_Args) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"arg\":")
	out.Int8(int8(in.Arg))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SecondService_EchoI8_Args) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson822f1558EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSecondServiceEchoI81(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SecondService_EchoI8_Args) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson822f1558EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSecondServiceEchoI81(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SecondService_EchoI8_Args) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson822f1558DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSecondServiceEchoI81(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SecondService_EchoI8_Args) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson822f1558DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSecondServiceEchoI81(l, v)
}
