package license

import (
	_ea "bytes"
	_e "compress/gzip"
	_g "crypto"
	_fg "crypto/aes"
	_da "crypto/cipher"
	_fe "crypto/rand"
	_ed "crypto/rsa"
	_ec "crypto/sha256"
	_a "crypto/sha512"
	_bed "crypto/x509"
	_cb "encoding/base64"
	_eaf "encoding/binary"
	_cae "encoding/hex"
	_be "encoding/json"
	_cbb "encoding/pem"
	_dg "errors"
	_dgd "fmt"
	_ca "io"
	_bc "io/ioutil"
	_eg "log"
	_fc "net"
	_cg "net/http"
	_b "os"
	_ga "path/filepath"
	_d "regexp"
	_f "sort"
	_dge "strings"
	_cd "sync"
	_gf "time"

	_cgg "gitee.com/greatmusicians/unioffice"
	_cda "gitee.com/greatmusicians/unioffice/common"
)

func _edd(_bgcc []byte) (_ca.Reader, error) {
	_gca := new(_ea.Buffer)
	_dcf := _e.NewWriter(_gca)
	_dcf.Write(_bgcc)
	_fad := _dcf.Close()
	if _fad != nil {
		return nil, _fad
	}
	return _gca, nil
}
func (_ge *LicenseKey) TypeToString() string {
	if _ge._fec {
		return "Metered\u0020subscription"
	}
	if _ge.Tier == LicenseTierUnlicensed {
		return "Unlicensed"
	}
	if _ge.Tier == LicenseTierCommunity {
		return "AGPLv3\u0020Open\u0020Source\u0020Community\u0020License"
	}
	if _ge.Tier == LicenseTierIndividual || _ge.Tier == "indie" {
		return "Commercial\u0020License\u0020\u002d\u0020Individual"
	}
	return "Commercial\u0020License\u0020\u002d\u0020Business"
}
func (_eec defaultStateHolder) loadState(_bdd string) (reportState, error) {
	_cdae := _aab()
	if len(_cdae) == 0 {
		return reportState{}, _dg.New("home\u0020dir\u0020not\u0020set")
	}
	_agc := _ga.Join(_cdae, "\u002eunidoc")
	_ega := _b.MkdirAll(_agc, 0777)
	if _ega != nil {
		return reportState{}, _ega
	}
	if len(_bdd) < 20 {
		return reportState{}, _dg.New("invalid\u0020key")
	}
	_cdc := []byte(_bdd)
	_fdg := _a.Sum512_256(_cdc[:20])
	_gfb := _cae.EncodeToString(_fdg[:])
	_gb := _ga.Join(_agc, _gfb)
	_ffd, _ega := _bc.ReadFile(_gb)
	if _ega != nil {
		if _b.IsNotExist(_ega) {
			return reportState{}, nil
		}
		_cgg.Log("ERROR:\u0020\u0025v\u000a", _ega)
		return reportState{}, _dg.New("invalid\u0020data")
	}
	const _dbeb = "ha9NK8]RbL\u002am4LKW"
	_ffd, _ega = _eac([]byte(_dbeb), _ffd)
	if _ega != nil {
		return reportState{}, _ega
	}
	var _ecg reportState
	_ega = _be.Unmarshal(_ffd, &_ecg)
	if _ega != nil {
		_cgg.Log("ERROR:\u0020Invalid\u0020data:\u0020\u0025v\u000a", _ega)
		return reportState{}, _dg.New("invalid\u0020data")
	}
	return _ecg, nil
}

const _ac = "305c300d06092a864886f70d0101010500034b003048024100b87eafb6c07499eb97cc9d3565ecf3168196301907c841addc665086bb3ed8eb12d9da26cafa96450146da8bd0ccf155fcacc686955ef0302fa44aa3ec89417b0203010001"

var _bgb = MakeUnlicensedKey()
var _dadc map[string]struct{}
var _gg = _gf.Date(2010, 1, 1, 0, 0, 0, 0, _gf.UTC)

func GetMeteredState() (MeteredStatus, error) {
	if _bgb == nil {
		return MeteredStatus{}, _dg.New("license\u0020key\u0020not\u0020set")
	}
	if !_bgb._fec || len(_bgb._fb) == 0 {
		return MeteredStatus{}, _dg.New("api key\u0020not\u0020set")
	}
	_egg, _dce := _bfa.loadState(_bgb._fb)
	if _dce != nil {
		_cgg.Log("ERROR:\u0020\u0025v\u000a", _dce)
		return MeteredStatus{}, _dce
	}
	if _egg.Docs > 0 {
		_eade := _adc("", "", true)
		if _eade != nil {
			return MeteredStatus{}, _eade
		}
	}
	_fga.Lock()
	defer _fga.Unlock()
	_gdee := _fca()
	_gdee._aafe = _bgb._fb
	_aff, _dce := _gdee.getStatus()
	if _dce != nil {
		return MeteredStatus{}, _dce
	}
	if !_aff.Valid {
		return MeteredStatus{}, _dg.New("key\u0020not\u0020valid")
	}
	_dbe := MeteredStatus{OK: true, Credits: _aff.OrgCredits, Used: _aff.OrgUsed}
	return _dbe, nil
}

const (
	_cf  = "\u002d\u002d\u002d--BEGIN\u0020UNIDOC\u0020LICENSE\u0020KEY\u002d\u002d\u002d\u002d\u002d"
	_eda = "\u002d\u002d\u002d\u002d\u002dEND\u0020UNIDOC LICENSE\u0020KEY\u002d\u002d\u002d\u002d\u002d"
)

func _ad(_bd string, _bg []byte) (string, error) {
	_edc, _ := _cbb.Decode([]byte(_bd))
	if _edc == nil {
		return "", _dgd.Errorf("PrivKey\u0020failed")
	}
	_ab, _cc := _bed.ParsePKCS1PrivateKey(_edc.Bytes)
	if _cc != nil {
		return "", _cc
	}
	_fd := _a.New()
	_fd.Write(_bg)
	_fcc := _fd.Sum(nil)
	_bf, _cc := _ed.SignPKCS1v15(_fe.Reader, _ab, _g.SHA512, _fcc)
	if _cc != nil {
		return "", _cc
	}
	_abd := _cb.StdEncoding.EncodeToString(_bg)
	_abd += "\u000a\u002b\u000a"
	_abd += _cb.StdEncoding.EncodeToString(_bf)
	return _abd, nil
}

type LegacyLicense struct {
	Name        string
	Signature   string `json:",omitempty"`
	Expiration  _gf.Time
	LicenseType LegacyLicenseType
}

func _dc(_dga string) (LicenseKey, error) {
	var _dd LicenseKey
	_dad, _cdb := _cac(_cf, _eda, _dga)
	if _cdb != nil {
		return _dd, _cdb
	}
	_gad, _cdb := _ba(_bcg, _dad)
	if _cdb != nil {
		return _dd, _cdb
	}
	_cdb = _be.Unmarshal(_gad, &_dd)
	if _cdb != nil {
		return _dd, _cdb
	}
	_dd.CreatedAt = _gf.Unix(_dd.CreatedAtInt, 0)
	if _dd.ExpiresAtInt > 0 {
		_ddg := _gf.Unix(_dd.ExpiresAtInt, 0)
		_dd.ExpiresAt = _ddg
	}
	return _dd, nil
}
func _aegg(_cfe *_cg.Response) (_ca.ReadCloser, error) {
	var _gdc error
	var _daad _ca.ReadCloser
	switch _dge.ToLower(_cfe.Header.Get("Content\u002dEncoding")) {
	case "gzip":
		_daad, _gdc = _e.NewReader(_cfe.Body)
		if _gdc != nil {
			return _daad, _gdc
		}
		defer _daad.Close()
	default:
		_daad = _cfe.Body
	}
	return _daad, nil
}
func _ddea(_bgcf *_cg.Response) ([]byte, error) {
	var _dbg []byte
	_ccb, _abeg := _aegg(_bgcf)
	if _abeg != nil {
		return _dbg, _abeg
	}
	return _bc.ReadAll(_ccb)
}
func _eac(_babe, _dae []byte) ([]byte, error) {
	_add := make([]byte, _cb.URLEncoding.DecodedLen(len(_dae)))
	_eed, _cggd := _cb.URLEncoding.Decode(_add, _dae)
	if _cggd != nil {
		return nil, _cggd
	}
	_add = _add[:_eed]
	_def, _cggd := _fg.NewCipher(_babe)
	if _cggd != nil {
		return nil, _cggd
	}
	if len(_add) < _fg.BlockSize {
		return nil, _dg.New("ciphertext\u0020too\u0020short")
	}
	_bdcd := _add[:_fg.BlockSize]
	_add = _add[_fg.BlockSize:]
	_defa := _da.NewCFBDecrypter(_def, _bdcd)
	_defa.XORKeyStream(_add, _add)
	return _add, nil
}
func SetMeteredKey(apiKey string) error {
	if len(apiKey) == 0 {
		_cgg.Log("Metered\u0020License\u0020API\u0020Key\u0020must\u0020not\u0020be\u0020empty\u000a")
		_cgg.Log("\u002d\u0020Grab\u0020one\u0020in\u0020the\u0020Free\u0020Tier at\u0020https:\u002f\u002fcloud\u002eunidoc\u002eio\u000a")
		return _dgd.Errorf("metered\u0020license\u0020api\u0020key\u0020must\u0020not\u0020be\u0020empty:\u0020create one\u0020at\u0020https:\u002f\u002fcloud\u002eunidoc.io")
	}
	if _bgb != nil && (_bgb._fec || _bgb.Tier != LicenseTierUnlicensed) {
		_cgg.Log("ERROR:\u0020Cannot\u0020set\u0020license\u0020key twice\u0020\u002d\u0020Should just\u0020initialize\u0020once\u000a")
		return _dg.New("license\u0020key\u0020already\u0020set")
	}
	_cgf := _fca()
	_cgf._aafe = apiKey
	_fgd, _fef := _cgf.getStatus()
	if _fef != nil {
		return _fef
	}
	if !_fgd.Valid {
		return _dg.New("key\u0020not\u0020valid")
	}
	_bdg := &LicenseKey{_fec: true, _fb: apiKey}
	_bgb = _bdg
	return nil
}

type LicenseKey struct {
	LicenseId    string   `json:"license_id"`
	CustomerId   string   `json:"customer_id"`
	CustomerName string   `json:"customer_name"`
	Tier         string   `json:"tier"`
	CreatedAt    _gf.Time `json:"-"`
	CreatedAtInt int64    `json:"created_at"`
	ExpiresAt    _gf.Time `json:"-"`
	ExpiresAtInt int64    `json:"expires_at"`
	CreatedBy    string   `json:"created_by"`
	CreatorName  string   `json:"creator_name"`
	CreatorEmail string   `json:"creator_email"`
	UniPDF       bool     `json:"unipdf"`
	UniOffice    bool     `json:"unioffice"`
	UniHTML      bool     `json:"unihtml"`
	Trial        bool     `json:"trial"`
	_fec         bool
	_fb          string
}

func _adc(_eag string, _eefc string, _cca bool) error {
	if _bgb == nil {
		return _dg.New("no\u0020license\u0020key")
	}
	if !_bgb._fec || len(_bgb._fb) == 0 {
		return nil
	}
	if len(_eag) == 0 && !_cca {
		return _dg.New("docKey\u0020not\u0020set")
	}
	_fga.Lock()
	defer _fga.Unlock()
	if _dadc == nil {
		_dadc = map[string]struct{}{}
	}
	if _ace == nil {
		_ace = map[string]int{}
	}
	_ece := 0
	if !_cca {
		_, _faf := _dadc[_eag]
		if !_faf {
			_dadc[_eag] = struct{}{}
			_ece++
		}
		if _ece == 0 {
			return nil
		}
		_ace[_eefc]++
	}
	_abea := _gf.Now()
	_egd, _adce := _bfa.loadState(_bgb._fb)
	if _adce != nil {
		_cgg.Log("ERROR:\u0020\u0025v\u000a", _adce)
		return _adce
	}
	if _egd.Usage == nil {
		_egd.Usage = map[string]int{}
	}
	for _dgb, _edac := range _ace {
		_egd.Usage[_dgb] += _edac
	}
	_ace = nil
	const _fccf = 24 * _gf.Hour
	const _ceb = 3 * 24 * _gf.Hour
	if len(_egd.Instance) == 0 || _abea.Sub(_egd.LastReported) > _fccf || (_egd.LimitDocs && _egd.RemainingDocs <= _egd.Docs+int64(_ece)) || _cca {
		_ged, _dgdg := _b.Hostname()
		if _dgdg != nil {
			return _dgdg
		}
		_bde := _egd.Docs
		_bege, _baab, _dgdg := _gec()
		if _dgdg != nil {
			return _dgdg
		}
		_f.Strings(_baab)
		_f.Strings(_bege)
		_caf, _dgdg := _bcbd()
		if _dgdg != nil {
			return _dgdg
		}
		_bbe := false
		for _, _ddc := range _baab {
			if _ddc == _caf.String() {
				_bbe = true
			}
		}
		if !_bbe {
			_baab = append(_baab, _caf.String())
		}
		_caa := _fca()
		_caa._aafe = _bgb._fb
		_bde += int64(_ece)
		_bdff := meteredUsageCheckinForm{Instance: _egd.Instance, Next: _egd.Next, UsageNumber: int(_bde), NumFailed: _egd.NumErrors, Hostname: _ged, LocalIP: _dge.Join(_baab, "\u002c\u0020"), MacAddress: _dge.Join(_bege, "\u002c\u0020"), Package: "unioffice", PackageVersion: _cda.Version, Usage: _egd.Usage}
		if len(_bege) == 0 {
			_bdff.MacAddress = "none"
		}
		_cbg := int64(0)
		_fdb := _egd.NumErrors
		_gdg := _abea
		_cfb := 0
		_geb := _egd.LimitDocs
		_ecbe, _dgdg := _caa.checkinUsage(_bdff)
		if _dgdg != nil {
			if _abea.Sub(_egd.LastReported) > _ceb {
				return _dg.New("too\u0020long\u0020since\u0020last\u0020successful checkin")
			}
			_cbg = _bde
			_fdb++
			_gdg = _egd.LastReported
		} else {
			_geb = _ecbe.LimitDocs
			_cfb = _ecbe.RemainingDocs
			_fdb = 0
		}
		if len(_ecbe.Instance) == 0 {
			_ecbe.Instance = _bdff.Instance
		}
		if len(_ecbe.Next) == 0 {
			_ecbe.Next = _bdff.Next
		}
		_dgdg = _bfa.updateState(_caa._aafe, _ecbe.Instance, _ecbe.Next, int(_cbg), _geb, _cfb, int(_fdb), _gdg, nil)
		if _dgdg != nil {
			return _dgdg
		}
		if !_ecbe.Success {
			return _dgd.Errorf("error:\u0020\u0025s", _ecbe.Message)
		}
	} else {
		_adce = _bfa.updateState(_bgb._fb, _egd.Instance, _egd.Next, int(_egd.Docs)+_ece, _egd.LimitDocs, int(_egd.RemainingDocs), int(_egd.NumErrors), _egd.LastReported, _egd.Usage)
		if _adce != nil {
			return _adce
		}
	}
	return nil
}

var _bfa stateLoader = defaultStateHolder{}

type meteredUsageCheckinResp struct {
	Instance      string `json:"inst"`
	Next          string `json:"next"`
	Success       bool   `json:"success"`
	Message       string `json:"message"`
	RemainingDocs int    `json:"rd"`
	LimitDocs     bool   `json:"ld"`
}

func init() {
	_deb, _cef := _cae.DecodeString(_ac)
	if _cef != nil {
		_eg.Fatalf("error reading\u0020key:\u0020\u0025s", _cef)
	}
	_cegf, _cef := _bed.ParsePKIXPublicKey(_deb)
	if _cef != nil {
		_eg.Fatalf("error reading\u0020key:\u0020\u0025s", _cef)
	}
	_bgaf = _cegf.(*_ed.PublicKey)
}
func (_bfb defaultStateHolder) updateState(_faa, _dadf, _ded string, _eba int, _aaa bool, _ecd int, _ebac int, _bbf _gf.Time, _aeg map[string]int) error {
	_beg := _aab()
	if len(_beg) == 0 {
		return _dg.New("home\u0020dir\u0020not\u0020set")
	}
	_abe := _ga.Join(_beg, "\u002eunidoc")
	_cgb := _b.MkdirAll(_abe, 0777)
	if _cgb != nil {
		return _cgb
	}
	if len(_faa) < 20 {
		return _dg.New("invalid\u0020key")
	}
	_dgf := []byte(_faa)
	_ccgf := _a.Sum512_256(_dgf[:20])
	_bgg := _cae.EncodeToString(_ccgf[:])
	_cab := _ga.Join(_abe, _bgg)
	var _ecc reportState
	_ecc.Docs = int64(_eba)
	_ecc.NumErrors = int64(_ebac)
	_ecc.LimitDocs = _aaa
	_ecc.RemainingDocs = int64(_ecd)
	_ecc.LastWritten = _gf.Now().UTC()
	_ecc.LastReported = _bbf
	_ecc.Instance = _dadf
	_ecc.Next = _ded
	_ecc.Usage = _aeg
	_debe, _cgb := _be.Marshal(_ecc)
	if _cgb != nil {
		return _cgb
	}
	const _agb = "ha9NK8]RbL\u002am4LKW"
	_debe, _cgb = _bab([]byte(_agb), _debe)
	if _cgb != nil {
		return _cgb
	}
	_cgb = _bc.WriteFile(_cab, _debe, 0600)
	if _cgb != nil {
		return _cgb
	}
	return nil
}

const _dfc = "UNIOFFICE_LICENSE_PATH"

func _ba(_gfd string, _aa string) ([]byte, error) {
	var (
		_dab int
		_bcf string
	)
	for _, _bcf = range []string{"\u000a\u002b\u000a", "\u000d\u000a\u002b\r\u000a", "\u0020\u002b\u0020"} {
		if _dab = _dge.Index(_aa, _bcf); _dab != -1 {
			break
		}
	}
	if _dab == -1 {
		return nil, _dgd.Errorf("invalid\u0020input\u002c signature\u0020separator")
	}
	_de := _aa[:_dab]
	_ce := _dab + len(_bcf)
	_bdc := _aa[_ce:]
	if _de == "" || _bdc == "" {
		return nil, _dgd.Errorf("invalid\u0020input,\u0020missing\u0020original or\u0020signature")
	}
	_fge, _cea := _cb.StdEncoding.DecodeString(_de)
	if _cea != nil {
		return nil, _dgd.Errorf("invalid\u0020input\u0020original")
	}
	_fea, _cea := _cb.StdEncoding.DecodeString(_bdc)
	if _cea != nil {
		return nil, _dgd.Errorf("invalid\u0020input\u0020signature")
	}
	_af, _ := _cbb.Decode([]byte(_gfd))
	if _af == nil {
		return nil, _dgd.Errorf("PubKey\u0020failed")
	}
	_ff, _cea := _bed.ParsePKIXPublicKey(_af.Bytes)
	if _cea != nil {
		return nil, _cea
	}
	_daf := _ff.(*_ed.PublicKey)
	if _daf == nil {
		return nil, _dgd.Errorf("PubKey\u0020conversion\u0020failed")
	}
	_fgf := _a.New()
	_fgf.Write(_fge)
	_ffg := _fgf.Sum(nil)
	_cea = _ed.VerifyPKCS1v15(_daf, _g.SHA512, _ffg, _fea)
	if _cea != nil {
		return nil, _cea
	}
	return _fge, nil
}
func (_dde *LicenseKey) ToString() string {
	if _dde._fec {
		return "Metered\u0020subscription"
	}
	_ae := _dgd.Sprintf("License\u0020Id:\u0020\u0025s\u000a", _dde.LicenseId)
	_ae += _dgd.Sprintf("Customer\u0020Id:\u0020\u0025s\u000a", _dde.CustomerId)
	_ae += _dgd.Sprintf("Customer\u0020Name:\u0020\u0025s\u000a", _dde.CustomerName)
	_ae += _dgd.Sprintf("Tier:\u0020\u0025s\n", _dde.Tier)
	_ae += _dgd.Sprintf("Created\u0020At:\u0020\u0025s\u000a", _cda.UtcTimeFormat(_dde.CreatedAt))
	if _dde.ExpiresAt.IsZero() {
		_ae += "Expires\u0020At:\u0020Never\u000a"
	} else {
		_ae += _dgd.Sprintf("Expires\u0020At:\u0020\u0025s\u000a", _cda.UtcTimeFormat(_dde.ExpiresAt))
	}
	_ae += _dgd.Sprintf("Creator:\u0020\u0025s\u0020<\u0025s\u003e\u000a", _dde.CreatorName, _dde.CreatorEmail)
	return _ae
}
func (_eaa *meteredClient) checkinUsage(_gde meteredUsageCheckinForm) (meteredUsageCheckinResp, error) {
	_gde.Package = "unioffice"
	_gde.PackageVersion = _cda.Version
	var _bdf meteredUsageCheckinResp
	_bcbg := _eaa._ccg + "\u002fmetered\u002fusage_checkin"
	_cdf, _fac := _be.Marshal(_gde)
	if _fac != nil {
		return _bdf, _fac
	}
	_egf, _fac := _edd(_cdf)
	if _fac != nil {
		return _bdf, _fac
	}
	_ceaa, _fac := _cg.NewRequest("POST", _bcbg, _egf)
	if _fac != nil {
		return _bdf, _fac
	}
	_ceaa.Header.Add("Content\u002dType", "application\u002fjson")
	_ceaa.Header.Add("Content\u002dEncoding", "gzip")
	_ceaa.Header.Add("Accept\u002dEncoding", "gzip")
	_ceaa.Header.Add("X-API\u002dKEY", _eaa._aafe)
	_bgf, _fac := _eaa._afb.Do(_ceaa)
	if _fac != nil {
		return _bdf, _fac
	}
	defer _bgf.Body.Close()
	if _bgf.StatusCode != 200 {
		return _bdf, _dgd.Errorf("failed\u0020to\u0020checkin\u002c\u0020status\u0020code\u0020is:\u0020\u0025d", _bgf.StatusCode)
	}
	_gfa, _fac := _ddea(_bgf)
	if _fac != nil {
		return _bdf, _fac
	}
	_fac = _be.Unmarshal(_gfa, &_bdf)
	if _fac != nil {
		return _bdf, _fac
	}
	return _bdf, nil
}
func (_ggf *LicenseKey) Validate() error {
	if _ggf._fec {
		return nil
	}
	if len(_ggf.LicenseId) < 10 {
		return _dgd.Errorf("invalid license:\u0020License\u0020Id")
	}
	if len(_ggf.CustomerId) < 10 {
		return _dgd.Errorf("invalid\u0020license:\u0020Customer Id")
	}
	if len(_ggf.CustomerName) < 1 {
		return _dgd.Errorf("invalid\u0020license:\u0020Customer\u0020Name")
	}
	if _gg.After(_ggf.CreatedAt) {
		return _dgd.Errorf("invalid\u0020license:\u0020Created At\u0020is invalid")
	}
	if _ggf.ExpiresAt.IsZero() {
		_fda := _ggf.CreatedAt.AddDate(1, 0, 0)
		if _aacd.After(_fda) {
			_fda = _aacd
		}
		_ggf.ExpiresAt = _fda
	}
	if _ggf.CreatedAt.After(_ggf.ExpiresAt) {
		return _dgd.Errorf("invalid\u0020license:\u0020Created\u0020At cannot be Greater\u0020than\u0020Expires\u0020At")
	}
	if _ggf.isExpired() {
		return _dgd.Errorf("invalid\u0020license:\u0020The license\u0020has\u0020already\u0020expired")
	}
	if len(_ggf.CreatorName) < 1 {
		return _dgd.Errorf("invalid\u0020license:\u0020Creator\u0020name")
	}
	if len(_ggf.CreatorEmail) < 1 {
		return _dgd.Errorf("invalid\u0020license:\u0020Creator\u0020email")
	}
	if _ggf.CreatedAt.After(_bga) {
		if !_ggf.UniOffice {
			return _dgd.Errorf("invalid\u0020license:\u0020This\u0020UniDoc\u0020key\u0020is\u0020invalid\u0020for\u0020UniOffice")
		}
	}
	return nil
}

type reportState struct {
	Instance      string         `json:"inst"`
	Next          string         `json:"n"`
	Docs          int64          `json:"d"`
	NumErrors     int64          `json:"e"`
	LimitDocs     bool           `json:"ld"`
	RemainingDocs int64          `json:"rd"`
	LastReported  _gf.Time       `json:"lr"`
	LastWritten   _gf.Time       `json:"lw"`
	Usage         map[string]int `json:"u"`
}

func GenRefId(prefix string) (string, error) {
	var _bcff _ea.Buffer
	_bcff.WriteString(prefix)
	_geg := make([]byte, 8+16)
	_gccg := _gf.Now().UTC().UnixNano()
	_eaf.BigEndian.PutUint64(_geg, uint64(_gccg))
	_, _aae := _fe.Read(_geg[8:])
	if _aae != nil {
		return "", _aae
	}
	_bcff.WriteString(_cae.EncodeToString(_geg))
	return _bcff.String(), nil
}

var _ace map[string]int

func (_afg LegacyLicense) Verify(pubKey *_ed.PublicKey) error {
	_aed := _afg
	_aed.Signature = ""
	_bgc := _ea.Buffer{}
	_gd := _be.NewEncoder(&_bgc)
	if _bcb := _gd.Encode(_aed); _bcb != nil {
		return _bcb
	}
	_ee, _dadd := _cae.DecodeString(_afg.Signature)
	if _dadd != nil {
		return _dadd
	}
	_adg := _ec.Sum256(_bgc.Bytes())
	_dadd = _ed.VerifyPKCS1v15(pubKey, _g.SHA256, _adg[:], _ee)
	return _dadd
}

type stateLoader interface {
	loadState(_bb string) (reportState, error)
	updateState(_agd, _bfd, _baa string, _egfa int, _ccf bool, _ade int, _gcc int, _bac _gf.Time, _eefe map[string]int) error
}

var _fga = &_cd.Mutex{}

func _gec() ([]string, []string, error) {
	_gdfd, _gac := _fc.Interfaces()
	if _gac != nil {
		return nil, nil, _gac
	}
	var _cacc []string
	var _ebc []string
	for _, _bdb := range _gdfd {
		if _bdb.Flags&_fc.FlagUp == 0 || _ea.Equal(_bdb.HardwareAddr, nil) {
			continue
		}
		_afc, _cff := _bdb.Addrs()
		if _cff != nil {
			return nil, nil, _cff
		}
		_bgab := 0
		for _, _dea := range _afc {
			var _cag _fc.IP
			switch _bce := _dea.(type) {
			case *_fc.IPNet:
				_cag = _bce.IP
			case *_fc.IPAddr:
				_cag = _bce.IP
			}
			if _cag.IsLoopback() {
				continue
			}
			if _cag.To4() == nil {
				continue
			}
			_ebc = append(_ebc, _cag.String())
			_bgab++
		}
		_acc := _bdb.HardwareAddr.String()
		if _acc != "" && _bgab > 0 {
			_cacc = append(_cacc, _acc)
		}
	}
	return _cacc, _ebc, nil
}

type meteredStatusResp struct {
	Valid        bool  `json:"valid"`
	OrgCredits   int64 `json:"org_credits"`
	OrgUsed      int64 `json:"org_used"`
	OrgRemaining int64 `json:"org_remaining"`
}

var _cabc = false

func (_cfc *meteredClient) getStatus() (meteredStatusResp, error) {
	var _aag meteredStatusResp
	_ag := _cfc._ccg + "\u002fmetered\u002fstatus"
	var _dda meteredStatusForm
	_gdf, _cgge := _be.Marshal(_dda)
	if _cgge != nil {
		return _aag, _cgge
	}
	_gee, _cgge := _edd(_gdf)
	if _cgge != nil {
		return _aag, _cgge
	}
	_cbe, _cgge := _cg.NewRequest("POST", _ag, _gee)
	if _cgge != nil {
		return _aag, _cgge
	}
	_cbe.Header.Add("Content\u002dType", "application\u002fjson")
	_cbe.Header.Add("Content\u002dEncoding", "gzip")
	_cbe.Header.Add("Accept\u002dEncoding", "gzip")
	_cbe.Header.Add("X-API\u002dKEY", _cfc._aafe)
	_eef, _cgge := _cfc._afb.Do(_cbe)
	if _cgge != nil {
		return _aag, _cgge
	}
	defer _eef.Body.Close()
	if _eef.StatusCode != 200 {
		return _aag, _dgd.Errorf("failed\u0020to\u0020checkin\u002c\u0020status\u0020code\u0020is:\u0020\u0025d", _eef.StatusCode)
	}
	_df, _cgge := _ddea(_eef)
	if _cgge != nil {
		return _aag, _cgge
	}
	_cgge = _be.Unmarshal(_df, &_aag)
	if _cgge != nil {
		return _aag, _cgge
	}
	return _aag, nil
}

type meteredUsageCheckinForm struct {
	Instance       string         `json:"inst"`
	Next           string         `json:"next"`
	UsageNumber    int            `json:"usage_number"`
	NumFailed      int64          `json:"num_failed"`
	Hostname       string         `json:"hostname"`
	LocalIP        string         `json:"local_ip"`
	MacAddress     string         `json:"mac_address"`
	Package        string         `json:"package"`
	PackageVersion string         `json:"package_version"`
	Usage          map[string]int `json:"u"`
}

const (
	LicenseTierUnlicensed = "unlicensed"
	LicenseTierCommunity  = "community"
	LicenseTierIndividual = "individual"
	LicenseTierBusiness   = "business"
)

func (_fa *LicenseKey) getExpiryDateToCompare() _gf.Time {
	if _fa.Trial {
		return _gf.Now().UTC()
	}
	return _cda.ReleasedAt
}
func _bcbd() (_fc.IP, error) {
	_bdfd, _agcd := _fc.Dial("udp", "8\u002e8\u002e8\u002e8:80")
	if _agcd != nil {
		return nil, _agcd
	}
	defer _bdfd.Close()
	_fdc := _bdfd.LocalAddr().(*_fc.UDPAddr)
	return _fdc.IP, nil
}
func _cac(_ceg string, _ffc string, _aac string) (string, error) {
	_db := _dge.Index(_aac, _ceg)
	if _db == -1 {
		return "", _dgd.Errorf("header not\u0020found")
	}
	_eae := _dge.Index(_aac, _ffc)
	if _eae == -1 {
		return "", _dgd.Errorf("footer not\u0020found")
	}
	_gc := _db + len(_ceg) + 1
	return _aac[_gc : _eae-1], nil
}
func (_fed *LicenseKey) isExpired() bool { return _fed.getExpiryDateToCompare().After(_fed.ExpiresAt) }
func TrackUse(useKey string) {
	if _bgb == nil {
		return
	}
	if !_bgb._fec || len(_bgb._fb) == 0 {
		return
	}
	if len(useKey) == 0 {
		return
	}
	_fga.Lock()
	defer _fga.Unlock()
	if _ace == nil {
		_ace = map[string]int{}
	}
	_ace[useKey]++
}

type meteredStatusForm struct{}
type defaultStateHolder struct{}

const _bae = "UNIOFFICE_CUSTOMER_NAME"

func GetLicenseKey() *LicenseKey {
	if _bgb == nil {
		return nil
	}
	_faab := *_bgb
	return &_faab
}
func _fca() *meteredClient {
	_adgg := meteredClient{_ccg: "https:\u002f/cloud\u002eunidoc\u002eio/api", _afb: &_cg.Client{Timeout: 30 * _gf.Second}}
	if _cdac := _b.Getenv("UNIDOC_LICENSE_SERVER_URL"); _dge.HasPrefix(_cdac, "http") {
		_adgg._ccg = _cdac
	}
	return &_adgg
}
func SetLicenseKey(content string, customerName string) error {
	if _cabc {
		return nil
	}
	_gag, _fdd := _dc(content)
	if _fdd != nil {
		_cgg.Log("License\u0020code\u0020decode\u0020error:\u0020\u0025v\u000a", _fdd)
		return _fdd
	}
	if !_dge.EqualFold(_gag.CustomerName, customerName) {
		_cgg.Log("License\u0020code\u0020issue\u0020\u002d Customer\u0020name\u0020mismatch\u002c\u0020expected\u0020\u0027\u0025s\u0027,\u0020but\u0020got \u0027\u0025s\u0027\u000a", customerName, _gag.CustomerName)
		return _dgd.Errorf("customer\u0020name\u0020mismatch\u002c\u0020expected\u0020\u0027\u0025s\u0027\u002c\u0020but\u0020got\u0020\u0027\u0025s'", customerName, _gag.CustomerName)
	}
	_fdd = _gag.Validate()
	if _fdd != nil {
		_cgg.Log("License\u0020code\u0020validation\u0020error:\u0020\u0025v\u000a", _fdd)
		return _fdd
	}
	_bgb = &_gag
	return nil
}

type LegacyLicenseType byte

var _bga = _gf.Date(2019, 6, 6, 0, 0, 0, 0, _gf.UTC)

func MakeUnlicensedKey() *LicenseKey {
	_ddd := LicenseKey{}
	_ddd.CustomerName = "Unlicensed"
	_ddd.Tier = LicenseTierUnlicensed
	_ddd.CreatedAt = _gf.Now().UTC()
	_ddd.CreatedAtInt = _ddd.CreatedAt.Unix()
	return &_ddd
}
func _bab(_gce, _gfe []byte) ([]byte, error) {
	_bfc, _ef := _fg.NewCipher(_gce)
	if _ef != nil {
		return nil, _ef
	}
	_ffge := make([]byte, _fg.BlockSize+len(_gfe))
	_gbf := _ffge[:_fg.BlockSize]
	if _, _fdcd := _ca.ReadFull(_fe.Reader, _gbf); _fdcd != nil {
		return nil, _fdcd
	}
	_cga := _da.NewCFBEncrypter(_bfc, _gbf)
	_cga.XORKeyStream(_ffge[_fg.BlockSize:], _gfe)
	_cdaf := make([]byte, _cb.URLEncoding.EncodedLen(len(_ffge)))
	_cb.URLEncoding.Encode(_cdaf, _ffge)
	return _cdaf, nil
}
func init() {
	_caaa := _b.Getenv(_dfc)
	_bdbc := _b.Getenv(_bae)
	if len(_caaa) == 0 || len(_bdbc) == 0 {
		return
	}
	_cba, _bbc := _bc.ReadFile(_caaa)
	if _bbc != nil {
		_cgg.Log("Unable\u0020to\u0020read\u0020license\u0020code\u0020file:\u0020\u0025v\u000a", _bbc)
		return
	}
	_bbc = SetLicenseKey(string(_cba), _bdbc)
	if _bbc != nil {
		_cgg.Log("Unable\u0020to\u0020load\u0020license\u0020code:\u0020\u0025v\u000a", _bbc)
		return
	}
}
func Track(docKey string, useKey string) error { return _adc(docKey, useKey, false) }

var _bgaf *_ed.PublicKey

const _bcg = "\u000a\u002d\u002d\u002d\u002d\u002dBEGIN PUBLIC\u0020KEY\u002d\u002d\u002d\u002d\u002d\u000aMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAmFUiyd7b5XjpkP5Rap4w\u000aDc1dyzIQ4LekxrvytnEMpNUbo6iA74V8ruZOvrScsf2QeN9\u002fqrUG8qEbUWdoEYq+\u000aotFNAFNxlGbxbDHcdGVaM0OXdXgDyL5aIEagL0c5pwjIdPGIn46f78eMJ\u002bJkdcpD\nDJaqYXdrz5KeshjSiIaa7menBIAXS4UFxNfHhN0HCYZYqQG7bK+s5rRHonydNWEG\u000aH8Myvr2pya2KrMumfmAxUB6fenC\u002f4O0Wr8gfPOU8RitmbDvQPIRXOL4vTBrBdbaA\u000a9nwNP\u002bi\u002f\u002f20MT2bxmeWB\u002bgpcEhGpXZ733azQxrC3J4v3CZmENStDK\u002fKDSPKUGfu6\u000afwIDAQAB\u000a\u002d\u002d\u002d\u002d\u002dEND\u0020PUBLIC KEY\u002d\u002d\u002d\u002d\u002d\n"

type meteredClient struct {
	_ccg  string
	_aafe string
	_afb  *_cg.Client
}

var _aacd = _gf.Date(2020, 1, 1, 0, 0, 0, 0, _gf.UTC)

func _aab() string {
	_dded := _b.Getenv("HOME")
	if len(_dded) == 0 {
		_dded, _ = _b.UserHomeDir()
	}
	return _dded
}
func SetLegacyLicenseKey(s string) error {
	_cabb := _d.MustCompile("\u005cs")
	s = _cabb.ReplaceAllString(s, "")
	var _fafc _ca.Reader
	_fafc = _dge.NewReader(s)
	_fafc = _cb.NewDecoder(_cb.RawURLEncoding, _fafc)
	_fafc, _gfad := _e.NewReader(_fafc)
	if _gfad != nil {
		return _gfad
	}
	_cfd := _be.NewDecoder(_fafc)
	_gaa := &LegacyLicense{}
	if _egc := _cfd.Decode(_gaa); _egc != nil {
		return _egc
	}
	if _cfbc := _gaa.Verify(_bgaf); _cfbc != nil {
		return _dg.New("license\u0020validatin\u0020error")
	}
	if _gaa.Expiration.Before(_cda.ReleasedAt) {
		return _dg.New("license\u0020expired")
	}
	_fdga := _gf.Now().UTC()
	_gfbb := LicenseKey{}
	_gfbb.CreatedAt = _fdga
	_gfbb.CustomerId = "Legacy"
	_gfbb.CustomerName = _gaa.Name
	_gfbb.Tier = LicenseTierBusiness
	_gfbb.ExpiresAt = _gaa.Expiration
	_gfbb.CreatorName = "UniDoc\u0020support"
	_gfbb.CreatorEmail = "support\u0040unidoc\u002eio"
	_gfbb.UniOffice = true
	_bgb = &_gfbb
	return nil
}
func (_daa *LicenseKey) IsLicensed() bool {
	if _daa == nil {
		return false
	}
	return _daa.Tier != LicenseTierUnlicensed || _daa._fec
}

type MeteredStatus struct {
	OK      bool
	Credits int64
	Used    int64
}
