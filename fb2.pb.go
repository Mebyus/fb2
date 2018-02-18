// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fb2.proto

/*
Package fb2 is a generated protocol buffer package.

It is generated from these files:
	fb2.proto

It has these top-level messages:
	PFB2
	Description
	TitleInfo
	Coverpage
	Image
	AuthorType
*/
package fb2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PFB2 struct {
	ID          string       `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	Description *Description `protobuf:"bytes,2,opt,name=Description" json:"Description,omitempty"`
}

func (m *PFB2) Reset()                    { *m = PFB2{} }
func (m *PFB2) String() string            { return proto.CompactTextString(m) }
func (*PFB2) ProtoMessage()               {}
func (*PFB2) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PFB2) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *PFB2) GetDescription() *Description {
	if m != nil {
		return m.Description
	}
	return nil
}

type Description struct {
	TitleInfo *TitleInfo `protobuf:"bytes,1,opt,name=TitleInfo" json:"TitleInfo,omitempty"`
}

func (m *Description) Reset()                    { *m = Description{} }
func (m *Description) String() string            { return proto.CompactTextString(m) }
func (*Description) ProtoMessage()               {}
func (*Description) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Description) GetTitleInfo() *TitleInfo {
	if m != nil {
		return m.TitleInfo
	}
	return nil
}

type TitleInfo struct {
	Genre      []string      `protobuf:"bytes,1,rep,name=Genre" json:"Genre,omitempty"`
	GenreType  []string      `protobuf:"bytes,2,rep,name=GenreType" json:"GenreType,omitempty"`
	Author     []*AuthorType `protobuf:"bytes,3,rep,name=Author" json:"Author,omitempty"`
	BookTitle  string        `protobuf:"bytes,4,opt,name=BookTitle" json:"BookTitle,omitempty"`
	Annotation string        `protobuf:"bytes,5,opt,name=Annotation" json:"Annotation,omitempty"`
	Keywords   string        `protobuf:"bytes,6,opt,name=Keywords" json:"Keywords,omitempty"`
	Date       string        `protobuf:"bytes,7,opt,name=Date" json:"Date,omitempty"`
	Coverpage  *Coverpage    `protobuf:"bytes,8,opt,name=Coverpage" json:"Coverpage,omitempty"`
	Lang       string        `protobuf:"bytes,9,opt,name=Lang" json:"Lang,omitempty"`
	SrcLang    string        `protobuf:"bytes,10,opt,name=SrcLang" json:"SrcLang,omitempty"`
	Translator *AuthorType   `protobuf:"bytes,11,opt,name=Translator" json:"Translator,omitempty"`
	Sequence   string        `protobuf:"bytes,12,opt,name=Sequence" json:"Sequence,omitempty"`
}

func (m *TitleInfo) Reset()                    { *m = TitleInfo{} }
func (m *TitleInfo) String() string            { return proto.CompactTextString(m) }
func (*TitleInfo) ProtoMessage()               {}
func (*TitleInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *TitleInfo) GetGenre() []string {
	if m != nil {
		return m.Genre
	}
	return nil
}

func (m *TitleInfo) GetGenreType() []string {
	if m != nil {
		return m.GenreType
	}
	return nil
}

func (m *TitleInfo) GetAuthor() []*AuthorType {
	if m != nil {
		return m.Author
	}
	return nil
}

func (m *TitleInfo) GetBookTitle() string {
	if m != nil {
		return m.BookTitle
	}
	return ""
}

func (m *TitleInfo) GetAnnotation() string {
	if m != nil {
		return m.Annotation
	}
	return ""
}

func (m *TitleInfo) GetKeywords() string {
	if m != nil {
		return m.Keywords
	}
	return ""
}

func (m *TitleInfo) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

func (m *TitleInfo) GetCoverpage() *Coverpage {
	if m != nil {
		return m.Coverpage
	}
	return nil
}

func (m *TitleInfo) GetLang() string {
	if m != nil {
		return m.Lang
	}
	return ""
}

func (m *TitleInfo) GetSrcLang() string {
	if m != nil {
		return m.SrcLang
	}
	return ""
}

func (m *TitleInfo) GetTranslator() *AuthorType {
	if m != nil {
		return m.Translator
	}
	return nil
}

func (m *TitleInfo) GetSequence() string {
	if m != nil {
		return m.Sequence
	}
	return ""
}

type Coverpage struct {
	Image *Image `protobuf:"bytes,1,opt,name=Image" json:"Image,omitempty"`
}

func (m *Coverpage) Reset()                    { *m = Coverpage{} }
func (m *Coverpage) String() string            { return proto.CompactTextString(m) }
func (*Coverpage) ProtoMessage()               {}
func (*Coverpage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Coverpage) GetImage() *Image {
	if m != nil {
		return m.Image
	}
	return nil
}

type Image struct {
	Href string `protobuf:"bytes,1,opt,name=Href" json:"Href,omitempty"`
}

func (m *Image) Reset()                    { *m = Image{} }
func (m *Image) String() string            { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()               {}
func (*Image) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Image) GetHref() string {
	if m != nil {
		return m.Href
	}
	return ""
}

type AuthorType struct {
	FirstName  string `protobuf:"bytes,1,opt,name=FirstName" json:"FirstName,omitempty"`
	MiddleName string `protobuf:"bytes,2,opt,name=MiddleName" json:"MiddleName,omitempty"`
	LastName   string `protobuf:"bytes,3,opt,name=LastName" json:"LastName,omitempty"`
	Nickname   string `protobuf:"bytes,4,opt,name=Nickname" json:"Nickname,omitempty"`
	HomePage   string `protobuf:"bytes,5,opt,name=HomePage" json:"HomePage,omitempty"`
	Email      string `protobuf:"bytes,6,opt,name=Email" json:"Email,omitempty"`
}

func (m *AuthorType) Reset()                    { *m = AuthorType{} }
func (m *AuthorType) String() string            { return proto.CompactTextString(m) }
func (*AuthorType) ProtoMessage()               {}
func (*AuthorType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *AuthorType) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *AuthorType) GetMiddleName() string {
	if m != nil {
		return m.MiddleName
	}
	return ""
}

func (m *AuthorType) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *AuthorType) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *AuthorType) GetHomePage() string {
	if m != nil {
		return m.HomePage
	}
	return ""
}

func (m *AuthorType) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func init() {
	proto.RegisterType((*PFB2)(nil), "fb2.PFB2")
	proto.RegisterType((*Description)(nil), "fb2.Description")
	proto.RegisterType((*TitleInfo)(nil), "fb2.TitleInfo")
	proto.RegisterType((*Coverpage)(nil), "fb2.Coverpage")
	proto.RegisterType((*Image)(nil), "fb2.Image")
	proto.RegisterType((*AuthorType)(nil), "fb2.AuthorType")
}

func init() { proto.RegisterFile("fb2.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 428 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x53, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x96, 0xed, 0x24, 0xad, 0xc7, 0xa8, 0xa0, 0x15, 0x87, 0x15, 0x20, 0x64, 0xf9, 0x42, 0x0e,
	0x50, 0x24, 0x73, 0xe4, 0xd4, 0x12, 0x4a, 0x03, 0xa5, 0xaa, 0xdc, 0xbc, 0xc0, 0x26, 0x99, 0x14,
	0xab, 0xf6, 0xae, 0x59, 0x6f, 0x41, 0x79, 0x2f, 0xde, 0x0f, 0x34, 0xb3, 0x8e, 0x6d, 0xd4, 0xdb,
	0x7e, 0x3f, 0x3b, 0xfe, 0x66, 0x66, 0x0d, 0xf1, 0x6e, 0x9d, 0x9f, 0x36, 0xd6, 0x38, 0x23, 0xa2,
	0xdd, 0x3a, 0xcf, 0xbe, 0xc2, 0xe4, 0xe6, 0xe2, 0x3c, 0x17, 0x27, 0x10, 0x2e, 0x17, 0x32, 0x48,
	0x83, 0x79, 0x5c, 0x84, 0xcb, 0x85, 0xc8, 0x21, 0x59, 0x60, 0xbb, 0xb1, 0x65, 0xe3, 0x4a, 0xa3,
	0x65, 0x98, 0x06, 0xf3, 0x24, 0x7f, 0x76, 0x4a, 0xb7, 0x47, 0x7c, 0x31, 0x36, 0x65, 0x1f, 0xff,
	0xbb, 0x23, 0xde, 0x42, 0xbc, 0x2a, 0x5d, 0x85, 0x4b, 0xbd, 0x33, 0x5c, 0x39, 0xc9, 0x4f, 0xb8,
	0x40, 0xcf, 0x16, 0x83, 0x21, 0xfb, 0x1b, 0x8e, 0xec, 0xe2, 0x39, 0x4c, 0xbf, 0xa0, 0xb6, 0x28,
	0x83, 0x34, 0x9a, 0xc7, 0x85, 0x07, 0xe2, 0x15, 0xc4, 0x7c, 0x58, 0xed, 0x1b, 0x94, 0x21, 0x2b,
	0x03, 0x21, 0xde, 0xc0, 0xec, 0xec, 0xc1, 0xfd, 0x30, 0x56, 0x46, 0x69, 0x34, 0x4f, 0xf2, 0xa7,
	0xfc, 0x31, 0x4f, 0x91, 0xa1, 0xe8, 0x64, 0x2a, 0x73, 0x6e, 0xcc, 0x3d, 0x7f, 0x4d, 0x4e, 0xb8,
	0xe5, 0x81, 0x10, 0xaf, 0x01, 0xce, 0xb4, 0x36, 0x4e, 0x71, 0xe3, 0x53, 0x96, 0x47, 0x8c, 0x78,
	0x01, 0xc7, 0xdf, 0x70, 0xff, 0xdb, 0xd8, 0x6d, 0x2b, 0x67, 0xac, 0xf6, 0x58, 0x08, 0x98, 0x2c,
	0x94, 0x43, 0x79, 0xc4, 0x3c, 0x9f, 0x69, 0x0c, 0x9f, 0xcc, 0x2f, 0xb4, 0x8d, 0xba, 0x43, 0x79,
	0x3c, 0x1a, 0x43, 0xcf, 0x16, 0x83, 0x81, 0x2a, 0x5c, 0x29, 0x7d, 0x27, 0x63, 0x5f, 0x81, 0xce,
	0x42, 0xc2, 0xd1, 0xad, 0xdd, 0x30, 0x0d, 0x4c, 0x1f, 0xa0, 0x78, 0x0f, 0xb0, 0xb2, 0x4a, 0xb7,
	0x95, 0x72, 0xc6, 0xca, 0x84, 0x8b, 0x3f, 0x6a, 0x7b, 0x64, 0xa1, 0xf0, 0xb7, 0xf8, 0xf3, 0x01,
	0xf5, 0x06, 0xe5, 0x13, 0x1f, 0xfe, 0x80, 0xb3, 0x77, 0xa3, 0xa0, 0x22, 0x85, 0xe9, 0xb2, 0xa6,
	0xc4, 0x7e, 0x71, 0xc0, 0x45, 0x99, 0x29, 0xbc, 0x90, 0xbd, 0xec, 0x1c, 0x14, 0xf9, 0xd2, 0xe2,
	0xae, 0x7b, 0x3c, 0x7c, 0xce, 0xfe, 0x04, 0x00, 0x43, 0x04, 0x9a, 0xf8, 0x45, 0x69, 0x5b, 0x77,
	0xad, 0x6a, 0xec, 0x7c, 0x03, 0x41, 0x13, 0xff, 0x5e, 0x6e, 0xb7, 0x15, 0xb2, 0x1c, 0xfa, 0x89,
	0x0f, 0x0c, 0x85, 0xbe, 0x52, 0xdd, 0xe5, 0xc8, 0x87, 0x3e, 0x60, 0xd2, 0xae, 0xcb, 0xcd, 0xbd,
	0x26, 0xcd, 0xaf, 0xb2, 0xc7, 0xa4, 0x5d, 0x9a, 0x1a, 0x6f, 0xa8, 0x0d, 0xbf, 0xc7, 0x1e, 0xd3,
	0x03, 0xfb, 0x5c, 0xab, 0xb2, 0xea, 0x56, 0xe8, 0xc1, 0x7a, 0xc6, 0x7f, 0xc6, 0x87, 0x7f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x02, 0x34, 0xa1, 0xd2, 0x26, 0x03, 0x00, 0x00,
}