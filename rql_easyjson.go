// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package rql

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

func easyjson4bc42f5bDecodeGithubComA8mXRql(in *jlexer.Lexer, out *Query) {
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
		case "limit":
			out.Limit = int(in.Int())
		case "offset":
			out.Offset = int(in.Int())
		case "sort":
			if in.IsNull() {
				in.Skip()
				out.Sort = nil
			} else {
				in.Delim('[')
				if out.Sort == nil {
					if !in.IsDelim(']') {
						out.Sort = make([]string, 0, 4)
					} else {
						out.Sort = []string{}
					}
				} else {
					out.Sort = (out.Sort)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Sort = append(out.Sort, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "filter":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Filter = make(map[string]interface{})
				} else {
					out.Filter = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v2 interface{}
					if m, ok := v2.(easyjson.Unmarshaler); ok {
						m.UnmarshalEasyJSON(in)
					} else if m, ok := v2.(json.Unmarshaler); ok {
						_ = m.UnmarshalJSON(in.Raw())
					} else {
						v2 = in.Interface()
					}
					(out.Filter)[key] = v2
					in.WantComma()
				}
				in.Delim('}')
			}
		default:
			in.AddError(&jlexer.LexerError{
				Offset: in.GetPos(),
				Reason: "unknown field",
				Data:   key,
			})
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson4bc42f5bEncodeGithubComA8mXRql(out *jwriter.Writer, in Query) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Limit != 0 {
		const prefix string = ",\"limit\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Limit))
	}
	if in.Offset != 0 {
		const prefix string = ",\"offset\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Offset))
	}
	if len(in.Sort) != 0 {
		const prefix string = ",\"sort\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v3, v4 := range in.Sort {
				if v3 > 0 {
					out.RawByte(',')
				}
				out.String(string(v4))
			}
			out.RawByte(']')
		}
	}
	if len(in.Filter) != 0 {
		const prefix string = ",\"filter\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('{')
			v5First := true
			for v5Name, v5Value := range in.Filter {
				if v5First {
					v5First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v5Name))
				out.RawByte(':')
				if m, ok := v5Value.(easyjson.Marshaler); ok {
					m.MarshalEasyJSON(out)
				} else if m, ok := v5Value.(json.Marshaler); ok {
					out.Raw(m.MarshalJSON())
				} else {
					out.Raw(json.Marshal(v5Value))
				}
			}
			out.RawByte('}')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Query) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson4bc42f5bEncodeGithubComA8mXRql(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Query) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson4bc42f5bEncodeGithubComA8mXRql(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Query) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson4bc42f5bDecodeGithubComA8mXRql(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Query) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson4bc42f5bDecodeGithubComA8mXRql(l, v)
}
