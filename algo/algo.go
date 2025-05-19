package algo

import "strconv"

func RepeatString(s string, cnt int) string {
	if cnt <= 0 {
		return ""
	}
	_dd := make([]byte, len(s)*cnt)
	_ab := []byte(s)
	for _gc := 0; _gc < cnt; _gc++ {
		copy(_dd[_gc:], _ab)
	}
	return string(_dd)
}
func _d(_bf byte) bool { return _bf >= '0' && _bf <= '9' }

// NaturalLess compares two strings in a human manner so rId2 sorts less than rId10
func NaturalLess(lhs, rhs string) bool {
	_bd, _gd := 0, 0
	for _bd < len(lhs) && _gd < len(rhs) {
		_a := lhs[_bd]
		_gb := rhs[_gd]
		_aa := _d(_a)
		_be := _d(_gb)
		switch {
		case _aa && !_be:
			return true
		case !_aa && _be:
			return false
		case !_aa && !_be:
			if _a != _gb {
				return _a < _gb
			}
			_bd++
			_gd++
		default:
			_ag := _bd + 1
			_c := _gd + 1
			for _ag < len(lhs) && _d(lhs[_ag]) {
				_ag++
			}
			for _c < len(rhs) && _d(rhs[_c]) {
				_c++
			}
			_e, _ := strconv.ParseUint(lhs[_bd:_ag], 10, 64)
			_bdc, _ := strconv.ParseUint(rhs[_bd:_c], 10, 64)
			if _e != _bdc {
				return _e < _bdc
			}
			_bd = _ag
			_gd = _c
		}
	}
	return len(lhs) < len(rhs)
}
