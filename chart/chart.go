package chart

import (
	_f "fmt"
	_d "math/rand"

	_fe "gitee.com/greatmusicians/unioffice"
	_fef "gitee.com/greatmusicians/unioffice/color"
	_fega "gitee.com/greatmusicians/unioffice/drawing"
	_ee "gitee.com/greatmusicians/unioffice/measurement"
	_feg "gitee.com/greatmusicians/unioffice/schema/soo/dml"
	_e "gitee.com/greatmusicians/unioffice/schema/soo/dml/chart"
)

// AddSeries adds a default series to a Surface chart.
func (_gbbc SurfaceChart) AddSeries() SurfaceChartSeries {
	_gce := _gbbc.nextColor(len(_gbbc._gcaa.Ser))
	_gag := _e.NewCT_SurfaceSer()
	_gbbc._gcaa.Ser = append(_gbbc._gcaa.Ser, _gag)
	_gag.Idx.ValAttr = uint32(len(_gbbc._gcaa.Ser) - 1)
	_gag.Order.ValAttr = uint32(len(_gbbc._gcaa.Ser) - 1)
	_baf := SurfaceChartSeries{_gag}
	_baf.InitializeDefaults()
	_baf.Properties().LineProperties().SetSolidFill(_gce)
	return _baf
}
func (_cff CategoryAxis) SetPosition(p _e.ST_AxPos) {
	_cff._aac.AxPos = _e.NewCT_AxPos()
	_cff._aac.AxPos.ValAttr = p
}

// SetDirection changes the direction of the bar chart (bar or column).
func (_bfe Bar3DChart) SetDirection(d _e.ST_BarDir) { _bfe._daa.BarDir.ValAttr = d }

// SetValues sets values directly on a source.
func (_beae NumberDataSource) SetValues(v []float64) {
	_beae.ensureChoice()
	_beae._dce.Choice.NumRef = nil
	_beae._dce.Choice.NumLit = _e.NewCT_NumData()
	_beae._dce.Choice.NumLit.PtCount = _e.NewCT_UnsignedInt()
	_beae._dce.Choice.NumLit.PtCount.ValAttr = uint32(len(v))
	for _gdfc, _bded := range v {
		_beae._dce.Choice.NumLit.Pt = append(_beae._dce.Choice.NumLit.Pt, &_e.CT_NumVal{IdxAttr: uint32(_gdfc), V: _f.Sprintf("\u0025g", _bded)})
	}
}

// AddStockChart adds a new stock chart.
func (_cec Chart) AddStockChart() StockChart {
	_df := _e.NewCT_PlotAreaChoice()
	_cec._gdd.Chart.PlotArea.Choice = append(_cec._gdd.Chart.PlotArea.Choice, _df)
	_df.StockChart = _e.NewCT_StockChart()
	_bca := StockChart{_agc: _df.StockChart}
	_bca.InitializeDefaults()
	return _bca
}

// X returns the inner wrapped XML type.
func (_afcc LineChartSeries) X() *_e.CT_LineSer { return _afcc._afc }

// Axis is the interface implemented by different axes when assigning to a
// chart.
type Axis interface{ AxisID() uint32 }
type NumberDataSource struct{ _dce *_e.CT_NumDataSource }
type Line3DChart struct {
	chartBase
	_cbd *_e.CT_Line3DChart
}

// BubbleChartSeries is a series to be used on a Bubble chart.
type BubbleChartSeries struct{ _gd *_e.CT_BubbleSer }

// AddDateAxis adds a value axis to the chart.
func (_dgce Chart) AddDateAxis() DateAxis {
	_dcb := _e.NewCT_DateAx()
	if _dgce._gdd.Chart.PlotArea.CChoice == nil {
		_dgce._gdd.Chart.PlotArea.CChoice = _e.NewCT_PlotAreaChoice1()
	}
	_dcb.AxId = _e.NewCT_UnsignedInt()
	_dcb.AxId.ValAttr = 0x7FFFFFFF & _d.Uint32()
	_dgce._gdd.Chart.PlotArea.CChoice.DateAx = append(_dgce._gdd.Chart.PlotArea.CChoice.DateAx, _dcb)
	_dcb.Delete = _e.NewCT_Boolean()
	_dcb.Delete.ValAttr = _fe.Bool(false)
	_dcb.Scaling = _e.NewCT_Scaling()
	_dcb.Scaling.Orientation = _e.NewCT_Orientation()
	_dcb.Scaling.Orientation.ValAttr = _e.ST_OrientationMinMax
	_dcb.Choice = &_e.EG_AxSharedChoice{}
	_dcb.Choice.Crosses = _e.NewCT_Crosses()
	_dcb.Choice.Crosses.ValAttr = _e.ST_CrossesAutoZero
	_gafb := DateAxis{_dcb}
	_gafb.MajorGridLines().Properties().LineProperties().SetSolidFill(_fef.LightGray)
	_gafb.SetMajorTickMark(_e.ST_TickMarkOut)
	_gafb.SetMinorTickMark(_e.ST_TickMarkIn)
	_gafb.SetTickLabelPosition(_e.ST_TickLblPosNextTo)
	_gafb.Properties().LineProperties().SetSolidFill(_fef.Black)
	_gafb.SetPosition(_e.ST_AxPosL)
	return _gafb
}

// RadarChart is an Radar chart that has a shaded Radar underneath a curve.
type RadarChart struct {
	chartBase
	_ebcd *_e.CT_RadarChart
}

// X returns the inner wrapped XML type.
func (_aa BubbleChartSeries) X() *_e.CT_BubbleSer { return _aa._gd }

// Values returns the value data source.
func (_da AreaChartSeries) Values() NumberDataSource {
	if _da._caa.Val == nil {
		_da._caa.Val = _e.NewCT_NumDataSource()
	}
	return MakeNumberDataSource(_da._caa.Val)
}

// CategoryAxis returns the category data source.
func (_cab BarChartSeries) CategoryAxis() CategoryAxisDataSource {
	if _cab._fc.Cat == nil {
		_cab._fc.Cat = _e.NewCT_AxDataSource()
	}
	return MakeAxisDataSource(_cab._fc.Cat)
}
func (_fg Area3DChart) AddAxis(axis Axis) {
	_bfa := _e.NewCT_UnsignedInt()
	_bfa.ValAttr = axis.AxisID()
	_fg._cg.AxId = append(_fg._cg.AxId, _bfa)
}

// Properties returns the line chart series shape properties.
func (_febd SurfaceChartSeries) Properties() _fega.ShapeProperties {
	if _febd._gdda.SpPr == nil {
		_febd._gdda.SpPr = _feg.NewCT_ShapeProperties()
	}
	return _fega.MakeShapeProperties(_febd._gdda.SpPr)
}

// AddSeries adds a default series to a Scatter chart.
func (_efaf ScatterChart) AddSeries() ScatterChartSeries {
	_dgf := _efaf.nextColor(len(_efaf._dbbd.Ser))
	_bcb := _e.NewCT_ScatterSer()
	_efaf._dbbd.Ser = append(_efaf._dbbd.Ser, _bcb)
	_bcb.Idx.ValAttr = uint32(len(_efaf._dbbd.Ser) - 1)
	_bcb.Order.ValAttr = uint32(len(_efaf._dbbd.Ser) - 1)
	_gfef := ScatterChartSeries{_bcb}
	_gfef.InitializeDefaults()
	_gfef.Marker().Properties().LineProperties().SetSolidFill(_dgf)
	_gfef.Marker().Properties().SetSolidFill(_dgf)
	return _gfef
}
func (_bdag DataLabels) SetShowLegendKey(b bool) {
	_bdag.ensureChoice()
	_bdag._bga.Choice.ShowLegendKey = _e.NewCT_Boolean()
	_bdag._bga.Choice.ShowLegendKey.ValAttr = _fe.Bool(b)
}

// AddPieChart adds a new pie chart to a chart.
func (_ece Chart) AddPieChart() PieChart {
	_bec := _e.NewCT_PlotAreaChoice()
	_ece._gdd.Chart.PlotArea.Choice = append(_ece._gdd.Chart.PlotArea.Choice, _bec)
	_bec.PieChart = _e.NewCT_PieChart()
	_ebf := PieChart{_bddb: _bec.PieChart}
	_ebf.InitializeDefaults()
	return _ebf
}

type chartBase struct{}

// X returns the inner wrapped XML type.
func (_cfe Bar3DChart) X() *_e.CT_Bar3DChart { return _cfe._daa }

// InitializeDefaults the bar chart to its defaults
func (_bad DoughnutChart) InitializeDefaults() {
	_bad._gebe.VaryColors = _e.NewCT_Boolean()
	_bad._gebe.VaryColors.ValAttr = _fe.Bool(true)
	_bad._gebe.HoleSize = _e.NewCT_HoleSize()
	_bad._gebe.HoleSize.ValAttr = &_e.ST_HoleSize{}
	_bad._gebe.HoleSize.ValAttr.ST_HoleSizeUByte = _fe.Uint8(50)
}
func (_fbc Title) ParagraphProperties() _fega.ParagraphProperties {
	if _fbc._afaa.Tx == nil {
		_fbc.SetText("")
	}
	if _fbc._afaa.Tx.Choice.Rich.P[0].PPr == nil {
		_fbc._afaa.Tx.Choice.Rich.P[0].PPr = _feg.NewCT_TextParagraphProperties()
	}
	return _fega.MakeParagraphProperties(_fbc._afaa.Tx.Choice.Rich.P[0].PPr)
}

type Legend struct{ _eca *_e.CT_Legend }
type nullAxis byte

// Values returns the value data source.
func (_ddb BarChartSeries) Values() NumberDataSource {
	if _ddb._fc.Val == nil {
		_ddb._fc.Val = _e.NewCT_NumDataSource()
	}
	return MakeNumberDataSource(_ddb._fc.Val)
}
func (_fee DateAxis) SetPosition(p _e.ST_AxPos) {
	_fee._gggda.AxPos = _e.NewCT_AxPos()
	_fee._gggda.AxPos.ValAttr = p
}

// X returns the inner wrapped XML type.
func (_efb SurfaceChartSeries) X() *_e.CT_SurfaceSer { return _efb._gdda }
func (_dba ScatterChart) InitializeDefaults() {
	_dba._dbbd.ScatterStyle.ValAttr = _e.ST_ScatterStyleMarker
}

// Labels returns the data label properties.
func (_gga LineChartSeries) Labels() DataLabels {
	if _gga._afc.DLbls == nil {
		_gga._afc.DLbls = _e.NewCT_DLbls()
	}
	return MakeDataLabels(_gga._afc.DLbls)
}
func MakeChart(x *_e.ChartSpace) Chart { return Chart{x} }

// AddSeries adds a default series to an area chart.
func (_fd Area3DChart) AddSeries() AreaChartSeries {
	_bf := _fd.nextColor(len(_fd._cg.Ser))
	_ce := _e.NewCT_AreaSer()
	_fd._cg.Ser = append(_fd._cg.Ser, _ce)
	_ce.Idx.ValAttr = uint32(len(_fd._cg.Ser) - 1)
	_ce.Order.ValAttr = uint32(len(_fd._cg.Ser) - 1)
	_g := AreaChartSeries{_ce}
	_g.InitializeDefaults()
	_g.Properties().SetSolidFill(_bf)
	return _g
}

// SetIndex sets the index of the series
func (_cad LineChartSeries) SetIndex(idx uint32) { _cad._afc.Idx.ValAttr = idx }

// SetOrder sets the order of the series
func (_egbc LineChartSeries) SetOrder(idx uint32) { _egbc._afc.Order.ValAttr = idx }

// InitializeDefaults the bar chart to its defaults
func (_db Bar3DChart) InitializeDefaults() { _db.SetDirection(_e.ST_BarDirCol) }

// X returns the inner wrapped XML type.
func (_ddg GridLines) X() *_e.CT_ChartLines { return _ddg._fgcf }

// AddAxis adds an axis to a line chart.
func (_efgg LineChart) AddAxis(axis Axis) {
	_feed := _e.NewCT_UnsignedInt()
	_feed.ValAttr = axis.AxisID()
	_efgg._gafd.AxId = append(_efgg._gafd.AxId, _feed)
}

type DateAxis struct{ _gggda *_e.CT_DateAx }

func (_ebc DataLabels) ensureChoice() {
	if _ebc._bga.Choice == nil {
		_ebc._bga.Choice = _e.NewCT_DLblsChoice()
	}
}
func (_fbbe DateAxis) AxisID() uint32 { return _fbbe._gggda.AxId.ValAttr }
func (_bgd StockChart) AddAxis(axis Axis) {
	_fbbb := _e.NewCT_UnsignedInt()
	_fbbb.ValAttr = axis.AxisID()
	_bgd._agc.AxId = append(_bgd._agc.AxId, _fbbb)
}

// X returns the inner wrapped XML type.
func (_gdeb ScatterChart) X() *_e.CT_ScatterChart { return _gdeb._dbbd }

// Properties returns the bar chart series shape properties.
func (_egac RadarChartSeries) Properties() _fega.ShapeProperties {
	if _egac._eed.SpPr == nil {
		_egac._eed.SpPr = _feg.NewCT_ShapeProperties()
	}
	return _fega.MakeShapeProperties(_egac._eed.SpPr)
}

// Marker returns the marker properties.
func (_becc LineChartSeries) Marker() Marker {
	if _becc._afc.Marker == nil {
		_becc._afc.Marker = _e.NewCT_Marker()
	}
	return MakeMarker(_becc._afc.Marker)
}

// InitializeDefaults the Bubble chart to its defaults
func (_fec BubbleChart) InitializeDefaults() {}

// DoughnutChart is a Doughnut chart.
type DoughnutChart struct {
	chartBase
	_gebe *_e.CT_DoughnutChart
}

// Properties returns the chart's shape properties.
func (_gda Chart) Properties() _fega.ShapeProperties {
	if _gda._gdd.SpPr == nil {
		_gda._gdd.SpPr = _feg.NewCT_ShapeProperties()
	}
	return _fega.MakeShapeProperties(_gda._gdd.SpPr)
}
func (_feb CategoryAxis) AxisID() uint32 { return _feb._aac.AxId.ValAttr }

// X returns the inner wrapped XML type.
func (_ec BarChart) X() *_e.CT_BarChart { return _ec._fa }
func (_cb CategoryAxis) SetTickLabelPosition(p _e.ST_TickLblPos) {
	if p == _e.ST_TickLblPosUnset {
		_cb._aac.TickLblPos = nil
	} else {
		_cb._aac.TickLblPos = _e.NewCT_TickLblPos()
		_cb._aac.TickLblPos.ValAttr = p
	}
}

// PieOfPieChart is a Pie chart with an extra Pie chart.
type PieOfPieChart struct {
	chartBase
	_ggac *_e.CT_OfPieChart
}

func (_eaa SurfaceChartSeries) CategoryAxis() CategoryAxisDataSource {
	if _eaa._gdda.Cat == nil {
		_eaa._gdda.Cat = _e.NewCT_AxDataSource()
	}
	return MakeAxisDataSource(_eaa._gdda.Cat)
}
func (_ggae LineChartSeries) Values() NumberDataSource {
	if _ggae._afc.Val == nil {
		_ggae._afc.Val = _e.NewCT_NumDataSource()
	}
	return MakeNumberDataSource(_ggae._afc.Val)
}
func (_dgea LineChartSeries) CategoryAxis() CategoryAxisDataSource {
	if _dgea._afc.Cat == nil {
		_dgea._afc.Cat = _e.NewCT_AxDataSource()
	}
	return MakeAxisDataSource(_dgea._afc.Cat)
}
func MakeNumberDataSource(x *_e.CT_NumDataSource) NumberDataSource { return NumberDataSource{x} }
func (_bb chartBase) nextColor(_bee int) _fef.Color                { return _fgc[_bee%len(_fgc)] }

// X returns the inner wrapped XML type.
func (_ddf ScatterChartSeries) X() *_e.CT_ScatterSer { return _ddf._egge }

// SetOrder sets the order of the series
func (_aed SurfaceChartSeries) SetOrder(idx uint32) { _aed._gdda.Order.ValAttr = idx }

// X returns the inner wrapped XML type.
func (_dca Legend) X() *_e.CT_Legend { return _dca._eca }

// InitializeDefaults initializes a Bubble chart series to the default values.
func (_aae BubbleChartSeries) InitializeDefaults() {}

// X returns the inner wrapped XML type.
func (_cgce SurfaceChart) X() *_e.CT_SurfaceChart { return _cgce._gcaa }
func MakeDataLabels(x *_e.CT_DLbls) DataLabels    { return DataLabels{x} }
func (_bfc DateAxis) Properties() _fega.ShapeProperties {
	if _bfc._gggda.SpPr == nil {
		_bfc._gggda.SpPr = _feg.NewCT_ShapeProperties()
	}
	return _fega.MakeShapeProperties(_bfc._gggda.SpPr)
}

// SetOrder sets the order of the series
func (_dcf ScatterChartSeries) SetOrder(idx uint32) { _dcf._egge.Order.ValAttr = idx }
func (_caf ValueAxis) SetCrosses(axis Axis)         { _caf._cdf.CrossAx.ValAttr = axis.AxisID() }

// X returns the inner wrapped XML type.
func (_ggbg PieChartSeries) X() *_e.CT_PieSer { return _ggbg._faag }
func (_edb Legend) SetPosition(p _e.ST_LegendPos) {
	if p == _e.ST_LegendPosUnset {
		_edb._eca.LegendPos = nil
	} else {
		_edb._eca.LegendPos = _e.NewCT_LegendPos()
		_edb._eca.LegendPos.ValAttr = p
	}
}

// SetLabelReference is used to set the source data to a range of cells
// containing strings.
func (_ac CategoryAxisDataSource) SetLabelReference(s string) {
	_ac._dgd.Choice = _e.NewCT_AxDataSourceChoice()
	_ac._dgd.Choice.StrRef = _e.NewCT_StrRef()
	_ac._dgd.Choice.StrRef.F = s
}

// InitializeDefaults the bar chart to its defaults
func (_adcg RadarChart) InitializeDefaults() { _adcg._ebcd.RadarStyle.ValAttr = _e.ST_RadarStyleMarker }

// InitializeDefaults the bar chart to its defaults
func (_bd AreaChart) InitializeDefaults() {}
func (_acf Marker) SetSize(sz uint8) {
	_acf._gdf.Size = _e.NewCT_MarkerSize()
	_acf._gdf.Size.ValAttr = &sz
}
func MakeCategoryAxis(x *_e.CT_CatAx) CategoryAxis { return CategoryAxis{x} }

// AddBar3DChart adds a new 3D bar chart to a chart.
func (_ggc Chart) AddBar3DChart() Bar3DChart {
	_ad(_ggc._gdd.Chart)
	_egb := _e.NewCT_PlotAreaChoice()
	_ggc._gdd.Chart.PlotArea.Choice = append(_ggc._gdd.Chart.PlotArea.Choice, _egb)
	_egb.Bar3DChart = _e.NewCT_Bar3DChart()
	_egb.Bar3DChart.Grouping = _e.NewCT_BarGrouping()
	_egb.Bar3DChart.Grouping.ValAttr = _e.ST_BarGroupingStandard
	_de := Bar3DChart{_daa: _egb.Bar3DChart}
	_de.InitializeDefaults()
	return _de
}

// AddScatterChart adds a scatter (X/Y) chart.
func (_agf Chart) AddScatterChart() ScatterChart {
	_dac := _e.NewCT_PlotAreaChoice()
	_agf._gdd.Chart.PlotArea.Choice = append(_agf._gdd.Chart.PlotArea.Choice, _dac)
	_dac.ScatterChart = _e.NewCT_ScatterChart()
	_bef := ScatterChart{_dbbd: _dac.ScatterChart}
	_bef.InitializeDefaults()
	return _bef
}

// InitializeDefaults the bar chart to its defaults
func (_fdgb PieChart) InitializeDefaults() {
	_fdgb._bddb.VaryColors = _e.NewCT_Boolean()
	_fdgb._bddb.VaryColors.ValAttr = _fe.Bool(true)
}

// BarChartSeries is a series to be used on a bar chart.
type BarChartSeries struct{ _fc *_e.CT_BarSer }

// AddAxis adds an axis to a line chart.
func (_bdb Line3DChart) AddAxis(axis Axis) {
	_dcbe := _e.NewCT_UnsignedInt()
	_dcbe.ValAttr = axis.AxisID()
	_bdb._cbd.AxId = append(_bdb._cbd.AxId, _dcbe)
}
func (_abge GridLines) Properties() _fega.ShapeProperties {
	if _abge._fgcf.SpPr == nil {
		_abge._fgcf.SpPr = _feg.NewCT_ShapeProperties()
	}
	return _fega.MakeShapeProperties(_abge._fgcf.SpPr)
}
func (_dbb DataLabels) SetShowValue(b bool) {
	_dbb.ensureChoice()
	_dbb._bga.Choice.ShowVal = _e.NewCT_Boolean()
	_dbb._bga.Choice.ShowVal.ValAttr = _fe.Bool(b)
}

// AddSurface3DChart adds a new 3D surface chart to a chart.
func (_ag Chart) AddSurface3DChart() Surface3DChart {
	_ada := _e.NewCT_PlotAreaChoice()
	_ag._gdd.Chart.PlotArea.Choice = append(_ag._gdd.Chart.PlotArea.Choice, _ada)
	_ada.Surface3DChart = _e.NewCT_Surface3DChart()
	_ad(_ag._gdd.Chart)
	_adc := Surface3DChart{_abfa: _ada.Surface3DChart}
	_adc.InitializeDefaults()
	return _adc
}

// AddDoughnutChart adds a new doughnut (pie with a hole in the center) chart to a chart.
func (_ae Chart) AddDoughnutChart() DoughnutChart {
	_ffg := _e.NewCT_PlotAreaChoice()
	_ae._gdd.Chart.PlotArea.Choice = append(_ae._gdd.Chart.PlotArea.Choice, _ffg)
	_ffg.DoughnutChart = _e.NewCT_DoughnutChart()
	_gggd := DoughnutChart{_gebe: _ffg.DoughnutChart}
	_gggd.InitializeDefaults()
	return _gggd
}

// AddAxis adds an axis to a Surface chart.
func (_ccb SurfaceChart) AddAxis(axis Axis) {
	_gdcc := _e.NewCT_UnsignedInt()
	_gdcc.ValAttr = axis.AxisID()
	_ccb._gcaa.AxId = append(_ccb._gcaa.AxId, _gdcc)
}

// Properties returns the bar chart series shape properties.
func (_cfeg BarChartSeries) Properties() _fega.ShapeProperties {
	if _cfeg._fc.SpPr == nil {
		_cfeg._fc.SpPr = _feg.NewCT_ShapeProperties()
	}
	return _fega.MakeShapeProperties(_cfeg._fc.SpPr)
}
func (_acg ScatterChartSeries) InitializeDefaults() {
	_acg.Properties().LineProperties().SetNoFill()
	_acg.Marker().SetSymbol(_e.ST_MarkerStyleAuto)
	_acg.Labels().SetShowLegendKey(false)
	_acg.Labels().SetShowValue(true)
	_acg.Labels().SetShowPercent(false)
	_acg.Labels().SetShowCategoryName(false)
	_acg.Labels().SetShowSeriesName(false)
	_acg.Labels().SetShowLeaderLines(false)
}
func MakeTitle(x *_e.CT_Title) Title { return Title{x} }
func (_dfa DataLabels) SetShowSeriesName(b bool) {
	_dfa.ensureChoice()
	_dfa._bga.Choice.ShowSerName = _e.NewCT_Boolean()
	_dfa._bga.Choice.ShowSerName.ValAttr = _fe.Bool(b)
}

// StockChart is a 2D Stock chart.
type StockChart struct {
	chartBase
	_agc *_e.CT_StockChart
}

// AddPieOfPieChart adds a new pie chart to a chart.
func (_gfe Chart) AddPieOfPieChart() PieOfPieChart {
	_dda := _e.NewCT_PlotAreaChoice()
	_gfe._gdd.Chart.PlotArea.Choice = append(_gfe._gdd.Chart.PlotArea.Choice, _dda)
	_dda.OfPieChart = _e.NewCT_OfPieChart()
	_efg := PieOfPieChart{_ggac: _dda.OfPieChart}
	_efg.InitializeDefaults()
	return _efg
}

// InitializeDefaults initializes an area series to the default values.
func (_cgc AreaChartSeries) InitializeDefaults() {}

// PieChartSeries is a series to be used on an Pie chart.
type PieChartSeries struct{ _faag *_e.CT_PieSer }

// SetIndex sets the index of the series
func (_ddcg SurfaceChartSeries) SetIndex(idx uint32) { _ddcg._gdda.Idx.ValAttr = idx }

// X returns the inner wrapped XML type.
func (_ccc DateAxis) X() *_e.CT_DateAx { return _ccc._gggda }

// AddAxis adds an axis to a Surface chart.
func (_gggc Surface3DChart) AddAxis(axis Axis) {
	_gfeg := _e.NewCT_UnsignedInt()
	_gfeg.ValAttr = axis.AxisID()
	_gggc._abfa.AxId = append(_gggc._abfa.AxId, _gfeg)
}

type Marker struct{ _gdf *_e.CT_Marker }

// SetType sets the type the secone pie to either pie or bar
func (_ggf PieOfPieChart) SetType(t _e.ST_OfPieType) { _ggf._ggac.OfPieType.ValAttr = t }
func (_fdg CategoryAxis) SetMinorTickMark(m _e.ST_TickMark) {
	if m == _e.ST_TickMarkUnset {
		_fdg._aac.MinorTickMark = nil
	} else {
		_fdg._aac.MinorTickMark = _e.NewCT_TickMark()
		_fdg._aac.MinorTickMark.ValAttr = m
	}
}
func (_gggb Chart) AddSeriesAxis() SeriesAxis {
	_ega := _e.NewCT_SerAx()
	if _gggb._gdd.Chart.PlotArea.CChoice == nil {
		_gggb._gdd.Chart.PlotArea.CChoice = _e.NewCT_PlotAreaChoice1()
	}
	_ega.AxId = _e.NewCT_UnsignedInt()
	_ega.AxId.ValAttr = 0x7FFFFFFF & _d.Uint32()
	_gggb._gdd.Chart.PlotArea.CChoice.SerAx = append(_gggb._gdd.Chart.PlotArea.CChoice.SerAx, _ega)
	_ega.Delete = _e.NewCT_Boolean()
	_ega.Delete.ValAttr = _fe.Bool(false)
	_aeg := MakeSeriesAxis(_ega)
	_aeg.InitializeDefaults()
	return _aeg
}

// LineChartSeries is the data series for a line chart.
type LineChartSeries struct{ _afc *_e.CT_LineSer }

// X returns the inner wrapped XML type.
func (_dbf PieChart) X() *_e.CT_PieChart { return _dbf._bddb }

// AddLine3DChart adds a new 3D line chart to a chart.
func (_ccd Chart) AddLine3DChart() Line3DChart {
	_ad(_ccd._gdd.Chart)
	_gc := _e.NewCT_PlotAreaChoice()
	_ccd._gdd.Chart.PlotArea.Choice = append(_ccd._gdd.Chart.PlotArea.Choice, _gc)
	_gc.Line3DChart = _e.NewCT_Line3DChart()
	_gc.Line3DChart.Grouping = _e.NewCT_Grouping()
	_gc.Line3DChart.Grouping.ValAttr = _e.ST_GroupingStandard
	return Line3DChart{_cbd: _gc.Line3DChart}
}

// Index returns the index of the series
func (_fbg ScatterChartSeries) Index() uint32 { return _fbg._egge.Idx.ValAttr }

// Chart is a generic chart.
type Chart struct{ _gdd *_e.ChartSpace }

func (_bcg DataLabels) SetShowPercent(b bool) {
	_bcg.ensureChoice()
	_bcg._bga.Choice.ShowPercent = _e.NewCT_Boolean()
	_bcg._bga.Choice.ShowPercent.ValAttr = _fe.Bool(b)
}

// Pie3DChart is a Pie3D chart.
type Pie3DChart struct {
	chartBase
	_cacc *_e.CT_Pie3DChart
}

// InitializeDefaults initializes a bar chart series to the default values.
func (_daaf BarChartSeries) InitializeDefaults() {}

// AddSeries adds a default series to an Pie chart.
func (_fgg PieChart) AddSeries() PieChartSeries {
	_gba := _e.NewCT_PieSer()
	_fgg._bddb.Ser = append(_fgg._bddb.Ser, _gba)
	_gba.Idx.ValAttr = uint32(len(_fgg._bddb.Ser) - 1)
	_gba.Order.ValAttr = uint32(len(_fgg._bddb.Ser) - 1)
	_bbf := PieChartSeries{_gba}
	_bbf.InitializeDefaults()
	return _bbf
}

// Properties returns the bar chart series shape properties.
func (_fae PieChartSeries) Properties() _fega.ShapeProperties {
	if _fae._faag.SpPr == nil {
		_fae._faag.SpPr = _feg.NewCT_ShapeProperties()
	}
	return _fega.MakeShapeProperties(_fae._faag.SpPr)
}

// SetExplosion sets the value that the segements of the pie are 'exploded' by
func (_aec PieChartSeries) SetExplosion(v uint32) {
	_aec._faag.Explosion = _e.NewCT_UnsignedInt()
	_aec._faag.Explosion.ValAttr = v
}

// X returns the inner wrapped XML type.
func (_cfb RadarChart) X() *_e.CT_RadarChart { return _cfb._ebcd }
func (_adca Legend) InitializeDefaults() {
	_adca.SetPosition(_e.ST_LegendPosR)
	_adca.SetOverlay(false)
	_adca.Properties().SetNoFill()
	_adca.Properties().LineProperties().SetNoFill()
}

// X returns the inner wrapped XML type.
func (_cac Marker) X() *_e.CT_Marker { return _cac._gdf }

// AddSeries adds a default series to an Doughnut chart.
func (_gca DoughnutChart) AddSeries() PieChartSeries {
	_fca := _e.NewCT_PieSer()
	_gca._gebe.Ser = append(_gca._gebe.Ser, _fca)
	_fca.Idx.ValAttr = uint32(len(_gca._gebe.Ser) - 1)
	_fca.Order.ValAttr = uint32(len(_gca._gebe.Ser) - 1)
	_bfb := PieChartSeries{_fca}
	_bfb.InitializeDefaults()
	return _bfb
}

// Values returns the value data source.
func (_dbg BubbleChartSeries) Values() NumberDataSource {
	if _dbg._gd.YVal == nil {
		_dbg._gd.YVal = _e.NewCT_NumDataSource()
	}
	return MakeNumberDataSource(_dbg._gd.YVal)
}
func (_ggaa ValueAxis) SetTickLabelPosition(p _e.ST_TickLblPos) {
	if p == _e.ST_TickLblPosUnset {
		_ggaa._cdf.TickLblPos = nil
	} else {
		_ggaa._cdf.TickLblPos = _e.NewCT_TickLblPos()
		_ggaa._cdf.TickLblPos.ValAttr = p
	}
}

// AddRadarChart adds a new radar chart to a chart.
func (_eb Chart) AddRadarChart() RadarChart {
	_ecd := _e.NewCT_PlotAreaChoice()
	_eb._gdd.Chart.PlotArea.Choice = append(_eb._gdd.Chart.PlotArea.Choice, _ecd)
	_ecd.RadarChart = _e.NewCT_RadarChart()
	_acb := RadarChart{_ebcd: _ecd.RadarChart}
	_acb.InitializeDefaults()
	return _acb
}
func (_bbeb ScatterChartSeries) SetSmooth(b bool) {
	_bbeb._egge.Smooth = _e.NewCT_Boolean()
	_bbeb._egge.Smooth.ValAttr = &b
}

// AddLineChart adds a new line chart to a chart.
func (_dgc Chart) AddLineChart() LineChart {
	_bdae := _e.NewCT_PlotAreaChoice()
	_dgc._gdd.Chart.PlotArea.Choice = append(_dgc._gdd.Chart.PlotArea.Choice, _bdae)
	_bdae.LineChart = _e.NewCT_LineChart()
	_bdae.LineChart.Grouping = _e.NewCT_Grouping()
	_bdae.LineChart.Grouping.ValAttr = _e.ST_GroupingStandard
	return LineChart{_gafd: _bdae.LineChart}
}

// AddSeries adds a default series to a bar chart.
func (_cd Bar3DChart) AddSeries() BarChartSeries {
	_fb := _cd.nextColor(len(_cd._daa.Ser))
	_ed := _e.NewCT_BarSer()
	_cd._daa.Ser = append(_cd._daa.Ser, _ed)
	_ed.Idx.ValAttr = uint32(len(_cd._daa.Ser) - 1)
	_ed.Order.ValAttr = uint32(len(_cd._daa.Ser) - 1)
	_eeg := BarChartSeries{_ed}
	_eeg.InitializeDefaults()
	_eeg.Properties().SetSolidFill(_fb)
	return _eeg
}

// InitializeDefaults initializes an Pie series to the default values.
func (_abfb PieChartSeries) InitializeDefaults() {}
func (_acd Marker) Properties() _fega.ShapeProperties {
	if _acd._gdf.SpPr == nil {
		_acd._gdf.SpPr = _feg.NewCT_ShapeProperties()
	}
	return _fega.MakeShapeProperties(_acd._gdf.SpPr)
}
func (_bcgf LineChartSeries) InitializeDefaults() {
	_bcgf.Properties().LineProperties().SetWidth(1 * _ee.Point)
	_bcgf.Properties().LineProperties().SetSolidFill(_fef.Black)
	_bcgf.Properties().LineProperties().SetJoin(_fega.LineJoinRound)
	_bcgf.Marker().SetSymbol(_e.ST_MarkerStyleNone)
	_bcgf.Labels().SetShowLegendKey(false)
	_bcgf.Labels().SetShowValue(false)
	_bcgf.Labels().SetShowPercent(false)
	_bcgf.Labels().SetShowCategoryName(false)
	_bcgf.Labels().SetShowSeriesName(false)
	_bcgf.Labels().SetShowLeaderLines(false)
}

// SetText sets the series text.
func (_egg RadarChartSeries) SetText(s string) {
	_egg._eed.Tx = _e.NewCT_SerTx()
	_egg._eed.Tx.Choice.V = &s
}

// X returns the inner wrapped XML type.
func (_fcd SeriesAxis) X() *_e.CT_SerAx { return _fcd._faac }

// X returns the inner wrapped XML type.
func (_fbf BubbleChart) X() *_e.CT_BubbleChart { return _fbf._gb }

// SetText sets the series text
func (_ddgd ScatterChartSeries) SetText(s string) {
	_ddgd._egge.Tx = _e.NewCT_SerTx()
	_ddgd._egge.Tx.Choice.V = &s
}

// ScatterChartSeries is the data series for a scatter chart.
type ScatterChartSeries struct{ _egge *_e.CT_ScatterSer }

// AddCategoryAxis adds a category axis.
func (_bcc Chart) AddCategoryAxis() CategoryAxis {
	_cgg := _e.NewCT_CatAx()
	if _bcc._gdd.Chart.PlotArea.CChoice == nil {
		_bcc._gdd.Chart.PlotArea.CChoice = _e.NewCT_PlotAreaChoice1()
	}
	_cgg.AxId = _e.NewCT_UnsignedInt()
	_cgg.AxId.ValAttr = 0x7FFFFFFF & _d.Uint32()
	_bcc._gdd.Chart.PlotArea.CChoice.CatAx = append(_bcc._gdd.Chart.PlotArea.CChoice.CatAx, _cgg)
	_cgg.Auto = _e.NewCT_Boolean()
	_cgg.Auto.ValAttr = _fe.Bool(true)
	_cgg.Delete = _e.NewCT_Boolean()
	_cgg.Delete.ValAttr = _fe.Bool(false)
	_gcf := MakeCategoryAxis(_cgg)
	_gcf.InitializeDefaults()
	return _gcf
}
func (_egc DateAxis) SetMinorTickMark(m _e.ST_TickMark) {
	if m == _e.ST_TickMarkUnset {
		_egc._gggda.MinorTickMark = nil
	} else {
		_egc._gggda.MinorTickMark = _e.NewCT_TickMark()
		_egc._gggda.MinorTickMark.ValAttr = m
	}
}

// InitializeDefaults the bar chart to its defaults
func (_fce Pie3DChart) InitializeDefaults() {
	_fce._cacc.VaryColors = _e.NewCT_Boolean()
	_fce._cacc.VaryColors.ValAttr = _fe.Bool(true)
}

// X returns the inner wrapped XML type.
func (_b Area3DChart) X() *_e.CT_Area3DChart { return _b._cg }

// AddBarChart adds a new bar chart to a chart.
func (_dfe Chart) AddBarChart() BarChart {
	_cbf := _e.NewCT_PlotAreaChoice()
	_dfe._gdd.Chart.PlotArea.Choice = append(_dfe._gdd.Chart.PlotArea.Choice, _cbf)
	_cbf.BarChart = _e.NewCT_BarChart()
	_cbf.BarChart.Grouping = _e.NewCT_BarGrouping()
	_cbf.BarChart.Grouping.ValAttr = _e.ST_BarGroupingStandard
	_abg := BarChart{_fa: _cbf.BarChart}
	_abg.InitializeDefaults()
	return _abg
}

// SetText sets the series text.
func (_gg BubbleChartSeries) SetText(s string) {
	_gg._gd.Tx = _e.NewCT_SerTx()
	_gg._gd.Tx.Choice.V = &s
}

var _fgc = []_fef.Color{_fef.RGB(0x33, 0x66, 0xcc), _fef.RGB(0xDC, 0x39, 0x12), _fef.RGB(0xFF, 0x99, 0x00), _fef.RGB(0x10, 0x96, 0x18), _fef.RGB(0x99, 0x00, 0x99), _fef.RGB(0x3B, 0x3E, 0xAC), _fef.RGB(0x00, 0x99, 0xC6), _fef.RGB(0xDD, 0x44, 0x77), _fef.RGB(0x66, 0xAA, 0x00), _fef.RGB(0xB8, 0x2E, 0x2E), _fef.RGB(0x31, 0x63, 0x95), _fef.RGB(0x99, 0x44, 0x99), _fef.RGB(0x22, 0xAA, 0x99), _fef.RGB(0xAA, 0xAA, 0x11), _fef.RGB(0x66, 0x33, 0xCC), _fef.RGB(0xE6, 0x73, 0x00), _fef.RGB(0x8B, 0x07, 0x07), _fef.RGB(0x32, 0x92, 0x62), _fef.RGB(0x55, 0x74, 0xA6), _fef.RGB(0x3B, 0x3E, 0xAC)}

// Marker returns the marker properties.
func (_faaa ScatterChartSeries) Marker() Marker {
	if _faaa._egge.Marker == nil {
		_faaa._egge.Marker = _e.NewCT_Marker()
	}
	return MakeMarker(_faaa._egge.Marker)
}
func MakeMarker(x *_e.CT_Marker) Marker { return Marker{x} }

// CreateEmptyNumberCache creates an empty number cache, which is used sometimes
// to increase file format compatibility.  It should actually contain the
// computed cell data, but just creating an empty one is good enough.
func (_geg NumberDataSource) CreateEmptyNumberCache() {
	_geg.ensureChoice()
	if _geg._dce.Choice.NumRef == nil {
		_geg._dce.Choice.NumRef = _e.NewCT_NumRef()
	}
	_geg._dce.Choice.NumLit = nil
	_geg._dce.Choice.NumRef.NumCache = _e.NewCT_NumData()
	_geg._dce.Choice.NumRef.NumCache.PtCount = _e.NewCT_UnsignedInt()
	_geg._dce.Choice.NumRef.NumCache.PtCount.ValAttr = 0
}
func (_accg Marker) SetSymbol(s _e.ST_MarkerStyle) {
	if s == _e.ST_MarkerStyleUnset {
		_accg._gdf.Symbol = nil
	} else {
		_accg._gdf.Symbol = _e.NewCT_MarkerStyle()
		_accg._gdf.Symbol.ValAttr = s
	}
}

var NullAxis Axis = nullAxis(0)

type DataLabels struct{ _bga *_e.CT_DLbls }

// InitializeDefaults the Stock chart to its defaults
func (_aba StockChart) InitializeDefaults() {
	_aba._agc.HiLowLines = _e.NewCT_ChartLines()
	_aba._agc.UpDownBars = _e.NewCT_UpDownBars()
	_aba._agc.UpDownBars.GapWidth = _e.NewCT_GapAmount()
	_aba._agc.UpDownBars.GapWidth.ValAttr = &_e.ST_GapAmount{}
	_aba._agc.UpDownBars.GapWidth.ValAttr.ST_GapAmountUShort = _fe.Uint16(150)
	_aba._agc.UpDownBars.UpBars = _e.NewCT_UpDownBar()
	_aba._agc.UpDownBars.DownBars = _e.NewCT_UpDownBar()
}
func (_cbdf SurfaceChartSeries) Values() NumberDataSource {
	if _cbdf._gdda.Val == nil {
		_cbdf._gdda.Val = _e.NewCT_NumDataSource()
	}
	_bdbc := MakeNumberDataSource(_cbdf._gdda.Val)
	_bdbc.CreateEmptyNumberCache()
	return _bdbc
}

// AddSeries adds a default series to an Pie3D chart.
func (_bfed Pie3DChart) AddSeries() PieChartSeries {
	_bfcb := _e.NewCT_PieSer()
	_bfed._cacc.Ser = append(_bfed._cacc.Ser, _bfcb)
	_bfcb.Idx.ValAttr = uint32(len(_bfed._cacc.Ser) - 1)
	_bfcb.Order.ValAttr = uint32(len(_bfed._cacc.Ser) - 1)
	_acag := PieChartSeries{_bfcb}
	_acag.InitializeDefaults()
	return _acag
}

// BubbleChart is a 2D Bubble chart.
type BubbleChart struct {
	chartBase
	_gb *_e.CT_BubbleChart
}

func (_gbbg DataLabels) SetShowCategoryName(b bool) {
	_gbbg.ensureChoice()
	_gbbg._bga.Choice.ShowCatName = _e.NewCT_Boolean()
	_gbbg._bga.Choice.ShowCatName.ValAttr = _fe.Bool(b)
}

// AddAxis adds an axis to a Scatter chart.
func (_aega ScatterChart) AddAxis(axis Axis) {
	_dcag := _e.NewCT_UnsignedInt()
	_dcag.ValAttr = axis.AxisID()
	_aega._dbbd.AxId = append(_aega._dbbd.AxId, _dcag)
}

// SetText sets the series text.
func (_cdg PieChartSeries) SetText(s string) {
	_cdg._faag.Tx = _e.NewCT_SerTx()
	_cdg._faag.Tx.Choice.V = &s
}

// CategoryAxis returns the category data source.
func (_dafc BubbleChartSeries) CategoryAxis() CategoryAxisDataSource {
	if _dafc._gd.XVal == nil {
		_dafc._gd.XVal = _e.NewCT_AxDataSource()
	}
	return MakeAxisDataSource(_dafc._gd.XVal)
}
func (_abgc Legend) Properties() _fega.ShapeProperties {
	if _abgc._eca.SpPr == nil {
		_abgc._eca.SpPr = _feg.NewCT_ShapeProperties()
	}
	return _fega.MakeShapeProperties(_abgc._eca.SpPr)
}
func (_dbc ScatterChartSeries) CategoryAxis() CategoryAxisDataSource {
	if _dbc._egge.XVal == nil {
		_dbc._egge.XVal = _e.NewCT_AxDataSource()
	}
	return MakeAxisDataSource(_dbc._egge.XVal)
}

// CategoryAxisDataSource specifies the data for an axis.  It's commonly used with
// SetReference to set the axis data to a range of cells.
type CategoryAxisDataSource struct{ _dgd *_e.CT_AxDataSource }

func (_fbfg ValueAxis) Properties() _fega.ShapeProperties {
	if _fbfg._cdf.SpPr == nil {
		_fbfg._cdf.SpPr = _feg.NewCT_ShapeProperties()
	}
	return _fega.MakeShapeProperties(_fbfg._cdf.SpPr)
}

// AddSeries adds a default series to a Stock chart.
func (_fcc StockChart) AddSeries() LineChartSeries {
	_gcfa := _e.NewCT_LineSer()
	_fcc._agc.Ser = append(_fcc._agc.Ser, _gcfa)
	_gcfa.Idx.ValAttr = uint32(len(_fcc._agc.Ser) - 1)
	_gcfa.Order.ValAttr = uint32(len(_fcc._agc.Ser) - 1)
	_gcc := LineChartSeries{_gcfa}
	_gcc.Values().CreateEmptyNumberCache()
	_gcc.Properties().LineProperties().SetNoFill()
	return _gcc
}

// SetHoleSize controls the hole size in the pie chart and is measured in percent.
func (_beb DoughnutChart) SetHoleSize(pct uint8) {
	if _beb._gebe.HoleSize == nil {
		_beb._gebe.HoleSize = _e.NewCT_HoleSize()
	}
	if _beb._gebe.HoleSize.ValAttr == nil {
		_beb._gebe.HoleSize.ValAttr = &_e.ST_HoleSize{}
	}
	_beb._gebe.HoleSize.ValAttr.ST_HoleSizeUByte = &pct
}

// RemoveLegend removes the legend if the chart has one.
func (_aca Chart) RemoveLegend() { _aca._gdd.Chart.Legend = nil }

// InitializeDefaults initializes an Radar series to the default values.
func (_badg RadarChartSeries) InitializeDefaults() {}
func (_bddc SurfaceChartSeries) InitializeDefaults() {
	_bddc.Properties().LineProperties().SetWidth(1 * _ee.Point)
	_bddc.Properties().LineProperties().SetSolidFill(_fef.Black)
	_bddc.Properties().LineProperties().SetJoin(_fega.LineJoinRound)
}

// SetDirection changes the direction of the bar chart (bar or column).
func (_fdb BarChart) SetDirection(d _e.ST_BarDir) { _fdb._fa.BarDir.ValAttr = d }
func (_bab Surface3DChart) InitializeDefaults() {
	_bab._abfa.Wireframe = _e.NewCT_Boolean()
	_bab._abfa.Wireframe.ValAttr = _fe.Bool(false)
	_bab._abfa.BandFmts = _e.NewCT_BandFmts()
	for _cfeb := 0; _cfeb < 15; _cfeb++ {
		_gddb := _e.NewCT_BandFmt()
		_gddb.Idx.ValAttr = uint32(_cfeb)
		_gddb.SpPr = _feg.NewCT_ShapeProperties()
		_dceg := _fega.MakeShapeProperties(_gddb.SpPr)
		_dceg.SetSolidFill(_bab.nextColor(_cfeb))
		_bab._abfa.BandFmts.BandFmt = append(_bab._abfa.BandFmts.BandFmt, _gddb)
	}
}

// X returns the inner wrapped XML type.
func (_fcdg Title) X() *_e.CT_Title { return _fcdg._afaa }

// BarChart is a 2D bar chart.
type BarChart struct {
	chartBase
	_fa *_e.CT_BarChart
}

func (_ace SeriesAxis) InitializeDefaults() {}

// X returns the inner wrapped XML type.
func (_fed Pie3DChart) X() *_e.CT_Pie3DChart { return _fed._cacc }

// AreaChartSeries is a series to be used on an area chart.
type AreaChartSeries struct{ _caa *_e.CT_AreaSer }

func (_dea ValueAxis) MajorGridLines() GridLines {
	if _dea._cdf.MajorGridlines == nil {
		_dea._cdf.MajorGridlines = _e.NewCT_ChartLines()
	}
	return GridLines{_dea._cdf.MajorGridlines}
}
func (_baef DateAxis) MajorGridLines() GridLines {
	if _baef._gggda.MajorGridlines == nil {
		_baef._gggda.MajorGridlines = _e.NewCT_ChartLines()
	}
	return GridLines{_baef._gggda.MajorGridlines}
}

// Properties returns the Bubble chart series shape properties.
func (_ffb BubbleChartSeries) Properties() _fega.ShapeProperties {
	if _ffb._gd.SpPr == nil {
		_ffb._gd.SpPr = _feg.NewCT_ShapeProperties()
	}
	return _fega.MakeShapeProperties(_ffb._gd.SpPr)
}

// X returns the inner wrapped XML type.
func (_fedc ValueAxis) X() *_e.CT_ValAx { return _fedc._cdf }
func (_cga SeriesAxis) AxisID() uint32  { return _cga._faac.AxId.ValAttr }
func (_fcb NumberDataSource) SetReference(s string) {
	_fcb.ensureChoice()
	if _fcb._dce.Choice.NumRef == nil {
		_fcb._dce.Choice.NumRef = _e.NewCT_NumRef()
	}
	_fcb._dce.Choice.NumRef.F = s
}

// AddSeries adds a default series to a line chart.
func (_fab Line3DChart) AddSeries() LineChartSeries {
	_bcaf := _fab.nextColor(len(_fab._cbd.Ser))
	_gggdb := _e.NewCT_LineSer()
	_fab._cbd.Ser = append(_fab._cbd.Ser, _gggdb)
	_gggdb.Idx.ValAttr = uint32(len(_fab._cbd.Ser) - 1)
	_gggdb.Order.ValAttr = uint32(len(_fab._cbd.Ser) - 1)
	_ea := LineChartSeries{_gggdb}
	_ea.InitializeDefaults()
	_ea.Properties().LineProperties().SetSolidFill(_bcaf)
	_ea.Properties().SetSolidFill(_bcaf)
	return _ea
}
func (_ecab LineChartSeries) SetSmooth(b bool) {
	_ecab._afc.Smooth = _e.NewCT_Boolean()
	_ecab._afc.Smooth.ValAttr = &b
}

// AddSeries adds a default series to a bar chart.
func (_ga BarChart) AddSeries() BarChartSeries {
	_efe := _ga.nextColor(len(_ga._fa.Ser))
	_fff := _e.NewCT_BarSer()
	_ga._fa.Ser = append(_ga._fa.Ser, _fff)
	_fff.Idx.ValAttr = uint32(len(_ga._fa.Ser) - 1)
	_fff.Order.ValAttr = uint32(len(_ga._fa.Ser) - 1)
	_dc := BarChartSeries{_fff}
	_dc.InitializeDefaults()
	_dc.Properties().SetSolidFill(_efe)
	return _dc
}
func (_beec DateAxis) SetTickLabelPosition(p _e.ST_TickLblPos) {
	if p == _e.ST_TickLblPosUnset {
		_beec._gggda.TickLblPos = nil
	} else {
		_beec._gggda.TickLblPos = _e.NewCT_TickLblPos()
		_beec._gggda.TickLblPos.ValAttr = p
	}
}

// RadarChartSeries is a series to be used on an Radar chart.
type RadarChartSeries struct{ _eed *_e.CT_RadarSer }

// Labels returns the data label properties.
func (_bbe ScatterChartSeries) Labels() DataLabels {
	if _bbe._egge.DLbls == nil {
		_bbe._egge.DLbls = _e.NewCT_DLbls()
	}
	return MakeDataLabels(_bbe._egge.DLbls)
}

type CategoryAxis struct{ _aac *_e.CT_CatAx }

// AddAreaChart adds a new area chart to a chart.
func (_gge Chart) AddAreaChart() AreaChart {
	_bgg := _e.NewCT_PlotAreaChoice()
	_gge._gdd.Chart.PlotArea.Choice = append(_gge._gdd.Chart.PlotArea.Choice, _bgg)
	_bgg.AreaChart = _e.NewCT_AreaChart()
	_dad := AreaChart{_cc: _bgg.AreaChart}
	_dad.InitializeDefaults()
	return _dad
}

// InitializeDefaults the bar chart to its defaults
func (_abbg PieOfPieChart) InitializeDefaults() {
	_abbg._ggac.VaryColors = _e.NewCT_Boolean()
	_abbg._ggac.VaryColors.ValAttr = _fe.Bool(true)
	_abbg.SetType(_e.ST_OfPieTypePie)
	_abbg._ggac.SecondPieSize = _e.NewCT_SecondPieSize()
	_abbg._ggac.SecondPieSize.ValAttr = &_e.ST_SecondPieSize{}
	_abbg._ggac.SecondPieSize.ValAttr.ST_SecondPieSizeUShort = _fe.Uint16(75)
	_fceb := _e.NewCT_ChartLines()
	_fceb.SpPr = _feg.NewCT_ShapeProperties()
	_bgb := _fega.MakeShapeProperties(_fceb.SpPr)
	_bgb.LineProperties().SetSolidFill(_fef.Auto)
	_abbg._ggac.SerLines = append(_abbg._ggac.SerLines, _fceb)
}

// SetText sets the series text
func (_eag LineChartSeries) SetText(s string) {
	_eag._afc.Tx = _e.NewCT_SerTx()
	_eag._afc.Tx.Choice.V = &s
}

// SurfaceChart is a 3D surface chart, viewed from the top-down.
type SurfaceChart struct {
	chartBase
	_gcaa *_e.CT_SurfaceChart
}

// CategoryAxis returns the category data source.
func (_efa PieChartSeries) CategoryAxis() CategoryAxisDataSource {
	if _efa._faag.Cat == nil {
		_efa._faag.Cat = _e.NewCT_AxDataSource()
	}
	return MakeAxisDataSource(_efa._faag.Cat)
}
func (_bgeb ValueAxis) SetPosition(p _e.ST_AxPos) {
	_bgeb._cdf.AxPos = _e.NewCT_AxPos()
	_bgeb._cdf.AxPos.ValAttr = p
}
func (_adbb Title) InitializeDefaults() {
	_adbb.SetText("Title")
	_adbb.RunProperties().SetSize(16 * _ee.Point)
	_adbb.RunProperties().SetSolidFill(_fef.Black)
	_adbb.RunProperties().SetFont("Calib\u0020ri")
	_adbb.RunProperties().SetBold(false)
}

// X returns the inner wrapped XML type.
func (_acad RadarChartSeries) X() *_e.CT_RadarSer { return _acad._eed }

type SurfaceChartSeries struct{ _gdda *_e.CT_SurfaceSer }

// AddSurfaceChart adds a new surface chart to a chart.
func (_geb Chart) AddSurfaceChart() SurfaceChart {
	_baaa := _e.NewCT_PlotAreaChoice()
	_geb._gdd.Chart.PlotArea.Choice = append(_geb._gdd.Chart.PlotArea.Choice, _baaa)
	_baaa.SurfaceChart = _e.NewCT_SurfaceChart()
	_ad(_geb._gdd.Chart)
	_geb._gdd.Chart.View3D.RotX.ValAttr = _fe.Int8(90)
	_geb._gdd.Chart.View3D.RotY.ValAttr = _fe.Uint16(0)
	_geb._gdd.Chart.View3D.Perspective = _e.NewCT_Perspective()
	_geb._gdd.Chart.View3D.Perspective.ValAttr = _fe.Uint8(0)
	_gbd := SurfaceChart{_gcaa: _baaa.SurfaceChart}
	_gbd.InitializeDefaults()
	return _gbd
}
func (_efd RadarChart) AddAxis(axis Axis) {
	_aff := _e.NewCT_UnsignedInt()
	_aff.ValAttr = axis.AxisID()
	_efd._ebcd.AxId = append(_efd._ebcd.AxId, _aff)
}

type Title struct{ _afaa *_e.CT_Title }
type GridLines struct{ _fgcf *_e.CT_ChartLines }

func (_bg Bar3DChart) AddAxis(axis Axis) {
	_bae := _e.NewCT_UnsignedInt()
	_bae.ValAttr = axis.AxisID()
	_bg._daa.AxId = append(_bg._daa.AxId, _bae)
}
func (_caef SeriesAxis) SetCrosses(axis Axis) { _caef._faac.CrossAx.ValAttr = axis.AxisID() }

// X returns the inner wrapped XML type.
func (_cfa StockChart) X() *_e.CT_StockChart { return _cfa._agc }

// SetText sets the series text
func (_fgf SurfaceChartSeries) SetText(s string) {
	_fgf._gdda.Tx = _e.NewCT_SerTx()
	_fgf._gdda.Tx.Choice.V = &s
}

// RemoveTitle removes any existing title from the chart.
func (_ebg Chart) RemoveTitle() {
	_ebg._gdd.Chart.Title = nil
	_ebg._gdd.Chart.AutoTitleDeleted = _e.NewCT_Boolean()
	_ebg._gdd.Chart.AutoTitleDeleted.ValAttr = _fe.Bool(true)
}
func (_bdf AreaChart) AddAxis(axis Axis) {
	_dg := _e.NewCT_UnsignedInt()
	_dg.ValAttr = axis.AxisID()
	_bdf._cc.AxId = append(_bdf._cc.AxId, _dg)
}

// AddSeries adds a default series to a Surface chart.
func (_afa Surface3DChart) AddSeries() SurfaceChartSeries {
	_efga := _afa.nextColor(len(_afa._abfa.Ser))
	_dbff := _e.NewCT_SurfaceSer()
	_afa._abfa.Ser = append(_afa._abfa.Ser, _dbff)
	_dbff.Idx.ValAttr = uint32(len(_afa._abfa.Ser) - 1)
	_dbff.Order.ValAttr = uint32(len(_afa._abfa.Ser) - 1)
	_ggd := SurfaceChartSeries{_dbff}
	_ggd.InitializeDefaults()
	_ggd.Properties().LineProperties().SetSolidFill(_efga)
	return _ggd
}
func (_bggb ValueAxis) AxisID() uint32 { return _bggb._cdf.AxId.ValAttr }

// X returns the inner wrapped XML type.
func (_be AreaChartSeries) X() *_e.CT_AreaSer { return _be._caa }

type ValueAxis struct{ _cdf *_e.CT_ValAx }

// Order returns the order of the series
func (_dbd LineChartSeries) Order() uint32 { return _dbd._afc.Order.ValAttr }
func (_cgd BarChart) AddAxis(axis Axis) {
	_fba := _e.NewCT_UnsignedInt()
	_fba.ValAttr = axis.AxisID()
	_cgd._fa.AxId = append(_cgd._fa.AxId, _fba)
}

// Properties returns the line chart series shape properties.
func (_fea LineChartSeries) Properties() _fega.ShapeProperties {
	if _fea._afc.SpPr == nil {
		_fea._afc.SpPr = _feg.NewCT_ShapeProperties()
	}
	return _fega.MakeShapeProperties(_fea._afc.SpPr)
}

// Properties returns the line chart series shape properties.
func (_fbba ScatterChartSeries) Properties() _fega.ShapeProperties {
	if _fbba._egge.SpPr == nil {
		_fbba._egge.SpPr = _feg.NewCT_ShapeProperties()
	}
	return _fega.MakeShapeProperties(_fbba._egge.SpPr)
}

// SetNumberReference is used to set the source data to a range of cells containing
// numbers.
func (_gbf CategoryAxisDataSource) SetNumberReference(s string) {
	_gbf._dgd.Choice = _e.NewCT_AxDataSourceChoice()
	_gbf._dgd.Choice.NumRef = _e.NewCT_NumRef()
	_gbf._dgd.Choice.NumRef.F = s
}

// SetDisplayBlanksAs controls how missing values are displayed.
func (_befg Chart) SetDisplayBlanksAs(v _e.ST_DispBlanksAs) {
	_befg._gdd.Chart.DispBlanksAs = _e.NewCT_DispBlanksAs()
	_befg._gdd.Chart.DispBlanksAs.ValAttr = v
}
func (_fgec DataLabels) SetPosition(p _e.ST_DLblPos) {
	_fgec.ensureChoice()
	_fgec._bga.Choice.DLblPos = _e.NewCT_DLblPos()
	_fgec._bga.Choice.DLblPos.ValAttr = p
}

// Values returns the value data source.
func (_deg RadarChartSeries) Values() NumberDataSource {
	if _deg._eed.Val == nil {
		_deg._eed.Val = _e.NewCT_NumDataSource()
	}
	return MakeNumberDataSource(_deg._eed.Val)
}

// SetValues is used to set the source data to a set of values.
func (_abb CategoryAxisDataSource) SetValues(v []string) {
	_abb._dgd.Choice = _e.NewCT_AxDataSourceChoice()
	_abb._dgd.Choice.StrLit = _e.NewCT_StrData()
	_abb._dgd.Choice.StrLit.PtCount = _e.NewCT_UnsignedInt()
	_abb._dgd.Choice.StrLit.PtCount.ValAttr = uint32(len(v))
	for _eg, _dge := range v {
		_abb._dgd.Choice.StrLit.Pt = append(_abb._dgd.Choice.StrLit.Pt, &_e.CT_StrVal{IdxAttr: uint32(_eg), V: _dge})
	}
}
func (_aaea DateAxis) SetCrosses(axis Axis) { _aaea._gggda.CrossAx.ValAttr = axis.AxisID() }

// InitializeDefaults the bar chart to its defaults
func (_ef BarChart) InitializeDefaults() { _ef.SetDirection(_e.ST_BarDirCol) }

// X returns the inner wrapped XML type.
func (_gbe Chart) X() *_e.ChartSpace { return _gbe._gdd }

// CategoryAxis returns the category data source.
func (_bge RadarChartSeries) CategoryAxis() CategoryAxisDataSource {
	if _bge._eed.Cat == nil {
		_bge._eed.Cat = _e.NewCT_AxDataSource()
	}
	return MakeAxisDataSource(_bge._eed.Cat)
}
func (_faa BubbleChart) AddAxis(axis Axis) {
	_ab := _e.NewCT_UnsignedInt()
	_ab.ValAttr = axis.AxisID()
	_faa._gb.AxId = append(_faa._gb.AxId, _ab)
}

// AddBubbleChart adds a new bubble chart.
func (_def Chart) AddBubbleChart() BubbleChart {
	_gdc := _e.NewCT_PlotAreaChoice()
	_def._gdd.Chart.PlotArea.Choice = append(_def._gdd.Chart.PlotArea.Choice, _gdc)
	_gdc.BubbleChart = _e.NewCT_BubbleChart()
	_cbg := BubbleChart{_gb: _gdc.BubbleChart}
	_cbg.InitializeDefaults()
	return _cbg
}

// X returns the inner wrapped XML type.
func (_gcfe PieOfPieChart) X() *_e.CT_OfPieChart { return _gcfe._ggac }
func (_bdg ValueAxis) SetMinorTickMark(m _e.ST_TickMark) {
	if m == _e.ST_TickMarkUnset {
		_bdg._cdf.MinorTickMark = nil
	} else {
		_bdg._cdf.MinorTickMark = _e.NewCT_TickMark()
		_bdg._cdf.MinorTickMark.ValAttr = m
	}
}
func (_egda SurfaceChart) InitializeDefaults() {
	_egda._gcaa.Wireframe = _e.NewCT_Boolean()
	_egda._gcaa.Wireframe.ValAttr = _fe.Bool(false)
	_egda._gcaa.BandFmts = _e.NewCT_BandFmts()
	for _ccf := 0; _ccf < 15; _ccf++ {
		_fffd := _e.NewCT_BandFmt()
		_fffd.Idx.ValAttr = uint32(_ccf)
		_fffd.SpPr = _feg.NewCT_ShapeProperties()
		_dcfe := _fega.MakeShapeProperties(_fffd.SpPr)
		_dcfe.SetSolidFill(_egda.nextColor(_ccf))
		_egda._gcaa.BandFmts.BandFmt = append(_egda._gcaa.BandFmts.BandFmt, _fffd)
	}
}

// AddLegend adds a legend to a chart, replacing any existing legend.
func (_aea Chart) AddLegend() Legend {
	_aea._gdd.Chart.Legend = _e.NewCT_Legend()
	_beff := MakeLegend(_aea._gdd.Chart.Legend)
	_beff.InitializeDefaults()
	return _beff
}
func (_adb DateAxis) SetMajorTickMark(m _e.ST_TickMark) {
	if m == _e.ST_TickMarkUnset {
		_adb._gggda.MajorTickMark = nil
	} else {
		_adb._gggda.MajorTickMark = _e.NewCT_TickMark()
		_adb._gggda.MajorTickMark.ValAttr = m
	}
}
func MakeValueAxis(x *_e.CT_ValAx) ValueAxis { return ValueAxis{x} }
func _ad(_gde *_e.CT_Chart) {
	_gde.View3D = _e.NewCT_View3D()
	_gde.View3D.RotX = _e.NewCT_RotX()
	_gde.View3D.RotX.ValAttr = _fe.Int8(15)
	_gde.View3D.RotY = _e.NewCT_RotY()
	_gde.View3D.RotY.ValAttr = _fe.Uint16(20)
	_gde.View3D.RAngAx = _e.NewCT_Boolean()
	_gde.View3D.RAngAx.ValAttr = _fe.Bool(false)
	_gde.Floor = _e.NewCT_Surface()
	_gde.Floor.Thickness = _e.NewCT_Thickness()
	_gde.Floor.Thickness.ValAttr.Uint32 = _fe.Uint32(0)
	_gde.SideWall = _e.NewCT_Surface()
	_gde.SideWall.Thickness = _e.NewCT_Thickness()
	_gde.SideWall.Thickness.ValAttr.Uint32 = _fe.Uint32(0)
	_gde.BackWall = _e.NewCT_Surface()
	_gde.BackWall.Thickness = _e.NewCT_Thickness()
	_gde.BackWall.Thickness.ValAttr.Uint32 = _fe.Uint32(0)
}

type ScatterChart struct {
	chartBase
	_dbbd *_e.CT_ScatterChart
}

// X returns the inner wrapped XML type.
func (_beg LineChart) X() *_e.CT_LineChart { return _beg._gafd }

// AddTitle sets a new title on the chart.
func (_ffge Chart) AddTitle() Title {
	_ffge._gdd.Chart.Title = _e.NewCT_Title()
	_ffge._gdd.Chart.Title.Overlay = _e.NewCT_Boolean()
	_ffge._gdd.Chart.Title.Overlay.ValAttr = _fe.Bool(false)
	_ffge._gdd.Chart.AutoTitleDeleted = _e.NewCT_Boolean()
	_ffge._gdd.Chart.AutoTitleDeleted.ValAttr = _fe.Bool(false)
	_fge := MakeTitle(_ffge._gdd.Chart.Title)
	_fge.InitializeDefaults()
	return _fge
}

// Surface3DChart is a 3D view of a surface chart.
type Surface3DChart struct {
	chartBase
	_abfa *_e.CT_Surface3DChart
}

func (_ffc CategoryAxis) SetMajorTickMark(m _e.ST_TickMark) {
	if m == _e.ST_TickMarkUnset {
		_ffc._aac.MajorTickMark = nil
	} else {
		_ffc._aac.MajorTickMark = _e.NewCT_TickMark()
		_ffc._aac.MajorTickMark.ValAttr = m
	}
}

// AddArea3DChart adds a new area chart to a chart.
func (_bff Chart) AddArea3DChart() Area3DChart {
	_ad(_bff._gdd.Chart)
	_bea := _e.NewCT_PlotAreaChoice()
	_bff._gdd.Chart.PlotArea.Choice = append(_bff._gdd.Chart.PlotArea.Choice, _bea)
	_bea.Area3DChart = _e.NewCT_Area3DChart()
	_fbb := Area3DChart{_cg: _bea.Area3DChart}
	_fbb.InitializeDefaults()
	return _fbb
}

// Values returns the bubble size data source.
func (_ddc BubbleChartSeries) BubbleSizes() NumberDataSource {
	if _ddc._gd.BubbleSize == nil {
		_ddc._gd.BubbleSize = _e.NewCT_NumDataSource()
	}
	return MakeNumberDataSource(_ddc._gd.BubbleSize)
}
func MakeSeriesAxis(x *_e.CT_SerAx) SeriesAxis { return SeriesAxis{x} }
func (_baa nullAxis) AxisID() uint32           { return 0 }

// X returns the inner wrapped XML type.
func (_a AreaChart) X() *_e.CT_AreaChart { return _a._cc }

// Index returns the index of the series
func (_egd LineChartSeries) Index() uint32 { return _egd._afc.Idx.ValAttr }

// PieChart is a Pie chart.
type PieChart struct {
	chartBase
	_bddb *_e.CT_PieChart
}

func (_ggb NumberDataSource) ensureChoice() {
	if _ggb._dce.Choice == nil {
		_ggb._dce.Choice = _e.NewCT_NumDataSourceChoice()
	}
}

// Area3DChart is an area chart that has a shaded area underneath a curve.
type Area3DChart struct {
	chartBase
	_cg *_e.CT_Area3DChart
}

func (_dcd CategoryAxis) InitializeDefaults() {
	_dcd.SetPosition(_e.ST_AxPosB)
	_dcd.SetMajorTickMark(_e.ST_TickMarkOut)
	_dcd.SetMinorTickMark(_e.ST_TickMarkIn)
	_dcd.SetTickLabelPosition(_e.ST_TickLblPosNextTo)
	_dcd.MajorGridLines().Properties().LineProperties().SetSolidFill(_fef.LightGray)
	_dcd.Properties().LineProperties().SetSolidFill(_fef.Black)
}

// AddValueAxis adds a value axis to the chart.
func (_dfee Chart) AddValueAxis() ValueAxis {
	_cae := _e.NewCT_ValAx()
	if _dfee._gdd.Chart.PlotArea.CChoice == nil {
		_dfee._gdd.Chart.PlotArea.CChoice = _e.NewCT_PlotAreaChoice1()
	}
	_cae.AxId = _e.NewCT_UnsignedInt()
	_cae.AxId.ValAttr = 0x7FFFFFFF & _d.Uint32()
	_dfee._gdd.Chart.PlotArea.CChoice.ValAx = append(_dfee._gdd.Chart.PlotArea.CChoice.ValAx, _cae)
	_cae.Delete = _e.NewCT_Boolean()
	_cae.Delete.ValAttr = _fe.Bool(false)
	_cae.Scaling = _e.NewCT_Scaling()
	_cae.Scaling.Orientation = _e.NewCT_Orientation()
	_cae.Scaling.Orientation.ValAttr = _e.ST_OrientationMinMax
	_cae.Choice = &_e.EG_AxSharedChoice{}
	_cae.Choice.Crosses = _e.NewCT_Crosses()
	_cae.Choice.Crosses.ValAttr = _e.ST_CrossesAutoZero
	_cae.CrossBetween = _e.NewCT_CrossBetween()
	_cae.CrossBetween.ValAttr = _e.ST_CrossBetweenBetween
	_acc := MakeValueAxis(_cae)
	_acc.MajorGridLines().Properties().LineProperties().SetSolidFill(_fef.LightGray)
	_acc.SetMajorTickMark(_e.ST_TickMarkOut)
	_acc.SetMinorTickMark(_e.ST_TickMarkIn)
	_acc.SetTickLabelPosition(_e.ST_TickLblPosNextTo)
	_acc.Properties().LineProperties().SetSolidFill(_fef.Black)
	_acc.SetPosition(_e.ST_AxPosL)
	return _acc
}

// AddSeries adds a default series to an Radar chart.
func (_aee RadarChart) AddSeries() RadarChartSeries {
	_ebe := _aee.nextColor(len(_aee._ebcd.Ser))
	_bfbf := _e.NewCT_RadarSer()
	_aee._ebcd.Ser = append(_aee._ebcd.Ser, _bfbf)
	_bfbf.Idx.ValAttr = uint32(len(_aee._ebcd.Ser) - 1)
	_bfbf.Order.ValAttr = uint32(len(_aee._ebcd.Ser) - 1)
	_dbgg := RadarChartSeries{_bfbf}
	_dbgg.InitializeDefaults()
	_dbgg.Properties().SetSolidFill(_ebe)
	return _dbgg
}

// X returns the inner wrapped XML type.
func (_gbeg Line3DChart) X() *_e.CT_Line3DChart { return _gbeg._cbd }
func MakeLegend(l *_e.CT_Legend) Legend         { return Legend{l} }

// Values returns the value data source.
func (_fdbb PieChartSeries) Values() NumberDataSource {
	if _fdbb._faag.Val == nil {
		_fdbb._faag.Val = _e.NewCT_NumDataSource()
	}
	return MakeNumberDataSource(_fdbb._faag.Val)
}

// SetText sets the series text.
func (_ced AreaChartSeries) SetText(s string) {
	_ced._caa.Tx = _e.NewCT_SerTx()
	_ced._caa.Tx.Choice.V = &s
}
func (_bc CategoryAxis) SetCrosses(axis Axis) {
	_bc._aac.Choice = _e.NewEG_AxSharedChoice()
	_bc._aac.Choice.Crosses = _e.NewCT_Crosses()
	_bc._aac.Choice.Crosses.ValAttr = _e.ST_CrossesAutoZero
	_bc._aac.CrossAx.ValAttr = axis.AxisID()
}

// Properties returns the bar chart series shape properties.
func (_dd AreaChartSeries) Properties() _fega.ShapeProperties {
	if _dd._caa.SpPr == nil {
		_dd._caa.SpPr = _feg.NewCT_ShapeProperties()
	}
	return _fega.MakeShapeProperties(_dd._caa.SpPr)
}

// InitializeDefaults the bar chart to its defaults
func (_ca Area3DChart) InitializeDefaults() {}

// SetIndex sets the index of the series
func (_feba ScatterChartSeries) SetIndex(idx uint32) { _feba._egge.Idx.ValAttr = idx }
func (_ggg CategoryAxis) MajorGridLines() GridLines {
	if _ggg._aac.MajorGridlines == nil {
		_ggg._aac.MajorGridlines = _e.NewCT_ChartLines()
	}
	return GridLines{_ggg._aac.MajorGridlines}
}
func (_bgf Title) RunProperties() _fega.RunProperties {
	if _bgf._afaa.Tx == nil {
		_bgf.SetText("")
	}
	if _bgf._afaa.Tx.Choice.Rich.P[0].EG_TextRun[0].R.RPr == nil {
		_bgf._afaa.Tx.Choice.Rich.P[0].EG_TextRun[0].R.RPr = _feg.NewCT_TextCharacterProperties()
	}
	return _fega.MakeRunProperties(_bgf._afaa.Tx.Choice.Rich.P[0].EG_TextRun[0].R.RPr)
}

type SeriesAxis struct{ _faac *_e.CT_SerAx }

// AddSeries adds a default series to an Pie chart.
func (_afb PieOfPieChart) AddSeries() PieChartSeries {
	_ecea := _e.NewCT_PieSer()
	_afb._ggac.Ser = append(_afb._ggac.Ser, _ecea)
	_ecea.Idx.ValAttr = uint32(len(_afb._ggac.Ser) - 1)
	_ecea.Order.ValAttr = uint32(len(_afb._ggac.Ser) - 1)
	_fcg := PieChartSeries{_ecea}
	_fcg.InitializeDefaults()
	return _fcg
}

// Bar3DChart is a 3D bar chart.
type Bar3DChart struct {
	chartBase
	_daa *_e.CT_Bar3DChart
}

func (_fdd ScatterChartSeries) Values() NumberDataSource {
	if _fdd._egge.YVal == nil {
		_fdd._egge.YVal = _e.NewCT_NumDataSource()
	}
	return MakeNumberDataSource(_fdd._egge.YVal)
}

// X returns the inner wrapped XML type.
func (_ebea Surface3DChart) X() *_e.CT_Surface3DChart { return _ebea._abfa }

// SetText sets the series text.
func (_fga BarChartSeries) SetText(s string) {
	_fga._fc.Tx = _e.NewCT_SerTx()
	_fga._fc.Tx.Choice.V = &s
}
func (_cbfb Title) SetText(s string) {
	if _cbfb._afaa.Tx == nil {
		_cbfb._afaa.Tx = _e.NewCT_Tx()
	}
	if _cbfb._afaa.Tx.Choice.Rich == nil {
		_cbfb._afaa.Tx.Choice.Rich = _feg.NewCT_TextBody()
	}
	var _gbda *_feg.CT_TextParagraph
	if len(_cbfb._afaa.Tx.Choice.Rich.P) == 0 {
		_gbda = _feg.NewCT_TextParagraph()
		_cbfb._afaa.Tx.Choice.Rich.P = []*_feg.CT_TextParagraph{_gbda}
	} else {
		_gbda = _cbfb._afaa.Tx.Choice.Rich.P[0]
	}
	var _cabc *_feg.EG_TextRun
	if len(_gbda.EG_TextRun) == 0 {
		_cabc = _feg.NewEG_TextRun()
		_gbda.EG_TextRun = []*_feg.EG_TextRun{_cabc}
	} else {
		_cabc = _gbda.EG_TextRun[0]
	}
	if _cabc.R == nil {
		_cabc.R = _feg.NewCT_RegularTextRun()
	}
	_cabc.R.T = s
}

// Index returns the index of the series
func (_fde SurfaceChartSeries) Index() uint32 { return _fde._gdda.Idx.ValAttr }
func (_abf DataLabels) SetShowLeaderLines(b bool) {
	_abf.ensureChoice()
	_abf._bga.Choice.ShowLeaderLines = _e.NewCT_Boolean()
	_abf._bga.Choice.ShowLeaderLines.ValAttr = _fe.Bool(b)
}
func (_fgfb ValueAxis) SetMajorTickMark(m _e.ST_TickMark) {
	if m == _e.ST_TickMarkUnset {
		_fgfb._cdf.MajorTickMark = nil
	} else {
		_fgfb._cdf.MajorTickMark = _e.NewCT_TickMark()
		_fgfb._cdf.MajorTickMark.ValAttr = m
	}
}

// Order returns the order of the series
func (_gcfec SurfaceChartSeries) Order() uint32 { return _gcfec._gdda.Order.ValAttr }

// MakeAxisDataSource constructs an AxisDataSource wrapper.
func MakeAxisDataSource(x *_e.CT_AxDataSource) CategoryAxisDataSource {
	return CategoryAxisDataSource{x}
}

type LineChart struct {
	chartBase
	_gafd *_e.CT_LineChart
}

// AddPie3DChart adds a new pie chart to a chart.
func (_gaf Chart) AddPie3DChart() Pie3DChart {
	_ad(_gaf._gdd.Chart)
	_aab := _e.NewCT_PlotAreaChoice()
	_gaf._gdd.Chart.PlotArea.Choice = append(_gaf._gdd.Chart.PlotArea.Choice, _aab)
	_aab.Pie3DChart = _e.NewCT_Pie3DChart()
	_gbb := Pie3DChart{_cacc: _aab.Pie3DChart}
	_gbb.InitializeDefaults()
	return _gbb
}

// AddSeries adds a default series to an area chart.
func (_ff AreaChart) AddSeries() AreaChartSeries {
	_cf := _ff.nextColor(len(_ff._cc.Ser))
	_af := _e.NewCT_AreaSer()
	_ff._cc.Ser = append(_ff._cc.Ser, _af)
	_af.Idx.ValAttr = uint32(len(_ff._cc.Ser) - 1)
	_af.Order.ValAttr = uint32(len(_ff._cc.Ser) - 1)
	_ba := AreaChartSeries{_af}
	_ba.InitializeDefaults()
	_ba.Properties().SetSolidFill(_cf)
	return _ba
}

// AreaChart is an area chart that has a shaded area underneath a curve.
type AreaChart struct {
	chartBase
	_cc *_e.CT_AreaChart
}

// AddSeries adds a default series to a line chart.
func (_fcf LineChart) AddSeries() LineChartSeries {
	_gbeb := _fcf.nextColor(len(_fcf._gafd.Ser))
	_gcb := _e.NewCT_LineSer()
	_fcf._gafd.Ser = append(_fcf._gafd.Ser, _gcb)
	_gcb.Idx.ValAttr = uint32(len(_fcf._gafd.Ser) - 1)
	_gcb.Order.ValAttr = uint32(len(_fcf._gafd.Ser) - 1)
	_daag := LineChartSeries{_gcb}
	_daag.InitializeDefaults()
	_daag.Properties().LineProperties().SetSolidFill(_gbeb)
	return _daag
}
func (_gf CategoryAxis) Properties() _fega.ShapeProperties {
	if _gf._aac.SpPr == nil {
		_gf._aac.SpPr = _feg.NewCT_ShapeProperties()
	}
	return _fega.MakeShapeProperties(_gf._aac.SpPr)
}

// X returns the inner wrapped XML type.
func (_gbg DoughnutChart) X() *_e.CT_DoughnutChart { return _gbg._gebe }

// CategoryAxis returns the category data source.
func (_ge AreaChartSeries) CategoryAxis() CategoryAxisDataSource {
	if _ge._caa.Cat == nil {
		_ge._caa.Cat = _e.NewCT_AxDataSource()
	}
	return MakeAxisDataSource(_ge._caa.Cat)
}
func (_dbe Legend) SetOverlay(b bool) {
	_dbe._eca.Overlay = _e.NewCT_Boolean()
	_dbe._eca.Overlay.ValAttr = _fe.Bool(b)
}

// X returns the inner wrapped XML type.
func (_bda BarChartSeries) X() *_e.CT_BarSer { return _bda._fc }

// AddSeries adds a default series to a Bubble chart.
func (_bdd BubbleChart) AddSeries() BubbleChartSeries {
	_eff := _bdd.nextColor(len(_bdd._gb.Ser))
	_daf := _e.NewCT_BubbleSer()
	_bdd._gb.Ser = append(_bdd._gb.Ser, _daf)
	_daf.Idx.ValAttr = uint32(len(_bdd._gb.Ser) - 1)
	_daf.Order.ValAttr = uint32(len(_bdd._gb.Ser) - 1)
	_bde := BubbleChartSeries{_daf}
	_bde.InitializeDefaults()
	_bde.Properties().SetSolidFill(_eff)
	return _bde
}

// Order returns the order of the series
func (_cgde ScatterChartSeries) Order() uint32 { return _cgde._egge.Order.ValAttr }
