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

package core_properties

import (
	_g "encoding/xml"
	_aa "fmt"
	_bc "gitee.com/greatmusicians/unioffice"
	_b "time"
)

func (_ae *CT_CoreProperties) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _ae.Category != nil {
		_eb := _g.StartElement{Name: _g.Name{Local: "cp:category"}}
		_bc.AddPreserveSpaceAttr(&_eb, *_ae.Category)
		e.EncodeElement(_ae.Category, _eb)
	}
	if _ae.ContentStatus != nil {
		_d := _g.StartElement{Name: _g.Name{Local: "cp:contentStatus"}}
		_bc.AddPreserveSpaceAttr(&_d, *_ae.ContentStatus)
		e.EncodeElement(_ae.ContentStatus, _d)
	}
	if _ae.Created != nil {
		_de := _g.StartElement{Name: _g.Name{Local: "dcterms:created"}}
		e.EncodeElement(_ae.Created, _de)
	}
	if _ae.Creator != nil {
		_ga := _g.StartElement{Name: _g.Name{Local: "dc:creator"}}
		e.EncodeElement(_ae.Creator, _ga)
	}
	if _ae.Description != nil {
		_f := _g.StartElement{Name: _g.Name{Local: "dc:description"}}
		e.EncodeElement(_ae.Description, _f)
	}
	if _ae.Identifier != nil {
		_c := _g.StartElement{Name: _g.Name{Local: "dc:identifier"}}
		e.EncodeElement(_ae.Identifier, _c)
	}
	if _ae.Keywords != nil {
		_be := _g.StartElement{Name: _g.Name{Local: "cp:keywords"}}
		e.EncodeElement(_ae.Keywords, _be)
	}
	if _ae.Language != nil {
		_ge := _g.StartElement{Name: _g.Name{Local: "dc:language"}}
		e.EncodeElement(_ae.Language, _ge)
	}
	if _ae.LastModifiedBy != nil {
		_gad := _g.StartElement{Name: _g.Name{Local: "cp:lastModifiedBy"}}
		_bc.AddPreserveSpaceAttr(&_gad, *_ae.LastModifiedBy)
		e.EncodeElement(_ae.LastModifiedBy, _gad)
	}
	if _ae.LastPrinted != nil {
		_gef := _g.StartElement{Name: _g.Name{Local: "cp:lastPrinted"}}
		e.EncodeElement(_ae.LastPrinted, _gef)
	}
	if _ae.Modified != nil {
		_bf := _g.StartElement{Name: _g.Name{Local: "dcterms:modified"}}
		e.EncodeElement(_ae.Modified, _bf)
	}
	if _ae.Revision != nil {
		_ebe := _g.StartElement{Name: _g.Name{Local: "cp:revision"}}
		_bc.AddPreserveSpaceAttr(&_ebe, *_ae.Revision)
		e.EncodeElement(_ae.Revision, _ebe)
	}
	if _ae.Subject != nil {
		_beb := _g.StartElement{Name: _g.Name{Local: "dc:subject"}}
		e.EncodeElement(_ae.Subject, _beb)
	}
	if _ae.Title != nil {
		_fb := _g.StartElement{Name: _g.Name{Local: "dc:title"}}
		e.EncodeElement(_ae.Title, _fb)
	}
	if _ae.Version != nil {
		_ag := _g.StartElement{Name: _g.Name{Local: "cp:version"}}
		_bc.AddPreserveSpaceAttr(&_ag, *_ae.Version)
		e.EncodeElement(_ae.Version, _ag)
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

// Validate validates the CoreProperties and its children
func (_ggc *CoreProperties) Validate() error { return _ggc.ValidateWithPath("CoreProperties") }
func (_dcc *CoreProperties) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_dcc.CT_CoreProperties = *NewCT_CoreProperties()
_bd:
	for {
		_aaa, _eg := d.Token()
		if _eg != nil {
			return _eg
		}
		switch _ef := _aaa.(type) {
		case _g.StartElement:
			switch _ef.Name {
			case _g.Name{Space: "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg/package\u002f2006\u002fmetadata\u002fcore\u002dproperties", Local: "category"}:
				_dcc.Category = new(string)
				if _dead := d.DecodeElement(_dcc.Category, &_ef); _dead != nil {
					return _dead
				}
			case _g.Name{Space: "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg/package\u002f2006\u002fmetadata\u002fcore\u002dproperties", Local: "contentStatus"}:
				_dcc.ContentStatus = new(string)
				if _gabb := d.DecodeElement(_dcc.ContentStatus, &_ef); _gabb != nil {
					return _gabb
				}
			case _g.Name{Space: "http:\u002f/purl\u002eorg\u002fdc\u002fterms/", Local: "created"}:
				_dcc.Created = new(_bc.XSDAny)
				if _ggg := d.DecodeElement(_dcc.Created, &_ef); _ggg != nil {
					return _ggg
				}
			case _g.Name{Space: "http:\u002f\u002fpurl\u002eorg/dc\u002felements\u002f1\u002e1\u002f", Local: "creator"}:
				_dcc.Creator = new(_bc.XSDAny)
				if _edfe := d.DecodeElement(_dcc.Creator, &_ef); _edfe != nil {
					return _edfe
				}
			case _g.Name{Space: "http:\u002f\u002fpurl\u002eorg/dc\u002felements\u002f1\u002e1\u002f", Local: "description"}:
				_dcc.Description = new(_bc.XSDAny)
				if _ggb := d.DecodeElement(_dcc.Description, &_ef); _ggb != nil {
					return _ggb
				}
			case _g.Name{Space: "http:\u002f\u002fpurl\u002eorg/dc\u002felements\u002f1\u002e1\u002f", Local: "identifier"}:
				_dcc.Identifier = new(_bc.XSDAny)
				if _cb := d.DecodeElement(_dcc.Identifier, &_ef); _cb != nil {
					return _cb
				}
			case _g.Name{Space: "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg/package\u002f2006\u002fmetadata\u002fcore\u002dproperties", Local: "keywords"}:
				_dcc.Keywords = NewCT_Keywords()
				if _dgbf := d.DecodeElement(_dcc.Keywords, &_ef); _dgbf != nil {
					return _dgbf
				}
			case _g.Name{Space: "http:\u002f\u002fpurl\u002eorg/dc\u002felements\u002f1\u002e1\u002f", Local: "language"}:
				_dcc.Language = new(_bc.XSDAny)
				if _gfgb := d.DecodeElement(_dcc.Language, &_ef); _gfgb != nil {
					return _gfgb
				}
			case _g.Name{Space: "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg/package\u002f2006\u002fmetadata\u002fcore\u002dproperties", Local: "lastModifiedBy"}:
				_dcc.LastModifiedBy = new(string)
				if _bbb := d.DecodeElement(_dcc.LastModifiedBy, &_ef); _bbb != nil {
					return _bbb
				}
			case _g.Name{Space: "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg/package\u002f2006\u002fmetadata\u002fcore\u002dproperties", Local: "lastPrinted"}:
				_dcc.LastPrinted = new(_b.Time)
				if _gd := d.DecodeElement(_dcc.LastPrinted, &_ef); _gd != nil {
					return _gd
				}
			case _g.Name{Space: "http:\u002f/purl\u002eorg\u002fdc\u002fterms/", Local: "modified"}:
				_dcc.Modified = new(_bc.XSDAny)
				if _dab := d.DecodeElement(_dcc.Modified, &_ef); _dab != nil {
					return _dab
				}
			case _g.Name{Space: "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg/package\u002f2006\u002fmetadata\u002fcore\u002dproperties", Local: "revision"}:
				_dcc.Revision = new(string)
				if _cbf := d.DecodeElement(_dcc.Revision, &_ef); _cbf != nil {
					return _cbf
				}
			case _g.Name{Space: "http:\u002f\u002fpurl\u002eorg/dc\u002felements\u002f1\u002e1\u002f", Local: "subject"}:
				_dcc.Subject = new(_bc.XSDAny)
				if _egc := d.DecodeElement(_dcc.Subject, &_ef); _egc != nil {
					return _egc
				}
			case _g.Name{Space: "http:\u002f\u002fpurl\u002eorg/dc\u002felements\u002f1\u002e1\u002f", Local: "title"}:
				_dcc.Title = new(_bc.XSDAny)
				if _cecf := d.DecodeElement(_dcc.Title, &_ef); _cecf != nil {
					return _cecf
				}
			case _g.Name{Space: "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg/package\u002f2006\u002fmetadata\u002fcore\u002dproperties", Local: "version"}:
				_dcc.Version = new(string)
				if _dfd := d.DecodeElement(_dcc.Version, &_ef); _dfd != nil {
					return _dfd
				}
			default:
				_bc.Log("skipping\u0020unsupported\u0020element\u0020on\u0020CoreProperties\u0020\u0025v", _ef.Name)
				if _aea := d.Skip(); _aea != nil {
					return _aea
				}
			}
		case _g.EndElement:
			break _bd
		case _g.CharData:
		}
	}
	return nil
}
func (_da *CT_Keyword) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	if _da.LangAttr != nil {
		start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "xml:lang"}, Value: _aa.Sprintf("\u0025v", *_da.LangAttr)})
	}
	e.EncodeElement(_da.Content, start)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

// Validate validates the CT_CoreProperties and its children
func (_dbfd *CT_CoreProperties) Validate() error { return _dbfd.ValidateWithPath("CT_CoreProperties") }

type CT_Keyword struct {
	LangAttr *string
	Content  string
}

// ValidateWithPath validates the CT_CoreProperties and its children, prefixing error messages with path
func (_dg *CT_CoreProperties) ValidateWithPath(path string) error {
	if _dg.Keywords != nil {
		if _df := _dg.Keywords.ValidateWithPath(path + "\u002fKeywords"); _df != nil {
			return _df
		}
	}
	return nil
}
func NewCT_CoreProperties() *CT_CoreProperties { _e := &CT_CoreProperties{}; return _e }
func (_acc *CT_Keywords) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	if _acc.LangAttr != nil {
		start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "xml:lang"}, Value: _aa.Sprintf("\u0025v", *_acc.LangAttr)})
	}
	e.EncodeToken(start)
	if _acc.Value != nil {
		_ea := _g.StartElement{Name: _g.Name{Local: "cp:value"}}
		for _, _ede := range _acc.Value {
			e.EncodeElement(_ede, _ea)
		}
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}
func (_ed *CT_CoreProperties) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_ff:
	for {
		_cd, _fd := d.Token()
		if _fd != nil {
			return _fd
		}
		switch _ce := _cd.(type) {
		case _g.StartElement:
			switch _ce.Name {
			case _g.Name{Space: "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg/package\u002f2006\u002fmetadata\u002fcore\u002dproperties", Local: "category"}:
				_ed.Category = new(string)
				if _dc := d.DecodeElement(_ed.Category, &_ce); _dc != nil {
					return _dc
				}
			case _g.Name{Space: "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg/package\u002f2006\u002fmetadata\u002fcore\u002dproperties", Local: "contentStatus"}:
				_ed.ContentStatus = new(string)
				if _dea := d.DecodeElement(_ed.ContentStatus, &_ce); _dea != nil {
					return _dea
				}
			case _g.Name{Space: "http:\u002f/purl\u002eorg\u002fdc\u002fterms/", Local: "created"}:
				_ed.Created = new(_bc.XSDAny)
				if _ebf := d.DecodeElement(_ed.Created, &_ce); _ebf != nil {
					return _ebf
				}
			case _g.Name{Space: "http:\u002f\u002fpurl\u002eorg/dc\u002felements\u002f1\u002e1\u002f", Local: "creator"}:
				_ed.Creator = new(_bc.XSDAny)
				if _fa := d.DecodeElement(_ed.Creator, &_ce); _fa != nil {
					return _fa
				}
			case _g.Name{Space: "http:\u002f\u002fpurl\u002eorg/dc\u002felements\u002f1\u002e1\u002f", Local: "description"}:
				_ed.Description = new(_bc.XSDAny)
				if _eda := d.DecodeElement(_ed.Description, &_ce); _eda != nil {
					return _eda
				}
			case _g.Name{Space: "http:\u002f\u002fpurl\u002eorg/dc\u002felements\u002f1\u002e1\u002f", Local: "identifier"}:
				_ed.Identifier = new(_bc.XSDAny)
				if _ac := d.DecodeElement(_ed.Identifier, &_ce); _ac != nil {
					return _ac
				}
			case _g.Name{Space: "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg/package\u002f2006\u002fmetadata\u002fcore\u002dproperties", Local: "keywords"}:
				_ed.Keywords = NewCT_Keywords()
				if _cg := d.DecodeElement(_ed.Keywords, &_ce); _cg != nil {
					return _cg
				}
			case _g.Name{Space: "http:\u002f\u002fpurl\u002eorg/dc\u002felements\u002f1\u002e1\u002f", Local: "language"}:
				_ed.Language = new(_bc.XSDAny)
				if _deg := d.DecodeElement(_ed.Language, &_ce); _deg != nil {
					return _deg
				}
			case _g.Name{Space: "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg/package\u002f2006\u002fmetadata\u002fcore\u002dproperties", Local: "lastModifiedBy"}:
				_ed.LastModifiedBy = new(string)
				if _db := d.DecodeElement(_ed.LastModifiedBy, &_ce); _db != nil {
					return _db
				}
			case _g.Name{Space: "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg/package\u002f2006\u002fmetadata\u002fcore\u002dproperties", Local: "lastPrinted"}:
				_ed.LastPrinted = new(_b.Time)
				if _cc := d.DecodeElement(_ed.LastPrinted, &_ce); _cc != nil {
					return _cc
				}
			case _g.Name{Space: "http:\u002f/purl\u002eorg\u002fdc\u002fterms/", Local: "modified"}:
				_ed.Modified = new(_bc.XSDAny)
				if _dbf := d.DecodeElement(_ed.Modified, &_ce); _dbf != nil {
					return _dbf
				}
			case _g.Name{Space: "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg/package\u002f2006\u002fmetadata\u002fcore\u002dproperties", Local: "revision"}:
				_ed.Revision = new(string)
				if _gg := d.DecodeElement(_ed.Revision, &_ce); _gg != nil {
					return _gg
				}
			case _g.Name{Space: "http:\u002f\u002fpurl\u002eorg/dc\u002felements\u002f1\u002e1\u002f", Local: "subject"}:
				_ed.Subject = new(_bc.XSDAny)
				if _bg := d.DecodeElement(_ed.Subject, &_ce); _bg != nil {
					return _bg
				}
			case _g.Name{Space: "http:\u002f\u002fpurl\u002eorg/dc\u002felements\u002f1\u002e1\u002f", Local: "title"}:
				_ed.Title = new(_bc.XSDAny)
				if _bgg := d.DecodeElement(_ed.Title, &_ce); _bgg != nil {
					return _bgg
				}
			case _g.Name{Space: "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg/package\u002f2006\u002fmetadata\u002fcore\u002dproperties", Local: "version"}:
				_ed.Version = new(string)
				if _gf := d.DecodeElement(_ed.Version, &_ce); _gf != nil {
					return _gf
				}
			default:
				_bc.Log("skipping unsupported\u0020element\u0020on\u0020CT_CoreProperties\u0020\u0025v", _ce.Name)
				if _fae := d.Skip(); _fae != nil {
					return _fae
				}
			}
		case _g.EndElement:
			break _ff
		case _g.CharData:
		}
	}
	return nil
}

type CT_Keywords struct {
	LangAttr *string
	Value    []*CT_Keyword
}

// Validate validates the CT_Keywords and its children
func (_bb *CT_Keywords) Validate() error { return _bb.ValidateWithPath("CT_Keywords") }

type CoreProperties struct{ CT_CoreProperties }

func NewCT_Keywords() *CT_Keywords { _fbd := &CT_Keywords{}; return _fbd }

// Validate validates the CT_Keyword and its children
func (_bff *CT_Keyword) Validate() error { return _bff.ValidateWithPath("CT_Keyword") }
func (_bfc *CT_Keyword) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	for _, _dcg := range start.Attr {
		if _dcg.Name.Space == "http:\u002f\u002fwww\u002ew3.org/XML\u002f1998/namespace" && _dcg.Name.Local == "lang" {
			_gc, _bgb := _dcg.Value, error(nil)
			if _bgb != nil {
				return _bgb
			}
			_bfc.LangAttr = &_gc
			continue
		}
	}
	for {
		_ad, _dgc := d.Token()
		if _dgc != nil {
			return _aa.Errorf("parsing\u0020CT_Keyword:\u0020%s", _dgc)
		}
		if _ab, _dbfdb := _ad.(_g.CharData); _dbfdb {
			_bfc.Content = string(_ab)
		}
		if _gfc, _cee := _ad.(_g.EndElement); _cee && _gfc.Name == start.Name {
			break
		}
	}
	return nil
}

// ValidateWithPath validates the CT_Keywords and its children, prefixing error messages with path
func (_cec *CT_Keywords) ValidateWithPath(path string) error {
	for _dgb, _bfca := range _cec.Value {
		if _gb := _bfca.ValidateWithPath(_aa.Sprintf("\u0025s\u002fValue\u005b\u0025d\u005d", path, _dgb)); _gb != nil {
			return _gb
		}
	}
	return nil
}
func NewCoreProperties() *CoreProperties {
	_fffa := &CoreProperties{}
	_fffa.CT_CoreProperties = *NewCT_CoreProperties()
	return _fffa
}

// ValidateWithPath validates the CoreProperties and its children, prefixing error messages with path
func (_fdg *CoreProperties) ValidateWithPath(path string) error {
	if _gfca := _fdg.CT_CoreProperties.ValidateWithPath(path); _gfca != nil {
		return _gfca
	}
	return nil
}
func (_edc *CoreProperties) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "xmlns"}, Value: "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg/package\u002f2006\u002fmetadata\u002fcore\u002dproperties"})
	start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "xmlns:cp"}, Value: "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg/package\u002f2006\u002fmetadata\u002fcore\u002dproperties"})
	start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "xmlns:dc"}, Value: "http:\u002f\u002fpurl\u002eorg/dc\u002felements\u002f1\u002e1\u002f"})
	start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "xmlns:dcterms"}, Value: "http:\u002f/purl\u002eorg\u002fdc\u002fterms/"})
//	start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "xmlns:xxml"}, Value: "http:\u002f\u002fwww\u002ew3.org/XML\u002f1998/namespace"})
	start.Name.Local = "cp:coreProperties"
	return _edc.CT_CoreProperties.MarshalXML(e, start)
}

type CT_CoreProperties struct {
	Category       *string
	ContentStatus  *string
	Created        *_bc.XSDAny
	Creator        *_bc.XSDAny
	Description    *_bc.XSDAny
	Identifier     *_bc.XSDAny
	Keywords       *CT_Keywords
	Language       *_bc.XSDAny
	LastModifiedBy *string
	LastPrinted    *_b.Time
	Modified       *_bc.XSDAny
	Revision       *string
	Subject        *_bc.XSDAny
	Title          *_bc.XSDAny
	Version        *string
}

// ValidateWithPath validates the CT_Keyword and its children, prefixing error messages with path
func (_cgb *CT_Keyword) ValidateWithPath(path string) error { return nil }
func NewCT_Keyword() *CT_Keyword                            { _fc := &CT_Keyword{}; return _fc }
func (_fce *CT_Keywords) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	for _, _gfg := range start.Attr {
		if _gfg.Name.Space == "http:\u002f\u002fwww\u002ew3.org/XML\u002f1998/namespace" && _gfg.Name.Local == "lang" {
			_aee, _abf := _gfg.Value, error(nil)
			if _abf != nil {
				return _abf
			}
			_fce.LangAttr = &_aee
			continue
		}
	}
_edf:
	for {
		_gab, _ca := d.Token()
		if _ca != nil {
			return _ca
		}
		switch _ebg := _gab.(type) {
		case _g.StartElement:
			switch _ebg.Name {
			case _g.Name{Space: "http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg/package\u002f2006\u002fmetadata\u002fcore\u002dproperties", Local: "value"}:
				_fde := NewCT_Keyword()
				if _agd := d.DecodeElement(_fde, &_ebg); _agd != nil {
					return _agd
				}
				_fce.Value = append(_fce.Value, _fde)
			default:
				_bc.Log("skipping\u0020unsupported\u0020element\u0020on\u0020CT_Keywords\u0020\u0025v", _ebg.Name)
				if _ebfc := d.Skip(); _ebfc != nil {
					return _ebfc
				}
			}
		case _g.EndElement:
			break _edf
		case _g.CharData:
		}
	}
	return nil
}
func init() {
	_bc.RegisterConstructor("http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg/package\u002f2006\u002fmetadata\u002fcore\u002dproperties", "CT_CoreProperties", NewCT_CoreProperties)
	_bc.RegisterConstructor("http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg/package\u002f2006\u002fmetadata\u002fcore\u002dproperties", "CT_Keywords", NewCT_Keywords)
	_bc.RegisterConstructor("http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg/package\u002f2006\u002fmetadata\u002fcore\u002dproperties", "CT_Keyword", NewCT_Keyword)
	_bc.RegisterConstructor("http:\u002f\u002fschemas\u002eopenxmlformats\u002eorg/package\u002f2006\u002fmetadata\u002fcore\u002dproperties", "coreProperties", NewCoreProperties)
}
