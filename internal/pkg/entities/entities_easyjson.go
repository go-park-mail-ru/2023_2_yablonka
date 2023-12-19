// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package entities

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	dto "server/internal/pkg/dto"
	time "time"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson3e8ab7adDecodeServerInternalPkgEntities(in *jlexer.Lexer, out *Workspace) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = uint64(in.Uint64())
		case "name":
			out.Name = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "date_created":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.DateCreated).UnmarshalJSON(data))
			}
		case "users":
			if in.IsNull() {
				in.Skip()
				out.Users = nil
			} else {
				in.Delim('[')
				if out.Users == nil {
					if !in.IsDelim(']') {
						out.Users = make([]dto.UserPublicInfo, 0, 1)
					} else {
						out.Users = []dto.UserPublicInfo{}
					}
				} else {
					out.Users = (out.Users)[:0]
				}
				for !in.IsDelim(']') {
					var v1 dto.UserPublicInfo
					(v1).UnmarshalEasyJSON(in)
					out.Users = append(out.Users, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "boards":
			if in.IsNull() {
				in.Skip()
				out.Boards = nil
			} else {
				in.Delim('[')
				if out.Boards == nil {
					if !in.IsDelim(']') {
						out.Boards = make([]Board, 0, 0)
					} else {
						out.Boards = []Board{}
					}
				} else {
					out.Boards = (out.Boards)[:0]
				}
				for !in.IsDelim(']') {
					var v2 Board
					(v2).UnmarshalEasyJSON(in)
					out.Boards = append(out.Boards, v2)
					in.WantComma()
				}
				in.Delim(']')
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
func easyjson3e8ab7adEncodeServerInternalPkgEntities(out *jwriter.Writer, in Workspace) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.ID))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"date_created\":"
		out.RawString(prefix)
		out.Raw((in.DateCreated).MarshalJSON())
	}
	{
		const prefix string = ",\"users\":"
		out.RawString(prefix)
		if in.Users == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v3, v4 := range in.Users {
				if v3 > 0 {
					out.RawByte(',')
				}
				(v4).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"boards\":"
		out.RawString(prefix)
		if in.Boards == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Boards {
				if v5 > 0 {
					out.RawByte(',')
				}
				(v6).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Workspace) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3e8ab7adEncodeServerInternalPkgEntities(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Workspace) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3e8ab7adEncodeServerInternalPkgEntities(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Workspace) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3e8ab7adDecodeServerInternalPkgEntities(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Workspace) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3e8ab7adDecodeServerInternalPkgEntities(l, v)
}
func easyjson3e8ab7adDecodeServerInternalPkgEntities1(in *jlexer.Lexer, out *User) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "user_id":
			out.ID = uint64(in.Uint64())
		case "email":
			out.Email = string(in.String())
		case "password_hash":
			out.PasswordHash = string(in.String())
		case "name":
			if in.IsNull() {
				in.Skip()
				out.Name = nil
			} else {
				if out.Name == nil {
					out.Name = new(string)
				}
				*out.Name = string(in.String())
			}
		case "surname":
			if in.IsNull() {
				in.Skip()
				out.Surname = nil
			} else {
				if out.Surname == nil {
					out.Surname = new(string)
				}
				*out.Surname = string(in.String())
			}
		case "avatar_url":
			if in.IsNull() {
				in.Skip()
				out.AvatarURL = nil
			} else {
				if out.AvatarURL == nil {
					out.AvatarURL = new(string)
				}
				*out.AvatarURL = string(in.String())
			}
		case "description":
			if in.IsNull() {
				in.Skip()
				out.Description = nil
			} else {
				if out.Description == nil {
					out.Description = new(string)
				}
				*out.Description = string(in.String())
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
func easyjson3e8ab7adEncodeServerInternalPkgEntities1(out *jwriter.Writer, in User) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"user_id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.ID))
	}
	{
		const prefix string = ",\"email\":"
		out.RawString(prefix)
		out.String(string(in.Email))
	}
	{
		const prefix string = ",\"password_hash\":"
		out.RawString(prefix)
		out.String(string(in.PasswordHash))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		if in.Name == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Name))
		}
	}
	{
		const prefix string = ",\"surname\":"
		out.RawString(prefix)
		if in.Surname == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Surname))
		}
	}
	{
		const prefix string = ",\"avatar_url\":"
		out.RawString(prefix)
		if in.AvatarURL == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.AvatarURL))
		}
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		if in.Description == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Description))
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v User) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3e8ab7adEncodeServerInternalPkgEntities1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v User) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3e8ab7adEncodeServerInternalPkgEntities1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *User) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3e8ab7adDecodeServerInternalPkgEntities1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *User) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3e8ab7adDecodeServerInternalPkgEntities1(l, v)
}
func easyjson3e8ab7adDecodeServerInternalPkgEntities2(in *jlexer.Lexer, out *Task) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = uint64(in.Uint64())
		case "list_id":
			out.ListID = uint64(in.Uint64())
		case "date_created":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.DateCreated).UnmarshalJSON(data))
			}
		case "name":
			out.Name = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "list_position":
			out.ListPosition = uint64(in.Uint64())
		case "start":
			if in.IsNull() {
				in.Skip()
				out.Start = nil
			} else {
				if out.Start == nil {
					out.Start = new(time.Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.Start).UnmarshalJSON(data))
				}
			}
		case "end":
			if in.IsNull() {
				in.Skip()
				out.End = nil
			} else {
				if out.End == nil {
					out.End = new(time.Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.End).UnmarshalJSON(data))
				}
			}
		case "users":
			if in.IsNull() {
				in.Skip()
				out.Users = nil
			} else {
				in.Delim('[')
				if out.Users == nil {
					if !in.IsDelim(']') {
						out.Users = make([]uint64, 0, 8)
					} else {
						out.Users = []uint64{}
					}
				} else {
					out.Users = (out.Users)[:0]
				}
				for !in.IsDelim(']') {
					var v7 uint64
					v7 = uint64(in.Uint64())
					out.Users = append(out.Users, v7)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "checklists":
			if in.IsNull() {
				in.Skip()
				out.Checklists = nil
			} else {
				in.Delim('[')
				if out.Checklists == nil {
					if !in.IsDelim(']') {
						out.Checklists = make([]uint64, 0, 8)
					} else {
						out.Checklists = []uint64{}
					}
				} else {
					out.Checklists = (out.Checklists)[:0]
				}
				for !in.IsDelim(']') {
					var v8 uint64
					v8 = uint64(in.Uint64())
					out.Checklists = append(out.Checklists, v8)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "comments":
			if in.IsNull() {
				in.Skip()
				out.Comments = nil
			} else {
				in.Delim('[')
				if out.Comments == nil {
					if !in.IsDelim(']') {
						out.Comments = make([]uint64, 0, 8)
					} else {
						out.Comments = []uint64{}
					}
				} else {
					out.Comments = (out.Comments)[:0]
				}
				for !in.IsDelim(']') {
					var v9 uint64
					v9 = uint64(in.Uint64())
					out.Comments = append(out.Comments, v9)
					in.WantComma()
				}
				in.Delim(']')
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
func easyjson3e8ab7adEncodeServerInternalPkgEntities2(out *jwriter.Writer, in Task) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.ID))
	}
	{
		const prefix string = ",\"list_id\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.ListID))
	}
	{
		const prefix string = ",\"date_created\":"
		out.RawString(prefix)
		out.Raw((in.DateCreated).MarshalJSON())
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"list_position\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.ListPosition))
	}
	{
		const prefix string = ",\"start\":"
		out.RawString(prefix)
		if in.Start == nil {
			out.RawString("null")
		} else {
			out.Raw((*in.Start).MarshalJSON())
		}
	}
	{
		const prefix string = ",\"end\":"
		out.RawString(prefix)
		if in.End == nil {
			out.RawString("null")
		} else {
			out.Raw((*in.End).MarshalJSON())
		}
	}
	{
		const prefix string = ",\"users\":"
		out.RawString(prefix)
		if in.Users == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v10, v11 := range in.Users {
				if v10 > 0 {
					out.RawByte(',')
				}
				out.Uint64(uint64(v11))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"checklists\":"
		out.RawString(prefix)
		if in.Checklists == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v12, v13 := range in.Checklists {
				if v12 > 0 {
					out.RawByte(',')
				}
				out.Uint64(uint64(v13))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"comments\":"
		out.RawString(prefix)
		if in.Comments == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v14, v15 := range in.Comments {
				if v14 > 0 {
					out.RawByte(',')
				}
				out.Uint64(uint64(v15))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Task) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3e8ab7adEncodeServerInternalPkgEntities2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Task) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3e8ab7adEncodeServerInternalPkgEntities2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Task) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3e8ab7adDecodeServerInternalPkgEntities2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Task) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3e8ab7adDecodeServerInternalPkgEntities2(l, v)
}
func easyjson3e8ab7adDecodeServerInternalPkgEntities3(in *jlexer.Lexer, out *Session) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "SessionID":
			out.SessionID = string(in.String())
		case "UserID":
			out.UserID = uint64(in.Uint64())
		case "ExpiryDate":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.ExpiryDate).UnmarshalJSON(data))
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
func easyjson3e8ab7adEncodeServerInternalPkgEntities3(out *jwriter.Writer, in Session) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"SessionID\":"
		out.RawString(prefix[1:])
		out.String(string(in.SessionID))
	}
	{
		const prefix string = ",\"UserID\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.UserID))
	}
	{
		const prefix string = ",\"ExpiryDate\":"
		out.RawString(prefix)
		out.Raw((in.ExpiryDate).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Session) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3e8ab7adEncodeServerInternalPkgEntities3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Session) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3e8ab7adEncodeServerInternalPkgEntities3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Session) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3e8ab7adDecodeServerInternalPkgEntities3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Session) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3e8ab7adDecodeServerInternalPkgEntities3(l, v)
}
func easyjson3e8ab7adDecodeServerInternalPkgEntities4(in *jlexer.Lexer, out *Role) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = uint64(in.Uint64())
		case "name":
			out.Name = string(in.String())
		case "description":
			if in.IsNull() {
				in.Skip()
				out.Description = nil
			} else {
				if out.Description == nil {
					out.Description = new(string)
				}
				*out.Description = string(in.String())
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
func easyjson3e8ab7adEncodeServerInternalPkgEntities4(out *jwriter.Writer, in Role) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.ID))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		if in.Description == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Description))
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Role) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3e8ab7adEncodeServerInternalPkgEntities4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Role) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3e8ab7adEncodeServerInternalPkgEntities4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Role) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3e8ab7adDecodeServerInternalPkgEntities4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Role) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3e8ab7adDecodeServerInternalPkgEntities4(l, v)
}
func easyjson3e8ab7adDecodeServerInternalPkgEntities5(in *jlexer.Lexer, out *QuestionType) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = uint64(in.Uint64())
		case "name":
			out.Name = string(in.String())
		case "max_rating":
			out.MaxRating = uint64(in.Uint64())
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
func easyjson3e8ab7adEncodeServerInternalPkgEntities5(out *jwriter.Writer, in QuestionType) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.ID))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"max_rating\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.MaxRating))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v QuestionType) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3e8ab7adEncodeServerInternalPkgEntities5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v QuestionType) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3e8ab7adEncodeServerInternalPkgEntities5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *QuestionType) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3e8ab7adDecodeServerInternalPkgEntities5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *QuestionType) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3e8ab7adDecodeServerInternalPkgEntities5(l, v)
}
func easyjson3e8ab7adDecodeServerInternalPkgEntities6(in *jlexer.Lexer, out *List) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = uint64(in.Uint64())
		case "board_id":
			out.BoardID = uint64(in.Uint64())
		case "name":
			out.Name = string(in.String())
		case "description":
			if in.IsNull() {
				in.Skip()
				out.Description = nil
			} else {
				if out.Description == nil {
					out.Description = new(string)
				}
				*out.Description = string(in.String())
			}
		case "list_position":
			out.ListPosition = uint64(in.Uint64())
		case "tasks":
			if in.IsNull() {
				in.Skip()
				out.Tasks = nil
			} else {
				in.Delim('[')
				if out.Tasks == nil {
					if !in.IsDelim(']') {
						out.Tasks = make([]Task, 0, 0)
					} else {
						out.Tasks = []Task{}
					}
				} else {
					out.Tasks = (out.Tasks)[:0]
				}
				for !in.IsDelim(']') {
					var v16 Task
					(v16).UnmarshalEasyJSON(in)
					out.Tasks = append(out.Tasks, v16)
					in.WantComma()
				}
				in.Delim(']')
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
func easyjson3e8ab7adEncodeServerInternalPkgEntities6(out *jwriter.Writer, in List) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.ID))
	}
	{
		const prefix string = ",\"board_id\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.BoardID))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		if in.Description == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Description))
		}
	}
	{
		const prefix string = ",\"list_position\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.ListPosition))
	}
	{
		const prefix string = ",\"tasks\":"
		out.RawString(prefix)
		if in.Tasks == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v17, v18 := range in.Tasks {
				if v17 > 0 {
					out.RawByte(',')
				}
				(v18).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v List) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3e8ab7adEncodeServerInternalPkgEntities6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v List) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3e8ab7adEncodeServerInternalPkgEntities6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *List) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3e8ab7adDecodeServerInternalPkgEntities6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *List) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3e8ab7adDecodeServerInternalPkgEntities6(l, v)
}
func easyjson3e8ab7adDecodeServerInternalPkgEntities7(in *jlexer.Lexer, out *Comment) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = uint64(in.Uint64())
		case "user_id":
			out.UserID = uint64(in.Uint64())
		case "task_id":
			out.TaskID = uint64(in.Uint64())
		case "text":
			out.Text = string(in.String())
		case "date_created":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.DateCreated).UnmarshalJSON(data))
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
func easyjson3e8ab7adEncodeServerInternalPkgEntities7(out *jwriter.Writer, in Comment) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.ID))
	}
	{
		const prefix string = ",\"user_id\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.UserID))
	}
	{
		const prefix string = ",\"task_id\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.TaskID))
	}
	{
		const prefix string = ",\"text\":"
		out.RawString(prefix)
		out.String(string(in.Text))
	}
	{
		const prefix string = ",\"date_created\":"
		out.RawString(prefix)
		out.Raw((in.DateCreated).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Comment) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3e8ab7adEncodeServerInternalPkgEntities7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Comment) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3e8ab7adEncodeServerInternalPkgEntities7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Comment) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3e8ab7adDecodeServerInternalPkgEntities7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Comment) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3e8ab7adDecodeServerInternalPkgEntities7(l, v)
}
func easyjson3e8ab7adDecodeServerInternalPkgEntities8(in *jlexer.Lexer, out *ChecklistItem) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = uint64(in.Uint64())
		case "checklist_id":
			out.ChecklistID = uint64(in.Uint64())
		case "name":
			out.Name = string(in.String())
		case "done":
			out.Done = bool(in.Bool())
		case "list_position":
			out.ListPosition = uint64(in.Uint64())
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
func easyjson3e8ab7adEncodeServerInternalPkgEntities8(out *jwriter.Writer, in ChecklistItem) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.ID))
	}
	{
		const prefix string = ",\"checklist_id\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.ChecklistID))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"done\":"
		out.RawString(prefix)
		out.Bool(bool(in.Done))
	}
	{
		const prefix string = ",\"list_position\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.ListPosition))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ChecklistItem) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3e8ab7adEncodeServerInternalPkgEntities8(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ChecklistItem) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3e8ab7adEncodeServerInternalPkgEntities8(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ChecklistItem) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3e8ab7adDecodeServerInternalPkgEntities8(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ChecklistItem) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3e8ab7adDecodeServerInternalPkgEntities8(l, v)
}
func easyjson3e8ab7adDecodeServerInternalPkgEntities9(in *jlexer.Lexer, out *Checklist) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = uint64(in.Uint64())
		case "task_id":
			out.TaskID = uint64(in.Uint64())
		case "name":
			out.Name = string(in.String())
		case "list_position":
			out.ListPosition = uint64(in.Uint64())
		case "items":
			if in.IsNull() {
				in.Skip()
				out.Items = nil
			} else {
				in.Delim('[')
				if out.Items == nil {
					if !in.IsDelim(']') {
						out.Items = make([]uint64, 0, 8)
					} else {
						out.Items = []uint64{}
					}
				} else {
					out.Items = (out.Items)[:0]
				}
				for !in.IsDelim(']') {
					var v19 uint64
					v19 = uint64(in.Uint64())
					out.Items = append(out.Items, v19)
					in.WantComma()
				}
				in.Delim(']')
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
func easyjson3e8ab7adEncodeServerInternalPkgEntities9(out *jwriter.Writer, in Checklist) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.ID))
	}
	{
		const prefix string = ",\"task_id\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.TaskID))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"list_position\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.ListPosition))
	}
	{
		const prefix string = ",\"items\":"
		out.RawString(prefix)
		if in.Items == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v20, v21 := range in.Items {
				if v20 > 0 {
					out.RawByte(',')
				}
				out.Uint64(uint64(v21))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Checklist) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3e8ab7adEncodeServerInternalPkgEntities9(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Checklist) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3e8ab7adEncodeServerInternalPkgEntities9(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Checklist) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3e8ab7adDecodeServerInternalPkgEntities9(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Checklist) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3e8ab7adDecodeServerInternalPkgEntities9(l, v)
}
func easyjson3e8ab7adDecodeServerInternalPkgEntities10(in *jlexer.Lexer, out *CSRF) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Token":
			out.Token = string(in.String())
		case "UserID":
			out.UserID = uint64(in.Uint64())
		case "ExpirationDate":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.ExpirationDate).UnmarshalJSON(data))
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
func easyjson3e8ab7adEncodeServerInternalPkgEntities10(out *jwriter.Writer, in CSRF) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Token\":"
		out.RawString(prefix[1:])
		out.String(string(in.Token))
	}
	{
		const prefix string = ",\"UserID\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.UserID))
	}
	{
		const prefix string = ",\"ExpirationDate\":"
		out.RawString(prefix)
		out.Raw((in.ExpirationDate).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CSRF) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3e8ab7adEncodeServerInternalPkgEntities10(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CSRF) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3e8ab7adEncodeServerInternalPkgEntities10(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CSRF) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3e8ab7adDecodeServerInternalPkgEntities10(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CSRF) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3e8ab7adDecodeServerInternalPkgEntities10(l, v)
}
func easyjson3e8ab7adDecodeServerInternalPkgEntities11(in *jlexer.Lexer, out *CSATQuestion) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = uint64(in.Uint64())
		case "type":
			out.TypeID = string(in.String())
		case "content":
			out.Content = string(in.String())
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
func easyjson3e8ab7adEncodeServerInternalPkgEntities11(out *jwriter.Writer, in CSATQuestion) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.ID))
	}
	{
		const prefix string = ",\"type\":"
		out.RawString(prefix)
		out.String(string(in.TypeID))
	}
	{
		const prefix string = ",\"content\":"
		out.RawString(prefix)
		out.String(string(in.Content))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CSATQuestion) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3e8ab7adEncodeServerInternalPkgEntities11(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CSATQuestion) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3e8ab7adEncodeServerInternalPkgEntities11(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CSATQuestion) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3e8ab7adDecodeServerInternalPkgEntities11(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CSATQuestion) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3e8ab7adDecodeServerInternalPkgEntities11(l, v)
}
func easyjson3e8ab7adDecodeServerInternalPkgEntities12(in *jlexer.Lexer, out *Board) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "board_id":
			out.ID = uint64(in.Uint64())
		case "name":
			out.Name = string(in.String())
		case "owner":
			(out.Owner).UnmarshalEasyJSON(in)
		case "thumbnail_url":
			if in.IsNull() {
				in.Skip()
				out.ThumbnailURL = nil
			} else {
				if out.ThumbnailURL == nil {
					out.ThumbnailURL = new(string)
				}
				*out.ThumbnailURL = string(in.String())
			}
		case "date_created":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.DateCreated).UnmarshalJSON(data))
			}
		case "users":
			if in.IsNull() {
				in.Skip()
				out.Users = nil
			} else {
				in.Delim('[')
				if out.Users == nil {
					if !in.IsDelim(']') {
						out.Users = make([]dto.UserPublicInfo, 0, 1)
					} else {
						out.Users = []dto.UserPublicInfo{}
					}
				} else {
					out.Users = (out.Users)[:0]
				}
				for !in.IsDelim(']') {
					var v22 dto.UserPublicInfo
					(v22).UnmarshalEasyJSON(in)
					out.Users = append(out.Users, v22)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "lists":
			if in.IsNull() {
				in.Skip()
				out.Lists = nil
			} else {
				in.Delim('[')
				if out.Lists == nil {
					if !in.IsDelim(']') {
						out.Lists = make([]List, 0, 0)
					} else {
						out.Lists = []List{}
					}
				} else {
					out.Lists = (out.Lists)[:0]
				}
				for !in.IsDelim(']') {
					var v23 List
					(v23).UnmarshalEasyJSON(in)
					out.Lists = append(out.Lists, v23)
					in.WantComma()
				}
				in.Delim(']')
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
func easyjson3e8ab7adEncodeServerInternalPkgEntities12(out *jwriter.Writer, in Board) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"board_id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.ID))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"owner\":"
		out.RawString(prefix)
		(in.Owner).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"thumbnail_url\":"
		out.RawString(prefix)
		if in.ThumbnailURL == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.ThumbnailURL))
		}
	}
	{
		const prefix string = ",\"date_created\":"
		out.RawString(prefix)
		out.Raw((in.DateCreated).MarshalJSON())
	}
	{
		const prefix string = ",\"users\":"
		out.RawString(prefix)
		if in.Users == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v24, v25 := range in.Users {
				if v24 > 0 {
					out.RawByte(',')
				}
				(v25).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"lists\":"
		out.RawString(prefix)
		if in.Lists == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v26, v27 := range in.Lists {
				if v26 > 0 {
					out.RawByte(',')
				}
				(v27).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Board) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3e8ab7adEncodeServerInternalPkgEntities12(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Board) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3e8ab7adEncodeServerInternalPkgEntities12(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Board) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3e8ab7adDecodeServerInternalPkgEntities12(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Board) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3e8ab7adDecodeServerInternalPkgEntities12(l, v)
}
