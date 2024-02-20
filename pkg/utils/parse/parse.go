package parse

import "strconv"

func ParseBool(val string, defVal bool) bool {
	_val, err := strconv.ParseBool(val)
	if err != nil {
		return defVal
	}
	return _val
}

func ParseInt(val string, defVal int) int {
	_val, err := strconv.Atoi(val)
	if err != nil {
		return defVal
	}
	return _val
}

func ParseInt64(val string, defVal int64) int64 {
	_val, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return defVal
	}
	return _val
}

func ParseFloat64(val string, defVal float64) float64 {
	_val, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return defVal
	}
	return _val
}

func ParseUint(val string, defVal uint) uint {
	_val, err := strconv.ParseUint(val, 10, 0)
	if err != nil {
		return defVal
	}
	return uint(_val)
}
