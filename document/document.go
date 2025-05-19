/*

Package document provides creation, reading, and writing of ECMA 376 Open
Office XML documents.

Example:

	doc := document.New()
	para := doc.AddParagraph()
	run := para.AddRun()
	run.SetText("foo")
	doc.SaveToFile("foo.docx")
*/
package document

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"log"
	manthrand "math/rand"
	"os"
	"path/filepath"
	"strings"
	"unicode"

	unioffice "gitee.com/greatmusicians/unioffice"
	"gitee.com/greatmusicians/unioffice/color"
	"gitee.com/greatmusicians/unioffice/common"
	"gitee.com/greatmusicians/unioffice/common/tempstorage"
	"gitee.com/greatmusicians/unioffice/measurement"
	"gitee.com/greatmusicians/unioffice/schema/schemas.microsoft.com/office/activeX"
	"gitee.com/greatmusicians/unioffice/schema/soo/dml"
	"gitee.com/greatmusicians/unioffice/schema/soo/dml/picture"
	"gitee.com/greatmusicians/unioffice/schema/soo/ofc/sharedTypes"
	"gitee.com/greatmusicians/unioffice/schema/soo/pkg/relationships"
	"gitee.com/greatmusicians/unioffice/schema/soo/wml"
	"gitee.com/greatmusicians/unioffice/zippkg"
)

func (p Paragraph) addEndBookmark(_cefd int64) *wml.CT_MarkupRange {
	_cgdd := wml.NewEG_PContent()
	p.WParagraph.EG_PContent = append(p.WParagraph.EG_PContent, _cgdd)
	_bafb := wml.NewEG_ContentRunContent()
	_bffag := wml.NewEG_RunLevelElts()
	_egeaa := wml.NewEG_RangeMarkupElements()
	_eae := wml.NewCT_MarkupRange()
	_eae.IdAttr = _cefd
	_egeaa.BookmarkEnd = _eae
	_cgdd.EG_ContentRunContent = append(_cgdd.EG_ContentRunContent, _bafb)
	_bafb.EG_RunLevelElts = append(_bafb.EG_RunLevelElts, _bffag)
	_bffag.EG_RangeMarkupElements = append(_bffag.EG_RangeMarkupElements, _egeaa)
	return _eae
}

// SetBottom sets the bottom border to a specified type, color and thickness.
func (c CellBorders) SetBottom(t wml.ST_Border, co color.Color, thickness measurement.Distance) {
	c.WBorders.Bottom = wml.NewCT_Border()
	setBorder(c.WBorders.Bottom, t, co, thickness)
}

// Spacing returns the paragraph spacing settings.
func (p ParagraphProperties) Spacing() ParagraphSpacing {
	if p.Properties.Spacing == nil {
		p.Properties.Spacing = wml.NewCT_Spacing()
	}
	return ParagraphSpacing{p.Properties.Spacing}
}

// SetDefaultValue sets the default value of a FormFieldTypeDropDown. For
// FormFieldTypeDropDown, the value must be one of the fields possible values.
func (f FormField) SetDefaultValue(v string) {
	if f.WData.DdList != nil {
		for _geed, _acd := range f.PossibleValues() {
			if _acd == v {
				f.WData.DdList.Default = wml.NewCT_DecimalNumber()
				f.WData.DdList.Default.ValAttr = int64(_geed)
				break
			}
		}
	}
}

// GetColor returns the color.Color object representing the run color.
func (r RunProperties) GetColor() color.Color {
	if _bafd := r.WProperties.Color; _bafd != nil {
		_eage := _bafd.ValAttr
		if _eage.ST_HexColorRGB != nil {
			return color.FromHex(*_eage.ST_HexColorRGB)
		}
	}
	return color.Color{}
}

// Table is a table within a document.
type Table struct {
	Document *Document
	WTable   *wml.CT_Tbl
}

// Append appends a document d0 to a document d1. All settings, headers and footers remain the same as in the document d0 if they exist there, otherwise they are taken from the d1.
func (d *Document) Append(d0 *Document) error {
	_fbcd, err := d0.Copy()
	if err != nil {
		return err
	}
	d.DocBase = d.DocBase.Append(_fbcd.DocBase)
	if _fbcd.Document.ConformanceAttr != sharedTypes.ST_ConformanceClassStrict {
		d.Document.ConformanceAttr = _fbcd.Document.ConformanceAttr
	}
	_gdc := d._fbb.X().Relationship
	_fda := _fbcd._fbb.X().Relationship
	_cdg := _fbcd.Document.Body
	_baad := map[string]string{}
	_gfeg := map[int64]int64{}
	_dbb := map[int64]int64{}
	for _, _fcgb := range _fda {
		_bcd := true
		_dcag := _fcgb.IdAttr
		_fcdg := _fcgb.TargetAttr
		_cddb := _fcgb.TypeAttr
		_cccae := _cddb == unioffice.ImageType
		_cge := _cddb == unioffice.HyperLinkType
		var _cfd string
		for _, _feae := range _gdc {
			if _feae.TypeAttr == _cddb && _feae.TargetAttr == _fcdg {
				_bcd = false
				_cfd = _feae.IdAttr
				break
			}
		}
		if _cccae {
			_dcge := "word\u002f" + _fcdg
			for _, _fgcf := range _fbcd.DocBase.Images {
				if _fgcf.Target() == _dcge {
					_gggf, _gdef := common.ImageFromStorage(_fgcf.Path())
					if _gdef != nil {
						return _gdef
					}
					_bgc, _gdef := d.AddImage(_gggf)
					if _gdef != nil {
						return _gdef
					}
					_cfd = _bgc.RelID()
					break
				}
			}
		} else if _bcd {
			if _cge {
				_gdfb := d._fbb.AddHyperlink(_fcdg)
				_cfd = common.Relationship(_gdfb).ID()
			} else {
				_cabe := d._fbb.AddRelationship(_fcdg, _cddb)
				_cfd = _cabe.X().IdAttr
			}
		}
		if _dcag != _cfd {
			_baad[_dcag] = _cfd
		}
	}
	if _cdg.SectPr != nil {
		for _, _fcdb := range _cdg.SectPr.EG_HdrFtrReferences {
			if _fcdb.HeaderReference != nil {
				if _dgag, _edfe := _baad[_fcdb.HeaderReference.IdAttr]; _edfe {
					_fcdb.HeaderReference.IdAttr = _dgag
					d._ddc = append(d._ddc, common.NewRelationships())
				}
			} else if _fcdb.FooterReference != nil {
				if _dfdd, _cgg := _baad[_fcdb.FooterReference.IdAttr]; _cgg {
					_fcdb.FooterReference.IdAttr = _dfdd
					d._fcbd = append(d._fcbd, common.NewRelationships())
				}
			}
		}
	}
	_fcab, _cfde := d.WEndnotes, _fbcd.WEndnotes
	if _fcab != nil {
		if _cfde != nil {
			if _fcab.Endnote != nil {
				if _cfde.Endnote != nil {
					_agdd := int64(len(_fcab.Endnote) + 1)
					for _, _bef := range _cfde.Endnote {
						_ddcg := _bef.IdAttr
						if _ddcg > 0 {
							_bef.IdAttr = _agdd
							_fcab.Endnote = append(_fcab.Endnote, _bef)
							_dbb[_ddcg] = _agdd
							_agdd++
						}
					}
				}
			} else {
				_fcab.Endnote = _cfde.Endnote
			}
		}
	} else if _cfde != nil {
		_fcab = _cfde
	}
	d.WEndnotes = _fcab
	_bgbgf, _bfff := d.WFootnotes, _fbcd.WFootnotes
	if _bgbgf != nil {
		if _bfff != nil {
			if _bgbgf.Footnote != nil {
				if _bfff.Footnote != nil {
					_fcbdc := int64(len(_bgbgf.Footnote) + 1)
					for _, _fcfd := range _bfff.Footnote {
						_fedf := _fcfd.IdAttr
						if _fedf > 0 {
							_fcfd.IdAttr = _fcbdc
							_bgbgf.Footnote = append(_bgbgf.Footnote, _fcfd)
							_gfeg[_fedf] = _fcbdc
							_fcbdc++
						}
					}
				}
			} else {
				_bgbgf.Footnote = _bfff.Footnote
			}
		}
	} else if _bfff != nil {
		_bgbgf = _bfff
	}
	d.WFootnotes = _bgbgf
	for _, _bege := range _cdg.EG_BlockLevelElts {
		for _, _bcfg := range _bege.EG_ContentBlockContent {
			for _, _cged := range _bcfg.P {
				_bfbf(_cged, _baad)
				_baf(_cged, _baad)
				_gfb(_cged, _gfeg, _dbb)
			}
			for _, _acbg := range _bcfg.Tbl {
				_efag(_acbg, _baad)
				_fdfb(_acbg, _baad)
				_ebcg(_acbg, _gfeg, _dbb)
			}
		}
	}
	d.Document.Body.EG_BlockLevelElts = append(d.Document.Body.EG_BlockLevelElts, _fbcd.Document.Body.EG_BlockLevelElts...)
	if d.Document.Body.SectPr == nil {
		d.Document.Body.SectPr = _fbcd.Document.Body.SectPr
	} else {
		var _dgfc, _fcbf bool
		for _, _cbeb := range d.Document.Body.SectPr.EG_HdrFtrReferences {
			if _cbeb.HeaderReference != nil {
				_dgfc = true
			} else if _cbeb.FooterReference != nil {
				_fcbf = true
			}
		}
		if !_dgfc {
			for _, _ccd := range _fbcd.Document.Body.SectPr.EG_HdrFtrReferences {
				if _ccd.HeaderReference != nil {
					d.Document.Body.SectPr.EG_HdrFtrReferences = append(d.Document.Body.SectPr.EG_HdrFtrReferences, _ccd)
					break
				}
			}
		}
		if !_fcbf {
			for _, _cffg := range _fbcd.Document.Body.SectPr.EG_HdrFtrReferences {
				if _cffg.FooterReference != nil {
					d.Document.Body.SectPr.EG_HdrFtrReferences = append(d.Document.Body.SectPr.EG_HdrFtrReferences, _cffg)
					break
				}
			}
		}
		if d.Document.Body.SectPr.Cols == nil && _fbcd.Document.Body.SectPr.Cols != nil {
			d.Document.Body.SectPr.Cols = _fbcd.Document.Body.SectPr.Cols
		}
	}
	_dbgd := d.Numbering.WNumbering
	_fce := _fbcd.Numbering.WNumbering
	if _dbgd != nil {
		if _fce != nil {
			_dbgd.NumPicBullet = append(_dbgd.NumPicBullet, _fce.NumPicBullet...)
			_dbgd.AbstractNum = append(_dbgd.AbstractNum, _fce.AbstractNum...)
			_dbgd.Num = append(_dbgd.Num, _fce.Num...)
		}
	} else if _fce != nil {
		_dbgd = _fce
	}
	d.Numbering.WNumbering = _dbgd
	if d.Styles.WStyles == nil && _fbcd.Styles.WStyles != nil {
		d.Styles.WStyles = _fbcd.Styles.WStyles
	}
	d.DTheme = append(d.DTheme, _fbcd.DTheme...)
	d.Ocx = append(d.Ocx, _fbcd.Ocx...)
	if len(d.WHeader) == 0 {
		d.WHeader = _fbcd.WHeader
	}
	if len(d.WFooter) == 0 {
		d.WFooter = _fbcd.WFooter
	}
	_bdd := d.WWebSettings
	_gadg := _fbcd.WWebSettings
	if _bdd != nil {
		if _gadg != nil {
			if _bdd.Divs != nil {
				if _gadg.Divs != nil {
					_bdd.Divs.Div = append(_bdd.Divs.Div, _gadg.Divs.Div...)
				}
			} else {
				_bdd.Divs = _gadg.Divs
			}
		}
		_bdd.Frameset = nil
	} else if _gadg != nil {
		_bdd = _gadg
		_bdd.Frameset = nil
	}
	d.WWebSettings = _bdd
	_gffb := d.WFonts
	_cebd := _fbcd.WFonts
	if _gffb != nil {
		if _cebd != nil {
			if _gffb.Font != nil {
				if _cebd.Font != nil {
					for _, _edbg := range _cebd.Font {
						_befe := true
						for _, _fdbb := range _gffb.Font {
							if _fdbb.NameAttr == _edbg.NameAttr {
								_befe = false
								break
							}
						}
						if _befe {
							_gffb.Font = append(_gffb.Font, _edbg)
						}
					}
				}
			} else {
				_gffb.Font = _cebd.Font
			}
		}
	} else if _cebd != nil {
		_gffb = _cebd
	}
	d.WFonts = _gffb
	return nil
}

var _dfff = [...]uint8{0, 20, 37, 58, 79}

// SetSize sets the size of the displayed image on the page.
func (i InlineDrawing) SetSize(w, h measurement.Distance) {
	i.WInlineDrawing.Extent.CxAttr = int64(float64(w*measurement.Pixel72) / measurement.EMU)
	i.WInlineDrawing.Extent.CyAttr = int64(float64(h*measurement.Pixel72) / measurement.EMU)
}

const _acga = "FormFieldTypeUnknownFormFieldTypeTextFormFieldTypeCheckBoxFormFieldTypeDropDown"

// SetRightPct sets the cell right margin
func (c CellMargins) SetRightPct(pct float64) {
	c.WMargins.Right = wml.NewCT_TblWidth()
	setTableMarginPercent(c.WMargins.Right, pct)
}

// Footnote is an individual footnote reference within the document.
type Footnote struct {
	Document  *Document
	WFootnote *wml.CT_FtnEdn
}

// SetColor sets the text color.
func (r RunProperties) SetColor(c color.Color) {
	r.WProperties.Color = wml.NewCT_Color()
	r.WProperties.Color.ValAttr.ST_HexColorRGB = c.AsRGBString()
}

// SetSemiHidden controls if the style is hidden in the UI.
func (s Style) SetSemiHidden(b bool) {
	if b {
		s.WStyle.SemiHidden = wml.NewCT_OnOff()
	} else {
		s.WStyle.SemiHidden = nil
	}
}

// SetUpdateFieldsOnOpen controls if fields are recalculated upon opening the
// document. This is useful for things like a table of contents as the library
// only adds the field code and relies on Word/LibreOffice to actually compute
// the content.
func (s Settings) SetUpdateFieldsOnOpen(b bool) {
	if !b {
		s.WSettings.UpdateFields = nil
	} else {
		s.WSettings.UpdateFields = wml.NewCT_OnOff()
	}
}

// GetTargetByRelId returns a target path with the associated relation ID in the
// document.
func (d *Document) GetTargetByRelId(idAttr string) string {
	return d._fbb.GetTargetByRelId(idAttr)
}

// AddField adds a field (automatically computed text) to the document.
func (r Run) AddField(code string) { r.AddFieldWithFormatting(code, "", true) }
func (d *Document) insertTable(p Paragraph, _agad bool) Table {
	_dcga := d.Document.Body
	if _dcga == nil {
		return d.AddTable()
	}
	_fbbg := p.X()
	for _fcfg, _efg := range _dcga.EG_BlockLevelElts {
		for _, _bff := range _efg.EG_ContentBlockContent {
			for _ead, _accd := range _bff.P {
				if _accd == _fbbg {
					_bab := wml.NewCT_Tbl()
					_ecae := wml.NewEG_BlockLevelElts()
					_abe := wml.NewEG_ContentBlockContent()
					_ecae.EG_ContentBlockContent = append(_ecae.EG_ContentBlockContent, _abe)
					_abe.Tbl = append(_abe.Tbl, _bab)
					_dcga.EG_BlockLevelElts = append(_dcga.EG_BlockLevelElts, nil)
					if _agad {
						copy(_dcga.EG_BlockLevelElts[_fcfg+1:], _dcga.EG_BlockLevelElts[_fcfg:])
						_dcga.EG_BlockLevelElts[_fcfg] = _ecae
						if _ead != 0 {
							_ebeg := wml.NewEG_BlockLevelElts()
							_cea := wml.NewEG_ContentBlockContent()
							_ebeg.EG_ContentBlockContent = append(_ebeg.EG_ContentBlockContent, _cea)
							_cea.P = _bff.P[:_ead]
							_dcga.EG_BlockLevelElts = append(_dcga.EG_BlockLevelElts, nil)
							copy(_dcga.EG_BlockLevelElts[_fcfg+1:], _dcga.EG_BlockLevelElts[_fcfg:])
							_dcga.EG_BlockLevelElts[_fcfg] = _ebeg
						}
						_bff.P = _bff.P[_ead:]
					} else {
						copy(_dcga.EG_BlockLevelElts[_fcfg+2:], _dcga.EG_BlockLevelElts[_fcfg+1:])
						_dcga.EG_BlockLevelElts[_fcfg+1] = _ecae
						if _ead != len(_bff.P)-1 {
							_cbbd := wml.NewEG_BlockLevelElts()
							_dbg := wml.NewEG_ContentBlockContent()
							_cbbd.EG_ContentBlockContent = append(_cbbd.EG_ContentBlockContent, _dbg)
							_dbg.P = _bff.P[_ead+1:]
							_dcga.EG_BlockLevelElts = append(_dcga.EG_BlockLevelElts, nil)
							copy(_dcga.EG_BlockLevelElts[_fcfg+3:], _dcga.EG_BlockLevelElts[_fcfg+2:])
							_dcga.EG_BlockLevelElts[_fcfg+2] = _cbbd
						}
						_bff.P = _bff.P[:_ead+1]
					}
					return Table{d, _bab}
				}
			}
			for _, _abg := range _bff.Tbl {
				_aeea := _ecd(_abg, _fbbg, _agad)
				if _aeea != nil {
					return Table{d, _aeea}
				}
			}
		}
	}
	return d.AddTable()
}

// MultiLevelType returns the multilevel type, or ST_MultiLevelTypeUnset if not set.
func (n NumberingDefinition) MultiLevelType() wml.ST_MultiLevelType {
	if n.WDefinition.MultiLevelType != nil {
		return n.WDefinition.MultiLevelType.ValAttr
	} else {
		return wml.ST_MultiLevelTypeUnset
	}
}

// RemoveRun removes a child run from a paragraph.
func (p Paragraph) RemoveRun(r Run) {
	for _, _cabdg := range p.WParagraph.EG_PContent {
		for _afbc, _bdfg := range _cabdg.EG_ContentRunContent {
			if _bdfg.R == r.WRun {
				copy(_cabdg.EG_ContentRunContent[_afbc:], _cabdg.EG_ContentRunContent[_afbc+1:])
				_cabdg.EG_ContentRunContent = _cabdg.EG_ContentRunContent[0 : len(_cabdg.EG_ContentRunContent)-1]
			}
			if _bdfg.Sdt != nil && _bdfg.Sdt.SdtContent != nil {
				for _ecgfg, _aafe := range _bdfg.Sdt.SdtContent.EG_ContentRunContent {
					if _aafe.R == r.WRun {
						copy(_bdfg.Sdt.SdtContent.EG_ContentRunContent[_ecgfg:], _bdfg.Sdt.SdtContent.EG_ContentRunContent[_ecgfg+1:])
						_bdfg.Sdt.SdtContent.EG_ContentRunContent = _bdfg.Sdt.SdtContent.EG_ContentRunContent[0 : len(_bdfg.Sdt.SdtContent.EG_ContentRunContent)-1]
					}
				}
			}
		}
	}
}

// GetImageByRelID returns an ImageRef with the associated relation ID in the
// document.
func (d *Document) GetImageByRelID(relID string) (common.ImageRef, bool) {
	for _, _gfaa := range d.Images {
		if _gfaa.RelID() == relID {
			return _gfaa, true
		}
	}
	return common.ImageRef{}, false
}

// GetFooter gets a section Footer for given type
func (s Section) GetFooter(t wml.ST_HdrFtr) (Footer, bool) {
	for _, _cabeg := range s.WSection.EG_HdrFtrReferences {
		if _cabeg.FooterReference.TypeAttr == t {
			for _, f := range s.Document.Footers() {
				_fead := s.Document._fbb.FindRIDForN(f.Index(), unioffice.FooterType)
				if _fead == _cabeg.FooterReference.IdAttr {
					return f, true
				}
			}
		}
	}
	return Footer{}, false
}

// BoldValue returns the precise nature of the bold setting (unset, off or on).
func (r RunProperties) BoldValue() OnOffValue { return getOnOffValue(r.WProperties.B) }

// DocText is an array of extracted text items which has some methods for representing extracted text.
type DocText struct{ Items []TextItem }

// Color returns the style's Color.
func (r RunProperties) Color() Color {
	if r.WProperties.Color == nil {
		r.WProperties.Color = wml.NewCT_Color()
	}
	return Color{r.WProperties.Color}
}

// AddTable adds a new table to the document body.
func (d *Document) AddTable() Table {
	_fcf := wml.NewEG_BlockLevelElts()
	d.Document.Body.EG_BlockLevelElts = append(d.Document.Body.EG_BlockLevelElts, _fcf)
	_fag := wml.NewEG_ContentBlockContent()
	_fcf.EG_ContentBlockContent = append(_fcf.EG_ContentBlockContent, _fag)
	_efe := wml.NewCT_Tbl()
	_fag.Tbl = append(_fag.Tbl, _efe)
	return Table{d, _efe}
}

func _baf(_adb *wml.CT_P, _efdd map[string]string) {
	for _, _bba := range _adb.EG_PContent {
		if _bba.Hyperlink != nil && _bba.Hyperlink.IdAttr != nil {
			if _afe, _cfeg := _efdd[*_bba.Hyperlink.IdAttr]; _cfeg {
				*_bba.Hyperlink.IdAttr = _afe
			}
		}
	}
}

// AddParagraph adds a paragraph to the endnote.
func (e Endnote) AddParagraph() Paragraph {
	_fdae := wml.NewEG_ContentBlockContent()
	_becb := len(e.WEndnote.EG_BlockLevelElts[0].EG_ContentBlockContent)
	e.WEndnote.EG_BlockLevelElts[0].EG_ContentBlockContent = append(e.WEndnote.EG_BlockLevelElts[0].EG_ContentBlockContent, _fdae)
	_gada := wml.NewCT_P()
	var _abag *wml.CT_String
	if _becb != 0 {
		_debe := len(e.WEndnote.EG_BlockLevelElts[0].EG_ContentBlockContent[_becb-1].P)
		_abag = e.WEndnote.EG_BlockLevelElts[0].EG_ContentBlockContent[_becb-1].P[_debe-1].PPr.PStyle
	} else {
		_abag = wml.NewCT_String()
		_abag.ValAttr = "Endnote"
	}
	_fdae.P = append(_fdae.P, _gada)
	p := Paragraph{e.Document, _gada}
	p.WParagraph.PPr = wml.NewCT_PPr()
	p.WParagraph.PPr.PStyle = _abag
	p.WParagraph.PPr.RPr = wml.NewCT_ParaRPr()
	return p
}

// AddParagraph adds a paragraph to the header.
func (h Header) AddParagraph() Paragraph {
	_ecbe := wml.NewEG_ContentBlockContent()
	h.WHeader.EG_ContentBlockContent = append(h.WHeader.EG_ContentBlockContent, _ecbe)
	_bcfb := wml.NewCT_P()
	_ecbe.P = append(_ecbe.P, _bcfb)
	return Paragraph{h.Document, _bcfb}
}

// AddFieldWithFormatting adds a field (automatically computed text) to the
// document with field specifc formatting.
func (r Run) AddFieldWithFormatting(code string, fmt string, isDirty bool) {
	_fcef := r.newIC()
	_fcef.FldChar = wml.NewCT_FldChar()
	_fcef.FldChar.FldCharTypeAttr = wml.ST_FldCharTypeBegin
	if isDirty {
		_fcef.FldChar.DirtyAttr = &sharedTypes.ST_OnOff{}
		_fcef.FldChar.DirtyAttr.Bool = unioffice.Bool(true)
	}
	_fcef = r.newIC()
	_fcef.InstrText = wml.NewCT_Text()
	if fmt != "" {
		_fcef.InstrText.Content = code + "\u0020" + fmt
	} else {
		_fcef.InstrText.Content = code
	}
	_fcef = r.newIC()
	_fcef.FldChar = wml.NewCT_FldChar()
	_fcef.FldChar.FldCharTypeAttr = wml.ST_FldCharTypeEnd
}

// Text returns the underlying tet in the run.
func (r Run) Text() string {
	if len(r.WRun.EG_RunInnerContent) == 0 {
		return ""
	}
	_gdaaa := bytes.Buffer{}
	for _, _cgceb := range r.WRun.EG_RunInnerContent {
		if _cgceb.T != nil {
			_gdaaa.WriteString(_cgceb.T.Content)
		}
		if _cgceb.Tab != nil {
			_gdaaa.WriteByte('\t')
		}
	}
	return _gdaaa.String()
}

// SetPageBreakBefore controls if there is a page break before this paragraph.
func (p ParagraphProperties) SetPageBreakBefore(b bool) {
	if !b {
		p.Properties.PageBreakBefore = nil
	} else {
		p.Properties.PageBreakBefore = wml.NewCT_OnOff()
	}
}

// SetShading controls the cell shading.
func (c CellProperties) SetShading(shd wml.ST_Shd, foreground, fill color.Color) {
	if shd == wml.ST_ShdUnset {
		c.WProperties.Shd = nil
	} else {
		c.WProperties.Shd = wml.NewCT_Shd()
		c.WProperties.Shd.ValAttr = shd
		c.WProperties.Shd.ColorAttr = &wml.ST_HexColor{}
		if foreground.IsAuto() {
			c.WProperties.Shd.ColorAttr.ST_HexColorAuto = wml.ST_HexColorAutoAuto
		} else {
			c.WProperties.Shd.ColorAttr.ST_HexColorRGB = foreground.AsRGBString()
		}
		c.WProperties.Shd.FillAttr = &wml.ST_HexColor{}
		if fill.IsAuto() {
			c.WProperties.Shd.FillAttr.ST_HexColorAuto = wml.ST_HexColorAutoAuto
		} else {
			c.WProperties.Shd.FillAttr.ST_HexColorRGB = fill.AsRGBString()
		}
	}
}

// TableProperties returns the table style properties.
func (s Style) TableProperties() TableStyleProperties {
	if s.WStyle.TblPr == nil {
		s.WStyle.TblPr = wml.NewCT_TblPrBase()
	}
	return TableStyleProperties{s.WStyle.TblPr}
}

// RunProperties returns the run style properties.
func (s Style) RunProperties() RunProperties {
	if s.WStyle.RPr == nil {
		s.WStyle.RPr = wml.NewCT_RPr()
	}
	return RunProperties{s.WStyle.RPr}
}

// SetVerticalBanding controls the conditional formatting for vertical banding.
func (t TableLook) SetVerticalBanding(on bool) {
	if !on {
		t.WTableLook.NoVBandAttr = &sharedTypes.ST_OnOff{}
		t.WTableLook.NoVBandAttr.ST_OnOff1 = sharedTypes.ST_OnOff1On
	} else {
		t.WTableLook.NoVBandAttr = &sharedTypes.ST_OnOff{}
		t.WTableLook.NoVBandAttr.ST_OnOff1 = sharedTypes.ST_OnOff1Off
	}
}

// HasEndnotes returns a bool based on the presence or abscence of endnotes within
// the document.
func (d *Document) HasEndnotes() bool { return d.WEndnotes != nil }

// SetRight sets the right border to a specified type, color and thickness.
func (t TableBorders) SetRight(b wml.ST_Border, c color.Color, thickness measurement.Distance) {
	t.WBorders.Right = wml.NewCT_Border()
	setBorder(t.WBorders.Right, b, c, thickness)
}

// SetSpacing sets the spacing that comes before and after the paragraph.
// Deprecated: See Spacing() instead which allows finer control.
func (p ParagraphProperties) SetSpacing(before, after measurement.Distance) {
	if p.Properties.Spacing == nil {
		p.Properties.Spacing = wml.NewCT_Spacing()
	}
	p.Properties.Spacing.BeforeAttr = &sharedTypes.ST_TwipsMeasure{}
	p.Properties.Spacing.BeforeAttr.ST_UnsignedDecimalNumber = unioffice.Uint64(uint64(before / measurement.Twips))
	p.Properties.Spacing.AfterAttr = &sharedTypes.ST_TwipsMeasure{}
	p.Properties.Spacing.AfterAttr.ST_UnsignedDecimalNumber = unioffice.Uint64(uint64(after / measurement.Twips))
}

// AnchoredDrawing is an absolutely positioned image within a document page.
type AnchoredDrawing struct {
	Document         *Document
	WAnchoredDrawing *wml.WdAnchor
}

// RemoveParagraph removes a paragraph from the footnote.
func (f Footnote) RemoveParagraph(p Paragraph) {
	for _, _gabc := range f.content() {
		for _abf, _cccd := range _gabc.P {
			if _cccd == p.WParagraph {
				copy(_gabc.P[_abf:], _gabc.P[_abf+1:])
				_gabc.P = _gabc.P[0 : len(_gabc.P)-1]
				return
			}
		}
	}
}

// SetWidthPercent sets the table to a width percentage.
func (t TableProperties) SetWidthPercent(pct float64) {
	t.WProperties.TblW = wml.NewCT_TblWidth()
	t.WProperties.TblW.TypeAttr = wml.ST_TblWidthPct
	t.WProperties.TblW.WAttr = &wml.ST_MeasurementOrPercent{}
	t.WProperties.TblW.WAttr.ST_DecimalNumberOrPercent = &wml.ST_DecimalNumberOrPercent{}
	t.WProperties.TblW.WAttr.ST_DecimalNumberOrPercent.ST_UnqualifiedPercentage = unioffice.Int64(int64(pct * 50))
}

// SetCSTheme sets the font complex script theme.
func (f Fonts) SetCSTheme(t wml.ST_Theme) { f.WFonts.CsthemeAttr = t }

// X returns the inner wrapped XML type.
func (h Header) X() *wml.Hdr { return h.WHeader }

// SetUnhideWhenUsed controls if a semi hidden style becomes visible when used.
func (s Style) SetUnhideWhenUsed(b bool) {
	if b {
		s.WStyle.UnhideWhenUsed = wml.NewCT_OnOff()
	} else {
		s.WStyle.UnhideWhenUsed = nil
	}
}
func parseBookmarkList2(content *wml.EG_ContentBlockContent) []Bookmark {
	result := []Bookmark{}
	for _, _acgf := range content.P {
		for _, _agbc := range _acgf.EG_PContent {
			for _, _egfd := range _agbc.EG_ContentRunContent {
				for _, _dcgc := range _egfd.EG_RunLevelElts {
					for _, _ada := range _dcgc.EG_RangeMarkupElements {
						if _ada.BookmarkStart != nil {
							result = append(result, Bookmark{_ada.BookmarkStart})
						}
					}
				}
			}
		}
	}
	for _, _gafg := range content.EG_RunLevelElts {
		for _, _ggef := range _gafg.EG_RangeMarkupElements {
			if _ggef.BookmarkStart != nil {
				result = append(result, Bookmark{_ggef.BookmarkStart})
			}
		}
	}
	for _, _bbd := range content.Tbl {
		for _, _caefg := range _bbd.EG_ContentRowContent {
			for _, _ffcf := range _caefg.Tr {
				for _, _bfb := range _ffcf.EG_ContentCellContent {
					for _, _fddb := range _bfb.Tc {
						for _, _adca := range _fddb.EG_BlockLevelElts {
							for _, _fbce := range _adca.EG_ContentBlockContent {
								result = append(result, parseBookmarkList2(_fbce)...)
							}
						}
					}
				}
			}
		}
	}
	return result
}

// Name returns the name of the bookmark whcih is the document unique ID that
// identifies the bookmark.
func (b Bookmark) Name() string { return b.WBookmark.NameAttr }

// StructuredDocumentTag are a tagged bit of content in a document.
type StructuredDocumentTag struct {
	Document     *Document
	WTaggedBlock *wml.CT_SdtBlock
}

func (p Paragraph) addEndFldChar() *wml.CT_FldChar {
	_dgffd := p.addFldChar()
	_dgffd.FldCharTypeAttr = wml.ST_FldCharTypeEnd
	return _dgffd
}

// Underline returns the type of paragraph underline.
func (p ParagraphProperties) Underline() wml.ST_Underline {
	if _bgeg := p.Properties.RPr.U; _bgeg != nil {
		return _bgeg.ValAttr
	}
	return 0
}

// SetNumberingDefinition sets the numbering definition ID via a NumberingDefinition
// defined in numbering.xml
func (p Paragraph) SetNumberingDefinition(nd NumberingDefinition) {
	p.ensurePPr()
	if p.WParagraph.PPr.NumPr == nil {
		p.WParagraph.PPr.NumPr = wml.NewCT_NumPr()
	}
	_aaca := wml.NewCT_DecimalNumber()
	_gagf := int64(-1)
	for _, _ffed := range p.Document.Numbering.WNumbering.Num {
		if _ffed.AbstractNumId != nil && _ffed.AbstractNumId.ValAttr == nd.AbstractNumberID() {
			_gagf = _ffed.NumIdAttr
		}
	}
	if _gagf == -1 {
		_gcedg := wml.NewCT_Num()
		p.Document.Numbering.WNumbering.Num = append(p.Document.Numbering.WNumbering.Num, _gcedg)
		_gcedg.NumIdAttr = int64(len(p.Document.Numbering.WNumbering.Num))
		_gcedg.AbstractNumId = wml.NewCT_DecimalNumber()
		_gcedg.AbstractNumId.ValAttr = nd.AbstractNumberID()
	}
	_aaca.ValAttr = _gagf
	p.WParagraph.PPr.NumPr.NumId = _aaca
}

// AddDrawingAnchored adds an anchored (floating) drawing from an ImageRef.
func (r Run) AddDrawingAnchored(img common.ImageRef) (AnchoredDrawing, error) {
	_cgaaf := r.newIC()
	_cgaaf.Drawing = wml.NewCT_Drawing()
	_bbdg := wml.NewWdAnchor()
	_dddag := AnchoredDrawing{r.Document, _bbdg}
	_bbdg.SimplePosAttr = unioffice.Bool(false)
	_bbdg.AllowOverlapAttr = true
	_bbdg.CNvGraphicFramePr = dml.NewCT_NonVisualGraphicFrameProperties()
	_cgaaf.Drawing.Anchor = append(_cgaaf.Drawing.Anchor, _bbdg)
	_bbdg.Graphic = dml.NewGraphic()
	_bbdg.Graphic.GraphicData = dml.NewCT_GraphicalObjectData()
	_bbdg.Graphic.GraphicData.UriAttr = "http:\u002f/schemas.openxmlformats\u002eorg\u002fdrawingml\u002f2006\u002fpicture"
	_bbdg.SimplePos.XAttr.ST_CoordinateUnqualified = unioffice.Int64(0)
	_bbdg.SimplePos.YAttr.ST_CoordinateUnqualified = unioffice.Int64(0)
	_bbdg.PositionH.RelativeFromAttr = wml.WdST_RelFromHPage
	_bbdg.PositionH.Choice = &wml.WdCT_PosHChoice{}
	_bbdg.PositionH.Choice.PosOffset = unioffice.Int32(0)
	_bbdg.PositionV.RelativeFromAttr = wml.WdST_RelFromVPage
	_bbdg.PositionV.Choice = &wml.WdCT_PosVChoice{}
	_bbdg.PositionV.Choice.PosOffset = unioffice.Int32(0)
	_bbdg.Extent.CxAttr = int64(float64(img.Size().X*measurement.Pixel72) / measurement.EMU)
	_bbdg.Extent.CyAttr = int64(float64(img.Size().Y*measurement.Pixel72) / measurement.EMU)
	_bbdg.Choice = &wml.WdEG_WrapTypeChoice{}
	_bbdg.Choice.WrapSquare = wml.NewWdCT_WrapSquare()
	_bbdg.Choice.WrapSquare.WrapTextAttr = wml.WdST_WrapTextBothSides
	_fcfc := 0x7FFFFFFF & manthrand.Uint32()
	_bbdg.DocPr.IdAttr = _fcfc
	_dggb := picture.NewPic()
	_dggb.NvPicPr.CNvPr.IdAttr = _fcfc
	_deec := img.RelID()
	if _deec == "" {
		return _dddag, errors.New("couldn\u0027t\u0020find\u0020reference\u0020to\u0020image\u0020within\u0020document\u0020relations")
	}
	_bbdg.Graphic.GraphicData.Any = append(_bbdg.Graphic.GraphicData.Any, _dggb)
	_dggb.BlipFill = dml.NewCT_BlipFillProperties()
	_dggb.BlipFill.Blip = dml.NewCT_Blip()
	_dggb.BlipFill.Blip.EmbedAttr = &_deec
	_dggb.BlipFill.Stretch = dml.NewCT_StretchInfoProperties()
	_dggb.BlipFill.Stretch.FillRect = dml.NewCT_RelativeRect()
	_dggb.SpPr = dml.NewCT_ShapeProperties()
	_dggb.SpPr.Xfrm = dml.NewCT_Transform2D()
	_dggb.SpPr.Xfrm.Off = dml.NewCT_Point2D()
	_dggb.SpPr.Xfrm.Off.XAttr.ST_CoordinateUnqualified = unioffice.Int64(0)
	_dggb.SpPr.Xfrm.Off.YAttr.ST_CoordinateUnqualified = unioffice.Int64(0)
	_dggb.SpPr.Xfrm.Ext = dml.NewCT_PositiveSize2D()
	_dggb.SpPr.Xfrm.Ext.CxAttr = int64(img.Size().X * measurement.Point)
	_dggb.SpPr.Xfrm.Ext.CyAttr = int64(img.Size().Y * measurement.Point)
	_dggb.SpPr.PrstGeom = dml.NewCT_PresetGeometry2D()
	_dggb.SpPr.PrstGeom.PrstAttr = dml.ST_ShapeTypeRect
	return _dddag, nil
}

// OpenTemplate opens a document, removing all content so it can be used as a
// template.  Since Word removes unused styles from a document upon save, to
// create a template in Word add a paragraph with every style of interest.  When
// opened with OpenTemplate the document's styles will be available but the
// content will be gone.
func OpenTemplate(filename string) (*Document, error) {
	_gafe, _dfc := Open(filename)
	if _dfc != nil {
		return nil, _dfc
	}
	_gafe.Document.Body = wml.NewCT_Body()
	return _gafe, nil
}

// SetConformance sets conformance attribute of the document
// as one of these values from gitee.com/greatmusicians/unioffice/schema/soo/ofc/sharedTypes:
// ST_ConformanceClassUnset, ST_ConformanceClassStrict or ST_ConformanceClassTransitional.
func (d Document) SetConformance(conformanceAttr sharedTypes.ST_ConformanceClass) {
	d.Document.ConformanceAttr = conformanceAttr
}

// SetImprint sets the run to imprinted text.
func (r RunProperties) SetImprint(b bool) {
	if !b {
		r.WProperties.Imprint = nil
	} else {
		r.WProperties.Imprint = wml.NewCT_OnOff()
	}
}

// SetInsideVertical sets the interior vertical borders to a specified type, color and thickness.
func (t TableBorders) SetInsideVertical(b wml.ST_Border, c color.Color, thickness measurement.Distance) {
	t.WBorders.InsideV = wml.NewCT_Border()
	setBorder(t.WBorders.InsideV, b, c, thickness)
}

// Shadow returns true if paragraph shadow is on.
func (p ParagraphProperties) Shadow() bool { return checkAttributeSet(p.Properties.RPr.Shadow) }

// X returns the inner wrapped XML type.
func (t TableStyleProperties) X() *wml.CT_TblPrBase { return t.WProperties }

// SetBasedOn sets the style that this style is based on.
func (s Style) SetBasedOn(name string) {
	if name == "" {
		s.WStyle.BasedOn = nil
	} else {
		s.WStyle.BasedOn = wml.NewCT_String()
		s.WStyle.BasedOn.ValAttr = name
	}
}

// RStyle returns the name of character style.
// It is defined here http://officeopenxml.com/WPstyleCharStyles.php
func (p ParagraphProperties) RStyle() string {
	if p.Properties.RPr.RStyle != nil {
		return p.Properties.RPr.RStyle.ValAttr
	}
	return ""
}
func _cgff(_cfgda *dml.CT_Blip, _fdec map[string]string) {
	if _cfgda.EmbedAttr != nil {
		if _eeca, _ebcgc := _fdec[*_cfgda.EmbedAttr]; _ebcgc {
			*_cfgda.EmbedAttr = _eeca
		}
	}
}

// SetSize sets size attribute for a FormFieldTypeCheckBox in pt.
func (f FormField) SetSize(size uint64) {
	size *= 2
	if f.WData.CheckBox != nil {
		f.WData.CheckBox.Choice = wml.NewCT_FFCheckBoxChoice()
		f.WData.CheckBox.Choice.Size = wml.NewCT_HpsMeasure()
		f.WData.CheckBox.Choice.Size.ValAttr = wml.ST_HpsMeasure{ST_UnsignedDecimalNumber: &size}
	}
}

// Caps returns true if run font is capitalized.
func (r RunProperties) Caps() bool { return checkAttributeSet(r.WProperties.Caps) }

// CellMargins are the margins for an individual cell.
type CellMargins struct{ WMargins *wml.CT_TcMar }

// SetItalic sets the run to italic.
func (r RunProperties) SetItalic(b bool) {
	if !b {
		r.WProperties.I = nil
		r.WProperties.ICs = nil
	} else {
		r.WProperties.I = wml.NewCT_OnOff()
		r.WProperties.ICs = wml.NewCT_OnOff()
	}
}

// SetOrigin sets the origin of the image.  It defaults to ST_RelFromHPage and
// ST_RelFromVPage
func (a AnchoredDrawing) SetOrigin(h wml.WdST_RelFromH, v wml.WdST_RelFromV) {
	a.WAnchoredDrawing.PositionH.RelativeFromAttr = h
	a.WAnchoredDrawing.PositionV.RelativeFromAttr = v
}

// AddHeader creates a header associated with the document, but doesn't add it
// to the document for display.
func (d *Document) AddHeader() Header {
	header := wml.NewHdr()
	d.WHeader = append(d.WHeader, header)
	_gbb := fmt.Sprintf("header\u0025d\u002exml", len(d.WHeader))
	d._fbb.AddRelationship(_gbb, unioffice.HeaderType)
	d.ContentTypes.AddOverride("\u002fword\u002f"+_gbb, "application\u002fvnd.openxmlformats\u002dofficedocument\u002ewordprocessingml\u002eheader\u002bxml")
	d._ddc = append(d._ddc, common.NewRelationships())
	return Header{d, header}
}

// SetASCIITheme sets the font ASCII Theme.
func (f Fonts) SetASCIITheme(t wml.ST_Theme) { f.WFonts.AsciiThemeAttr = t }

// SetWidthPercent sets the cell to a width percentage.
func (c CellProperties) SetWidthPercent(pct float64) {
	c.WProperties.TcW = wml.NewCT_TblWidth()
	c.WProperties.TcW.TypeAttr = wml.ST_TblWidthPct
	c.WProperties.TcW.WAttr = &wml.ST_MeasurementOrPercent{}
	c.WProperties.TcW.WAttr.ST_DecimalNumberOrPercent = &wml.ST_DecimalNumberOrPercent{}
	c.WProperties.TcW.WAttr.ST_DecimalNumberOrPercent.ST_UnqualifiedPercentage = unioffice.Int64(int64(pct * 50))
}

// Paragraphs returns all of the paragraphs in the document body including tables.
func (d *Document) Paragraphs() []Paragraph {
	result := []Paragraph{}
	if d.Document.Body == nil {
		return nil
	}
	for _, _fgab := range d.Document.Body.EG_BlockLevelElts {
		for _, _ddd := range _fgab.EG_ContentBlockContent {
			for _, _dbc := range _ddd.P {
				result = append(result, Paragraph{d, _dbc})
			}
		}
	}
	for _, _cbf := range d.Tables() {
		for _, _cfbb := range _cbf.Rows() {
			for _, _beg := range _cfbb.Cells() {
				result = append(result, _beg.Paragraphs()...)
			}
		}
	}
	return result
}

// SetLastColumn controls the conditional formatting for the last column in a table.
func (t TableLook) SetLastColumn(on bool) {
	if !on {
		t.WTableLook.LastColumnAttr = &sharedTypes.ST_OnOff{}
		t.WTableLook.LastColumnAttr.ST_OnOff1 = sharedTypes.ST_OnOff1Off
	} else {
		t.WTableLook.LastColumnAttr = &sharedTypes.ST_OnOff{}
		t.WTableLook.LastColumnAttr.ST_OnOff1 = sharedTypes.ST_OnOff1On
	}
}

// AddLevel adds a new numbering level to a NumberingDefinition.
func (n NumberingDefinition) AddLevel() NumberingLevel {
	_feaf := wml.NewCT_Lvl()
	_feaf.Start = &wml.CT_DecimalNumber{ValAttr: 1}
	_feaf.IlvlAttr = int64(len(n.WDefinition.Lvl))
	n.WDefinition.Lvl = append(n.WDefinition.Lvl, _feaf)
	return NumberingLevel{_feaf}
}

func (p Paragraph) ensurePPr() {
	if p.WParagraph.PPr == nil {
		p.WParagraph.PPr = wml.NewCT_PPr()
	}
}

// Numbering is the document wide numbering styles contained in numbering.xml.
type Numbering struct{ WNumbering *wml.Numbering }

// ParagraphProperties returns the paragraph properties controlling text formatting within the table.
func (t TableConditionalFormatting) ParagraphProperties() ParagraphStyleProperties {
	if t.WFormat.PPr == nil {
		t.WFormat.PPr = wml.NewCT_PPrGeneral()
	}
	return ParagraphStyleProperties{t.WFormat.PPr}
}

// SetStyle sets the style of a paragraph.
func (p ParagraphProperties) SetStyle(s string) {
	if s == "" {
		p.Properties.PStyle = nil
	} else {
		p.Properties.PStyle = wml.NewCT_String()
		p.Properties.PStyle.ValAttr = s
	}
}
func (e Endnote) id() int64 { return e.WEndnote.IdAttr }

// ParagraphProperties returns the paragraph style properties.
func (s Style) ParagraphProperties() ParagraphStyleProperties {
	if s.WStyle.PPr == nil {
		s.WStyle.PPr = wml.NewCT_PPrGeneral()
	}
	return ParagraphStyleProperties{s.WStyle.PPr}
}

// StyleID returns the style ID.
func (s Style) StyleID() string {
	if s.WStyle.StyleIdAttr == nil {
		return ""
	}
	return *s.WStyle.StyleIdAttr
}

// SetValue sets the width value.
func (t TableWidth) SetValue(m measurement.Distance) {
	t.WWidth.WAttr = &wml.ST_MeasurementOrPercent{}
	t.WWidth.WAttr.ST_DecimalNumberOrPercent = &wml.ST_DecimalNumberOrPercent{}
	t.WWidth.WAttr.ST_DecimalNumberOrPercent.ST_UnqualifiedPercentage = unioffice.Int64(int64(m / measurement.Twips))
	t.WWidth.TypeAttr = wml.ST_TblWidthDxa
}

// SetAfterAuto controls if spacing after a paragraph is automatically determined.
func (p ParagraphSpacing) SetAfterAuto(b bool) {
	if b {
		p.WSpacing.AfterAutospacingAttr = &sharedTypes.ST_OnOff{}
		p.WSpacing.AfterAutospacingAttr.Bool = unioffice.Bool(true)
	} else {
		p.WSpacing.AfterAutospacingAttr = nil
	}
}

// Fonts returns the style's Fonts.
func (r RunProperties) Fonts() Fonts {
	if r.WProperties.RFonts == nil {
		r.WProperties.RFonts = wml.NewCT_Fonts()
	}
	return Fonts{r.WProperties.RFonts}
}

// SetHeader sets a section header.
func (s Section) SetHeader(h Header, t wml.ST_HdrFtr) {
	_eccge := wml.NewEG_HdrFtrReferences()
	s.WSection.EG_HdrFtrReferences = append(s.WSection.EG_HdrFtrReferences, _eccge)
	_eccge.HeaderReference = wml.NewCT_HdrFtrRef()
	_eccge.HeaderReference.TypeAttr = t
	_bfg := s.Document._fbb.FindRIDForN(h.Index(), unioffice.HeaderType)
	if _bfg == "" {
		log.Print("unable\u0020to\u0020determine\u0020header ID")
	}
	_eccge.HeaderReference.IdAttr = _bfg
}

// X returns the inner wrapped XML type.
func (p Paragraph) X() *wml.CT_P { return p.WParagraph }

// RemoveParagraph removes a paragraph from a footer.
func (h Header) RemoveParagraph(p Paragraph) {
	for _, _ebga := range h.WHeader.EG_ContentBlockContent {
		for _deg, _cddbf := range _ebga.P {
			if _cddbf == p.WParagraph {
				copy(_ebga.P[_deg:], _ebga.P[_deg+1:])
				_ebga.P = _ebga.P[0 : len(_ebga.P)-1]
				return
			}
		}
	}
}

// X returns the inner wrapped XML type.
func (c CellProperties) X() *wml.CT_TcPr { return c.WProperties }

// SetAll sets all of the borders to a given value.
func (c CellBorders) SetAll(t wml.ST_Border, co color.Color, thickness measurement.Distance) {
	c.SetBottom(t, co, thickness)
	c.SetLeft(t, co, thickness)
	c.SetRight(t, co, thickness)
	c.SetTop(t, co, thickness)
	c.SetInsideHorizontal(t, co, thickness)
	c.SetInsideVertical(t, co, thickness)
}

// NewNumbering constructs a new numbering.
func NewNumbering() Numbering { n := wml.NewNumbering(); return Numbering{n} }

// X returns the inner wrapped XML type.
func (_abaf Footnote) X() *wml.CT_FtnEdn { return _abaf.WFootnote }

func _ebcg(_aaeb *wml.CT_Tbl, _gcfc, _aagg map[int64]int64) {
	for _, _cbda := range _aaeb.EG_ContentRowContent {
		for _, _egcc := range _cbda.Tr {
			for _, _aecc := range _egcc.EG_ContentCellContent {
				for _, _cdf := range _aecc.Tc {
					for _, _baca := range _cdf.EG_BlockLevelElts {
						for _, _cce := range _baca.EG_ContentBlockContent {
							for _, _gfbf := range _cce.P {
								_gfb(_gfbf, _gcfc, _aagg)
							}
							for _, _cga := range _cce.Tbl {
								_ebcg(_cga, _gcfc, _aagg)
							}
						}
					}
				}
			}
		}
	}
}

// AddParagraph adds a paragraph to the footer.
func (f Footer) AddParagraph() Paragraph {
	_gece := wml.NewEG_ContentBlockContent()
	f.WFooter.EG_ContentBlockContent = append(f.WFooter.EG_ContentBlockContent, _gece)
	_eegc := wml.NewCT_P()
	_gece.P = append(_gece.P, _eegc)
	return Paragraph{f.Document, _eegc}
}

// StructuredDocumentTags returns the structured document tags in the document
// which are commonly used in document templates.
func (d *Document) StructuredDocumentTags() []StructuredDocumentTag {
	tagList := []StructuredDocumentTag{}
	for _, _fbe := range d.Document.Body.EG_BlockLevelElts {
		for _, _bbec := range _fbe.EG_ContentBlockContent {
			if _bbec.Sdt != nil {
				tagList = append(tagList, StructuredDocumentTag{d, _bbec.Sdt})
			}
		}
	}
	return tagList
}

// GetImage returns the ImageRef associated with an AnchoredDrawing.
func (a AnchoredDrawing) GetImage() (common.ImageRef, bool) {
	_ca := a.WAnchoredDrawing.Graphic.GraphicData.Any
	if len(_ca) > 0 {
		_ggaa, _dfb := _ca[0].(*picture.Pic)
		if _dfb {
			if _ggaa.BlipFill != nil && _ggaa.BlipFill.Blip != nil && _ggaa.BlipFill.Blip.EmbedAttr != nil {
				return a.Document.GetImageByRelID(*_ggaa.BlipFill.Blip.EmbedAttr)
			}
		}
	}
	return common.ImageRef{}, false
}

// SetPossibleValues sets possible values for a FormFieldTypeDropDown.
func (f FormField) SetPossibleValues(values []string) {
	if f.WData.DdList != nil {
		for _, _gegf := range values {
			_aggd := wml.NewCT_String()
			_aggd.ValAttr = _gegf
			f.WData.DdList.ListEntry = append(f.WData.DdList.ListEntry, _aggd)
		}
	}
}

// ItalicValue returns the precise nature of the italic setting (unset, off or on).
func (r RunProperties) ItalicValue() OnOffValue { return getOnOffValue(r.WProperties.I) }

// SizeMeasure returns font with its measure which can be mm, cm, in, pt, pc or pi.
func (p ParagraphProperties) SizeMeasure() string {
	if _cbca := p.Properties.RPr.Sz; _cbca != nil {
		_caag := _cbca.ValAttr
		if _caag.ST_PositiveUniversalMeasure != nil {
			return *_caag.ST_PositiveUniversalMeasure
		}
	}
	return ""
}

// Settings controls the document settings.
type Settings struct{ WSettings *wml.Settings }

// SetEndIndent controls the end indentation.
func (p ParagraphProperties) SetEndIndent(m measurement.Distance) {
	if p.Properties.Ind == nil {
		p.Properties.Ind = wml.NewCT_Ind()
	}
	if m == measurement.Zero {
		p.Properties.Ind.EndAttr = nil
	} else {
		p.Properties.Ind.EndAttr = &wml.ST_SignedTwipsMeasure{}
		p.Properties.Ind.EndAttr.Int64 = unioffice.Int64(int64(m / measurement.Twips))
	}
}

// ClearContent clears any content in the run (text, tabs, breaks, etc.)
func (r Run) ClearContent() { r.WRun.EG_RunInnerContent = nil }

// InsertParagraphBefore adds a new empty paragraph before the relativeTo
// paragraph.
func (d *Document) InsertParagraphBefore(relativeTo Paragraph) Paragraph {
	return d.insertParagraph(relativeTo, true)
}

// SetKeepOnOnePage controls if all lines in a paragraph are kept on the same
// page.
func (p ParagraphStyleProperties) SetKeepOnOnePage(b bool) {
	if !b {
		p.WProperties.KeepLines = nil
	} else {
		p.WProperties.KeepLines = wml.NewCT_OnOff()
	}
}

func parseTextItemList(content []*wml.EG_ContentBlockContent, tableInfo *TableInfo) []TextItem {
	result := []TextItem{}
	for _, _ccdd := range content {
		if _beed := _ccdd.Sdt; _beed != nil {
			if _egdc := _beed.SdtContent; _egdc != nil {
				result = append(result, _gecdc(_egdc.P, tableInfo, nil)...)
			}
		}
		result = append(result, _gecdc(_ccdd.P, tableInfo, nil)...)
		for _, _bea := range _ccdd.Tbl {
			for _bccd, _eee := range _bea.EG_ContentRowContent {
				for _, _dcgac := range _eee.Tr {
					for _dfag, _fddf := range _dcgac.EG_ContentCellContent {
						for _, _gbed := range _fddf.Tc {
							_fad := &TableInfo{Table: _bea, Row: _dcgac, Cell: _gbed, RowIndex: _bccd, ColIndex: _dfag}
							for _, _gdag := range _gbed.EG_BlockLevelElts {
								result = append(result, parseTextItemList(_gdag.EG_ContentBlockContent, _fad)...)
							}
						}
					}
				}
			}
		}
	}
	return result
}

// AddStyle adds a new empty style.
func (s Styles) AddStyle(styleID string, t wml.ST_StyleType, isDefault bool) Style {
	style := wml.NewCT_Style()
	style.TypeAttr = t
	if isDefault {
		style.DefaultAttr = &sharedTypes.ST_OnOff{}
		style.DefaultAttr.Bool = unioffice.Bool(isDefault)
	}
	style.StyleIdAttr = unioffice.String(styleID)
	s.WStyles.Style = append(s.WStyles.Style, style)
	return Style{style}
}

// Paragraphs returns the paragraphs defined in a header.
func (h Header) Paragraphs() []Paragraph {
	_gbcc := []Paragraph{}
	for _, _cdbb := range h.WHeader.EG_ContentBlockContent {
		for _, _cecf := range _cdbb.P {
			_gbcc = append(_gbcc, Paragraph{h.Document, _cecf})
		}
	}
	for _, _bfea := range h.Tables() {
		for _, _afcg := range _bfea.Rows() {
			for _, _gaef := range _afcg.Cells() {
				_gbcc = append(_gbcc, _gaef.Paragraphs()...)
			}
		}
	}
	return _gbcc
}

// SetLineSpacing sets the spacing between lines in a paragraph.
func (p ParagraphSpacing) SetLineSpacing(d measurement.Distance, rule wml.ST_LineSpacingRule) {
	if rule == wml.ST_LineSpacingRuleUnset {
		p.WSpacing.LineRuleAttr = wml.ST_LineSpacingRuleUnset
		p.WSpacing.LineAttr = nil
	} else {
		p.WSpacing.LineRuleAttr = rule
		p.WSpacing.LineAttr = &wml.ST_SignedTwipsMeasure{}
		p.WSpacing.LineAttr.Int64 = unioffice.Int64(int64(d / measurement.Twips))
	}
}

// Paragraphs returns the paragraphs within a structured document tag.
func (s StructuredDocumentTag) Paragraphs() []Paragraph {
	if s.WTaggedBlock.SdtContent == nil {
		return nil
	}
	_cbad := []Paragraph{}
	for _, _gfbc := range s.WTaggedBlock.SdtContent.P {
		_cbad = append(_cbad, Paragraph{s.Document, _gfbc})
	}
	return _cbad
}

// CellProperties returns the cell properties.
func (t TableConditionalFormatting) CellProperties() CellProperties {
	if t.WFormat.TcPr == nil {
		t.WFormat.TcPr = wml.NewCT_TcPr()
	}
	return CellProperties{t.WFormat.TcPr}
}

// X returns the inner wrapped XML type.
func (s Settings) X() *wml.Settings { return s.WSettings }

// SetHANSITheme sets the font H ANSI Theme.
func (f Fonts) SetHANSITheme(t wml.ST_Theme) { f.WFonts.HAnsiThemeAttr = t }

// SetPrimaryStyle marks the style as a primary style.
func (s Style) SetPrimaryStyle(b bool) {
	if b {
		s.WStyle.QFormat = wml.NewCT_OnOff()
	} else {
		s.WStyle.QFormat = nil
	}
}

// Properties returns the row properties.
func (r Row) Properties() RowProperties {
	if r.WRow.TrPr == nil {
		r.WRow.TrPr = wml.NewCT_TrPr()
	}
	return RowProperties{r.WRow.TrPr}
}

// Document is a text document that can be written out in the OOXML .docx
// format. It can be opened from a file on disk and modified, or created from
// scratch.
type Document struct {
	common.DocBase
	Document       *wml.Document
	Settings       Settings
	Numbering      Numbering
	Styles         Styles
	WHeader        []*wml.Hdr
	_ddc           []common.Relationships
	WFooter        []*wml.Ftr
	_fcbd          []common.Relationships
	_fbb           common.Relationships
	DTheme         []*dml.Theme
	WWebSettings   *wml.WebSettings
	WFonts         *wml.Fonts
	WEndnotes      *wml.Endnotes
	WFootnotes     *wml.Footnotes
	Ocx            []*activeX.Ocx
	UnknownMeaning string
}

// ComplexSizeMeasure returns font with its measure which can be mm, cm, in, pt, pc or pi.
func (r RunProperties) ComplexSizeMeasure() string {
	if _cgeb := r.WProperties.SzCs; _cgeb != nil {
		_facad := _cgeb.ValAttr
		if _facad.ST_PositiveUniversalMeasure != nil {
			return *_facad.ST_PositiveUniversalMeasure
		}
	}
	return ""
}

// X returns the internally wrapped *wml.CT_SectPr.
func (s Section) X() *wml.CT_SectPr { return s.WSection }

// New constructs an empty document that content can be added to.
func New() *Document {
	d := &Document{Document: wml.NewDocument()}
	d.ContentTypes = common.NewContentTypes()
	d.Document.Body = wml.NewCT_Body()
	d.Document.ConformanceAttr = sharedTypes.ST_ConformanceClassTransitional
	d._fbb = common.NewRelationships()
	d.AppProperties = common.NewAppProperties()
	d.CoreProperties = common.NewCoreProperties()
	d.ContentTypes.AddOverride("\u002fword\u002fdocument\u002exml", "application/vnd\u002eopenxmlformats\u002dofficedocument\u002ewordprocessingml\u002edocument\u002emain\u002bxml")
	d.Settings = NewSettings()
	d._fbb.AddRelationship("settings\u002exml", unioffice.SettingsType)
	d.ContentTypes.AddOverride("\u002fword\u002fsettings\u002exml", "application\u002fvnd\u002eopenxmlformats\u002dofficedocument\u002ewordprocessingml.settings\u002bxml")
	d.Rels = common.NewRelationships()
	d.Rels.AddRelationship(unioffice.RelativeFilename(unioffice.DocTypeDocument, "", unioffice.CorePropertiesType, 0), unioffice.CorePropertiesType)
	d.Rels.AddRelationship("docProps\u002fapp\u002exml", unioffice.ExtendedPropertiesType)
	d.Rels.AddRelationship("word\u002fdocument\u002exml", unioffice.OfficeDocumentType)
	d.Numbering = NewNumbering()
	d.Numbering.InitializeDefault()
	d.ContentTypes.AddOverride("\u002fword/numbering\u002exml", "application\u002fvnd\u002eopenxmlformats\u002dofficedocument\u002ewordprocessingml\u002enumbering\u002bxml")
	d._fbb.AddRelationship("numbering\u002exml", unioffice.NumberingType)
	d.Styles = NewStyles()
	d.Styles.InitializeDefault()
	d.ContentTypes.AddOverride("\u002fword\u002fstyles\u002exml", "application\u002fvnd.openxmlformats\u002dofficedocument\u002ewordprocessingml\u002estyles\u002bxml")
	d._fbb.AddRelationship("styles\u002exml", unioffice.StylesType)
	d.Document.Body = wml.NewCT_Body()
	return d
}

// SetEmboss sets the run to embossed text.
func (r RunProperties) SetEmboss(b bool) {
	if !b {
		r.WProperties.Emboss = nil
	} else {
		r.WProperties.Emboss = wml.NewCT_OnOff()
	}
}

// SetCellSpacing sets the cell spacing within a table.
func (t TableProperties) SetCellSpacing(m measurement.Distance) {
	t.WProperties.TblCellSpacing = wml.NewCT_TblWidth()
	t.WProperties.TblCellSpacing.TypeAttr = wml.ST_TblWidthDxa
	t.WProperties.TblCellSpacing.WAttr = &wml.ST_MeasurementOrPercent{}
	t.WProperties.TblCellSpacing.WAttr.ST_DecimalNumberOrPercent = &wml.ST_DecimalNumberOrPercent{}
	t.WProperties.TblCellSpacing.WAttr.ST_DecimalNumberOrPercent.ST_UnqualifiedPercentage = unioffice.Int64(int64(m / measurement.Dxa))
}

// GetHeader gets a section Header for given type t [ST_HdrFtrDefault, ST_HdrFtrEven, ST_HdrFtrFirst]
func (s Section) GetHeader(t wml.ST_HdrFtr) (Header, bool) {
	for _, _cgea := range s.WSection.EG_HdrFtrReferences {
		if _cgea.HeaderReference.TypeAttr == t {
			for _, _aabgd := range s.Document.Headers() {
				_bgcf := s.Document._fbb.FindRIDForN(_aabgd.Index(), unioffice.HeaderType)
				if _bgcf == _cgea.HeaderReference.IdAttr {
					return _aabgd, true
				}
			}
		}
	}
	return Header{}, false
}

// Strike returns true if paragraph is striked.
func (p ParagraphProperties) Strike() bool { return checkAttributeSet(p.Properties.RPr.Strike) }

// SetFirstColumn controls the conditional formatting for the first column in a table.
func (t TableLook) SetFirstColumn(on bool) {
	if !on {
		t.WTableLook.FirstColumnAttr = &sharedTypes.ST_OnOff{}
		t.WTableLook.FirstColumnAttr.ST_OnOff1 = sharedTypes.ST_OnOff1Off
	} else {
		t.WTableLook.FirstColumnAttr = &sharedTypes.ST_OnOff{}
		t.WTableLook.FirstColumnAttr.ST_OnOff1 = sharedTypes.ST_OnOff1On
	}
}

// SetAlignment sets the alignment of a table within the page.
func (t TableProperties) SetAlignment(align wml.ST_JcTable) {
	if align == wml.ST_JcTableUnset {
		t.WProperties.Jc = nil
	} else {
		t.WProperties.Jc = wml.NewCT_JcTable()
		t.WProperties.Jc.ValAttr = align
	}
}

// MergeFields returns the list of all mail merge fields found in the document.
func (d Document) MergeFields() []string {
	_cefa := map[string]struct{}{}
	for _, v := range d.mergeFields() {
		_cefa[v._cbbg] = struct{}{}
	}
	result := []string{}
	for k := range _cefa {
		result = append(result, k)
	}
	return result
}

// RStyle returns the name of character style.
// It is defined here http://officeopenxml.com/WPstyleCharStyles.php
func (r RunProperties) RStyle() string {
	if r.WProperties.RStyle != nil {
		return r.WProperties.RStyle.ValAttr
	}
	return ""
}

// RemoveFootnote removes a footnote from both the paragraph and the document
// the requested footnote must be anchored on the paragraph being referenced.
func (p Paragraph) RemoveFootnote(id int64) {
	_gbfga := p.Document.WFootnotes
	_gecde := 0
	_gbfga.CT_Footnotes.Footnote[_gecde] = nil
	_gbfga.CT_Footnotes.Footnote[_gecde] = _gbfga.CT_Footnotes.Footnote[len(_gbfga.CT_Footnotes.Footnote)-1]
	_gbfga.CT_Footnotes.Footnote = _gbfga.CT_Footnotes.Footnote[:len(_gbfga.CT_Footnotes.Footnote)-1]
	var fRun Run
	for _, _cedd := range p.Runs() {
		if _fegf, _bbecc := _cedd.IsFootnote(); _fegf {
			if _bbecc == id {
				fRun = _cedd
			}
		}
	}
	p.RemoveRun(fRun)
}

// CharacterSpacingMeasure returns paragraph characters spacing with its measure which can be mm, cm, in, pt, pc or pi.
func (r RunProperties) CharacterSpacingMeasure() string {
	if _dgba := r.WProperties.Spacing; _dgba != nil {
		_dafba := _dgba.ValAttr
		if _dafba.ST_UniversalMeasure != nil {
			return *_dafba.ST_UniversalMeasure
		}
	}
	return ""
}

// SetKeepOnOnePage controls if all lines in a paragraph are kept on the same
// page.
func (p ParagraphProperties) SetKeepOnOnePage(b bool) {
	if !b {
		p.Properties.KeepLines = nil
	} else {
		p.Properties.KeepLines = wml.NewCT_OnOff()
	}
}

// DrawingInfo is used for keep information about a drawing wrapping a textbox where the text is located.
type DrawingInfo struct {
	WDrawing *wml.CT_Drawing
	Width    int64
	Height   int64
}

// SetKerning sets the run's font kerning.
func (r RunProperties) SetKerning(size measurement.Distance) {
	r.WProperties.Kern = wml.NewCT_HpsMeasure()
	r.WProperties.Kern.ValAttr.ST_UnsignedDecimalNumber = unioffice.Uint64(uint64(size / measurement.HalfPoint))
}

// Endnote returns the endnote based on the ID; this can be used nicely with
// the run.IsEndnote() functionality.
func (d *Document) Endnote(id int64) Endnote {
	for _, _eaf := range d.Endnotes() {
		if _eaf.id() == id {
			return _eaf
		}
	}
	return Endnote{}
}

// SetInsideHorizontal sets the interior horizontal borders to a specified type, color and thickness.
func (t TableBorders) SetInsideHorizontal(b wml.ST_Border, c color.Color, thickness measurement.Distance) {
	t.WBorders.InsideH = wml.NewCT_Border()
	setBorder(t.WBorders.InsideH, b, c, thickness)
}

// RemoveEndnote removes a endnote from both the paragraph and the document
// the requested endnote must be anchored on the paragraph being referenced.
func (p Paragraph) RemoveEndnote(id int64) {
	_bacf := p.Document.WEndnotes
	_gabb := 0
	_bacf.CT_Endnotes.Endnote[_gabb] = nil
	_bacf.CT_Endnotes.Endnote[_gabb] = _bacf.CT_Endnotes.Endnote[len(_bacf.CT_Endnotes.Endnote)-1]
	_bacf.CT_Endnotes.Endnote = _bacf.CT_Endnotes.Endnote[:len(_bacf.CT_Endnotes.Endnote)-1]
	var _cagd Run
	for _, _bcgcd := range p.Runs() {
		if _ccbb, _cdbc := _bcgcd.IsEndnote(); _ccbb {
			if _cdbc == id {
				_cagd = _bcgcd
			}
		}
	}
	p.RemoveRun(_cagd)
}

func (s Styles) initializeStyleDefaults() {
	_becf := s.AddStyle("Normal", wml.ST_StyleTypeParagraph, true)
	_becf.SetName("Normal")
	_becf.SetPrimaryStyle(true)
	_dbad := s.AddStyle("DefaultParagraphFont", wml.ST_StyleTypeCharacter, true)
	_dbad.SetName("Default\u0020Paragraph\u0020Font")
	_dbad.SetUISortOrder(1)
	_dbad.SetSemiHidden(true)
	_dbad.SetUnhideWhenUsed(true)
	_egeg := s.AddStyle("TitleChar", wml.ST_StyleTypeCharacter, false)
	_egeg.SetName("Title\u0020Char")
	_egeg.SetBasedOn(_dbad.StyleID())
	_egeg.SetLinkedStyle("Title")
	_egeg.SetUISortOrder(10)
	_egeg.RunProperties().Fonts().SetASCIITheme(wml.ST_ThemeMajorAscii)
	_egeg.RunProperties().Fonts().SetEastAsiaTheme(wml.ST_ThemeMajorEastAsia)
	_egeg.RunProperties().Fonts().SetHANSITheme(wml.ST_ThemeMajorHAnsi)
	_egeg.RunProperties().Fonts().SetCSTheme(wml.ST_ThemeMajorBidi)
	_egeg.RunProperties().SetSize(28 * measurement.Point)
	_egeg.RunProperties().SetKerning(14 * measurement.Point)
	_egeg.RunProperties().SetCharacterSpacing(-10 * measurement.Twips)
	_adggb := s.AddStyle("Title", wml.ST_StyleTypeParagraph, false)
	_adggb.SetName("Title")
	_adggb.SetBasedOn(_becf.StyleID())
	_adggb.SetNextStyle(_becf.StyleID())
	_adggb.SetLinkedStyle(_egeg.StyleID())
	_adggb.SetUISortOrder(10)
	_adggb.SetPrimaryStyle(true)
	_adggb.ParagraphProperties().SetContextualSpacing(true)
	_adggb.RunProperties().Fonts().SetASCIITheme(wml.ST_ThemeMajorAscii)
	_adggb.RunProperties().Fonts().SetEastAsiaTheme(wml.ST_ThemeMajorEastAsia)
	_adggb.RunProperties().Fonts().SetHANSITheme(wml.ST_ThemeMajorHAnsi)
	_adggb.RunProperties().Fonts().SetCSTheme(wml.ST_ThemeMajorBidi)
	_adggb.RunProperties().SetSize(28 * measurement.Point)
	_adggb.RunProperties().SetKerning(14 * measurement.Point)
	_adggb.RunProperties().SetCharacterSpacing(-10 * measurement.Twips)
	_ageg := s.AddStyle("TableNormal", wml.ST_StyleTypeTable, false)
	_ageg.SetName("Normal\u0020Table")
	_ageg.SetUISortOrder(99)
	_ageg.SetSemiHidden(true)
	_ageg.SetUnhideWhenUsed(true)
	_ageg.X().TblPr = wml.NewCT_TblPrBase()
	_deeg := NewTableWidth()
	_ageg.X().TblPr.TblInd = _deeg.X()
	_deeg.SetValue(0 * measurement.Dxa)
	_ageg.X().TblPr.TblCellMar = wml.NewCT_TblCellMar()
	_deeg = NewTableWidth()
	_ageg.X().TblPr.TblCellMar.Top = _deeg.X()
	_deeg.SetValue(0 * measurement.Dxa)
	_deeg = NewTableWidth()
	_ageg.X().TblPr.TblCellMar.Bottom = _deeg.X()
	_deeg.SetValue(0 * measurement.Dxa)
	_deeg = NewTableWidth()
	_ageg.X().TblPr.TblCellMar.Left = _deeg.X()
	_deeg.SetValue(108 * measurement.Dxa)
	_deeg = NewTableWidth()
	_ageg.X().TblPr.TblCellMar.Right = _deeg.X()
	_deeg.SetValue(108 * measurement.Dxa)
	_dadb := s.AddStyle("NoList", wml.ST_StyleTypeNumbering, false)
	_dadb.SetName("No\u0020List")
	_dadb.SetUISortOrder(1)
	_dadb.SetSemiHidden(true)
	_dadb.SetUnhideWhenUsed(true)
	_gaed := []measurement.Distance{16, 13, 12, 11, 11, 11, 11, 11, 11}
	_ffdgd := []measurement.Distance{240, 40, 40, 40, 40, 40, 40, 40, 40}
	for _fffd := 0; _fffd < 9; _fffd++ {
		_efbcc := fmt.Sprintf("Heading\u0025d", _fffd+1)
		_egaab := s.AddStyle(_efbcc+"Char", wml.ST_StyleTypeCharacter, false)
		_egaab.SetName(fmt.Sprintf("Heading\u0020\u0025d\u0020Char", _fffd+1))
		_egaab.SetBasedOn(_dbad.StyleID())
		_egaab.SetLinkedStyle(_efbcc)
		_egaab.SetUISortOrder(9 + _fffd)
		_egaab.RunProperties().SetSize(_gaed[_fffd] * measurement.Point)
		_ffeb := s.AddStyle(_efbcc, wml.ST_StyleTypeParagraph, false)
		_ffeb.SetName(fmt.Sprintf("heading\u0020\u0025d", _fffd+1))
		_ffeb.SetNextStyle(_becf.StyleID())
		_ffeb.SetLinkedStyle(_ffeb.StyleID())
		_ffeb.SetUISortOrder(9 + _fffd)
		_ffeb.SetPrimaryStyle(true)
		_ffeb.ParagraphProperties().SetKeepNext(true)
		_ffeb.ParagraphProperties().SetSpacing(_ffdgd[_fffd]*measurement.Twips, 0)
		_ffeb.ParagraphProperties().SetOutlineLevel(_fffd)
		_ffeb.RunProperties().SetSize(_gaed[_fffd] * measurement.Point)
	}
}

// Tables returns the tables defined in the footer.
func (f Footer) Tables() []Table {
	result := []Table{}
	if f.WFooter == nil {
		return nil
	}
	for _, v := range f.WFooter.EG_ContentBlockContent {
		result = append(result, f.Document.tables(v)...)
	}
	return result
}

// ExtractFromFooter returns text from the document footer as an array of TextItems.
func ExtractFromFooter(footer *wml.Ftr) []TextItem {
	return parseTextItemList(footer.EG_ContentBlockContent, nil)
}

func (e Endnote) content() []*wml.EG_ContentBlockContent {
	var result []*wml.EG_ContentBlockContent
	for _, v := range e.WEndnote.EG_BlockLevelElts {
		result = append(result, v.EG_ContentBlockContent...)
	}
	return result
}

// SizeValue returns the value of paragraph font size in points.
func (p ParagraphProperties) SizeValue() float64 {
	if _fcbfe := p.Properties.RPr.Sz; _fcbfe != nil {
		_fbged := _fcbfe.ValAttr
		if _fbged.ST_UnsignedDecimalNumber != nil {
			return float64(*_fbged.ST_UnsignedDecimalNumber) / 2
		}
	}
	return 0.0
}

// AddFooter creates a Footer associated with the document, but doesn't add it
// to the document for display.
func (d *Document) AddFooter() Footer {
	footer := wml.NewFtr()
	d.WFooter = append(d.WFooter, footer)
	_fgea := fmt.Sprintf("footer\u0025d\u002exml", len(d.WFooter))
	d._fbb.AddRelationship(_fgea, unioffice.FooterType)
	d.ContentTypes.AddOverride("\u002fword\u002f"+_fgea, "application\u002fvnd.openxmlformats\u002dofficedocument\u002ewordprocessingml\u002efooter\u002bxml")
	d._fcbd = append(d._fcbd, common.NewRelationships())
	return Footer{d, footer}
}

// SetFirstRow controls the conditional formatting for the first row in a table.
func (t TableLook) SetFirstRow(on bool) {
	if !on {
		t.WTableLook.FirstRowAttr = &sharedTypes.ST_OnOff{}
		t.WTableLook.FirstRowAttr.ST_OnOff1 = sharedTypes.ST_OnOff1Off
	} else {
		t.WTableLook.FirstRowAttr = &sharedTypes.ST_OnOff{}
		t.WTableLook.FirstRowAttr.ST_OnOff1 = sharedTypes.ST_OnOff1On
	}
}

func (f FormFieldType) String() string {
	if f >= FormFieldType(len(_dfff)-1) {
		return fmt.Sprintf("FormFieldType\u0028\u0025d\u0029", f)
	}
	return _acga[_dfff[f]:_dfff[f+1]]
}

// SetWidth sets the cell width to a specified width.
func (c CellProperties) SetWidth(d measurement.Distance) {
	c.WProperties.TcW = wml.NewCT_TblWidth()
	c.WProperties.TcW.TypeAttr = wml.ST_TblWidthDxa
	c.WProperties.TcW.WAttr = &wml.ST_MeasurementOrPercent{}
	c.WProperties.TcW.WAttr.ST_DecimalNumberOrPercent = &wml.ST_DecimalNumberOrPercent{}
	c.WProperties.TcW.WAttr.ST_DecimalNumberOrPercent.ST_UnqualifiedPercentage = unioffice.Int64(int64(d / measurement.Twips))
}

// SetHAlignment sets the horizontal alignment for an anchored image.
func (a AnchoredDrawing) SetHAlignment(h wml.WdST_AlignH) {
	a.WAnchoredDrawing.PositionH.Choice = &wml.WdCT_PosHChoice{}
	a.WAnchoredDrawing.PositionH.Choice.Align = h
}

// InsertRowAfter inserts a row after another row
func (t Table) InsertRowAfter(r Row) Row {
	for _aegb, _eaddfd := range t.WTable.EG_ContentRowContent {
		if len(_eaddfd.Tr) > 0 && r.X() == _eaddfd.Tr[0] {
			_bacfg := wml.NewEG_ContentRowContent()
			if len(t.WTable.EG_ContentRowContent) < _aegb+2 {
				return t.AddRow()
			}
			t.WTable.EG_ContentRowContent = append(t.WTable.EG_ContentRowContent, nil)
			copy(t.WTable.EG_ContentRowContent[_aegb+2:], t.WTable.EG_ContentRowContent[_aegb+1:])
			t.WTable.EG_ContentRowContent[_aegb+1] = _bacfg
			_eafd := wml.NewCT_Row()
			_bacfg.Tr = append(_bacfg.Tr, _eafd)
			return Row{t.Document, _eafd}
		}
	}
	return t.AddRow()
}

// Caps returns true if paragraph font is capitalized.
func (p ParagraphProperties) Caps() bool { return checkAttributeSet(p.Properties.RPr.Caps) }

// SetFirstLineIndent controls the indentation of the first line in a paragraph.
func (p ParagraphProperties) SetFirstLineIndent(m measurement.Distance) {
	if p.Properties.Ind == nil {
		p.Properties.Ind = wml.NewCT_Ind()
	}
	if m == measurement.Zero {
		p.Properties.Ind.FirstLineAttr = nil
	} else {
		p.Properties.Ind.FirstLineAttr = &sharedTypes.ST_TwipsMeasure{}
		p.Properties.Ind.FirstLineAttr.ST_UnsignedDecimalNumber = unioffice.Uint64(uint64(m / measurement.Twips))
	}
}

// RemoveParagraph removes a paragraph from a footer.
func (f Footer) RemoveParagraph(p Paragraph) {
	for _, _agfb := range f.WFooter.EG_ContentBlockContent {
		for _dfee, _acgc := range _agfb.P {
			if _acgc == p.WParagraph {
				copy(_agfb.P[_dfee:], _agfb.P[_dfee+1:])
				_agfb.P = _agfb.P[0 : len(_agfb.P)-1]
				return
			}
		}
	}
}

// SetBottom sets the bottom border to a specified type, color and thickness.
func (t TableBorders) SetBottom(b wml.ST_Border, c color.Color, thickness measurement.Distance) {
	t.WBorders.Bottom = wml.NewCT_Border()
	setBorder(t.WBorders.Bottom, b, c, thickness)
}

// AddCheckBox adds checkbox form field to the paragraph and returns it.
func (p Paragraph) AddCheckBox(name string) FormField {
	_ggfe := p.addFldCharsForField(name, "FORMCHECKBOX")
	_ggfe.WData.CheckBox = wml.NewCT_FFCheckBox()
	return _ggfe
}

// X returns the inner wrapped XML type.
func (t TableProperties) X() *wml.CT_TblPr { return t.WProperties }

// SizeMeasure returns font with its measure which can be mm, cm, in, pt, pc or pi.
func (r RunProperties) SizeMeasure() string {
	if _bebgf := r.WProperties.Sz; _bebgf != nil {
		_gbdea := _bebgf.ValAttr
		if _gbdea.ST_PositiveUniversalMeasure != nil {
			return *_gbdea.ST_PositiveUniversalMeasure
		}
	}
	return ""
}

// Underline returns the type of run underline.
func (r RunProperties) Underline() wml.ST_Underline {
	if _bgcg := r.WProperties.U; _bgcg != nil {
		return _bgcg.ValAttr
	}
	return 0
}

// SetBottomPct sets the cell bottom margin
func (c CellMargins) SetBottomPct(pct float64) {
	c.WMargins.Bottom = wml.NewCT_TblWidth()
	setTableMarginPercent(c.WMargins.Bottom, pct)
}

// ParagraphSpacing controls the spacing for a paragraph and its lines.
type ParagraphSpacing struct{ WSpacing *wml.CT_Spacing }

// TextItem is used for keeping text with references to a paragraph and run or a table, a row and a cell where it is located.
type TextItem struct {
	Text        string
	DrawingInfo *DrawingInfo
	WParagraph  *wml.CT_P
	WHyperlink  *wml.CT_Hyperlink
	WRun        *wml.CT_R
	TableInfo   *TableInfo
}

// SetEffect sets a text effect on the run.
func (r RunProperties) SetEffect(e wml.ST_TextEffect) {
	if e == wml.ST_TextEffectUnset {
		r.WProperties.Effect = nil
	} else {
		r.WProperties.Effect = wml.NewCT_TextEffect()
		r.WProperties.Effect.ValAttr = wml.ST_TextEffectShimmer
	}
}

// Footnote returns the footnote based on the ID; this can be used nicely with
// the run.IsFootnote() functionality.
func (d *Document) Footnote(id int64) Footnote {
	for _, _acae := range d.Footnotes() {
		if _acae.id() == id {
			return _acae
		}
	}
	return Footnote{}
}

// AddDropdownList adds dropdown list form field to the paragraph and returns it.
func (p Paragraph) AddDropdownList(name string) FormField {
	_bgd := p.addFldCharsForField(name, "FORMDROPDOWN")
	_bgd.WData.DdList = wml.NewCT_FFDDList()
	return _bgd
}

func getDocumentFromReader(readAt io.ReaderAt, _aacb int64, _ecdd string) (*Document, error) {
	const _aae = "document\u002eRead"
	d := New()
	d.Numbering.WNumbering = nil
	if len(_ecdd) > 0 {
		d.UnknownMeaning = _ecdd
	}
	_ebb, err := tempstorage.TempDir("unioffice-docx")
	if err != nil {
		return nil, err
	}
	d.TmpPath = _ebb
	_bgg, err := zip.NewReader(readAt, _aacb)
	if err != nil {
		return nil, fmt.Errorf("parsing\u0020zip:\u0020\u0025s", err)
	}
	_bcgd := []*zip.File{}
	_bcgd = append(_bcgd, _bgg.File...)
	_dcaf := false
	for _, _ggfd := range _bcgd {
		if _ggfd.FileHeader.Name == "docProps\u002fcustom\u002exml" {
			_dcaf = true
			break
		}
	}
	if _dcaf {
		d.CreateCustomProperties()
	}
	_gae := d.Document.ConformanceAttr
	_ggcc := zippkg.DecodeMap{}
	_ggcc.SetOnNewRelationshipFunc(d.onNewRelationship)
	_ggcc.AddTarget(unioffice.ContentTypesFilename, d.ContentTypes.X(), "", 0)
	_ggcc.AddTarget(unioffice.BaseRelsFilename, d.Rels.X(), "", 0)
	if _ebeb := _ggcc.Decode(_bcgd); _ebeb != nil {
		return nil, _ebeb
	}
	d.Document.ConformanceAttr = _gae
	for _, _bae := range _bcgd {
		if _bae == nil {
			continue
		}
		if _feba := d.AddExtraFileFromZip(_bae); _feba != nil {
			return nil, _feba
		}
	}
	if _dcaf {
		_cfbg := false
		for _, _ecdb := range d.Rels.X().Relationship {
			if _ecdb.TargetAttr == "docProps\u002fcustom\u002exml" {
				_cfbg = true
				break
			}
		}
		if !_cfbg {
			d.AddCustomRelationships()
		}
	}
	return d, nil
}

// BodySection returns the default body section used for all preceding
// paragraphs until the previous Section. If there is no previous sections, the
// body section applies to the entire document.
func (d *Document) BodySection() Section {
	if d.Document.Body.SectPr == nil {
		d.Document.Body.SectPr = wml.NewCT_SectPr()
	}
	return Section{d, d.Document.Body.SectPr}
}

// Paragraphs returns the paragraphs defined in the cell.
func (c Cell) Paragraphs() []Paragraph {
	result := []Paragraph{}
	for _, _eb := range c.WCell.EG_BlockLevelElts {
		for _, _ac := range _eb.EG_ContentBlockContent {
			for _, _fge := range _ac.P {
				result = append(result, Paragraph{c.Document, _fge})
			}
		}
	}
	return result
}

// X returns the inner wrapped XML type.
func (c Cell) X() *wml.CT_Tc { return c.WCell }

// AddImage adds an image to the document package, returning a reference that
// can be used to add the image to a run and place it in the document contents.
func (h Header) AddImage(i common.Image) (common.ImageRef, error) {
	var _eeeb common.Relationships
	for _gfca, _daefc := range h.Document.WHeader {
		if _daefc == h.WHeader {
			_eeeb = h.Document._ddc[_gfca]
		}
	}
	_fdff := common.MakeImageRef(i, &h.Document.DocBase, _eeeb)
	if i.Data == nil && i.Path == "" {
		return _fdff, errors.New("image\u0020must have\u0020data\u0020or\u0020a\u0020path")
	}
	if i.Format == "" {
		return _fdff, errors.New("image\u0020must have\u0020a\u0020valid\u0020format")
	}
	if i.Size.X == 0 || i.Size.Y == 0 {
		return _fdff, errors.New("image\u0020must\u0020have a valid\u0020size")
	}
	h.Document.Images = append(h.Document.Images, _fdff)
	_acab := fmt.Sprintf("media\u002fimage\u0025d\u002e\u0025s", len(h.Document.Images), i.Format)
	_dafc := _eeeb.AddRelationship(_acab, unioffice.ImageType)
	_fdff.SetRelID(_dafc.X().IdAttr)
	return _fdff, nil
}

// Italic returns true if run font is italic.
func (r RunProperties) Italic() bool {
	_fded := r.WProperties
	return checkAttributeSet(_fded.I) || checkAttributeSet(_fded.ICs)
}

// SetCharacterSpacing sets the run's Character Spacing Adjustment.
func (r RunProperties) SetCharacterSpacing(size measurement.Distance) {
	r.WProperties.Spacing = wml.NewCT_SignedTwipsMeasure()
	r.WProperties.Spacing.ValAttr.Int64 = unioffice.Int64(int64(size / measurement.Twips))
}

// SetTopPct sets the cell top margin
func (c CellMargins) SetTopPct(pct float64) {
	c.WMargins.Top = wml.NewCT_TblWidth()
	setTableMarginPercent(c.WMargins.Top, pct)
}

// SetInsideHorizontal sets the interior horizontal borders to a specified type, color and thickness.
func (c CellBorders) SetInsideHorizontal(t wml.ST_Border, co color.Color, thickness measurement.Distance) {
	c.WBorders.InsideH = wml.NewCT_Border()
	setBorder(c.WBorders.InsideH, t, co, thickness)
}

// SetTarget sets the URL target of the hyperlink.
func (h HyperLink) SetTarget(url string) {
	_gbda := h.Document.AddHyperlink(url)
	h.WHyperLink.IdAttr = unioffice.String(common.Relationship(_gbda).ID())
	h.WHyperLink.AnchorAttr = nil
}

// AddHyperLink adds a new hyperlink to a parapgraph.
func (p Paragraph) AddHyperLink() HyperLink {
	_beaf := wml.NewEG_PContent()
	p.WParagraph.EG_PContent = append(p.WParagraph.EG_PContent, _beaf)
	_beaf.Hyperlink = wml.NewCT_Hyperlink()
	return HyperLink{p.Document, _beaf.Hyperlink}
}

// ClearColor clears the text color.
func (r RunProperties) ClearColor() { r.WProperties.Color = nil }

// AddTextInput adds text input form field to the paragraph and returns it.
func (p Paragraph) AddTextInput(name string) FormField {
	_cebc := p.addFldCharsForField(name, "FORMTEXT")
	_cebc.WData.TextInput = wml.NewCT_FFTextInput()
	return _cebc
}

// read reads a document from an io.Reader.
func Read(r io.ReaderAt, size int64) (*Document, error) { return getDocumentFromReader(r, size, "") }

// SetAlignment controls the paragraph alignment
func (p ParagraphStyleProperties) SetAlignment(align wml.ST_Jc) {
	if align == wml.ST_JcUnset {
		p.WProperties.Jc = nil
	} else {
		p.WProperties.Jc = wml.NewCT_Jc()
		p.WProperties.Jc.ValAttr = align
	}
}

// Bold returns true if run font is bold.
func (r RunProperties) Bold() bool {
	_bceb := r.WProperties
	return checkAttributeSet(_bceb.B) || checkAttributeSet(_bceb.BCs)
}

// SetLinkedStyle sets the style that this style is linked to.
func (s Style) SetLinkedStyle(name string) {
	if name == "" {
		s.WStyle.Link = nil
	} else {
		s.WStyle.Link = wml.NewCT_String()
		s.WStyle.Link.ValAttr = name
	}
}

// Open opens and reads a document from a file (.docx).
func Open(filename string) (*Document, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error\u0020opening\u0020\u0025s: \u0025s", filename, err)
	}
	defer f.Close()
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return nil, fmt.Errorf("error\u0020opening\u0020\u0025s: \u0025s", filename, err)
	}
	return Read(f, fileInfo.Size())
}

func (p Paragraph) addFldChar() *wml.CT_FldChar {
	_adad := p.AddRun()
	_abce := _adad.X()
	_cfede := wml.NewEG_RunInnerContent()
	_egbb := wml.NewCT_FldChar()
	_cfede.FldChar = _egbb
	_abce.EG_RunInnerContent = append(_abce.EG_RunInnerContent, _cfede)
	return _egbb
}

// Clear clears all content within a header
func (h Header) Clear() { h.WHeader.EG_ContentBlockContent = nil }

// AbstractNumberID returns the ID that is unique within all numbering
// definitions that is used to assign the definition to a paragraph.
func (n NumberingDefinition) AbstractNumberID() int64 { return n.WDefinition.AbstractNumIdAttr }

// SetHeight allows controlling the height of a row within a table.
func (r RowProperties) SetHeight(ht measurement.Distance, rule wml.ST_HeightRule) {
	if rule == wml.ST_HeightRuleUnset {
		r.WProperties.TrHeight = nil
	} else {
		height := wml.NewCT_Height()
		height.HRuleAttr = rule
		height.ValAttr = &sharedTypes.ST_TwipsMeasure{}
		height.ValAttr.ST_UnsignedDecimalNumber = unioffice.Uint64(uint64(ht / measurement.Twips))
		r.WProperties.TrHeight = []*wml.CT_Height{height}
	}
}

// X returns the inner wrapped type
func (c CellBorders) X() *wml.CT_TcBorders { return c.WBorders }

// NumberingLevel is the definition for numbering for a particular level within
// a NumberingDefinition.
type NumberingLevel struct{ WLevel *wml.CT_Lvl }

func (f Footnote) content() []*wml.EG_ContentBlockContent {
	var _gca []*wml.EG_ContentBlockContent
	for _, _agga := range f.WFootnote.EG_BlockLevelElts {
		_gca = append(_gca, _agga.EG_ContentBlockContent...)
	}
	return _gca
}

// IsEndnote returns a bool based on whether the run has a
// footnote or not. Returns both a bool as to whether it has
// a footnote as well as the ID of the footnote.
func (r Run) IsEndnote() (bool, int64) {
	if r.WRun.EG_RunInnerContent != nil {
		if r.WRun.EG_RunInnerContent[0].EndnoteReference != nil {
			return true, r.WRun.EG_RunInnerContent[0].EndnoteReference.IdAttr
		}
	}
	return false, 0
}

// X returns the inner wrapped XML type.
func (c Color) X() *wml.CT_Color { return c.WColor }

// SetTargetBookmark sets the bookmark target of the hyperlink.
func (h HyperLink) SetTargetBookmark(bm Bookmark) {
	h.WHyperLink.AnchorAttr = unioffice.String(bm.Name())
	h.WHyperLink.IdAttr = nil
}

// HyperLink is a link within a document.
type HyperLink struct {
	Document   *Document
	WHyperLink *wml.CT_Hyperlink
}

// AddParagraph adds a paragraph to the footnote.
func (f Footnote) AddParagraph() Paragraph {
	_fbaa := wml.NewEG_ContentBlockContent()
	_dafbf := len(f.WFootnote.EG_BlockLevelElts[0].EG_ContentBlockContent)
	f.WFootnote.EG_BlockLevelElts[0].EG_ContentBlockContent = append(f.WFootnote.EG_BlockLevelElts[0].EG_ContentBlockContent, _fbaa)
	_aedaf := wml.NewCT_P()
	var _fbcde *wml.CT_String
	if _dafbf != 0 {
		_cgfg := len(f.WFootnote.EG_BlockLevelElts[0].EG_ContentBlockContent[_dafbf-1].P)
		_fbcde = f.WFootnote.EG_BlockLevelElts[0].EG_ContentBlockContent[_dafbf-1].P[_cgfg-1].PPr.PStyle
	} else {
		_fbcde = wml.NewCT_String()
		_fbcde.ValAttr = "Footnote"
	}
	_fbaa.P = append(_fbaa.P, _aedaf)
	result := Paragraph{f.Document, _aedaf}
	result.WParagraph.PPr = wml.NewCT_PPr()
	result.WParagraph.PPr.PStyle = _fbcde
	result.WParagraph.PPr.RPr = wml.NewCT_ParaRPr()
	return result
}

// SetKeepNext controls if the paragraph is kept with the next paragraph.
func (p ParagraphStyleProperties) SetKeepNext(b bool) {
	if !b {
		p.WProperties.KeepNext = nil
	} else {
		p.WProperties.KeepNext = wml.NewCT_OnOff()
	}
}

// SetSize sets the font size for a run.
func (r RunProperties) SetSize(size measurement.Distance) {
	r.WProperties.Sz = wml.NewCT_HpsMeasure()
	r.WProperties.Sz.ValAttr.ST_UnsignedDecimalNumber = unioffice.Uint64(uint64(size / measurement.HalfPoint))
	r.WProperties.SzCs = wml.NewCT_HpsMeasure()
	r.WProperties.SzCs.ValAttr.ST_UnsignedDecimalNumber = unioffice.Uint64(uint64(size / measurement.HalfPoint))
}
func (d Document) mergeFields() []mergeFieldInfo {
	pList := []Paragraph{}
	fieldInfoList := []mergeFieldInfo{}
	for _, t := range d.Tables() {
		for _, r := range t.Rows() {
			for _, c := range r.Cells() {
				pList = append(pList, c.Paragraphs()...)
			}
		}
	}
	pList = append(pList, d.Paragraphs()...)
	for _, p := range pList {
		_efaf := p.Runs()
		_gdbd := -1
		_efafa := -1
		_egec := -1
		fieldInfo := mergeFieldInfo{}
		for _, pContent := range p.WParagraph.EG_PContent {
			for _, f := range pContent.FldSimple {
				if strings.Contains(f.InstrAttr, "MERGEFIELD") {
					_fdfg := _bggc(f.InstrAttr)
					_fdfg._bgfa = true
					_fdfg._aeab = p
					_fdfg._aefd = pContent
					fieldInfoList = append(fieldInfoList, _fdfg)
				}
			}
		}
		for i := 0; i < len(_efaf); i++ {
			_afece := _efaf[i]
			for _, _bbfgc := range _afece.X().EG_RunInnerContent {
				if _bbfgc.FldChar != nil {
					switch _bbfgc.FldChar.FldCharTypeAttr {
					case wml.ST_FldCharTypeBegin:
						_gdbd = i
					case wml.ST_FldCharTypeSeparate:
						_efafa = i
					case wml.ST_FldCharTypeEnd:
						_egec = i
						if fieldInfo._cbbg != "" {
							fieldInfo._aeab = p
							fieldInfo._ebgd = _gdbd
							fieldInfo._fdg = _egec
							fieldInfo._aabb = _efafa
							fieldInfoList = append(fieldInfoList, fieldInfo)
						}
						_gdbd = -1
						_efafa = -1
						_egec = -1
						fieldInfo = mergeFieldInfo{}
					}
				} else if _bbfgc.InstrText != nil && strings.Contains(_bbfgc.InstrText.Content, "MERGEFIELD") {
					if _gdbd != -1 && _egec == -1 {
						fieldInfo = _bggc(_bbfgc.InstrText.Content)
					}
				}
			}
		}
	}
	return fieldInfoList
}

// Index returns the index of the header within the document.  This is used to
// form its zip packaged filename as well as to match it with its relationship
// ID.
func (h Header) Index() int {
	for k, v := range h.Document.WHeader {
		if v == h.WHeader {
			return k
		}
	}
	return -1
}

// SaveToFile writes the document out to a file.
func (d *Document) SaveToFile(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return d.Save(f)
}

// Italic returns true if paragraph font is italic.
func (p ParagraphProperties) Italic() bool {
	r := p.Properties.RPr
	return checkAttributeSet(r.I) || checkAttributeSet(r.ICs)
}

// FormFieldType is the type of the form field.
//go:generate stringer -type=FormFieldType
type FormFieldType byte

// ParagraphStyles returns only the paragraph styles.
func (s Styles) ParagraphStyles() []Style {
	result := []Style{}
	for _, _edbeg := range s.WStyles.Style {
		if _edbeg.TypeAttr != wml.ST_StyleTypeParagraph {
			continue
		}
		result = append(result, Style{_edbeg})
	}
	return result
}
func (r Run) newIC() *wml.EG_RunInnerContent {
	c := wml.NewEG_RunInnerContent()
	r.WRun.EG_RunInnerContent = append(r.WRun.EG_RunInnerContent, c)
	return c
}

// RemoveParagraph removes a paragraph from a document.
func (d *Document) RemoveParagraph(p Paragraph) {
	if d.Document.Body == nil {
		return
	}
	for _, _ddb := range d.Document.Body.EG_BlockLevelElts {
		for _, _gfdc := range _ddb.EG_ContentBlockContent {
			for _ddcf, _bag := range _gfdc.P {
				if _bag == p.WParagraph {
					copy(_gfdc.P[_ddcf:], _gfdc.P[_ddcf+1:])
					_gfdc.P = _gfdc.P[0 : len(_gfdc.P)-1]
					return
				}
			}
			if _gfdc.Sdt != nil && _gfdc.Sdt.SdtContent != nil && _gfdc.Sdt.SdtContent.P != nil {
				for _bedb, _ceb := range _gfdc.Sdt.SdtContent.P {
					if _ceb == p.WParagraph {
						copy(_gfdc.P[_bedb:], _gfdc.P[_bedb+1:])
						_gfdc.P = _gfdc.P[0 : len(_gfdc.P)-1]
						return
					}
				}
			}
		}
	}
	for _, _fgd := range d.Tables() {
		for _, _cafc := range _fgd.Rows() {
			for _, _faf := range _cafc.Cells() {
				for _, _ffgbe := range _faf.WCell.EG_BlockLevelElts {
					for _, _agfc := range _ffgbe.EG_ContentBlockContent {
						for _ega, _ggge := range _agfc.P {
							if _ggge == p.WParagraph {
								copy(_agfc.P[_ega:], _agfc.P[_ega+1:])
								_agfc.P = _agfc.P[0 : len(_agfc.P)-1]
								return
							}
						}
					}
				}
			}
		}
	}
	for _, v := range d.Headers() {
		v.RemoveParagraph(p)
	}
	for _, v := range d.Footers() {
		v.RemoveParagraph(p)
	}
}

// SetRowBandSize sets the number of Rows in the row band
func (t TableStyleProperties) SetRowBandSize(rows int64) {
	t.WProperties.TblStyleRowBandSize = wml.NewCT_DecimalNumber()
	t.WProperties.TblStyleRowBandSize.ValAttr = rows
}

// AddBreak adds a line break to a run.
func (r Run) AddBreak() { _gbfaa := r.newIC(); _gbfaa.Br = wml.NewCT_Br() }

// Clear resets the numbering.
func (n Numbering) Clear() {
	n.WNumbering.AbstractNum = nil
	n.WNumbering.Num = nil
	n.WNumbering.NumIdMacAtCleanup = nil
	n.WNumbering.NumPicBullet = nil
}

// SetColumnBandSize sets the number of Columns in the column band
func (t TableStyleProperties) SetColumnBandSize(cols int64) {
	t.WProperties.TblStyleColBandSize = wml.NewCT_DecimalNumber()
	t.WProperties.TblStyleColBandSize.ValAttr = cols
}
func _gfb(_fdde *wml.CT_P, _cbfa, _fcc map[int64]int64) {
	for _, _aeef := range _fdde.EG_PContent {
		for _, _bggd := range _aeef.EG_ContentRunContent {
			if _bggd.R != nil {
				for _, _edff := range _bggd.R.EG_RunInnerContent {
					_cef := _edff.EndnoteReference
					if _cef != nil && _cef.IdAttr > 0 {
						if _ced, _fdba := _fcc[_cef.IdAttr]; _fdba {
							_cef.IdAttr = _ced
						}
					}
					_cfab := _edff.FootnoteReference
					if _cfab != nil && _cfab.IdAttr > 0 {
						if _gbea, _bagg := _cbfa[_cfab.IdAttr]; _bagg {
							_cfab.IdAttr = _gbea
						}
					}
				}
			}
		}
	}
}

// X returns the inner wrapped XML type.
func (s Style) X() *wml.CT_Style { return s.WStyle }

// Levels returns all of the numbering levels defined in the definition.
func (n NumberingDefinition) Levels() []NumberingLevel {
	result := []NumberingLevel{}
	for _, v := range n.WDefinition.Lvl {
		result = append(result, NumberingLevel{v})
	}
	return result
}

// SetWindowControl controls if the first or last line of the paragraph is
// allowed to dispay on a separate page.
func (p ParagraphProperties) SetWindowControl(b bool) {
	if !b {
		p.Properties.WidowControl = nil
	} else {
		p.Properties.WidowControl = wml.NewCT_OnOff()
	}
}

// EastAsiaFont returns the name of run font family for East Asia.
func (r RunProperties) EastAsiaFont() string {
	if _dccc := r.WProperties.RFonts; _dccc != nil {
		if _dccc.EastAsiaAttr != nil {
			return *_dccc.EastAsiaAttr
		}
	}
	return ""
}

// X returns the inner wrapped XML type.
func (i InlineDrawing) X() *wml.WdInline { return i.WInlineDrawing }

// InitializeDefault constructs a default numbering.
func (n Numbering) InitializeDefault() {
	_gbdc := wml.NewCT_AbstractNum()
	_gbdc.MultiLevelType = wml.NewCT_MultiLevelType()
	_gbdc.MultiLevelType.ValAttr = wml.ST_MultiLevelTypeHybridMultilevel
	n.WNumbering.AbstractNum = append(n.WNumbering.AbstractNum, _gbdc)
	_gbdc.AbstractNumIdAttr = 1
	const _bdef = 720
	const _bfef = 720
	const _ggced = 360
	for _gade := 0; _gade < 9; _gade++ {
		_abde := wml.NewCT_Lvl()
		_abde.IlvlAttr = int64(_gade)
		_abde.Start = wml.NewCT_DecimalNumber()
		_abde.Start.ValAttr = 1
		_abde.NumFmt = wml.NewCT_NumFmt()
		_abde.NumFmt.ValAttr = wml.ST_NumberFormatBullet
		_abde.Suff = wml.NewCT_LevelSuffix()
		_abde.Suff.ValAttr = wml.ST_LevelSuffixNothing
		_abde.LvlText = wml.NewCT_LevelText()
		_abde.LvlText.ValAttr = unioffice.String("\uf0b7")
		_abde.LvlJc = wml.NewCT_Jc()
		_abde.LvlJc.ValAttr = wml.ST_JcLeft
		_abde.RPr = wml.NewCT_RPr()
		_abde.RPr.RFonts = wml.NewCT_Fonts()
		_abde.RPr.RFonts.AsciiAttr = unioffice.String("Symbol")
		_abde.RPr.RFonts.HAnsiAttr = unioffice.String("Symbol")
		_abde.RPr.RFonts.HintAttr = wml.ST_HintDefault
		_abde.PPr = wml.NewCT_PPrGeneral()
		_ddecd := int64(_gade*_bfef + _bdef)
		_abde.PPr.Ind = wml.NewCT_Ind()
		_abde.PPr.Ind.LeftAttr = &wml.ST_SignedTwipsMeasure{}
		_abde.PPr.Ind.LeftAttr.Int64 = unioffice.Int64(_ddecd)
		_abde.PPr.Ind.HangingAttr = &sharedTypes.ST_TwipsMeasure{}
		_abde.PPr.Ind.HangingAttr.ST_UnsignedDecimalNumber = unioffice.Uint64(uint64(_ggced))
		_gbdc.Lvl = append(_gbdc.Lvl, _abde)
	}
	_egca := wml.NewCT_Num()
	_egca.NumIdAttr = 1
	_egca.AbstractNumId = wml.NewCT_DecimalNumber()
	_egca.AbstractNumId.ValAttr = 1
	n.WNumbering.Num = append(n.WNumbering.Num, _egca)
}

// SetInsideVertical sets the interior vertical borders to a specified type, color and thickness.
func (_gcd CellBorders) SetInsideVertical(t wml.ST_Border, c color.Color, thickness measurement.Distance) {
	_gcd.WBorders.InsideV = wml.NewCT_Border()
	setBorder(_gcd.WBorders.InsideV, t, c, thickness)
}

// SetCellSpacingPercent sets the cell spacing within a table to a percent width.
func (_cdab TableStyleProperties) SetCellSpacingPercent(pct float64) {
	_cdab.WProperties.TblCellSpacing = wml.NewCT_TblWidth()
	_cdab.WProperties.TblCellSpacing.TypeAttr = wml.ST_TblWidthPct
	_cdab.WProperties.TblCellSpacing.WAttr = &wml.ST_MeasurementOrPercent{}
	_cdab.WProperties.TblCellSpacing.WAttr.ST_DecimalNumberOrPercent = &wml.ST_DecimalNumberOrPercent{}
	_cdab.WProperties.TblCellSpacing.WAttr.ST_DecimalNumberOrPercent.ST_UnqualifiedPercentage = unioffice.Int64(int64(pct * 50))
}

// AddTab adds tab to a run and can be used with the the Paragraph's tab stops.
func (_ddbec Run) AddTab() { _geae := _ddbec.newIC(); _geae.Tab = wml.NewCT_Empty() }

// SetWidth sets the table with to a specified width.
func (_ggbff TableProperties) SetWidth(d measurement.Distance) {
	_ggbff.WProperties.TblW = wml.NewCT_TblWidth()
	_ggbff.WProperties.TblW.TypeAttr = wml.ST_TblWidthDxa
	_ggbff.WProperties.TblW.WAttr = &wml.ST_MeasurementOrPercent{}
	_ggbff.WProperties.TblW.WAttr.ST_DecimalNumberOrPercent = &wml.ST_DecimalNumberOrPercent{}
	_ggbff.WProperties.TblW.WAttr.ST_DecimalNumberOrPercent.ST_UnqualifiedPercentage = unioffice.Int64(int64(d / measurement.Twips))
}

// Validate validates the structure and in cases where it't possible, the ranges
// of elements within a document. A validation error dones't mean that the
// document won't work in MS Word or LibreOffice, but it's worth checking into.
func (_bega *Document) Validate() error {
	if _bega == nil || _bega.Document == nil {
		return errors.New("document\u0020not\u0020initialized correctly\u002c\u0020nil\u0020base")
	}
	for _, _cddd := range []func() error{_bega.validateTableCells, _bega.validateBookmarks} {
		if _febae := _cddd(); _febae != nil {
			return _febae
		}
	}
	if _efc := _bega.Document.Validate(); _efc != nil {
		return _efc
	}
	return nil
}

// AddBookmark adds a bookmark to a document that can then be used from a hyperlink. Name is a document
// unique name that identifies the bookmark so it can be referenced from hyperlinks.
func (_fbdf Paragraph) AddBookmark(name string) Bookmark {
	_bggbc := wml.NewEG_PContent()
	_agfcf := wml.NewEG_ContentRunContent()
	_bggbc.EG_ContentRunContent = append(_bggbc.EG_ContentRunContent, _agfcf)
	_gfcc := wml.NewEG_RunLevelElts()
	_agfcf.EG_RunLevelElts = append(_agfcf.EG_RunLevelElts, _gfcc)
	_faef := wml.NewEG_RangeMarkupElements()
	_eff := wml.NewCT_Bookmark()
	_faef.BookmarkStart = _eff
	_gfcc.EG_RangeMarkupElements = append(_gfcc.EG_RangeMarkupElements, _faef)
	_faef = wml.NewEG_RangeMarkupElements()
	_faef.BookmarkEnd = wml.NewCT_MarkupRange()
	_gfcc.EG_RangeMarkupElements = append(_gfcc.EG_RangeMarkupElements, _faef)
	_fbdf.WParagraph.EG_PContent = append(_fbdf.WParagraph.EG_PContent, _bggbc)
	_fgeb := Bookmark{_eff}
	_fgeb.SetName(name)
	return _fgeb
}

// NewSettings constructs a new empty Settings
func NewSettings() Settings {
	_bbfc := wml.NewSettings()
	_bbfc.Compat = wml.NewCT_Compat()
	_fbgcf := wml.NewCT_CompatSetting()
	_fbgcf.NameAttr = unioffice.String("compatibilityMode")
	_fbgcf.UriAttr = unioffice.String("http:\u002f\u002fschemas\u002emicrosoft\u002ecom\u002foffice/word")
	_fbgcf.ValAttr = unioffice.String("15")
	_bbfc.Compat.CompatSetting = append(_bbfc.Compat.CompatSetting, _fbgcf)
	return Settings{_bbfc}
}

// Properties returns the numbering level paragraph properties.
func (_bfdc NumberingLevel) Properties() ParagraphStyleProperties {
	if _bfdc.WLevel.PPr == nil {
		_bfdc.WLevel.PPr = wml.NewCT_PPrGeneral()
	}
	return ParagraphStyleProperties{_bfdc.WLevel.PPr}
}

// SetShadow sets the run to shadowed text.
func (_gcgf RunProperties) SetShadow(b bool) {
	if !b {
		_gcgf.WProperties.Shadow = nil
	} else {
		_gcgf.WProperties.Shadow = wml.NewCT_OnOff()
	}
}

// SetTop sets the cell top margin
func (_dda CellMargins) SetTop(d measurement.Distance) {
	_dda.WMargins.Top = wml.NewCT_TblWidth()
	setTableMarginDistance(_dda.WMargins.Top, d)
}

// IsItalic returns true if the run has been set to italics.
func (_ddedcc RunProperties) IsItalic() bool { return _ddedcc.ItalicValue() == OnOffValueOn }

// X returns the inner wrapped XML type.
func (_efbc Numbering) X() *wml.Numbering { return _efbc.WNumbering }

// NumberingDefinition defines a numbering definition for a list of pragraphs.
type NumberingDefinition struct{ WDefinition *wml.CT_AbstractNum }

// AddParagraph adds a paragraph to the table cell.
func (_fcg Cell) AddParagraph() Paragraph {
	_ae := wml.NewEG_BlockLevelElts()
	_fcg.WCell.EG_BlockLevelElts = append(_fcg.WCell.EG_BlockLevelElts, _ae)
	_aeb := wml.NewEG_ContentBlockContent()
	_ae.EG_ContentBlockContent = append(_ae.EG_ContentBlockContent, _aeb)
	_dcg := wml.NewCT_P()
	_aeb.P = append(_aeb.P, _dcg)
	return Paragraph{_fcg.Document, _dcg}
}

// X returns the inner wrapped XML type.
func (_afag NumberingLevel) X() *wml.CT_Lvl { return _afag.WLevel }

// SetThemeColor sets the color from the theme.
func (_bfd Color) SetThemeColor(t wml.ST_ThemeColor) { _bfd.WColor.ThemeColorAttr = t }

// SetHeadingLevel sets a heading level and style based on the level to a
// paragraph.  The default styles for a new unioffice document support headings
// from level 1 to 8.
func (_eeb ParagraphProperties) SetHeadingLevel(idx int) {
	_eeb.SetStyle(fmt.Sprintf("Heading\u0025d", idx))
	if _eeb.Properties.NumPr == nil {
		_eeb.Properties.NumPr = wml.NewCT_NumPr()
	}
	_eeb.Properties.NumPr.Ilvl = wml.NewCT_DecimalNumber()
	_eeb.Properties.NumPr.Ilvl.ValAttr = int64(idx)
}

// X returns the inner wrapped XML type.
func (_fcfcd Styles) X() *wml.Styles { return _fcfcd.WStyles }

// RightToLeft returns true if run text goes from right to left.
func (_cbgc RunProperties) RightToLeft() bool { return checkAttributeSet(_cbgc.WProperties.Rtl) }

// SetRight sets the right border to a specified type, color and thickness.
func (_cff CellBorders) SetRight(t wml.ST_Border, c color.Color, thickness measurement.Distance) {
	_cff.WBorders.Right = wml.NewCT_Border()
	setBorder(_cff.WBorders.Right, t, c, thickness)
}

// X returns the inner wrapped XML type.
func (_bagd ParagraphStyleProperties) X() *wml.CT_PPrGeneral { return _bagd.WProperties }

// SetLineSpacing sets the spacing between lines in a paragraph.
func (_gfae Paragraph) SetLineSpacing(d measurement.Distance, rule wml.ST_LineSpacingRule) {
	_gfae.ensurePPr()
	if _gfae.WParagraph.PPr.Spacing == nil {
		_gfae.WParagraph.PPr.Spacing = wml.NewCT_Spacing()
	}
	_ecdbc := _gfae.WParagraph.PPr.Spacing
	if rule == wml.ST_LineSpacingRuleUnset {
		_ecdbc.LineRuleAttr = wml.ST_LineSpacingRuleUnset
		_ecdbc.LineAttr = nil
	} else {
		_ecdbc.LineRuleAttr = rule
		_ecdbc.LineAttr = &wml.ST_SignedTwipsMeasure{}
		_ecdbc.LineAttr.Int64 = unioffice.Int64(int64(d / measurement.Twips))
	}
}

// VerticalAlign returns the value of run vertical align.
func (_dabed RunProperties) VerticalAlignment() sharedTypes.ST_VerticalAlignRun {
	if _bfefc := _dabed.WProperties.VertAlign; _bfefc != nil {
		return _bfefc.ValAttr
	}
	return 0
}

// TableStyleProperties are table properties as defined in a style.
type TableStyleProperties struct{ WProperties *wml.CT_TblPrBase }

// AddEndnote will create a new endnote and attach it to the Paragraph in the
// location at the end of the previous run (endnotes create their own run within
// the paragraph. The text given to the function is simply a convenience helper,
// paragraphs and runs can always be added to the text of the endnote later.
func (_aggg Paragraph) AddEndnote(text string) Endnote {
	var _gecda int64
	if _aggg.Document.HasEndnotes() {
		for _, _efagc := range _aggg.Document.Endnotes() {
			if _efagc.id() > _gecda {
				_gecda = _efagc.id()
			}
		}
		_gecda++
	} else {
		_gecda = 0
		_aggg.Document.WEndnotes = &wml.Endnotes{}
	}
	_fbdg := wml.NewCT_FtnEdn()
	_bcgfb := wml.NewCT_FtnEdnRef()
	_bcgfb.IdAttr = _gecda
	_aggg.Document.WEndnotes.CT_Endnotes.Endnote = append(_aggg.Document.WEndnotes.CT_Endnotes.Endnote, _fbdg)
	_dee := _aggg.AddRun()
	_dbdeb := _dee.Properties()
	_dbdeb.SetStyle("EndnoteAnchor")
	_dee.WRun.EG_RunInnerContent = []*wml.EG_RunInnerContent{wml.NewEG_RunInnerContent()}
	_dee.WRun.EG_RunInnerContent[0].EndnoteReference = _bcgfb
	_faga := Endnote{_aggg.Document, _fbdg}
	_faga.WEndnote.IdAttr = _gecda
	_faga.WEndnote.EG_BlockLevelElts = []*wml.EG_BlockLevelElts{wml.NewEG_BlockLevelElts()}
	_ebge := _faga.AddParagraph()
	_ebge.Properties().SetStyle("Endnote")
	_ebge.WParagraph.PPr.RPr = wml.NewCT_ParaRPr()
	_ageb := _ebge.AddRun()
	_ageb.AddTab()
	_ageb.AddText(text)
	return _faga
}

// TableLook returns the table look, or conditional formatting applied to a table style.
func (_gaece TableProperties) TableLook() TableLook {
	if _gaece.WProperties.TblLook == nil {
		_gaece.WProperties.TblLook = wml.NewCT_TblLook()
	}
	return TableLook{_gaece.WProperties.TblLook}
}

// X returns the inner wrapped XML type.
func (_eadg Footer) X() *wml.Ftr { return _eadg.WFooter }

// Outline returns true if paragraph outline is on.
func (_gbbbe ParagraphProperties) Outline() bool {
	return checkAttributeSet(_gbbbe.Properties.RPr.Outline)
}

// Footers returns the footers defined in the document.
func (_bbe *Document) Footers() []Footer {
	_dae := []Footer{}
	for _, _bg := range _bbe.WFooter {
		_dae = append(_dae, Footer{_bbe, _bg})
	}
	return _dae
}

// SetOffset sets the offset of the image relative to the origin, which by
// default this is the top-left corner of the page. Offset is incompatible with
// SetAlignment, whichever is called last is applied.
func (_bee AnchoredDrawing) SetOffset(x, y measurement.Distance) {
	_bee.SetXOffset(x)
	_bee.SetYOffset(y)
}

// SetStyle sets the table style name.
func (_fbecd TableProperties) SetStyle(name string) {
	if name == "" {
		_fbecd.WProperties.TblStyle = nil
	} else {
		_fbecd.WProperties.TblStyle = wml.NewCT_String()
		_fbecd.WProperties.TblStyle.ValAttr = name
	}
}

// X returns the inner wrapped XML type.
func (_gfe *Document) X() *wml.Document { return _gfe.Document }

// GetImageObjByRelId returns a common.Image with the associated relation ID in the
// document.
func (_gagg *Document) GetImageObjByRelId(relId string) (common.Image, error) {
	_fbced := _gagg._fbb.GetTargetByRelId(relId)
	return _gagg.DocBase.GetImageBytesByTarget(_fbced)
}

// AddRun adds a run to a paragraph.
func (_deae Paragraph) AddRun() Run {
	_gdbc := wml.NewEG_PContent()
	_deae.WParagraph.EG_PContent = append(_deae.WParagraph.EG_PContent, _gdbc)
	_abdc := wml.NewEG_ContentRunContent()
	_gdbc.EG_ContentRunContent = append(_gdbc.EG_ContentRunContent, _abdc)
	_gbfe := wml.NewCT_R()
	_abdc.R = _gbfe
	return Run{_deae.Document, _gbfe}
}

// SetKeepWithNext controls if this paragraph should be kept with the next.
func (_eadb ParagraphProperties) SetKeepWithNext(b bool) {
	if !b {
		_eadb.Properties.KeepNext = nil
	} else {
		_eadb.Properties.KeepNext = wml.NewCT_OnOff()
	}
}

// ParagraphProperties are the properties for a paragraph.
type ParagraphProperties struct {
	Document   *Document
	Properties *wml.CT_PPr
}

// RunProperties returns the run properties controlling text formatting within the table.
func (_fddbb TableConditionalFormatting) RunProperties() RunProperties {
	if _fddbb.WFormat.RPr == nil {
		_fddbb.WFormat.RPr = wml.NewCT_RPr()
	}
	return RunProperties{_fddbb.WFormat.RPr}
}

// InsertRunBefore inserts a run in the paragraph before the relative run.
func (_cffff Paragraph) InsertRunBefore(relativeTo Run) Run {
	return _cffff.insertRun(relativeTo, true)
}

// Bold returns true if paragraph font is bold.
func (_beee ParagraphProperties) Bold() bool {
	_edbeb := _beee.Properties.RPr
	return checkAttributeSet(_edbeb.B) || checkAttributeSet(_edbeb.BCs)
}

// X returns the inner wrapped XML type.
func (_bfdaf Fonts) X() *wml.CT_Fonts { return _bfdaf.WFonts }

// ComplexSizeValue returns the value of run font size for complex fonts in points.
func (_bccg RunProperties) ComplexSizeValue() float64 {
	if _bdebd := _bccg.WProperties.SzCs; _bdebd != nil {
		_fcgbd := _bdebd.ValAttr
		if _fcgbd.ST_UnsignedDecimalNumber != nil {
			return float64(*_fcgbd.ST_UnsignedDecimalNumber) / 2
		}
	}
	return 0.0
}

// AddRow adds a row to a table.
func (_cdbcc Table) AddRow() Row {
	_eegdg := wml.NewEG_ContentRowContent()
	_cdbcc.WTable.EG_ContentRowContent = append(_cdbcc.WTable.EG_ContentRowContent, _eegdg)
	_bagc := wml.NewCT_Row()
	_eegdg.Tr = append(_eegdg.Tr, _bagc)
	return Row{_cdbcc.Document, _bagc}
}

// Endnotes returns the endnotes defined in the document.
func (_dbgg *Document) Endnotes() []Endnote {
	_gac := []Endnote{}
	for _, _cac := range _dbgg.WEndnotes.CT_Endnotes.Endnote {
		_gac = append(_gac, Endnote{_dbgg, _cac})
	}
	return _gac
}

// SetCellSpacingAuto sets the cell spacing within a table to automatic.
func (_facae TableStyleProperties) SetCellSpacingAuto() {
	_facae.WProperties.TblCellSpacing = wml.NewCT_TblWidth()
	_facae.WProperties.TblCellSpacing.TypeAttr = wml.ST_TblWidthAuto
}

// GetImage returns the ImageRef associated with an InlineDrawing.
func (_ddbd InlineDrawing) GetImage() (common.ImageRef, bool) {
	_dabbc := _ddbd.WInlineDrawing.Graphic.GraphicData.Any
	if len(_dabbc) > 0 {
		_cgbe, _degg := _dabbc[0].(*picture.Pic)
		if _degg {
			if _cgbe.BlipFill != nil && _cgbe.BlipFill.Blip != nil && _cgbe.BlipFill.Blip.EmbedAttr != nil {
				return _ddbd.Document.GetImageByRelID(*_cgbe.BlipFill.Blip.EmbedAttr)
			}
		}
	}
	return common.ImageRef{}, false
}

// AddTabStop adds a tab stop to the paragraph.
func (_cacd ParagraphStyleProperties) AddTabStop(position measurement.Distance, justificaton wml.ST_TabJc, leader wml.ST_TabTlc) {
	if _cacd.WProperties.Tabs == nil {
		_cacd.WProperties.Tabs = wml.NewCT_Tabs()
	}
	_ggffd := wml.NewCT_TabStop()
	_ggffd.LeaderAttr = leader
	_ggffd.ValAttr = justificaton
	_ggffd.PosAttr.Int64 = unioffice.Int64(int64(position / measurement.Twips))
	_cacd.WProperties.Tabs.Tab = append(_cacd.WProperties.Tabs.Tab, _ggffd)
}

// Clear removes all of the content from within a run.
func (_dgdg Run) Clear() { _dgdg.WRun.EG_RunInnerContent = nil }

/*
有一个语句
if !license.GetLicenseKey().IsLicensed() && !_gba {
	报错退出
}
据此判断，该变量为true时就可以不在乎是否有license而继续使用
似乎是调试模式的意思，开启调试，就不检查license了
*/
var _gba = false

// Color controls the run or styles color.
type Color struct{ WColor *wml.CT_Color }

// SetLeftIndent controls the left indent of the paragraph.
func (_eabd ParagraphStyleProperties) SetLeftIndent(m measurement.Distance) {
	if _eabd.WProperties.Ind == nil {
		_eabd.WProperties.Ind = wml.NewCT_Ind()
	}
	if m == measurement.Zero {
		_eabd.WProperties.Ind.LeftAttr = nil
	} else {
		_eabd.WProperties.Ind.LeftAttr = &wml.ST_SignedTwipsMeasure{}
		_eabd.WProperties.Ind.LeftAttr.Int64 = unioffice.Int64(int64(m / measurement.Twips))
	}
}

// SetAlignment positions an anchored image via alignment.  Offset is
// incompatible with SetOffset, whichever is called last is applied.
func (_fcb AnchoredDrawing) SetAlignment(h wml.WdST_AlignH, v wml.WdST_AlignV) {
	_fcb.SetHAlignment(h)
	_fcb.SetVAlignment(v)
}

// Style returns the style for a paragraph, or an empty string if it is unset.
func (_afee Paragraph) Style() string {
	if _afee.WParagraph.PPr != nil && _afee.WParagraph.PPr.PStyle != nil {
		return _afee.WParagraph.PPr.PStyle.ValAttr
	}
	return ""
}

// SetEnabled marks a FormField as enabled or disabled.
func (_gdaa FormField) SetEnabled(enabled bool) {
	_egabg := wml.NewCT_OnOff()
	_egabg.ValAttr = &sharedTypes.ST_OnOff{Bool: &enabled}
	_gdaa.WData.Enabled = []*wml.CT_OnOff{_egabg}
}

// Borders allows manipulation of the table borders.
func (_ffca TableStyleProperties) Borders() TableBorders {
	if _ffca.WProperties.TblBorders == nil {
		_ffca.WProperties.TblBorders = wml.NewCT_TblBorders()
	}
	return TableBorders{_ffca.WProperties.TblBorders}
}

// SetHighlight highlights text in a specified color.
func (_cccad RunProperties) SetHighlight(c wml.ST_HighlightColor) {
	_cccad.WProperties.Highlight = wml.NewCT_Highlight()
	_cccad.WProperties.Highlight.ValAttr = c
}
func _abgf(_adbc *wml.CT_P, _ggac *wml.CT_Hyperlink, _eggbf *TableInfo, _gdcg *DrawingInfo, _cega []*wml.EG_ContentRunContent) []TextItem {
	_gcg := []TextItem{}
	for _, _cggf := range _cega {
		if _efec := _cggf.R; _efec != nil {
			_aedb := bytes.NewBuffer([]byte{})
			for _, _ccea := range _efec.EG_RunInnerContent {
				if _ccea.T != nil && _ccea.T.Content != "" {
					_aedb.WriteString(_ccea.T.Content)
				}
			}
			_gcg = append(_gcg, TextItem{Text: _aedb.String(), DrawingInfo: _gdcg, WParagraph: _adbc, WHyperlink: _ggac, WRun: _efec, TableInfo: _eggbf})
			for _, _gbdf := range _efec.Extra {
				if _decg, _ggada := _gbdf.(*wml.AlternateContentRun); _ggada {
					_bcgf := &DrawingInfo{WDrawing: _decg.Choice.Drawing}
					for _, _bbff := range _bcgf.WDrawing.Anchor {
						for _, _edad := range _bbff.Graphic.GraphicData.Any {
							if _fbec, _gfc := _edad.(*wml.WdWsp); _gfc {
								if _fbec.WChoice != nil {
									if _ddea := _fbec.SpPr; _ddea != nil {
										if _edea := _ddea.Xfrm; _edea != nil {
											if _gbc := _edea.Ext; _gbc != nil {
												_bcgf.Width = _gbc.CxAttr
												_bcgf.Height = _gbc.CyAttr
											}
										}
									}
									for _, _acfc := range _fbec.WChoice.Txbx.TxbxContent.EG_ContentBlockContent {
										_gcg = append(_gcg, _gecdc(_acfc.P, _eggbf, _bcgf)...)
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return _gcg
}

// ComplexSizeMeasure returns font with its measure which can be mm, cm, in, pt, pc or pi.
func (_fcce ParagraphProperties) ComplexSizeMeasure() string {
	if _dffgf := _fcce.Properties.RPr.SzCs; _dffgf != nil {
		_dgdef := _dffgf.ValAttr
		if _dgdef.ST_PositiveUniversalMeasure != nil {
			return *_dgdef.ST_PositiveUniversalMeasure
		}
	}
	return ""
}

// SizeValue returns the value of run font size in points.
func (_cfedb RunProperties) SizeValue() float64 {
	if _ggbc := _cfedb.WProperties.Sz; _ggbc != nil {
		_cegb := _ggbc.ValAttr
		if _cegb.ST_UnsignedDecimalNumber != nil {
			return float64(*_cegb.ST_UnsignedDecimalNumber) / 2
		}
	}
	return 0.0
}

// SetWidthAuto sets the the cell width to automatic.
func (_gda CellProperties) SetWidthAuto() {
	_gda.WProperties.TcW = wml.NewCT_TblWidth()
	_gda.WProperties.TcW.TypeAttr = wml.ST_TblWidthAuto
}

// SetColor sets a specific color or auto.
func (_agg Color) SetColor(v color.Color) {
	if v.IsAuto() {
		_agg.WColor.ValAttr.ST_HexColorAuto = wml.ST_HexColorAutoAuto
		_agg.WColor.ValAttr.ST_HexColorRGB = nil
	} else {
		_agg.WColor.ValAttr.ST_HexColorAuto = wml.ST_HexColorAutoUnset
		_agg.WColor.ValAttr.ST_HexColorRGB = v.AsRGBString()
	}
}
func (_cag *Document) validateBookmarks() error {
	_cddf := make(map[string]struct{})
	for _, _gff := range _cag.Bookmarks() {
		if _, _dbde := _cddf[_gff.Name()]; _dbde {
			return fmt.Errorf("duplicate\u0020bookmark\u0020\u0025s found", _gff.Name())
		}
		_cddf[_gff.Name()] = struct{}{}
	}
	return nil
}

// FormFields extracts all of the fields from a document.  They can then be
// manipulated via the methods on the field and the document saved.
func (_daec *Document) FormFields() []FormField {
	_dac := []FormField{}
	for _, _dcgbc := range _daec.Paragraphs() {
		_fgce := _dcgbc.Runs()
		for _ccca, _bffg := range _fgce {
			for _, _ceag := range _bffg.WRun.EG_RunInnerContent {
				if _ceag.FldChar == nil || _ceag.FldChar.FfData == nil {
					continue
				}
				if _ceag.FldChar.FldCharTypeAttr == wml.ST_FldCharTypeBegin {
					if len(_ceag.FldChar.FfData.Name) == 0 || _ceag.FldChar.FfData.Name[0].ValAttr == nil {
						continue
					}
					_cba := FormField{WData: _ceag.FldChar.FfData}
					if _ceag.FldChar.FfData.TextInput != nil {
						for _efdg := _ccca + 1; _efdg < len(_fgce)-1; _efdg++ {
							if len(_fgce[_efdg].WRun.EG_RunInnerContent) == 0 {
								continue
							}
							_agge := _fgce[_efdg].WRun.EG_RunInnerContent[0]
							if _agge.FldChar != nil && _agge.FldChar.FldCharTypeAttr == wml.ST_FldCharTypeSeparate {
								if len(_fgce[_efdg+1].WRun.EG_RunInnerContent) == 0 {
									continue
								}
								if _fgce[_efdg+1].WRun.EG_RunInnerContent[0].FldChar == nil {
									_cba._fdea = _fgce[_efdg+1].WRun.EG_RunInnerContent[0]
									break
								}
							}
						}
					}
					_dac = append(_dac, _cba)
				}
			}
		}
	}
	return _dac
}

// Name returns the name of the field.
func (_fbeb FormField) Name() string { return *_fbeb.WData.Name[0].ValAttr }

// CharacterSpacingMeasure returns paragraph characters spacing with its measure which can be mm, cm, in, pt, pc or pi.
func (_gage ParagraphProperties) CharacterSpacingMeasure() string {
	if _gdbb := _gage.Properties.RPr.Spacing; _gdbb != nil {
		_aabg := _gdbb.ValAttr
		if _aabg.ST_UniversalMeasure != nil {
			return *_aabg.ST_UniversalMeasure
		}
	}
	return ""
}

// SetDoubleStrikeThrough sets the run to double strike-through.
func (_bdfe RunProperties) SetDoubleStrikeThrough(b bool) {
	if !b {
		_bdfe.WProperties.Dstrike = nil
	} else {
		_bdfe.WProperties.Dstrike = wml.NewCT_OnOff()
	}
}

// TableLook is the conditional formatting associated with a table style that
// has been assigned to a table.
type TableLook struct{ WTableLook *wml.CT_TblLook }

// HasFootnotes returns a bool based on the presence or abscence of footnotes within
// the document.
func (_afg *Document) HasFootnotes() bool { return _afg.WFootnotes != nil }

// SetLastRow controls the conditional formatting for the last row in a table.
// This is called the 'Total' row within Word.
func (_gbfd TableLook) SetLastRow(on bool) {
	if !on {
		_gbfd.WTableLook.LastRowAttr = &sharedTypes.ST_OnOff{}
		_gbfd.WTableLook.LastRowAttr.ST_OnOff1 = sharedTypes.ST_OnOff1Off
	} else {
		_gbfd.WTableLook.LastRowAttr = &sharedTypes.ST_OnOff{}
		_gbfd.WTableLook.LastRowAttr.ST_OnOff1 = sharedTypes.ST_OnOff1On
	}
}

// VerticalAlign returns the value of paragraph vertical align.
func (_edagb ParagraphProperties) VerticalAlignment() sharedTypes.ST_VerticalAlignRun {
	if _gbee := _edagb.Properties.RPr.VertAlign; _gbee != nil {
		return _gbee.ValAttr
	}
	return 0
}
func (_efeb Paragraph) addBeginFldChar(_dgbe string) *wml.CT_FFData {
	_cacc := _efeb.addFldChar()
	_cacc.FldCharTypeAttr = wml.ST_FldCharTypeBegin
	_cacc.FfData = wml.NewCT_FFData()
	_bdad := wml.NewCT_FFName()
	_bdad.ValAttr = &_dgbe
	_cacc.FfData.Name = []*wml.CT_FFName{_bdad}
	return _cacc.FfData
}

// SetBefore sets the spacing that comes before the paragraph.
func (_gbde ParagraphSpacing) SetBefore(before measurement.Distance) {
	_gbde.WSpacing.BeforeAttr = &sharedTypes.ST_TwipsMeasure{}
	_gbde.WSpacing.BeforeAttr.ST_UnsignedDecimalNumber = unioffice.Uint64(uint64(before / measurement.Twips))
}

// SetSpacing sets the spacing that comes before and after the paragraph.
func (_gccb ParagraphStyleProperties) SetSpacing(before, after measurement.Distance) {
	if _gccb.WProperties.Spacing == nil {
		_gccb.WProperties.Spacing = wml.NewCT_Spacing()
	}
	if before == measurement.Zero {
		_gccb.WProperties.Spacing.BeforeAttr = nil
	} else {
		_gccb.WProperties.Spacing.BeforeAttr = &sharedTypes.ST_TwipsMeasure{}
		_gccb.WProperties.Spacing.BeforeAttr.ST_UnsignedDecimalNumber = unioffice.Uint64(uint64(before / measurement.Twips))
	}
	if after == measurement.Zero {
		_gccb.WProperties.Spacing.AfterAttr = nil
	} else {
		_gccb.WProperties.Spacing.AfterAttr = &sharedTypes.ST_TwipsMeasure{}
		_gccb.WProperties.Spacing.AfterAttr.ST_UnsignedDecimalNumber = unioffice.Uint64(uint64(after / measurement.Twips))
	}
}

// AddPageBreak adds a page break to a run.
func (_gggb Run) AddPageBreak() {
	_aceg := _gggb.newIC()
	_aceg.Br = wml.NewCT_Br()
	_aceg.Br.TypeAttr = wml.ST_BrTypePage
}

// Text returns text from the document as one string separated with line breaks.
func (_ggce *DocText) Text() string {
	_edef := bytes.NewBuffer([]byte{})
	for _, _gaeb := range _ggce.Items {
		if _gaeb.Text != "" {
			_edef.WriteString(_gaeb.Text)
			_edef.WriteString("\u000a")
		}
	}
	return _edef.String()
}

const (
	FieldCurrentPage   = "PAGE"
	FieldNumberOfPages = "NUMPAGES"
	FieldDate          = "DATE"
	FieldCreateDate    = "CREATEDATE"
	FieldEditTime      = "EDITTIME"
	FieldPrintDate     = "PRINTDATE"
	FieldSaveDate      = "SAVEDATE"
	FieldTIme          = "TIME"
	FieldTOC           = "TOC"
)

// Cells returns the cells defined in the table.
func (_eaff Row) Cells() []Cell {
	_febb := []Cell{}
	for _, _aacc := range _eaff.WRow.EG_ContentCellContent {
		for _, _bccc := range _aacc.Tc {
			_febb = append(_febb, Cell{_eaff.Document, _bccc})
		}
		if _aacc.Sdt != nil && _aacc.Sdt.SdtContent != nil {
			for _, _baea := range _aacc.Sdt.SdtContent.Tc {
				_febb = append(_febb, Cell{_eaff.Document, _baea})
			}
		}
	}
	return _febb
}
func (_dgff *Document) InsertTableBefore(relativeTo Paragraph) Table {
	return _dgff.insertTable(relativeTo, true)
}

// Run is a run of text within a paragraph that shares the same formatting.
type Run struct {
	Document *Document
	WRun     *wml.CT_R
}

// SetUnderline controls underline for a run style.
func (_eegd RunProperties) SetUnderline(style wml.ST_Underline, c color.Color) {
	if style == wml.ST_UnderlineUnset {
		_eegd.WProperties.U = nil
	} else {
		_eegd.WProperties.U = wml.NewCT_Underline()
		_eegd.WProperties.U.ColorAttr = &wml.ST_HexColor{}
		_eegd.WProperties.U.ColorAttr.ST_HexColorRGB = c.AsRGBString()
		_eegd.WProperties.U.ValAttr = style
	}
}

// SetBeforeAuto controls if spacing before a paragraph is automatically determined.
func (_dgfff ParagraphSpacing) SetBeforeAuto(b bool) {
	if b {
		_dgfff.WSpacing.BeforeAutospacingAttr = &sharedTypes.ST_OnOff{}
		_dgfff.WSpacing.BeforeAutospacingAttr.Bool = unioffice.Bool(true)
	} else {
		_dgfff.WSpacing.BeforeAutospacingAttr = nil
	}
}

// SetLeft sets the left border to a specified type, color and thickness.
func (_fdf CellBorders) SetLeft(t wml.ST_Border, c color.Color, thickness measurement.Distance) {
	_fdf.WBorders.Left = wml.NewCT_Border()
	setBorder(_fdf.WBorders.Left, t, c, thickness)
}

// SetToolTip sets the tooltip text for a hyperlink.
func (_dffa HyperLink) SetToolTip(text string) {
	if text == "" {
		_dffa.WHyperLink.TooltipAttr = nil
	} else {
		_dffa.WHyperLink.TooltipAttr = unioffice.String(text)
	}
}

// ParagraphStyleProperties is the styling information for a paragraph.
type ParagraphStyleProperties struct{ WProperties *wml.CT_PPrGeneral }

// Fonts allows manipulating a style or run's fonts.
type Fonts struct{ WFonts *wml.CT_Fonts }

// Runs returns all of the runs in a paragraph.
func (_ecfd Paragraph) Runs() []Run {
	_eacf := []Run{}
	for _, _fbfb := range _ecfd.WParagraph.EG_PContent {
		for _, _eggg := range _fbfb.EG_ContentRunContent {
			if _eggg.R != nil {
				_eacf = append(_eacf, Run{_ecfd.Document, _eggg.R})
			}
			if _eggg.Sdt != nil && _eggg.Sdt.SdtContent != nil {
				for _, _ecged := range _eggg.Sdt.SdtContent.EG_ContentRunContent {
					if _ecged.R != nil {
						_eacf = append(_eacf, Run{_ecfd.Document, _ecged.R})
					}
				}
			}
		}
	}
	return _eacf
}

// DrawingAnchored returns a slice of AnchoredDrawings.
func (_babgc Run) DrawingAnchored() []AnchoredDrawing {
	_ccfc := []AnchoredDrawing{}
	for _, _fagdc := range _babgc.WRun.EG_RunInnerContent {
		if _fagdc.Drawing == nil {
			continue
		}
		for _, _edde := range _fagdc.Drawing.Anchor {
			_ccfc = append(_ccfc, AnchoredDrawing{_babgc.Document, _edde})
		}
	}
	return _ccfc
}
func _bfbf(_egff *wml.CT_P, _ddce map[string]string) {
	for _, _dfdg := range _egff.EG_PContent {
		for _, _bbbd := range _dfdg.EG_ContentRunContent {
			if _bbbd.R != nil {
				for _, _bbgd := range _bbbd.R.EG_RunInnerContent {
					_ebdd := _bbgd.Drawing
					if _ebdd != nil {
						for _, _cgce := range _ebdd.Anchor {
							for _, _bebb := range _cgce.Graphic.GraphicData.Any {
								switch _bgfc := _bebb.(type) {
								case *picture.Pic:
									if _bgfc.BlipFill != nil && _bgfc.BlipFill.Blip != nil {
										_cgff(_bgfc.BlipFill.Blip, _ddce)
									}
								default:
								}
							}
						}
						for _, _ded := range _ebdd.Inline {
							for _, _deba := range _ded.Graphic.GraphicData.Any {
								switch _eade := _deba.(type) {
								case *picture.Pic:
									if _eade.BlipFill != nil && _eade.BlipFill.Blip != nil {
										_cgff(_eade.BlipFill.Blip, _ddce)
									}
								default:
								}
							}
						}
					}
				}
			}
		}
	}
}

// Font returns the name of paragraph font family.
func (_cgad ParagraphProperties) Font() string {
	if _accg := _cgad.Properties.RPr.RFonts; _accg != nil {
		if _accg.AsciiAttr != nil {
			return *_accg.AsciiAttr
		} else if _accg.HAnsiAttr != nil {
			return *_accg.HAnsiAttr
		} else if _accg.CsAttr != nil {
			return *_accg.CsAttr
		}
	}
	return ""
}

// SetName marks sets a name attribute for a FormField.
func (_cgd FormField) SetName(name string) {
	_gaea := wml.NewCT_FFName()
	_gaea.ValAttr = &name
	_cgd.WData.Name = []*wml.CT_FFName{_gaea}
}
func _gecdc(_cfda []*wml.CT_P, _aeffb *TableInfo, _eadf *DrawingInfo) []TextItem {
	_gfad := []TextItem{}
	for _, _dbe := range _cfda {
		_gfad = append(_gfad, _dafb(_dbe, nil, _aeffb, _eadf, _dbe.EG_PContent)...)
	}
	return _gfad
}

// GetDocRelTargetByID returns TargetAttr of document relationship given its IdAttr.
func (_fcaf *Document) GetDocRelTargetByID(idAttr string) string {
	for _, _gfdb := range _fcaf._fbb.X().Relationship {
		if _gfdb.IdAttr == idAttr {
			return _gfdb.TargetAttr
		}
	}
	return ""
}

// Header is a header for a document section.
type Header struct {
	Document *Document
	WHeader  *wml.Hdr
}

func checkAttributeSet(a *wml.CT_OnOff) bool { return a != nil }

// SetStyle sets the font size.
func (r RunProperties) SetStyle(style string) {
	if style == "" {
		r.WProperties.RStyle = nil
	} else {
		r.WProperties.RStyle = wml.NewCT_String()
		r.WProperties.RStyle.ValAttr = style
	}
}

// SetVerticalAlignment sets the vertical alignment of content within a table cell.
func (_egf CellProperties) SetVerticalAlignment(align wml.ST_VerticalJc) {
	if align == wml.ST_VerticalJcUnset {
		_egf.WProperties.VAlign = nil
	} else {
		_egf.WProperties.VAlign = wml.NewCT_VerticalJc()
		_egf.WProperties.VAlign.ValAttr = align
	}
}

// Index returns the index of the footer within the document.  This is used to
// form its zip packaged filename as well as to match it with its relationship
// ID.
func (_eba Footer) Index() int {
	for _fcbg, _gggd := range _eba.Document.WFooter {
		if _gggd == _eba.WFooter {
			return _fcbg
		}
	}
	return -1
}

// Footer is a footer for a document section.
type Footer struct {
	Document *Document
	WFooter  *wml.Ftr
}

// Clear clears the styes.
func (_fabg Styles) Clear() {
	_fabg.WStyles.DocDefaults = nil
	_fabg.WStyles.LatentStyles = nil
	_fabg.WStyles.Style = nil
}

// X returns the inner wrapped XML type.
func (_fdfa Row) X() *wml.CT_Row { return _fdfa.WRow }

// SetThemeShade sets the shade based off the theme color.
func (_abb Color) SetThemeShade(s uint8) {
	_bfca := fmt.Sprintf("\u002502x", s)
	_abb.WColor.ThemeShadeAttr = &_bfca
}

// SetFirstLineIndent controls the indentation of the first line in a paragraph.
func (_badf Paragraph) SetFirstLineIndent(m measurement.Distance) {
	_badf.ensurePPr()
	_ecbd := _badf.WParagraph.PPr
	if _ecbd.Ind == nil {
		_ecbd.Ind = wml.NewCT_Ind()
	}
	if m == measurement.Zero {
		_ecbd.Ind.FirstLineAttr = nil
	} else {
		_ecbd.Ind.FirstLineAttr = &sharedTypes.ST_TwipsMeasure{}
		_ecbd.Ind.FirstLineAttr.ST_UnsignedDecimalNumber = unioffice.Uint64(uint64(m / measurement.Twips))
	}
}

// CharacterSpacingValue returns the value of characters spacing in twips (1/20 of point).
func (_cddfa ParagraphProperties) CharacterSpacingValue() int64 {
	if _gfbe := _cddfa.Properties.RPr.Spacing; _gfbe != nil {
		_dbge := _gfbe.ValAttr
		if _dbge.Int64 != nil {
			return *_dbge.Int64
		}
	}
	return int64(0)
}

// SetEastAsiaTheme sets the font East Asia Theme.
func (_cgb Fonts) SetEastAsiaTheme(t wml.ST_Theme) { _cgb.WFonts.EastAsiaThemeAttr = t }

// SetAll sets all of the borders to a given value.
func (_geee TableBorders) SetAll(t wml.ST_Border, c color.Color, thickness measurement.Distance) {
	_geee.SetBottom(t, c, thickness)
	_geee.SetLeft(t, c, thickness)
	_geee.SetRight(t, c, thickness)
	_geee.SetTop(t, c, thickness)
	_geee.SetInsideHorizontal(t, c, thickness)
	_geee.SetInsideVertical(t, c, thickness)
}

// SetAllCaps sets the run to all caps.
func (_cbaa RunProperties) SetAllCaps(b bool) {
	if !b {
		_cbaa.WProperties.Caps = nil
	} else {
		_cbaa.WProperties.Caps = wml.NewCT_OnOff()
	}
}

// Borders allows manipulation of the table borders.
func (_dbaa TableProperties) Borders() TableBorders {
	if _dbaa.WProperties.TblBorders == nil {
		_dbaa.WProperties.TblBorders = wml.NewCT_TblBorders()
	}
	return TableBorders{_dbaa.WProperties.TblBorders}
}

// AddCell adds a cell to a row and returns it
func (_fadd Row) AddCell() Cell {
	_cbed := wml.NewEG_ContentCellContent()
	_fadd.WRow.EG_ContentCellContent = append(_fadd.WRow.EG_ContentCellContent, _cbed)
	_eebd := wml.NewCT_Tc()
	_cbed.Tc = append(_cbed.Tc, _eebd)
	return Cell{_fadd.Document, _eebd}
}

// SetName sets the name of the image, visible in the properties of the image
// within Word.
func (_bda AnchoredDrawing) SetName(name string) {
	_bda.WAnchoredDrawing.DocPr.NameAttr = name
	for _, _ggc := range _bda.WAnchoredDrawing.Graphic.GraphicData.Any {
		if _egd, _ba := _ggc.(*picture.Pic); _ba {
			_egd.NvPicPr.CNvPr.DescrAttr = unioffice.String(name)
		}
	}
}

// NewTableWidth returns a newly intialized TableWidth
func NewTableWidth() TableWidth { return TableWidth{wml.NewCT_TblWidth()} }

// SetText sets the text to be used in bullet mode.
func (_feab NumberingLevel) SetText(t string) {
	if t == "" {
		_feab.WLevel.LvlText = nil
	} else {
		_feab.WLevel.LvlText = wml.NewCT_LevelText()
		_feab.WLevel.LvlText.ValAttr = unioffice.String(t)
	}
}

// AddTable adds a table to the table cell.
func (_bc Cell) AddTable() Table {
	_ga := wml.NewEG_BlockLevelElts()
	_bc.WCell.EG_BlockLevelElts = append(_bc.WCell.EG_BlockLevelElts, _ga)
	_bbg := wml.NewEG_ContentBlockContent()
	_ga.EG_ContentBlockContent = append(_ga.EG_ContentBlockContent, _bbg)
	_agf := wml.NewCT_Tbl()
	_bbg.Tbl = append(_bbg.Tbl, _agf)
	return Table{_bc.Document, _agf}
}

// AddHyperlink adds a hyperlink to a document. Adding the hyperlink to a document
// and setting it on a cell is more efficient than setting hyperlinks directly
// on a cell.
func (_faeg Document) AddHyperlink(url string) common.Hyperlink { return _faeg._fbb.AddHyperlink(url) }

// SetRight sets the cell right margin
func (_bdf CellMargins) SetRight(d measurement.Distance) {
	_bdf.WMargins.Right = wml.NewCT_TblWidth()
	setTableMarginDistance(_bdf.WMargins.Right, d)
}

// SetChecked marks a FormFieldTypeCheckBox as checked or unchecked.
func (_addgb FormField) SetChecked(b bool) {
	if _addgb.WData.CheckBox == nil {
		return
	}
	if !b {
		_addgb.WData.CheckBox.Checked = nil
	} else {
		_addgb.WData.CheckBox.Checked = wml.NewCT_OnOff()
	}
}
func _bggc(_gebf string) mergeFieldInfo {
	_aaad := []string{}
	_degc := bytes.Buffer{}
	_cbac := -1
	for _ebcaa, _ebbff := range _gebf {
		switch _ebbff {
		case ' ':
			if _degc.Len() != 0 {
				_aaad = append(_aaad, _degc.String())
			}
			_degc.Reset()
		case '"':
			if _cbac != -1 {
				_aaad = append(_aaad, _gebf[_cbac+1:_ebcaa])
				_cbac = -1
			} else {
				_cbac = _ebcaa
			}
		default:
			_degc.WriteRune(_ebbff)
		}
	}
	if _degc.Len() != 0 {
		_aaad = append(_aaad, _degc.String())
	}
	_abeb := mergeFieldInfo{}
	for _ceaa := 0; _ceaa < len(_aaad)-1; _ceaa++ {
		_aada := _aaad[_ceaa]
		switch _aada {
		case "MERGEFIELD":
			_abeb._cbbg = _aaad[_ceaa+1]
			_ceaa++
		case "\u005cf":
			_abeb._dffc = _aaad[_ceaa+1]
			_ceaa++
		case "\u005cb":
			_abeb._edac = _aaad[_ceaa+1]
			_ceaa++
		case "\u005c\u002a":
			switch _aaad[_ceaa+1] {
			case "Upper":
				_abeb._dcca = true
			case "Lower":
				_abeb._dgfbe = true
			case "Caps":
				_abeb._bge = true
			case "FirstCap":
				_abeb._bdge = true
			}
			_ceaa++
		}
	}
	return _abeb
}

// UnderlineColor returns the hex color value of run underline.
func (_cgedd RunProperties) UnderlineColor() string {
	if _acde := _cgedd.WProperties.U; _acde != nil {
		_acgbd := _acde.ColorAttr
		if _acgbd != nil && _acgbd.ST_HexColorRGB != nil {
			return *_acgbd.ST_HexColorRGB
		}
	}
	return ""
}

// SetHangingIndent controls the hanging indent of the paragraph.
func (_fcga ParagraphStyleProperties) SetHangingIndent(m measurement.Distance) {
	if _fcga.WProperties.Ind == nil {
		_fcga.WProperties.Ind = wml.NewCT_Ind()
	}
	if m == measurement.Zero {
		_fcga.WProperties.Ind.HangingAttr = nil
	} else {
		_fcga.WProperties.Ind.HangingAttr = &sharedTypes.ST_TwipsMeasure{}
		_fcga.WProperties.Ind.HangingAttr.ST_UnsignedDecimalNumber = unioffice.Uint64(uint64(m / measurement.Twips))
	}
}

// DoubleStrike returns true if run is double striked.
func (_bdfee RunProperties) DoubleStrike() bool { return checkAttributeSet(_bdfee.WProperties.Dstrike) }

// ExtractText returns text from the document as a DocText object.
func (_bcdd *Document) ExtractText() *DocText {
	_ebgc := []TextItem{}
	for _, _ddda := range _bcdd.Document.Body.EG_BlockLevelElts {
		_ebgc = append(_ebgc, parseTextItemList(_ddda.EG_ContentBlockContent, nil)...)
	}
	return &DocText{Items: _ebgc}
}

// SetNumberingLevel sets the numbering level of a paragraph.  If used, then the
// NumberingDefinition must also be set via SetNumberingDefinition or
// SetNumberingDefinitionByID.
func (_gbcf Paragraph) SetNumberingLevel(listLevel int) {
	_gbcf.ensurePPr()
	if _gbcf.WParagraph.PPr.NumPr == nil {
		_gbcf.WParagraph.PPr.NumPr = wml.NewCT_NumPr()
	}
	_eege := wml.NewCT_DecimalNumber()
	_eege.ValAttr = int64(listLevel)
	_gbcf.WParagraph.PPr.NumPr.Ilvl = _eege
}

// AddDrawingInline adds an inline drawing from an ImageRef.
func (_ffdg Run) AddDrawingInline(img common.ImageRef) (InlineDrawing, error) {
	_gaac := _ffdg.newIC()
	_gaac.Drawing = wml.NewCT_Drawing()
	_cgced := wml.NewWdInline()
	_fefg := InlineDrawing{_ffdg.Document, _cgced}
	_cgced.CNvGraphicFramePr = dml.NewCT_NonVisualGraphicFrameProperties()
	_gaac.Drawing.Inline = append(_gaac.Drawing.Inline, _cgced)
	_cgced.Graphic = dml.NewGraphic()
	_cgced.Graphic.GraphicData = dml.NewCT_GraphicalObjectData()
	_cgced.Graphic.GraphicData.UriAttr = "http:\u002f/schemas.openxmlformats\u002eorg\u002fdrawingml\u002f2006\u002fpicture"
	_cgced.DistTAttr = unioffice.Uint32(0)
	_cgced.DistLAttr = unioffice.Uint32(0)
	_cgced.DistBAttr = unioffice.Uint32(0)
	_cgced.DistRAttr = unioffice.Uint32(0)
	_cgced.Extent.CxAttr = int64(float64(img.Size().X*measurement.Pixel72) / measurement.EMU)
	_cgced.Extent.CyAttr = int64(float64(img.Size().Y*measurement.Pixel72) / measurement.EMU)
	_befd := 0x7FFFFFFF & manthrand.Uint32()
	_cgced.DocPr.IdAttr = _befd
	_edbd := picture.NewPic()
	_edbd.NvPicPr.CNvPr.IdAttr = _befd
	_eedb := img.RelID()
	if _eedb == "" {
		return _fefg, errors.New("couldn\u0027t\u0020find\u0020reference\u0020to\u0020image\u0020within\u0020document\u0020relations")
	}
	_cgced.Graphic.GraphicData.Any = append(_cgced.Graphic.GraphicData.Any, _edbd)
	_edbd.BlipFill = dml.NewCT_BlipFillProperties()
	_edbd.BlipFill.Blip = dml.NewCT_Blip()
	_edbd.BlipFill.Blip.EmbedAttr = &_eedb
	_edbd.BlipFill.Stretch = dml.NewCT_StretchInfoProperties()
	_edbd.BlipFill.Stretch.FillRect = dml.NewCT_RelativeRect()
	_edbd.SpPr = dml.NewCT_ShapeProperties()
	_edbd.SpPr.Xfrm = dml.NewCT_Transform2D()
	_edbd.SpPr.Xfrm.Off = dml.NewCT_Point2D()
	_edbd.SpPr.Xfrm.Off.XAttr.ST_CoordinateUnqualified = unioffice.Int64(0)
	_edbd.SpPr.Xfrm.Off.YAttr.ST_CoordinateUnqualified = unioffice.Int64(0)
	_edbd.SpPr.Xfrm.Ext = dml.NewCT_PositiveSize2D()
	_edbd.SpPr.Xfrm.Ext.CxAttr = int64(img.Size().X * measurement.Point)
	_edbd.SpPr.Xfrm.Ext.CyAttr = int64(img.Size().Y * measurement.Point)
	_edbd.SpPr.PrstGeom = dml.NewCT_PresetGeometry2D()
	_edbd.SpPr.PrstGeom.PrstAttr = dml.ST_ShapeTypeRect
	return _fefg, nil
}

// RemoveParagraph removes a paragraph from the endnote.
func (_cggd Endnote) RemoveParagraph(p Paragraph) {
	for _, _cacf := range _cggd.content() {
		for _afec, _eddd := range _cacf.P {
			if _eddd == p.WParagraph {
				copy(_cacf.P[_afec:], _cacf.P[_afec+1:])
				_cacf.P = _cacf.P[0 : len(_cacf.P)-1]
				return
			}
		}
	}
}

// SetBold sets the run to bold.
func (_fedb RunProperties) SetBold(b bool) {
	if !b {
		_fedb.WProperties.B = nil
		_fedb.WProperties.BCs = nil
	} else {
		_fedb.WProperties.B = wml.NewCT_OnOff()
		_fedb.WProperties.BCs = wml.NewCT_OnOff()
	}
}

// SetName sets the name of the bookmark. This is the name that is used to
// reference the bookmark from hyperlinks.
func (_dgd Bookmark) SetName(name string) { _dgd.WBookmark.NameAttr = name }

// Paragraphs returns the paragraphs defined in a footnote.
func (_cdddd Footnote) Paragraphs() []Paragraph {
	_afcc := []Paragraph{}
	for _, _fbag := range _cdddd.content() {
		for _, _egcdd := range _fbag.P {
			_afcc = append(_afcc, Paragraph{_cdddd.Document, _egcdd})
		}
	}
	return _afcc
}

// SetVerticalMerge controls the vertical merging of cells.
func (_fga CellProperties) SetVerticalMerge(mergeVal wml.ST_Merge) {
	if mergeVal == wml.ST_MergeUnset {
		_fga.WProperties.VMerge = nil
	} else {
		_fga.WProperties.VMerge = wml.NewCT_VMerge()
		_fga.WProperties.VMerge.ValAttr = mergeVal
	}
}
func (_bdba *Document) validateTableCells() error {
	for _, _gffe := range _bdba.Document.Body.EG_BlockLevelElts {
		for _, _edbe := range _gffe.EG_ContentBlockContent {
			for _, _gege := range _edbe.Tbl {
				for _, _gab := range _gege.EG_ContentRowContent {
					for _, _agac := range _gab.Tr {
						_eag := false
						for _, _gef := range _agac.EG_ContentCellContent {
							_aaa := false
							for _, _afcb := range _gef.Tc {
								_eag = true
								for _, _fca := range _afcb.EG_BlockLevelElts {
									for _, _ebbg := range _fca.EG_ContentBlockContent {
										if len(_ebbg.P) > 0 {
											_aaa = true
											break
										}
									}
								}
							}
							if !_aaa {
								return errors.New("table\u0020cell\u0020must\u0020contain\u0020a\u0020paragraph")
							}
						}
						if !_eag {
							return errors.New("table\u0020row\u0020must\u0020contain\u0020a\u0020cell")
						}
					}
				}
			}
		}
	}
	return nil
}

// SetCellSpacingAuto sets the cell spacing within a table to automatic.
func (_gdgdc TableProperties) SetCellSpacingAuto() {
	_gdgdc.WProperties.TblCellSpacing = wml.NewCT_TblWidth()
	_gdgdc.WProperties.TblCellSpacing.TypeAttr = wml.ST_TblWidthAuto
}

// AddImage adds an image to the document package, returning a reference that
// can be used to add the image to a run and place it in the document contents.
func (_deb *Document) AddImage(i common.Image) (common.ImageRef, error) {
	_ebce := common.MakeImageRef(i, &_deb.DocBase, _deb._fbb)
	if i.Data == nil && i.Path == "" {
		return _ebce, errors.New("image\u0020must have\u0020data\u0020or\u0020a\u0020path")
	}
	if i.Format == "" {
		return _ebce, errors.New("image\u0020must have\u0020a\u0020valid\u0020format")
	}
	if i.Size.X == 0 || i.Size.Y == 0 {
		return _ebce, errors.New("image\u0020must\u0020have a valid\u0020size")
	}
	if i.Path != "" {
		_dcgab := tempstorage.Add(i.Path)
		if _dcgab != nil {
			return _ebce, _dcgab
		}
	}
	_deb.Images = append(_deb.Images, _ebce)
	_dgcb := fmt.Sprintf("media\u002fimage\u0025d\u002e\u0025s", len(_deb.Images), i.Format)
	_fgg := _deb._fbb.AddRelationship(_dgcb, unioffice.ImageType)
	_deb.ContentTypes.EnsureDefault("png", "image\u002fpng")
	_deb.ContentTypes.EnsureDefault("jpeg", "image\u002fjpeg")
	_deb.ContentTypes.EnsureDefault("jpg", "image\u002fjpeg")
	_deb.ContentTypes.EnsureDefault("wmf", "image\u002fx\u002dwmf")
	_deb.ContentTypes.EnsureDefault(i.Format, "image\u002f"+i.Format)
	_ebce.SetRelID(_fgg.X().IdAttr)
	_ebce.SetTarget(_dgcb)
	return _ebce, nil
}

// X returns the inner wrapped XML type.
func (_ef AnchoredDrawing) X() *wml.WdAnchor { return _ef.WAnchoredDrawing }

// TableConditionalFormatting controls the conditional formatting within a table
// style.
type TableConditionalFormatting struct{ WFormat *wml.CT_TblStylePr }

// X returns the inner wrapped XML type.
func (r RunProperties) X() *wml.CT_RPr { return r.WProperties }

// Emboss returns true if paragraph emboss is on.
func (p ParagraphProperties) Emboss() bool {
	return checkAttributeSet(p.Properties.RPr.Emboss)
}

// SetFontFamily sets the Ascii & HAnsi fonly family for a run.
func (r RunProperties) SetFontFamily(family string) {
	if r.WProperties.RFonts == nil {
		r.WProperties.RFonts = wml.NewCT_Fonts()
	}
	r.WProperties.RFonts.AsciiAttr = unioffice.String(family)
	r.WProperties.RFonts.HAnsiAttr = unioffice.String(family)
	r.WProperties.RFonts.EastAsiaAttr = unioffice.String(family)
}

// AddFootnote will create a new footnote and attach it to the Paragraph in the
// location at the end of the previous run (footnotes create their own run within
// the paragraph). The text given to the function is simply a convenience helper,
// paragraphs and runs can always be added to the text of the footnote later.
func (p Paragraph) AddFootnote(text string) Footnote {
	var _cedc int64
	if p.Document.HasFootnotes() {
		for _, _cgge := range p.Document.Footnotes() {
			if _cgge.id() > _cedc {
				_cedc = _cgge.id()
			}
		}
		_cedc++
	} else {
		_cedc = 0
		p.Document.WFootnotes = &wml.Footnotes{}
		p.Document.WFootnotes.CT_Footnotes = wml.CT_Footnotes{}
		p.Document.WFootnotes.Footnote = make([]*wml.CT_FtnEdn, 0)
	}
	_ggec := wml.NewCT_FtnEdn()
	_faca := wml.NewCT_FtnEdnRef()
	_faca.IdAttr = _cedc
	p.Document.WFootnotes.CT_Footnotes.Footnote = append(p.Document.WFootnotes.CT_Footnotes.Footnote, _ggec)
	_fgcc := p.AddRun()
	_bbffe := _fgcc.Properties()
	_bbffe.SetStyle("FootnoteAnchor")
	_fgcc.WRun.EG_RunInnerContent = []*wml.EG_RunInnerContent{wml.NewEG_RunInnerContent()}
	_fgcc.WRun.EG_RunInnerContent[0].FootnoteReference = _faca
	_dfbb := Footnote{p.Document, _ggec}
	_dfbb.WFootnote.IdAttr = _cedc
	_dfbb.WFootnote.EG_BlockLevelElts = []*wml.EG_BlockLevelElts{wml.NewEG_BlockLevelElts()}
	_deadc := _dfbb.AddParagraph()
	_deadc.Properties().SetStyle("Footnote")
	_deadc.WParagraph.PPr.RPr = wml.NewCT_ParaRPr()
	_ebaa := _deadc.AddRun()
	_ebaa.AddTab()
	_ebaa.AddText(text)
	return _dfbb
}

// SetAfter sets the spacing that comes after the paragraph.
func (p ParagraphSpacing) SetAfter(after measurement.Distance) {
	p.WSpacing.AfterAttr = &sharedTypes.ST_TwipsMeasure{}
	p.WSpacing.AfterAttr.ST_UnsignedDecimalNumber = unioffice.Uint64(uint64(after / measurement.Twips))
}

// IsFootnote returns a bool based on whether the run has a
// footnote or not. Returns both a bool as to whether it has
// a footnote as well as the ID of the footnote.
func (r Run) IsFootnote() (bool, int64) {
	if r.WRun.EG_RunInnerContent != nil {
		if r.WRun.EG_RunInnerContent[0].FootnoteReference != nil {
			return true, r.WRun.EG_RunInnerContent[0].FootnoteReference.IdAttr
		}
	}
	return false, 0
}

// X returns the inner wrapped XML type.
func (n NumberingDefinition) X() *wml.CT_AbstractNum { return n.WDefinition }

// Font returns the name of run font family.
func (r RunProperties) Font() string {
	if _edaf := r.WProperties.RFonts; _edaf != nil {
		if _edaf.AsciiAttr != nil {
			return *_edaf.AsciiAttr
		} else if _edaf.HAnsiAttr != nil {
			return *_edaf.HAnsiAttr
		} else if _edaf.CsAttr != nil {
			return *_edaf.CsAttr
		}
	}
	return ""
}

// Styles returns all styles.
func (s Styles) Styles() []Style {
	_gefb := []Style{}
	for _, _bgge := range s.WStyles.Style {
		_gefb = append(_gefb, Style{_bgge})
	}
	return _gefb
}

// AddTabStop adds a tab stop to the paragraph.  It controls the position of text when using Run.AddTab()
func (p ParagraphProperties) AddTabStop(position measurement.Distance, justificaton wml.ST_TabJc, leader wml.ST_TabTlc) {
	if p.Properties.Tabs == nil {
		p.Properties.Tabs = wml.NewCT_Tabs()
	}
	_caaaf := wml.NewCT_TabStop()
	_caaaf.LeaderAttr = leader
	_caaaf.ValAttr = justificaton
	_caaaf.PosAttr.Int64 = unioffice.Int64(int64(position / measurement.Twips))
	p.Properties.Tabs.Tab = append(p.Properties.Tabs.Tab, _caaaf)
}

// IsBold returns true if the run has been set to bold.
func (r RunProperties) IsBold() bool { return r.BoldValue() == OnOffValueOn }

func (p Paragraph) addSeparateFldChar() *wml.CT_FldChar {
	_adge := p.addFldChar()
	_adge.FldCharTypeAttr = wml.ST_FldCharTypeSeparate
	return _adge
}

// Properties returns the table properties.
func (t Table) Properties() TableProperties {
	if t.WTable.TblPr == nil {
		t.WTable.TblPr = wml.NewCT_TblPr()
	}
	return TableProperties{t.WTable.TblPr}
}

// IsChecked returns true if a FormFieldTypeCheckBox is checked.
func (f FormField) IsChecked() bool {
	if f.WData.CheckBox == nil {
		return false
	}
	if f.WData.CheckBox.Checked != nil {
		return true
	}
	return false
}

// Borders allows controlling individual cell borders.
func (c CellProperties) Borders() CellBorders {
	if c.WProperties.TcBorders == nil {
		c.WProperties.TcBorders = wml.NewCT_TcBorders()
	}
	return CellBorders{c.WProperties.TcBorders}
}

// Tables returns the tables defined in the document.
func (d *Document) Tables() []Table {
	_aabe := []Table{}
	if d.Document.Body == nil {
		return nil
	}
	for _, _add := range d.Document.Body.EG_BlockLevelElts {
		for _, _cdd := range _add.EG_ContentBlockContent {
			_aabe = append(_aabe, d.tables(_cdd)...)
		}
	}
	return _aabe
}

// SetColumnSpan sets the number of Grid Columns Spanned by the Cell.  This is used
// to give the appearance of merged cells.
func (c CellProperties) SetColumnSpan(cols int) {
	if cols == 0 {
		c.WProperties.GridSpan = nil
	} else {
		c.WProperties.GridSpan = wml.NewCT_DecimalNumber()
		c.WProperties.GridSpan.ValAttr = int64(cols)
	}
}

// X returns the inner wrapped XML type.
func (t Table) X() *wml.CT_Tbl { return t.WTable }

// InitializeDefault constructs the default styles.
func (s Styles) InitializeDefault() {
	s.initializeDocDefaults()
	s.initializeStyleDefaults()
}

// SetYOffset sets the Y offset for an image relative to the origin.
func (a AnchoredDrawing) SetYOffset(y measurement.Distance) {
	a.WAnchoredDrawing.PositionV.Choice = &wml.WdCT_PosVChoice{}
	a.WAnchoredDrawing.PositionV.Choice.PosOffset = unioffice.Int32(int32(y / measurement.EMU))
}

// SetSize sets the size of the displayed image on the page.
func (a AnchoredDrawing) SetSize(w, h measurement.Distance) {
	a.WAnchoredDrawing.Extent.CxAttr = int64(float64(w*measurement.Pixel72) / measurement.EMU)
	a.WAnchoredDrawing.Extent.CyAttr = int64(float64(h*measurement.Pixel72) / measurement.EMU)
}

// Name returns the name of the style if set.
func (s Style) Name() string {
	if s.WStyle.Name == nil {
		return ""
	}
	return s.WStyle.Name.ValAttr
}

func setTableMarginPercent(_bdg *wml.CT_TblWidth, pct float64) {
	_bdg.TypeAttr = wml.ST_TblWidthPct
	_bdg.WAttr = &wml.ST_MeasurementOrPercent{}
	_bdg.WAttr.ST_DecimalNumberOrPercent = &wml.ST_DecimalNumberOrPercent{}
	_bdg.WAttr.ST_DecimalNumberOrPercent.ST_UnqualifiedPercentage = unioffice.Int64(int64(pct * 50))
}

// Row is a row within a table within a document.
type Row struct {
	Document *Document
	WRow     *wml.CT_Row
}

// SetVerticalAlignment controls the vertical alignment of the run, this is used
// to control if text is superscript/subscript.
func (r RunProperties) SetVerticalAlignment(v sharedTypes.ST_VerticalAlignRun) {
	if v == sharedTypes.ST_VerticalAlignRunUnset {
		r.WProperties.VertAlign = nil
	} else {
		r.WProperties.VertAlign = wml.NewCT_VerticalAlignRun()
		r.WProperties.VertAlign.ValAttr = v
	}
}

// Style returns the style for a paragraph, or an empty string if it is unset.
func (p ParagraphProperties) Style() string {
	if p.Properties.PStyle != nil {
		return p.Properties.PStyle.ValAttr
	}
	return ""
}

// TableConditionalFormatting returns a conditional formatting object of a given
// type.  Calling this method repeatedly will return the same object.
func (s Style) TableConditionalFormatting(typ wml.ST_TblStyleOverrideType) TableConditionalFormatting {
	for _, _bfad := range s.WStyle.TblStylePr {
		if _bfad.TypeAttr == typ {
			return TableConditionalFormatting{_bfad}
		}
	}
	_fbecf := wml.NewCT_TblStylePr()
	_fbecf.TypeAttr = typ
	s.WStyle.TblStylePr = append(s.WStyle.TblStylePr, _fbecf)
	return TableConditionalFormatting{_fbecf}
}

// DoubleStrike returns true if paragraph is double striked.
func (p ParagraphProperties) DoubleStrike() bool {
	return checkAttributeSet(p.Properties.RPr.Dstrike)
}

// AddDefinition adds a new numbering definition.
func (n Numbering) AddDefinition() NumberingDefinition {
	_gafeea := wml.NewCT_Num()
	_eaddf := int64(1)
	for _, v := range n.Definitions() {
		if v.AbstractNumberID() >= _eaddf {
			_eaddf = v.AbstractNumberID() + 1
		}
	}
	_abge := int64(1)
	for _, _cegc := range n.X().Num {
		if _cegc.NumIdAttr >= _abge {
			_abge = _cegc.NumIdAttr + 1
		}
	}
	_gafeea.NumIdAttr = _abge
	_gafeea.AbstractNumId = wml.NewCT_DecimalNumber()
	_gafeea.AbstractNumId.ValAttr = _eaddf
	_gaeg := wml.NewCT_AbstractNum()
	_gaeg.AbstractNumIdAttr = _eaddf
	n.WNumbering.AbstractNum = append(n.WNumbering.AbstractNum, _gaeg)
	n.WNumbering.Num = append(n.WNumbering.Num, _gafeea)
	return NumberingDefinition{_gaeg}
}

// GetColor returns the color.Color object representing the run color.
func (p ParagraphProperties) GetColor() color.Color {
	if _aggad := p.Properties.RPr.Color; _aggad != nil {
		_cfdd := _aggad.ValAttr
		if _cfdd.ST_HexColorRGB != nil {
			return color.FromHex(*_cfdd.ST_HexColorRGB)
		}
	}
	return color.Color{}
}

// CellBorders are the borders for an individual
type CellBorders struct{ WBorders *wml.CT_TcBorders }

// SetLayout controls the table layout. wml.ST_TblLayoutTypeAutofit corresponds
// to "Automatically resize to fit contents" being checked, while
// wml.ST_TblLayoutTypeFixed corresponds to it being unchecked.
func (t TableProperties) SetLayout(l wml.ST_TblLayoutType) {
	if l == wml.ST_TblLayoutTypeUnset || l == wml.ST_TblLayoutTypeAutofit {
		t.WProperties.TblLayout = nil
	} else {
		t.WProperties.TblLayout = wml.NewCT_TblLayoutType()
		t.WProperties.TblLayout.TypeAttr = l
	}
}

// Styles is the document wide styles contained in styles.xml.
type Styles struct{ WStyles *wml.Styles }

// SetCalcOnExit marks if a FormField should be CalcOnExit or not.
func (f FormField) SetCalcOnExit(calcOnExit bool) {
	_efeg := wml.NewCT_OnOff()
	_efeg.ValAttr = &sharedTypes.ST_OnOff{Bool: &calcOnExit}
	f.WData.CalcOnExit = []*wml.CT_OnOff{_efeg}
}

// TableInfo is used for keep information about a table, a row and a cell where the text is located.
type TableInfo struct {
	Table    *wml.CT_Tbl
	Row      *wml.CT_Row
	Cell     *wml.CT_Tc
	RowIndex int
	ColIndex int
}

// ComplexSizeValue returns the value of paragraph font size for complex fonts in points.
func (p ParagraphProperties) ComplexSizeValue() float64 {
	if _dgea := p.Properties.RPr.SzCs; _dgea != nil {
		_bbfa := _dgea.ValAttr
		if _bbfa.ST_UnsignedDecimalNumber != nil {
			return float64(*_bbfa.ST_UnsignedDecimalNumber) / 2
		}
	}
	return 0.0
}

// AddImage adds an image to the document package, returning a reference that
// can be used to add the image to a run and place it in the document contents.
func (f Footer) AddImage(i common.Image) (common.ImageRef, error) {
	var _eed common.Relationships
	for k, v := range f.Document.WFooter {
		if v == f.WFooter {
			_eed = f.Document._fcbd[k]
		}
	}
	_ggff := common.MakeImageRef(i, &f.Document.DocBase, _eed)
	if i.Data == nil && i.Path == "" {
		return _ggff, errors.New("image\u0020must have\u0020data\u0020or\u0020a\u0020path")
	}
	if i.Format == "" {
		return _ggff, errors.New("image\u0020must have\u0020a\u0020valid\u0020format")
	}
	if i.Size.X == 0 || i.Size.Y == 0 {
		return _ggff, errors.New("image\u0020must\u0020have a valid\u0020size")
	}
	f.Document.Images = append(f.Document.Images, _ggff)
	_agef := fmt.Sprintf("media\u002fimage\u0025d\u002e\u0025s", len(f.Document.Images), i.Format)
	_egea := _eed.AddRelationship(_agef, unioffice.ImageType)
	_ggff.SetRelID(_egea.X().IdAttr)
	return _ggff, nil
}

// SetTextWrapNone unsets text wrapping so the image can float on top of the
// text. When used in conjunction with X/Y Offset relative to the page it can be
// used to place a logo at the top of a page at an absolute position that
// doesn't interfere with text.
func (a AnchoredDrawing) SetTextWrapNone() {
	a.WAnchoredDrawing.Choice = &wml.WdEG_WrapTypeChoice{}
	a.WAnchoredDrawing.Choice.WrapNone = wml.NewWdCT_WrapNone()
}

func getOnOffValue(coo *wml.CT_OnOff) OnOffValue {
	if coo == nil {
		return OnOffValueUnset
	}
	if coo.ValAttr != nil && coo.ValAttr.Bool != nil && !*coo.ValAttr.Bool {
		return OnOffValueOff
	}
	return OnOffValueOn
}

// SetValue sets the value of a FormFieldTypeText or FormFieldTypeDropDown. For
// FormFieldTypeDropDown, the value must be one of the fields possible values.
func (f FormField) SetValue(v string) {
	if f.WData.DdList != nil {
		for k1, v1 := range f.PossibleValues() {
			if v1 == v {
				f.WData.DdList.Result = wml.NewCT_DecimalNumber()
				f.WData.DdList.Result.ValAttr = int64(k1)
				break
			}
		}
	} else if f.WData.TextInput != nil {
		f._fdea.T = wml.NewCT_Text()
		f._fdea.T.Content = v
	}
}

func _dafb(_ggad *wml.CT_P, _fbcf *wml.CT_Hyperlink, _geea *TableInfo, _fbf *DrawingInfo, _fbbfd []*wml.EG_PContent) []TextItem {
	if len(_fbbfd) == 0 {
		return []TextItem{{Text: "", DrawingInfo: _fbf, WParagraph: _ggad, WHyperlink: _fbcf, WRun: nil, TableInfo: _geea}}
	}
	_edffd := []TextItem{}
	for _, _ccgg := range _fbbfd {
		for _, _fdbg := range _ccgg.FldSimple {
			if _fdbg != nil {
				_edffd = append(_edffd, _dafb(_ggad, _fbcf, _geea, _fbf, _fdbg.EG_PContent)...)
			}
		}
		if _bfa := _ccgg.Hyperlink; _bfa != nil {
			_edffd = append(_edffd, _abgf(_ggad, _bfa, _geea, _fbf, _bfa.EG_ContentRunContent)...)
		}
		_edffd = append(_edffd, _abgf(_ggad, nil, _geea, _fbf, _ccgg.EG_ContentRunContent)...)
	}
	return _edffd
}

// Paragraph is a paragraph within a document.
type Paragraph struct {
	Document   *Document
	WParagraph *wml.CT_P
}

// Headers returns the headers defined in the document.
func (d *Document) Headers() []Header {
	_cdc := []Header{}
	for _, _gcfb := range d.WHeader {
		_cdc = append(_cdc, Header{d, _gcfb})
	}
	return _cdc
}

// GetStyleByID returns Style by it's IdAttr.
func (d *Document) GetStyleByID(id string) Style {
	for _, _ccad := range d.Styles.WStyles.Style {
		if _ccad.StyleIdAttr != nil && *_ccad.StyleIdAttr == id {
			return Style{_ccad}
		}
	}
	return Style{}
}

// Section is the beginning of a new section.
type Section struct {
	Document *Document
	WSection *wml.CT_SectPr
}

// Properties returns the paragraph properties.
func (p Paragraph) Properties() ParagraphProperties {
	p.ensurePPr()
	return ParagraphProperties{p.Document, p.WParagraph.PPr}
}

// SetLeftPct sets the cell left margin
func (c CellMargins) SetLeftPct(pct float64) {
	c.WMargins.Left = wml.NewCT_TblWidth()
	setTableMarginPercent(c.WMargins.Left, pct)
}

// SetTop sets the top border to a specified type, color and thickness.
func (t TableBorders) SetTop(b wml.ST_Border, c color.Color, thickness measurement.Distance) {
	t.WBorders.Top = wml.NewCT_Border()
	setBorder(t.WBorders.Top, b, c, thickness)
}

// SetContextualSpacing controls whether to Ignore Spacing Above and Below When
// Using Identical Styles
func (p ParagraphStyleProperties) SetContextualSpacing(b bool) {
	if !b {
		p.WProperties.ContextualSpacing = nil
	} else {
		p.WProperties.ContextualSpacing = wml.NewCT_OnOff()
	}
}

// SetOutline sets the run to outlined text.
func (r RunProperties) SetOutline(b bool) {
	if !b {
		r.WProperties.Outline = nil
	} else {
		r.WProperties.Outline = wml.NewCT_OnOff()
	}
}

// SetTableIndent sets the Table Indent from the Leading Margin
func (t TableStyleProperties) SetTableIndent(ind measurement.Distance) {
	t.WProperties.TblInd = wml.NewCT_TblWidth()
	t.WProperties.TblInd.TypeAttr = wml.ST_TblWidthDxa
	t.WProperties.TblInd.WAttr = &wml.ST_MeasurementOrPercent{}
	t.WProperties.TblInd.WAttr.ST_DecimalNumberOrPercent = &wml.ST_DecimalNumberOrPercent{}
	t.WProperties.TblInd.WAttr.ST_DecimalNumberOrPercent.ST_UnqualifiedPercentage = unioffice.Int64(int64(ind / measurement.Dxa))
}

// Value returns the tring value of a FormFieldTypeText or FormFieldTypeDropDown.
func (f FormField) Value() string {
	if f.WData.TextInput != nil && f._fdea.T != nil {
		return f._fdea.T.Content
	} else if f.WData.DdList != nil && f.WData.DdList.Result != nil {
		_efgee := f.PossibleValues()
		_fgff := int(f.WData.DdList.Result.ValAttr)
		if _fgff < len(_efgee) {
			return _efgee[_fgff]
		}
	} else if f.WData.CheckBox != nil {
		if f.IsChecked() {
			return "true"
		}
		return "false"
	}
	return ""
}

// SetMultiLevelType sets the multilevel type.
func (n NumberingDefinition) SetMultiLevelType(t wml.ST_MultiLevelType) {
	if t == wml.ST_MultiLevelTypeUnset {
		n.WDefinition.MultiLevelType = nil
	} else {
		n.WDefinition.MultiLevelType = wml.NewCT_MultiLevelType()
		n.WDefinition.MultiLevelType.ValAttr = t
	}
}

// RightToLeft returns true if paragraph text goes from right to left.
func (p ParagraphProperties) RightToLeft() bool {
	return checkAttributeSet(p.Properties.RPr.Rtl)
}

// NewStyles constructs a new empty Styles
func NewStyles() Styles { return Styles{wml.NewStyles()} }

// X returns the inner wrapped XML type.
func (p ParagraphProperties) X() *wml.CT_PPr { return p.Properties }

// SetXOffset sets the X offset for an image relative to the origin.
func (a AnchoredDrawing) SetXOffset(x measurement.Distance) {
	a.WAnchoredDrawing.PositionH.Choice = &wml.WdCT_PosHChoice{}
	a.WAnchoredDrawing.PositionH.Choice.PosOffset = unioffice.Int32(int32(x / measurement.EMU))
}

// RowProperties are the properties for a row within a table
type RowProperties struct{ WProperties *wml.CT_TrPr }

func setBorder(cb *wml.CT_Border, sb wml.ST_Border, co color.Color, d measurement.Distance) {
	cb.ValAttr = sb
	cb.ColorAttr = &wml.ST_HexColor{}
	if co.IsAuto() {
		cb.ColorAttr.ST_HexColorAuto = wml.ST_HexColorAutoAuto
	} else {
		cb.ColorAttr.ST_HexColorRGB = co.AsRGBString()
	}
	if d != measurement.Zero {
		cb.SzAttr = unioffice.Uint64(uint64(d / measurement.Point * 8))
	}
}

// TableBorders allows manipulation of borders on a table.
type TableBorders struct{ WBorders *wml.CT_TblBorders }

// Shadow returns true if run shadow is on.
func (r RunProperties) Shadow() bool { return checkAttributeSet(r.WProperties.Shadow) }

// InsertParagraphAfter adds a new empty paragraph after the relativeTo
// paragraph.
func (d *Document) InsertParagraphAfter(relativeTo Paragraph) Paragraph {
	return d.insertParagraph(relativeTo, false)
}

// EastAsiaFont returns the name of paragraph font family for East Asia.
func (p ParagraphProperties) EastAsiaFont() string {
	if _ebf := p.Properties.RPr.RFonts; _ebf != nil {
		if _ebf.EastAsiaAttr != nil {
			return *_ebf.EastAsiaAttr
		}
	}
	return ""
}

// InsertRowBefore inserts a row before another row
func (t Table) InsertRowBefore(r Row) Row {
	for k, v := range t.WTable.EG_ContentRowContent {
		if len(v.Tr) > 0 && r.X() == v.Tr[0] {
			newcontent := wml.NewEG_ContentRowContent()
			t.WTable.EG_ContentRowContent = append(t.WTable.EG_ContentRowContent, nil)
			copy(t.WTable.EG_ContentRowContent[k+1:], t.WTable.EG_ContentRowContent[k:])
			t.WTable.EG_ContentRowContent[k] = newcontent
			newrow := wml.NewCT_Row()
			newcontent.Tr = append(newcontent.Tr, newrow)
			return Row{t.Document, newrow}
		}
	}
	return t.AddRow()
}

// SetStrikeThrough sets the run to strike-through.
func (r RunProperties) SetStrikeThrough(b bool) {
	if !b {
		r.WProperties.Strike = nil
	} else {
		r.WProperties.Strike = wml.NewCT_OnOff()
	}
}
func (p Paragraph) addStartBookmark(id int64, name string) *wml.CT_Bookmark {
	_acffc := wml.NewEG_PContent()
	p.WParagraph.EG_PContent = append(p.WParagraph.EG_PContent, _acffc)
	_cbfag := wml.NewEG_ContentRunContent()
	_efgg := wml.NewEG_RunLevelElts()
	_dcbe := wml.NewEG_RangeMarkupElements()
	wbookmark := wml.NewCT_Bookmark()
	wbookmark.NameAttr = name
	wbookmark.IdAttr = id
	_dcbe.BookmarkStart = wbookmark
	_acffc.EG_ContentRunContent = append(_acffc.EG_ContentRunContent, _cbfag)
	_cbfag.EG_RunLevelElts = append(_cbfag.EG_RunLevelElts, _efgg)
	_efgg.EG_RangeMarkupElements = append(_efgg.EG_RangeMarkupElements, _dcbe)
	return wbookmark
}

// MailMerge finds mail merge fields and replaces them with the text provided.  It also removes
// the mail merge source info from the document settings.
func (d *Document) MailMerge(mergeContent map[string]string) {
	fieldInfoList := d.mergeFields()
	runMap := map[Paragraph][]Run{}
	for _, fieldInfo := range fieldInfoList {
		_fbcfb, _fggg := mergeContent[fieldInfo._cbbg]
		if _fggg {
			if fieldInfo._dcca {
				_fbcfb = strings.ToUpper(_fbcfb)
			} else if fieldInfo._dgfbe {
				_fbcfb = strings.ToLower(_fbcfb)
			} else if fieldInfo._bge {
				_fbcfb = strings.Title(_fbcfb)
			} else if fieldInfo._bdge {
				_aeed := bytes.Buffer{}
				for _eddda, _dfba := range _fbcfb {
					if _eddda == 0 {
						_aeed.WriteRune(unicode.ToUpper(_dfba))
					} else {
						_aeed.WriteRune(_dfba)
					}
				}
				_fbcfb = _aeed.String()
			}
			if _fbcfb != "" && fieldInfo._edac != "" {
				_fbcfb = fieldInfo._edac + _fbcfb
			}
			if _fbcfb != "" && fieldInfo._dffc != "" {
				_fbcfb = _fbcfb + fieldInfo._dffc
			}
		}
		if fieldInfo._bgfa {
			if len(fieldInfo._aefd.FldSimple) == 1 && len(fieldInfo._aefd.FldSimple[0].EG_PContent) == 1 && len(fieldInfo._aefd.FldSimple[0].EG_PContent[0].EG_ContentRunContent) == 1 {
				_gcc := &wml.EG_ContentRunContent{}
				_gcc.R = fieldInfo._aefd.FldSimple[0].EG_PContent[0].EG_ContentRunContent[0].R
				fieldInfo._aefd.FldSimple = nil
				_efafb := Run{d, _gcc.R}
				_efafb.ClearContent()
				_efafb.AddText(_fbcfb)
				fieldInfo._aefd.EG_ContentRunContent = append(fieldInfo._aefd.EG_ContentRunContent, _gcc)
			}
		} else {
			_gbag := fieldInfo._aeab.Runs()
			for _cgeg := fieldInfo._ebgd; _cgeg <= fieldInfo._fdg; _cgeg++ {
				if _cgeg == fieldInfo._aabb+1 {
					_gbag[_cgeg].ClearContent()
					_gbag[_cgeg].AddText(_fbcfb)
				} else {
					runMap[fieldInfo._aeab] = append(runMap[fieldInfo._aeab], _gbag[_cgeg])
				}
			}
		}
	}
	for k1, v1 := range runMap {
		for _, v2 := range v1 {
			k1.RemoveRun(v2)
		}
	}
	d.Settings.RemoveMailMerge()
}

func setTableMarginDistance(w *wml.CT_TblWidth, _dcf measurement.Distance) {
	w.TypeAttr = wml.ST_TblWidthDxa
	w.WAttr = &wml.ST_MeasurementOrPercent{}
	w.WAttr.ST_DecimalNumberOrPercent = &wml.ST_DecimalNumberOrPercent{}
	w.WAttr.ST_DecimalNumberOrPercent.ST_UnqualifiedPercentage = unioffice.Int64(int64(_dcf / measurement.Dxa))
}

func (d *Document) InsertTableAfter(relativeTo Paragraph) Table {
	return d.insertTable(relativeTo, false)
}

// X returns the inner wml.CT_TblBorders
func (t TableBorders) X() *wml.CT_TblBorders { return t.WBorders }

// SetStartIndent controls the start indent of the paragraph.
func (p ParagraphStyleProperties) SetStartIndent(m measurement.Distance) {
	if p.WProperties.Ind == nil {
		p.WProperties.Ind = wml.NewCT_Ind()
	}
	if m == measurement.Zero {
		p.WProperties.Ind.StartAttr = nil
	} else {
		p.WProperties.Ind.StartAttr = &wml.ST_SignedTwipsMeasure{}
		p.WProperties.Ind.StartAttr.Int64 = unioffice.Int64(int64(m / measurement.Twips))
	}
}

// X returns the inner wrapped XML type.
func (t TableLook) X() *wml.CT_TblLook { return t.WTableLook }

// RemoveMailMerge removes any mail merge settings
func (s Settings) RemoveMailMerge() { s.WSettings.MailMerge = nil }

// SetHangingIndent controls the indentation of the non-first lines in a paragraph.
func (p ParagraphProperties) SetHangingIndent(m measurement.Distance) {
	if p.Properties.Ind == nil {
		p.Properties.Ind = wml.NewCT_Ind()
	}
	if m == measurement.Zero {
		p.Properties.Ind.HangingAttr = nil
	} else {
		p.Properties.Ind.HangingAttr = &sharedTypes.ST_TwipsMeasure{}
		p.Properties.Ind.HangingAttr.ST_UnsignedDecimalNumber = unioffice.Uint64(uint64(m / measurement.Twips))
	}
}

// Copy makes a deep copy of the document by saving and reading it back.
// It can be useful to avoid sharing common data between two documents.
func (d *Document) Copy() (*Document, error) {
	buffer := bytes.NewBuffer([]byte{})
	err := d.save(buffer, d.UnknownMeaning)
	if err != nil {
		return nil, err
	}
	bufferBytes := buffer.Bytes()
	reader := bytes.NewReader(bufferBytes)
	return getDocumentFromReader(reader, int64(reader.Len()), d.UnknownMeaning)
}

// UnderlineColor returns the hex color value of paragraph underline.
func (p ParagraphProperties) UnderlineColor() string {
	if u := p.Properties.RPr.U; u != nil {
		color := u.ColorAttr
		if color != nil && color.ST_HexColorRGB != nil {
			return *color.ST_HexColorRGB
		}
	}
	return ""
}

// SetTextWrapSquare sets the text wrap to square with a given wrap type.
func (a AnchoredDrawing) SetTextWrapSquare(t wml.WdST_WrapText) {
	a.WAnchoredDrawing.Choice = &wml.WdEG_WrapTypeChoice{}
	a.WAnchoredDrawing.Choice.WrapSquare = wml.NewWdCT_WrapSquare()
	a.WAnchoredDrawing.Choice.WrapSquare.WrapTextAttr = t
}

type mergeFieldInfo struct {
	_cbbg              string
	_dffc              string
	_edac              string
	_dcca              bool
	_dgfbe             bool
	_bdge              bool
	_bge               bool
	_aeab              Paragraph
	_ebgd, _aabb, _fdg int
	_aefd              *wml.EG_PContent
	_bgfa              bool
}

// AddParagraph adds a new paragraph to the document body.
func (d *Document) AddParagraph() Paragraph {
	_fbc := wml.NewEG_BlockLevelElts()
	d.Document.Body.EG_BlockLevelElts = append(d.Document.Body.EG_BlockLevelElts, _fbc)
	_cfe := wml.NewEG_ContentBlockContent()
	_fbc.EG_ContentBlockContent = append(_fbc.EG_ContentBlockContent, _cfe)
	_ffgb := wml.NewCT_P()
	_cfe.P = append(_cfe.P, _ffgb)
	return Paragraph{d, _ffgb}
}

// SetName sets the name of the style.
func (s Style) SetName(name string) {
	s.WStyle.Name = wml.NewCT_String()
	s.WStyle.Name.ValAttr = name
}

// SetStrict is a shortcut for document.SetConformance,
// as one of these values from gitee.com/greatmusicians/unioffice/schema/soo/ofc/sharedTypes:
// ST_ConformanceClassUnset, ST_ConformanceClassStrict or ST_ConformanceClassTransitional.
func (d Document) SetStrict(strict bool) {
	if strict {
		d.Document.ConformanceAttr = sharedTypes.ST_ConformanceClassStrict
	} else {
		d.Document.ConformanceAttr = sharedTypes.ST_ConformanceClassTransitional
	}
}

// CellProperties are a table cells properties within a document.
type CellProperties struct{ WProperties *wml.CT_TcPr }

func _ecd(_ddg *wml.CT_Tbl, _edd *wml.CT_P, _faa bool) *wml.CT_Tbl {
	for _, _bcg := range _ddg.EG_ContentRowContent {
		for _, _fdca := range _bcg.Tr {
			for _, _fcbb := range _fdca.EG_ContentCellContent {
				for _, _ee := range _fcbb.Tc {
					for _ece, _ebg := range _ee.EG_BlockLevelElts {
						for _, _egb := range _ebg.EG_ContentBlockContent {
							for _fgae, _defb := range _egb.P {
								if _defb == _edd {
									_ffg := wml.NewEG_BlockLevelElts()
									_caa := wml.NewEG_ContentBlockContent()
									_ffg.EG_ContentBlockContent = append(_ffg.EG_ContentBlockContent, _caa)
									_egdfb := wml.NewCT_Tbl()
									_caa.Tbl = append(_caa.Tbl, _egdfb)
									_ee.EG_BlockLevelElts = append(_ee.EG_BlockLevelElts, nil)
									if _faa {
										copy(_ee.EG_BlockLevelElts[_ece+1:], _ee.EG_BlockLevelElts[_ece:])
										_ee.EG_BlockLevelElts[_ece] = _ffg
										if _fgae != 0 {
											_dfeb := wml.NewEG_BlockLevelElts()
											_ebd := wml.NewEG_ContentBlockContent()
											_dfeb.EG_ContentBlockContent = append(_dfeb.EG_ContentBlockContent, _ebd)
											_ebd.P = _egb.P[:_fgae]
											_ee.EG_BlockLevelElts = append(_ee.EG_BlockLevelElts, nil)
											copy(_ee.EG_BlockLevelElts[_ece+1:], _ee.EG_BlockLevelElts[_ece:])
											_ee.EG_BlockLevelElts[_ece] = _dfeb
										}
										_egb.P = _egb.P[_fgae:]
									} else {
										copy(_ee.EG_BlockLevelElts[_ece+2:], _ee.EG_BlockLevelElts[_ece+1:])
										_ee.EG_BlockLevelElts[_ece+1] = _ffg
										if _fgae != len(_egb.P)-1 {
											_ebc := wml.NewEG_BlockLevelElts()
											_eddc := wml.NewEG_ContentBlockContent()
											_ebc.EG_ContentBlockContent = append(_ebc.EG_ContentBlockContent, _eddc)
											_eddc.P = _egb.P[_fgae+1:]
											_ee.EG_BlockLevelElts = append(_ee.EG_BlockLevelElts, nil)
											copy(_ee.EG_BlockLevelElts[_ece+3:], _ee.EG_BlockLevelElts[_ece+2:])
											_ee.EG_BlockLevelElts[_ece+2] = _ebc
										} else {
											_ddf := wml.NewEG_BlockLevelElts()
											_ddae := wml.NewEG_ContentBlockContent()
											_ddf.EG_ContentBlockContent = append(_ddf.EG_ContentBlockContent, _ddae)
											_ddae.P = []*wml.CT_P{wml.NewCT_P()}
											_ee.EG_BlockLevelElts = append(_ee.EG_BlockLevelElts, nil)
											copy(_ee.EG_BlockLevelElts[_ece+3:], _ee.EG_BlockLevelElts[_ece+2:])
											_ee.EG_BlockLevelElts[_ece+2] = _ddf
										}
										_egb.P = _egb.P[:_fgae+1]
									}
									return _egdfb
								}
							}
							for _, _fgc := range _egb.Tbl {
								_gad := _ecd(_fgc, _edd, _faa)
								if _gad != nil {
									return _gad
								}
							}
						}
					}
				}
			}
		}
	}
	return nil
}

// X returns the inner wrapped XML type.
func (b Bookmark) X() *wml.CT_Bookmark { return b.WBookmark }

// Footnotes returns the footnotes defined in the document.
func (d *Document) Footnotes() []Footnote {
	result := []Footnote{}
	for _, v := range d.WFootnotes.CT_Footnotes.Footnote {
		result = append(result, Footnote{d, v})
	}
	return result
}

// InsertRunAfter inserts a run in the paragraph after the relative run.
func (p Paragraph) InsertRunAfter(relativeTo Run) Run { return p.insertRun(relativeTo, false) }

// SetHorizontalBanding controls the conditional formatting for horizontal banding.
func (t TableLook) SetHorizontalBanding(on bool) {
	if !on {
		t.WTableLook.NoHBandAttr = &sharedTypes.ST_OnOff{}
		t.WTableLook.NoHBandAttr.ST_OnOff1 = sharedTypes.ST_OnOff1On
	} else {
		t.WTableLook.NoHBandAttr = &sharedTypes.ST_OnOff{}
		t.WTableLook.NoHBandAttr.ST_OnOff1 = sharedTypes.ST_OnOff1Off
	}
}

// CharacterSpacingValue returns the value of run's characters spacing in twips (1/20 of point).
func (r RunProperties) CharacterSpacingValue() int64 {
	if s := r.WProperties.Spacing; s != nil {
		value := s.ValAttr
		if value.Int64 != nil {
			return *value.Int64
		}
	}
	return int64(0)
}

// ExtractFromHeader returns text from the document header as an array of TextItems.
func ExtractFromHeader(header *wml.Hdr) []TextItem {
	return parseTextItemList(header.EG_ContentBlockContent, nil)
}

// AddSection adds a new document section with an optional section break.  If t
// is ST_SectionMarkUnset, then no break will be inserted.
func (p ParagraphProperties) AddSection(t wml.ST_SectionMark) Section {
	p.Properties.SectPr = wml.NewCT_SectPr()
	if t != wml.ST_SectionMarkUnset {
		p.Properties.SectPr.Type = wml.NewCT_SectType()
		p.Properties.SectPr.Type.ValAttr = t
	}
	return Section{p.Document, p.Properties.SectPr}
}

// Rows returns the rows defined in the table.
func (t Table) Rows() []Row {
	result := []Row{}
	for _, v1 := range t.WTable.EG_ContentRowContent {
		for _, v2 := range v1.Tr {
			result = append(result, Row{t.Document, v2})
		}
		if v1.Sdt != nil && v1.Sdt.SdtContent != nil {
			for _, v2 := range v1.Sdt.SdtContent.Tr {
				result = append(result, Row{t.Document, v2})
			}
		}
	}
	return result
}
func (p Paragraph) addInstrText(text string) *wml.CT_Text {
	run := p.AddRun()
	wrun := run.X()
	content := wml.NewEG_RunInnerContent()
	wtext := wml.NewCT_Text()
	_fffaf := "preserve"
	wtext.SpaceAttr = &_fffaf
	wtext.Content = " " + text + " "
	content.InstrText = wtext
	wrun.EG_RunInnerContent = append(wrun.EG_RunInnerContent, content)
	return wtext
}

// SetNextStyle sets the style that the next paragraph will use.
func (s Style) SetNextStyle(name string) {
	if name == "" {
		s.WStyle.Next = nil
	} else {
		s.WStyle.Next = wml.NewCT_String()
		s.WStyle.Next.ValAttr = name
	}
}

// Style is a style within the styles.xml file.
type Style struct{ WStyle *wml.CT_Style }

// Endnote is an individual endnote reference within the document.
type Endnote struct {
	Document *Document
	WEndnote *wml.CT_FtnEdn
}

// AddRun adds a run of text to a hyperlink. This is the text that will be linked.
func (h HyperLink) AddRun() Run {
	text := wml.NewEG_ContentRunContent()
	h.WHyperLink.EG_ContentRunContent = append(h.WHyperLink.EG_ContentRunContent, text)
	run := wml.NewCT_R()
	text.R = run
	return Run{h.Document, run}
}

// Type returns the type of the style.
func (s Style) Type() wml.ST_StyleType { return s.WStyle.TypeAttr }

// Paragraphs returns the paragraphs defined in a footer.
func (f Footer) Paragraphs() []Paragraph {
	result := []Paragraph{}
	for _, v1 := range f.WFooter.EG_ContentBlockContent {
		for _, v2 := range v1.P {
			result = append(result, Paragraph{f.Document, v2})
		}
	}
	for _, v1 := range f.Tables() {
		for _, v2 := range v1.Rows() {
			for _, v3 := range v2.Cells() {
				result = append(result, v3.Paragraphs()...)
			}
		}
	}
	return result
}

// X returns the inner wrapped XML type.
func (r Run) X() *wml.CT_R { return r.WRun }

// SetOutlineLevel sets the outline level of this style.
func (p ParagraphStyleProperties) SetOutlineLevel(lvl int) {
	p.WProperties.OutlineLvl = wml.NewCT_DecimalNumber()
	p.WProperties.OutlineLvl.ValAttr = int64(lvl)
}

// SetWidthAuto sets the the table width to automatic.
func (t TableProperties) SetWidthAuto() {
	t.WProperties.TblW = wml.NewCT_TblWidth()
	t.WProperties.TblW.TypeAttr = wml.ST_TblWidthAuto
}

func _fdfb(_dcfg *wml.CT_Tbl, _baec map[string]string) {
	for _, _efae := range _dcfg.EG_ContentRowContent {
		for _, _eagd := range _efae.Tr {
			for _, _edae := range _eagd.EG_ContentCellContent {
				for _, _bada := range _edae.Tc {
					for _, _cfgdg := range _bada.EG_BlockLevelElts {
						for _, _bdbf := range _cfgdg.EG_ContentBlockContent {
							for _, _cdb := range _bdbf.P {
								_baf(_cdb, _baec)
							}
							for _, _febc := range _bdbf.Tbl {
								_fdfb(_febc, _baec)
							}
						}
					}
				}
			}
		}
	}
}

// SetNumberingDefinitionByID sets the numbering definition ID directly, which must
// match an ID defined in numbering.xml
func (p Paragraph) SetNumberingDefinitionByID(abstractNumberID int64) {
	p.ensurePPr()
	if p.WParagraph.PPr.NumPr == nil {
		p.WParagraph.PPr.NumPr = wml.NewCT_NumPr()
	}
	id := wml.NewCT_DecimalNumber()
	id.ValAttr = int64(abstractNumberID)
	p.WParagraph.PPr.NumPr.NumId = id
}

// Save writes the document to an io.Writer in the Zip package format.
func (d *Document) Save(w io.Writer) error { return d.save(w, d.UnknownMeaning) }

// SetTop sets the top border to a specified type, color and thickness.
func (c CellBorders) SetTop(t wml.ST_Border, co color.Color, thickness measurement.Distance) {
	c.WBorders.Top = wml.NewCT_Border()
	setBorder(c.WBorders.Top, t, co, thickness)
}

// Cell is a table cell within a document (not a spreadsheet)
type Cell struct {
	Document *Document
	WCell    *wml.CT_Tc
}

func (p Paragraph) insertRun(run Run, _abfa bool) Run {
	for _, v1 := range p.WParagraph.EG_PContent {
		for k2, v2 := range v1.EG_ContentRunContent {
			if v2.R == run.X() {
				newrun := wml.NewCT_R()
				v1.EG_ContentRunContent = append(v1.EG_ContentRunContent, nil)
				if _abfa {
					copy(v1.EG_ContentRunContent[k2+1:], v1.EG_ContentRunContent[k2:])
					v1.EG_ContentRunContent[k2] = wml.NewEG_ContentRunContent()
					v1.EG_ContentRunContent[k2].R = newrun
				} else {
					copy(v1.EG_ContentRunContent[k2+2:], v1.EG_ContentRunContent[k2+1:])
					v1.EG_ContentRunContent[k2+1] = wml.NewEG_ContentRunContent()
					v1.EG_ContentRunContent[k2+1].R = newrun
				}
				return Run{p.Document, newrun}
			}
			if v2.Sdt != nil && v2.Sdt.SdtContent != nil {
				for _, v3 := range v2.Sdt.SdtContent.EG_ContentRunContent {
					if v3.R == run.X() {
						newrun := wml.NewCT_R()
						v2.Sdt.SdtContent.EG_ContentRunContent = append(v2.Sdt.SdtContent.EG_ContentRunContent, nil)
						if _abfa {
							copy(v2.Sdt.SdtContent.EG_ContentRunContent[k2+1:], v2.Sdt.SdtContent.EG_ContentRunContent[k2:])
							v2.Sdt.SdtContent.EG_ContentRunContent[k2] = wml.NewEG_ContentRunContent()
							v2.Sdt.SdtContent.EG_ContentRunContent[k2].R = newrun
						} else {
							copy(v2.Sdt.SdtContent.EG_ContentRunContent[k2+2:], v2.Sdt.SdtContent.EG_ContentRunContent[k2+1:])
							v2.Sdt.SdtContent.EG_ContentRunContent[k2+1] = wml.NewEG_ContentRunContent()
							v2.Sdt.SdtContent.EG_ContentRunContent[k2+1].R = newrun
						}
						return Run{p.Document, newrun}
					}
				}
			}
		}
	}
	return p.AddRun()
}

// AddText adds tet to a run.
func (r Run) AddText(s string) {
	content := wml.NewEG_RunInnerContent()
	r.WRun.EG_RunInnerContent = append(r.WRun.EG_RunInnerContent, content)
	content.T = wml.NewCT_Text()
	if unioffice.NeedsSpacePreserve(s) {
		_febf := "preserve"
		content.T.SpaceAttr = &_febf
	}
	content.T.Content = s
}

// Properties returns the cell properties.
func (c Cell) Properties() CellProperties {
	if c.WCell.TcPr == nil {
		c.WCell.TcPr = wml.NewCT_TcPr()
	}
	return CellProperties{c.WCell.TcPr}
}

// Tables returns the tables defined in the header.
func (h Header) Tables() []Table {
	result := []Table{}
	if h.WHeader == nil {
		return nil
	}
	for _, _fbedd := range h.WHeader.EG_ContentBlockContent {
		result = append(result, h.Document.tables(_fbedd)...)
	}
	return result
}

// OnOffValue represents an on/off value that can also be unset
type OnOffValue byte

// RunProperties returns the RunProperties controlling numbering level font, etc.
func (n NumberingLevel) RunProperties() RunProperties {
	if n.WLevel.RPr == nil {
		n.WLevel.RPr = wml.NewCT_RPr()
	}
	return RunProperties{n.WLevel.RPr}
}

func (d *Document) tables(content *wml.EG_ContentBlockContent) []Table {
	result := []Table{}
	for _, v1 := range content.Tbl {
		result = append(result, Table{d, v1})
		for _, v2 := range v1.EG_ContentRowContent {
			for _, v3 := range v2.Tr {
				for _, v4 := range v3.EG_ContentCellContent {
					for _, v5 := range v4.Tc {
						for _, v6 := range v5.EG_BlockLevelElts {
							for _, v7 := range v6.EG_ContentBlockContent {
								result = append(result, d.tables(v7)...)
							}
						}
					}
				}
			}
		}
	}
	return result
}

const (
	FormFieldTypeUnknown FormFieldType = iota
	FormFieldTypeText
	FormFieldTypeCheckBox
	FormFieldTypeDropDown
)

// Margins allows controlling individual cell margins.
func (c CellProperties) Margins() CellMargins {
	if c.WProperties.TcMar == nil {
		c.WProperties.TcMar = wml.NewCT_TcMar()
	}
	return CellMargins{c.WProperties.TcMar}
}

// TableProperties are the properties for a table within a document
type TableProperties struct{ WProperties *wml.CT_TblPr }

func (d *Document) onNewRelationship(decodeMap *zippkg.DecodeMap, _edeb, _ddab string, zipFileList []*zip.File, relation *relationships.Relationship, target zippkg.Target) error {
	docType := unioffice.DocTypeDocument
	switch _ddab {
	case unioffice.OfficeDocumentType, unioffice.OfficeDocumentTypeStrict:
		d.Document = wml.NewDocument()
		decodeMap.AddTarget(_edeb, d.Document, _ddab, 0)
		decodeMap.AddTarget(zippkg.RelationsPathFor(_edeb), d._fbb.X(), _ddab, 0)
		relation.TargetAttr = unioffice.RelativeFilename(docType, target.Typ, _ddab, 0)
	case unioffice.CorePropertiesType:
		decodeMap.AddTarget(_edeb, d.CoreProperties.X(), _ddab, 0)
		relation.TargetAttr = unioffice.RelativeFilename(docType, target.Typ, _ddab, 0)
	case unioffice.CustomPropertiesType:
		decodeMap.AddTarget(_edeb, d.CustomProperties.X(), _ddab, 0)
		relation.TargetAttr = unioffice.RelativeFilename(docType, target.Typ, _ddab, 0)
	case unioffice.ExtendedPropertiesType, unioffice.ExtendedPropertiesTypeStrict:
		decodeMap.AddTarget(_edeb, d.AppProperties.X(), _ddab, 0)
		relation.TargetAttr = unioffice.RelativeFilename(docType, target.Typ, _ddab, 0)
	case unioffice.ThumbnailType, unioffice.ThumbnailTypeStrict:
		for k, v := range zipFileList {
			if v == nil {
				continue
			}
			if v.Name == _edeb {
				closer, err := v.Open()
				if err != nil {
					return fmt.Errorf("error reading thumbnail: %s", err)
				}
				d.Thumbnail, _, err = image.Decode(closer)
				closer.Close()
				if err != nil {
					return fmt.Errorf("error decoding thumbnail: %s", err)
				}
				zipFileList[k] = nil
			}
		}
	case unioffice.SettingsType, unioffice.SettingsTypeStrict:
		decodeMap.AddTarget(_edeb, d.Settings.X(), _ddab, 0)
		relation.TargetAttr = unioffice.RelativeFilename(docType, target.Typ, _ddab, 0)
	case unioffice.NumberingType, unioffice.NumberingTypeStrict:
		d.Numbering = NewNumbering()
		decodeMap.AddTarget(_edeb, d.Numbering.X(), _ddab, 0)
		relation.TargetAttr = unioffice.RelativeFilename(docType, target.Typ, _ddab, 0)
	case unioffice.StylesType, unioffice.StylesTypeStrict:
		d.Styles.Clear()
		decodeMap.AddTarget(_edeb, d.Styles.X(), _ddab, 0)
		relation.TargetAttr = unioffice.RelativeFilename(docType, target.Typ, _ddab, 0)
	case unioffice.HeaderType, unioffice.HeaderTypeStrict:
		_afd := wml.NewHdr()
		decodeMap.AddTarget(_edeb, _afd, _ddab, uint32(len(d.WHeader)))
		d.WHeader = append(d.WHeader, _afd)
		relation.TargetAttr = unioffice.RelativeFilename(docType, target.Typ, _ddab, len(d.WHeader))
		_feda := common.NewRelationships()
		decodeMap.AddTarget(zippkg.RelationsPathFor(_edeb), _feda.X(), _ddab, 0)
		d._ddc = append(d._ddc, _feda)
	case unioffice.FooterType, unioffice.FooterTypeStrict:
		_bcgdb := wml.NewFtr()
		decodeMap.AddTarget(_edeb, _bcgdb, _ddab, uint32(len(d.WFooter)))
		d.WFooter = append(d.WFooter, _bcgdb)
		relation.TargetAttr = unioffice.RelativeFilename(docType, target.Typ, _ddab, len(d.WFooter))
		_accdg := common.NewRelationships()
		decodeMap.AddTarget(zippkg.RelationsPathFor(_edeb), _accdg.X(), _ddab, 0)
		d._fcbd = append(d._fcbd, _accdg)
	case unioffice.ThemeType, unioffice.ThemeTypeStrict:
		_ebbe := dml.NewTheme()
		decodeMap.AddTarget(_edeb, _ebbe, _ddab, uint32(len(d.DTheme)))
		d.DTheme = append(d.DTheme, _ebbe)
		relation.TargetAttr = unioffice.RelativeFilename(docType, target.Typ, _ddab, len(d.DTheme))
	case unioffice.WebSettingsType, unioffice.WebSettingsTypeStrict:
		d.WWebSettings = wml.NewWebSettings()
		decodeMap.AddTarget(_edeb, d.WWebSettings, _ddab, 0)
		relation.TargetAttr = unioffice.RelativeFilename(docType, target.Typ, _ddab, 0)
	case unioffice.FontTableType, unioffice.FontTableTypeStrict:
		d.WFonts = wml.NewFonts()
		decodeMap.AddTarget(_edeb, d.WFonts, _ddab, 0)
		relation.TargetAttr = unioffice.RelativeFilename(docType, target.Typ, _ddab, 0)
	case unioffice.EndNotesType, unioffice.EndNotesTypeStrict:
		d.WEndnotes = wml.NewEndnotes()
		decodeMap.AddTarget(_edeb, d.WEndnotes, _ddab, 0)
		relation.TargetAttr = unioffice.RelativeFilename(docType, target.Typ, _ddab, 0)
	case unioffice.FootNotesType, unioffice.FootNotesTypeStrict:
		d.WFootnotes = wml.NewFootnotes()
		decodeMap.AddTarget(_edeb, d.WFootnotes, _ddab, 0)
		relation.TargetAttr = unioffice.RelativeFilename(docType, target.Typ, _ddab, 0)
	case unioffice.ImageType, unioffice.ImageTypeStrict:
		var _bca common.ImageRef
		for k, v := range zipFileList {
			if v == nil {
				continue
			}
			if v.Name == _edeb {
				_bebc, err := zippkg.ExtractToDiskTmp(v, d.TmpPath)
				if err != nil {
					return err
				}
				_ecb, err := common.ImageFromStorage(_bebc)
				if err != nil {
					return err
				}
				_bca = common.MakeImageRef(_ecb, &d.DocBase, d._fbb)
				zipFileList[k] = nil
			}
		}
		_edec := "." + strings.ToLower(_bca.Format())
		relation.TargetAttr = unioffice.RelativeFilename(docType, target.Typ, _ddab, len(d.Images)+1)
		if _bad := filepath.Ext(relation.TargetAttr); _bad != _edec {
			relation.TargetAttr = relation.TargetAttr[0:len(relation.TargetAttr)-len(_bad)] + _edec
		}
		_bca.SetTarget("word/" + relation.TargetAttr)
		d.Images = append(d.Images, _bca)
	case unioffice.ControlType, unioffice.ControlTypeStrict:
		_eeg := activeX.NewOcx()
		decodeMap.AddTarget(_edeb, _eeg, _ddab, uint32(len(d.Ocx)))
		d.Ocx = append(d.Ocx, _eeg)
		relation.TargetAttr = unioffice.RelativeFilename(docType, target.Typ, _ddab, len(d.Ocx))
	default:
		unioffice.Log("unsupported relationship type: %s tgt: %s", _ddab, _edeb)
	}
	return nil
}

// PossibleValues returns the possible values for a FormFieldTypeDropDown.
func (f FormField) PossibleValues() []string {
	if f.WData.DdList == nil {
		return nil
	}
	result := []string{}
	for _, v := range f.WData.DdList.ListEntry {
		if v == nil {
			continue
		}
		result = append(result, v.ValAttr)
	}
	return result
}

// SetFooter sets a section footer.
func (s Section) SetFooter(f Footer, t wml.ST_HdrFtr) {
	r := wml.NewEG_HdrFtrReferences()
	s.WSection.EG_HdrFtrReferences = append(s.WSection.EG_HdrFtrReferences, r)
	r.FooterReference = wml.NewCT_HdrFtrRef()
	r.FooterReference.TypeAttr = t
	relationshipID := s.Document._fbb.FindRIDForN(f.Index(), unioffice.FooterType)
	if relationshipID == "" {
		log.Print("unable\u0020to\u0020determine\u0020footer ID")
	}
	r.FooterReference.IdAttr = relationshipID
}

const (
	OnOffValueUnset OnOffValue = iota
	OnOffValueOff
	OnOffValueOn
)

// X returns the inner wrapped XML type.
func (t TableConditionalFormatting) X() *wml.CT_TblStylePr { return t.WFormat }

// SetLeft sets the cell left margin
func (c CellMargins) SetLeft(d measurement.Distance) {
	c.WMargins.Left = wml.NewCT_TblWidth()
	setTableMarginDistance(c.WMargins.Left, d)
}

// Clear clears all content within a footer
func (f Footer) Clear() { f.WFooter.EG_ContentBlockContent = nil }

// Emboss returns true if run emboss is on.
func (r RunProperties) Emboss() bool { return checkAttributeSet(r.WProperties.Emboss) }

func (s Styles) initializeDocDefaults() {
	s.WStyles.DocDefaults = wml.NewCT_DocDefaults()
	s.WStyles.DocDefaults.RPrDefault = wml.NewCT_RPrDefault()
	s.WStyles.DocDefaults.RPrDefault.RPr = wml.NewCT_RPr()
	properties := RunProperties{s.WStyles.DocDefaults.RPrDefault.RPr}
	properties.SetSize(12 * measurement.Point)
	properties.Fonts().SetASCIITheme(wml.ST_ThemeMajorAscii)
	properties.Fonts().SetEastAsiaTheme(wml.ST_ThemeMajorEastAsia)
	properties.Fonts().SetHANSITheme(wml.ST_ThemeMajorHAnsi)
	properties.Fonts().SetCSTheme(wml.ST_ThemeMajorBidi)
	properties.X().Lang = wml.NewCT_Language()
	properties.X().Lang.ValAttr = unioffice.String("en\u002dUS")
	properties.X().Lang.EastAsiaAttr = unioffice.String("en\u002dUS")
	properties.X().Lang.BidiAttr = unioffice.String("ar\u002dSA")
	s.WStyles.DocDefaults.PPrDefault = wml.NewCT_PPrDefault()
}

func (f Footnote) id() int64 { return f.WFootnote.IdAttr }

// RunProperties controls run styling properties
type RunProperties struct{ WProperties *wml.CT_RPr }

// InlineDrawing is an inlined image within a run.
type InlineDrawing struct {
	Document       *Document
	WInlineDrawing *wml.WdInline
}

// X returns the inner wrapped XML type.
func (t TableWidth) X() *wml.CT_TblWidth { return t.WWidth }

// SetStyle sets the style of a paragraph and is identical to setting it on the
// paragraph's Properties()
func (p Paragraph) SetStyle(s string) {
	p.ensurePPr()
	if s == "" {
		p.WParagraph.PPr.PStyle = nil
	} else {
		p.WParagraph.PPr.PStyle = wml.NewCT_String()
		p.WParagraph.PPr.PStyle.ValAttr = s
	}
}

// SetAlignment controls the paragraph alignment
func (p ParagraphProperties) SetAlignment(align wml.ST_Jc) {
	if align == wml.ST_JcUnset {
		p.Properties.Jc = nil
	} else {
		p.Properties.Jc = wml.NewCT_Jc()
		p.Properties.Jc.ValAttr = align
	}
}

// SetStartPct sets the cell start margin
func (c CellMargins) SetStartPct(pct float64) {
	c.WMargins.Start = wml.NewCT_TblWidth()
	setTableMarginPercent(c.WMargins.Start, pct)
}

// SetUISortOrder controls the order the style is displayed in the UI.
func (s Style) SetUISortOrder(order int) {
	s.WStyle.UiPriority = wml.NewCT_DecimalNumber()
	s.WStyle.UiPriority.ValAttr = int64(order)
}

func (p Paragraph) addFldCharsForField(data, _efaeg string) FormField {
	ffdata := p.addBeginFldChar(data)
	field := FormField{WData: ffdata}
	bookmarkList := p.Document.Bookmarks()
	bookmarkCount := int64(len(bookmarkList))
	if data != "" {
		p.addStartBookmark(bookmarkCount, data)
	}
	p.addInstrText(_efaeg)
	p.addSeparateFldChar()
	if _efaeg == "FORMTEXT" {
		_cabg := p.AddRun()
		_dbbf := wml.NewEG_RunInnerContent()
		_cabg.WRun.EG_RunInnerContent = []*wml.EG_RunInnerContent{_dbbf}
		field._fdea = _dbbf
	}
	p.addEndFldChar()
	if data != "" {
		p.addEndBookmark(bookmarkCount)
	}
	return field
}

// Outline returns true if run outline is on.
func (r RunProperties) Outline() bool { return checkAttributeSet(r.WProperties.Outline) }

// SetCellSpacingPercent sets the cell spacing within a table to a percent width.
func (t TableProperties) SetCellSpacingPercent(pct float64) {
	t.WProperties.TblCellSpacing = wml.NewCT_TblWidth()
	t.WProperties.TblCellSpacing.TypeAttr = wml.ST_TblWidthPct
	t.WProperties.TblCellSpacing.WAttr = &wml.ST_MeasurementOrPercent{}
	t.WProperties.TblCellSpacing.WAttr.ST_DecimalNumberOrPercent = &wml.ST_DecimalNumberOrPercent{}
	t.WProperties.TblCellSpacing.WAttr.ST_DecimalNumberOrPercent.ST_UnqualifiedPercentage = unioffice.Int64(int64(pct * 50))
}

// Properties returns the run properties.
func (r Run) Properties() RunProperties {
	if r.WRun.RPr == nil {
		r.WRun.RPr = wml.NewCT_RPr()
	}
	return RunProperties{r.WRun.RPr}
}

// SetStartIndent controls the start indentation.
func (p ParagraphProperties) SetStartIndent(m measurement.Distance) {
	if p.Properties.Ind == nil {
		p.Properties.Ind = wml.NewCT_Ind()
	}
	if m == measurement.Zero {
		p.Properties.Ind.StartAttr = nil
	} else {
		p.Properties.Ind.StartAttr = &wml.ST_SignedTwipsMeasure{}
		p.Properties.Ind.StartAttr.Int64 = unioffice.Int64(int64(m / measurement.Twips))
	}
}

// Bookmark is a bookmarked location within a document that can be referenced
// with a hyperlink.
type Bookmark struct{ WBookmark *wml.CT_Bookmark }

// Type returns the type of the field.
func (f FormField) Type() FormFieldType {
	if f.WData.TextInput != nil {
		return FormFieldTypeText
	} else if f.WData.CheckBox != nil {
		return FormFieldTypeCheckBox
	} else if f.WData.DdList != nil {
		return FormFieldTypeDropDown
	}
	return FormFieldTypeUnknown
}

// TableWidth controls width values in table settings.
type TableWidth struct{ WWidth *wml.CT_TblWidth }

// SetFormat sets the numbering format.
func (n NumberingLevel) SetFormat(f wml.ST_NumberFormat) {
	if n.WLevel.NumFmt == nil {
		n.WLevel.NumFmt = wml.NewCT_NumFmt()
	}
	n.WLevel.NumFmt.ValAttr = f
}

// SetBottom sets the cell bottom margin
func (c CellMargins) SetBottom(d measurement.Distance) {
	c.WMargins.Bottom = wml.NewCT_TblWidth()
	setTableMarginDistance(c.WMargins.Bottom, d)
}

// SetVAlignment sets the vertical alignment for an anchored image.
func (a AnchoredDrawing) SetVAlignment(v wml.WdST_AlignV) {
	a.WAnchoredDrawing.PositionV.Choice = &wml.WdCT_PosVChoice{}
	a.WAnchoredDrawing.PositionV.Choice.Align = v
}

// FormField is a form within a document. It references the document, so changes
// to the form field wil be reflected in the document if it is saved.
type FormField struct {
	WData *wml.CT_FFData
	_fdea *wml.EG_RunInnerContent
}

func _efag(table *wml.CT_Tbl, _ccf map[string]string) {
	for _, _adgd := range table.EG_ContentRowContent {
		for _, _ebca := range _adgd.Tr {
			for _, _gace := range _ebca.EG_ContentCellContent {
				for _, _dgdeg := range _gace.Tc {
					for _, _acbb := range _dgdeg.EG_BlockLevelElts {
						for _, _egdg := range _acbb.EG_ContentBlockContent {
							for _, _dcdd := range _egdg.P {
								_bfbf(_dcdd, _ccf)
							}
							for _, _begec := range _egdg.Tbl {
								_efag(_begec, _ccf)
							}
						}
					}
				}
			}
		}
	}
}

// SetLeft sets the left border to a specified type, color and thickness.
func (t TableBorders) SetLeft(b wml.ST_Border, c color.Color, thickness measurement.Distance) {
	t.WBorders.Left = wml.NewCT_Border()
	setBorder(t.WBorders.Left, b, c, thickness)
}

// SetAlignment sets the paragraph alignment
func (n NumberingLevel) SetAlignment(j wml.ST_Jc) {
	if j == wml.ST_JcUnset {
		n.WLevel.LvlJc = nil
	} else {
		n.WLevel.LvlJc = wml.NewCT_Jc()
		n.WLevel.LvlJc.ValAttr = j
	}
}

// SetTargetByRef sets the URL target of the hyperlink and is more efficient if a link
// destination will be used many times.
func (h HyperLink) SetTargetByRef(link common.Hyperlink) {
	h.WHyperLink.IdAttr = unioffice.String(common.Relationship(link).ID())
	h.WHyperLink.AnchorAttr = nil
}

// SetStart sets the cell start margin
func (c CellMargins) SetStart(d measurement.Distance) {
	c.WMargins.Start = wml.NewCT_TblWidth()
	setTableMarginDistance(c.WMargins.Start, d)
}

// SetSmallCaps sets the run to small caps.
func (r RunProperties) SetSmallCaps(b bool) {
	if !b {
		r.WProperties.SmallCaps = nil
	} else {
		r.WProperties.SmallCaps = wml.NewCT_OnOff()
	}
}

// SetPageMargins sets the page margins for a section
func (s Section) SetPageMargins(top, right, bottom, left, header, footer, gutter measurement.Distance) {
	margin := wml.NewCT_PageMar()
	margin.TopAttr.Int64 = unioffice.Int64(int64(top / measurement.Twips))
	margin.BottomAttr.Int64 = unioffice.Int64(int64(bottom / measurement.Twips))
	margin.RightAttr.ST_UnsignedDecimalNumber = unioffice.Uint64(uint64(right / measurement.Twips))
	margin.LeftAttr.ST_UnsignedDecimalNumber = unioffice.Uint64(uint64(left / measurement.Twips))
	margin.HeaderAttr.ST_UnsignedDecimalNumber = unioffice.Uint64(uint64(header / measurement.Twips))
	margin.FooterAttr.ST_UnsignedDecimalNumber = unioffice.Uint64(uint64(footer / measurement.Twips))
	margin.GutterAttr.ST_UnsignedDecimalNumber = unioffice.Uint64(uint64(gutter / measurement.Twips))
	s.WSection.PgMar = margin
}

// X returns the inner wrapped XML type.
func (e Endnote) X() *wml.CT_FtnEdn { return e.WEndnote }

func (d *Document) save(writer io.Writer, _dcgb string) error {
	const _egdf = "document:d.Save"
	if _cbe := d.Document.Validate(); _cbe != nil {
		unioffice.Log("validation error in document: %s", _cbe)
	}
	docType := unioffice.DocTypeDocument
	if len(d.UnknownMeaning) == 0 {
		if len(_dcgb) > 0 {
			d.UnknownMeaning = _dcgb
		}
	}
	zipWriter := zip.NewWriter(writer)
	defer zipWriter.Close()
	if err := zippkg.MarshalXML(zipWriter, unioffice.BaseRelsFilename, d.Rels.X()); err != nil {
		return err
	}
	if err := zippkg.MarshalXMLByType(zipWriter, docType, unioffice.ExtendedPropertiesType, d.AppProperties.X()); err != nil {
		return err
	}
	if err := zippkg.MarshalXMLByType(zipWriter, docType, unioffice.CorePropertiesType, d.CoreProperties.X()); err != nil {
		return err
	}
	if d.CustomProperties.X() != nil {
		if err := zippkg.MarshalXMLByType(zipWriter, docType, unioffice.CustomPropertiesType, d.CustomProperties.X()); err != nil {
			return err
		}
	}
	if d.Thumbnail != nil {
		w, err := zipWriter.Create("docProps/thumbnail.jpeg")
		if err != nil {
			return err
		}
		if err := jpeg.Encode(w, d.Thumbnail, nil); err != nil {
			return err
		}
	}
	if err := zippkg.MarshalXMLByType(zipWriter, docType, unioffice.SettingsType, d.Settings.X()); err != nil {
		return err
	}
	filename := unioffice.AbsoluteFilename(docType, unioffice.OfficeDocumentType, 0)
	if err := zippkg.MarshalXML(zipWriter, filename, d.Document); err != nil {
		return err
	}
	if err := zippkg.MarshalXML(zipWriter, zippkg.RelationsPathFor(filename), d._fbb.X()); err != nil {
		return err
	}
	if d.Numbering.X() != nil {
		if err := zippkg.MarshalXMLByType(zipWriter, docType, unioffice.NumberingType, d.Numbering.X()); err != nil {
			return err
		}
	}
	if err := zippkg.MarshalXMLByType(zipWriter, docType, unioffice.StylesType, d.Styles.X()); err != nil {
		return err
	}
	if d.WWebSettings != nil {
		if err := zippkg.MarshalXMLByType(zipWriter, docType, unioffice.WebSettingsType, d.WWebSettings); err != nil {
			return err
		}
	}
	if d.WFonts != nil {
		if err := zippkg.MarshalXMLByType(zipWriter, docType, unioffice.FontTableType, d.WFonts); err != nil {
			return err
		}
	}
	if d.WEndnotes != nil {
		if err := zippkg.MarshalXMLByType(zipWriter, docType, unioffice.EndNotesType, d.WEndnotes); err != nil {
			return err
		}
	}
	if d.WFootnotes != nil {
		if err := zippkg.MarshalXMLByType(zipWriter, docType, unioffice.FootNotesType, d.WFootnotes); err != nil {
			return err
		}
	}
	for k, v := range d.DTheme {
		if err := zippkg.MarshalXMLByTypeIndex(zipWriter, docType, unioffice.ThemeType, k+1, v); err != nil {
			return err
		}
	}
	for k, v := range d.Ocx {
		if err := zippkg.MarshalXMLByTypeIndex(zipWriter, docType, unioffice.ControlType, k+1, v); err != nil {
			return err
		}
	}
	for k, v := range d.WHeader {
		fielname := unioffice.AbsoluteFilename(docType, unioffice.HeaderType, k+1)
		if err := zippkg.MarshalXML(zipWriter, fielname, v); err != nil {
			return err
		}
		if !d._ddc[k].IsEmpty() {
			zippkg.MarshalXML(zipWriter, zippkg.RelationsPathFor(fielname), d._ddc[k].X())
		}
	}
	for k, v := range d.WFooter {
		filename := unioffice.AbsoluteFilename(docType, unioffice.FooterType, k+1)
		if err := zippkg.MarshalXMLByTypeIndex(zipWriter, docType, unioffice.FooterType, k+1, v); err != nil {
			return err
		}
		if !d._fcbd[k].IsEmpty() {
			zippkg.MarshalXML(zipWriter, zippkg.RelationsPathFor(filename), d._fcbd[k].X())
		}
	}
	for k, v := range d.Images {
		if err := common.AddImageToZip(zipWriter, v, k+1, unioffice.DocTypeDocument); err != nil {
			return err
		}
	}
	if err := zippkg.MarshalXML(zipWriter, unioffice.ContentTypesFilename, d.ContentTypes.X()); err != nil {
		return err
	}
	if err := d.WriteExtraFiles(zipWriter); err != nil {
		return err
	}
	return zipWriter.Close()
}

// X returns the inner wrapped XML type.
func (h HyperLink) X() *wml.CT_Hyperlink { return h.WHyperLink }

// Close closes the document, removing any temporary files that might have been
// created when opening a document.
func (d *Document) Close() error {
	if d.TmpPath != "" {
		return tempstorage.RemoveAll(d.TmpPath)
	}
	return nil
}
func (d *Document) insertParagraph(p Paragraph, _dga bool) Paragraph {
	if d.Document.Body == nil {
		return d.AddParagraph()
	}
	_gbfa := p.X()
	for _, _dbgf := range d.Document.Body.EG_BlockLevelElts {
		for _, _ebbf := range _dbgf.EG_ContentBlockContent {
			for _gec, _gecd := range _ebbf.P {
				if _gecd == _gbfa {
					_eccd := wml.NewCT_P()
					_ebbf.P = append(_ebbf.P, nil)
					if _dga {
						copy(_ebbf.P[_gec+1:], _ebbf.P[_gec:])
						_ebbf.P[_gec] = _eccd
					} else {
						copy(_ebbf.P[_gec+2:], _ebbf.P[_gec+1:])
						_ebbf.P[_gec+1] = _eccd
					}
					return Paragraph{d, _eccd}
				}
			}
			for _, _debb := range _ebbf.Tbl {
				for _, _dag := range _debb.EG_ContentRowContent {
					for _, _ddfb := range _dag.Tr {
						for _, _dddg := range _ddfb.EG_ContentCellContent {
							for _, _ecfg := range _dddg.Tc {
								for _, _ffc := range _ecfg.EG_BlockLevelElts {
									for _, _gfgc := range _ffc.EG_ContentBlockContent {
										for _egc, _gefcd := range _gfgc.P {
											if _gefcd == _gbfa {
												_baeb := wml.NewCT_P()
												_gfgc.P = append(_gfgc.P, nil)
												if _dga {
													copy(_gfgc.P[_egc+1:], _gfgc.P[_egc:])
													_gfgc.P[_egc] = _baeb
												} else {
													copy(_gfgc.P[_egc+2:], _gfgc.P[_egc+1:])
													_gfgc.P[_egc+1] = _baeb
												}
												return Paragraph{d, _baeb}
											}
										}
									}
								}
							}
						}
					}
				}
			}
			if _ebbf.Sdt != nil && _ebbf.Sdt.SdtContent != nil && _ebbf.Sdt.SdtContent.P != nil {
				for _ggbd, _bcbd := range _ebbf.Sdt.SdtContent.P {
					if _bcbd == _gbfa {
						_fgaef := wml.NewCT_P()
						_ebbf.Sdt.SdtContent.P = append(_ebbf.Sdt.SdtContent.P, nil)
						if _dga {
							copy(_ebbf.Sdt.SdtContent.P[_ggbd+1:], _ebbf.Sdt.SdtContent.P[_ggbd:])
							_ebbf.Sdt.SdtContent.P[_ggbd] = _fgaef
						} else {
							copy(_ebbf.Sdt.SdtContent.P[_ggbd+2:], _ebbf.Sdt.SdtContent.P[_ggbd+1:])
							_ebbf.Sdt.SdtContent.P[_ggbd+1] = _fgaef
						}
						return Paragraph{d, _fgaef}
					}
				}
			}
		}
	}
	return d.AddParagraph()
}

// Bookmarks returns all of the bookmarks defined in the document.
func (d Document) Bookmarks() []Bookmark {
	if d.Document.Body == nil {
		return nil
	}
	result := []Bookmark{}
	for _, v1 := range d.Document.Body.EG_BlockLevelElts {
		for _, v2 := range v1.EG_ContentBlockContent {
			result = append(result, parseBookmarkList2(v2)...)
		}
	}
	return result
}

// Strike returns true if run is striked.
func (r RunProperties) Strike() bool { return checkAttributeSet(r.WProperties.Strike) }

// Definitions returns the defined numbering definitions.
func (n Numbering) Definitions() []NumberingDefinition {
	result := []NumberingDefinition{}
	for _, v := range n.WNumbering.AbstractNum {
		result = append(result, NumberingDefinition{v})
	}
	return result
}

// Paragraphs returns the paragraphs defined in an endnote.
func (e Endnote) Paragraphs() []Paragraph {
	result := []Paragraph{}
	for _, v1 := range e.content() {
		for _, v2 := range v1.P {
			result = append(result, Paragraph{e.Document, v2})
		}
	}
	return result
}
