//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

package custom_properties

import (
	_d "encoding/xml"
	_da "fmt"
	_ef "gitee.com/greatmusicians/unioffice"
	_b "gitee.com/greatmusicians/unioffice/schema/soo/ofc/docPropsVTypes"
	_bc "gitee.com/greatmusicians/unioffice/schema/soo/ofc/sharedTypes"
	_ff "strconv"
	_e "time"
)

// ValidateWithPath validates the CT_Properties and its children, prefixing error messages with path
func (_ge *CT_Properties) ValidateWithPath(path string) error {
	for _bf, _gd := range _ge.Property {
		if _bd := _gd.ValidateWithPath(_da.Sprintf("\u0025s\u002fProperty\u005b\u0025d\u005d", path, _bf)); _bd != nil {
			return _bd
		}
	}
	return nil
}
func NewCT_Properties() *CT_Properties { _g := &CT_Properties{}; return _g }
func (_de *CT_Properties) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	e.EncodeToken(start)
	if _de.Property != nil {
		_fa := _d.StartElement{Name: _d.Name{Local: "property"}}
		for _, _ee := range _de.Property {
			e.EncodeElement(_ee, _fa)
		}
	}
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}

// Validate validates the CT_Properties and its children
func (_cc *CT_Properties) Validate() error { return _cc.ValidateWithPath("CT_Properties") }

type Properties struct{ CT_Properties }

// Validate validates the CT_Property and its children
func (_bcg *CT_Property) Validate() error { return _bcg.ValidateWithPath("CT_Property") }
func NewCT_Property() *CT_Property {
	_fd := &CT_Property{}
	_fd.FmtidAttr = "\u007b00000000\u002d0000\u002d0000\u002d0000-000000000000\u007d"
	return _fd
}
func (_agg *CT_Property) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_agg.FmtidAttr = "\u007b00000000\u002d0000\u002d0000\u002d0000-000000000000\u007d"
	for _, _ccec := range start.Attr {
		if _ccec.Name.Local == "pid" {
			_dg, _dfa := _ff.ParseInt(_ccec.Value, 10, 32)
			if _dfa != nil {
				return _dfa
			}
			_agg.PidAttr = int32(_dg)
			continue
		}
		if _ccec.Name.Local == "linkTarget" {
			_ea, _aae := _ccec.Value, error(nil)
			if _aae != nil {
				return _aae
			}
			_agg.LinkTargetAttr = &_ea
			continue
		}
		if _ccec.Name.Local == "name" {
			_ecc, _bfa := _ccec.Value, error(nil)
			if _bfa != nil {
				return _bfa
			}
			_agg.NameAttr = &_ecc
			continue
		}
		if _ccec.Name.Local == "fmtid" {
			_ace, _ga := _ccec.Value, error(nil)
			if _ga != nil {
				return _ga
			}
			_agg.FmtidAttr = _ace
			continue
		}
	}
_bbg:
	for {
		_geb, _gceg := d.Token()
		if _gceg != nil {
			return _gceg
		}
		switch _cbf := _geb.(type) {
		case _d.StartElement:
			switch _cbf.Name {
			case _d.Name{Space: "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg\u002fofficeDocument\u002f2006\u002fdocPropsVTypes", Local: "vector"}, _d.Name{Space: "http:\u002f/purl\u002eoclc\u002eorg\u002fooxml\u002fofficeDocument\u002fdocPropsVTypes", Local: "vector"}:
				_agg.Vector = _b.NewVector()
				if _acb := d.DecodeElement(_agg.Vector, &_cbf); _acb != nil {
					return _acb
				}
			case _d.Name{Space: "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg\u002fofficeDocument\u002f2006\u002fdocPropsVTypes", Local: "array"}, _d.Name{Space: "http:\u002f/purl\u002eoclc\u002eorg\u002fooxml\u002fofficeDocument\u002fdocPropsVTypes", Local: "array"}:
				_agg.Array = _b.NewArray()
				if _dbf := d.DecodeElement(_agg.Array, &_cbf); _dbf != nil {
					return _dbf
				}
			case _d.Name{Space: "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg\u002fofficeDocument\u002f2006\u002fdocPropsVTypes", Local: "blob"}, _d.Name{Space: "http:\u002f/purl\u002eoclc\u002eorg\u002fooxml\u002fofficeDocument\u002fdocPropsVTypes", Local: "blob"}:
				_agg.Blob = new(string)
				if _ab := d.DecodeElement(_agg.Blob, &_cbf); _ab != nil {
					return _ab
				}
			case _d.Name{Space: "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg\u002fofficeDocument\u002f2006\u002fdocPropsVTypes", Local: "oblob"}, _d.Name{Space: "http:\u002f/purl\u002eoclc\u002eorg\u002fooxml\u002fofficeDocument\u002fdocPropsVTypes", Local: "oblob"}:
				_agg.Oblob = new(string)
				if _bdg := d.DecodeElement(_agg.Oblob, &_cbf); _bdg != nil {
					return _bdg
				}
			case _d.Name{Space: "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg\u002fofficeDocument\u002f2006\u002fdocPropsVTypes", Local: "empty"}, _d.Name{Space: "http:\u002f/purl\u002eoclc\u002eorg\u002fooxml\u002fofficeDocument\u002fdocPropsVTypes", Local: "empty"}:
				_agg.Empty = _b.NewEmpty()
				if _dga := d.DecodeElement(_agg.Empty, &_cbf); _dga != nil {
					return _dga
				}
			case _d.Name{Space: "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg\u002fofficeDocument\u002f2006\u002fdocPropsVTypes", Local: "null"}, _d.Name{Space: "http:\u002f/purl\u002eoclc\u002eorg\u002fooxml\u002fofficeDocument\u002fdocPropsVTypes", Local: "null"}:
				_agg.Null = _b.NewNull()
				if _fb := d.DecodeElement(_agg.Null, &_cbf); _fb != nil {
					return _fb
				}
			case _d.Name{Space: "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg\u002fofficeDocument\u002f2006\u002fdocPropsVTypes", Local: "i1"}, _d.Name{Space: "http:\u002f/purl\u002eoclc\u002eorg\u002fooxml\u002fofficeDocument\u002fdocPropsVTypes", Local: "i1"}:
				_agg.I1 = new(int8)
				if _abe := d.DecodeElement(_agg.I1, &_cbf); _abe != nil {
					return _abe
				}
			case _d.Name{Space: "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg\u002fofficeDocument\u002f2006\u002fdocPropsVTypes", Local: "i2"}, _d.Name{Space: "http:\u002f/purl\u002eoclc\u002eorg\u002fooxml\u002fofficeDocument\u002fdocPropsVTypes", Local: "i2"}:
				_agg.I2 = new(int16)
				if _gag := d.DecodeElement(_agg.I2, &_cbf); _gag != nil {
					return _gag
				}
			case _d.Name{Space: "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg\u002fofficeDocument\u002f2006\u002fdocPropsVTypes", Local: "i4"}, _d.Name{Space: "http:\u002f/purl\u002eoclc\u002eorg\u002fooxml\u002fofficeDocument\u002fdocPropsVTypes", Local: "i4"}:
				_agg.I4 = new(int32)
				if _fag := d.DecodeElement(_agg.I4, &_cbf); _fag != nil {
					return _fag
				}
			case _d.Name{Space: "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg\u002fofficeDocument\u002f2006\u002fdocPropsVTypes", Local: "i8"}, _d.Name{Space: "http:\u002f/purl\u002eoclc\u002eorg\u002fooxml\u002fofficeDocument\u002fdocPropsVTypes", Local: "i8"}:
				_agg.I8 = new(int64)
				if _faa := d.DecodeElement(_agg.I8, &_cbf); _faa != nil {
					return _faa
				}
			case _d.Name{Space: "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg\u002fofficeDocument\u002f2006\u002fdocPropsVTypes", Local: "int"}, _d.Name{Space: "http:\u002f/purl\u002eoclc\u002eorg\u002fooxml\u002fofficeDocument\u002fdocPropsVTypes", Local: "int"}:
				_agg.Int = new(int32)
				if _cbg := d.DecodeElement(_agg.Int, &_cbf); _cbg != nil {
					return _cbg
				}
			case _d.Name{Space: "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg\u002fofficeDocument\u002f2006\u002fdocPropsVTypes", Local: "ui1"}, _d.Name{Space: "http:\u002f/purl\u002eoclc\u002eorg\u002fooxml\u002fofficeDocument\u002fdocPropsVTypes", Local: "ui1"}:
				_agg.Ui1 = new(uint8)
				if _eac := d.DecodeElement(_agg.Ui1, &_cbf); _eac != nil {
					return _eac
				}
			case _d.Name{Space: "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg\u002fofficeDocument\u002f2006\u002fdocPropsVTypes", Local: "ui2"}, _d.Name{Space: "http:\u002f/purl\u002eoclc\u002eorg\u002fooxml\u002fofficeDocument\u002fdocPropsVTypes", Local: "ui2"}:
				_agg.Ui2 = new(uint16)
				if _ggea := d.DecodeElement(_agg.Ui2, &_cbf); _ggea != nil {
					return _ggea
				}
			case _d.Name{Space: "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg\u002fofficeDocument\u002f2006\u002fdocPropsVTypes", Local: "ui4"}, _d.Name{Space: "http:\u002f/purl\u002eoclc\u002eorg\u002fooxml\u002fofficeDocument\u002fdocPropsVTypes", Local: "ui4"}:
				_agg.Ui4 = new(uint32)
				if _ad := d.DecodeElement(_agg.Ui4, &_cbf); _ad != nil {
					return _ad
				}
			case _d.Name{Space: "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg\u002fofficeDocument\u002f2006\u002fdocPropsVTypes", Local: "ui8"}, _d.Name{Space: "http:\u002f/purl\u002eoclc\u002eorg\u002fooxml\u002fofficeDocument\u002fdocPropsVTypes", Local: "ui8"}:
				_agg.Ui8 = new(uint64)
				if _aea := d.DecodeElement(_agg.Ui8, &_cbf); _aea != nil {
					return _aea
				}
			case _d.Name{Space: "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg\u002fofficeDocument\u002f2006\u002fdocPropsVTypes", Local: "uint"}, _d.Name{Space: "http:\u002f/purl\u002eoclc\u002eorg\u002fooxml\u002fofficeDocument\u002fdocPropsVTypes", Local: "uint"}:
				_agg.Uint = new(uint32)
				if _cgc := d.DecodeElement(_agg.Uint, &_cbf); _cgc != nil {
					return _cgc
				}
			case _d.Name{Space: "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg\u002fofficeDocument\u002f2006\u002fdocPropsVTypes", Local: "r4"}, _d.Name{Space: "http:\u002f/purl\u002eoclc\u002eorg\u002fooxml\u002fofficeDocument\u002fdocPropsVTypes", Local: "r4"}:
				_agg.R4 = new(float32)
				if _cbb := d.DecodeElement(_agg.R4, &_cbf); _cbb != nil {
					return _cbb
				}
			case _d.Name{Space: "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg\u002fofficeDocument\u002f2006\u002fdocPropsVTypes", Local: "r8"}, _d.Name{Space: "http:\u002f/purl\u002eoclc\u002eorg\u002fooxml\u002fofficeDocument\u002fdocPropsVTypes", Local: "r8"}:
				_agg.R8 = new(float64)
				if _aef := d.DecodeElement(_agg.R8, &_cbf); _aef != nil {
					return _aef
				}
			case _d.Name{Space: "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg\u002fofficeDocument\u002f2006\u002fdocPropsVTypes", Local: "decimal"}, _d.Name{Space: "http:\u002f/purl\u002eoclc\u002eorg\u002fooxml\u002fofficeDocument\u002fdocPropsVTypes", Local: "decimal"}:
				_agg.Decimal = new(float64)
				if _fcb := d.DecodeElement(_agg.Decimal, &_cbf); _fcb != nil {
					return _fcb
				}
			case _d.Name{Space: "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg\u002fofficeDocument\u002f2006\u002fdocPropsVTypes", Local: "lpstr"}, _d.Name{Space: "http:\u002f/purl\u002eoclc\u002eorg\u002fooxml\u002fofficeDocument\u002fdocPropsVTypes", Local: "lpstr"}:
				_agg.Lpstr = new(string)
				if _degf := d.DecodeElement(_agg.Lpstr, &_cbf); _degf != nil {
					return _degf
				}
			case _d.Name{Space: "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg\u002fofficeDocument\u002f2006\u002fdocPropsVTypes", Local: "lpwstr"}, _d.Name{Space: "http:\u002f/purl\u002eoclc\u002eorg\u002fooxml\u002fofficeDocument\u002fdocPropsVTypes", Local: "lpwstr"}:
				_agg.Lpwstr = new(string)
				if _egb := d.DecodeElement(_agg.Lpwstr, &_cbf); _egb != nil {
					return _egb
				}
			case _d.Name{Space: "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg\u002fofficeDocument\u002f2006\u002fdocPropsVTypes", Local: "bstr"}, _d.Name{Space: "http:\u002f/purl\u002eoclc\u002eorg\u002fooxml\u002fofficeDocument\u002fdocPropsVTypes", Local: "bstr"}:
				_agg.Bstr = new(string)
				if _dfc := d.DecodeElement(_agg.Bstr, &_cbf); _dfc != nil {
					return _dfc
				}
			case _d.Name{Space: "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg\u002fofficeDocument\u002f2006\u002fdocPropsVTypes", Local: "date"}, _d.Name{Space: "http:\u002f/purl\u002eoclc\u002eorg\u002fooxml\u002fofficeDocument\u002fdocPropsVTypes", Local: "date"}:
				_agg.Date = new(_e.Time)
				if _eb := d.DecodeElement(_agg.Date, &_cbf); _eb != nil {
					return _eb
				}
			case _d.Name{Space: "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg\u002fofficeDocument\u002f2006\u002fdocPropsVTypes", Local: "filetime"}, _d.Name{Space: "http:\u002f/purl\u002eoclc\u002eorg\u002fooxml\u002fofficeDocument\u002fdocPropsVTypes", Local: "filetime"}:
				_agg.Filetime = new(_e.Time)
				if _cee := d.DecodeElement(_agg.Filetime, &_cbf); _cee != nil {
					return _cee
				}
			case _d.Name{Space: "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg\u002fofficeDocument\u002f2006\u002fdocPropsVTypes", Local: "bool"}, _d.Name{Space: "http:\u002f/purl\u002eoclc\u002eorg\u002fooxml\u002fofficeDocument\u002fdocPropsVTypes", Local: "bool"}:
				_agg.Bool = new(bool)
				if _ecb := d.DecodeElement(_agg.Bool, &_cbf); _ecb != nil {
					return _ecb
				}
			case _d.Name{Space: "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg\u002fofficeDocument\u002f2006\u002fdocPropsVTypes", Local: "cy"}, _d.Name{Space: "http:\u002f/purl\u002eoclc\u002eorg\u002fooxml\u002fofficeDocument\u002fdocPropsVTypes", Local: "cy"}:
				_agg.Cy = new(string)
				if _cbd := d.DecodeElement(_agg.Cy, &_cbf); _cbd != nil {
					return _cbd
				}
			case _d.Name{Space: "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg\u002fofficeDocument\u002f2006\u002fdocPropsVTypes", Local: "error"}, _d.Name{Space: "http:\u002f/purl\u002eoclc\u002eorg\u002fooxml\u002fofficeDocument\u002fdocPropsVTypes", Local: "error"}:
				_agg.Error = new(string)
				if _gdb := d.DecodeElement(_agg.Error, &_cbf); _gdb != nil {
					return _gdb
				}
			case _d.Name{Space: "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg\u002fofficeDocument\u002f2006\u002fdocPropsVTypes", Local: "stream"}, _d.Name{Space: "http:\u002f/purl\u002eoclc\u002eorg\u002fooxml\u002fofficeDocument\u002fdocPropsVTypes", Local: "stream"}:
				_agg.Stream = new(string)
				if _fac := d.DecodeElement(_agg.Stream, &_cbf); _fac != nil {
					return _fac
				}
			case _d.Name{Space: "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg\u002fofficeDocument\u002f2006\u002fdocPropsVTypes", Local: "ostream"}, _d.Name{Space: "http:\u002f/purl\u002eoclc\u002eorg\u002fooxml\u002fofficeDocument\u002fdocPropsVTypes", Local: "ostream"}:
				_agg.Ostream = new(string)
				if _faff := d.DecodeElement(_agg.Ostream, &_cbf); _faff != nil {
					return _faff
				}
			case _d.Name{Space: "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg\u002fofficeDocument\u002f2006\u002fdocPropsVTypes", Local: "storage"}, _d.Name{Space: "http:\u002f/purl\u002eoclc\u002eorg\u002fooxml\u002fofficeDocument\u002fdocPropsVTypes", Local: "storage"}:
				_agg.Storage = new(string)
				if _abb := d.DecodeElement(_agg.Storage, &_cbf); _abb != nil {
					return _abb
				}
			case _d.Name{Space: "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg\u002fofficeDocument\u002f2006\u002fdocPropsVTypes", Local: "ostorage"}, _d.Name{Space: "http:\u002f/purl\u002eoclc\u002eorg\u002fooxml\u002fofficeDocument\u002fdocPropsVTypes", Local: "ostorage"}:
				_agg.Ostorage = new(string)
				if _dgg := d.DecodeElement(_agg.Ostorage, &_cbf); _dgg != nil {
					return _dgg
				}
			case _d.Name{Space: "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg\u002fofficeDocument\u002f2006\u002fdocPropsVTypes", Local: "vstream"}, _d.Name{Space: "http:\u002f/purl\u002eoclc\u002eorg\u002fooxml\u002fofficeDocument\u002fdocPropsVTypes", Local: "vstream"}:
				_agg.Vstream = _b.NewVstream()
				if _cfa := d.DecodeElement(_agg.Vstream, &_cbf); _cfa != nil {
					return _cfa
				}
			case _d.Name{Space: "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg\u002fofficeDocument\u002f2006\u002fdocPropsVTypes", Local: "clsid"}, _d.Name{Space: "http:\u002f/purl\u002eoclc\u002eorg\u002fooxml\u002fofficeDocument\u002fdocPropsVTypes", Local: "clsid"}:
				_agg.Clsid = new(string)
				if _dea := d.DecodeElement(_agg.Clsid, &_cbf); _dea != nil {
					return _dea
				}
			default:
				_ef.Log("skipping\u0020unsupported\u0020element\u0020on\u0020CT_Property\u0020\u0025v", _cbf.Name)
				if _cgf := d.Skip(); _cgf != nil {
					return _cgf
				}
			}
		case _d.EndElement:
			break _bbg
		case _d.CharData:
		}
	}
	return nil
}
func (_ed *CT_Properties) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
_c:
	for {
		_cd, _bb := d.Token()
		if _bb != nil {
			return _bb
		}
		switch _gg := _cd.(type) {
		case _d.StartElement:
			switch _gg.Name {
			case _d.Name{Space: "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg/officeDocument\u002f2006/custom\u002dproperties", Local: "property"}:
				_fc := NewCT_Property()
				if _ce := d.DecodeElement(_fc, &_gg); _ce != nil {
					return _ce
				}
				_ed.Property = append(_ed.Property, _fc)
			default:
				_ef.Log("skipping unsupported element\u0020on\u0020CT_Properties \u0025v", _gg.Name)
				if _cf := d.Skip(); _cf != nil {
					return _cf
				}
			}
		case _d.EndElement:
			break _c
		case _d.CharData:
		}
	}
	return nil
}
func (_eg *CT_Property) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "fmtid"}, Value: _da.Sprintf("\u0025v", _eg.FmtidAttr)})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "pid"}, Value: _da.Sprintf("\u0025v", _eg.PidAttr)})
	if _eg.NameAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "name"}, Value: _da.Sprintf("\u0025v", *_eg.NameAttr)})
	}
	if _eg.LinkTargetAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "linkTarget"}, Value: _da.Sprintf("\u0025v", *_eg.LinkTargetAttr)})
	}
	e.EncodeToken(start)
	if _eg.Vector != nil {
		_bg := _d.StartElement{Name: _d.Name{Local: "vt:vector"}}
		e.EncodeElement(_eg.Vector, _bg)
	}
	if _eg.Array != nil {
		_dc := _d.StartElement{Name: _d.Name{Local: "vt:array"}}
		e.EncodeElement(_eg.Array, _dc)
	}
	if _eg.Blob != nil {
		_gc := _d.StartElement{Name: _d.Name{Local: "vt:blob"}}
		_ef.AddPreserveSpaceAttr(&_gc, *_eg.Blob)
		e.EncodeElement(_eg.Blob, _gc)
	}
	if _eg.Oblob != nil {
		_gb := _d.StartElement{Name: _d.Name{Local: "vt:oblob"}}
		_ef.AddPreserveSpaceAttr(&_gb, *_eg.Oblob)
		e.EncodeElement(_eg.Oblob, _gb)
	}
	if _eg.Empty != nil {
		_eef := _d.StartElement{Name: _d.Name{Local: "vt:empty"}}
		e.EncodeElement(_eg.Empty, _eef)
	}
	if _eg.Null != nil {
		_ec := _d.StartElement{Name: _d.Name{Local: "vt:null"}}
		e.EncodeElement(_eg.Null, _ec)
	}
	if _eg.I1 != nil {
		_gge := _d.StartElement{Name: _d.Name{Local: "vt:i1"}}
		e.EncodeElement(_eg.I1, _gge)
	}
	if _eg.I2 != nil {
		_a := _d.StartElement{Name: _d.Name{Local: "vt:i2"}}
		e.EncodeElement(_eg.I2, _a)
	}
	if _eg.I4 != nil {
		_bgc := _d.StartElement{Name: _d.Name{Local: "vt:i4"}}
		e.EncodeElement(_eg.I4, _bgc)
	}
	if _eg.I8 != nil {
		_df := _d.StartElement{Name: _d.Name{Local: "vt:i8"}}
		e.EncodeElement(_eg.I8, _df)
	}
	if _eg.Int != nil {
		_gf := _d.StartElement{Name: _d.Name{Local: "vt:int"}}
		e.EncodeElement(_eg.Int, _gf)
	}
	if _eg.Ui1 != nil {
		_gfd := _d.StartElement{Name: _d.Name{Local: "vt:ui1"}}
		e.EncodeElement(_eg.Ui1, _gfd)
	}
	if _eg.Ui2 != nil {
		_bda := _d.StartElement{Name: _d.Name{Local: "vt:ui2"}}
		e.EncodeElement(_eg.Ui2, _bda)
	}
	if _eg.Ui4 != nil {
		_cg := _d.StartElement{Name: _d.Name{Local: "vt:ui4"}}
		e.EncodeElement(_eg.Ui4, _cg)
	}
	if _eg.Ui8 != nil {
		_ae := _d.StartElement{Name: _d.Name{Local: "vt:ui8"}}
		e.EncodeElement(_eg.Ui8, _ae)
	}
	if _eg.Uint != nil {
		_edc := _d.StartElement{Name: _d.Name{Local: "vt:uint"}}
		e.EncodeElement(_eg.Uint, _edc)
	}
	if _eg.R4 != nil {
		_ag := _d.StartElement{Name: _d.Name{Local: "vt:r4"}}
		e.EncodeElement(_eg.R4, _ag)
	}
	if _eg.R8 != nil {
		_cb := _d.StartElement{Name: _d.Name{Local: "vt:r8"}}
		e.EncodeElement(_eg.R8, _cb)
	}
	if _eg.Decimal != nil {
		_dag := _d.StartElement{Name: _d.Name{Local: "vt:decimal"}}
		e.EncodeElement(_eg.Decimal, _dag)
	}
	if _eg.Lpstr != nil {
		_agd := _d.StartElement{Name: _d.Name{Local: "vt:lpstr"}}
		_ef.AddPreserveSpaceAttr(&_agd, *_eg.Lpstr)
		e.EncodeElement(_eg.Lpstr, _agd)
	}
	if _eg.Lpwstr != nil {
		_ac := _d.StartElement{Name: _d.Name{Local: "vt:lpwstr"}}
		_ef.AddPreserveSpaceAttr(&_ac, *_eg.Lpwstr)
		e.EncodeElement(_eg.Lpwstr, _ac)
	}
	if _eg.Bstr != nil {
		_dfb := _d.StartElement{Name: _d.Name{Local: "vt:bstr"}}
		_ef.AddPreserveSpaceAttr(&_dfb, *_eg.Bstr)
		e.EncodeElement(_eg.Bstr, _dfb)
	}
	if _eg.Date != nil {
		_deg := _d.StartElement{Name: _d.Name{Local: "vt:date"}}
		e.EncodeElement(_eg.Date, _deg)
	}
	if _eg.Filetime != nil {
		_faf := _d.StartElement{Name: _d.Name{Local: "vt:filetime"}}
		e.EncodeElement(_eg.Filetime, _faf)
	}
	if _eg.Bool != nil {
		_aa := _d.StartElement{Name: _d.Name{Local: "vt:bool"}}
		e.EncodeElement(_eg.Bool, _aa)
	}
	if _eg.Cy != nil {
		_fcf := _d.StartElement{Name: _d.Name{Local: "vt:cy"}}
		_ef.AddPreserveSpaceAttr(&_fcf, *_eg.Cy)
		e.EncodeElement(_eg.Cy, _fcf)
	}
	if _eg.Error != nil {
		_gce := _d.StartElement{Name: _d.Name{Local: "vt:error"}}
		_ef.AddPreserveSpaceAttr(&_gce, *_eg.Error)
		e.EncodeElement(_eg.Error, _gce)
	}
	if _eg.Stream != nil {
		_fe := _d.StartElement{Name: _d.Name{Local: "vt:stream"}}
		_ef.AddPreserveSpaceAttr(&_fe, *_eg.Stream)
		e.EncodeElement(_eg.Stream, _fe)
	}
	if _eg.Ostream != nil {
		_dbc := _d.StartElement{Name: _d.Name{Local: "vt:ostream"}}
		_ef.AddPreserveSpaceAttr(&_dbc, *_eg.Ostream)
		e.EncodeElement(_eg.Ostream, _dbc)
	}
	if _eg.Storage != nil {
		_fg := _d.StartElement{Name: _d.Name{Local: "vt:storage"}}
		_ef.AddPreserveSpaceAttr(&_fg, *_eg.Storage)
		e.EncodeElement(_eg.Storage, _fg)
	}
	if _eg.Ostorage != nil {
		_gcg := _d.StartElement{Name: _d.Name{Local: "vt:ostorage"}}
		_ef.AddPreserveSpaceAttr(&_gcg, *_eg.Ostorage)
		e.EncodeElement(_eg.Ostorage, _gcg)
	}
	if _eg.Vstream != nil {
		_cce := _d.StartElement{Name: _d.Name{Local: "vt:vstream"}}
		e.EncodeElement(_eg.Vstream, _cce)
	}
	if _eg.Clsid != nil {
		_fdc := _d.StartElement{Name: _d.Name{Local: "vt:clsid"}}
		_ef.AddPreserveSpaceAttr(&_fdc, *_eg.Clsid)
		e.EncodeElement(_eg.Clsid, _fdc)
	}
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}
func NewProperties() *Properties {
	_eeg := &Properties{}
	_eeg.CT_Properties = *NewCT_Properties()
	return _eeg
}

// Validate validates the Properties and its children
func (_acd *Properties) Validate() error { return _acd.ValidateWithPath("Properties") }

type CT_Property struct {
	FmtidAttr      string
	PidAttr        int32
	NameAttr       *string
	LinkTargetAttr *string
	Vector         *_b.Vector
	Array          *_b.Array
	Blob           *string
	Oblob          *string
	Empty          *_b.Empty
	Null           *_b.Null
	I1             *int8
	I2             *int16
	I4             *int32
	I8             *int64
	Int            *int32
	Ui1            *uint8
	Ui2            *uint16
	Ui4            *uint32
	Ui8            *uint64
	Uint           *uint32
	R4             *float32
	R8             *float64
	Decimal        *float64
	Lpstr          *string
	Lpwstr         *string
	Bstr           *string
	Date           *_e.Time
	Filetime       *_e.Time
	Bool           *bool
	Cy             *string
	Error          *string
	Stream         *string
	Ostream        *string
	Storage        *string
	Ostorage       *string
	Vstream        *_b.Vstream
	Clsid          *string
}
type CT_Properties struct{ Property []*CT_Property }

// ValidateWithPath validates the Properties and its children, prefixing error messages with path
func (_bde *Properties) ValidateWithPath(path string) error {
	if _ebe := _bde.CT_Properties.ValidateWithPath(path); _ebe != nil {
		return _ebe
	}
	return nil
}
func (_fcfa *Properties) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_fcfa.CT_Properties = *NewCT_Properties()
_gdd:
	for {
		_eec, _af := d.Token()
		if _af != nil {
			return _af
		}
		switch _afg := _eec.(type) {
		case _d.StartElement:
			switch _afg.Name {
			case _d.Name{Space: "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg/officeDocument\u002f2006/custom\u002dproperties", Local: "property"}:
				_eae := NewCT_Property()
				if _bfd := d.DecodeElement(_eae, &_afg); _bfd != nil {
					return _bfd
				}
				_fcfa.Property = append(_fcfa.Property, _eae)
			default:
				_ef.Log("skipping unsupported\u0020element\u0020on\u0020Properties\u0020\u0025v", _afg.Name)
				if _bgg := d.Skip(); _bgg != nil {
					return _bgg
				}
			}
		case _d.EndElement:
			break _gdd
		case _d.CharData:
		}
	}
	return nil
}

// ValidateWithPath validates the CT_Property and its children, prefixing error messages with path
func (_dcb *CT_Property) ValidateWithPath(path string) error {
	if !_bc.ST_GuidPatternRe.MatchString(_dcb.FmtidAttr) {
		return _da.Errorf("%s\u002fm\u002eFmtidAttr\u0020must match\u0020\u0027\u0025s\u0027 (have\u0020%v\u0029", path, _bc.ST_GuidPatternRe, _dcb.FmtidAttr)
	}
	if _dcb.Vector != nil {
		if _dbb := _dcb.Vector.ValidateWithPath(path + "\u002fVector"); _dbb != nil {
			return _dbb
		}
	}
	if _dcb.Array != nil {
		if _dcd := _dcb.Array.ValidateWithPath(path + "\u002fArray"); _dcd != nil {
			return _dcd
		}
	}
	if _dcb.Empty != nil {
		if _cfc := _dcb.Empty.ValidateWithPath(path + "\u002fEmpty"); _cfc != nil {
			return _cfc
		}
	}
	if _dcb.Null != nil {
		if _fdd := _dcb.Null.ValidateWithPath(path + "\u002fNull"); _fdd != nil {
			return _fdd
		}
	}
	if _dcb.Cy != nil {
		if !_b.ST_CyPatternRe.MatchString(*_dcb.Cy) {
			return _da.Errorf("\u0025s\u002fm\u002eCy\u0020must\u0020match\u0020\u0027%s\u0027\u0020\u0028have\u0020\u0025v\u0029", path, _b.ST_CyPatternRe, *_dcb.Cy)
		}
	}
	if _dcb.Error != nil {
		if !_b.ST_ErrorPatternRe.MatchString(*_dcb.Error) {
			return _da.Errorf("\u0025s/m\u002eError\u0020must\u0020match\u0020\u0027\u0025s' \u0028have\u0020\u0025v\u0029", path, _b.ST_ErrorPatternRe, *_dcb.Error)
		}
	}
	if _dcb.Vstream != nil {
		if _fgb := _dcb.Vstream.ValidateWithPath(path + "\u002fVstream"); _fgb != nil {
			return _fgb
		}
	}
	if _dcb.Clsid != nil {
		if !_bc.ST_GuidPatternRe.MatchString(*_dcb.Clsid) {
			return _da.Errorf("\u0025s/m\u002eClsid\u0020must\u0020match\u0020\u0027\u0025s' \u0028have\u0020\u0025v\u0029", path, _bc.ST_GuidPatternRe, *_dcb.Clsid)
		}
	}
	return nil
}
func (_gdbd *Properties) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "xmlns"}, Value: "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg/officeDocument\u002f2006/custom\u002dproperties"})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "xmlns:s"}, Value: "http:/\u002fschemas\u002eopenxmlformats\u002eorg/officeDocument\u002f2006\u002fsharedTypes"})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "xmlns:vt"}, Value: "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg\u002fofficeDocument\u002f2006\u002fdocPropsVTypes"})
//	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "xmlns:xxml"}, Value: "http:\u002f\u002fwww\u002ew3.org/XML\u002f1998/namespace"})
	start.Name.Local = "Properties"
	return _gdbd.CT_Properties.MarshalXML(e, start)
}
func init() {
	_ef.RegisterConstructor("http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg/officeDocument\u002f2006/custom\u002dproperties", "CT_Properties", NewCT_Properties)
	_ef.RegisterConstructor("http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg/officeDocument\u002f2006/custom\u002dproperties", "CT_Property", NewCT_Property)
	_ef.RegisterConstructor("http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg/officeDocument\u002f2006/custom\u002dproperties", "Properties", NewProperties)
}
