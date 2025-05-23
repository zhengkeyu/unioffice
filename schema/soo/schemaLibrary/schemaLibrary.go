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

package schemaLibrary

import (
	_b "encoding/xml"
	_c "fmt"
	_a "gitee.com/greatmusicians/unioffice"
)

func NewSchemaLibrary() *SchemaLibrary {
	_ggb := &SchemaLibrary{}
	_ggb.CT_SchemaLibrary = *NewCT_SchemaLibrary()
	return _ggb
}
func NewCT_Schema() *CT_Schema { _d := &CT_Schema{}; return _d }

// Validate validates the CT_Schema and its children
func (_cb *CT_Schema) Validate() error { return _cb.ValidateWithPath("CT_Schema") }

// Validate validates the SchemaLibrary and its children
func (_db *SchemaLibrary) Validate() error { return _db.ValidateWithPath("SchemaLibrary") }
func (_af *CT_Schema) UnmarshalXML(d *_b.Decoder, start _b.StartElement) error {
	for _, _gd := range start.Attr {
		if _gd.Name.Local == "uri" {
			_ag, _cf := _gd.Value, error(nil)
			if _cf != nil {
				return _cf
			}
			_af.UriAttr = &_ag
			continue
		}
		if _gd.Name.Local == "manifestLocation" {
			_fc, _gc := _gd.Value, error(nil)
			if _gc != nil {
				return _gc
			}
			_af.ManifestLocationAttr = &_fc
			continue
		}
		if _gd.Name.Local == "schemaLocation" {
			_ef, _dg := _gd.Value, error(nil)
			if _dg != nil {
				return _dg
			}
			_af.SchemaLocationAttr = &_ef
			continue
		}
		if _gd.Name.Local == "schemaLanguage" {
			_gg, _cfa := _gd.Value, error(nil)
			if _cfa != nil {
				return _cfa
			}
			_af.SchemaLanguageAttr = &_gg
			continue
		}
	}
	for {
		_dge, _da := d.Token()
		if _da != nil {
			return _c.Errorf("parsing\u0020CT_Schema:\u0020\u0025s", _da)
		}
		if _fca, _ac := _dge.(_b.EndElement); _ac && _fca.Name == start.Name {
			break
		}
	}
	return nil
}
func NewCT_SchemaLibrary() *CT_SchemaLibrary { _cab := &CT_SchemaLibrary{}; return _cab }
func (_efa *CT_SchemaLibrary) UnmarshalXML(d *_b.Decoder, start _b.StartElement) error {
_ge:
	for {
		_cdf, _df := d.Token()
		if _df != nil {
			return _df
		}
		switch _cff := _cdf.(type) {
		case _b.StartElement:
			switch _cff.Name {
			case _b.Name{Space: "http:\u002f\u002fschemas.openxmlformats\u002eorg/schemaLibrary\u002f2006\u002fmain", Local: "schema"}:
				_cbc := NewCT_Schema()
				if _cabb := d.DecodeElement(_cbc, &_cff); _cabb != nil {
					return _cabb
				}
				_efa.Schema = append(_efa.Schema, _cbc)
			default:
				_a.Log("skipping\u0020unsupported\u0020element\u0020on\u0020CT_SchemaLibrary\u0020\u0025v", _cff.Name)
				if _bc := d.Skip(); _bc != nil {
					return _bc
				}
			}
		case _b.EndElement:
			break _ge
		case _b.CharData:
		}
	}
	return nil
}

type SchemaLibrary struct{ CT_SchemaLibrary }

func (_fbf *SchemaLibrary) MarshalXML(e *_b.Encoder, start _b.StartElement) error {
	start.Attr = append(start.Attr, _b.Attr{Name: _b.Name{Local: "xmlns"}, Value: "http:\u002f\u002fschemas.openxmlformats\u002eorg/schemaLibrary\u002f2006\u002fmain"})
	start.Attr = append(start.Attr, _b.Attr{Name: _b.Name{Local: "xmlns:ma"}, Value: "http:\u002f\u002fschemas.openxmlformats\u002eorg/schemaLibrary\u002f2006\u002fmain"})
//	start.Attr = append(start.Attr, _b.Attr{Name: _b.Name{Local: "xmlns:xxml"}, Value: "http:\u002f\u002fwww\u002ew3.org/XML\u002f1998/namespace"})
	start.Name.Local = "ma:schemaLibrary"
	return _fbf.CT_SchemaLibrary.MarshalXML(e, start)
}

// Validate validates the CT_SchemaLibrary and its children
func (_ea *CT_SchemaLibrary) Validate() error { return _ea.ValidateWithPath("CT_SchemaLibrary") }

// ValidateWithPath validates the CT_Schema and its children, prefixing error messages with path
func (_ca *CT_Schema) ValidateWithPath(path string) error { return nil }

// ValidateWithPath validates the SchemaLibrary and its children, prefixing error messages with path
func (_ggf *SchemaLibrary) ValidateWithPath(path string) error {
	if _bb := _ggf.CT_SchemaLibrary.ValidateWithPath(path); _bb != nil {
		return _bb
	}
	return nil
}

// ValidateWithPath validates the CT_SchemaLibrary and its children, prefixing error messages with path
func (_fb *CT_SchemaLibrary) ValidateWithPath(path string) error {
	for _ec, _aff := range _fb.Schema {
		if _aca := _aff.ValidateWithPath(_c.Sprintf("\u0025s\u002fSchema\u005b\u0025d\u005d", path, _ec)); _aca != nil {
			return _aca
		}
	}
	return nil
}
func (_cabd *SchemaLibrary) UnmarshalXML(d *_b.Decoder, start _b.StartElement) error {
	_cabd.CT_SchemaLibrary = *NewCT_SchemaLibrary()
_fba:
	for {
		_cfc, _ee := d.Token()
		if _ee != nil {
			return _ee
		}
		switch _gb := _cfc.(type) {
		case _b.StartElement:
			switch _gb.Name {
			case _b.Name{Space: "http:\u002f\u002fschemas.openxmlformats\u002eorg/schemaLibrary\u002f2006\u002fmain", Local: "schema"}:
				_eb := NewCT_Schema()
				if _cg := d.DecodeElement(_eb, &_gb); _cg != nil {
					return _cg
				}
				_cabd.Schema = append(_cabd.Schema, _eb)
			default:
				_a.Log("skipping unsupported element\u0020on\u0020SchemaLibrary \u0025v", _gb.Name)
				if _ga := d.Skip(); _ga != nil {
					return _ga
				}
			}
		case _b.EndElement:
			break _fba
		case _b.CharData:
		}
	}
	return nil
}
func (_g *CT_Schema) MarshalXML(e *_b.Encoder, start _b.StartElement) error {
	if _g.UriAttr != nil {
		start.Attr = append(start.Attr, _b.Attr{Name: _b.Name{Local: "ma:uri"}, Value: _c.Sprintf("\u0025v", *_g.UriAttr)})
	}
	if _g.ManifestLocationAttr != nil {
		start.Attr = append(start.Attr, _b.Attr{Name: _b.Name{Local: "ma:manifestLocation"}, Value: _c.Sprintf("\u0025v", *_g.ManifestLocationAttr)})
	}
	if _g.SchemaLocationAttr != nil {
		start.Attr = append(start.Attr, _b.Attr{Name: _b.Name{Local: "ma:schemaLocation"}, Value: _c.Sprintf("\u0025v", *_g.SchemaLocationAttr)})
	}
	if _g.SchemaLanguageAttr != nil {
		start.Attr = append(start.Attr, _b.Attr{Name: _b.Name{Local: "ma:schemaLanguage"}, Value: _c.Sprintf("\u0025v", *_g.SchemaLanguageAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(_b.EndElement{Name: start.Name})
	return nil
}

type CT_SchemaLibrary struct{ Schema []*CT_Schema }

func (_dc *CT_SchemaLibrary) MarshalXML(e *_b.Encoder, start _b.StartElement) error {
	e.EncodeToken(start)
	if _dc.Schema != nil {
		_cd := _b.StartElement{Name: _b.Name{Local: "ma:schema"}}
		for _, _efe := range _dc.Schema {
			e.EncodeElement(_efe, _cd)
		}
	}
	e.EncodeToken(_b.EndElement{Name: start.Name})
	return nil
}

type CT_Schema struct {
	UriAttr              *string
	ManifestLocationAttr *string
	SchemaLocationAttr   *string
	SchemaLanguageAttr   *string
}

func init() {
	_a.RegisterConstructor("http:\u002f\u002fschemas.openxmlformats\u002eorg/schemaLibrary\u002f2006\u002fmain", "CT_Schema", NewCT_Schema)
	_a.RegisterConstructor("http:\u002f\u002fschemas.openxmlformats\u002eorg/schemaLibrary\u002f2006\u002fmain", "CT_SchemaLibrary", NewCT_SchemaLibrary)
	_a.RegisterConstructor("http:\u002f\u002fschemas.openxmlformats\u002eorg/schemaLibrary\u002f2006\u002fmain", "schemaLibrary", NewSchemaLibrary)
}
