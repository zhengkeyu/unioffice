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

package terms

import (
	_a "encoding/xml"
	_d "fmt"
	_dg "gitee.com/greatmusicians/unioffice"
	_f "gitee.com/greatmusicians/unioffice/schema/purl.org/dc/elements"
)

// Validate validates the TGN and its children
func (_dcfd *TGN) Validate() error { return _dcfd.ValidateWithPath("TGN") }
func (_abb *ISO639_2) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	start.Name.Local = "ISO639\u002d2"
	e.EncodeToken(start)
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}

// Validate validates the ISO3166 and its children
func (_fad *ISO3166) Validate() error { return _fad.ValidateWithPath("ISO3166") }

// Validate validates the MESH and its children
func (_fdf *MESH) Validate() error { return _fdf.ValidateWithPath("MESH") }

type LCSH struct{}

func (_dbd *ISO3166) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	for {
		_ccc, _aga := d.Token()
		if _aga != nil {
			return _d.Errorf("parsing\u0020ISO3166:\u0020\u0025s", _aga)
		}
		if _ccgf, _afb := _ccc.(_a.EndElement); _afb && _ccgf.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the W3CDTF and its children
func (_beae *W3CDTF) Validate() error { return _beae.ValidateWithPath("W3CDTF") }

// Validate validates the DCMIType and its children
func (_ce *DCMIType) Validate() error { return _ce.ValidateWithPath("DCMIType") }
func NewW3CDTF() *W3CDTF              { _fdg := &W3CDTF{}; return _fdg }

// Validate validates the Point and its children
func (_eefa *Point) Validate() error { return _eefa.ValidateWithPath("Point") }
func NewPoint() *Point               { _cdg := &Point{}; return _cdg }
func NewLCC() *LCC                   { _fbb := &LCC{}; return _fbb }
func (_gf *ElementOrRefinementContainer) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
_gg:
	for {
		_cd, _ee := d.Token()
		if _ee != nil {
			return _ee
		}
		switch _feb := _cd.(type) {
		case _a.StartElement:
			switch _feb.Name {
			case _a.Name{Space: "http:\u002f\u002fpurl\u002eorg/dc\u002felements\u002f1\u002e1\u002f", Local: "any"}:
				_ef := NewElementsAndRefinementsGroupChoice()
				if _eac := d.DecodeElement(&_ef.Any, &_feb); _eac != nil {
					return _eac
				}
				_gf.Choice = append(_gf.Choice, _ef)
			default:
				_dg.Log("skipping\u0020unsupported\u0020element on\u0020ElementOrRefinementContainer\u0020\u0025v", _feb.Name)
				if _ggf := d.Skip(); _ggf != nil {
					return _ggf
				}
			}
		case _a.EndElement:
			break _gg
		case _a.CharData:
		}
	}
	return nil
}

// ValidateWithPath validates the UDC and its children, prefixing error messages with path
func (_aef *UDC) ValidateWithPath(path string) error { return nil }
func (_aad *Period) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	for {
		_babb, _bgb := d.Token()
		if _bgb != nil {
			return _d.Errorf("parsing\u0020Period:\u0020\u0025s", _bgb)
		}
		if _ebf, _ced := _babb.(_a.EndElement); _ced && _ebf.Name == start.Name {
			break
		}
	}
	return nil
}
func NewElementOrRefinementContainer() *ElementOrRefinementContainer {
	_ea := &ElementOrRefinementContainer{}
	return _ea
}

type ElementsAndRefinementsGroupChoice struct{ Any []*_f.Any }

func (_babd *IMT) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	start.Name.Local = "IMT"
	e.EncodeToken(start)
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}
func NewUDC() *UDC { _dce := &UDC{}; return _dce }
func (_add *Point) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	start.Name.Local = "Point"
	e.EncodeToken(start)
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}
func NewPeriod() *Period { _aac := &Period{}; return _aac }

// Validate validates the LCC and its children
func (_ffe *LCC) Validate() error { return _ffe.ValidateWithPath("LCC") }

type IMT struct{}

func (_aeg *LCSH) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	start.Name.Local = "LCSH"
	e.EncodeToken(start)
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}
func (_fef *ElementOrRefinementContainer) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	start.Name.Local = "elementOrRefinementContainer"
	e.EncodeToken(start)
	if _fef.Choice != nil {
		for _, _cef := range _fef.Choice {
			_cef.MarshalXML(e, _a.StartElement{})
		}
	}
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}
func (_beg *ElementsAndRefinementsGroup) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
_ad:
	for {
		_ed, _aab := d.Token()
		if _aab != nil {
			return _aab
		}
		switch _baf := _ed.(type) {
		case _a.StartElement:
			switch _baf.Name {
			case _a.Name{Space: "http:\u002f\u002fpurl\u002eorg/dc\u002felements\u002f1\u002e1\u002f", Local: "any"}:
				_cec := NewElementsAndRefinementsGroupChoice()
				if _adc := d.DecodeElement(&_cec.Any, &_baf); _adc != nil {
					return _adc
				}
				_beg.Choice = append(_beg.Choice, _cec)
			default:
				_dg.Log("skipping\u0020unsupported\u0020element\u0020on\u0020ElementsAndRefinementsGroup\u0020\u0025v", _baf.Name)
				if _ae := d.Skip(); _ae != nil {
					return _ae
				}
			}
		case _a.EndElement:
			break _ad
		case _a.CharData:
		}
	}
	return nil
}
func (_fgb *RFC3066) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	for {
		_cfa, _aba := d.Token()
		if _aba != nil {
			return _d.Errorf("parsing\u0020RFC3066:\u0020\u0025s", _aba)
		}
		if _adff, _cdb := _cfa.(_a.EndElement); _cdb && _adff.Name == start.Name {
			break
		}
	}
	return nil
}

// ValidateWithPath validates the ElementsAndRefinementsGroup and its children, prefixing error messages with path
func (_cb *ElementsAndRefinementsGroup) ValidateWithPath(path string) error {
	for _eae, _ead := range _cb.Choice {
		if _fc := _ead.ValidateWithPath(_d.Sprintf("\u0025s\u002fChoice\u005b\u0025d\u005d", path, _eae)); _fc != nil {
			return _fc
		}
	}
	return nil
}
func (_cg *W3CDTF) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	for {
		_afg, _gfge := d.Token()
		if _gfge != nil {
			return _d.Errorf("parsing\u0020W3CDTF:\u0020\u0025s", _gfge)
		}
		if _dgg, _fee := _afg.(_a.EndElement); _fee && _dgg.Name == start.Name {
			break
		}
	}
	return nil
}
func NewRFC3066() *RFC3066 { _dfcd := &RFC3066{}; return _dfcd }

// Validate validates the ISO639_2 and its children
func (_efe *ISO639_2) Validate() error { return _efe.ValidateWithPath("ISO639_2") }

type ElementsAndRefinementsGroup struct {
	Choice []*ElementsAndRefinementsGroupChoice
}
type DCMIType struct{}

func (_gdd *LCC) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	start.Name.Local = "LCC"
	e.EncodeToken(start)
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}

// Validate validates the DDC and its children
func (_gbc *DDC) Validate() error { return _gbc.ValidateWithPath("DDC") }
func (_dca *DDC) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	start.Name.Local = "DDC"
	e.EncodeToken(start)
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}

// ValidateWithPath validates the ElementsAndRefinementsGroupChoice and its children, prefixing error messages with path
func (_cf *ElementsAndRefinementsGroupChoice) ValidateWithPath(path string) error {
	for _ade, _eba := range _cf.Any {
		if _ag := _eba.ValidateWithPath(_d.Sprintf("\u0025s\u002fAny\u005b\u0025d\u005d", path, _ade)); _ag != nil {
			return _ag
		}
	}
	return nil
}
func NewElementsAndRefinementsGroup() *ElementsAndRefinementsGroup {
	_fge := &ElementsAndRefinementsGroup{}
	return _fge
}
func (_de *Box) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	start.Name.Local = "Box"
	e.EncodeToken(start)
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}
func (_ebgg *Point) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	for {
		_eef, _dgc := d.Token()
		if _dgc != nil {
			return _d.Errorf("parsing\u0020Point:\u0020\u0025s", _dgc)
		}
		if _gfg, _fadb := _eef.(_a.EndElement); _fadb && _gfg.Name == start.Name {
			break
		}
	}
	return nil
}

type ISO639_2 struct{}

func (_fga *MESH) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	for {
		_aed, _gaa := d.Token()
		if _gaa != nil {
			return _d.Errorf("parsing\u0020MESH:\u0020\u0025s", _gaa)
		}
		if _ca, _dgf := _aed.(_a.EndElement); _dgf && _ca.Name == start.Name {
			break
		}
	}
	return nil
}

// ValidateWithPath validates the RFC1766 and its children, prefixing error messages with path
func (_gef *RFC1766) ValidateWithPath(path string) error { return nil }

// ValidateWithPath validates the Period and its children, prefixing error messages with path
func (_fbg *Period) ValidateWithPath(path string) error { return nil }
func (_efg *Period) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	start.Name.Local = "Period"
	e.EncodeToken(start)
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}
func (_ba *DCMIType) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	for {
		_bb, _gc := d.Token()
		if _gc != nil {
			return _d.Errorf("parsing\u0020DCMIType: \u0025s", _gc)
		}
		if _dgd, _dff := _bb.(_a.EndElement); _dff && _dgd.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the IMT and its children
func (_ebb *IMT) Validate() error { return _ebb.ValidateWithPath("IMT") }

type Point struct{}

func NewDCMIType() *DCMIType { _db := &DCMIType{}; return _db }

type MESH struct{}

func (_fde *LCSH) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	for {
		_efc, _ged := d.Token()
		if _ged != nil {
			return _d.Errorf("parsing\u0020LCSH:\u0020\u0025s", _ged)
		}
		if _egd, _ggfa := _efc.(_a.EndElement); _ggfa && _egd.Name == start.Name {
			break
		}
	}
	return nil
}

type RFC3066 struct{}

// Validate validates the LCSH and its children
func (_aeb *LCSH) Validate() error { return _aeb.ValidateWithPath("LCSH") }
func (_bcf *MESH) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	start.Name.Local = "MESH"
	e.EncodeToken(start)
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}
func NewLCSH() *LCSH { _gec := &LCSH{}; return _gec }

// ValidateWithPath validates the ElementOrRefinementContainer and its children, prefixing error messages with path
func (_bee *ElementOrRefinementContainer) ValidateWithPath(path string) error {
	for _fb, _bab := range _bee.Choice {
		if _dge := _bab.ValidateWithPath(_d.Sprintf("\u0025s\u002fChoice\u005b\u0025d\u005d", path, _fb)); _dge != nil {
			return _dge
		}
	}
	return nil
}

// ValidateWithPath validates the RFC3066 and its children, prefixing error messages with path
func (_ecc *RFC3066) ValidateWithPath(path string) error { return nil }
func NewMESH() *MESH                                     { _ggfb := &MESH{}; return _ggfb }
func (_fac *ISO639_2) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	for {
		_ga, _facc := d.Token()
		if _facc != nil {
			return _d.Errorf("parsing\u0020ISO639_2: \u0025s", _facc)
		}
		if _bac, _edc := _ga.(_a.EndElement); _edc && _bac.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the UDC and its children
func (_dfe *UDC) Validate() error { return _dfe.ValidateWithPath("UDC") }

type ElementOrRefinementContainer struct {
	Choice []*ElementsAndRefinementsGroupChoice
}
type Period struct{}

func (_bea *W3CDTF) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	start.Name.Local = "W3CDTF"
	e.EncodeToken(start)
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}
func (_ada *RFC3066) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	start.Name.Local = "RFC3066"
	e.EncodeToken(start)
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}
func NewRFC1766() *RFC1766 { _bec := &RFC1766{}; return _bec }

type RFC1766 struct{}

func (_eg *DDC) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	for {
		_aa, _gb := d.Token()
		if _gb != nil {
			return _d.Errorf("parsing\u0020DDC:\u0020\u0025s", _gb)
		}
		if _dd, _gd := _aa.(_a.EndElement); _gd && _dd.Name == start.Name {
			break
		}
	}
	return nil
}
func (_dde *ElementsAndRefinementsGroup) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	if _dde.Choice != nil {
		for _, _bf := range _dde.Choice {
			_bf.MarshalXML(e, _a.StartElement{})
		}
	}
	return nil
}

// ValidateWithPath validates the URI and its children, prefixing error messages with path
func (_babbf *URI) ValidateWithPath(path string) error { return nil }
func (_be *DCMIType) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	start.Name.Local = "DCMIType"
	e.EncodeToken(start)
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}
func (_cca *RFC1766) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	for {
		_fgf, _agd := d.Token()
		if _agd != nil {
			return _d.Errorf("parsing\u0020RFC1766:\u0020\u0025s", _agd)
		}
		if _dga, _abg := _fgf.(_a.EndElement); _abg && _dga.Name == start.Name {
			break
		}
	}
	return nil
}

// ValidateWithPath validates the DDC and its children, prefixing error messages with path
func (_fg *DDC) ValidateWithPath(path string) error { return nil }

// Validate validates the URI and its children
func (_fcbb *URI) Validate() error { return _fcbb.ValidateWithPath("URI") }

// ValidateWithPath validates the Point and its children, prefixing error messages with path
func (_adf *Point) ValidateWithPath(path string) error { return nil }

// ValidateWithPath validates the W3CDTF and its children, prefixing error messages with path
func (_dbb *W3CDTF) ValidateWithPath(path string) error { return nil }
func NewIMT() *IMT                                      { _af := &IMT{}; return _af }

type Box struct{}

func NewBox() *Box { _g := &Box{}; return _g }

type W3CDTF struct{}

// Validate validates the RFC1766 and its children
func (_fca *RFC1766) Validate() error { return _fca.ValidateWithPath("RFC1766") }
func (_dc *Box) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	for {
		_fe, _da := d.Token()
		if _da != nil {
			return _d.Errorf("parsing\u0020Box:\u0020\u0025s", _da)
		}
		if _c, _e := _fe.(_a.EndElement); _e && _c.Name == start.Name {
			break
		}
	}
	return nil
}
func NewTGN() *TGN { _eff := &TGN{}; return _eff }

// Validate validates the ElementsAndRefinementsGroup and its children
func (_dcf *ElementsAndRefinementsGroup) Validate() error {
	return _dcf.ValidateWithPath("ElementsAndRefinementsGroup")
}
func (_ffd *URI) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	start.Name.Local = "URI"
	e.EncodeToken(start)
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}
func NewDDC() *DDC { _cc := &DDC{}; return _cc }

// Validate validates the ElementsAndRefinementsGroupChoice and its children
func (_gga *ElementsAndRefinementsGroupChoice) Validate() error {
	return _gga.ValidateWithPath("ElementsAndRefinementsGroupChoice")
}
func (_ccd *TGN) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	start.Name.Local = "TGN"
	e.EncodeToken(start)
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}

// ValidateWithPath validates the ISO3166 and its children, prefixing error messages with path
func (_ebg *ISO3166) ValidateWithPath(path string) error { return nil }
func (_deg *URI) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	for {
		_cee, _ecf := d.Token()
		if _ecf != nil {
			return _d.Errorf("parsing\u0020URI:\u0020\u0025s", _ecf)
		}
		if _fadc, _afbb := _cee.(_a.EndElement); _afbb && _fadc.Name == start.Name {
			break
		}
	}
	return nil
}

// ValidateWithPath validates the LCC and its children, prefixing error messages with path
func (_fgeb *LCC) ValidateWithPath(path string) error { return nil }

// Validate validates the Box and its children
func (_ec *Box) Validate() error { return _ec.ValidateWithPath("Box") }
func (_bfd *UDC) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	for {
		_agf, _cbe := d.Token()
		if _cbe != nil {
			return _d.Errorf("parsing\u0020UDC:\u0020\u0025s", _cbe)
		}
		if _fcb, _ffea := _agf.(_a.EndElement); _ffea && _fcb.Name == start.Name {
			break
		}
	}
	return nil
}

// ValidateWithPath validates the DCMIType and its children, prefixing error messages with path
func (_eb *DCMIType) ValidateWithPath(path string) error { return nil }

// ValidateWithPath validates the Box and its children, prefixing error messages with path
func (_df *Box) ValidateWithPath(path string) error { return nil }

// ValidateWithPath validates the MESH and its children, prefixing error messages with path
func (_egb *MESH) ValidateWithPath(path string) error { return nil }

// ValidateWithPath validates the LCSH and its children, prefixing error messages with path
func (_dcb *LCSH) ValidateWithPath(path string) error { return nil }

// ValidateWithPath validates the TGN and its children, prefixing error messages with path
func (_geg *TGN) ValidateWithPath(path string) error { return nil }

// ValidateWithPath validates the ISO639_2 and its children, prefixing error messages with path
func (_bc *ISO639_2) ValidateWithPath(path string) error { return nil }
func NewISO3166() *ISO3166                               { _dfc := &ISO3166{}; return _dfc }
func (_agb *UDC) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	start.Name.Local = "UDC"
	e.EncodeToken(start)
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}

// Validate validates the RFC3066 and its children
func (_gac *RFC3066) Validate() error { return _gac.ValidateWithPath("RFC3066") }
func (_eeg *ElementsAndRefinementsGroupChoice) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	if _eeg.Any != nil {
		_gcg := _a.StartElement{Name: _a.Name{Local: "dc:any"}}
		for _, _dba := range _eeg.Any {
			e.EncodeElement(_dba, _gcg)
		}
	}
	return nil
}
func (_ccg *ElementsAndRefinementsGroupChoice) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
_aaed:
	for {
		_dec, _dbg := d.Token()
		if _dbg != nil {
			return _dbg
		}
		switch _ac := _dec.(type) {
		case _a.StartElement:
			switch _ac.Name {
			case _a.Name{Space: "http:\u002f\u002fpurl\u002eorg/dc\u002felements\u002f1\u002e1\u002f", Local: "any"}:
				_gde := _f.NewAny()
				if _cefb := d.DecodeElement(_gde, &_ac); _cefb != nil {
					return _cefb
				}
				_ccg.Any = append(_ccg.Any, _gde)
			default:
				_dg.Log("skipping\u0020unsupported element\u0020on\u0020ElementsAndRefinementsGroupChoice\u0020\u0025v", _ac.Name)
				if _dag := d.Skip(); _dag != nil {
					return _dag
				}
			}
		case _a.EndElement:
			break _aaed
		case _a.CharData:
		}
	}
	return nil
}
func NewElementsAndRefinementsGroupChoice() *ElementsAndRefinementsGroupChoice {
	_fa := &ElementsAndRefinementsGroupChoice{}
	return _fa
}

type ISO3166 struct{}

func (_bdg *RFC1766) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	start.Name.Local = "RFC1766"
	e.EncodeToken(start)
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}

// Validate validates the Period and its children
func (_dae *Period) Validate() error { return _dae.ValidateWithPath("Period") }

type TGN struct{}

// ValidateWithPath validates the IMT and its children, prefixing error messages with path
func (_dea *IMT) ValidateWithPath(path string) error { return nil }

type URI struct{}

func (_abe *TGN) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	for {
		_gfe, _abba := d.Token()
		if _abba != nil {
			return _d.Errorf("parsing\u0020TGN:\u0020\u0025s", _abba)
		}
		if _cde, _bcb := _gfe.(_a.EndElement); _bcb && _cde.Name == start.Name {
			break
		}
	}
	return nil
}
func (_gfd *ISO3166) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	start.Name.Local = "ISO3166"
	e.EncodeToken(start)
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}
func (_ab *IMT) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	for {
		_cdf, _bg := d.Token()
		if _bg != nil {
			return _d.Errorf("parsing\u0020IMT:\u0020\u0025s", _bg)
		}
		if _febf, _ge := _cdf.(_a.EndElement); _ge && _febf.Name == start.Name {
			break
		}
	}
	return nil
}

type UDC struct{}

func NewURI() *URI { _bda := &URI{}; return _bda }
func (_aaf *LCC) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	for {
		_ded, _ddf := d.Token()
		if _ddf != nil {
			return _d.Errorf("parsing\u0020LCC:\u0020\u0025s", _ddf)
		}
		if _agg, _geb := _ded.(_a.EndElement); _geb && _agg.Name == start.Name {
			break
		}
	}
	return nil
}

type LCC struct{}
type DDC struct{}

// Validate validates the ElementOrRefinementContainer and its children
func (_fd *ElementOrRefinementContainer) Validate() error {
	return _fd.ValidateWithPath("ElementOrRefinementContainer")
}
func NewISO639_2() *ISO639_2 { _cbg := &ISO639_2{}; return _cbg }
func init() {
	_dg.RegisterConstructor("http:\u002f/purl\u002eorg\u002fdc\u002fterms/", "LCSH", NewLCSH)
	_dg.RegisterConstructor("http:\u002f/purl\u002eorg\u002fdc\u002fterms/", "MESH", NewMESH)
	_dg.RegisterConstructor("http:\u002f/purl\u002eorg\u002fdc\u002fterms/", "DDC", NewDDC)
	_dg.RegisterConstructor("http:\u002f/purl\u002eorg\u002fdc\u002fterms/", "LCC", NewLCC)
	_dg.RegisterConstructor("http:\u002f/purl\u002eorg\u002fdc\u002fterms/", "UDC", NewUDC)
	_dg.RegisterConstructor("http:\u002f/purl\u002eorg\u002fdc\u002fterms/", "Period", NewPeriod)
	_dg.RegisterConstructor("http:\u002f/purl\u002eorg\u002fdc\u002fterms/", "W3CDTF", NewW3CDTF)
	_dg.RegisterConstructor("http:\u002f/purl\u002eorg\u002fdc\u002fterms/", "DCMIType", NewDCMIType)
	_dg.RegisterConstructor("http:\u002f/purl\u002eorg\u002fdc\u002fterms/", "IMT", NewIMT)
	_dg.RegisterConstructor("http:\u002f/purl\u002eorg\u002fdc\u002fterms/", "URI", NewURI)
	_dg.RegisterConstructor("http:\u002f/purl\u002eorg\u002fdc\u002fterms/", "ISO639\u002d2", NewISO639_2)
	_dg.RegisterConstructor("http:\u002f/purl\u002eorg\u002fdc\u002fterms/", "RFC1766", NewRFC1766)
	_dg.RegisterConstructor("http:\u002f/purl\u002eorg\u002fdc\u002fterms/", "RFC3066", NewRFC3066)
	_dg.RegisterConstructor("http:\u002f/purl\u002eorg\u002fdc\u002fterms/", "Point", NewPoint)
	_dg.RegisterConstructor("http:\u002f/purl\u002eorg\u002fdc\u002fterms/", "ISO3166", NewISO3166)
	_dg.RegisterConstructor("http:\u002f/purl\u002eorg\u002fdc\u002fterms/", "Box", NewBox)
	_dg.RegisterConstructor("http:\u002f/purl\u002eorg\u002fdc\u002fterms/", "TGN", NewTGN)
	_dg.RegisterConstructor("http:\u002f/purl\u002eorg\u002fdc\u002fterms/", "elementOrRefinementContainer", NewElementOrRefinementContainer)
	_dg.RegisterConstructor("http:\u002f/purl\u002eorg\u002fdc\u002fterms/", "elementsAndRefinementsGroup", NewElementsAndRefinementsGroup)
}
