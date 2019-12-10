package wax

import (
	"math"
	"strconv"
	"strings"
)

func ParseNaN32WithPayload(s string) (float32, error) {
	s = strings.Trim(s, " \r\n\t")
	sign := uint32(0)
	if strings.HasPrefix(s, "-") {
		sign = 1
	}
	i := strings.Index(s, ":")
	if i < 0 {
		return qNaN32(sign), nil
	}

	payload, err := strconv.ParseUint(strings.TrimPrefix(s[i:], ":"), 0, 32)
	if err != nil {
		return 0, err
	}

	return sNaN32(sign, uint32(payload)), nil
}

func qNaN32(sign uint32) float32 {
	x := math.Float32bits(float32(math.NaN()))
	x &= 0x7fc00000
	x |= ((sign & 0x1) << 31)
	return math.Float32frombits(x)
}

func sNaN32(sign, payload uint32) float32 {
	x := math.Float32bits(float32(math.NaN()))
	x &= 0x7f800000
	x |= ((sign & 0x1) << 31)
	x |= (payload & 0x007fffff)
	return math.Float32frombits(x)
}

func ParseNaN64WithPayload(s string) (float64, error) {
	s = strings.Trim(s, " \r\n\t")
	sign := uint64(0)
	if strings.HasPrefix(s, "-") {
		sign = 1
	}
	i := strings.Index(s, ":")
	if i < 0 {
		return qNaN64(sign), nil
	}

	payload, err := strconv.ParseUint(strings.TrimPrefix(s[i:], ":"), 0, 64)
	if err != nil {
		return 0, err
	}

	return sNaN64(sign, payload), nil
}

func qNaN64(sign uint64) float64 {
	x := math.Float64bits(math.NaN())
	x &= 0x7ff8000000000000
	x |= ((sign & 0x1) << 63)
	return math.Float64frombits(x)
}

func sNaN64(sign, payload uint64) float64 {
	x := math.Float64bits(math.NaN())
	x &= 0x7ff0000000000000
	x |= ((sign & 0x1) << 63)
	x |= (payload & 0x000fffffffffffff)
	return math.Float64frombits(x)
}
