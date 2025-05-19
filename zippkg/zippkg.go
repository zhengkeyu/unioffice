package zippkg

import (
	_g "archive/zip"
	_a "bytes"
	_bc "encoding/xml"
	_c "fmt"
	_gd "io"
	_fb "path"
	_f "sort"
	_e "strings"
	_b "time"

	_ae "gitee.com/greatmusicians/unioffice"
	_fd "gitee.com/greatmusicians/unioffice/algo"
	_eg "gitee.com/greatmusicians/unioffice/common/tempstorage"
	_eb "gitee.com/greatmusicians/unioffice/schema/soo/pkg/relationships"
)

func (_deb SelfClosingWriter) Write(b []byte) (int, error) {
	_fbff := 0
	_gdec := 0
	for _af := 0; _af < len(b)-2; _af++ {
		if b[_af] == '>' && b[_af+1] == '<' && b[_af+2] == '/' {
			_ed := []byte{}
			_ede := _af
			for _edee := _af; _edee >= 0; _edee-- {
				if b[_edee] == ' ' {
					_ede = _edee
				} else if b[_edee] == '<' {
					_ed = b[_edee+1 : _ede]
					break
				}
			}
			_fbfb := []byte{}
			for _abc := _af + 3; _abc < len(b); _abc++ {
				if b[_abc] == '>' {
					_fbfb = b[_af+3 : _abc]
					break
				}
			}
			if !_a.Equal(_ed, _fbfb) {
				continue
			}
			_agd, _eca := _deb.W.Write(b[_fbff:_af])
			if _eca != nil {
				return _gdec + _agd, _eca
			}
			_gdec += _agd
			_, _eca = _deb.W.Write(_bce)
			if _eca != nil {
				return _gdec, _eca
			}
			_gdec += 3
			for _bcb := _af + 2; _bcb < len(b) && b[_bcb] != '>'; _bcb++ {
				_gdec++
				_fbff = _bcb + 2
				_af = _fbff
			}
		}
	}
	_dad, _fg := _deb.W.Write(b[_fbff:])
	return _dad + _gdec, _fg
}
func (_fe *DecodeMap) IndexFor(path string) int          { return _fe._agf[path] }
func (_aeg *DecodeMap) RecordIndex(path string, idx int) { _aeg._agf[path] = idx }
func MarshalXMLByType(z *_g.Writer, dt _ae.DocType, typ string, v interface{}) error {
	_bgb := _ae.AbsoluteFilename(dt, typ, 0)
	return MarshalXML(z, _bgb, v)
}

// AddFileFromBytes takes a byte array and adds it at a given path to a zip file.
func AddFileFromBytes(z *_g.Writer, zipPath string, data []byte) error {
	_bf, _cb := z.Create(zipPath)
	if _cb != nil {
		return _c.Errorf("error creating \u0025s:\u0020\u0025s", zipPath, _cb)
	}
	_, _cb = _gd.Copy(_bf, _a.NewReader(data))
	return _cb
}

// Decode unmarshals the content of a *zip.File as XML to a given destination.
func Decode(f *_g.File, dest interface{}) error {
	_ecc, _ded := f.Open()
	if _ded != nil {
		return _c.Errorf("error\u0020reading\u0020\u0025s: \u0025s", f.Name, _ded)
	}
	defer _ecc.Close()
	_bg := _bc.NewDecoder(_ecc)
	if _gf := _bg.Decode(dest); _gf != nil {
		return _c.Errorf("error decoding \u0025s:\u0020\u0025s", f.Name, _gf)
	}
	if _aef, _aae := dest.(*_eb.Relationships); _aae {
		for _gbe, _ccg := range _aef.Relationship {
			switch _ccg.TypeAttr {
			case _ae.OfficeDocumentTypeStrict:
				_aef.Relationship[_gbe].TypeAttr = _ae.OfficeDocumentType
			case _ae.StylesTypeStrict:
				_aef.Relationship[_gbe].TypeAttr = _ae.StylesType
			case _ae.ThemeTypeStrict:
				_aef.Relationship[_gbe].TypeAttr = _ae.ThemeType
			case _ae.ControlTypeStrict:
				_aef.Relationship[_gbe].TypeAttr = _ae.ControlType
			case _ae.SettingsTypeStrict:
				_aef.Relationship[_gbe].TypeAttr = _ae.SettingsType
			case _ae.ImageTypeStrict:
				_aef.Relationship[_gbe].TypeAttr = _ae.ImageType
			case _ae.CommentsTypeStrict:
				_aef.Relationship[_gbe].TypeAttr = _ae.CommentsType
			case _ae.ThumbnailTypeStrict:
				_aef.Relationship[_gbe].TypeAttr = _ae.ThumbnailType
			case _ae.DrawingTypeStrict:
				_aef.Relationship[_gbe].TypeAttr = _ae.DrawingType
			case _ae.ChartTypeStrict:
				_aef.Relationship[_gbe].TypeAttr = _ae.ChartType
			case _ae.ExtendedPropertiesTypeStrict:
				_aef.Relationship[_gbe].TypeAttr = _ae.ExtendedPropertiesType
			case _ae.CustomXMLTypeStrict:
				_aef.Relationship[_gbe].TypeAttr = _ae.CustomXMLType
			case _ae.WorksheetTypeStrict:
				_aef.Relationship[_gbe].TypeAttr = _ae.WorksheetType
			case _ae.SharedStringsTypeStrict:
				_aef.Relationship[_gbe].TypeAttr = _ae.SharedStringsType
			case _ae.TableTypeStrict:
				_aef.Relationship[_gbe].TypeAttr = _ae.TableType
			case _ae.HeaderTypeStrict:
				_aef.Relationship[_gbe].TypeAttr = _ae.HeaderType
			case _ae.FooterTypeStrict:
				_aef.Relationship[_gbe].TypeAttr = _ae.FooterType
			case _ae.NumberingTypeStrict:
				_aef.Relationship[_gbe].TypeAttr = _ae.NumberingType
			case _ae.FontTableTypeStrict:
				_aef.Relationship[_gbe].TypeAttr = _ae.FontTableType
			case _ae.WebSettingsTypeStrict:
				_aef.Relationship[_gbe].TypeAttr = _ae.WebSettingsType
			case _ae.FootNotesTypeStrict:
				_aef.Relationship[_gbe].TypeAttr = _ae.FootNotesType
			case _ae.EndNotesTypeStrict:
				_aef.Relationship[_gbe].TypeAttr = _ae.EndNotesType
			case _ae.SlideTypeStrict:
				_aef.Relationship[_gbe].TypeAttr = _ae.SlideType
			case _ae.VMLDrawingTypeStrict:
				_aef.Relationship[_gbe].TypeAttr = _ae.VMLDrawingType
			}
		}
		_f.Slice(_aef.Relationship, func(_adg, _cgd int) bool {
			_dce := _aef.Relationship[_adg]
			_ggd := _aef.Relationship[_cgd]
			return _fd.NaturalLess(_dce.IdAttr, _ggd.IdAttr)
		})
	}
	return nil
}

// OnNewRelationshipFunc is called when a new relationship has been discovered.
//
// target is a resolved path that takes into account the location of the
// relationships file source and should be the path in the zip file.
//
// files are passed so non-XML files that can't be handled by AddTarget can be
// decoded directly (e.g. images)
//
// rel is the actual relationship so its target can be modified if the source
// target doesn't match where unioffice will write the file (e.g. read in
// 'xl/worksheets/MyWorksheet.xml' and we'll write out
// 'xl/worksheets/sheet1.xml')
type OnNewRelationshipFunc func(_aa *DecodeMap, _gg, _cf string, _ag []*_g.File, _ea *_eb.Relationship, _fbf Target) error

// SetOnNewRelationshipFunc sets the function to be called when a new
// relationship has been discovered.
func (_gge *DecodeMap) SetOnNewRelationshipFunc(fn OnNewRelationshipFunc) { _gge._cc = fn }

// DecodeMap is used to walk a tree of relationships, decoding files and passing
// control back to the document.
type DecodeMap struct {
	_eae map[string]Target
	_dc  map[*_eb.Relationships]string
	_fc  []Target
	_cc  OnNewRelationshipFunc
	_ga  map[string]struct{}
	_agf map[string]int
}

var _bce = []byte{'/', '>'}

// RelationsPathFor returns the relations path for a given filename.
func RelationsPathFor(path string) string {
	_gc := _e.Split(path, "\u002f")
	_ba := _e.Join(_gc[0:len(_gc)-1], "\u002f")
	_da := _gc[len(_gc)-1]
	_ba += "\u002f_rels\u002f"
	_da += "\u002erels"
	return _ba + _da
}

var _ccc = []byte{'\r', '\n'}

// Decode loops decoding targets registered with AddTarget and calling th
func (_cd *DecodeMap) Decode(files []*_g.File) error {
	_cg := 1
	for _cg > 0 {
		for len(_cd._fc) > 0 {
			_ecf := _cd._fc[len(_cd._fc)-1]
			_cd._fc = _cd._fc[0 : len(_cd._fc)-1]
			_ge := _ecf.Ifc.(*_eb.Relationships)
			for _, _ad := range _ge.Relationship {
				_ee, _ := _cd._dc[_ge]
				_cd._cc(_cd, _ee+_ad.TargetAttr, _ad.TypeAttr, files, _ad, _ecf)
			}
		}
		for _ebca, _gad := range files {
			if _gad == nil {
				continue
			}
			if _ege, _adf := _cd._eae[_gad.Name]; _adf {
				delete(_cd._eae, _gad.Name)
				if _gb := Decode(_gad, _ege.Ifc); _gb != nil {
					return _gb
				}
				files[_ebca] = nil
				if _gae, _gag := _ege.Ifc.(*_eb.Relationships); _gag {
					_cd._fc = append(_cd._fc, _ege)
					_bb, _ := _fb.Split(_fb.Clean(_gad.Name + "\u002f\u002e\u002e\u002f"))
					_cd._dc[_gae] = _bb
					_cg++
				}
			}
		}
		_cg--
	}
	return nil
}

// ExtractToDiskTmp extracts a zip file to a temporary file in a given path,
// returning the name of the file.
func ExtractToDiskTmp(f *_g.File, path string) (string, error) {
	_adfe, _cfa := _eg.TempFile(path, "zz")
	if _cfa != nil {
		return "", _cfa
	}
	defer _adfe.Close()
	_eac, _cfa := f.Open()
	if _cfa != nil {
		return "", _cfa
	}
	defer _eac.Close()
	_, _cfa = _gd.Copy(_adfe, _eac)
	if _cfa != nil {
		return "", _cfa
	}
	return _adfe.Name(), nil
}

type Target struct {
	Path  string
	Typ   string
	Ifc   interface{}
	Index uint32
}

// MarshalXML creates a file inside of a zip and marshals an object as xml, prefixing it
// with a standard XML header.
func MarshalXML(z *_g.Writer, filename string, v interface{}) error {
	_ab := &_g.FileHeader{}
	_ab.Method = _g.Deflate
	_ab.Name = filename
	_ab.SetModTime(_b.Now())
	_agb, _egc := z.CreateHeader(_ab)
	if _egc != nil {
		return _c.Errorf("creating\u0020\u0025s\u0020in\u0020zip:\u0020%s", filename, _egc)
	}
	_, _egc = _agb.Write([]byte(XMLHeader))
	if _egc != nil {
		return _c.Errorf("creating\u0020xml\u0020header\u0020to\u0020\u0025s: \u0025s", filename, _egc)
	}
	if _egc = _bc.NewEncoder(SelfClosingWriter{_agb}).Encode(v); _egc != nil {
		return _c.Errorf("marshaling\u0020\u0025s:\u0020\u0025s", filename, _egc)
	}
	_, _egc = _agb.Write(_ccc)
	return _egc
}

// AddTarget allows documents to register decode targets. Path is a path that
// will be found in the zip file and ifc is an XML element that the file will be
// unmarshaled to.  filePath is the absolute path to the target, ifc is the
// object to decode into, sourceFileType is the type of file that the reference
// was discovered in, and index is the index of the source file type.
func (_gde *DecodeMap) AddTarget(filePath string, ifc interface{}, sourceFileType string, idx uint32) bool {
	if _gde._eae == nil {
		_gde._eae = make(map[string]Target)
		_gde._dc = make(map[*_eb.Relationships]string)
		_gde._ga = make(map[string]struct{})
		_gde._agf = make(map[string]int)
	}
	_de := _fb.Clean(filePath)
	if _, _ebc := _gde._ga[_de]; _ebc {
		return false
	}
	_gde._ga[_de] = struct{}{}
	_gde._eae[_de] = Target{Path: filePath, Typ: sourceFileType, Ifc: ifc, Index: idx}
	return true
}
func MarshalXMLByTypeIndex(z *_g.Writer, dt _ae.DocType, typ string, idx int, v interface{}) error {
	_ca := _ae.AbsoluteFilename(dt, typ, idx)
	return MarshalXML(z, _ca, v)
}

const XMLHeader = "\u003c\u003fxml\u0020version\u003d\u00221\u002e0\"\u0020encoding=\u0022UTF\u002d8\u0022\u0020standalone\u003d\u0022yes\u0022\u003f\u003e" + "\u000a"

// SelfClosingWriter wraps a writer and replaces XML tags of the
// type <foo></foo> with <foo/>
type SelfClosingWriter struct{ W _gd.Writer }

// AddFileFromDisk reads a file from internal storage and adds it at a given path to a zip file.
// TODO: Rename to AddFileFromStorage in next major version release (v2).
// NOTE: If disk storage cannot be used, memory storage can be used instead by calling memstore.SetAsStorage().
func AddFileFromDisk(z *_g.Writer, zipPath, storagePath string) error {
	_ac, _fca := z.Create(zipPath)
	if _fca != nil {
		return _c.Errorf("error creating \u0025s:\u0020\u0025s", zipPath, _fca)
	}
	_fa, _fca := _eg.Open(storagePath)
	if _fca != nil {
		return _c.Errorf("error\u0020opening\u0020\u0025s: \u0025s", storagePath, _fca)
	}
	defer _fa.Close()
	_, _fca = _gd.Copy(_ac, _fa)
	return _fca
}
