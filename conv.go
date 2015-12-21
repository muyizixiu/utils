// Package utils provides common functions
package utils

import (
	"strconv"
)

//String convert numbers into string on decimal form
func String(params interface{}) string {
	var param_i int64
	var param_u uint64
	switch params.(type) {
	case int8:
		param_i = int64(params.(int8))
		goto inttag
	case int16:
		param_i = int64(params.(int16))
		goto inttag
	case int32:
		param_i = int64(params.(int32))
		goto inttag
	case int64:
		param_i = int64(params.(int64))
		goto inttag
	case int:
		param_i = int64(params.(int))
		goto inttag
	case uint:
		param_u = uint64(params.(uint))
	case uint8:
		param_u = uint64(params.(uint8))
	case uint16:
		param_u = uint64(params.(uint16))
	case uint32:
		param_u = uint64(params.(uint32))
	case uint64:
		param_u = uint64(params.(uint64))
	default:
		return ""
	}
	return strconv.FormatUint(param_u, 10)
inttag:
	return strconv.FormatInt(param_i, 10)
}

//Number convert string into number
//@param flag indicate whether the number is unsigned or singed,
//by default it is unsigned
func Number(params string, flag ...bool) (interface{}, error) {
	if len(flag) != 0 && !flag[0] {
		return strconv.ParseUint(params, 10, 64)
	}
	return strconv.ParseInt(params, 10, 64)
}

//Int convert a string into a int type
//if errors happen ,it returns zero and the error
func Int(params string) (int, error) {
	i, err := strconv.ParseInt(params, 10, 64)
	if err != nil {
		return 0, err
	}
	return int(i), nil
}
