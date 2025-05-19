package drawing

import (
	_f "gitee.com/greatmusicians/unioffice"
	_b "gitee.com/greatmusicians/unioffice/color"
	_fc "gitee.com/greatmusicians/unioffice/measurement"
	_fb "gitee.com/greatmusicians/unioffice/schema/soo/dml"
)

// SetAlign controls the paragraph alignment
func (_fbc ParagraphProperties) SetAlign(a _fb.ST_TextAlignType) { _fbc._ad.AlgnAttr = a }

// SetGeometry sets the shape type of the shape
func (_fbd ShapeProperties) SetGeometry(g _fb.ST_ShapeType) {
	if _fbd._fa.PrstGeom == nil {
		_fbd._fa.PrstGeom = _fb.NewCT_PresetGeometry2D()
	}
	_fbd._fa.PrstGeom.PrstAttr = g
}

// SetSize sets the font size of the run text
func (_ed RunProperties) SetSize(sz _fc.Distance) {
	_ed._eg.SzAttr = _f.Int32(int32(sz / _fc.HundredthPoint))
}

// Properties returns the paragraph properties.
func (_ge Paragraph) Properties() ParagraphProperties {
	if _ge._bbc.PPr == nil {
		_ge._bbc.PPr = _fb.NewCT_TextParagraphProperties()
	}
	return MakeParagraphProperties(_ge._bbc.PPr)
}

// SetText sets the run's text contents.
func (_ff Run) SetText(s string) {
	_ff._bea.Br = nil
	_ff._bea.Fld = nil
	if _ff._bea.R == nil {
		_ff._bea.R = _fb.NewCT_RegularTextRun()
	}
	_ff._bea.R.T = s
}
func (_bb LineProperties) SetSolidFill(c _b.Color) {
	_bb.clearFill()
	_bb._g.SolidFill = _fb.NewCT_SolidColorFillProperties()
	_bb._g.SolidFill.SrgbClr = _fb.NewCT_SRgbColor()
	_bb._g.SolidFill.SrgbClr.ValAttr = *c.AsRGBString()
}
func (_aac ShapeProperties) SetNoFill() {
	_aac.clearFill()
	_aac._fa.NoFill = _fb.NewCT_NoFillProperties()
}

// SetSolidFill controls the text color of a run.
func (_dbc RunProperties) SetSolidFill(c _b.Color) {
	_dbc._eg.NoFill = nil
	_dbc._eg.BlipFill = nil
	_dbc._eg.GradFill = nil
	_dbc._eg.GrpFill = nil
	_dbc._eg.PattFill = nil
	_dbc._eg.SolidFill = _fb.NewCT_SolidColorFillProperties()
	_dbc._eg.SolidFill.SrgbClr = _fb.NewCT_SRgbColor()
	_dbc._eg.SolidFill.SrgbClr.ValAttr = *c.AsRGBString()
}

// SetFlipHorizontal controls if the shape is flipped horizontally.
func (_bad ShapeProperties) SetFlipHorizontal(b bool) {
	_bad.ensureXfrm()
	if !b {
		_bad._fa.Xfrm.FlipHAttr = nil
	} else {
		_bad._fa.Xfrm.FlipHAttr = _f.Bool(true)
	}
}

// MakeParagraph constructs a new paragraph wrapper.
func MakeParagraph(x *_fb.CT_TextParagraph) Paragraph { return Paragraph{x} }

// SetFont controls the font of a run.
func (_gd RunProperties) SetFont(s string) {
	_gd._eg.Latin = _fb.NewCT_TextFont()
	_gd._eg.Latin.TypefaceAttr = s
}

// X returns the inner wrapped XML type.
func (_cb ParagraphProperties) X() *_fb.CT_TextParagraphProperties { return _cb._ad }

// SetFlipVertical controls if the shape is flipped vertically.
func (_bfb ShapeProperties) SetFlipVertical(b bool) {
	_bfb.ensureXfrm()
	if !b {
		_bfb._fa.Xfrm.FlipVAttr = nil
	} else {
		_bfb._fa.Xfrm.FlipVAttr = _f.Bool(true)
	}
}

// Paragraph is a paragraph within a document.
type Paragraph struct{ _bbc *_fb.CT_TextParagraph }

// Run is a run within a paragraph.
type Run struct{ _bea *_fb.EG_TextRun }

// ParagraphProperties allows controlling paragraph properties.
type ParagraphProperties struct {
	_ad *_fb.CT_TextParagraphProperties
}

// SetSize sets the width and height of the shape.
func (_da ShapeProperties) SetSize(w, h _fc.Distance) { _da.SetWidth(w); _da.SetHeight(h) }

// SetNumbered controls if bullets are numbered or not.
func (_ga ParagraphProperties) SetNumbered(scheme _fb.ST_TextAutonumberScheme) {
	if scheme == _fb.ST_TextAutonumberSchemeUnset {
		_ga._ad.BuAutoNum = nil
	} else {
		_ga._ad.BuAutoNum = _fb.NewCT_TextAutonumberBullet()
		_ga._ad.BuAutoNum.TypeAttr = scheme
	}
}

// X returns the inner wrapped XML type.
func (_ba ShapeProperties) X() *_fb.CT_ShapeProperties { return _ba._fa }

// RunProperties controls the run properties.
type RunProperties struct {
	_eg *_fb.CT_TextCharacterProperties
}

func (_ee ShapeProperties) LineProperties() LineProperties {
	if _ee._fa.Ln == nil {
		_ee._fa.Ln = _fb.NewCT_LineProperties()
	}
	return LineProperties{_ee._fa.Ln}
}

// MakeRun constructs a new Run wrapper.
func MakeRun(x *_fb.EG_TextRun) Run { return Run{x} }

const (
	LineJoinRound LineJoin = iota
	LineJoinBevel
	LineJoinMiter
)

func MakeShapeProperties(x *_fb.CT_ShapeProperties) ShapeProperties { return ShapeProperties{x} }
func (_ae ShapeProperties) SetSolidFill(c _b.Color) {
	_ae.clearFill()
	_ae._fa.SolidFill = _fb.NewCT_SolidColorFillProperties()
	_ae._fa.SolidFill.SrgbClr = _fb.NewCT_SRgbColor()
	_ae._fa.SolidFill.SrgbClr.ValAttr = *c.AsRGBString()
}

// SetBulletFont controls the font for the bullet character.
func (_db ParagraphProperties) SetBulletFont(f string) {
	if f == "" {
		_db._ad.BuFont = nil
	} else {
		_db._ad.BuFont = _fb.NewCT_TextFont()
		_db._ad.BuFont.TypefaceAttr = f
	}
}

type ShapeProperties struct{ _fa *_fb.CT_ShapeProperties }

// LineJoin is the type of line join
type LineJoin byte

// X returns the inner wrapped XML type.
func (_a LineProperties) X() *_fb.CT_LineProperties { return _a._g }

// X returns the inner wrapped XML type.
func (_dd Paragraph) X() *_fb.CT_TextParagraph { return _dd._bbc }

// GetPosition gets the position of the shape in EMU.
func (_dg ShapeProperties) GetPosition() (int64, int64) {
	_dg.ensureXfrm()
	if _dg._fa.Xfrm.Off == nil {
		_dg._fa.Xfrm.Off = _fb.NewCT_Point2D()
	}
	return *_dg._fa.Xfrm.Off.XAttr.ST_CoordinateUnqualified, *_dg._fa.Xfrm.Off.YAttr.ST_CoordinateUnqualified
}

// AddRun adds a new run to a paragraph.
func (_gb Paragraph) AddRun() Run {
	_c := MakeRun(_fb.NewEG_TextRun())
	_gb._bbc.EG_TextRun = append(_gb._bbc.EG_TextRun, _c.X())
	return _c
}

// AddBreak adds a new line break to a paragraph.
func (_be Paragraph) AddBreak() {
	_fbe := _fb.NewEG_TextRun()
	_fbe.Br = _fb.NewCT_TextLineBreak()
	_be._bbc.EG_TextRun = append(_be._bbc.EG_TextRun, _fbe)
}

// SetJoin sets the line join style.
func (_e LineProperties) SetJoin(e LineJoin) {
	_e._g.Round = nil
	_e._g.Miter = nil
	_e._g.Bevel = nil
	switch e {
	case LineJoinRound:
		_e._g.Round = _fb.NewCT_LineJoinRound()
	case LineJoinBevel:
		_e._g.Bevel = _fb.NewCT_LineJoinBevel()
	case LineJoinMiter:
		_e._g.Miter = _fb.NewCT_LineJoinMiterProperties()
	}
}

// X returns the inner wrapped XML type.
func (_ce Run) X() *_fb.EG_TextRun { return _ce._bea }

// MakeParagraphProperties constructs a new ParagraphProperties wrapper.
func MakeParagraphProperties(x *_fb.CT_TextParagraphProperties) ParagraphProperties {
	return ParagraphProperties{x}
}

// SetHeight sets the height of the shape.
func (_fcb ShapeProperties) SetHeight(h _fc.Distance) {
	_fcb.ensureXfrm()
	if _fcb._fa.Xfrm.Ext == nil {
		_fcb._fa.Xfrm.Ext = _fb.NewCT_PositiveSize2D()
	}
	_fcb._fa.Xfrm.Ext.CyAttr = int64(h / _fc.EMU)
}

// MakeRunProperties constructs a new RunProperties wrapper.
func MakeRunProperties(x *_fb.CT_TextCharacterProperties) RunProperties { return RunProperties{x} }
func (_gbb ShapeProperties) clearFill() {
	_gbb._fa.NoFill = nil
	_gbb._fa.BlipFill = nil
	_gbb._fa.GradFill = nil
	_gbb._fa.GrpFill = nil
	_gbb._fa.SolidFill = nil
	_gbb._fa.PattFill = nil
}

// SetBold controls the bolding of a run.
func (_ea RunProperties) SetBold(b bool) { _ea._eg.BAttr = _f.Bool(b) }

// SetBulletChar sets the bullet character for the paragraph.
func (_cg ParagraphProperties) SetBulletChar(c string) {
	if c == "" {
		_cg._ad.BuChar = nil
	} else {
		_cg._ad.BuChar = _fb.NewCT_TextCharBullet()
		_cg._ad.BuChar.CharAttr = c
	}
}
func (_aa LineProperties) SetNoFill() { _aa.clearFill(); _aa._g.NoFill = _fb.NewCT_NoFillProperties() }

// SetPosition sets the position of the shape.
func (_cd ShapeProperties) SetPosition(x, y _fc.Distance) {
	_cd.ensureXfrm()
	if _cd._fa.Xfrm.Off == nil {
		_cd._fa.Xfrm.Off = _fb.NewCT_Point2D()
	}
	_cd._fa.Xfrm.Off.XAttr.ST_CoordinateUnqualified = _f.Int64(int64(x / _fc.EMU))
	_cd._fa.Xfrm.Off.YAttr.ST_CoordinateUnqualified = _f.Int64(int64(y / _fc.EMU))
}

type LineProperties struct{ _g *_fb.CT_LineProperties }

// SetWidth sets the width of the shape.
func (_bff ShapeProperties) SetWidth(w _fc.Distance) {
	_bff.ensureXfrm()
	if _bff._fa.Xfrm.Ext == nil {
		_bff._fa.Xfrm.Ext = _fb.NewCT_PositiveSize2D()
	}
	_bff._fa.Xfrm.Ext.CxAttr = int64(w / _fc.EMU)
}

// SetLevel sets the level of indentation of a paragraph.
func (_eb ParagraphProperties) SetLevel(idx int32) { _eb._ad.LvlAttr = _f.Int32(idx) }

// SetWidth sets the line width, MS products treat zero as the minimum width
// that can be displayed.
func (_bg LineProperties) SetWidth(w _fc.Distance) { _bg._g.WAttr = _f.Int32(int32(w / _fc.EMU)) }

// Properties returns the run's properties.
func (_fca Run) Properties() RunProperties {
	if _fca._bea.R == nil {
		_fca._bea.R = _fb.NewCT_RegularTextRun()
	}
	if _fca._bea.R.RPr == nil {
		_fca._bea.R.RPr = _fb.NewCT_TextCharacterProperties()
	}
	return RunProperties{_fca._bea.R.RPr}
}
func (_dc ShapeProperties) ensureXfrm() {
	if _dc._fa.Xfrm == nil {
		_dc._fa.Xfrm = _fb.NewCT_Transform2D()
	}
}
func (_bf LineProperties) clearFill() {
	_bf._g.NoFill = nil
	_bf._g.GradFill = nil
	_bf._g.SolidFill = nil
	_bf._g.PattFill = nil
}
