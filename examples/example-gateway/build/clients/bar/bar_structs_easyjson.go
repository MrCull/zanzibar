// AUTOGENERATED FILE: easyjson marshaller/unmarshallers.

package barClient

import (
	json "encoding/json"
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

func easyjsonCa6a3ed2DecodeGithubComUberZanzibarExamplesExampleGatewayBuildClientsBar(in *jlexer.Lexer, out *TooManyArgsHTTPRequest) {
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
		case "Request":
			(out.Request).UnmarshalEasyJSON(in)
		case "Foo":
			(out.Foo).UnmarshalEasyJSON(in)
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
func easyjsonCa6a3ed2EncodeGithubComUberZanzibarExamplesExampleGatewayBuildClientsBar(out *jwriter.Writer, in TooManyArgsHTTPRequest) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"Request\":")
	(in.Request).MarshalEasyJSON(out)
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"Foo\":")
	(in.Foo).MarshalEasyJSON(out)
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v TooManyArgsHTTPRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCa6a3ed2EncodeGithubComUberZanzibarExamplesExampleGatewayBuildClientsBar(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v TooManyArgsHTTPRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCa6a3ed2EncodeGithubComUberZanzibarExamplesExampleGatewayBuildClientsBar(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *TooManyArgsHTTPRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCa6a3ed2DecodeGithubComUberZanzibarExamplesExampleGatewayBuildClientsBar(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *TooManyArgsHTTPRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCa6a3ed2DecodeGithubComUberZanzibarExamplesExampleGatewayBuildClientsBar(l, v)
}
func easyjsonCa6a3ed2DecodeGithubComUberZanzibarExamplesExampleGatewayBuildClientsBar1(in *jlexer.Lexer, out *NormalHTTPRequest) {
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
		case "Request":
			(out.Request).UnmarshalEasyJSON(in)
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
func easyjsonCa6a3ed2EncodeGithubComUberZanzibarExamplesExampleGatewayBuildClientsBar1(out *jwriter.Writer, in NormalHTTPRequest) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"Request\":")
	(in.Request).MarshalEasyJSON(out)
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v NormalHTTPRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCa6a3ed2EncodeGithubComUberZanzibarExamplesExampleGatewayBuildClientsBar1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v NormalHTTPRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCa6a3ed2EncodeGithubComUberZanzibarExamplesExampleGatewayBuildClientsBar1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *NormalHTTPRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCa6a3ed2DecodeGithubComUberZanzibarExamplesExampleGatewayBuildClientsBar1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *NormalHTTPRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCa6a3ed2DecodeGithubComUberZanzibarExamplesExampleGatewayBuildClientsBar1(l, v)
}
func easyjsonCa6a3ed2DecodeGithubComUberZanzibarExamplesExampleGatewayBuildClientsBar2(in *jlexer.Lexer, out *ArgNotStructHTTPRequest) {
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
		case "Request":
			out.Request = string(in.String())
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
func easyjsonCa6a3ed2EncodeGithubComUberZanzibarExamplesExampleGatewayBuildClientsBar2(out *jwriter.Writer, in ArgNotStructHTTPRequest) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"Request\":")
	out.String(string(in.Request))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ArgNotStructHTTPRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCa6a3ed2EncodeGithubComUberZanzibarExamplesExampleGatewayBuildClientsBar2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ArgNotStructHTTPRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCa6a3ed2EncodeGithubComUberZanzibarExamplesExampleGatewayBuildClientsBar2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ArgNotStructHTTPRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCa6a3ed2DecodeGithubComUberZanzibarExamplesExampleGatewayBuildClientsBar2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ArgNotStructHTTPRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCa6a3ed2DecodeGithubComUberZanzibarExamplesExampleGatewayBuildClientsBar2(l, v)
}