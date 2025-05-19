/*

Package unioffice provides creation, reading, and writing of ECMA 376 Office Open
XML documents, spreadsheets and presentations.  It is still early in
development, but is progressing quickly.  This library takes a slightly
different approach from others, in that it starts by trying to support all of
the ECMA-376 standard when marshaling/unmarshaling XML documents.  From there it
adds wrappers around the ECMA-376 derived types that provide a more convenient
interface.

The raw XML based types reside in the `schema/`` directory. These types are
always accessible from the wrapper types via a `X() method that returns the
raw type.  Except for the base documents (document.Document,
spreadsheet.Workbook and presentation.Presentation), the other wrapper types are
value types with non-pointer methods.  They exist solely to modify and return
data from one or more XML types.

The packages of interest are gitee.com/greatmusicians/unioffice/document,
unidoc/unioffice/spreadsheet and gitee.com/greatmusicians/unioffice/presentation.

*/
package unioffice

import (
	_cd "encoding/xml"
	_fg "errors"
	_cb "fmt"
	_d "log"
	_f "reflect"
	_cg "strings"
	_g "unicode"

	_ca "gitee.com/greatmusicians/unioffice/algo"
)

// XSDAny  is used to marshal/unmarshal xsd:any types in the OOXML schema.
type XSDAny struct {
	XMLName _cd.Name
	Attrs   []_cd.Attr
	Data    []byte
	Nodes   []*XSDAny
}

// Uint8 returns a copy of v as a pointer.
func Uint8(v uint8) *uint8 { _bfd := v; return &_bfd }

// Stringf formats according to a format specifier and returns a pointer to the
// resulting string.
func Stringf(f string, args ...interface{}) *string { _fba := _cb.Sprintf(f, args...); return &_fba }

// String returns a copy of v as a pointer.
func String(v string) *string { _beb := v; return &_beb }

// AddPreserveSpaceAttr adds an xml:space="preserve" attribute to a start
// element if it is required for the string s.
func AddPreserveSpaceAttr(se *_cd.StartElement, s string) {
	if NeedsSpacePreserve(s) {
		se.Attr = append(se.Attr, _cd.Attr{Name: _cd.Name{Local: "xml:space"}, Value: "preserve"})
	}
}

// Int8 returns a copy of v as a pointer.
func Int8(v int8) *int8 { _bfdd := v; return &_bfdd }

const (
	OfficeDocumentTypeStrict     = "http:\u002f\u002fpurl\u002eoclc\u002eorg\u002fooxml\u002fofficeDocument\u002frelationships\u002fofficeDocument"
	StylesTypeStrict             = "http:\u002f\u002fpurl\u002eoclc\u002eorg\u002fooxml\u002fofficeDocument\u002frelationships\u002fstyles"
	ThemeTypeStrict              = "http:\u002f\u002fpurl\u002eoclc\u002eorg\u002fooxml\u002fofficeDocument\u002frelationships\u002ftheme"
	ControlTypeStrict            = "http:\u002f\u002fpurl\u002eoclc\u002eorg\u002fooxml\u002fofficeDocument\u002frelationships\u002fcontrol"
	SettingsTypeStrict           = "http:\u002f\u002fpurl\u002eoclc\u002eorg\u002fooxml\u002fofficeDocument\u002frelationships\u002fsettings"
	ImageTypeStrict              = "http:\u002f\u002fpurl\u002eoclc\u002eorg\u002fooxml\u002fofficeDocument\u002frelationships\u002fimage"
	CommentsTypeStrict           = "http:\u002f\u002fpurl\u002eoclc\u002eorg\u002fooxml\u002fofficeDocument\u002frelationships\u002fcomments"
	ThumbnailTypeStrict          = "http:/\u002fpurl\u002eoclc\u002eorg/ooxml\u002fofficeDocument\u002frelationships\u002fmetadata\u002fthumbnail"
	DrawingTypeStrict            = "http:\u002f\u002fpurl\u002eoclc\u002eorg\u002fooxml\u002fofficeDocument\u002frelationships\u002fdrawing"
	ChartTypeStrict              = "http:\u002f\u002fpurl\u002eoclc\u002eorg\u002fooxml\u002fofficeDocument\u002frelationships\u002fchart"
	ExtendedPropertiesTypeStrict = "http:/\u002fpurl\u002eoclc\u002eorg/ooxml\u002fofficeDocument\u002frelationships\u002fextendedProperties"
	CustomXMLTypeStrict          = "http:/\u002fpurl.oclc\u002eorg\u002fooxml\u002fofficeDocument\u002frelationships\u002fcustomXml"
	WorksheetTypeStrict          = "http:/\u002fpurl.oclc\u002eorg\u002fooxml\u002fofficeDocument\u002frelationships\u002fworksheet"
	SharedStringsTypeStrict      = "http:\u002f\u002fpurl\u002eoclc\u002eorg\u002fooxml\u002fofficeDocument/relationships/sharedStrings"
	SharedStingsTypeStrict       = SharedStringsTypeStrict
	TableTypeStrict              = "http:\u002f\u002fpurl\u002eoclc\u002eorg\u002fooxml\u002fofficeDocument\u002frelationships\u002ftable"
	HeaderTypeStrict             = "http:\u002f\u002fpurl\u002eoclc\u002eorg\u002fooxml\u002fofficeDocument\u002frelationships\u002fheader"
	FooterTypeStrict             = "http:\u002f\u002fpurl\u002eoclc\u002eorg\u002fooxml\u002fofficeDocument\u002frelationships\u002ffooter"
	NumberingTypeStrict          = "http:/\u002fpurl.oclc\u002eorg\u002fooxml\u002fofficeDocument\u002frelationships\u002fnumbering"
	FontTableTypeStrict          = "http:/\u002fpurl.oclc\u002eorg\u002fooxml\u002fofficeDocument\u002frelationships\u002ffontTable"
	WebSettingsTypeStrict        = "http:\u002f/purl\u002eoclc\u002eorg\u002fooxml\u002fofficeDocument\u002frelationships\u002fwebSettings"
	FootNotesTypeStrict          = "http:/\u002fpurl.oclc\u002eorg\u002fooxml\u002fofficeDocument\u002frelationships\u002ffootnotes"
	EndNotesTypeStrict           = "http:\u002f\u002fpurl\u002eoclc\u002eorg\u002fooxml\u002fofficeDocument\u002frelationships\u002fendnotes"
	SlideTypeStrict              = "http:\u002f\u002fpurl\u002eoclc\u002eorg\u002fooxml\u002fofficeDocument\u002frelationships\u002fslide"
	VMLDrawingTypeStrict         = "http:\u002f\u002fpurl\u002eoclc\u002eorg\u002fooxml\u002fofficeDocument\u002frelationships\u002fvmlDrawing"
	OfficeDocumentType           = "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg\u002fofficeDocument\u002f2006\u002frelationships\u002fofficeDocument"
	StylesType                   = "http:/\u002fschemas\u002eopenxmlformats.org\u002fofficeDocument\u002f2006\u002frelationships\u002fstyles"
	ThemeType                    = "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg\u002fofficeDocument\u002f2006\u002frelationships/theme"
	ThemeContentType             = "application/vnd.openxmlformats\u002dofficedocument\u002etheme\u002bxml"
	SettingsType                 = "http:/\u002fschemas\u002eopenxmlformats\u002eorg\u002fofficeDocument\u002f2006\u002frelationships/settings"
	ImageType                    = "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg\u002fofficeDocument\u002f2006\u002frelationships/image"
	ControlType                  = "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg\u002fofficeDocument\u002f2006\u002frelationships\u002fcontrol"
	CommentsType                 = "http:/\u002fschemas\u002eopenxmlformats\u002eorg\u002fofficeDocument\u002f2006\u002frelationships/comments"
	CommentsContentType          = "application\u002fvnd\u002eopenxmlformats\u002dofficedocument.spreadsheetml\u002ecomments\u002bxml"
	ThumbnailType                = "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg\u002fpackage\u002f2006\u002frelationships\u002fmetadata\u002fthumbnail"
	DrawingType                  = "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg\u002fofficeDocument\u002f2006\u002frelationships\u002fdrawing"
	DrawingContentType           = "application\u002fvnd\u002eopenxmlformats\u002dofficedocument\u002edrawing\u002bxml"
	ChartType                    = "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg\u002fofficeDocument\u002f2006\u002frelationships/chart"
	ChartContentType             = "application/vnd\u002eopenxmlformats\u002dofficedocument\u002edrawingml\u002echart\u002bxml"
	HyperLinkType                = "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg\u002fofficeDocument\u002f2006\u002frelationships\u002fhyperlink"
	ExtendedPropertiesType       = "http:\u002f\u002fschemas\u002eopenxmlformats.org\u002fofficeDocument\u002f2006\u002frelationships\u002fextended\u002dproperties"
	CorePropertiesType           = "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg\u002fpackage\u002f2006\u002frelationships\u002fmetadata/core\u002dproperties"
	CustomPropertiesType         = "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg\u002fofficeDocument\u002f2006\u002frelationships\u002fcustom\u002dproperties"
	CustomXMLType                = "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg\u002fofficeDocument\u002f2006\u002frelationships\u002fcustomXml"
	TableStylesType              = "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg\u002fofficeDocument\u002f2006\u002frelationships\u002ftableStyles"
	ViewPropertiesType           = "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg\u002fofficeDocument\u002f2006\u002frelationships\u002fviewProps"
	WorksheetType                = "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg\u002fofficeDocument\u002f2006\u002frelationships\u002fworksheet"
	WorksheetContentType         = "application\u002fvnd.openxmlformats\u002dofficedocument\u002espreadsheetml\u002eworksheet\u002bxml"
	SharedStringsType            = "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg\u002fofficeDocument/2006/relationships\u002fsharedStrings"
	SharedStingsType             = SharedStringsType
	SharedStringsContentType     = "application\u002fvnd\u002eopenxmlformats\u002dofficedocument\u002espreadsheetml\u002esharedStrings\u002bxml"
	SMLStyleSheetContentType     = "application\u002fvnd\u002eopenxmlformats\u002dofficedocument\u002espreadsheetml\u002estyles\u002bxml"
	TableType                    = "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg\u002fofficeDocument\u002f2006\u002frelationships/table"
	TableContentType             = "application\u002fvnd\u002eopenxmlformats\u002dofficedocument\u002espreadsheetml\u002etable\u002bxml"
	HeaderType                   = "http:/\u002fschemas\u002eopenxmlformats.org\u002fofficeDocument\u002f2006\u002frelationships\u002fheader"
	FooterType                   = "http:/\u002fschemas\u002eopenxmlformats.org\u002fofficeDocument\u002f2006\u002frelationships\u002ffooter"
	NumberingType                = "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg\u002fofficeDocument\u002f2006\u002frelationships\u002fnumbering"
	FontTableType                = "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg\u002fofficeDocument\u002f2006\u002frelationships\u002ffontTable"
	WebSettingsType              = "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg\u002fofficeDocument\u002f2006\u002frelationships\u002fwebSettings"
	FootNotesType                = "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg\u002fofficeDocument\u002f2006\u002frelationships\u002ffootnotes"
	EndNotesType                 = "http:/\u002fschemas\u002eopenxmlformats\u002eorg\u002fofficeDocument\u002f2006\u002frelationships/endnotes"
	SlideType                    = "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg\u002fofficeDocument\u002f2006\u002frelationships/slide"
	SlideContentType             = "application\u002fvnd\u002eopenxmlformats\u002dofficedocument\u002epresentationml\u002eslide\u002bxml"
	SlideMasterType              = "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg\u002fofficeDocument\u002f2006\u002frelationships\u002fslideMaster"
	SlideMasterContentType       = "application\u002fvnd\u002eopenxmlformats\u002dofficedocument\u002epresentationml\u002eslideMaster\u002bxml"
	SlideLayoutType              = "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg\u002fofficeDocument\u002f2006\u002frelationships\u002fslideLayout"
	SlideLayoutContentType       = "application\u002fvnd\u002eopenxmlformats\u002dofficedocument\u002epresentationml\u002eslideLayout\u002bxml"
	PresentationPropertiesType   = "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg\u002fofficeDocument\u002f2006\u002frelationships\u002fpresProps"
	HandoutMasterType            = "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg\u002fofficeDocument/2006/relationships\u002fhandoutMaster"
	NotesMasterType              = "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg\u002fofficeDocument\u002f2006\u002frelationships\u002fnotesMaster"
	VMLDrawingType               = "http:\u002f\u002fschemas\u002eopenxmlformats.org\u002fofficeDocument\u002f2006\u002frelationships\u002fvmlDrawing"
	VMLDrawingContentType        = "application\u002fvnd\u002eopenxmlformats\u002dofficedocument\u002evmlDrawing"
)

// Float32 returns a copy of v as a pointer.
func Float32(v float32) *float32 { _ecc := v; return &_ecc }

// Bool returns a copy of v as a pointer.
func Bool(v bool) *bool { _cda := v; return &_cda }

// AbsoluteImageFilename returns the full path to an image from the root of the
// zip container.
func AbsoluteImageFilename(dt DocType, index int, fileExtension string) string {
	_cf := AbsoluteFilename(dt, ImageType, index)
	return _cf[0:len(_cf)-3] + fileExtension
}

// Int64 returns a copy of v as a pointer.
func Int64(v int64) *int64 { _da := v; return &_da }

// Uint16 returns a copy of v as a pointer.
func Uint16(v uint16) *uint16 { _afa := v; return &_afa }
func (_daa nsSet) applyToNode(_bc *any) {
	if _bc.XMLName.Space == "" {
		return
	}
	_ded := _daa.getPrefix(_bc.XMLName.Space)
	_bc.XMLName.Space = ""
	_bc.XMLName.Local = _ded + ":" + _bc.XMLName.Local
	_bbg := _bc.Attrs
	_bc.Attrs = nil
	for _, _eec := range _bbg {
		if _eec.Name.Space == "xmlns" {
			continue
		}
		if _eec.Name.Space != "" {
			_abc := _daa.getPrefix(_eec.Name.Space)
			_eec.Name.Space = ""
			_eec.Name.Local = _abc + ":" + _eec.Name.Local
		}
		_bc.Attrs = append(_bc.Attrs, _eec)
	}
	for _, _gcd := range _bc.Nodes {
		_daa.applyToNode(_gcd)
	}
}

// Uint64 returns a copy of v as a pointer.
func Uint64(v uint64) *uint64 { _dcd := v; return &_dcd }

// DocType represents one of the three document types supported (docx/xlsx/pptx)
type DocType byte

func _bdb(_cgd []*XSDAny) []*any {
	_gaa := []*any{}
	for _, _efd := range _cgd {
		_ddf := &any{}
		_ddf.XMLName = _efd.XMLName
		_dde := []_cd.Attr{}
		for _, _feg := range _efd.Attrs {
			if _feg.Name.Local != "xmlns" {
				_dde = append(_dde, _feg)
			}
		}
		_ddf.Attrs = _dde
		_ddf.Data = _efd.Data
		_ddf.Nodes = _bdb(_efd.Nodes)
		_gaa = append(_gaa, _ddf)
	}
	return _gaa
}

// Float64 returns a copy of v as a pointer.
func Float64(v float64) *float64 { _fbd := v; return &_fbd }

// Int32 returns a copy of v as a pointer.
func Int32(v int32) *int32 { _bbc := v; return &_bbc }

var _geb = map[string]bool{"w10": true, "w14": true, "wp14": true, "w15": true, "x15ac": true, "w16se": true, "w16cid": true, "w16": true, "w16cex": true}

func _gfb(_dad *any) {
	for _, _cfa := range _dad.Nodes {
		_gfb(_cfa)
	}
}

// CreateElement creates an element with the given namespace and name. It is
// used to unmarshal some xsd:any elements to the appropriate concrete type.
func CreateElement(start _cd.StartElement) (Any, error) {
	_ab, _dc := _ec[start.Name.Space+"\u002f"+start.Name.Local]
	if !_dc {
		_fge := &XSDAny{}
		return _fge, nil
	}
	_de := _f.ValueOf(_ab)
	_fb := _de.Call(nil)
	if len(_fb) != 1 {
		return nil, _cb.Errorf("constructor\u0020function\u0020should\u0020return\u0020one\u0020value\u002c\u0020got\u0020\u0025d", len(_fb))
	}
	_bf, _dc := _fb[0].Interface().(Any)
	if !_dc {
		return nil, _fg.New("constructor\u0020function\u0020should\u0020return\u0020any \u0027Any\u0027")
	}
	return _bf, nil
}

// Uint32 returns a copy of v as a pointer.
func Uint32(v uint32) *uint32 { _be := v; return &_be }

// UnmarshalXML implements the xml.Unmarshaler interface.
func (_ga *XSDAny) UnmarshalXML(d *_cd.Decoder, start _cd.StartElement) error {
	_fgd := any{}
	if _ee := d.DecodeElement(&_fgd, &start); _ee != nil {
		return _ee
	}
	_gfb(&_fgd)
	_ga.XMLName = _fgd.XMLName
	_ga.Attrs = _fgd.Attrs
	_ga.Data = _fgd.Data
	_ga.Nodes = _fa(_fgd.Nodes)
	return nil
}

// Any is the interface used for marshaling/unmarshaling xsd:any
type Any interface {
	MarshalXML(_a *_cd.Encoder, _db _cd.StartElement) error
	UnmarshalXML(_e *_cd.Decoder, _b _cd.StartElement) error
}

// NeedsSpacePreserve returns true if the string has leading or trailing space.
func NeedsSpacePreserve(s string) bool {
	if len(s) == 0 {
		return false
	}
	switch s[0] {
	case '\t', '\n', '\v', '\f', '\r', ' ', 0x85, 0xA0:
		return true
	}
	switch s[len(s)-1] {
	case '\t', '\n', '\v', '\f', '\r', ' ', 0x85, 0xA0:
		return true
	}
	return false
}

const (
	Unknown DocType = iota
	DocTypeSpreadsheet
	DocTypeDocument
	DocTypePresentation
)

type any struct {
	XMLName _cd.Name
	Attrs   []_cd.Attr `xml:",any,attr"`
	Nodes   []*any     `xml:",any"`
	Data    []byte     `xml:",chardata"`
}

func (_dae *XSDAny) collectNS(_ad *nsSet) {
	if _dae.XMLName.Space != "" {
		_ad.getPrefix(_dae.XMLName.Space)
	}
	for _, _aed := range _dae.Attrs {
		if _aed.Name.Space != "" && _aed.Name.Space != "xmlns" {
			_ad.getPrefix(_aed.Name.Space)
		}
	}
	for _, _fgec := range _dae.Nodes {
		_fgec.collectNS(_ad)
	}
}

// AbsoluteFilename returns the full path to a file from the root of the zip
// container. Index is used in some cases for files which there may be more than
// one of (e.g. worksheets/drawings/charts)
func AbsoluteFilename(dt DocType, typ string, index int) string {
	switch typ {
	case CorePropertiesType:
		return "docProps\u002fcore\u002exml"
	case CustomPropertiesType:
		return "docProps\u002fcustom\u002exml"
	case ExtendedPropertiesType, ExtendedPropertiesTypeStrict:
		return "docProps\u002fapp\u002exml"
	case ThumbnailType, ThumbnailTypeStrict:
		return "docProps\u002fthumbnail\u002ejpeg"
	case CustomXMLType:
		return _cb.Sprintf("customXml\u002fitem\u0025d.xml", index)
	case PresentationPropertiesType:
		return "ppt\u002fpresProps\u002exml"
	case ViewPropertiesType:
		switch dt {
		case DocTypePresentation:
			return "ppt\u002fviewProps\u002exml"
		case DocTypeSpreadsheet:
			return "xl/viewProps\u002exml"
		case DocTypeDocument:
			return "word\u002fviewProps\u002exml"
		}
	case TableStylesType:
		switch dt {
		case DocTypePresentation:
			return "ppt\u002ftableStyles\u002exml"
		case DocTypeSpreadsheet:
			return "xl\u002ftableStyles\u002exml"
		case DocTypeDocument:
			return "word\u002ftableStyles.xml"
		}
	case HyperLinkType:
		return ""
	case OfficeDocumentType, OfficeDocumentTypeStrict:
		switch dt {
		case DocTypeSpreadsheet:
			return "xl\u002fworkbook\u002exml"
		case DocTypeDocument:
			return "word\u002fdocument\u002exml"
		case DocTypePresentation:
			return "ppt\u002fpresentation.xml"
		default:
			Log("unsupported\u0020type \u0025s\u0020pair\u0020and\u0020\u0025v", typ, dt)
		}
	case ThemeType, ThemeTypeStrict, ThemeContentType:
		switch dt {
		case DocTypeSpreadsheet:
			return _cb.Sprintf("xl/theme\u002ftheme\u0025d.xml", index)
		case DocTypeDocument:
			return _cb.Sprintf("word/theme\u002ftheme\u0025d\u002exml", index)
		case DocTypePresentation:
			return _cb.Sprintf("ppt\u002ftheme\u002ftheme\u0025d\u002exml", index)
		default:
			Log("unsupported\u0020type \u0025s\u0020pair\u0020and\u0020\u0025v", typ, dt)
		}
	case StylesType, StylesTypeStrict:
		switch dt {
		case DocTypeSpreadsheet:
			return "xl\u002fstyles\u002exml"
		case DocTypeDocument:
			return "word\u002fstyles\u002exml"
		case DocTypePresentation:
			return "ppt\u002fstyles\u002exml"
		default:
			Log("unsupported\u0020type \u0025s\u0020pair\u0020and\u0020\u0025v", typ, dt)
		}
	case ChartType, ChartTypeStrict, ChartContentType:
		switch dt {
		case DocTypeSpreadsheet:
			return _cb.Sprintf("xl\u002fcharts\u002fchart\u0025d\u002exml", index)
		case DocTypeDocument:
			return _cb.Sprintf("word/charts\u002fchart\u0025d\u002exml", index)
		case DocTypePresentation:
			return _cb.Sprintf("ppt\u002fcharts\u002fchart\u0025d\u002exml", index)
		default:
			Log("unsupported\u0020type \u0025s\u0020pair\u0020and\u0020\u0025v", typ, dt)
		}
	case TableType, TableTypeStrict, TableContentType:
		return _cb.Sprintf("xl\u002ftables\u002ftable\u0025d\u002exml", index)
	case DrawingType, DrawingTypeStrict, DrawingContentType:
		switch dt {
		case DocTypeSpreadsheet:
			return _cb.Sprintf("xl\u002fdrawings\u002fdrawing\u0025d\u002exml", index)
		default:
			Log("unsupported\u0020type \u0025s\u0020pair\u0020and\u0020\u0025v", typ, dt)
		}
	case CommentsType, CommentsTypeStrict, CommentsContentType:
		switch dt {
		case DocTypeSpreadsheet:
			return _cb.Sprintf("xl\u002fcomments\u0025d\u002exml", index)
		default:
			Log("unsupported\u0020type \u0025s\u0020pair\u0020and\u0020\u0025v", typ, dt)
		}
	case VMLDrawingType, VMLDrawingTypeStrict, VMLDrawingContentType:
		switch dt {
		case DocTypeSpreadsheet:
			return _cb.Sprintf("xl\u002fdrawings\u002fvmlDrawing\u0025d\u002evml", index)
		default:
			Log("unsupported\u0020type \u0025s\u0020pair\u0020and\u0020\u0025v", typ, dt)
		}
	case ImageType, ImageTypeStrict:
		switch dt {
		case DocTypeDocument:
			return _cb.Sprintf("word/media\u002fimage\u0025d\u002epng", index)
		case DocTypeSpreadsheet:
			return _cb.Sprintf("xl/media\u002fimage\u0025d.png", index)
		case DocTypePresentation:
			return _cb.Sprintf("ppt\u002fmedia\u002fimage\u0025d\u002epng", index)
		default:
			Log("unsupported\u0020type \u0025s\u0020pair\u0020and\u0020\u0025v", typ, dt)
		}
	case WorksheetType, WorksheetTypeStrict, WorksheetContentType:
		return _cb.Sprintf("xl\u002fworksheets\u002fsheet\u0025d\u002exml", index)
	case SharedStringsType, SharedStringsTypeStrict, SharedStringsContentType:
		return "xl/sharedStrings.xml"
	case FontTableType, FontTableTypeStrict:
		return "word\u002ffontTable\u002exml"
	case EndNotesType, EndNotesTypeStrict:
		return "word\u002fendnotes\u002exml"
	case FootNotesType, FootNotesTypeStrict:
		return "word\u002ffootnotes\u002exml"
	case NumberingType, NumberingTypeStrict:
		return "word\u002fnumbering\u002exml"
	case WebSettingsType, WebSettingsTypeStrict:
		return "word\u002fwebSettings.xml"
	case SettingsType, SettingsTypeStrict:
		return "word\u002fsettings\u002exml"
	case HeaderType, HeaderTypeStrict:
		return _cb.Sprintf("word\u002fheader\u0025d\u002exml", index)
	case FooterType, FooterTypeStrict:
		return _cb.Sprintf("word\u002ffooter\u0025d\u002exml", index)
	case ControlType, ControlTypeStrict:
		switch dt {
		case DocTypeSpreadsheet:
			return _cb.Sprintf("xl\u002factiveX\u002factiveX\u0025d\u002exml", index)
		case DocTypeDocument:
			return _cb.Sprintf("word\u002factiveX\u002factiveX\u0025d.xml", index)
		case DocTypePresentation:
			return _cb.Sprintf("ppt\u002factiveX\u002factiveX\u0025d\u002exml", index)
		default:
			Log("unsupported\u0020type \u0025s\u0020pair\u0020and\u0020\u0025v", typ, dt)
		}
	case SlideType, SlideTypeStrict:
		return _cb.Sprintf("ppt\u002fslides\u002fslide\u0025d\u002exml", index)
	case SlideLayoutType:
		return _cb.Sprintf("ppt/slideLayouts/slideLayout\u0025d\u002exml", index)
	case SlideMasterType:
		return _cb.Sprintf("ppt/slideMasters/slideMaster\u0025d\u002exml", index)
	case HandoutMasterType:
		return _cb.Sprintf("ppt\u002fhandoutMasters\u002fhandoutMaster\u0025d\u002exml", index)
	case NotesMasterType:
		return _cb.Sprintf("ppt/notesMasters/notesMaster\u0025d\u002exml", index)
	default:
		Log("unsupported\u0020type\u0020\u0025s", typ)
	}
	return ""
}

// DisableLogging sets the Log function to a no-op so that any log messages are
// silently discarded.
func DisableLogging() { Log = func(string, ...interface{}) {} }

// RegisterConstructor registers a constructor function used for unmarshaling
// xsd:any elements.
func RegisterConstructor(ns, name string, fn interface{}) { _ec[ns+"\u002f"+name] = fn }

const (
	ContentTypesFilename = "\u005bContent_Types\u005d\u002exml"
	BaseRelsFilename     = "_rels\u002f\u002erels"
)

type nsSet struct {
	_cbb map[string]string
	_bg  map[string]string
	_cde []string
}

var _bbb = func() map[string]string {
	_beg := map[string]string{}
	for _gda, _fe := range _gfg {
		_beg[_fe] = _gda
	}
	return _beg
}()

// RelativeImageFilename returns an image filename relative to the source file referenced
// from a relationships file. It is identical to RelativeFilename but is used particularly for images
// in order to handle different image formats.
func RelativeImageFilename(dt DocType, relToTyp, typ string, index int, fileExtension string) string {
	_gd := RelativeFilename(dt, relToTyp, typ, index)
	return _gd[0:len(_gd)-3] + fileExtension
}
func _fa(_gg []*any) []*XSDAny {
	_ff := []*XSDAny{}
	for _, _bec := range _gg {
		_fee := &XSDAny{}
		_fee.XMLName = _bec.XMLName
		_fee.Attrs = _bec.Attrs
		_fee.Data = _bec.Data
		_fee.Nodes = _fa(_bec.Nodes)
		_ff = append(_ff, _fee)
	}
	return _ff
}

var _gfg = map[string]string{"a": "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg\u002fdrawingml\u002f2006\u002fmain", "dc": "http:\u002f\u002fpurl\u002eorg/dc\u002felements\u002f1\u002e1\u002f", "dcterms": "http:\u002f/purl\u002eorg\u002fdc\u002fterms/", "mc": "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg\u002fmarkup\u002dcompatibility\u002f2006", "mo": "http:\u002f/schemas.microsoft\u002ecom\u002foffice\u002fmac\u002foffice\u002f2008\u002fmain", "w": "http:\u002f\u002fschemas.openxmlformats\u002eorg\u002fwordprocessingml\u002f2006\u002fmain", "w10": "urn:schemas\u002dmicrosoft\u002dcom:office:word", "w14": "http:\u002f\u002fschemas.microsoft\u002ecom\u002foffice\u002fword\u002f2010\u002fwordml", "w15": "http:\u002f\u002fschemas.microsoft\u002ecom\u002foffice\u002fword\u002f2012\u002fwordml", "wne": "http:\u002f\u002fschemas.microsoft\u002ecom\u002foffice\u002fword\u002f2006\u002fwordml", "wp": "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg\u002fdrawingml\u002f2006\u002fwordprocessingDrawing", "wp14": "http:\u002f/schemas\u002emicrosoft\u002ecom\u002foffice\u002fword\u002f2010\u002fwordprocessingDrawing", "wpc": "http:\u002f\u002fschemas\u002emicrosoft\u002ecom\u002foffice\u002fword\u002f2010\u002fwordprocessingCanvas", "wpg": "http:/\u002fschemas\u002emicrosoft\u002ecom\u002foffice\u002fword\u002f2010\u002fwordprocessingGroup", "wpi": "http:\u002f\u002fschemas\u002emicrosoft\u002ecom/office\u002fword\u002f2010\u002fwordprocessingInk", "wps": "http:/\u002fschemas\u002emicrosoft\u002ecom\u002foffice\u002fword\u002f2010\u002fwordprocessingShape", "xsi": "http:/\u002fwww\u002ew3\u002eorg\u002f2001\u002fXMLSchema\u002dinstance", "x15ac": "http:\u002f\u002fschemas.microsoft\u002ecom\u002foffice\u002fspreadsheetml\u002f2010/11\u002fac", "w16se": "http:\u002f\u002fschemas\u002emicrosoft\u002ecom\u002foffice\u002fword\u002f2015\u002fwordml\u002fsymex", "w16cid": "http:\u002f/schemas.microsoft\u002ecom\u002foffice\u002fword\u002f2016\u002fwordml\u002fcid", "w16": "http:\u002f\u002fschemas.microsoft\u002ecom\u002foffice\u002fword\u002f2018\u002fwordml", "w16cex": "http:\u002f/schemas.microsoft\u002ecom\u002foffice\u002fword\u002f2018\u002fwordml\u002fcex", "xml": "http:\u002f\u002fwww\u002ew3.org/XML\u002f1998/namespace"}

func (_ag *nsSet) getPrefix(_dga string) string {
	if _dd, _gc := _bbb[_dga]; _gc {
		if _, _ecb := _ag._bg[_dd]; !_ecb {
			_ag._bg[_dd] = _dga
			_ag._cbb[_dga] = _dd
			_ag._cde = append(_ag._cde, _dd)
		}
		return _dd
	}
	_dga = _cg.TrimFunc(_dga, func(_agf rune) bool { return !_g.IsLetter(_agf) })
	if _ccd, _fbg := _ag._cbb[_dga]; _fbg {
		return _ccd
	}
	_ef := _cg.Split(_dga, "\u002f")
	_ef = _cg.Split(_ef[len(_ef)-1], ":")
	_aa := _ef[len(_ef)-1]
	_aae := 0
	_gdd := []byte{}
	for {
		if _aae < len(_aa) {
			_gdd = append(_gdd, _aa[_aae])
		} else {
			_gdd = append(_gdd, '_')
		}
		_aae++
		if _, _gb := _ag._bg[string(_gdd)]; !_gb {
			_ag._bg[string(_gdd)] = _dga
			_ag._cbb[_dga] = string(_gdd)
			_ag._cde = append(_ag._cde, string(_gdd))
			return string(_gdd)
		}
	}
}

// MarshalXML implements the xml.Marshaler interface.
func (_abaf *XSDAny) MarshalXML(e *_cd.Encoder, start _cd.StartElement) error {
	start.Name = _abaf.XMLName
	start.Attr = _abaf.Attrs
	_eg := any{}
	_eg.XMLName = _abaf.XMLName
	_eg.Attrs = _abaf.Attrs
	_eg.Data = _abaf.Data
	_eg.Nodes = _bdb(_abaf.Nodes)
	_efdg := []string{}
	_abd := false
	_dbb := nsSet{_cbb: map[string]string{}, _bg: map[string]string{}}
	_abaf.collectNS(&_dbb)
	_dbb.applyToNode(&_eg)
	for _, _egc := range _dbb._cde {
		if _, _fd := _geb[_egc]; _fd {
			_efdg = append(_efdg, _egc)
		}
		_cce := _dbb._bg[_egc]
		_eg.Attrs = append(_eg.Attrs, _cd.Attr{Name: _cd.Name{Local: "xmlns:" + _egc}, Value: _cce})
		if _egc == "mc" {
			_abd = true
		}
	}
	if _abd && len(_efdg) > 0 {
		_eg.Attrs = append(_eg.Attrs, _cd.Attr{Name: _cd.Name{Local: "mc:Ignorable"}, Value: _cg.Join(_efdg, "\u0020")})
	}
	return e.Encode(&_eg)
}

const MinGoVersion = _dg
const _dg = true

var _ec = map[string]interface{}{}

// RelativeFilename returns a filename relative to the source file referenced
// from a relationships file. Index is used in some cases for files which there
// may be more than one of (e.g. worksheets/drawings/charts)
func RelativeFilename(dt DocType, relToTyp, typ string, index int) string {
	_bd := AbsoluteFilename(dt, typ, index)
	if relToTyp == "" {
		return _bd
	}
	_cc := AbsoluteFilename(dt, relToTyp, index)
	_bb := _cg.Split(_cc, "\u002f")
	_af := _cg.Split(_bd, "\u002f")
	_gf := 0
	for _aba := 0; _aba < len(_bb); _aba++ {
		if _bb[_aba] == _af[_aba] {
			_gf++
		}
		if _aba+1 == len(_af) {
			break
		}
	}
	_bb = _bb[_gf:]
	_af = _af[_gf:]
	_ge := len(_bb) - 1
	if _ge > 0 {
		return _ca.RepeatString("\u002e\u002e\u002f", _ge) + _cg.Join(_af, "\u002f")
	}
	return _cg.Join(_af, "\u002f")
}

var Log = _d.Printf
